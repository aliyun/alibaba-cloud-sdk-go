package safe

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

// QueryRegionAz invokes the safe.QueryRegionAz API synchronously
func (client *Client) QueryRegionAz(request *QueryRegionAzRequest) (response *QueryRegionAzResponse, err error) {
	response = CreateQueryRegionAzResponse()
	err = client.DoAction(request, response)
	return
}

// QueryRegionAzWithChan invokes the safe.QueryRegionAz API asynchronously
func (client *Client) QueryRegionAzWithChan(request *QueryRegionAzRequest) (<-chan *QueryRegionAzResponse, <-chan error) {
	responseChan := make(chan *QueryRegionAzResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryRegionAz(request)
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

// QueryRegionAzWithCallback invokes the safe.QueryRegionAz API asynchronously
func (client *Client) QueryRegionAzWithCallback(request *QueryRegionAzRequest, callback func(response *QueryRegionAzResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryRegionAzResponse
		var err error
		defer close(result)
		response, err = client.QueryRegionAz(request)
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

// QueryRegionAzRequest is the request struct for api QueryRegionAz
type QueryRegionAzRequest struct {
	*requests.RpcRequest
	AuthKey      string           `position:"Body" name:"AuthKey"`
	ReqTimestamp requests.Integer `position:"Body" name:"ReqTimestamp"`
	Product      string           `position:"Body" name:"Product"`
	Limit        requests.Integer `position:"Body" name:"Limit"`
	AuthSign     string           `position:"Body" name:"AuthSign"`
	Page         requests.Integer `position:"Body" name:"Page"`
}

// QueryRegionAzResponse is the response struct for api QueryRegionAz
type QueryRegionAzResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryRegionAzRequest creates a request to invoke QueryRegionAz API
func CreateQueryRegionAzRequest() (request *QueryRegionAzRequest) {
	request = &QueryRegionAzRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Safe", "2022-01-17", "QueryRegionAz", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryRegionAzResponse creates a response to parse from QueryRegionAz response
func CreateQueryRegionAzResponse() (response *QueryRegionAzResponse) {
	response = &QueryRegionAzResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
