package drds

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

// DescribeDrdsInstance invokes the drds.DescribeDrdsInstance API synchronously
func (client *Client) DescribeDrdsInstance(request *DescribeDrdsInstanceRequest) (response *DescribeDrdsInstanceResponse, err error) {
	response = CreateDescribeDrdsInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDrdsInstanceWithChan invokes the drds.DescribeDrdsInstance API asynchronously
func (client *Client) DescribeDrdsInstanceWithChan(request *DescribeDrdsInstanceRequest) (<-chan *DescribeDrdsInstanceResponse, <-chan error) {
	responseChan := make(chan *DescribeDrdsInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDrdsInstance(request)
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

// DescribeDrdsInstanceWithCallback invokes the drds.DescribeDrdsInstance API asynchronously
func (client *Client) DescribeDrdsInstanceWithCallback(request *DescribeDrdsInstanceRequest, callback func(response *DescribeDrdsInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDrdsInstanceResponse
		var err error
		defer close(result)
		response, err = client.DescribeDrdsInstance(request)
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

// DescribeDrdsInstanceRequest is the request struct for api DescribeDrdsInstance
type DescribeDrdsInstanceRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

// DescribeDrdsInstanceResponse is the response struct for api DescribeDrdsInstance
type DescribeDrdsInstanceResponse struct {
	*responses.BaseResponse
	Success   bool                       `json:"Success" xml:"Success"`
	RequestId string                     `json:"RequestId" xml:"RequestId"`
	Data      DataInDescribeDrdsInstance `json:"Data" xml:"Data"`
}

// CreateDescribeDrdsInstanceRequest creates a request to invoke DescribeDrdsInstance API
func CreateDescribeDrdsInstanceRequest() (request *DescribeDrdsInstanceRequest) {
	request = &DescribeDrdsInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "DescribeDrdsInstance", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDrdsInstanceResponse creates a response to parse from DescribeDrdsInstance response
func CreateDescribeDrdsInstanceResponse() (response *DescribeDrdsInstanceResponse) {
	response = &DescribeDrdsInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
