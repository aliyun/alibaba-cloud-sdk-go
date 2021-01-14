package cdrs

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

// DetectTrajectoryRegularPattern invokes the cdrs.DetectTrajectoryRegularPattern API synchronously
func (client *Client) DetectTrajectoryRegularPattern(request *DetectTrajectoryRegularPatternRequest) (response *DetectTrajectoryRegularPatternResponse, err error) {
	response = CreateDetectTrajectoryRegularPatternResponse()
	err = client.DoAction(request, response)
	return
}

// DetectTrajectoryRegularPatternWithChan invokes the cdrs.DetectTrajectoryRegularPattern API asynchronously
func (client *Client) DetectTrajectoryRegularPatternWithChan(request *DetectTrajectoryRegularPatternRequest) (<-chan *DetectTrajectoryRegularPatternResponse, <-chan error) {
	responseChan := make(chan *DetectTrajectoryRegularPatternResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetectTrajectoryRegularPattern(request)
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

// DetectTrajectoryRegularPatternWithCallback invokes the cdrs.DetectTrajectoryRegularPattern API asynchronously
func (client *Client) DetectTrajectoryRegularPatternWithCallback(request *DetectTrajectoryRegularPatternRequest, callback func(response *DetectTrajectoryRegularPatternResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetectTrajectoryRegularPatternResponse
		var err error
		defer close(result)
		response, err = client.DetectTrajectoryRegularPattern(request)
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

// DetectTrajectoryRegularPatternRequest is the request struct for api DetectTrajectoryRegularPattern
type DetectTrajectoryRegularPatternRequest struct {
	*requests.RpcRequest
	PredictDate string `position:"Body" name:"PredictDate"`
	CorpId      string `position:"Body" name:"CorpId"`
	IdValue     string `position:"Body" name:"IdValue"`
	IdType      string `position:"Body" name:"IdType"`
}

// DetectTrajectoryRegularPatternResponse is the response struct for api DetectTrajectoryRegularPattern
type DetectTrajectoryRegularPatternResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Data      string `json:"Data" xml:"Data"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDetectTrajectoryRegularPatternRequest creates a request to invoke DetectTrajectoryRegularPattern API
func CreateDetectTrajectoryRegularPatternRequest() (request *DetectTrajectoryRegularPatternRequest) {
	request = &DetectTrajectoryRegularPatternRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CDRS", "2020-11-01", "DetectTrajectoryRegularPattern", "", "")
	request.Method = requests.POST
	return
}

// CreateDetectTrajectoryRegularPatternResponse creates a response to parse from DetectTrajectoryRegularPattern response
func CreateDetectTrajectoryRegularPatternResponse() (response *DetectTrajectoryRegularPatternResponse) {
	response = &DetectTrajectoryRegularPatternResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
