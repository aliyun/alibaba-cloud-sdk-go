package antiddos_public

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

// DescribeBgpPackByIp invokes the antiddos_public.DescribeBgpPackByIp API synchronously
func (client *Client) DescribeBgpPackByIp(request *DescribeBgpPackByIpRequest) (response *DescribeBgpPackByIpResponse, err error) {
	response = CreateDescribeBgpPackByIpResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBgpPackByIpWithChan invokes the antiddos_public.DescribeBgpPackByIp API asynchronously
func (client *Client) DescribeBgpPackByIpWithChan(request *DescribeBgpPackByIpRequest) (<-chan *DescribeBgpPackByIpResponse, <-chan error) {
	responseChan := make(chan *DescribeBgpPackByIpResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBgpPackByIp(request)
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

// DescribeBgpPackByIpWithCallback invokes the antiddos_public.DescribeBgpPackByIp API asynchronously
func (client *Client) DescribeBgpPackByIpWithCallback(request *DescribeBgpPackByIpRequest, callback func(response *DescribeBgpPackByIpResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBgpPackByIpResponse
		var err error
		defer close(result)
		response, err = client.DescribeBgpPackByIp(request)
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

// DescribeBgpPackByIpRequest is the request struct for api DescribeBgpPackByIp
type DescribeBgpPackByIpRequest struct {
	*requests.RpcRequest
	SourceIp     string `position:"Query" name:"SourceIp"`
	DdosRegionId string `position:"Query" name:"DdosRegionId"`
	Lang         string `position:"Query" name:"Lang"`
	Ip           string `position:"Query" name:"Ip"`
}

// DescribeBgpPackByIpResponse is the response struct for api DescribeBgpPackByIp
type DescribeBgpPackByIpResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	Code        int         `json:"Code" xml:"Code"`
	Success     bool        `json:"Success" xml:"Success"`
	DdosbgpInfo DdosbgpInfo `json:"DdosbgpInfo" xml:"DdosbgpInfo"`
}

// CreateDescribeBgpPackByIpRequest creates a request to invoke DescribeBgpPackByIp API
func CreateDescribeBgpPackByIpRequest() (request *DescribeBgpPackByIpRequest) {
	request = &DescribeBgpPackByIpRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("antiddos-public", "2017-05-18", "DescribeBgpPackByIp", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeBgpPackByIpResponse creates a response to parse from DescribeBgpPackByIp response
func CreateDescribeBgpPackByIpResponse() (response *DescribeBgpPackByIpResponse) {
	response = &DescribeBgpPackByIpResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
