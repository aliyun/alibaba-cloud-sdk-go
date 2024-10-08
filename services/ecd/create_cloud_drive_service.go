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

// CreateCloudDriveService invokes the ecd.CreateCloudDriveService API synchronously
func (client *Client) CreateCloudDriveService(request *CreateCloudDriveServiceRequest) (response *CreateCloudDriveServiceResponse, err error) {
	response = CreateCreateCloudDriveServiceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCloudDriveServiceWithChan invokes the ecd.CreateCloudDriveService API asynchronously
func (client *Client) CreateCloudDriveServiceWithChan(request *CreateCloudDriveServiceRequest) (<-chan *CreateCloudDriveServiceResponse, <-chan error) {
	responseChan := make(chan *CreateCloudDriveServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCloudDriveService(request)
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

// CreateCloudDriveServiceWithCallback invokes the ecd.CreateCloudDriveService API asynchronously
func (client *Client) CreateCloudDriveServiceWithCallback(request *CreateCloudDriveServiceRequest, callback func(response *CreateCloudDriveServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCloudDriveServiceResponse
		var err error
		defer close(result)
		response, err = client.CreateCloudDriveService(request)
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

// CreateCloudDriveServiceRequest is the request struct for api CreateCloudDriveService
type CreateCloudDriveServiceRequest struct {
	*requests.RpcRequest
	OfficeSiteId   string           `position:"Query" name:"OfficeSiteId"`
	CdsChargeType  string           `position:"Query" name:"CdsChargeType"`
	CenId          string           `position:"Query" name:"CenId"`
	GlobalStatus   string           `position:"Query" name:"GlobalStatus"`
	OfficeSiteType string           `position:"Query" name:"OfficeSiteType"`
	UserCount      requests.Integer `position:"Query" name:"UserCount"`
	EndUserId      *[]string        `position:"Query" name:"EndUserId"  type:"Repeated"`
	SolutionId     string           `position:"Query" name:"SolutionId"`
	UserMaxSize    requests.Integer `position:"Query" name:"UserMaxSize"`
	Period         requests.Integer `position:"Query" name:"Period"`
	AutoPay        requests.Boolean `position:"Query" name:"AutoPay"`
	DomainName     string           `position:"Query" name:"DomainName"`
	BizType        requests.Integer `position:"Query" name:"BizType"`
	PeriodUnit     string           `position:"Query" name:"PeriodUnit"`
	AutoRenew      requests.Boolean `position:"Query" name:"AutoRenew"`
	Name           string           `position:"Query" name:"Name"`
	MaxSize        requests.Integer `position:"Query" name:"MaxSize"`
}

// CreateCloudDriveServiceResponse is the response struct for api CreateCloudDriveService
type CreateCloudDriveServiceResponse struct {
	*responses.BaseResponse
	RequestId           string              `json:"RequestId" xml:"RequestId"`
	CdsId               string              `json:"CdsId" xml:"CdsId"`
	CdsName             string              `json:"CdsName" xml:"CdsName"`
	MaxSize             string              `json:"MaxSize" xml:"MaxSize"`
	OfficeSiteType      string              `json:"OfficeSiteType" xml:"OfficeSiteType"`
	CenId               string              `json:"CenId" xml:"CenId"`
	DomainName          string              `json:"DomainName" xml:"DomainName"`
	OrderId             string              `json:"OrderId" xml:"OrderId"`
	ErrorCode           string              `json:"ErrorCode" xml:"ErrorCode"`
	ConflictCdsAndOrder ConflictCdsAndOrder `json:"ConflictCdsAndOrder" xml:"ConflictCdsAndOrder"`
}

// CreateCreateCloudDriveServiceRequest creates a request to invoke CreateCloudDriveService API
func CreateCreateCloudDriveServiceRequest() (request *CreateCloudDriveServiceRequest) {
	request = &CreateCloudDriveServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "CreateCloudDriveService", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateCloudDriveServiceResponse creates a response to parse from CreateCloudDriveService response
func CreateCreateCloudDriveServiceResponse() (response *CreateCloudDriveServiceResponse) {
	response = &CreateCloudDriveServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
