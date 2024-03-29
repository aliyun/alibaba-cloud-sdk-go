package pairecservice

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

// ListABMetricGroups invokes the pairecservice.ListABMetricGroups API synchronously
func (client *Client) ListABMetricGroups(request *ListABMetricGroupsRequest) (response *ListABMetricGroupsResponse, err error) {
	response = CreateListABMetricGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// ListABMetricGroupsWithChan invokes the pairecservice.ListABMetricGroups API asynchronously
func (client *Client) ListABMetricGroupsWithChan(request *ListABMetricGroupsRequest) (<-chan *ListABMetricGroupsResponse, <-chan error) {
	responseChan := make(chan *ListABMetricGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListABMetricGroups(request)
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

// ListABMetricGroupsWithCallback invokes the pairecservice.ListABMetricGroups API asynchronously
func (client *Client) ListABMetricGroupsWithCallback(request *ListABMetricGroupsRequest, callback func(response *ListABMetricGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListABMetricGroupsResponse
		var err error
		defer close(result)
		response, err = client.ListABMetricGroups(request)
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

// ListABMetricGroupsRequest is the request struct for api ListABMetricGroups
type ListABMetricGroupsRequest struct {
	*requests.RoaRequest
	Realtime   requests.Boolean `position:"Query" name:"Realtime"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	SceneId    string           `position:"Query" name:"SceneId"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	SortBy     string           `position:"Query" name:"SortBy"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	Order      string           `position:"Query" name:"Order"`
}

// ListABMetricGroupsResponse is the response struct for api ListABMetricGroups
type ListABMetricGroupsResponse struct {
	*responses.BaseResponse
	RequestId      string               `json:"RequestId" xml:"RequestId"`
	TotalCount     int64                `json:"TotalCount" xml:"TotalCount"`
	ABMetricGroups []ABMetricGroupsItem `json:"ABMetricGroups" xml:"ABMetricGroups"`
}

// CreateListABMetricGroupsRequest creates a request to invoke ListABMetricGroups API
func CreateListABMetricGroupsRequest() (request *ListABMetricGroupsRequest) {
	request = &ListABMetricGroupsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "ListABMetricGroups", "/api/v1/abmetricgroups", "", "")
	request.Method = requests.GET
	return
}

// CreateListABMetricGroupsResponse creates a response to parse from ListABMetricGroups response
func CreateListABMetricGroupsResponse() (response *ListABMetricGroupsResponse) {
	response = &ListABMetricGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
