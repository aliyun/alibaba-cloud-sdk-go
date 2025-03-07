package cloudauth

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

// InsertWhiteListSetting invokes the cloudauth.InsertWhiteListSetting API synchronously
func (client *Client) InsertWhiteListSetting(request *InsertWhiteListSettingRequest) (response *InsertWhiteListSettingResponse, err error) {
	response = CreateInsertWhiteListSettingResponse()
	err = client.DoAction(request, response)
	return
}

// InsertWhiteListSettingWithChan invokes the cloudauth.InsertWhiteListSetting API asynchronously
func (client *Client) InsertWhiteListSettingWithChan(request *InsertWhiteListSettingRequest) (<-chan *InsertWhiteListSettingResponse, <-chan error) {
	responseChan := make(chan *InsertWhiteListSettingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InsertWhiteListSetting(request)
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

// InsertWhiteListSettingWithCallback invokes the cloudauth.InsertWhiteListSetting API asynchronously
func (client *Client) InsertWhiteListSettingWithCallback(request *InsertWhiteListSettingRequest, callback func(response *InsertWhiteListSettingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InsertWhiteListSettingResponse
		var err error
		defer close(result)
		response, err = client.InsertWhiteListSetting(request)
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

// InsertWhiteListSettingRequest is the request struct for api InsertWhiteListSetting
type InsertWhiteListSettingRequest struct {
	*requests.RpcRequest
	ValidDay    requests.Integer `position:"Query" name:"ValidDay"`
	Remark      string           `position:"Query" name:"Remark"`
	CertifyId   string           `position:"Query" name:"CertifyId"`
	CertNo      string           `position:"Query" name:"CertNo"`
	ServiceCode string           `position:"Query" name:"ServiceCode"`
	SceneId     requests.Integer `position:"Query" name:"SceneId"`
}

// InsertWhiteListSettingResponse is the response struct for api InsertWhiteListSetting
type InsertWhiteListSettingResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Code         string `json:"Code" xml:"Code"`
	Message      string `json:"Message" xml:"Message"`
	Success      bool   `json:"Success" xml:"Success"`
	ResultObject bool   `json:"ResultObject" xml:"ResultObject"`
}

// CreateInsertWhiteListSettingRequest creates a request to invoke InsertWhiteListSetting API
func CreateInsertWhiteListSettingRequest() (request *InsertWhiteListSettingRequest) {
	request = &InsertWhiteListSettingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "InsertWhiteListSetting", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateInsertWhiteListSettingResponse creates a response to parse from InsertWhiteListSetting response
func CreateInsertWhiteListSettingResponse() (response *InsertWhiteListSettingResponse) {
	response = &InsertWhiteListSettingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
