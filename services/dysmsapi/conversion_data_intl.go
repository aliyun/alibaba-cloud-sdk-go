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

// ConversionDataIntl invokes the dysmsapi.ConversionDataIntl API synchronously
func (client *Client) ConversionDataIntl(request *ConversionDataIntlRequest) (response *ConversionDataIntlResponse, err error) {
	response = CreateConversionDataIntlResponse()
	err = client.DoAction(request, response)
	return
}

// ConversionDataIntlWithChan invokes the dysmsapi.ConversionDataIntl API asynchronously
func (client *Client) ConversionDataIntlWithChan(request *ConversionDataIntlRequest) (<-chan *ConversionDataIntlResponse, <-chan error) {
	responseChan := make(chan *ConversionDataIntlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConversionDataIntl(request)
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

// ConversionDataIntlWithCallback invokes the dysmsapi.ConversionDataIntl API asynchronously
func (client *Client) ConversionDataIntlWithCallback(request *ConversionDataIntlRequest, callback func(response *ConversionDataIntlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConversionDataIntlResponse
		var err error
		defer close(result)
		response, err = client.ConversionDataIntl(request)
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

// ConversionDataIntlRequest is the request struct for api ConversionDataIntl
type ConversionDataIntlRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ExtendFunction       string           `position:"Query" name:"ExtendFunction"`
	RouteName            string           `position:"Query" name:"RouteName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ReportTime           requests.Integer `position:"Query" name:"ReportTime"`
	ConversionRate       string           `position:"Query" name:"ConversionRate"`
}

// ConversionDataIntlResponse is the response struct for api ConversionDataIntl
type ConversionDataIntlResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
}

// CreateConversionDataIntlRequest creates a request to invoke ConversionDataIntl API
func CreateConversionDataIntlRequest() (request *ConversionDataIntlRequest) {
	request = &ConversionDataIntlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dysmsapi", "2017-05-25", "ConversionDataIntl", "dysms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateConversionDataIntlResponse creates a response to parse from ConversionDataIntl response
func CreateConversionDataIntlResponse() (response *ConversionDataIntlResponse) {
	response = &ConversionDataIntlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
