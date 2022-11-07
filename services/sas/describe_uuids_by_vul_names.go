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

// DescribeUuidsByVulNames invokes the sas.DescribeUuidsByVulNames API synchronously
func (client *Client) DescribeUuidsByVulNames(request *DescribeUuidsByVulNamesRequest) (response *DescribeUuidsByVulNamesResponse, err error) {
	response = CreateDescribeUuidsByVulNamesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUuidsByVulNamesWithChan invokes the sas.DescribeUuidsByVulNames API asynchronously
func (client *Client) DescribeUuidsByVulNamesWithChan(request *DescribeUuidsByVulNamesRequest) (<-chan *DescribeUuidsByVulNamesResponse, <-chan error) {
	responseChan := make(chan *DescribeUuidsByVulNamesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUuidsByVulNames(request)
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

// DescribeUuidsByVulNamesWithCallback invokes the sas.DescribeUuidsByVulNames API asynchronously
func (client *Client) DescribeUuidsByVulNamesWithCallback(request *DescribeUuidsByVulNamesRequest, callback func(response *DescribeUuidsByVulNamesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUuidsByVulNamesResponse
		var err error
		defer close(result)
		response, err = client.DescribeUuidsByVulNames(request)
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

// DescribeUuidsByVulNamesRequest is the request struct for api DescribeUuidsByVulNames
type DescribeUuidsByVulNamesRequest struct {
	*requests.RpcRequest
	StatusList     string           `position:"Query" name:"StatusList"`
	TargetType     string           `position:"Query" name:"TargetType"`
	Remark         string           `position:"Query" name:"Remark"`
	Type           string           `position:"Query" name:"Type"`
	VpcInstanceIds string           `position:"Query" name:"VpcInstanceIds"`
	VulNames       *[]string        `position:"Query" name:"VulNames"  type:"Repeated"`
	SourceIp       string           `position:"Query" name:"SourceIp"`
	Tag            string           `position:"Query" name:"Tag"`
	Lang           string           `position:"Query" name:"Lang"`
	Level          string           `position:"Query" name:"Level"`
	GroupId        requests.Integer `position:"Query" name:"GroupId"`
	Dealed         string           `position:"Query" name:"Dealed"`
	FieldValue     string           `position:"Query" name:"FieldValue"`
	FieldName      string           `position:"Query" name:"FieldName"`
	SearchTags     string           `position:"Query" name:"SearchTags"`
	Necessity      string           `position:"Query" name:"Necessity"`
}

// DescribeUuidsByVulNamesResponse is the response struct for api DescribeUuidsByVulNames
type DescribeUuidsByVulNamesResponse struct {
	*responses.BaseResponse
	RequestId             string                 `json:"RequestId" xml:"RequestId"`
	MachineInfoStatistics []MachineInfoStatistic `json:"MachineInfoStatistics" xml:"MachineInfoStatistics"`
}

// CreateDescribeUuidsByVulNamesRequest creates a request to invoke DescribeUuidsByVulNames API
func CreateDescribeUuidsByVulNamesRequest() (request *DescribeUuidsByVulNamesRequest) {
	request = &DescribeUuidsByVulNamesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "DescribeUuidsByVulNames", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeUuidsByVulNamesResponse creates a response to parse from DescribeUuidsByVulNames response
func CreateDescribeUuidsByVulNamesResponse() (response *DescribeUuidsByVulNamesResponse) {
	response = &DescribeUuidsByVulNamesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
