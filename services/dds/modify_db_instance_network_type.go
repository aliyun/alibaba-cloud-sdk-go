package dds

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

// ModifyDBInstanceNetworkType invokes the dds.ModifyDBInstanceNetworkType API synchronously
func (client *Client) ModifyDBInstanceNetworkType(request *ModifyDBInstanceNetworkTypeRequest) (response *ModifyDBInstanceNetworkTypeResponse, err error) {
	response = CreateModifyDBInstanceNetworkTypeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBInstanceNetworkTypeWithChan invokes the dds.ModifyDBInstanceNetworkType API asynchronously
func (client *Client) ModifyDBInstanceNetworkTypeWithChan(request *ModifyDBInstanceNetworkTypeRequest) (<-chan *ModifyDBInstanceNetworkTypeResponse, <-chan error) {
	responseChan := make(chan *ModifyDBInstanceNetworkTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBInstanceNetworkType(request)
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

// ModifyDBInstanceNetworkTypeWithCallback invokes the dds.ModifyDBInstanceNetworkType API asynchronously
func (client *Client) ModifyDBInstanceNetworkTypeWithCallback(request *ModifyDBInstanceNetworkTypeRequest, callback func(response *ModifyDBInstanceNetworkTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBInstanceNetworkTypeResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBInstanceNetworkType(request)
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

// ModifyDBInstanceNetworkTypeRequest is the request struct for api ModifyDBInstanceNetworkType
type ModifyDBInstanceNetworkTypeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	NetworkType          string           `position:"Query" name:"NetworkType"`
	ClassicExpiredDays   requests.Integer `position:"Query" name:"ClassicExpiredDays"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	VSwitchId            string           `position:"Query" name:"VSwitchId"`
	RetainClassic        string           `position:"Query" name:"RetainClassic"`
	VpcId                string           `position:"Query" name:"VpcId"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
}

// ModifyDBInstanceNetworkTypeResponse is the response struct for api ModifyDBInstanceNetworkType
type ModifyDBInstanceNetworkTypeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBInstanceNetworkTypeRequest creates a request to invoke ModifyDBInstanceNetworkType API
func CreateModifyDBInstanceNetworkTypeRequest() (request *ModifyDBInstanceNetworkTypeRequest) {
	request = &ModifyDBInstanceNetworkTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "ModifyDBInstanceNetworkType", "dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDBInstanceNetworkTypeResponse creates a response to parse from ModifyDBInstanceNetworkType response
func CreateModifyDBInstanceNetworkTypeResponse() (response *ModifyDBInstanceNetworkTypeResponse) {
	response = &ModifyDBInstanceNetworkTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
