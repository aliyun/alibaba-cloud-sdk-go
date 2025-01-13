package r_kvstore

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

// CreateTCInstance invokes the r_kvstore.CreateTCInstance API synchronously
func (client *Client) CreateTCInstance(request *CreateTCInstanceRequest) (response *CreateTCInstanceResponse, err error) {
	response = CreateCreateTCInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTCInstanceWithChan invokes the r_kvstore.CreateTCInstance API asynchronously
func (client *Client) CreateTCInstanceWithChan(request *CreateTCInstanceRequest) (<-chan *CreateTCInstanceResponse, <-chan error) {
	responseChan := make(chan *CreateTCInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTCInstance(request)
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

// CreateTCInstanceWithCallback invokes the r_kvstore.CreateTCInstance API asynchronously
func (client *Client) CreateTCInstanceWithCallback(request *CreateTCInstanceRequest, callback func(response *CreateTCInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTCInstanceResponse
		var err error
		defer close(result)
		response, err = client.CreateTCInstance(request)
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

// CreateTCInstanceRequest is the request struct for api CreateTCInstance
type CreateTCInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer            `position:"Query" name:"ResourceOwnerId"`
	CouponNo             string                      `position:"Query" name:"CouponNo"`
	NetworkType          string                      `position:"Query" name:"NetworkType"`
	NeedEni              requests.Boolean            `position:"Query" name:"NeedEni"`
	ResourceGroupId      string                      `position:"Query" name:"ResourceGroupId"`
	Password             string                      `position:"Query" name:"Password"`
	SecurityToken        string                      `position:"Query" name:"SecurityToken"`
	Tag                  *[]CreateTCInstanceTag      `position:"Query" name:"Tag"  type:"Repeated"`
	BusinessInfo         string                      `position:"Query" name:"BusinessInfo"`
	AutoRenewPeriod      string                      `position:"Query" name:"AutoRenewPeriod"`
	Period               string                      `position:"Query" name:"Period"`
	DryRun               requests.Boolean            `position:"Query" name:"DryRun"`
	OwnerId              requests.Integer            `position:"Query" name:"OwnerId"`
	VSwitchId            string                      `position:"Query" name:"VSwitchId"`
	InstanceName         string                      `position:"Query" name:"InstanceName"`
	AutoRenew            string                      `position:"Query" name:"AutoRenew"`
	ZoneId               string                      `position:"Query" name:"ZoneId"`
	ImageId              string                      `position:"Query" name:"ImageId"`
	ClientToken          string                      `position:"Query" name:"ClientToken"`
	SecurityGroupId      string                      `position:"Query" name:"SecurityGroupId"`
	AutoUseCoupon        string                      `position:"Query" name:"AutoUseCoupon"`
	InstanceClass        string                      `position:"Query" name:"InstanceClass"`
	InstanceChargeType   string                      `position:"Query" name:"InstanceChargeType"`
	DeploymentSetId      string                      `position:"Query" name:"DeploymentSetId"`
	ResourceOwnerAccount string                      `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                      `position:"Query" name:"OwnerAccount"`
	DataDisk             *[]CreateTCInstanceDataDisk `position:"Query" name:"DataDisk"  type:"Repeated"`
	VpcId                string                      `position:"Query" name:"VpcId"`
}

// CreateTCInstanceTag is a repeated param struct in CreateTCInstanceRequest
type CreateTCInstanceTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateTCInstanceDataDisk is a repeated param struct in CreateTCInstanceRequest
type CreateTCInstanceDataDisk struct {
	Size             string `name:"Size"`
	PerformanceLevel string `name:"PerformanceLevel"`
	Category         string `name:"Category"`
}

// CreateTCInstanceResponse is the response struct for api CreateTCInstance
type CreateTCInstanceResponse struct {
	*responses.BaseResponse
	InstanceId string `json:"InstanceId" xml:"InstanceId"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	OrderId    int64  `json:"OrderId" xml:"OrderId"`
}

// CreateCreateTCInstanceRequest creates a request to invoke CreateTCInstance API
func CreateCreateTCInstanceRequest() (request *CreateTCInstanceRequest) {
	request = &CreateTCInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "CreateTCInstance", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateTCInstanceResponse creates a response to parse from CreateTCInstance response
func CreateCreateTCInstanceResponse() (response *CreateTCInstanceResponse) {
	response = &CreateTCInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
