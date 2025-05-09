package drds

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

// FlashbackRecycleBinTable invokes the drds.FlashbackRecycleBinTable API synchronously
func (client *Client) FlashbackRecycleBinTable(request *FlashbackRecycleBinTableRequest) (response *FlashbackRecycleBinTableResponse, err error) {
	response = CreateFlashbackRecycleBinTableResponse()
	err = client.DoAction(request, response)
	return
}

// FlashbackRecycleBinTableWithChan invokes the drds.FlashbackRecycleBinTable API asynchronously
func (client *Client) FlashbackRecycleBinTableWithChan(request *FlashbackRecycleBinTableRequest) (<-chan *FlashbackRecycleBinTableResponse, <-chan error) {
	responseChan := make(chan *FlashbackRecycleBinTableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FlashbackRecycleBinTable(request)
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

// FlashbackRecycleBinTableWithCallback invokes the drds.FlashbackRecycleBinTable API asynchronously
func (client *Client) FlashbackRecycleBinTableWithCallback(request *FlashbackRecycleBinTableRequest, callback func(response *FlashbackRecycleBinTableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FlashbackRecycleBinTableResponse
		var err error
		defer close(result)
		response, err = client.FlashbackRecycleBinTable(request)
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

// FlashbackRecycleBinTableRequest is the request struct for api FlashbackRecycleBinTable
type FlashbackRecycleBinTableRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbName         string `position:"Query" name:"DbName"`
	TableName      string `position:"Query" name:"TableName"`
}

// FlashbackRecycleBinTableResponse is the response struct for api FlashbackRecycleBinTable
type FlashbackRecycleBinTableResponse struct {
	*responses.BaseResponse
	Data      bool   `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateFlashbackRecycleBinTableRequest creates a request to invoke FlashbackRecycleBinTable API
func CreateFlashbackRecycleBinTableRequest() (request *FlashbackRecycleBinTableRequest) {
	request = &FlashbackRecycleBinTableRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "FlashbackRecycleBinTable", "drds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateFlashbackRecycleBinTableResponse creates a response to parse from FlashbackRecycleBinTable response
func CreateFlashbackRecycleBinTableResponse() (response *FlashbackRecycleBinTableResponse) {
	response = &FlashbackRecycleBinTableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
