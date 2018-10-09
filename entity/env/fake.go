// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package env

import (
	"github.com/ktr0731/evans/entity"
	"sync"
)

var (
	lockEnvironmentMockAddHeader    sync.RWMutex
	lockEnvironmentMockDSN          sync.RWMutex
	lockEnvironmentMockHeaders      sync.RWMutex
	lockEnvironmentMockMessage      sync.RWMutex
	lockEnvironmentMockMessages     sync.RWMutex
	lockEnvironmentMockPackages     sync.RWMutex
	lockEnvironmentMockRPC          sync.RWMutex
	lockEnvironmentMockRPCs         sync.RWMutex
	lockEnvironmentMockRemoveHeader sync.RWMutex
	lockEnvironmentMockService      sync.RWMutex
	lockEnvironmentMockServices     sync.RWMutex
	lockEnvironmentMockUsePackage   sync.RWMutex
	lockEnvironmentMockUseService   sync.RWMutex
)

// EnvironmentMock is a mock implementation of Environment.
//
//     func TestSomethingThatUsesEnvironment(t *testing.T) {
//
//         // make and configure a mocked Environment
//         mockedEnvironment := &EnvironmentMock{
//             AddHeaderFunc: func(header *entity.Header)  {
// 	               panic("TODO: mock out the AddHeader method")
//             },
//             DSNFunc: func() string {
// 	               panic("TODO: mock out the DSN method")
//             },
//             HeadersFunc: func() []*entity.Header {
// 	               panic("TODO: mock out the Headers method")
//             },
//             MessageFunc: func(name string) (entity.Message, error) {
// 	               panic("TODO: mock out the Message method")
//             },
//             MessagesFunc: func() ([]entity.Message, error) {
// 	               panic("TODO: mock out the Messages method")
//             },
//             PackagesFunc: func() []*entity.Package {
// 	               panic("TODO: mock out the Packages method")
//             },
//             RPCFunc: func(name string) (entity.RPC, error) {
// 	               panic("TODO: mock out the RPC method")
//             },
//             RPCsFunc: func() ([]entity.RPC, error) {
// 	               panic("TODO: mock out the RPCs method")
//             },
//             RemoveHeaderFunc: func(key string)  {
// 	               panic("TODO: mock out the RemoveHeader method")
//             },
//             ServiceFunc: func(name string) (entity.Service, error) {
// 	               panic("TODO: mock out the Service method")
//             },
//             ServicesFunc: func() ([]entity.Service, error) {
// 	               panic("TODO: mock out the Services method")
//             },
//             UsePackageFunc: func(name string) error {
// 	               panic("TODO: mock out the UsePackage method")
//             },
//             UseServiceFunc: func(name string) error {
// 	               panic("TODO: mock out the UseService method")
//             },
//         }
//
//         // TODO: use mockedEnvironment in code that requires Environment
//         //       and then make assertions.
//
//     }
type EnvironmentMock struct {
	// AddHeaderFunc mocks the AddHeader method.
	AddHeaderFunc func(header *entity.Header)

	// DSNFunc mocks the DSN method.
	DSNFunc func() string

	// HeadersFunc mocks the Headers method.
	HeadersFunc func() []*entity.Header

	// MessageFunc mocks the Message method.
	MessageFunc func(name string) (entity.Message, error)

	// MessagesFunc mocks the Messages method.
	MessagesFunc func() ([]entity.Message, error)

	// PackagesFunc mocks the Packages method.
	PackagesFunc func() []*entity.Package

	// RPCFunc mocks the RPC method.
	RPCFunc func(name string) (entity.RPC, error)

	// RPCsFunc mocks the RPCs method.
	RPCsFunc func() ([]entity.RPC, error)

	// RemoveHeaderFunc mocks the RemoveHeader method.
	RemoveHeaderFunc func(key string)

	// ServiceFunc mocks the Service method.
	ServiceFunc func(name string) (entity.Service, error)

	// ServicesFunc mocks the Services method.
	ServicesFunc func() ([]entity.Service, error)

	// UsePackageFunc mocks the UsePackage method.
	UsePackageFunc func(name string) error

	// UseServiceFunc mocks the UseService method.
	UseServiceFunc func(name string) error

	// calls tracks calls to the methods.
	calls struct {
		// AddHeader holds details about calls to the AddHeader method.
		AddHeader []struct {
			// Header is the header argument value.
			Header *entity.Header
		}
		// DSN holds details about calls to the DSN method.
		DSN []struct {
		}
		// Headers holds details about calls to the Headers method.
		Headers []struct {
		}
		// Message holds details about calls to the Message method.
		Message []struct {
			// Name is the name argument value.
			Name string
		}
		// Messages holds details about calls to the Messages method.
		Messages []struct {
		}
		// Packages holds details about calls to the Packages method.
		Packages []struct {
		}
		// RPC holds details about calls to the RPC method.
		RPC []struct {
			// Name is the name argument value.
			Name string
		}
		// RPCs holds details about calls to the RPCs method.
		RPCs []struct {
		}
		// RemoveHeader holds details about calls to the RemoveHeader method.
		RemoveHeader []struct {
			// Key is the key argument value.
			Key string
		}
		// Service holds details about calls to the Service method.
		Service []struct {
			// Name is the name argument value.
			Name string
		}
		// Services holds details about calls to the Services method.
		Services []struct {
		}
		// UsePackage holds details about calls to the UsePackage method.
		UsePackage []struct {
			// Name is the name argument value.
			Name string
		}
		// UseService holds details about calls to the UseService method.
		UseService []struct {
			// Name is the name argument value.
			Name string
		}
	}
}

// AddHeader calls AddHeaderFunc.
func (mock *EnvironmentMock) AddHeader(header *entity.Header) {
	if mock.AddHeaderFunc == nil {
		panic("EnvironmentMock.AddHeaderFunc: method is nil but Environment.AddHeader was just called")
	}
	callInfo := struct {
		Header *entity.Header
	}{
		Header: header,
	}
	lockEnvironmentMockAddHeader.Lock()
	mock.calls.AddHeader = append(mock.calls.AddHeader, callInfo)
	lockEnvironmentMockAddHeader.Unlock()
	mock.AddHeaderFunc(header)
}

// AddHeaderCalls gets all the calls that were made to AddHeader.
// Check the length with:
//     len(mockedEnvironment.AddHeaderCalls())
func (mock *EnvironmentMock) AddHeaderCalls() []struct {
	Header *entity.Header
} {
	var calls []struct {
		Header *entity.Header
	}
	lockEnvironmentMockAddHeader.RLock()
	calls = mock.calls.AddHeader
	lockEnvironmentMockAddHeader.RUnlock()
	return calls
}

// DSN calls DSNFunc.
func (mock *EnvironmentMock) DSN() string {
	if mock.DSNFunc == nil {
		panic("EnvironmentMock.DSNFunc: method is nil but Environment.DSN was just called")
	}
	callInfo := struct {
	}{}
	lockEnvironmentMockDSN.Lock()
	mock.calls.DSN = append(mock.calls.DSN, callInfo)
	lockEnvironmentMockDSN.Unlock()
	return mock.DSNFunc()
}

// DSNCalls gets all the calls that were made to DSN.
// Check the length with:
//     len(mockedEnvironment.DSNCalls())
func (mock *EnvironmentMock) DSNCalls() []struct {
} {
	var calls []struct {
	}
	lockEnvironmentMockDSN.RLock()
	calls = mock.calls.DSN
	lockEnvironmentMockDSN.RUnlock()
	return calls
}

// Headers calls HeadersFunc.
func (mock *EnvironmentMock) Headers() []*entity.Header {
	if mock.HeadersFunc == nil {
		panic("EnvironmentMock.HeadersFunc: method is nil but Environment.Headers was just called")
	}
	callInfo := struct {
	}{}
	lockEnvironmentMockHeaders.Lock()
	mock.calls.Headers = append(mock.calls.Headers, callInfo)
	lockEnvironmentMockHeaders.Unlock()
	return mock.HeadersFunc()
}

// HeadersCalls gets all the calls that were made to Headers.
// Check the length with:
//     len(mockedEnvironment.HeadersCalls())
func (mock *EnvironmentMock) HeadersCalls() []struct {
} {
	var calls []struct {
	}
	lockEnvironmentMockHeaders.RLock()
	calls = mock.calls.Headers
	lockEnvironmentMockHeaders.RUnlock()
	return calls
}

// Message calls MessageFunc.
func (mock *EnvironmentMock) Message(name string) (entity.Message, error) {
	if mock.MessageFunc == nil {
		panic("EnvironmentMock.MessageFunc: method is nil but Environment.Message was just called")
	}
	callInfo := struct {
		Name string
	}{
		Name: name,
	}
	lockEnvironmentMockMessage.Lock()
	mock.calls.Message = append(mock.calls.Message, callInfo)
	lockEnvironmentMockMessage.Unlock()
	return mock.MessageFunc(name)
}

// MessageCalls gets all the calls that were made to Message.
// Check the length with:
//     len(mockedEnvironment.MessageCalls())
func (mock *EnvironmentMock) MessageCalls() []struct {
	Name string
} {
	var calls []struct {
		Name string
	}
	lockEnvironmentMockMessage.RLock()
	calls = mock.calls.Message
	lockEnvironmentMockMessage.RUnlock()
	return calls
}

// Messages calls MessagesFunc.
func (mock *EnvironmentMock) Messages() ([]entity.Message, error) {
	if mock.MessagesFunc == nil {
		panic("EnvironmentMock.MessagesFunc: method is nil but Environment.Messages was just called")
	}
	callInfo := struct {
	}{}
	lockEnvironmentMockMessages.Lock()
	mock.calls.Messages = append(mock.calls.Messages, callInfo)
	lockEnvironmentMockMessages.Unlock()
	return mock.MessagesFunc()
}

// MessagesCalls gets all the calls that were made to Messages.
// Check the length with:
//     len(mockedEnvironment.MessagesCalls())
func (mock *EnvironmentMock) MessagesCalls() []struct {
} {
	var calls []struct {
	}
	lockEnvironmentMockMessages.RLock()
	calls = mock.calls.Messages
	lockEnvironmentMockMessages.RUnlock()
	return calls
}

// Packages calls PackagesFunc.
func (mock *EnvironmentMock) Packages() []*entity.Package {
	if mock.PackagesFunc == nil {
		panic("EnvironmentMock.PackagesFunc: method is nil but Environment.Packages was just called")
	}
	callInfo := struct {
	}{}
	lockEnvironmentMockPackages.Lock()
	mock.calls.Packages = append(mock.calls.Packages, callInfo)
	lockEnvironmentMockPackages.Unlock()
	return mock.PackagesFunc()
}

// PackagesCalls gets all the calls that were made to Packages.
// Check the length with:
//     len(mockedEnvironment.PackagesCalls())
func (mock *EnvironmentMock) PackagesCalls() []struct {
} {
	var calls []struct {
	}
	lockEnvironmentMockPackages.RLock()
	calls = mock.calls.Packages
	lockEnvironmentMockPackages.RUnlock()
	return calls
}

// RPC calls RPCFunc.
func (mock *EnvironmentMock) RPC(name string) (entity.RPC, error) {
	if mock.RPCFunc == nil {
		panic("EnvironmentMock.RPCFunc: method is nil but Environment.RPC was just called")
	}
	callInfo := struct {
		Name string
	}{
		Name: name,
	}
	lockEnvironmentMockRPC.Lock()
	mock.calls.RPC = append(mock.calls.RPC, callInfo)
	lockEnvironmentMockRPC.Unlock()
	return mock.RPCFunc(name)
}

// RPCCalls gets all the calls that were made to RPC.
// Check the length with:
//     len(mockedEnvironment.RPCCalls())
func (mock *EnvironmentMock) RPCCalls() []struct {
	Name string
} {
	var calls []struct {
		Name string
	}
	lockEnvironmentMockRPC.RLock()
	calls = mock.calls.RPC
	lockEnvironmentMockRPC.RUnlock()
	return calls
}

// RPCs calls RPCsFunc.
func (mock *EnvironmentMock) RPCs() ([]entity.RPC, error) {
	if mock.RPCsFunc == nil {
		panic("EnvironmentMock.RPCsFunc: method is nil but Environment.RPCs was just called")
	}
	callInfo := struct {
	}{}
	lockEnvironmentMockRPCs.Lock()
	mock.calls.RPCs = append(mock.calls.RPCs, callInfo)
	lockEnvironmentMockRPCs.Unlock()
	return mock.RPCsFunc()
}

// RPCsCalls gets all the calls that were made to RPCs.
// Check the length with:
//     len(mockedEnvironment.RPCsCalls())
func (mock *EnvironmentMock) RPCsCalls() []struct {
} {
	var calls []struct {
	}
	lockEnvironmentMockRPCs.RLock()
	calls = mock.calls.RPCs
	lockEnvironmentMockRPCs.RUnlock()
	return calls
}

// RemoveHeader calls RemoveHeaderFunc.
func (mock *EnvironmentMock) RemoveHeader(key string) {
	if mock.RemoveHeaderFunc == nil {
		panic("EnvironmentMock.RemoveHeaderFunc: method is nil but Environment.RemoveHeader was just called")
	}
	callInfo := struct {
		Key string
	}{
		Key: key,
	}
	lockEnvironmentMockRemoveHeader.Lock()
	mock.calls.RemoveHeader = append(mock.calls.RemoveHeader, callInfo)
	lockEnvironmentMockRemoveHeader.Unlock()
	mock.RemoveHeaderFunc(key)
}

// RemoveHeaderCalls gets all the calls that were made to RemoveHeader.
// Check the length with:
//     len(mockedEnvironment.RemoveHeaderCalls())
func (mock *EnvironmentMock) RemoveHeaderCalls() []struct {
	Key string
} {
	var calls []struct {
		Key string
	}
	lockEnvironmentMockRemoveHeader.RLock()
	calls = mock.calls.RemoveHeader
	lockEnvironmentMockRemoveHeader.RUnlock()
	return calls
}

// Service calls ServiceFunc.
func (mock *EnvironmentMock) Service(name string) (entity.Service, error) {
	if mock.ServiceFunc == nil {
		panic("EnvironmentMock.ServiceFunc: method is nil but Environment.Service was just called")
	}
	callInfo := struct {
		Name string
	}{
		Name: name,
	}
	lockEnvironmentMockService.Lock()
	mock.calls.Service = append(mock.calls.Service, callInfo)
	lockEnvironmentMockService.Unlock()
	return mock.ServiceFunc(name)
}

// ServiceCalls gets all the calls that were made to Service.
// Check the length with:
//     len(mockedEnvironment.ServiceCalls())
func (mock *EnvironmentMock) ServiceCalls() []struct {
	Name string
} {
	var calls []struct {
		Name string
	}
	lockEnvironmentMockService.RLock()
	calls = mock.calls.Service
	lockEnvironmentMockService.RUnlock()
	return calls
}

// Services calls ServicesFunc.
func (mock *EnvironmentMock) Services() ([]entity.Service, error) {
	if mock.ServicesFunc == nil {
		panic("EnvironmentMock.ServicesFunc: method is nil but Environment.Services was just called")
	}
	callInfo := struct {
	}{}
	lockEnvironmentMockServices.Lock()
	mock.calls.Services = append(mock.calls.Services, callInfo)
	lockEnvironmentMockServices.Unlock()
	return mock.ServicesFunc()
}

// ServicesCalls gets all the calls that were made to Services.
// Check the length with:
//     len(mockedEnvironment.ServicesCalls())
func (mock *EnvironmentMock) ServicesCalls() []struct {
} {
	var calls []struct {
	}
	lockEnvironmentMockServices.RLock()
	calls = mock.calls.Services
	lockEnvironmentMockServices.RUnlock()
	return calls
}

// UsePackage calls UsePackageFunc.
func (mock *EnvironmentMock) UsePackage(name string) error {
	if mock.UsePackageFunc == nil {
		panic("EnvironmentMock.UsePackageFunc: method is nil but Environment.UsePackage was just called")
	}
	callInfo := struct {
		Name string
	}{
		Name: name,
	}
	lockEnvironmentMockUsePackage.Lock()
	mock.calls.UsePackage = append(mock.calls.UsePackage, callInfo)
	lockEnvironmentMockUsePackage.Unlock()
	return mock.UsePackageFunc(name)
}

// UsePackageCalls gets all the calls that were made to UsePackage.
// Check the length with:
//     len(mockedEnvironment.UsePackageCalls())
func (mock *EnvironmentMock) UsePackageCalls() []struct {
	Name string
} {
	var calls []struct {
		Name string
	}
	lockEnvironmentMockUsePackage.RLock()
	calls = mock.calls.UsePackage
	lockEnvironmentMockUsePackage.RUnlock()
	return calls
}

// UseService calls UseServiceFunc.
func (mock *EnvironmentMock) UseService(name string) error {
	if mock.UseServiceFunc == nil {
		panic("EnvironmentMock.UseServiceFunc: method is nil but Environment.UseService was just called")
	}
	callInfo := struct {
		Name string
	}{
		Name: name,
	}
	lockEnvironmentMockUseService.Lock()
	mock.calls.UseService = append(mock.calls.UseService, callInfo)
	lockEnvironmentMockUseService.Unlock()
	return mock.UseServiceFunc(name)
}

// UseServiceCalls gets all the calls that were made to UseService.
// Check the length with:
//     len(mockedEnvironment.UseServiceCalls())
func (mock *EnvironmentMock) UseServiceCalls() []struct {
	Name string
} {
	var calls []struct {
		Name string
	}
	lockEnvironmentMockUseService.RLock()
	calls = mock.calls.UseService
	lockEnvironmentMockUseService.RUnlock()
	return calls
}
