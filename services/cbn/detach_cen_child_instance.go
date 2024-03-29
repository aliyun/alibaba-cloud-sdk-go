package cbn

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

// DetachCenChildInstance invokes the cbn.DetachCenChildInstance API synchronously
func (client *Client) DetachCenChildInstance(request *DetachCenChildInstanceRequest) (response *DetachCenChildInstanceResponse, err error) {
	response = CreateDetachCenChildInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// DetachCenChildInstanceWithChan invokes the cbn.DetachCenChildInstance API asynchronously
func (client *Client) DetachCenChildInstanceWithChan(request *DetachCenChildInstanceRequest) (<-chan *DetachCenChildInstanceResponse, <-chan error) {
	responseChan := make(chan *DetachCenChildInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachCenChildInstance(request)
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

// DetachCenChildInstanceWithCallback invokes the cbn.DetachCenChildInstance API asynchronously
func (client *Client) DetachCenChildInstanceWithCallback(request *DetachCenChildInstanceRequest, callback func(response *DetachCenChildInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachCenChildInstanceResponse
		var err error
		defer close(result)
		response, err = client.DetachCenChildInstance(request)
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

// DetachCenChildInstanceRequest is the request struct for api DetachCenChildInstance
type DetachCenChildInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CenId                 string           `position:"Query" name:"CenId"`
	CenOwnerId            requests.Integer `position:"Query" name:"CenOwnerId"`
	ChildInstanceRegionId string           `position:"Query" name:"ChildInstanceRegionId"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	ChildInstanceType     string           `position:"Query" name:"ChildInstanceType"`
	Version               string           `position:"Query" name:"Version"`
	ChildInstanceOwnerId  requests.Integer `position:"Query" name:"ChildInstanceOwnerId"`
	ChildInstanceId       string           `position:"Query" name:"ChildInstanceId"`
}

// DetachCenChildInstanceResponse is the response struct for api DetachCenChildInstance
type DetachCenChildInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDetachCenChildInstanceRequest creates a request to invoke DetachCenChildInstance API
func CreateDetachCenChildInstanceRequest() (request *DetachCenChildInstanceRequest) {
	request = &DetachCenChildInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "DetachCenChildInstance", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDetachCenChildInstanceResponse creates a response to parse from DetachCenChildInstance response
func CreateDetachCenChildInstanceResponse() (response *DetachCenChildInstanceResponse) {
	response = &DetachCenChildInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
