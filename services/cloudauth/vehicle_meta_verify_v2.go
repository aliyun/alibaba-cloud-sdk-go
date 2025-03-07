package cloudauth

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

// VehicleMetaVerifyV2 invokes the cloudauth.VehicleMetaVerifyV2 API synchronously
func (client *Client) VehicleMetaVerifyV2(request *VehicleMetaVerifyV2Request) (response *VehicleMetaVerifyV2Response, err error) {
	response = CreateVehicleMetaVerifyV2Response()
	err = client.DoAction(request, response)
	return
}

// VehicleMetaVerifyV2WithChan invokes the cloudauth.VehicleMetaVerifyV2 API asynchronously
func (client *Client) VehicleMetaVerifyV2WithChan(request *VehicleMetaVerifyV2Request) (<-chan *VehicleMetaVerifyV2Response, <-chan error) {
	responseChan := make(chan *VehicleMetaVerifyV2Response, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.VehicleMetaVerifyV2(request)
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

// VehicleMetaVerifyV2WithCallback invokes the cloudauth.VehicleMetaVerifyV2 API asynchronously
func (client *Client) VehicleMetaVerifyV2WithCallback(request *VehicleMetaVerifyV2Request, callback func(response *VehicleMetaVerifyV2Response, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *VehicleMetaVerifyV2Response
		var err error
		defer close(result)
		response, err = client.VehicleMetaVerifyV2(request)
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

// VehicleMetaVerifyV2Request is the request struct for api VehicleMetaVerifyV2
type VehicleMetaVerifyV2Request struct {
	*requests.RpcRequest
	VehicleType    string `position:"Query" name:"VehicleType"`
	ParamType      string `position:"Query" name:"ParamType"`
	VehicleNum     string `position:"Query" name:"VehicleNum"`
	IdentifyNum    string `position:"Query" name:"IdentifyNum"`
	VerifyMetaType string `position:"Query" name:"VerifyMetaType"`
	UserName       string `position:"Query" name:"UserName"`
}

// VehicleMetaVerifyV2Response is the response struct for api VehicleMetaVerifyV2
type VehicleMetaVerifyV2Response struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Code         string       `json:"Code" xml:"Code"`
	Message      string       `json:"Message" xml:"Message"`
	ResultObject ResultObject `json:"ResultObject" xml:"ResultObject"`
}

// CreateVehicleMetaVerifyV2Request creates a request to invoke VehicleMetaVerifyV2 API
func CreateVehicleMetaVerifyV2Request() (request *VehicleMetaVerifyV2Request) {
	request = &VehicleMetaVerifyV2Request{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "VehicleMetaVerifyV2", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateVehicleMetaVerifyV2Response creates a response to parse from VehicleMetaVerifyV2 response
func CreateVehicleMetaVerifyV2Response() (response *VehicleMetaVerifyV2Response) {
	response = &VehicleMetaVerifyV2Response{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
