package domain

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

// BidDomain invokes the domain.BidDomain API synchronously
func (client *Client) BidDomain(request *BidDomainRequest) (response *BidDomainResponse, err error) {
	response = CreateBidDomainResponse()
	err = client.DoAction(request, response)
	return
}

// BidDomainWithChan invokes the domain.BidDomain API asynchronously
func (client *Client) BidDomainWithChan(request *BidDomainRequest) (<-chan *BidDomainResponse, <-chan error) {
	responseChan := make(chan *BidDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BidDomain(request)
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

// BidDomainWithCallback invokes the domain.BidDomain API asynchronously
func (client *Client) BidDomainWithCallback(request *BidDomainRequest, callback func(response *BidDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BidDomainResponse
		var err error
		defer close(result)
		response, err = client.BidDomain(request)
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

// BidDomainRequest is the request struct for api BidDomain
type BidDomainRequest struct {
	*requests.RpcRequest
	AuctionId string         `position:"Body" name:"AuctionId"`
	MaxBid    requests.Float `position:"Body" name:"MaxBid"`
	Currency  string         `position:"Body" name:"Currency"`
}

// BidDomainResponse is the response struct for api BidDomain
type BidDomainResponse struct {
	*responses.BaseResponse
	AuctionId string `json:"AuctionId" xml:"AuctionId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateBidDomainRequest creates a request to invoke BidDomain API
func CreateBidDomainRequest() (request *BidDomainRequest) {
	request = &BidDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-02-08", "BidDomain", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBidDomainResponse creates a response to parse from BidDomain response
func CreateBidDomainResponse() (response *BidDomainResponse) {
	response = &BidDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
