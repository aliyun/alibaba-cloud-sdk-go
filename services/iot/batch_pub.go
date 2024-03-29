package iot

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

// BatchPub invokes the iot.BatchPub API synchronously
func (client *Client) BatchPub(request *BatchPubRequest) (response *BatchPubResponse, err error) {
	response = CreateBatchPubResponse()
	err = client.DoAction(request, response)
	return
}

// BatchPubWithChan invokes the iot.BatchPub API asynchronously
func (client *Client) BatchPubWithChan(request *BatchPubRequest) (<-chan *BatchPubResponse, <-chan error) {
	responseChan := make(chan *BatchPubResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchPub(request)
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

// BatchPubWithCallback invokes the iot.BatchPub API asynchronously
func (client *Client) BatchPubWithCallback(request *BatchPubRequest, callback func(response *BatchPubResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchPubResponse
		var err error
		defer close(result)
		response, err = client.BatchPub(request)
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

// BatchPubRequest is the request struct for api BatchPub
type BatchPubRequest struct {
	*requests.RpcRequest
	UserProp                  *[]BatchPubUserProp `position:"Query" name:"UserProp"  type:"Repeated"`
	ResponseTopicTemplateName string              `position:"Query" name:"ResponseTopicTemplateName"`
	MessageContent            string              `position:"Query" name:"MessageContent"`
	TopicTemplateName         string              `position:"Query" name:"TopicTemplateName"`
	Qos                       requests.Integer    `position:"Query" name:"Qos"`
	CorrelationData           string              `position:"Query" name:"CorrelationData"`
	IotInstanceId             string              `position:"Query" name:"IotInstanceId"`
	MessageExpiryInterval     requests.Integer    `position:"Query" name:"MessageExpiryInterval"`
	TopicShortName            string              `position:"Query" name:"TopicShortName"`
	PayloadFormatIndicator    requests.Integer    `position:"Query" name:"PayloadFormatIndicator"`
	ProductKey                string              `position:"Query" name:"ProductKey"`
	ContentType               string              `position:"Query" name:"ContentType"`
	Retained                  requests.Boolean    `position:"Query" name:"Retained"`
	ApiProduct                string              `position:"Body" name:"ApiProduct"`
	ApiRevision               string              `position:"Body" name:"ApiRevision"`
	DeviceName                *[]string           `position:"Query" name:"DeviceName"  type:"Repeated"`
}

// BatchPubUserProp is a repeated param struct in BatchPubRequest
type BatchPubUserProp struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// BatchPubResponse is the response struct for api BatchPub
type BatchPubResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateBatchPubRequest creates a request to invoke BatchPub API
func CreateBatchPubRequest() (request *BatchPubRequest) {
	request = &BatchPubRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "BatchPub", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBatchPubResponse creates a response to parse from BatchPub response
func CreateBatchPubResponse() (response *BatchPubResponse) {
	response = &BatchPubResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
