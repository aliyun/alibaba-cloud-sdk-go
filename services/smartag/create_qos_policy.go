package smartag

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

// CreateQosPolicy invokes the smartag.CreateQosPolicy API synchronously
func (client *Client) CreateQosPolicy(request *CreateQosPolicyRequest) (response *CreateQosPolicyResponse, err error) {
	response = CreateCreateQosPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// CreateQosPolicyWithChan invokes the smartag.CreateQosPolicy API asynchronously
func (client *Client) CreateQosPolicyWithChan(request *CreateQosPolicyRequest) (<-chan *CreateQosPolicyResponse, <-chan error) {
	responseChan := make(chan *CreateQosPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateQosPolicy(request)
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

// CreateQosPolicyWithCallback invokes the smartag.CreateQosPolicy API asynchronously
func (client *Client) CreateQosPolicyWithCallback(request *CreateQosPolicyRequest, callback func(response *CreateQosPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateQosPolicyResponse
		var err error
		defer close(result)
		response, err = client.CreateQosPolicy(request)
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

// CreateQosPolicyRequest is the request struct for api CreateQosPolicy
type CreateQosPolicyRequest struct {
	*requests.RpcRequest
	DpiGroupIds          *[]string        `position:"Query" name:"DpiGroupIds"  type:"Repeated"`
	ResourceOwnerId      requests.Integer `position:"Query"`
	SourcePortRange      string           `position:"Query"`
	SourceCidr           string           `position:"Query"`
	Description          string           `position:"Query"`
	StartTime            string           `position:"Query"`
	DestCidr             string           `position:"Query"`
	DpiSignatureIds      *[]string        `position:"Query" name:"DpiSignatureIds"  type:"Repeated"`
	QosId                string           `position:"Query"`
	ResourceOwnerAccount string           `position:"Query"`
	IpProtocol           string           `position:"Query"`
	OwnerAccount         string           `position:"Query"`
	EndTime              string           `position:"Query"`
	OwnerId              requests.Integer `position:"Query"`
	Priority             requests.Integer `position:"Query"`
	DestPortRange        string           `position:"Query"`
	Name                 string           `position:"Query"`
}

// CreateQosPolicyResponse is the response struct for api CreateQosPolicy
type CreateQosPolicyResponse struct {
	*responses.BaseResponse
	DestCidr        string                           `json:"DestCidr" xml:"DestCidr"`
	Description     string                           `json:"Description" xml:"Description"`
	SourcePortRange string                           `json:"SourcePortRange" xml:"SourcePortRange"`
	QosPolicyId     string                           `json:"QosPolicyId" xml:"QosPolicyId"`
	RequestId       string                           `json:"RequestId" xml:"RequestId"`
	EndTime         string                           `json:"EndTime" xml:"EndTime"`
	SourceCidr      string                           `json:"SourceCidr" xml:"SourceCidr"`
	Priority        int                              `json:"Priority" xml:"Priority"`
	StartTime       string                           `json:"StartTime" xml:"StartTime"`
	IpProtocol      string                           `json:"IpProtocol" xml:"IpProtocol"`
	QosId           string                           `json:"QosId" xml:"QosId"`
	DestPortRange   string                           `json:"DestPortRange" xml:"DestPortRange"`
	Name            string                           `json:"Name" xml:"Name"`
	DpiSignatureIds DpiSignatureIdsInCreateQosPolicy `json:"DpiSignatureIds" xml:"DpiSignatureIds"`
	DpiGroupIds     DpiGroupIdsInCreateQosPolicy     `json:"DpiGroupIds" xml:"DpiGroupIds"`
}

// CreateCreateQosPolicyRequest creates a request to invoke CreateQosPolicy API
func CreateCreateQosPolicyRequest() (request *CreateQosPolicyRequest) {
	request = &CreateQosPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "CreateQosPolicy", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateQosPolicyResponse creates a response to parse from CreateQosPolicy response
func CreateCreateQosPolicyResponse() (response *CreateQosPolicyResponse) {
	response = &CreateQosPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
