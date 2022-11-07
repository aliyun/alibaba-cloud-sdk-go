package cloudesl

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// QueryLayoutByDevice invokes the cloudesl.QueryLayoutByDevice API synchronously
func (client *Client) QueryLayoutByDevice(request *QueryLayoutByDeviceRequest) (response *QueryLayoutByDeviceResponse, err error) {
	response = CreateQueryLayoutByDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// QueryLayoutByDeviceWithChan invokes the cloudesl.QueryLayoutByDevice API asynchronously
func (client *Client) QueryLayoutByDeviceWithChan(request *QueryLayoutByDeviceRequest) (<-chan *QueryLayoutByDeviceResponse, <-chan error) {
	responseChan := make(chan *QueryLayoutByDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryLayoutByDevice(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// QueryLayoutByDeviceWithCallback invokes the cloudesl.QueryLayoutByDevice API asynchronously
func (client *Client) QueryLayoutByDeviceWithCallback(request *QueryLayoutByDeviceRequest, callback func(response *QueryLayoutByDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryLayoutByDeviceResponse
		var err error
		defer close(result)
		response, err = client.QueryLayoutByDevice(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// QueryLayoutByDeviceRequest is the request struct for api QueryLayoutByDevice
type QueryLayoutByDeviceRequest struct {
	*requests.RpcRequest
	StoreId         string           `position:"Body" name:"StoreId"`
	PageNumber      requests.Integer `position:"Body" name:"PageNumber"`
	TemplateVersion string           `position:"Body" name:"TemplateVersion"`
	EslBarCode      string           `position:"Body" name:"EslBarCode"`
	PageSize        requests.Integer `position:"Body" name:"PageSize"`
	EslModelId      string           `position:"Body" name:"EslModelId"`
}

// QueryLayoutByDeviceResponse is the response struct for api QueryLayoutByDevice
type QueryLayoutByDeviceResponse struct {
	*responses.BaseResponse
	RequestId      string   `json:"RequestId" xml:"RequestId"`
	Success        bool     `json:"Success" xml:"Success"`
	Message        string   `json:"Message" xml:"Message"`
	ErrorCode      string   `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string   `json:"ErrorMessage" xml:"ErrorMessage"`
	Code           string   `json:"Code" xml:"Code"`
	DynamicCode    string   `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string   `json:"DynamicMessage" xml:"DynamicMessage"`
	PageNumber     int      `json:"PageNumber" xml:"PageNumber"`
	PageSize       int      `json:"PageSize" xml:"PageSize"`
	TotalCount     int      `json:"TotalCount" xml:"TotalCount"`
	LayoutList     []Layout `json:"LayoutList" xml:"LayoutList"`
}

// CreateQueryLayoutByDeviceRequest creates a request to invoke QueryLayoutByDevice API
func CreateQueryLayoutByDeviceRequest() (request *QueryLayoutByDeviceRequest) {
	request = &QueryLayoutByDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "QueryLayoutByDevice", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryLayoutByDeviceResponse creates a response to parse from QueryLayoutByDevice response
func CreateQueryLayoutByDeviceResponse() (response *QueryLayoutByDeviceResponse) {
	response = &QueryLayoutByDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
