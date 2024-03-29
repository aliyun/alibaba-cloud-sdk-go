package config

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

// GetAggregateConfigDeliveryChannel invokes the config.GetAggregateConfigDeliveryChannel API synchronously
func (client *Client) GetAggregateConfigDeliveryChannel(request *GetAggregateConfigDeliveryChannelRequest) (response *GetAggregateConfigDeliveryChannelResponse, err error) {
	response = CreateGetAggregateConfigDeliveryChannelResponse()
	err = client.DoAction(request, response)
	return
}

// GetAggregateConfigDeliveryChannelWithChan invokes the config.GetAggregateConfigDeliveryChannel API asynchronously
func (client *Client) GetAggregateConfigDeliveryChannelWithChan(request *GetAggregateConfigDeliveryChannelRequest) (<-chan *GetAggregateConfigDeliveryChannelResponse, <-chan error) {
	responseChan := make(chan *GetAggregateConfigDeliveryChannelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAggregateConfigDeliveryChannel(request)
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

// GetAggregateConfigDeliveryChannelWithCallback invokes the config.GetAggregateConfigDeliveryChannel API asynchronously
func (client *Client) GetAggregateConfigDeliveryChannelWithCallback(request *GetAggregateConfigDeliveryChannelRequest, callback func(response *GetAggregateConfigDeliveryChannelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAggregateConfigDeliveryChannelResponse
		var err error
		defer close(result)
		response, err = client.GetAggregateConfigDeliveryChannel(request)
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

// GetAggregateConfigDeliveryChannelRequest is the request struct for api GetAggregateConfigDeliveryChannel
type GetAggregateConfigDeliveryChannelRequest struct {
	*requests.RpcRequest
	AggregatorId      string `position:"Query" name:"AggregatorId"`
	DeliveryChannelId string `position:"Query" name:"DeliveryChannelId"`
}

// GetAggregateConfigDeliveryChannelResponse is the response struct for api GetAggregateConfigDeliveryChannel
type GetAggregateConfigDeliveryChannelResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	DeliveryChannel DeliveryChannel `json:"DeliveryChannel" xml:"DeliveryChannel"`
}

// CreateGetAggregateConfigDeliveryChannelRequest creates a request to invoke GetAggregateConfigDeliveryChannel API
func CreateGetAggregateConfigDeliveryChannelRequest() (request *GetAggregateConfigDeliveryChannelRequest) {
	request = &GetAggregateConfigDeliveryChannelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "GetAggregateConfigDeliveryChannel", "", "")
	request.Method = requests.POST
	return
}

// CreateGetAggregateConfigDeliveryChannelResponse creates a response to parse from GetAggregateConfigDeliveryChannel response
func CreateGetAggregateConfigDeliveryChannelResponse() (response *GetAggregateConfigDeliveryChannelResponse) {
	response = &GetAggregateConfigDeliveryChannelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
