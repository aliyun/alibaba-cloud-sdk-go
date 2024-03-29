package qualitycheck

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

// UpdateSyncQualityCheckData invokes the qualitycheck.UpdateSyncQualityCheckData API synchronously
func (client *Client) UpdateSyncQualityCheckData(request *UpdateSyncQualityCheckDataRequest) (response *UpdateSyncQualityCheckDataResponse, err error) {
	response = CreateUpdateSyncQualityCheckDataResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateSyncQualityCheckDataWithChan invokes the qualitycheck.UpdateSyncQualityCheckData API asynchronously
func (client *Client) UpdateSyncQualityCheckDataWithChan(request *UpdateSyncQualityCheckDataRequest) (<-chan *UpdateSyncQualityCheckDataResponse, <-chan error) {
	responseChan := make(chan *UpdateSyncQualityCheckDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateSyncQualityCheckData(request)
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

// UpdateSyncQualityCheckDataWithCallback invokes the qualitycheck.UpdateSyncQualityCheckData API asynchronously
func (client *Client) UpdateSyncQualityCheckDataWithCallback(request *UpdateSyncQualityCheckDataRequest, callback func(response *UpdateSyncQualityCheckDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateSyncQualityCheckDataResponse
		var err error
		defer close(result)
		response, err = client.UpdateSyncQualityCheckData(request)
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

// UpdateSyncQualityCheckDataRequest is the request struct for api UpdateSyncQualityCheckData
type UpdateSyncQualityCheckDataRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
	BaseMeAgentId   requests.Integer `position:"Query" name:"BaseMeAgentId"`
}

// UpdateSyncQualityCheckDataResponse is the response struct for api UpdateSyncQualityCheckData
type UpdateSyncQualityCheckDataResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateUpdateSyncQualityCheckDataRequest creates a request to invoke UpdateSyncQualityCheckData API
func CreateUpdateSyncQualityCheckDataRequest() (request *UpdateSyncQualityCheckDataRequest) {
	request = &UpdateSyncQualityCheckDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "UpdateSyncQualityCheckData", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateSyncQualityCheckDataResponse creates a response to parse from UpdateSyncQualityCheckData response
func CreateUpdateSyncQualityCheckDataResponse() (response *UpdateSyncQualityCheckDataResponse) {
	response = &UpdateSyncQualityCheckDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
