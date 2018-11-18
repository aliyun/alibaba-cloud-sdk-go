package endpoints

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func TestLocationResolver_GetName(t *testing.T) {
	resolver := &LocationResolver{}
	assert.Equal(t, "location resolver", resolver.GetName())
}

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

func TestLocationResolver_TryResolve_LocationWith404(t *testing.T) {
	resolver := &LocationResolver{}
	resolveParam := &ResolveParam{
		LocationProduct: "ecs",
		RegionId:        "cn-hangzhou",
		Product:         "ecs",
		CommonApi: func(request *requests.CommonRequest) (response *responses.CommonResponse, err error) {
			response = responses.NewCommonResponse()
			responses.Unmarshal(response, makeHTTPResponse(404, "content"), "JSON")
			return
		},
	}
	endpoint, support, err := resolver.TryResolve(resolveParam)
	assert.Equal(t, "", endpoint)
	assert.Equal(t, false, support)
	assert.Nil(t, err)
}

func TestLocationResolver_TryResolve_LocationWith200InvalidJSON(t *testing.T) {
	resolver := &LocationResolver{}
	resolveParam := &ResolveParam{
		LocationProduct: "ecs",
		RegionId:        "cn-hangzhou",
		Product:         "ecs",
		CommonApi: func(request *requests.CommonRequest) (response *responses.CommonResponse, err error) {
			response = responses.NewCommonResponse()
			responses.Unmarshal(response, makeHTTPResponse(200, "content"), "JSON")
			return
		},
	}
	endpoint, support, err := resolver.TryResolve(resolveParam)
	assert.Equal(t, "", endpoint)
	assert.Equal(t, false, support)
	assert.Equal(t, "invalid character 'c' looking for beginning of value", err.Error())
}

func TestLocationResolver_TryResolve_LocationWith200ValidJSON(t *testing.T) {
	resolver := &LocationResolver{}
	resolveParam := &ResolveParam{
		LocationProduct: "ecs",
		RegionId:        "cn-hangzhou",
		Product:         "ecs",
		CommonApi: func(request *requests.CommonRequest) (response *responses.CommonResponse, err error) {
			response = responses.NewCommonResponse()
			responses.Unmarshal(response, makeHTTPResponse(200, `{"Code":"Success","RequestId":"request id"}`), "JSON")
			return
		},
	}
	endpoint, support, err := resolver.TryResolve(resolveParam)
	assert.Equal(t, "", endpoint)
	assert.Equal(t, false, support)
	assert.Nil(t, err)
	// assert.Equal(t, "json: cannot unmarshal array into Go struct field GetEndpointResponse.Endpoints of type endpoints.EndpointsObj", err.Error())
}

func TestLocationResolver_TryResolve_LocationWith200(t *testing.T) {
	resolver := &LocationResolver{}
	resolveParam := &ResolveParam{
		LocationProduct: "ecs",
		RegionId:        "cn-hangzhou",
		Product:         "ecs",
		CommonApi: func(request *requests.CommonRequest) (response *responses.CommonResponse, err error) {
			response = responses.NewCommonResponse()
			responses.Unmarshal(response, makeHTTPResponse(200, `{"Success":true,"RequestId":"request id","Endpoints":{"Endpoint":[]}}`), "JSON")
			return
		},
	}
	endpoint, support, err := resolver.TryResolve(resolveParam)
	assert.Equal(t, "", endpoint)
	assert.Equal(t, false, support)
	assert.Nil(t, err)
}

func TestLocationResolver_TryResolve_Location_With_Endpoint(t *testing.T) {
	resolver := &LocationResolver{}
	resolveParam := &ResolveParam{
		LocationProduct: "ecs",
		RegionId:        "cn-hangzhou",
		Product:         "ecs",
		CommonApi: func(request *requests.CommonRequest) (response *responses.CommonResponse, err error) {
			response = responses.NewCommonResponse()
			responses.Unmarshal(response, makeHTTPResponse(200, `{"Success":true,"RequestId":"request id","Endpoints":{"Endpoint":[{"Endpoint":"domain.com"}]}}`), "JSON")
			return
		},
	}
	endpoint, support, err := resolver.TryResolve(resolveParam)
	assert.Equal(t, "domain.com", endpoint)
	assert.Equal(t, true, support)
	assert.Nil(t, err)
	// hit the cache
	endpoint, support, err = resolver.TryResolve(resolveParam)
	assert.Equal(t, "domain.com", endpoint)
	assert.Equal(t, true, support)
	assert.Nil(t, err)
}

func TestLocationResolver_TryResolve_Location_With_Endpoint2(t *testing.T) {
	resolver := &LocationResolver{}
	resolveParam := &ResolveParam{
		LocationProduct: "ecs3",
		RegionId:        "cn-hangzhou",
		Product:         "ecs3",
		CommonApi: func(request *requests.CommonRequest) (response *responses.CommonResponse, err error) {
			response = responses.NewCommonResponse()
			responses.Unmarshal(response, makeHTTPResponse(200, `{
  "Endpoints":{
    "Endpoint":[
      {
        "Protocols":{
          "Protocols":["HTTP","HTTPS"]
        },
        "Type":"openAPI",
        "Namespace":"26842",
        "Id":"cn-beijing",
        "SerivceCode":"ecs",
        "Endpoint":"ecs-cn-hangzhou.aliyuncs.com"
      }
    ]
  },
  "RequestId":"B3B36D8E-5029-42E3-B1FB-9B687F7591DA",
  "Success":true
}`), "JSON")
			return
		},
	}
	endpoint, support, err := resolver.TryResolve(resolveParam)
	assert.Equal(t, "ecs-cn-hangzhou.aliyuncs.com", endpoint)
	assert.Equal(t, true, support)
	assert.Nil(t, err)
	// hit the cache
	endpoint, support, err = resolver.TryResolve(resolveParam)
	assert.Equal(t, "ecs-cn-hangzhou.aliyuncs.com", endpoint)
	assert.Equal(t, true, support)
	assert.Nil(t, err)
}

func TestLocationResolver_TryResolve_Location_With_EmptyEndpoint(t *testing.T) {
	resolver := &LocationResolver{}
	resolveParam := &ResolveParam{
		LocationProduct: "ecs2",
		RegionId:        "cn-hangzhou",
		Product:         "ecs2",
		CommonApi: func(request *requests.CommonRequest) (response *responses.CommonResponse, err error) {
			response = responses.NewCommonResponse()
			responses.Unmarshal(response, makeHTTPResponse(200, `{"Success":true,"RequestId":"request id","Endpoints":{"Endpoint":[{"Endpoint":""}]}}`), "JSON")
			return
		},
	}
	endpoint, support, err := resolver.TryResolve(resolveParam)
	assert.Equal(t, "", endpoint)
	assert.Equal(t, false, support)
	assert.Nil(t, err)
}
