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

// UpdateLiveSnapshotNotifyConfig invokes the live.UpdateLiveSnapshotNotifyConfig API synchronously
func (client *Client) UpdateLiveSnapshotNotifyConfig(request *UpdateLiveSnapshotNotifyConfigRequest) (response *UpdateLiveSnapshotNotifyConfigResponse, err error) {
	response = CreateUpdateLiveSnapshotNotifyConfigResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateLiveSnapshotNotifyConfigWithChan invokes the live.UpdateLiveSnapshotNotifyConfig API asynchronously
func (client *Client) UpdateLiveSnapshotNotifyConfigWithChan(request *UpdateLiveSnapshotNotifyConfigRequest) (<-chan *UpdateLiveSnapshotNotifyConfigResponse, <-chan error) {
	responseChan := make(chan *UpdateLiveSnapshotNotifyConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateLiveSnapshotNotifyConfig(request)
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

// UpdateLiveSnapshotNotifyConfigWithCallback invokes the live.UpdateLiveSnapshotNotifyConfig API asynchronously
func (client *Client) UpdateLiveSnapshotNotifyConfigWithCallback(request *UpdateLiveSnapshotNotifyConfigRequest, callback func(response *UpdateLiveSnapshotNotifyConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateLiveSnapshotNotifyConfigResponse
		var err error
		defer close(result)
		response, err = client.UpdateLiveSnapshotNotifyConfig(request)
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

// UpdateLiveSnapshotNotifyConfigRequest is the request struct for api UpdateLiveSnapshotNotifyConfig
type UpdateLiveSnapshotNotifyConfigRequest struct {
	*requests.RpcRequest
	NotifyReqAuth string           `position:"Query" name:"NotifyReqAuth"`
	NotifyUrl     string           `position:"Query" name:"NotifyUrl"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	NotifyAuthKey string           `position:"Query" name:"NotifyAuthKey"`
}

// UpdateLiveSnapshotNotifyConfigResponse is the response struct for api UpdateLiveSnapshotNotifyConfig
type UpdateLiveSnapshotNotifyConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateLiveSnapshotNotifyConfigRequest creates a request to invoke UpdateLiveSnapshotNotifyConfig API
func CreateUpdateLiveSnapshotNotifyConfigRequest() (request *UpdateLiveSnapshotNotifyConfigRequest) {
	request = &UpdateLiveSnapshotNotifyConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "UpdateLiveSnapshotNotifyConfig", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateLiveSnapshotNotifyConfigResponse creates a response to parse from UpdateLiveSnapshotNotifyConfig response
func CreateUpdateLiveSnapshotNotifyConfigResponse() (response *UpdateLiveSnapshotNotifyConfigResponse) {
	response = &UpdateLiveSnapshotNotifyConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
