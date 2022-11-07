package arms

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

// GetPrometheusGlobalView invokes the arms.GetPrometheusGlobalView API synchronously
func (client *Client) GetPrometheusGlobalView(request *GetPrometheusGlobalViewRequest) (response *GetPrometheusGlobalViewResponse, err error) {
	response = CreateGetPrometheusGlobalViewResponse()
	err = client.DoAction(request, response)
	return
}

// GetPrometheusGlobalViewWithChan invokes the arms.GetPrometheusGlobalView API asynchronously
func (client *Client) GetPrometheusGlobalViewWithChan(request *GetPrometheusGlobalViewRequest) (<-chan *GetPrometheusGlobalViewResponse, <-chan error) {
	responseChan := make(chan *GetPrometheusGlobalViewResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPrometheusGlobalView(request)
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

// GetPrometheusGlobalViewWithCallback invokes the arms.GetPrometheusGlobalView API asynchronously
func (client *Client) GetPrometheusGlobalViewWithCallback(request *GetPrometheusGlobalViewRequest, callback func(response *GetPrometheusGlobalViewResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPrometheusGlobalViewResponse
		var err error
		defer close(result)
		response, err = client.GetPrometheusGlobalView(request)
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

// GetPrometheusGlobalViewRequest is the request struct for api GetPrometheusGlobalView
type GetPrometheusGlobalViewRequest struct {
	*requests.RpcRequest
	GlobalViewClusterId string `position:"Query" name:"GlobalViewClusterId"`
}

// GetPrometheusGlobalViewResponse is the response struct for api GetPrometheusGlobalView
type GetPrometheusGlobalViewResponse struct {
	*responses.BaseResponse
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateGetPrometheusGlobalViewRequest creates a request to invoke GetPrometheusGlobalView API
func CreateGetPrometheusGlobalViewRequest() (request *GetPrometheusGlobalViewRequest) {
	request = &GetPrometheusGlobalViewRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "GetPrometheusGlobalView", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetPrometheusGlobalViewResponse creates a response to parse from GetPrometheusGlobalView response
func CreateGetPrometheusGlobalViewResponse() (response *GetPrometheusGlobalViewResponse) {
	response = &GetPrometheusGlobalViewResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
