package endpoints

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/stretchr/testify/assert"
)

// cases from later commit
func TestLocationResolver_TryResolve_EmptyLocationProduct(t *testing.T) {
	resolver := &LocationResolver{}
	resolveParam := &ResolveParam{}
	endpoint, support, err := resolver.TryResolve(resolveParam)
	assert.Nil(t, err)
	assert.Equal(t, "", endpoint)
	assert.Equal(t, false, support)
}

func TestLocationResolver_TryResolve_LocationWithError(t *testing.T) {
	resolver := &LocationResolver{}
	resolveParam := &ResolveParam{
		LocationProduct: "ecs",
		RegionId:        "cn-hangzhou",
		Product:         "ecs",
		CommonApi: func(request *requests.CommonRequest) (response *responses.CommonResponse, err error) {
			err = errors.NewClientError("SDK.MockError", "Mock error", nil)
			return
		},
	}
	endpoint, support, err := resolver.TryResolve(resolveParam)
	assert.Equal(t, "", endpoint)
	assert.Equal(t, false, support)
	assert.Equal(t, "[SDK.MockError] Mock error", err.Error())
}

func TestLocationResolver_TryResolve_Location_With_Endpoint(t *testing.T) {

	endpoint, support, err := resovleSucc(0)
	assert.Equal(t, "domain.com", endpoint)
	assert.Equal(t, true, support)
	assert.Nil(t, err)

	// hit the cache
	endpoint, support, err = resovleSucc(0)
	assert.Equal(t, "domain.com", endpoint)
	assert.Equal(t, true, support)
	assert.Nil(t, err)
}

func makeHTTPResponse(statusCode int, content string) (res *http.Response) {
	header := make(http.Header)
	status := strconv.Itoa(statusCode)
	res = &http.Response{
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		Header:     header,
		StatusCode: statusCode,
		Status:     status + " " + http.StatusText(statusCode),
	}
	res.Header = make(http.Header)
	res.Body = ioutil.NopCloser(bytes.NewReader([]byte(content)))
	return
}

func resovleSucc(i int) (ep string, isSupport bool, err error) {
	resolver := &LocationResolver{}
	resolveParam := &ResolveParam{
		LocationProduct: "ecs",
		RegionId:        fmt.Sprintf("cn-hangzhou%d", i),
		Product:         "ecs",
		CommonApi: func(request *requests.CommonRequest) (response *responses.CommonResponse, err error) {
			response = responses.NewCommonResponse()
			responses.Unmarshal(response, makeHTTPResponse(200, `{"Success":true,"RequestId":"request id","Endpoints":{"Endpoint":[{"Endpoint":"domain.com"}]}}`), "JSON")
			return
		},
	}
	endpoint, support, err := resolver.TryResolve(resolveParam)
	return endpoint, support, err
}

// concurrent cases
func TestResolveConcurrent(t *testing.T) {

	cnt := 50
	var wg sync.WaitGroup
	for i := 0; i < cnt; i++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			cachedKey := fmt.Sprintf("ecs#cn-hangzhou%d", k)
			for j := 0; j < 50; j++ {
				endpoint, support, err := resovleSucc(k)
				assert.Equal(t, "domain.com", endpointCache.Get(cachedKey))
				assert.Equal(t, "domain.com", endpoint)
				assert.Equal(t, true, support)
				assert.Nil(t, err)
			}
		}(i)
	}
	wg.Wait()
	assert.Equal(t, 50, len(endpointCache.cache))

	// hit cache and concurrent get
	for i := 0; i < cnt; i++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			cachedKey := fmt.Sprintf("ecs#cn-hangzhou%d", k)
			for j := 0; j < cnt; j++ {
				assert.Equal(t, "domain.com", endpointCache.Get(cachedKey))
				endpoint, support, err := resovleSucc(k)
				assert.Equal(t, "domain.com", endpoint)
				assert.Equal(t, true, support)
				assert.Nil(t, err)
			}
		}(i)
		wg.Wait()
	}
}
