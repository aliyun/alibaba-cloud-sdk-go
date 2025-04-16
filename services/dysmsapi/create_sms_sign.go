package dysmsapi

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

// CreateSmsSign invokes the dysmsapi.CreateSmsSign API synchronously
func (client *Client) CreateSmsSign(request *CreateSmsSignRequest) (response *CreateSmsSignResponse, err error) {
	response = CreateCreateSmsSignResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSmsSignWithChan invokes the dysmsapi.CreateSmsSign API asynchronously
func (client *Client) CreateSmsSignWithChan(request *CreateSmsSignRequest) (<-chan *CreateSmsSignResponse, <-chan error) {
	responseChan := make(chan *CreateSmsSignResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSmsSign(request)
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

// CreateSmsSignWithCallback invokes the dysmsapi.CreateSmsSign API asynchronously
func (client *Client) CreateSmsSignWithCallback(request *CreateSmsSignRequest, callback func(response *CreateSmsSignResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSmsSignResponse
		var err error
		defer close(result)
		response, err = client.CreateSmsSign(request)
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

// CreateSmsSignRequest is the request struct for api CreateSmsSign
type CreateSmsSignRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ApplySceneContent    string           `position:"Query" name:"ApplySceneContent"`
	MoreData             string           `position:"Query" name:"MoreData"`
	Remark               string           `position:"Query" name:"Remark"`
	ExtendFunction       string           `position:"Query" name:"ExtendFunction"`
	SignName             string           `position:"Query" name:"SignName"`
	ApplySource          string           `position:"Query" name:"ApplySource"`
	RouteName            string           `position:"Query" name:"RouteName"`
	QualificationId      requests.Integer `position:"Query" name:"QualificationId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ProdCode             string           `position:"Query" name:"ProdCode"`
	SignType             requests.Integer `position:"Query" name:"SignType"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ThirdParty           requests.Boolean `position:"Query" name:"ThirdParty"`
	SignSource           requests.Integer `position:"Query" name:"SignSource"`
}

// CreateSmsSignResponse is the response struct for api CreateSmsSign
type CreateSmsSignResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	SignName  string `json:"SignName" xml:"SignName"`
	OrderId   string `json:"OrderId" xml:"OrderId"`
}

// CreateCreateSmsSignRequest creates a request to invoke CreateSmsSign API
func CreateCreateSmsSignRequest() (request *CreateSmsSignRequest) {
	request = &CreateSmsSignRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dysmsapi", "2017-05-25", "CreateSmsSign", "dysms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateSmsSignResponse creates a response to parse from CreateSmsSign response
func CreateCreateSmsSignResponse() (response *CreateSmsSignResponse) {
	response = &CreateSmsSignResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
