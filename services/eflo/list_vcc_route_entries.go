package eflo

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

// ListVccRouteEntries invokes the eflo.ListVccRouteEntries API synchronously
func (client *Client) ListVccRouteEntries(request *ListVccRouteEntriesRequest) (response *ListVccRouteEntriesResponse, err error) {
	response = CreateListVccRouteEntriesResponse()
	err = client.DoAction(request, response)
	return
}

// ListVccRouteEntriesWithChan invokes the eflo.ListVccRouteEntries API asynchronously
func (client *Client) ListVccRouteEntriesWithChan(request *ListVccRouteEntriesRequest) (<-chan *ListVccRouteEntriesResponse, <-chan error) {
	responseChan := make(chan *ListVccRouteEntriesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVccRouteEntries(request)
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

// ListVccRouteEntriesWithCallback invokes the eflo.ListVccRouteEntries API asynchronously
func (client *Client) ListVccRouteEntriesWithCallback(request *ListVccRouteEntriesRequest, callback func(response *ListVccRouteEntriesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVccRouteEntriesResponse
		var err error
		defer close(result)
		response, err = client.ListVccRouteEntries(request)
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

// ListVccRouteEntriesRequest is the request struct for api ListVccRouteEntries
type ListVccRouteEntriesRequest struct {
	*requests.RpcRequest
	IgnoreDetailedRouteEntry requests.Boolean `position:"Body" name:"IgnoreDetailedRouteEntry"`
	PageNumber               requests.Integer `position:"Body" name:"PageNumber"`
	RouteType                string           `position:"Body" name:"RouteType"`
	ResourceGroupId          string           `position:"Body" name:"ResourceGroupId"`
	PageSize                 requests.Integer `position:"Body" name:"PageSize"`
	NextHopId                string           `position:"Body" name:"NextHopId"`
	VccId                    string           `position:"Body" name:"VccId"`
	NextHopType              string           `position:"Body" name:"NextHopType"`
	VpdRouteEntryId          string           `position:"Body" name:"VpdRouteEntryId"`
	DestinationCidrBlock     string           `position:"Body" name:"DestinationCidrBlock"`
	EnablePage               requests.Boolean `position:"Body" name:"EnablePage"`
	Status                   string           `position:"Body" name:"Status"`
}

// ListVccRouteEntriesResponse is the response struct for api ListVccRouteEntries
type ListVccRouteEntriesResponse struct {
	*responses.BaseResponse
	Code      int     `json:"Code" xml:"Code"`
	Message   string  `json:"Message" xml:"Message"`
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Content   Content `json:"Content" xml:"Content"`
}

// CreateListVccRouteEntriesRequest creates a request to invoke ListVccRouteEntries API
func CreateListVccRouteEntriesRequest() (request *ListVccRouteEntriesRequest) {
	request = &ListVccRouteEntriesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("eflo", "2022-05-30", "ListVccRouteEntries", "eflo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListVccRouteEntriesResponse creates a response to parse from ListVccRouteEntries response
func CreateListVccRouteEntriesResponse() (response *ListVccRouteEntriesResponse) {
	response = &ListVccRouteEntriesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
