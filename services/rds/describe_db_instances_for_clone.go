package rds

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

// DescribeDBInstancesForClone invokes the rds.DescribeDBInstancesForClone API synchronously
func (client *Client) DescribeDBInstancesForClone(request *DescribeDBInstancesForCloneRequest) (response *DescribeDBInstancesForCloneResponse, err error) {
	response = CreateDescribeDBInstancesForCloneResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBInstancesForCloneWithChan invokes the rds.DescribeDBInstancesForClone API asynchronously
func (client *Client) DescribeDBInstancesForCloneWithChan(request *DescribeDBInstancesForCloneRequest) (<-chan *DescribeDBInstancesForCloneResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstancesForCloneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstancesForClone(request)
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

// DescribeDBInstancesForCloneWithCallback invokes the rds.DescribeDBInstancesForClone API asynchronously
func (client *Client) DescribeDBInstancesForCloneWithCallback(request *DescribeDBInstancesForCloneRequest, callback func(response *DescribeDBInstancesForCloneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstancesForCloneResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstancesForClone(request)
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

// DescribeDBInstancesForCloneRequest is the request struct for api DescribeDBInstancesForClone
type DescribeDBInstancesForCloneRequest struct {
	*requests.RpcRequest
	ConnectionMode       string           `position:"Query" name:"ConnectionMode"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	NodeType             string           `position:"Query" name:"NodeType"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	SearchKey            string           `position:"Query" name:"SearchKey"`
	EngineVersion        string           `position:"Query" name:"EngineVersion"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	Expired              string           `position:"Query" name:"Expired"`
	Engine               string           `position:"Query" name:"Engine"`
	CurrentInstanceId    string           `position:"Query" name:"CurrentInstanceId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	DBInstanceStatus     string           `position:"Query" name:"DBInstanceStatus"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ProxyId              string           `position:"Query" name:"proxyId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBInstanceType       string           `position:"Query" name:"DBInstanceType"`
	DBInstanceClass      string           `position:"Query" name:"DBInstanceClass"`
	VSwitchId            string           `position:"Query" name:"VSwitchId"`
	VpcId                string           `position:"Query" name:"VpcId"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
	PayType              string           `position:"Query" name:"PayType"`
	InstanceNetworkType  string           `position:"Query" name:"InstanceNetworkType"`
}

// DescribeDBInstancesForCloneResponse is the response struct for api DescribeDBInstancesForClone
type DescribeDBInstancesForCloneResponse struct {
	*responses.BaseResponse
	RequestId        string                             `json:"RequestId" xml:"RequestId"`
	PageNumber       int                                `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  int                                `json:"PageRecordCount" xml:"PageRecordCount"`
	TotalRecordCount int                                `json:"TotalRecordCount" xml:"TotalRecordCount"`
	Items            ItemsInDescribeDBInstancesForClone `json:"Items" xml:"Items"`
}

// CreateDescribeDBInstancesForCloneRequest creates a request to invoke DescribeDBInstancesForClone API
func CreateDescribeDBInstancesForCloneRequest() (request *DescribeDBInstancesForCloneRequest) {
	request = &DescribeDBInstancesForCloneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstancesForClone", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBInstancesForCloneResponse creates a response to parse from DescribeDBInstancesForClone response
func CreateDescribeDBInstancesForCloneResponse() (response *DescribeDBInstancesForCloneResponse) {
	response = &DescribeDBInstancesForCloneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
