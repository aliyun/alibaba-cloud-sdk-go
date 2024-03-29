package linkvisual

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

// UpdateRtmpKey invokes the linkvisual.UpdateRtmpKey API synchronously
func (client *Client) UpdateRtmpKey(request *UpdateRtmpKeyRequest) (response *UpdateRtmpKeyResponse, err error) {
	response = CreateUpdateRtmpKeyResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateRtmpKeyWithChan invokes the linkvisual.UpdateRtmpKey API asynchronously
func (client *Client) UpdateRtmpKeyWithChan(request *UpdateRtmpKeyRequest) (<-chan *UpdateRtmpKeyResponse, <-chan error) {
	responseChan := make(chan *UpdateRtmpKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateRtmpKey(request)
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

// UpdateRtmpKeyWithCallback invokes the linkvisual.UpdateRtmpKey API asynchronously
func (client *Client) UpdateRtmpKeyWithCallback(request *UpdateRtmpKeyRequest, callback func(response *UpdateRtmpKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateRtmpKeyResponse
		var err error
		defer close(result)
		response, err = client.UpdateRtmpKey(request)
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

// UpdateRtmpKeyRequest is the request struct for api UpdateRtmpKey
type UpdateRtmpKeyRequest struct {
	*requests.RpcRequest
	PushAuthKey       string           `position:"Query" name:"PushAuthKey"`
	IotId             string           `position:"Query" name:"IotId"`
	IotInstanceId     string           `position:"Query" name:"IotInstanceId"`
	PushKeyExpireTime requests.Integer `position:"Query" name:"PushKeyExpireTime"`
	PullAuthKey       string           `position:"Query" name:"PullAuthKey"`
	ProductKey        string           `position:"Query" name:"ProductKey"`
	ApiProduct        string           `position:"Body" name:"ApiProduct"`
	ApiRevision       string           `position:"Body" name:"ApiRevision"`
	DeviceName        string           `position:"Query" name:"DeviceName"`
	PullKeyExpireTime requests.Integer `position:"Query" name:"PullKeyExpireTime"`
}

// UpdateRtmpKeyResponse is the response struct for api UpdateRtmpKey
type UpdateRtmpKeyResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateUpdateRtmpKeyRequest creates a request to invoke UpdateRtmpKey API
func CreateUpdateRtmpKeyRequest() (request *UpdateRtmpKeyRequest) {
	request = &UpdateRtmpKeyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "UpdateRtmpKey", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateRtmpKeyResponse creates a response to parse from UpdateRtmpKey response
func CreateUpdateRtmpKeyResponse() (response *UpdateRtmpKeyResponse) {
	response = &UpdateRtmpKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
