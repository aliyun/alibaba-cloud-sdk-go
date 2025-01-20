package polardb

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

// ModifyDBClusterAndNodesParameters invokes the polardb.ModifyDBClusterAndNodesParameters API synchronously
func (client *Client) ModifyDBClusterAndNodesParameters(request *ModifyDBClusterAndNodesParametersRequest) (response *ModifyDBClusterAndNodesParametersResponse, err error) {
	response = CreateModifyDBClusterAndNodesParametersResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDBClusterAndNodesParametersWithChan invokes the polardb.ModifyDBClusterAndNodesParameters API asynchronously
func (client *Client) ModifyDBClusterAndNodesParametersWithChan(request *ModifyDBClusterAndNodesParametersRequest) (<-chan *ModifyDBClusterAndNodesParametersResponse, <-chan error) {
	responseChan := make(chan *ModifyDBClusterAndNodesParametersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDBClusterAndNodesParameters(request)
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

// ModifyDBClusterAndNodesParametersWithCallback invokes the polardb.ModifyDBClusterAndNodesParameters API asynchronously
func (client *Client) ModifyDBClusterAndNodesParametersWithCallback(request *ModifyDBClusterAndNodesParametersRequest, callback func(response *ModifyDBClusterAndNodesParametersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDBClusterAndNodesParametersResponse
		var err error
		defer close(result)
		response, err = client.ModifyDBClusterAndNodesParameters(request)
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

// ModifyDBClusterAndNodesParametersRequest is the request struct for api ModifyDBClusterAndNodesParameters
type ModifyDBClusterAndNodesParametersRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PlannedEndTime                 string           `position:"Query" name:"PlannedEndTime"`
	DBNodeIds                      string           `position:"Query" name:"DBNodeIds"`
	ClearBinlog                    requests.Boolean `position:"Query" name:"ClearBinlog"`
	ParameterGroupId               string           `position:"Query" name:"ParameterGroupId"`
	ResourceOwnerAccount           string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId                    string           `position:"Query" name:"DBClusterId"`
	OwnerAccount                   string           `position:"Query" name:"OwnerAccount"`
	OwnerId                        requests.Integer `position:"Query" name:"OwnerId"`
	PlannedStartTime               string           `position:"Query" name:"PlannedStartTime"`
	StandbyClusterIdListNeedToSync string           `position:"Query" name:"StandbyClusterIdListNeedToSync"`
	Parameters                     string           `position:"Query" name:"Parameters"`
	FromTimeService                requests.Boolean `position:"Query" name:"FromTimeService"`
}

// ModifyDBClusterAndNodesParametersResponse is the response struct for api ModifyDBClusterAndNodesParameters
type ModifyDBClusterAndNodesParametersResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDBClusterAndNodesParametersRequest creates a request to invoke ModifyDBClusterAndNodesParameters API
func CreateModifyDBClusterAndNodesParametersRequest() (request *ModifyDBClusterAndNodesParametersRequest) {
	request = &ModifyDBClusterAndNodesParametersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "ModifyDBClusterAndNodesParameters", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDBClusterAndNodesParametersResponse creates a response to parse from ModifyDBClusterAndNodesParameters response
func CreateModifyDBClusterAndNodesParametersResponse() (response *ModifyDBClusterAndNodesParametersResponse) {
	response = &ModifyDBClusterAndNodesParametersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
