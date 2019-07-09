package authv2

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/ONSdigital/go-ns/common"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateDatasetAuthorisationParameters(t *testing.T) {
	getRequestVars = func(r2 *http.Request) map[string]string {
		return map[string]string{datasetIDKey: datasetIDKey}
	}

	Convey("should return the expected UserDatasetParameters", t, func() {
		req, err := http.NewRequest("GET", "http://localhost:8080", nil)
		req.Header.Set(common.FlorenceHeaderKey, common.FlorenceHeaderKey)
		req.Header.Set(CollectionIDHeader, CollectionIDHeader)
		So(err, ShouldBeNil)

		getRequestVars = func(r2 *http.Request) map[string]string {
			return map[string]string{datasetIDKey: datasetIDKey}
		}

		actual, err := createDatasetAuthorisationParameters(req)
		So(err, ShouldBeNil)

		params, ok := actual.(*UserDatasetParameters)
		So(ok, ShouldBeTrue)

		expected := &UserDatasetParameters{UserToken: common.FlorenceHeaderKey, CollectionID: CollectionIDHeader, DatasetID: datasetIDKey}
		So(params, ShouldResemble, expected)
	})

	Convey("should return the expected ServiceDatasetParameters", t, func() {
		req, err := http.NewRequest("GET", "http://localhost:8080", nil)
		req.Header.Set(common.AuthHeaderKey, common.AuthHeaderKey)
		So(err, ShouldBeNil)

		actual, err := createDatasetAuthorisationParameters(req)
		So(err, ShouldBeNil)

		params, ok := actual.(*ServiceDatasetParameters)
		So(ok, ShouldBeTrue)

		expected := &ServiceDatasetParameters{ServiceToken: common.AuthHeaderKey, DatasetID: datasetIDKey}
		So(params, ShouldResemble, expected)
	})

	Convey("should return the expected error if both user and service tokens are nil", t, func() {
		req, err := http.NewRequest("GET", "http://localhost:8080", nil)
		So(err, ShouldBeNil)

		actual, err := createDatasetAuthorisationParameters(req)
		So(err, ShouldResemble, noUserOrServiceAuthTokenProvidedError)
		So(actual, ShouldBeNil)
	})
}

func TestUserDatasetParameters_NewGetDatasetPermissionsRequest(t *testing.T) {
	host := "http:localhost:666"

	Convey("should create expected get user dataset permissions request", t, func() {
		params := &UserDatasetParameters{
			UserToken:    common.FlorenceHeaderKey,
			CollectionID: CollectionIDHeader,
			DatasetID:    datasetIDKey,
		}

		req, err := params.NewGetDatasetPermissionsRequest(host)
		So(err, ShouldBeNil)

		So(req.URL.String(), ShouldResemble, fmt.Sprintf(userDatasetPermissionsURL, host, datasetIDKey, CollectionIDHeader))
		So(req.Header.Get(common.FlorenceHeaderKey), ShouldEqual, common.FlorenceHeaderKey)
	})

	Convey("should return expected error if fails to create request", t, func() {
		params := &UserDatasetParameters{
			UserToken:    common.FlorenceHeaderKey,
			CollectionID: CollectionIDHeader,
			DatasetID:    datasetIDKey,
		}

		req, err := params.NewGetDatasetPermissionsRequest("!@£$%^&*()_")
		So(req, ShouldBeNil)

		permErr, ok := err.(Error)
		So(ok, ShouldBeTrue)
		So(permErr.Status, ShouldEqual, 500)
		So(permErr.Message, ShouldEqual, "error creating new get user dataset permissions http request")
	})

	Convey("should return expected error if host empty", t, func() {
		params := &UserDatasetParameters{
			UserToken:    common.FlorenceHeaderKey,
			CollectionID: CollectionIDHeader,
			DatasetID:    datasetIDKey,
		}

		req, err := params.NewGetDatasetPermissionsRequest("")
		So(err, ShouldResemble, hostRequiredButEmptyError)
		So(req, ShouldBeNil)
	})

}

func TestServiceDatasetParameters_NewGetDatasetPermissionsRequest(t *testing.T) {
	host := "http:localhost:666"

	Convey("should create expected get service dataset permissions request", t, func() {
		params := &ServiceDatasetParameters{
			ServiceToken: common.AuthHeaderKey,
			DatasetID:    datasetIDKey,
		}

		req, err := params.NewGetDatasetPermissionsRequest(host)
		So(err, ShouldBeNil)

		So(req.URL.String(), ShouldResemble, fmt.Sprintf(serviceDatasetPermissionsURL, host, datasetIDKey))
		So(req.Header.Get(common.AuthHeaderKey), ShouldEqual, common.AuthHeaderKey)
	})

	Convey("should return expected error if host empty", t, func() {
		params := &ServiceDatasetParameters{
			ServiceToken:    common.AuthHeaderKey,
			DatasetID:    datasetIDKey,
		}

		req, err := params.NewGetDatasetPermissionsRequest("")
		So(err, ShouldResemble, hostRequiredButEmptyError)
		So(req, ShouldBeNil)
	})

	Convey("should return expected error if fails to create request", t, func() {
		params := &ServiceDatasetParameters{
			ServiceToken:    common.AuthHeaderKey,
			DatasetID:    datasetIDKey,
		}

		req, err := params.NewGetDatasetPermissionsRequest("!@£$%^&*()_")
		So(req, ShouldBeNil)

		permErr, ok := err.(Error)
		So(ok, ShouldBeTrue)
		So(permErr.Status, ShouldEqual, 500)
		So(permErr.Message, ShouldEqual, "error creating new get service dataset permissions http request")
	})
}
