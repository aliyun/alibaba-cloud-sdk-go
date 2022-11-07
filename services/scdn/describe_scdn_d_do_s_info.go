package scdn

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

// DescribeScdnDDoSInfo invokes the scdn.DescribeScdnDDoSInfo API synchronously
func (client *Client) DescribeScdnDDoSInfo(request *DescribeScdnDDoSInfoRequest) (response *DescribeScdnDDoSInfoResponse, err error) {
	response = CreateDescribeScdnDDoSInfoResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnDDoSInfoWithChan invokes the scdn.DescribeScdnDDoSInfo API asynchronously
func (client *Client) DescribeScdnDDoSInfoWithChan(request *DescribeScdnDDoSInfoRequest) (<-chan *DescribeScdnDDoSInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnDDoSInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnDDoSInfo(request)
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

// DescribeScdnDDoSInfoWithCallback invokes the scdn.DescribeScdnDDoSInfo API asynchronously
func (client *Client) DescribeScdnDDoSInfoWithCallback(request *DescribeScdnDDoSInfoRequest, callback func(response *DescribeScdnDDoSInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnDDoSInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnDDoSInfo(request)
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

// DescribeScdnDDoSInfoRequest is the request struct for api DescribeScdnDDoSInfo
type DescribeScdnDDoSInfoRequest struct {
	*requests.RpcRequest
}

// DescribeScdnDDoSInfoResponse is the response struct for api DescribeScdnDDoSInfo
type DescribeScdnDDoSInfoResponse struct {
	*responses.BaseResponse
	SecBandwidth     int    `json:"SecBandwidth" xml:"SecBandwidth"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
	ElasticBandwidth int    `json:"ElasticBandwidth" xml:"ElasticBandwidth"`
}

// CreateDescribeScdnDDoSInfoRequest creates a request to invoke DescribeScdnDDoSInfo API
func CreateDescribeScdnDDoSInfoRequest() (request *DescribeScdnDDoSInfoRequest) {
	request = &DescribeScdnDDoSInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnDDoSInfo", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeScdnDDoSInfoResponse creates a response to parse from DescribeScdnDDoSInfo response
func CreateDescribeScdnDDoSInfoResponse() (response *DescribeScdnDDoSInfoResponse) {
	response = &DescribeScdnDDoSInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
