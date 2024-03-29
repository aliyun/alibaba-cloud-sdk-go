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

// ChangeAuction invokes the domain.ChangeAuction API synchronously
func (client *Client) ChangeAuction(request *ChangeAuctionRequest) (response *ChangeAuctionResponse, err error) {
	response = CreateChangeAuctionResponse()
	err = client.DoAction(request, response)
	return
}

// ChangeAuctionWithChan invokes the domain.ChangeAuction API asynchronously
func (client *Client) ChangeAuctionWithChan(request *ChangeAuctionRequest) (<-chan *ChangeAuctionResponse, <-chan error) {
	responseChan := make(chan *ChangeAuctionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ChangeAuction(request)
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

// ChangeAuctionWithCallback invokes the domain.ChangeAuction API asynchronously
func (client *Client) ChangeAuctionWithCallback(request *ChangeAuctionRequest, callback func(response *ChangeAuctionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ChangeAuctionResponse
		var err error
		defer close(result)
		response, err = client.ChangeAuction(request)
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

// ChangeAuctionRequest is the request struct for api ChangeAuction
type ChangeAuctionRequest struct {
	*requests.RpcRequest
	AuctionList *[]ChangeAuctionAuctionList `position:"Body" name:"AuctionList"  type:"Repeated"`
}

// ChangeAuctionAuctionList is a repeated param struct in ChangeAuctionRequest
type ChangeAuctionAuctionList struct {
	Winner       string                                `name:"Winner"`
	ReserveRange string                                `name:"ReserveRange"`
	DomainName   string                                `name:"DomainName"`
	EndTime      string                                `name:"EndTime"`
	TimeLeft     string                                `name:"TimeLeft"`
	BidRecords   *[]ChangeAuctionAuctionListBidRecords `name:"BidRecords" type:"Repeated"`
	IsReserve    string                                `name:"IsReserve"`
	Status       string                                `name:"Status"`
	WinnerPrice  string                                `name:"WinnerPrice"`
	ReservePrice string                                `name:"ReservePrice"`
}

// ChangeAuctionAuctionListBidRecords is a repeated param struct in ChangeAuctionRequest
type ChangeAuctionAuctionListBidRecords struct {
	CreateTime string `name:"CreateTime"`
	Price      string `name:"Price"`
	UserId     string `name:"UserId"`
}

// ChangeAuctionResponse is the response struct for api ChangeAuction
type ChangeAuctionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateChangeAuctionRequest creates a request to invoke ChangeAuction API
func CreateChangeAuctionRequest() (request *ChangeAuctionRequest) {
	request = &ChangeAuctionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-02-08", "ChangeAuction", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateChangeAuctionResponse creates a response to parse from ChangeAuction response
func CreateChangeAuctionResponse() (response *ChangeAuctionResponse) {
	response = &ChangeAuctionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
