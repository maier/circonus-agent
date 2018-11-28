// Code generated by moq; DO NOT EDIT
// github.com/matryer/moq

package check

import (
	"sync"

	"github.com/circonus-labs/go-apiclient"
)

var (
	lockAPIMockCreateCheckBundle        sync.RWMutex
	lockAPIMockFetchBroker              sync.RWMutex
	lockAPIMockFetchBrokers             sync.RWMutex
	lockAPIMockFetchCheckBundle         sync.RWMutex
	lockAPIMockFetchCheckBundleMetrics  sync.RWMutex
	lockAPIMockGet                      sync.RWMutex
	lockAPIMockSearchCheckBundles       sync.RWMutex
	lockAPIMockUpdateCheckBundle        sync.RWMutex
	lockAPIMockUpdateCheckBundleMetrics sync.RWMutex
)

// APIMock is a mock implementation of API.
//
//     func TestSomethingThatUsesAPI(t *testing.T) {
//
//         // make and configure a mocked API
//         mockedAPI := &APIMock{
//             CreateCheckBundleFunc: func(cfg *apiclient.CheckBundle) (*apiclient.CheckBundle, error) {
// 	               panic("TODO: mock out the CreateCheckBundle method")
//             },
//             FetchBrokerFunc: func(cid apiclient.CIDType) (*apiclient.Broker, error) {
// 	               panic("TODO: mock out the FetchBroker method")
//             },
//             FetchBrokersFunc: func() (*[]apiclient.Broker, error) {
// 	               panic("TODO: mock out the FetchBrokers method")
//             },
//             FetchCheckBundleFunc: func(cid apiclient.CIDType) (*apiclient.CheckBundle, error) {
// 	               panic("TODO: mock out the FetchCheckBundle method")
//             },
//             FetchCheckBundleMetricsFunc: func(cid apiclient.CIDType) (*apiclient.CheckBundleMetrics, error) {
// 	               panic("TODO: mock out the FetchCheckBundleMetrics method")
//             },
//             GetFunc: func(url string) ([]byte, error) {
// 	               panic("TODO: mock out the Get method")
//             },
//             SearchCheckBundlesFunc: func(searchCriteria *apiclient.SearchQueryType, filterCriteria *map[string][]string) (*[]apiclient.CheckBundle, error) {
// 	               panic("TODO: mock out the SearchCheckBundles method")
//             },
//             UpdateCheckBundleFunc: func(cfg *apiclient.CheckBundle) (*apiclient.CheckBundle, error) {
// 	               panic("TODO: mock out the UpdateCheckBundle method")
//             },
//             UpdateCheckBundleMetricsFunc: func(cfg *apiclient.CheckBundleMetrics) (*apiclient.CheckBundleMetrics, error) {
// 	               panic("TODO: mock out the UpdateCheckBundleMetrics method")
//             },
//         }
//
//         // TODO: use mockedAPI in code that requires API
//         //       and then make assertions.
//
//     }
type APIMock struct {
	// CreateCheckBundleFunc mocks the CreateCheckBundle method.
	CreateCheckBundleFunc func(cfg *apiclient.CheckBundle) (*apiclient.CheckBundle, error)

	// FetchBrokerFunc mocks the FetchBroker method.
	FetchBrokerFunc func(cid apiclient.CIDType) (*apiclient.Broker, error)

	// FetchBrokersFunc mocks the FetchBrokers method.
	FetchBrokersFunc func() (*[]apiclient.Broker, error)

	// FetchCheckBundleFunc mocks the FetchCheckBundle method.
	FetchCheckBundleFunc func(cid apiclient.CIDType) (*apiclient.CheckBundle, error)

	// FetchCheckBundleMetricsFunc mocks the FetchCheckBundleMetrics method.
	FetchCheckBundleMetricsFunc func(cid apiclient.CIDType) (*apiclient.CheckBundleMetrics, error)

	// GetFunc mocks the Get method.
	GetFunc func(url string) ([]byte, error)

	// SearchCheckBundlesFunc mocks the SearchCheckBundles method.
	SearchCheckBundlesFunc func(searchCriteria *apiclient.SearchQueryType, filterCriteria *apiclient.SearchFilterType) (*[]apiclient.CheckBundle, error)

	// UpdateCheckBundleFunc mocks the UpdateCheckBundle method.
	UpdateCheckBundleFunc func(cfg *apiclient.CheckBundle) (*apiclient.CheckBundle, error)

	// UpdateCheckBundleMetricsFunc mocks the UpdateCheckBundleMetrics method.
	UpdateCheckBundleMetricsFunc func(cfg *apiclient.CheckBundleMetrics) (*apiclient.CheckBundleMetrics, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateCheckBundle holds details about calls to the CreateCheckBundle method.
		CreateCheckBundle []struct {
			// Cfg is the cfg argument value.
			Cfg *apiclient.CheckBundle
		}
		// FetchBroker holds details about calls to the FetchBroker method.
		FetchBroker []struct {
			// Cid is the cid argument value.
			Cid apiclient.CIDType
		}
		// FetchBrokers holds details about calls to the FetchBrokers method.
		FetchBrokers []struct {
		}
		// FetchCheckBundle holds details about calls to the FetchCheckBundle method.
		FetchCheckBundle []struct {
			// Cid is the cid argument value.
			Cid apiclient.CIDType
		}
		// FetchCheckBundleMetrics holds details about calls to the FetchCheckBundleMetrics method.
		FetchCheckBundleMetrics []struct {
			// Cid is the cid argument value.
			Cid apiclient.CIDType
		}
		// Get holds details about calls to the Get method.
		Get []struct {
			// URL is the url argument value.
			URL string
		}
		// SearchCheckBundles holds details about calls to the SearchCheckBundles method.
		SearchCheckBundles []struct {
			// SearchCriteria is the searchCriteria argument value.
			SearchCriteria *apiclient.SearchQueryType
			// FilterCriteria is the filterCriteria argument value.
			FilterCriteria *apiclient.SearchFilterType
		}
		// UpdateCheckBundle holds details about calls to the UpdateCheckBundle method.
		UpdateCheckBundle []struct {
			// Cfg is the cfg argument value.
			Cfg *apiclient.CheckBundle
		}
		// UpdateCheckBundleMetrics holds details about calls to the UpdateCheckBundleMetrics method.
		UpdateCheckBundleMetrics []struct {
			// Cfg is the cfg argument value.
			Cfg *apiclient.CheckBundleMetrics
		}
	}
}

// CreateCheckBundle calls CreateCheckBundleFunc.
func (mock *APIMock) CreateCheckBundle(cfg *apiclient.CheckBundle) (*apiclient.CheckBundle, error) {
	if mock.CreateCheckBundleFunc == nil {
		panic("moq: APIMock.CreateCheckBundleFunc is nil but API.CreateCheckBundle was just called")
	}
	callInfo := struct {
		Cfg *apiclient.CheckBundle
	}{
		Cfg: cfg,
	}
	lockAPIMockCreateCheckBundle.Lock()
	mock.calls.CreateCheckBundle = append(mock.calls.CreateCheckBundle, callInfo)
	lockAPIMockCreateCheckBundle.Unlock()
	return mock.CreateCheckBundleFunc(cfg)
}

// CreateCheckBundleCalls gets all the calls that were made to CreateCheckBundle.
// Check the length with:
//     len(mockedAPI.CreateCheckBundleCalls())
func (mock *APIMock) CreateCheckBundleCalls() []struct {
	Cfg *apiclient.CheckBundle
} {
	var calls []struct {
		Cfg *apiclient.CheckBundle
	}
	lockAPIMockCreateCheckBundle.RLock()
	calls = mock.calls.CreateCheckBundle
	lockAPIMockCreateCheckBundle.RUnlock()
	return calls
}

// FetchBroker calls FetchBrokerFunc.
func (mock *APIMock) FetchBroker(cid apiclient.CIDType) (*apiclient.Broker, error) {
	if mock.FetchBrokerFunc == nil {
		panic("moq: APIMock.FetchBrokerFunc is nil but API.FetchBroker was just called")
	}
	callInfo := struct {
		Cid apiclient.CIDType
	}{
		Cid: cid,
	}
	lockAPIMockFetchBroker.Lock()
	mock.calls.FetchBroker = append(mock.calls.FetchBroker, callInfo)
	lockAPIMockFetchBroker.Unlock()
	return mock.FetchBrokerFunc(cid)
}

// FetchBrokerCalls gets all the calls that were made to FetchBroker.
// Check the length with:
//     len(mockedAPI.FetchBrokerCalls())
func (mock *APIMock) FetchBrokerCalls() []struct {
	Cid apiclient.CIDType
} {
	var calls []struct {
		Cid apiclient.CIDType
	}
	lockAPIMockFetchBroker.RLock()
	calls = mock.calls.FetchBroker
	lockAPIMockFetchBroker.RUnlock()
	return calls
}

// FetchBrokers calls FetchBrokersFunc.
func (mock *APIMock) FetchBrokers() (*[]apiclient.Broker, error) {
	if mock.FetchBrokersFunc == nil {
		panic("moq: APIMock.FetchBrokersFunc is nil but API.FetchBrokers was just called")
	}
	callInfo := struct {
	}{}
	lockAPIMockFetchBrokers.Lock()
	mock.calls.FetchBrokers = append(mock.calls.FetchBrokers, callInfo)
	lockAPIMockFetchBrokers.Unlock()
	return mock.FetchBrokersFunc()
}

// FetchBrokersCalls gets all the calls that were made to FetchBrokers.
// Check the length with:
//     len(mockedAPI.FetchBrokersCalls())
func (mock *APIMock) FetchBrokersCalls() []struct {
} {
	var calls []struct {
	}
	lockAPIMockFetchBrokers.RLock()
	calls = mock.calls.FetchBrokers
	lockAPIMockFetchBrokers.RUnlock()
	return calls
}

// FetchCheckBundle calls FetchCheckBundleFunc.
func (mock *APIMock) FetchCheckBundle(cid apiclient.CIDType) (*apiclient.CheckBundle, error) {
	if mock.FetchCheckBundleFunc == nil {
		panic("moq: APIMock.FetchCheckBundleFunc is nil but API.FetchCheckBundle was just called")
	}
	callInfo := struct {
		Cid apiclient.CIDType
	}{
		Cid: cid,
	}
	lockAPIMockFetchCheckBundle.Lock()
	mock.calls.FetchCheckBundle = append(mock.calls.FetchCheckBundle, callInfo)
	lockAPIMockFetchCheckBundle.Unlock()
	return mock.FetchCheckBundleFunc(cid)
}

// FetchCheckBundleCalls gets all the calls that were made to FetchCheckBundle.
// Check the length with:
//     len(mockedAPI.FetchCheckBundleCalls())
func (mock *APIMock) FetchCheckBundleCalls() []struct {
	Cid apiclient.CIDType
} {
	var calls []struct {
		Cid apiclient.CIDType
	}
	lockAPIMockFetchCheckBundle.RLock()
	calls = mock.calls.FetchCheckBundle
	lockAPIMockFetchCheckBundle.RUnlock()
	return calls
}

// FetchCheckBundleMetrics calls FetchCheckBundleMetricsFunc.
func (mock *APIMock) FetchCheckBundleMetrics(cid apiclient.CIDType) (*apiclient.CheckBundleMetrics, error) {
	if mock.FetchCheckBundleMetricsFunc == nil {
		panic("moq: APIMock.FetchCheckBundleMetricsFunc is nil but API.FetchCheckBundleMetrics was just called")
	}
	callInfo := struct {
		Cid apiclient.CIDType
	}{
		Cid: cid,
	}
	lockAPIMockFetchCheckBundleMetrics.Lock()
	mock.calls.FetchCheckBundleMetrics = append(mock.calls.FetchCheckBundleMetrics, callInfo)
	lockAPIMockFetchCheckBundleMetrics.Unlock()
	return mock.FetchCheckBundleMetricsFunc(cid)
}

// FetchCheckBundleMetricsCalls gets all the calls that were made to FetchCheckBundleMetrics.
// Check the length with:
//     len(mockedAPI.FetchCheckBundleMetricsCalls())
func (mock *APIMock) FetchCheckBundleMetricsCalls() []struct {
	Cid apiclient.CIDType
} {
	var calls []struct {
		Cid apiclient.CIDType
	}
	lockAPIMockFetchCheckBundleMetrics.RLock()
	calls = mock.calls.FetchCheckBundleMetrics
	lockAPIMockFetchCheckBundleMetrics.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *APIMock) Get(url string) ([]byte, error) {
	if mock.GetFunc == nil {
		panic("moq: APIMock.GetFunc is nil but API.Get was just called")
	}
	callInfo := struct {
		URL string
	}{
		URL: url,
	}
	lockAPIMockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	lockAPIMockGet.Unlock()
	return mock.GetFunc(url)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedAPI.GetCalls())
func (mock *APIMock) GetCalls() []struct {
	URL string
} {
	var calls []struct {
		URL string
	}
	lockAPIMockGet.RLock()
	calls = mock.calls.Get
	lockAPIMockGet.RUnlock()
	return calls
}

// SearchCheckBundles calls SearchCheckBundlesFunc.
func (mock *APIMock) SearchCheckBundles(searchCriteria *apiclient.SearchQueryType, filterCriteria *apiclient.SearchFilterType) (*[]apiclient.CheckBundle, error) {
	if mock.SearchCheckBundlesFunc == nil {
		panic("moq: APIMock.SearchCheckBundlesFunc is nil but API.SearchCheckBundles was just called")
	}
	callInfo := struct {
		SearchCriteria *apiclient.SearchQueryType
		FilterCriteria *apiclient.SearchFilterType
	}{
		SearchCriteria: searchCriteria,
		FilterCriteria: filterCriteria,
	}
	lockAPIMockSearchCheckBundles.Lock()
	mock.calls.SearchCheckBundles = append(mock.calls.SearchCheckBundles, callInfo)
	lockAPIMockSearchCheckBundles.Unlock()
	return mock.SearchCheckBundlesFunc(searchCriteria, filterCriteria)
}

// SearchCheckBundlesCalls gets all the calls that were made to SearchCheckBundles.
// Check the length with:
//     len(mockedAPI.SearchCheckBundlesCalls())
func (mock *APIMock) SearchCheckBundlesCalls() []struct {
	SearchCriteria *apiclient.SearchQueryType
	FilterCriteria *apiclient.SearchFilterType
} {
	var calls []struct {
		SearchCriteria *apiclient.SearchQueryType
		FilterCriteria *apiclient.SearchFilterType
	}
	lockAPIMockSearchCheckBundles.RLock()
	calls = mock.calls.SearchCheckBundles
	lockAPIMockSearchCheckBundles.RUnlock()
	return calls
}

// UpdateCheckBundle calls UpdateCheckBundleFunc.
func (mock *APIMock) UpdateCheckBundle(cfg *apiclient.CheckBundle) (*apiclient.CheckBundle, error) {
	if mock.UpdateCheckBundleFunc == nil {
		panic("moq: APIMock.UpdateCheckBundleFunc is nil but API.UpdateCheckBundle was just called")
	}
	callInfo := struct {
		Cfg *apiclient.CheckBundle
	}{
		Cfg: cfg,
	}
	lockAPIMockUpdateCheckBundle.Lock()
	mock.calls.UpdateCheckBundle = append(mock.calls.UpdateCheckBundle, callInfo)
	lockAPIMockUpdateCheckBundle.Unlock()
	return mock.UpdateCheckBundleFunc(cfg)
}

// UpdateCheckBundleCalls gets all the calls that were made to UpdateCheckBundle.
// Check the length with:
//     len(mockedAPI.UpdateCheckBundleCalls())
func (mock *APIMock) UpdateCheckBundleCalls() []struct {
	Cfg *apiclient.CheckBundle
} {
	var calls []struct {
		Cfg *apiclient.CheckBundle
	}
	lockAPIMockUpdateCheckBundle.RLock()
	calls = mock.calls.UpdateCheckBundle
	lockAPIMockUpdateCheckBundle.RUnlock()
	return calls
}

// UpdateCheckBundleMetrics calls UpdateCheckBundleMetricsFunc.
func (mock *APIMock) UpdateCheckBundleMetrics(cfg *apiclient.CheckBundleMetrics) (*apiclient.CheckBundleMetrics, error) {
	if mock.UpdateCheckBundleMetricsFunc == nil {
		panic("moq: APIMock.UpdateCheckBundleMetricsFunc is nil but API.UpdateCheckBundleMetrics was just called")
	}
	callInfo := struct {
		Cfg *apiclient.CheckBundleMetrics
	}{
		Cfg: cfg,
	}
	lockAPIMockUpdateCheckBundleMetrics.Lock()
	mock.calls.UpdateCheckBundleMetrics = append(mock.calls.UpdateCheckBundleMetrics, callInfo)
	lockAPIMockUpdateCheckBundleMetrics.Unlock()
	return mock.UpdateCheckBundleMetricsFunc(cfg)
}

// UpdateCheckBundleMetricsCalls gets all the calls that were made to UpdateCheckBundleMetrics.
// Check the length with:
//     len(mockedAPI.UpdateCheckBundleMetricsCalls())
func (mock *APIMock) UpdateCheckBundleMetricsCalls() []struct {
	Cfg *apiclient.CheckBundleMetrics
} {
	var calls []struct {
		Cfg *apiclient.CheckBundleMetrics
	}
	lockAPIMockUpdateCheckBundleMetrics.RLock()
	calls = mock.calls.UpdateCheckBundleMetrics
	lockAPIMockUpdateCheckBundleMetrics.RUnlock()
	return calls
}
