package ecd

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

// DescribeOfficeSites invokes the ecd.DescribeOfficeSites API synchronously
func (client *Client) DescribeOfficeSites(request *DescribeOfficeSitesRequest) (response *DescribeOfficeSitesResponse, err error) {
	response = CreateDescribeOfficeSitesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeOfficeSitesWithChan invokes the ecd.DescribeOfficeSites API asynchronously
func (client *Client) DescribeOfficeSitesWithChan(request *DescribeOfficeSitesRequest) (<-chan *DescribeOfficeSitesResponse, <-chan error) {
	responseChan := make(chan *DescribeOfficeSitesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeOfficeSites(request)
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

// DescribeOfficeSitesWithCallback invokes the ecd.DescribeOfficeSites API asynchronously
func (client *Client) DescribeOfficeSitesWithCallback(request *DescribeOfficeSitesRequest, callback func(response *DescribeOfficeSitesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeOfficeSitesResponse
		var err error
		defer close(result)
		response, err = client.DescribeOfficeSites(request)
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

// DescribeOfficeSitesRequest is the request struct for api DescribeOfficeSites
type DescribeOfficeSitesRequest struct {
	*requests.RpcRequest
	OfficeSiteId         *[]string        `position:"Query" name:"OfficeSiteId"  type:"Repeated"`
	VpcType              string           `position:"Query" name:"VpcType"`
	NmVersion            string           `position:"Query" name:"NmVersion"`
	EnableInternetAccess requests.Boolean `position:"Query" name:"EnableInternetAccess"`
	VpcNotNone           requests.Boolean `position:"Query" name:"VpcNotNone"`
	OfficeSiteType       string           `position:"Query" name:"OfficeSiteType"`
	NextToken            string           `position:"Query" name:"NextToken"`
	SecurityProtection   string           `position:"Query" name:"SecurityProtection"`
	Name                 string           `position:"Query" name:"Name"`
	MaxResults           requests.Integer `position:"Query" name:"MaxResults"`
	Status               string           `position:"Query" name:"Status"`
}

// DescribeOfficeSitesResponse is the response struct for api DescribeOfficeSites
type DescribeOfficeSitesResponse struct {
	*responses.BaseResponse
	NextToken   string       `json:"NextToken" xml:"NextToken"`
	RequestId   string       `json:"RequestId" xml:"RequestId"`
	TotalCount  int          `json:"TotalCount" xml:"TotalCount"`
	OfficeSites []OfficeSite `json:"OfficeSites" xml:"OfficeSites"`
}

// CreateDescribeOfficeSitesRequest creates a request to invoke DescribeOfficeSites API
func CreateDescribeOfficeSitesRequest() (request *DescribeOfficeSitesRequest) {
	request = &DescribeOfficeSitesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DescribeOfficeSites", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeOfficeSitesResponse creates a response to parse from DescribeOfficeSites response
func CreateDescribeOfficeSitesResponse() (response *DescribeOfficeSitesResponse) {
	response = &DescribeOfficeSitesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
