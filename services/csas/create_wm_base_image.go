package csas

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

// CreateWmBaseImage invokes the csas.CreateWmBaseImage API synchronously
func (client *Client) CreateWmBaseImage(request *CreateWmBaseImageRequest) (response *CreateWmBaseImageResponse, err error) {
	response = CreateCreateWmBaseImageResponse()
	err = client.DoAction(request, response)
	return
}

// CreateWmBaseImageWithChan invokes the csas.CreateWmBaseImage API asynchronously
func (client *Client) CreateWmBaseImageWithChan(request *CreateWmBaseImageRequest) (<-chan *CreateWmBaseImageResponse, <-chan error) {
	responseChan := make(chan *CreateWmBaseImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateWmBaseImage(request)
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

// CreateWmBaseImageWithCallback invokes the csas.CreateWmBaseImage API asynchronously
func (client *Client) CreateWmBaseImageWithCallback(request *CreateWmBaseImageRequest, callback func(response *CreateWmBaseImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateWmBaseImageResponse
		var err error
		defer close(result)
		response, err = client.CreateWmBaseImage(request)
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

// CreateWmBaseImageRequest is the request struct for api CreateWmBaseImage
type CreateWmBaseImageRequest struct {
	*requests.RpcRequest
	Scale          requests.Integer `position:"Body" name:"Scale"`
	ApiType        string           `position:"Body" name:"ApiType"`
	WmInfoUint     string           `position:"Body" name:"WmInfoUint"`
	WmInfoSize     requests.Integer `position:"Body" name:"WmInfoSize"`
	WmInfoBytesB64 string           `position:"Body" name:"WmInfoBytesB64"`
	WmType         string           `position:"Body" name:"WmType"`
	Width          requests.Integer `position:"Body" name:"Width"`
	Opacity        requests.Integer `position:"Body" name:"Opacity"`
	Height         requests.Integer `position:"Body" name:"Height"`
}

// CreateWmBaseImageResponse is the response struct for api CreateWmBaseImage
type CreateWmBaseImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateWmBaseImageRequest creates a request to invoke CreateWmBaseImage API
func CreateCreateWmBaseImageRequest() (request *CreateWmBaseImageRequest) {
	request = &CreateWmBaseImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("csas", "2023-01-20", "CreateWmBaseImage", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateWmBaseImageResponse creates a response to parse from CreateWmBaseImage response
func CreateCreateWmBaseImageResponse() (response *CreateWmBaseImageResponse) {
	response = &CreateWmBaseImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
