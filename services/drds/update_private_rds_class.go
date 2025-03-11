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

// UpdatePrivateRdsClass invokes the drds.UpdatePrivateRdsClass API synchronously
func (client *Client) UpdatePrivateRdsClass(request *UpdatePrivateRdsClassRequest) (response *UpdatePrivateRdsClassResponse, err error) {
	response = CreateUpdatePrivateRdsClassResponse()
	err = client.DoAction(request, response)
	return
}

// UpdatePrivateRdsClassWithChan invokes the drds.UpdatePrivateRdsClass API asynchronously
func (client *Client) UpdatePrivateRdsClassWithChan(request *UpdatePrivateRdsClassRequest) (<-chan *UpdatePrivateRdsClassResponse, <-chan error) {
	responseChan := make(chan *UpdatePrivateRdsClassResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdatePrivateRdsClass(request)
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

// UpdatePrivateRdsClassWithCallback invokes the drds.UpdatePrivateRdsClass API asynchronously
func (client *Client) UpdatePrivateRdsClassWithCallback(request *UpdatePrivateRdsClassRequest, callback func(response *UpdatePrivateRdsClassResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdatePrivateRdsClassResponse
		var err error
		defer close(result)
		response, err = client.UpdatePrivateRdsClass(request)
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

// UpdatePrivateRdsClassRequest is the request struct for api UpdatePrivateRdsClass
type UpdatePrivateRdsClassRequest struct {
	*requests.RpcRequest
	Storage        string           `position:"Query" name:"Storage"`
	AutoUseCoupon  requests.Boolean `position:"Query" name:"AutoUseCoupon"`
	DrdsInstanceId string           `position:"Query" name:"DrdsInstanceId"`
	RdsClass       string           `position:"Query" name:"RdsClass"`
	PrePayDuration requests.Integer `position:"Query" name:"PrePayDuration"`
	DBInstanceId   string           `position:"Query" name:"DBInstanceId"`
}

// UpdatePrivateRdsClassResponse is the response struct for api UpdatePrivateRdsClass
type UpdatePrivateRdsClassResponse struct {
	*responses.BaseResponse
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUpdatePrivateRdsClassRequest creates a request to invoke UpdatePrivateRdsClass API
func CreateUpdatePrivateRdsClassRequest() (request *UpdatePrivateRdsClassRequest) {
	request = &UpdatePrivateRdsClassRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "UpdatePrivateRdsClass", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdatePrivateRdsClassResponse creates a response to parse from UpdatePrivateRdsClass response
func CreateUpdatePrivateRdsClassResponse() (response *UpdatePrivateRdsClassResponse) {
	response = &UpdatePrivateRdsClassResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
