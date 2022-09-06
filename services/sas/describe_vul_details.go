package sas

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

// DescribeVulDetails invokes the sas.DescribeVulDetails API synchronously
func (client *Client) DescribeVulDetails(request *DescribeVulDetailsRequest) (response *DescribeVulDetailsResponse, err error) {
	response = CreateDescribeVulDetailsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVulDetailsWithChan invokes the sas.DescribeVulDetails API asynchronously
func (client *Client) DescribeVulDetailsWithChan(request *DescribeVulDetailsRequest) (<-chan *DescribeVulDetailsResponse, <-chan error) {
	responseChan := make(chan *DescribeVulDetailsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVulDetails(request)
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

// DescribeVulDetailsWithCallback invokes the sas.DescribeVulDetails API asynchronously
func (client *Client) DescribeVulDetailsWithCallback(request *DescribeVulDetailsRequest, callback func(response *DescribeVulDetailsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVulDetailsResponse
		var err error
		defer close(result)
		response, err = client.DescribeVulDetails(request)
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

// DescribeVulDetailsRequest is the request struct for api DescribeVulDetails
type DescribeVulDetailsRequest struct {
	*requests.RpcRequest
	Type      string `position:"Query" name:"Type"`
	AliasName string `position:"Query" name:"AliasName"`
	SourceIp  string `position:"Query" name:"SourceIp"`
	Name      string `position:"Query" name:"Name"`
	Lang      string `position:"Query" name:"Lang"`
}

// DescribeVulDetailsResponse is the response struct for api DescribeVulDetails
type DescribeVulDetailsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Cves      []Cve  `json:"Cves" xml:"Cves"`
}

// CreateDescribeVulDetailsRequest creates a request to invoke DescribeVulDetails API
func CreateDescribeVulDetailsRequest() (request *DescribeVulDetailsRequest) {
	request = &DescribeVulDetailsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeVulDetails", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVulDetailsResponse creates a response to parse from DescribeVulDetails response
func CreateDescribeVulDetailsResponse() (response *DescribeVulDetailsResponse) {
	response = &DescribeVulDetailsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
