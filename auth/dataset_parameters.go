package auth

import (
	"fmt"
	"net/http"

	"github.com/ONSdigital/go-ns/common"
)

// Parameters is an interface that encapsulates the context of the authorisation check being applied.
// The purpose of the implementing struct is to define what information to extract from the incoming http.Request and to
// generate the correct outgoing http.Request to Permissions API. The fields required and the structure of the outgoing
// HTTP request is defined by the implementing struct.
type Parameters interface {
	CreateGetPermissionsRequest(host string) (*http.Request, error)
}

// UserDatasetParameters is an implementation of Parameters. Requires a user auth token, collection ID and dataset ID
// and checks a user is authorised to access the specified dataset in the specified collection.
type UserDatasetParameters struct {
	UserAuthToken string
	CollectionID  string
	DatasetID     string
}

// ServiceDatasetParameters is an implementation of Parameters. Requires a service account auth token and dataset ID.
// Checks a service account is authorised to access the specified dataset.
type ServiceDatasetParameters struct {
	ServiceAuthToken string
	DatasetID        string
}

// DatasetParameterFactory is an implementation of ParameterFactory. Creates Parameters for requesting user & service
// account permissions for a CMD dataset.
type DatasetParameterFactory struct {
	DatasetIDKey       string
	GetRequestVarsFunc func(r *http.Request) map[string]string
}

// CreateParameters fulfilling the ParameterFactory interface. Generates:
// 	- A UserDatasetParameters instance if the request contains a user auth token header.
// 	- Or a ServiceDatasetParameters instance if the request contains the service account auth token header.
// If neither header is present then returns noUserOrServiceAuthTokenProvidedError
func (f *DatasetParameterFactory) CreateParameters(req *http.Request) (Parameters, error) {
	userAuthToken := req.Header.Get(common.FlorenceHeaderKey)
	serviceAuthToken := req.Header.Get(common.AuthHeaderKey)
	collectionID := req.Header.Get(CollectionIDHeader)
	datasetID := f.GetRequestVarsFunc(req)[f.DatasetIDKey]

	if userAuthToken != "" {
		return newUserDatasetParameters(userAuthToken, collectionID, datasetID), nil
	}

	if serviceAuthToken != "" {
		return newServiceParameters(serviceAuthToken, datasetID), nil
	}

	return nil, noUserOrServiceAuthTokenProvidedError
}

func newUserDatasetParameters(userToken string, collectionID string, datasetID string) Parameters {
	return &UserDatasetParameters{
		UserAuthToken: userToken,
		CollectionID:  collectionID,
		DatasetID:     datasetID,
	}
}

func newServiceParameters(serviceToken string, datasetID string) Parameters {
	return &ServiceDatasetParameters{
		ServiceAuthToken: serviceToken,
		DatasetID:        datasetID,
	}
}

// CreateGetPermissionsRequest fulfilling the Parameters interface - creates a Permissions API request to get user
// dataset permissions.
func (params *UserDatasetParameters) CreateGetPermissionsRequest(host string) (*http.Request, error) {
	if host == "" {
		return nil, hostRequiredButEmptyError
	}

	url := fmt.Sprintf(userDatasetPermissionsURL, host, params.DatasetID, params.CollectionID)
	httpRequest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, Error{
			Cause:   err,
			Status:  500,
			Message: "error creating new get user dataset permissions http request",
		}
	}

	httpRequest.Header.Set(common.FlorenceHeaderKey, params.UserAuthToken)
	return httpRequest, nil
}

// CreateGetPermissionsRequest fulfilling the Parameters interface - creates a Permissions API request to get service
// account dataset permissions.
func (params *ServiceDatasetParameters) CreateGetPermissionsRequest(host string) (*http.Request, error) {
	if host == "" {
		return nil, hostRequiredButEmptyError
	}

	url := fmt.Sprintf(serviceDatasetPermissionsURL, host, params.DatasetID)
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, Error{
			Cause:   err,
			Status:  500,
			Message: "error creating new get service dataset permissions http request",
		}
	}

	r.Header.Set(common.AuthHeaderKey, params.ServiceAuthToken)
	return r, nil
}