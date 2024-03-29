package live

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

// ModifyLiveMessageAppAudit invokes the live.ModifyLiveMessageAppAudit API synchronously
func (client *Client) ModifyLiveMessageAppAudit(request *ModifyLiveMessageAppAuditRequest) (response *ModifyLiveMessageAppAuditResponse, err error) {
	response = CreateModifyLiveMessageAppAuditResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyLiveMessageAppAuditWithChan invokes the live.ModifyLiveMessageAppAudit API asynchronously
func (client *Client) ModifyLiveMessageAppAuditWithChan(request *ModifyLiveMessageAppAuditRequest) (<-chan *ModifyLiveMessageAppAuditResponse, <-chan error) {
	responseChan := make(chan *ModifyLiveMessageAppAuditResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyLiveMessageAppAudit(request)
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

// ModifyLiveMessageAppAuditWithCallback invokes the live.ModifyLiveMessageAppAudit API asynchronously
func (client *Client) ModifyLiveMessageAppAuditWithCallback(request *ModifyLiveMessageAppAuditRequest, callback func(response *ModifyLiveMessageAppAuditResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyLiveMessageAppAuditResponse
		var err error
		defer close(result)
		response, err = client.ModifyLiveMessageAppAudit(request)
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

// ModifyLiveMessageAppAuditRequest is the request struct for api ModifyLiveMessageAppAudit
type ModifyLiveMessageAppAuditRequest struct {
	*requests.RpcRequest
	DataCenter string           `position:"Query" name:"DataCenter"`
	AppId      string           `position:"Query" name:"AppId"`
	AuditType  requests.Integer `position:"Query" name:"AuditType"`
	AuditUrl   string           `position:"Query" name:"AuditUrl"`
}

// ModifyLiveMessageAppAuditResponse is the response struct for api ModifyLiveMessageAppAudit
type ModifyLiveMessageAppAuditResponse struct {
	*responses.BaseResponse
	RequestId               string `json:"RequestId" xml:"RequestId"`
	AppId                   string `json:"AppId" xml:"AppId"`
	AppSign                 string `json:"AppSign" xml:"AppSign"`
	AuditType               int    `json:"AuditType" xml:"AuditType"`
	AuditUrl                string `json:"AuditUrl" xml:"AuditUrl"`
	AuditNeedAuthentication bool   `json:"AuditNeedAuthentication" xml:"AuditNeedAuthentication"`
}

// CreateModifyLiveMessageAppAuditRequest creates a request to invoke ModifyLiveMessageAppAudit API
func CreateModifyLiveMessageAppAuditRequest() (request *ModifyLiveMessageAppAuditRequest) {
	request = &ModifyLiveMessageAppAuditRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "ModifyLiveMessageAppAudit", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyLiveMessageAppAuditResponse creates a response to parse from ModifyLiveMessageAppAudit response
func CreateModifyLiveMessageAppAuditResponse() (response *ModifyLiveMessageAppAuditResponse) {
	response = &ModifyLiveMessageAppAuditResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
