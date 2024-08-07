package mts

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

// SubmitCopyrightExtractJob invokes the mts.SubmitCopyrightExtractJob API synchronously
func (client *Client) SubmitCopyrightExtractJob(request *SubmitCopyrightExtractJobRequest) (response *SubmitCopyrightExtractJobResponse, err error) {
	response = CreateSubmitCopyrightExtractJobResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitCopyrightExtractJobWithChan invokes the mts.SubmitCopyrightExtractJob API asynchronously
func (client *Client) SubmitCopyrightExtractJobWithChan(request *SubmitCopyrightExtractJobRequest) (<-chan *SubmitCopyrightExtractJobResponse, <-chan error) {
	responseChan := make(chan *SubmitCopyrightExtractJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitCopyrightExtractJob(request)
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

// SubmitCopyrightExtractJobWithCallback invokes the mts.SubmitCopyrightExtractJob API asynchronously
func (client *Client) SubmitCopyrightExtractJobWithCallback(request *SubmitCopyrightExtractJobRequest, callback func(response *SubmitCopyrightExtractJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitCopyrightExtractJobResponse
		var err error
		defer close(result)
		response, err = client.SubmitCopyrightExtractJob(request)
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

// SubmitCopyrightExtractJobRequest is the request struct for api SubmitCopyrightExtractJob
type SubmitCopyrightExtractJobRequest struct {
	*requests.RpcRequest
	Params   string `position:"Query" name:"Params"`
	Url      string `position:"Query" name:"Url"`
	Input    string `position:"Query" name:"Input"`
	UserData string `position:"Query" name:"UserData"`
	CallBack string `position:"Query" name:"CallBack"`
}

// SubmitCopyrightExtractJobResponse is the response struct for api SubmitCopyrightExtractJob
type SubmitCopyrightExtractJobResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Message    string `json:"Message" xml:"Message"`
	StatusCode int64  `json:"StatusCode" xml:"StatusCode"`
	Data       Data   `json:"Data" xml:"Data"`
}

// CreateSubmitCopyrightExtractJobRequest creates a request to invoke SubmitCopyrightExtractJob API
func CreateSubmitCopyrightExtractJobRequest() (request *SubmitCopyrightExtractJobRequest) {
	request = &SubmitCopyrightExtractJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "SubmitCopyrightExtractJob", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitCopyrightExtractJobResponse creates a response to parse from SubmitCopyrightExtractJob response
func CreateSubmitCopyrightExtractJobResponse() (response *SubmitCopyrightExtractJobResponse) {
	response = &SubmitCopyrightExtractJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
