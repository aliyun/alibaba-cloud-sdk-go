package polardb

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

// ModifyDBNodeClass invokes the polardb.ModifyDBNodeClass API synchronously
func (client *Client) ModifyDBNodeClass(request *ModifyDBNodeClassRequest) (response *ModifyDBNodeClassResponse, err error) {
	response = CreateModifyDBNodeClassResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBNodeClassWithChan invokes the polardb.ModifyDBNodeClass API asynchronously
func (client *Client) ModifyDBNodeClassWithChan(request *ModifyDBNodeClassRequest) (<-chan *ModifyDBNodeClassResponse, <-chan error) {
	responseChan := make(chan *ModifyDBNodeClassResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBNodeClass(request)
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

// ModifyDBNodeClassWithCallback invokes the polardb.ModifyDBNodeClass API asynchronously
func (client *Client) ModifyDBNodeClassWithCallback(request *ModifyDBNodeClassRequest, callback func(response *ModifyDBNodeClassResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBNodeClassResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBNodeClass(request)
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

// ModifyDBNodeClassRequest is the request struct for api ModifyDBNodeClass
type ModifyDBNodeClassRequest struct {
	*requests.RpcRequest
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken            string           `position:"Query" name:"ClientToken"`
	PlannedEndTime         string           `position:"Query" name:"PlannedEndTime"`
	DBNodeType             string           `position:"Query" name:"DBNodeType"`
	DBNodeTargetClass      string           `position:"Query" name:"DBNodeTargetClass"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId            string           `position:"Query" name:"DBClusterId"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	PlannedFlashingOffTime string           `position:"Query" name:"PlannedFlashingOffTime"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
	PlannedStartTime       string           `position:"Query" name:"PlannedStartTime"`
	ModifyType             string           `position:"Query" name:"ModifyType"`
	SubCategory            string           `position:"Query" name:"SubCategory"`
}

// ModifyDBNodeClassResponse is the response struct for api ModifyDBNodeClass
type ModifyDBNodeClassResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	DBClusterId string `json:"DBClusterId" xml:"DBClusterId"`
	OrderId     string `json:"OrderId" xml:"OrderId"`
}

// CreateModifyDBNodeClassRequest creates a request to invoke ModifyDBNodeClass API
func CreateModifyDBNodeClassRequest() (request *ModifyDBNodeClassRequest) {
	request = &ModifyDBNodeClassRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "ModifyDBNodeClass", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDBNodeClassResponse creates a response to parse from ModifyDBNodeClass response
func CreateModifyDBNodeClassResponse() (response *ModifyDBNodeClassResponse) {
	response = &ModifyDBNodeClassResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
