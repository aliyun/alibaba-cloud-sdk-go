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

// DescribeLiveGrtnDuration invokes the live.DescribeLiveGrtnDuration API synchronously
func (client *Client) DescribeLiveGrtnDuration(request *DescribeLiveGrtnDurationRequest) (response *DescribeLiveGrtnDurationResponse, err error) {
	response = CreateDescribeLiveGrtnDurationResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveGrtnDurationWithChan invokes the live.DescribeLiveGrtnDuration API asynchronously
func (client *Client) DescribeLiveGrtnDurationWithChan(request *DescribeLiveGrtnDurationRequest) (<-chan *DescribeLiveGrtnDurationResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveGrtnDurationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveGrtnDuration(request)
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

// DescribeLiveGrtnDurationWithCallback invokes the live.DescribeLiveGrtnDuration API asynchronously
func (client *Client) DescribeLiveGrtnDurationWithCallback(request *DescribeLiveGrtnDurationRequest, callback func(response *DescribeLiveGrtnDurationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveGrtnDurationResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveGrtnDuration(request)
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

// DescribeLiveGrtnDurationRequest is the request struct for api DescribeLiveGrtnDuration
type DescribeLiveGrtnDurationRequest struct {
	*requests.RpcRequest
	StartTime string           `position:"Query" name:"StartTime"`
	Area      string           `position:"Query" name:"Area"`
	EndTime   string           `position:"Query" name:"EndTime"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
	AppId     string           `position:"Query" name:"AppId"`
	Interval  string           `position:"Query" name:"Interval"`
}

// DescribeLiveGrtnDurationResponse is the response struct for api DescribeLiveGrtnDuration
type DescribeLiveGrtnDurationResponse struct {
	*responses.BaseResponse
	RequestId        string                                     `json:"RequestId" xml:"RequestId"`
	StreamDetailData StreamDetailDataInDescribeLiveGrtnDuration `json:"StreamDetailData" xml:"StreamDetailData"`
}

// CreateDescribeLiveGrtnDurationRequest creates a request to invoke DescribeLiveGrtnDuration API
func CreateDescribeLiveGrtnDurationRequest() (request *DescribeLiveGrtnDurationRequest) {
	request = &DescribeLiveGrtnDurationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeLiveGrtnDuration", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeLiveGrtnDurationResponse creates a response to parse from DescribeLiveGrtnDuration response
func CreateDescribeLiveGrtnDurationResponse() (response *DescribeLiveGrtnDurationResponse) {
	response = &DescribeLiveGrtnDurationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
