package permissions

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/ONSdigital/dp-permissions/permissions/mocks"
	"github.com/ONSdigital/go-ns/common"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetErrorFromResponse(t *testing.T) {

	type scenario struct {
		desc          string
		body          []byte
		readerErr     error
		status        int
		assertErrFunc func(err Error) bool
	}

	scenarios := []scenario{
		{
			desc:      "Should return the expected error for a valid error entity response body",
			body:      toJson(t, errorEntity{"unauthorized"}),
			readerErr: nil,
			status:    401,
			assertErrFunc: func(err Error) bool {
				return reflect.DeepEqual(err, Error{Cause: nil, Status: 401, Message: "unauthorized"})
			},
		},
		{
			desc:      "Should return error with status 500 if read body returns an error",
			body:      nil,
			readerErr: errors.New("pop!"),
			status:    401,
			assertErrFunc: func(err Error) bool {
				return reflect.DeepEqual(err, Error{
					Cause:   errors.New("pop!"),
					Status:  500,
					Message: "internal server error failed reading get permissions error response body",
				})
			},
		},
		{
			desc:      "Should return error with status 500 if unmarshal body to error entity fails",
			body:      toJson(t, []int{1, 2, 3, 4, 5}),
			readerErr: nil,
			status:    401,
			assertErrFunc: func(err Error) bool {
				_, ok := err.Cause.(*json.UnmarshalTypeError)
				return ok &&
					err.Status == 500 &&
					err.Message == "internal server error failed unmarshalling get permissions error response body"
			},
		},
	}

	for i, s := range scenarios {
		Convey(fmt.Sprintf("%d) %s", i, s.desc), t, func() {
			resp := getMockErrorResponse(s.status, s.body, s.readerErr)

			err := getErrorFromResponse(resp)

			permErr, ok := err.(Error)
			So(ok, ShouldBeTrue)
			So(s.assertErrFunc(permErr), ShouldBeTrue)
		})
	}
}

func TestUnmarshalPermissions(t *testing.T) {

	type scenario struct {
		desc          string
		body          []byte
		err           error
		crud          *CRUD
		perms         callerPermissions
		assertErrFunc func(err error) bool
	}

	scenarios := []scenario{
		{
			desc: "should return expected error if read response body fails",
			body: nil,
			err:  errors.New("reader error"),
			crud: nil,
			assertErrFunc: func(err error) bool {
				permErr, ok := err.(Error);
				return ok &&
					permErr.Status == 500 &&
					permErr.Message == "internal server error failed reading get permissions error response body"
			},
		},
		{
			desc: "should return expected error if response body not valid permissions json",
			body: toJson(t, 666),
			crud: nil,
			assertErrFunc: func(err error) bool {
				if permErr, ok := err.(Error); ok {
					_, isJsonErr := permErr.Cause.(*json.UnmarshalTypeError)
					return isJsonErr &&
						permErr.Status == 500 &&
						permErr.Message == "internal server error failed marshalling response to permissions"

				}
				return false
			},
		},
		{
			desc: "should return CRUD for permissions json [Create, Read, Update,  Delete]",
			body: toJson(t, callerPermissions{List: []permission{Create, Read, Update, Delete}}),
			err:  nil,
			crud: &CRUD{Create: true, Read: true, Update: true, Delete: true},
			assertErrFunc: func(err error) bool {
				return err == nil
			},
		},
		{
			desc: "should return R for permissions json [Read]",
			body: toJson(t, callerPermissions{List: []permission{Read}}),
			err:  nil,
			crud: &CRUD{Create: false, Read: true, Update: false, Delete: false},
			assertErrFunc: func(err error) bool {
				return err == nil
			},
		},
		{
			desc: "should return expected error if caller has no permissions",
			body: toJson(t, callerPermissions{List: []permission{}}),
			err:  nil,
			crud: nil,
			assertErrFunc: func(err error) bool {
				permErr, ok := err.(Error)
				return ok && permErr.Status == 403 && permErr.Message == "forbidden"
			},
		},
	}

	for i, s := range scenarios {
		Convey(fmt.Sprintf("%d) %s", i, s.desc), t, func() {
			reader := &mocks.ReadCloser{
				GetEntityFunc: func() (i []byte, e error) {
					return s.body, s.err
				},
			}

			crud, err := unmarshalPermissions(reader)
			So(crud, ShouldResemble, s.crud)
			So(s.assertErrFunc(err), ShouldBeTrue)
		})
	}
}

func TestGetPermissionsRequest(t *testing.T) {

	type scenario struct {
		desc          string
		permissions   *Permissions
		serviceT      string
		userT         string
		collectionID  string
		datasetID     string
		AssertReqFunc func(r *http.Request) bool
		AssertErrFunc func(err error) bool
	}

	scenarios := []scenario{
		{
			desc:         "should return the expected error if the checker has not been configured with a host",
			permissions:  &Permissions{},
			serviceT:     "",
			userT:        "",
			collectionID: "",
			datasetID:    "",
			AssertReqFunc: func(r *http.Request) bool {
				return r == nil
			},
			AssertErrFunc: func(err error) bool {
				return reflect.DeepEqual(err, Error{
					Status:  500,
					Message: "error creating permissionsList request host not configured",
				})
			},
		},
		{
			desc:         "should return the expected request if the check is correctly configured",
			permissions:  &Permissions{host: "http://localhost:8082/permissionsList"},
			serviceT:     "111",
			userT:        "222",
			collectionID: "333",
			datasetID:    "444",
			AssertErrFunc: func(err error) bool {
				return err == nil
			},
			AssertReqFunc: func(r *http.Request) bool {
				return r != nil &&
					r.Header.Get(common.AuthHeaderKey) == "111" &&
					r.Header.Get(common.FlorenceHeaderKey) == "222" &&
					r.URL.Query().Get("collection_id") == "333" &&
					r.URL.Query().Get("dataset_id") == "444"
			},
		},
	}

	for i, s := range scenarios {
		Convey(fmt.Sprintf("%d) %s", i, s.desc), t, func() {
			r, err := s.permissions.getPermissionsRequest(s.serviceT, s.userT, s.collectionID, s.datasetID)
			So(s.AssertReqFunc(r), ShouldBeTrue)
			So(s.AssertErrFunc(err), ShouldBeTrue)
		})
	}
}

func getMockErrorResponse(status int, b []byte, err error) *http.Response {
	return &http.Response{
		StatusCode: status,
		Body: &mocks.ReadCloser{
			GetEntityFunc: func() ([]byte, error) {
				return b, err
			},
		},
	}
}

func toJson(t *testing.T, i interface{}) []byte {
	b, err := json.Marshal(i)
	if err != nil {
		t.Fatalf("failed to marshal object to json: %s", err.Error())
	}
	return b
}
