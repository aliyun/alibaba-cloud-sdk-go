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

// CheckClickhouseToRDS invokes the clickhouse.CheckClickhouseToRDS API synchronously
func (client *Client) CheckClickhouseToRDS(request *CheckClickhouseToRDSRequest) (response *CheckClickhouseToRDSResponse, err error) {
	response = CreateCheckClickhouseToRDSResponse()
	err = client.DoAction(request, response)
	return
}

// CheckClickhouseToRDSWithChan invokes the clickhouse.CheckClickhouseToRDS API asynchronously
func (client *Client) CheckClickhouseToRDSWithChan(request *CheckClickhouseToRDSRequest) (<-chan *CheckClickhouseToRDSResponse, <-chan error) {
	responseChan := make(chan *CheckClickhouseToRDSResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckClickhouseToRDS(request)
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

// CheckClickhouseToRDSWithCallback invokes the clickhouse.CheckClickhouseToRDS API asynchronously
func (client *Client) CheckClickhouseToRDSWithCallback(request *CheckClickhouseToRDSRequest, callback func(response *CheckClickhouseToRDSResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckClickhouseToRDSResponse
		var err error
		defer close(result)
		response, err = client.CheckClickhouseToRDS(request)
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

// CheckClickhouseToRDSRequest is the request struct for api CheckClickhouseToRDS
type CheckClickhouseToRDSRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	RdsVpcId             string           `position:"Query" name:"RdsVpcId"`
	CkPassword           string           `position:"Query" name:"CkPassword"`
	RdsPassword          string           `position:"Query" name:"RdsPassword"`
	CkUserName           string           `position:"Query" name:"CkUserName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DbClusterId          string           `position:"Query" name:"DbClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	RdsId                string           `position:"Query" name:"RdsId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ClickhousePort       requests.Integer `position:"Query" name:"ClickhousePort"`
	RdsPort              requests.Integer `position:"Query" name:"RdsPort"`
	RdsVpcUrl            string           `position:"Query" name:"RdsVpcUrl"`
	RdsUserName          string           `position:"Query" name:"RdsUserName"`
}

// CheckClickhouseToRDSResponse is the response struct for api CheckClickhouseToRDS
type CheckClickhouseToRDSResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Status    bool   `json:"Status" xml:"Status"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
}

// CreateCheckClickhouseToRDSRequest creates a request to invoke CheckClickhouseToRDS API
func CreateCheckClickhouseToRDSRequest() (request *CheckClickhouseToRDSRequest) {
	request = &CheckClickhouseToRDSRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("clickhouse", "2019-11-11", "CheckClickhouseToRDS", "service", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCheckClickhouseToRDSResponse creates a response to parse from CheckClickhouseToRDS response
func CreateCheckClickhouseToRDSResponse() (response *CheckClickhouseToRDSResponse) {
	response = &CheckClickhouseToRDSResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
