package clickhouse

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

// CreateDBInstance invokes the clickhouse.CreateDBInstance API synchronously
func (client *Client) CreateDBInstance(request *CreateDBInstanceRequest) (response *CreateDBInstanceResponse, err error) {
	response = CreateCreateDBInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDBInstanceWithChan invokes the clickhouse.CreateDBInstance API asynchronously
func (client *Client) CreateDBInstanceWithChan(request *CreateDBInstanceRequest) (<-chan *CreateDBInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateDBInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDBInstance(request)
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

// CreateDBInstanceWithCallback invokes the clickhouse.CreateDBInstance API asynchronously
func (client *Client) CreateDBInstanceWithCallback(request *CreateDBInstanceRequest, callback func(response *CreateDBInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDBInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateDBInstance(request)
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

// CreateDBInstanceRequest is the request struct for api CreateDBInstance
type CreateDBInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBClusterDescription string           `position:"Query" name:"DBClusterDescription"`
	SourceDBClusterId    string           `position:"Query" name:"SourceDBClusterId"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	ZondIdBak2           string           `position:"Query" name:"ZondIdBak2"`
	DbNodeStorageType    string           `position:"Query" name:"DbNodeStorageType"`
	EncryptionType       string           `position:"Query" name:"EncryptionType"`
	ZoneIdBak            string           `position:"Query" name:"ZoneIdBak"`
	Period               string           `position:"Query" name:"Period"`
	BackupSetID          string           `position:"Query" name:"BackupSetID"`
	EncryptionKey        string           `position:"Query" name:"EncryptionKey"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBNodeGroupCount     string           `position:"Query" name:"DBNodeGroupCount"`
	VSwitchId            string           `position:"Query" name:"VSwitchId"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	DBClusterCategory    string           `position:"Query" name:"DBClusterCategory"`
	DBClusterNetworkType string           `position:"Query" name:"DBClusterNetworkType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	DBClusterVersion     string           `position:"Query" name:"DBClusterVersion"`
	DBClusterClass       string           `position:"Query" name:"DBClusterClass"`
	VSwitchBak           string           `position:"Query" name:"VSwitchBak"`
	UsedTime             string           `position:"Query" name:"UsedTime"`
	VSwitchBak2          string           `position:"Query" name:"VSwitchBak2"`
	DBNodeStorage        string           `position:"Query" name:"DBNodeStorage"`
	VPCId                string           `position:"Query" name:"VPCId"`
	PayType              string           `position:"Query" name:"PayType"`
}

// CreateDBInstanceResponse is the response struct for api CreateDBInstance
type CreateDBInstanceResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	DBClusterId string `json:"DBClusterId" xml:"DBClusterId"`
	OrderId     string `json:"OrderId" xml:"OrderId"`
}

// CreateCreateDBInstanceRequest creates a request to invoke CreateDBInstance API
func CreateCreateDBInstanceRequest() (request *CreateDBInstanceRequest) {
	request = &CreateDBInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("clickhouse", "2019-11-11", "CreateDBInstance", "service", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDBInstanceResponse creates a response to parse from CreateDBInstance response
func CreateCreateDBInstanceResponse() (response *CreateDBInstanceResponse) {
	response = &CreateDBInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
