// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package auth

import (
	"context"
	"net/http"
	"sync"
)

var (
	lockClienterMockGetCallerPermissions sync.RWMutex
	lockClienterMockGetPermissions       sync.RWMutex
)

// ClienterMock is a mock implementation of Clienter.
//
//     func TestSomethingThatUsesClienter(t *testing.T) {
//
//         // make and configure a mocked Clienter
//         mockedClienter := &ClienterMock{
//             GetCallerPermissionsFunc: func(ctx context.Context, params Parameters) (*Permissions, error) {
// 	               panic("TODO: mock out the GetCallerPermissions method")
//             },
//             GetPermissionsFunc: func(ctx context.Context, getPermissionsRequest *http.Request) (*Permissions, error) {
// 	               panic("TODO: mock out the GetPermissions method")
//             },
//         }
//
//         // TODO: use mockedClienter in code that requires Clienter
//         //       and then make assertions.
//
//     }
type ClienterMock struct {
	// GetCallerPermissionsFunc mocks the GetCallerPermissions method.
	GetCallerPermissionsFunc func(ctx context.Context, params Parameters) (*Permissions, error)

	// GetPermissionsFunc mocks the GetPermissions method.
	GetPermissionsFunc func(ctx context.Context, getPermissionsRequest *http.Request) (*Permissions, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetCallerPermissions holds details about calls to the GetCallerPermissions method.
		GetCallerPermissions []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Params is the params argument value.
			Params Parameters
		}
		// GetPermissions holds details about calls to the GetPermissions method.
		GetPermissions []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// GetPermissionsRequest is the getPermissionsRequest argument value.
			GetPermissionsRequest *http.Request
		}
	}
}

// GetCallerPermissions calls GetCallerPermissionsFunc.
func (mock *ClienterMock) GetCallerPermissions(ctx context.Context, params Parameters) (*Permissions, error) {
	if mock.GetCallerPermissionsFunc == nil {
		panic("moq: ClienterMock.GetCallerPermissionsFunc is nil but Clienter.GetCallerPermissions was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Params Parameters
	}{
		Ctx:    ctx,
		Params: params,
	}
	lockClienterMockGetCallerPermissions.Lock()
	mock.calls.GetCallerPermissions = append(mock.calls.GetCallerPermissions, callInfo)
	lockClienterMockGetCallerPermissions.Unlock()
	return mock.GetCallerPermissionsFunc(ctx, params)
}

// GetCallerPermissionsCalls gets all the calls that were made to GetCallerPermissions.
// Check the length with:
//     len(mockedClienter.GetCallerPermissionsCalls())
func (mock *ClienterMock) GetCallerPermissionsCalls() []struct {
	Ctx    context.Context
	Params Parameters
} {
	var calls []struct {
		Ctx    context.Context
		Params Parameters
	}
	lockClienterMockGetCallerPermissions.RLock()
	calls = mock.calls.GetCallerPermissions
	lockClienterMockGetCallerPermissions.RUnlock()
	return calls
}

// GetPermissions calls GetPermissionsFunc.
func (mock *ClienterMock) GetPermissions(ctx context.Context, getPermissionsRequest *http.Request) (*Permissions, error) {
	if mock.GetPermissionsFunc == nil {
		panic("moq: ClienterMock.GetPermissionsFunc is nil but Clienter.GetPermissions was just called")
	}
	callInfo := struct {
		Ctx                   context.Context
		GetPermissionsRequest *http.Request
	}{
		Ctx:                   ctx,
		GetPermissionsRequest: getPermissionsRequest,
	}
	lockClienterMockGetPermissions.Lock()
	mock.calls.GetPermissions = append(mock.calls.GetPermissions, callInfo)
	lockClienterMockGetPermissions.Unlock()
	return mock.GetPermissionsFunc(ctx, getPermissionsRequest)
}

// GetPermissionsCalls gets all the calls that were made to GetPermissions.
// Check the length with:
//     len(mockedClienter.GetPermissionsCalls())
func (mock *ClienterMock) GetPermissionsCalls() []struct {
	Ctx                   context.Context
	GetPermissionsRequest *http.Request
} {
	var calls []struct {
		Ctx                   context.Context
		GetPermissionsRequest *http.Request
	}
	lockClienterMockGetPermissions.RLock()
	calls = mock.calls.GetPermissions
	lockClienterMockGetPermissions.RUnlock()
	return calls
}

var (
	lockVerifierMockCheckAuthorisation sync.RWMutex
)

// VerifierMock is a mock implementation of Verifier.
//
//     func TestSomethingThatUsesVerifier(t *testing.T) {
//
//         // make and configure a mocked Verifier
//         mockedVerifier := &VerifierMock{
//             CheckAuthorisationFunc: func(ctx context.Context, callerPermissions *Permissions, requiredPermissions *Permissions) error {
// 	               panic("TODO: mock out the CheckAuthorisation method")
//             },
//         }
//
//         // TODO: use mockedVerifier in code that requires Verifier
//         //       and then make assertions.
//
//     }
type VerifierMock struct {
	// CheckAuthorisationFunc mocks the CheckAuthorisation method.
	CheckAuthorisationFunc func(ctx context.Context, callerPermissions *Permissions, requiredPermissions *Permissions) error

	// calls tracks calls to the methods.
	calls struct {
		// CheckAuthorisation holds details about calls to the CheckAuthorisation method.
		CheckAuthorisation []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// CallerPermissions is the callerPermissions argument value.
			CallerPermissions *Permissions
			// RequiredPermissions is the requiredPermissions argument value.
			RequiredPermissions *Permissions
		}
	}
}

// CheckAuthorisation calls CheckAuthorisationFunc.
func (mock *VerifierMock) CheckAuthorisation(ctx context.Context, callerPermissions *Permissions, requiredPermissions *Permissions) error {
	if mock.CheckAuthorisationFunc == nil {
		panic("moq: VerifierMock.CheckAuthorisationFunc is nil but Verifier.CheckAuthorisation was just called")
	}
	callInfo := struct {
		Ctx                 context.Context
		CallerPermissions   *Permissions
		RequiredPermissions *Permissions
	}{
		Ctx:                 ctx,
		CallerPermissions:   callerPermissions,
		RequiredPermissions: requiredPermissions,
	}
	lockVerifierMockCheckAuthorisation.Lock()
	mock.calls.CheckAuthorisation = append(mock.calls.CheckAuthorisation, callInfo)
	lockVerifierMockCheckAuthorisation.Unlock()
	return mock.CheckAuthorisationFunc(ctx, callerPermissions, requiredPermissions)
}

// CheckAuthorisationCalls gets all the calls that were made to CheckAuthorisation.
// Check the length with:
//     len(mockedVerifier.CheckAuthorisationCalls())
func (mock *VerifierMock) CheckAuthorisationCalls() []struct {
	Ctx                 context.Context
	CallerPermissions   *Permissions
	RequiredPermissions *Permissions
} {
	var calls []struct {
		Ctx                 context.Context
		CallerPermissions   *Permissions
		RequiredPermissions *Permissions
	}
	lockVerifierMockCheckAuthorisation.RLock()
	calls = mock.calls.CheckAuthorisation
	lockVerifierMockCheckAuthorisation.RUnlock()
	return calls
}

var (
	lockHTTPClienterMockDo sync.RWMutex
)

// HTTPClienterMock is a mock implementation of HTTPClienter.
//
//     func TestSomethingThatUsesHTTPClienter(t *testing.T) {
//
//         // make and configure a mocked HTTPClienter
//         mockedHTTPClienter := &HTTPClienterMock{
//             DoFunc: func(ctx context.Context, req *http.Request) (*http.Response, error) {
// 	               panic("TODO: mock out the Do method")
//             },
//         }
//
//         // TODO: use mockedHTTPClienter in code that requires HTTPClienter
//         //       and then make assertions.
//
//     }
type HTTPClienterMock struct {
	// DoFunc mocks the Do method.
	DoFunc func(ctx context.Context, req *http.Request) (*http.Response, error)

	// calls tracks calls to the methods.
	calls struct {
		// Do holds details about calls to the Do method.
		Do []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Req is the req argument value.
			Req *http.Request
		}
	}
}

// Do calls DoFunc.
func (mock *HTTPClienterMock) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	if mock.DoFunc == nil {
		panic("moq: HTTPClienterMock.DoFunc is nil but HTTPClienter.Do was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Req *http.Request
	}{
		Ctx: ctx,
		Req: req,
	}
	lockHTTPClienterMockDo.Lock()
	mock.calls.Do = append(mock.calls.Do, callInfo)
	lockHTTPClienterMockDo.Unlock()
	return mock.DoFunc(ctx, req)
}

// DoCalls gets all the calls that were made to Do.
// Check the length with:
//     len(mockedHTTPClienter.DoCalls())
func (mock *HTTPClienterMock) DoCalls() []struct {
	Ctx context.Context
	Req *http.Request
} {
	var calls []struct {
		Ctx context.Context
		Req *http.Request
	}
	lockHTTPClienterMockDo.RLock()
	calls = mock.calls.Do
	lockHTTPClienterMockDo.RUnlock()
	return calls
}

var (
	lockParametersMockCreateGetPermissionsRequest sync.RWMutex
)

// ParametersMock is a mock implementation of Parameters.
//
//     func TestSomethingThatUsesParameters(t *testing.T) {
//
//         // make and configure a mocked Parameters
//         mockedParameters := &ParametersMock{
//             CreateGetPermissionsRequestFunc: func(host string) (*http.Request, error) {
// 	               panic("TODO: mock out the CreateGetPermissionsRequest method")
//             },
//         }
//
//         // TODO: use mockedParameters in code that requires Parameters
//         //       and then make assertions.
//
//     }
type ParametersMock struct {
	// CreateGetPermissionsRequestFunc mocks the CreateGetPermissionsRequest method.
	CreateGetPermissionsRequestFunc func(host string) (*http.Request, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateGetPermissionsRequest holds details about calls to the CreateGetPermissionsRequest method.
		CreateGetPermissionsRequest []struct {
			// Host is the host argument value.
			Host string
		}
	}
}

// CreateGetPermissionsRequest calls CreateGetPermissionsRequestFunc.
func (mock *ParametersMock) CreateGetPermissionsRequest(host string) (*http.Request, error) {
	if mock.CreateGetPermissionsRequestFunc == nil {
		panic("moq: ParametersMock.CreateGetPermissionsRequestFunc is nil but Parameters.CreateGetPermissionsRequest was just called")
	}
	callInfo := struct {
		Host string
	}{
		Host: host,
	}
	lockParametersMockCreateGetPermissionsRequest.Lock()
	mock.calls.CreateGetPermissionsRequest = append(mock.calls.CreateGetPermissionsRequest, callInfo)
	lockParametersMockCreateGetPermissionsRequest.Unlock()
	return mock.CreateGetPermissionsRequestFunc(host)
}

// CreateGetPermissionsRequestCalls gets all the calls that were made to CreateGetPermissionsRequest.
// Check the length with:
//     len(mockedParameters.CreateGetPermissionsRequestCalls())
func (mock *ParametersMock) CreateGetPermissionsRequestCalls() []struct {
	Host string
} {
	var calls []struct {
		Host string
	}
	lockParametersMockCreateGetPermissionsRequest.RLock()
	calls = mock.calls.CreateGetPermissionsRequest
	lockParametersMockCreateGetPermissionsRequest.RUnlock()
	return calls
}

var (
	lockParameterFactoryMockCreateParameters sync.RWMutex
)

// ParameterFactoryMock is a mock implementation of ParameterFactory.
//
//     func TestSomethingThatUsesParameterFactory(t *testing.T) {
//
//         // make and configure a mocked ParameterFactory
//         mockedParameterFactory := &ParameterFactoryMock{
//             CreateParametersFunc: func(req *http.Request) (Parameters, error) {
// 	               panic("TODO: mock out the CreateParameters method")
//             },
//         }
//
//         // TODO: use mockedParameterFactory in code that requires ParameterFactory
//         //       and then make assertions.
//
//     }
type ParameterFactoryMock struct {
	// CreateParametersFunc mocks the CreateParameters method.
	CreateParametersFunc func(req *http.Request) (Parameters, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateParameters holds details about calls to the CreateParameters method.
		CreateParameters []struct {
			// Req is the req argument value.
			Req *http.Request
		}
	}
}

// CreateParameters calls CreateParametersFunc.
func (mock *ParameterFactoryMock) CreateParameters(req *http.Request) (Parameters, error) {
	if mock.CreateParametersFunc == nil {
		panic("moq: ParameterFactoryMock.CreateParametersFunc is nil but ParameterFactory.CreateParameters was just called")
	}
	callInfo := struct {
		Req *http.Request
	}{
		Req: req,
	}
	lockParameterFactoryMockCreateParameters.Lock()
	mock.calls.CreateParameters = append(mock.calls.CreateParameters, callInfo)
	lockParameterFactoryMockCreateParameters.Unlock()
	return mock.CreateParametersFunc(req)
}

// CreateParametersCalls gets all the calls that were made to CreateParameters.
// Check the length with:
//     len(mockedParameterFactory.CreateParametersCalls())
func (mock *ParameterFactoryMock) CreateParametersCalls() []struct {
	Req *http.Request
} {
	var calls []struct {
		Req *http.Request
	}
	lockParameterFactoryMockCreateParameters.RLock()
	calls = mock.calls.CreateParameters
	lockParameterFactoryMockCreateParameters.RUnlock()
	return calls
}

var (
	lockGetPermissionsRequestBuilderMockNewPermissionsRequest sync.RWMutex
)

// GetPermissionsRequestBuilderMock is a mock implementation of GetPermissionsRequestBuilder.
//
//     func TestSomethingThatUsesGetPermissionsRequestBuilder(t *testing.T) {
//
//         // make and configure a mocked GetPermissionsRequestBuilder
//         mockedGetPermissionsRequestBuilder := &GetPermissionsRequestBuilderMock{
//             NewPermissionsRequestFunc: func(req *http.Request) (*http.Request, error) {
// 	               panic("TODO: mock out the NewPermissionsRequest method")
//             },
//         }
//
//         // TODO: use mockedGetPermissionsRequestBuilder in code that requires GetPermissionsRequestBuilder
//         //       and then make assertions.
//
//     }
type GetPermissionsRequestBuilderMock struct {
	// NewPermissionsRequestFunc mocks the NewPermissionsRequest method.
	NewPermissionsRequestFunc func(req *http.Request) (*http.Request, error)

	// calls tracks calls to the methods.
	calls struct {
		// NewPermissionsRequest holds details about calls to the NewPermissionsRequest method.
		NewPermissionsRequest []struct {
			// Req is the req argument value.
			Req *http.Request
		}
	}
}

// NewPermissionsRequest calls NewPermissionsRequestFunc.
func (mock *GetPermissionsRequestBuilderMock) NewPermissionsRequest(req *http.Request) (*http.Request, error) {
	if mock.NewPermissionsRequestFunc == nil {
		panic("moq: GetPermissionsRequestBuilderMock.NewPermissionsRequestFunc is nil but GetPermissionsRequestBuilder.NewPermissionsRequest was just called")
	}
	callInfo := struct {
		Req *http.Request
	}{
		Req: req,
	}
	lockGetPermissionsRequestBuilderMockNewPermissionsRequest.Lock()
	mock.calls.NewPermissionsRequest = append(mock.calls.NewPermissionsRequest, callInfo)
	lockGetPermissionsRequestBuilderMockNewPermissionsRequest.Unlock()
	return mock.NewPermissionsRequestFunc(req)
}

// NewPermissionsRequestCalls gets all the calls that were made to NewPermissionsRequest.
// Check the length with:
//     len(mockedGetPermissionsRequestBuilder.NewPermissionsRequestCalls())
func (mock *GetPermissionsRequestBuilderMock) NewPermissionsRequestCalls() []struct {
	Req *http.Request
} {
	var calls []struct {
		Req *http.Request
	}
	lockGetPermissionsRequestBuilderMockNewPermissionsRequest.RLock()
	calls = mock.calls.NewPermissionsRequest
	lockGetPermissionsRequestBuilderMockNewPermissionsRequest.RUnlock()
	return calls
}
