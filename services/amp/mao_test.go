package amp

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

// MaoTest invokes the amp.MaoTest API synchronously
func (client *Client) MaoTest(request *MaoTestRequest) (response *MaoTestResponse, err error) {
	response = CreateMaoTestResponse()
	err = client.DoAction(request, response)
	return
}

// MaoTestWithChan invokes the amp.MaoTest API asynchronously
func (client *Client) MaoTestWithChan(request *MaoTestRequest) (<-chan *MaoTestResponse, <-chan error) {
	responseChan := make(chan *MaoTestResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MaoTest(request)
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

// MaoTestWithCallback invokes the amp.MaoTest API asynchronously
func (client *Client) MaoTestWithCallback(request *MaoTestRequest, callback func(response *MaoTestResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MaoTestResponse
		var err error
		defer close(result)
		response, err = client.MaoTest(request)
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

// MaoTestRequest is the request struct for api MaoTest
type MaoTestRequest struct {
	*requests.RoaRequest
	AA string `position:"Query" name:"AA"`
}

// MaoTestResponse is the response struct for api MaoTest
type MaoTestResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
}

// CreateMaoTestRequest creates a request to invoke MaoTest API
func CreateMaoTestRequest() (request *MaoTestRequest) {
	request = &MaoTestRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("amp", "2020-07-08", "MaoTest", "/rmo/efe", "ServiceCode", "openAPI")
	request.Method = requests.POST
	return
}

// CreateMaoTestResponse creates a response to parse from MaoTest response
func CreateMaoTestResponse() (response *MaoTestResponse) {
	response = &MaoTestResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
