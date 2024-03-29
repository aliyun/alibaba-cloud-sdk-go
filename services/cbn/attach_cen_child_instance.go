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

// AttachCenChildInstance invokes the cbn.AttachCenChildInstance API synchronously
func (client *Client) AttachCenChildInstance(request *AttachCenChildInstanceRequest) (response *AttachCenChildInstanceResponse, err error) {
	response = CreateAttachCenChildInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// AttachCenChildInstanceWithChan invokes the cbn.AttachCenChildInstance API asynchronously
func (client *Client) AttachCenChildInstanceWithChan(request *AttachCenChildInstanceRequest) (<-chan *AttachCenChildInstanceResponse, <-chan error) {
	responseChan := make(chan *AttachCenChildInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AttachCenChildInstance(request)
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

// AttachCenChildInstanceWithCallback invokes the cbn.AttachCenChildInstance API asynchronously
func (client *Client) AttachCenChildInstanceWithCallback(request *AttachCenChildInstanceRequest, callback func(response *AttachCenChildInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AttachCenChildInstanceResponse
		var err error
		defer close(result)
		response, err = client.AttachCenChildInstance(request)
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

// AttachCenChildInstanceRequest is the request struct for api AttachCenChildInstance
type AttachCenChildInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CenId                 string           `position:"Query" name:"CenId"`
	ChildInstanceRegionId string           `position:"Query" name:"ChildInstanceRegionId"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	ChildInstanceType     string           `position:"Query" name:"ChildInstanceType"`
	Version               string           `position:"Query" name:"Version"`
	ChildInstanceOwnerId  requests.Integer `position:"Query" name:"ChildInstanceOwnerId"`
	ChildInstanceId       string           `position:"Query" name:"ChildInstanceId"`
}

// AttachCenChildInstanceResponse is the response struct for api AttachCenChildInstance
type AttachCenChildInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAttachCenChildInstanceRequest creates a request to invoke AttachCenChildInstance API
func CreateAttachCenChildInstanceRequest() (request *AttachCenChildInstanceRequest) {
	request = &AttachCenChildInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "AttachCenChildInstance", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAttachCenChildInstanceResponse creates a response to parse from AttachCenChildInstance response
func CreateAttachCenChildInstanceResponse() (response *AttachCenChildInstanceResponse) {
	response = &AttachCenChildInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
