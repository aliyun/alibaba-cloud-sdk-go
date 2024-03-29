package dds

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

// ModifyNodeSpec invokes the dds.ModifyNodeSpec API synchronously
func (client *Client) ModifyNodeSpec(request *ModifyNodeSpecRequest) (response *ModifyNodeSpecResponse, err error) {
	response = CreateModifyNodeSpecResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyNodeSpecWithChan invokes the dds.ModifyNodeSpec API asynchronously
func (client *Client) ModifyNodeSpecWithChan(request *ModifyNodeSpecRequest) (<-chan *ModifyNodeSpecResponse, <-chan error) {
	responseChan := make(chan *ModifyNodeSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyNodeSpec(request)
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

// ModifyNodeSpecWithCallback invokes the dds.ModifyNodeSpec API asynchronously
func (client *Client) ModifyNodeSpecWithCallback(request *ModifyNodeSpecRequest, callback func(response *ModifyNodeSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyNodeSpecResponse
		var err error
		defer close(result)
		response, err = client.ModifyNodeSpec(request)
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

// ModifyNodeSpecRequest is the request struct for api ModifyNodeSpec
type ModifyNodeSpecRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ReadonlyReplicas     requests.Integer `position:"Query" name:"ReadonlyReplicas"`
	CouponNo             string           `position:"Query" name:"CouponNo"`
	NodeClass            string           `position:"Query" name:"NodeClass"`
	EffectiveTime        string           `position:"Query" name:"EffectiveTime"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	SwitchTime           string           `position:"Query" name:"SwitchTime"`
	NodeId               string           `position:"Query" name:"NodeId"`
	BusinessInfo         string           `position:"Query" name:"BusinessInfo"`
	AutoPay              requests.Boolean `position:"Query" name:"AutoPay"`
	FromApp              string           `position:"Query" name:"FromApp"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	NodeStorage          requests.Integer `position:"Query" name:"NodeStorage"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	OrderType            string           `position:"Query" name:"OrderType"`
}

// ModifyNodeSpecResponse is the response struct for api ModifyNodeSpec
type ModifyNodeSpecResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	OrderId   string `json:"OrderId" xml:"OrderId"`
}

// CreateModifyNodeSpecRequest creates a request to invoke ModifyNodeSpec API
func CreateModifyNodeSpecRequest() (request *ModifyNodeSpecRequest) {
	request = &ModifyNodeSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "ModifyNodeSpec", "dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyNodeSpecResponse creates a response to parse from ModifyNodeSpec response
func CreateModifyNodeSpecResponse() (response *ModifyNodeSpecResponse) {
	response = &ModifyNodeSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
