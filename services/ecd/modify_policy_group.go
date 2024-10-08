package ecd

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

// ModifyPolicyGroup invokes the ecd.ModifyPolicyGroup API synchronously
func (client *Client) ModifyPolicyGroup(request *ModifyPolicyGroupRequest) (response *ModifyPolicyGroupResponse, err error) {
	response = CreateModifyPolicyGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyPolicyGroupWithChan invokes the ecd.ModifyPolicyGroup API asynchronously
func (client *Client) ModifyPolicyGroupWithChan(request *ModifyPolicyGroupRequest) (<-chan *ModifyPolicyGroupResponse, <-chan error) {
	responseChan := make(chan *ModifyPolicyGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyPolicyGroup(request)
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

// ModifyPolicyGroupWithCallback invokes the ecd.ModifyPolicyGroup API asynchronously
func (client *Client) ModifyPolicyGroupWithCallback(request *ModifyPolicyGroupRequest, callback func(response *ModifyPolicyGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyPolicyGroupResponse
		var err error
		defer close(result)
		response, err = client.ModifyPolicyGroup(request)
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

// ModifyPolicyGroupRequest is the request struct for api ModifyPolicyGroup
type ModifyPolicyGroupRequest struct {
	*requests.RpcRequest
	RevokeSecurityPolicyRule      *[]ModifyPolicyGroupRevokeSecurityPolicyRule    `position:"Query" name:"RevokeSecurityPolicyRule"  type:"Repeated"`
	PrinterRedirection            string                                          `position:"Query" name:"PrinterRedirection"`
	WyAssistant                   string                                          `position:"Query" name:"WyAssistant"`
	LocalDrive                    string                                          `position:"Query" name:"LocalDrive"`
	AuthorizeSecurityPolicyRule   *[]ModifyPolicyGroupAuthorizeSecurityPolicyRule `position:"Query" name:"AuthorizeSecurityPolicyRule"  type:"Repeated"`
	WatermarkSecurity             string                                          `position:"Query" name:"WatermarkSecurity"`
	Clipboard                     string                                          `position:"Query" name:"Clipboard"`
	EndUserApplyAdminCoordinate   string                                          `position:"Query" name:"EndUserApplyAdminCoordinate"`
	UsbRedirect                   string                                          `position:"Query" name:"UsbRedirect"`
	RecordingStartTime            string                                          `position:"Query" name:"RecordingStartTime"`
	NetRedirectRule               *[]ModifyPolicyGroupNetRedirectRule             `position:"Query" name:"NetRedirectRule"  type:"Repeated"`
	WatermarkColor                requests.Integer                                `position:"Query" name:"WatermarkColor"`
	MemoryRateLimit               requests.Integer                                `position:"Query" name:"MemoryRateLimit"`
	QualityEnhancement            string                                          `position:"Query" name:"QualityEnhancement"`
	Watermark                     string                                          `position:"Query" name:"Watermark"`
	CpuSingleRateLimit            requests.Integer                                `position:"Query" name:"CpuSingleRateLimit"`
	DomainResolveRule             *[]ModifyPolicyGroupDomainResolveRule           `position:"Query" name:"DomainResolveRule"  type:"Repeated"`
	WatermarkPower                string                                          `position:"Query" name:"WatermarkPower"`
	PolicyGroupId                 string                                          `position:"Query" name:"PolicyGroupId"`
	CpuDownGradeDuration          requests.Integer                                `position:"Query" name:"CpuDownGradeDuration"`
	ClientType                    *[]ModifyPolicyGroupClientType                  `position:"Query" name:"ClientType"  type:"Repeated"`
	DeviceRedirects               *[]ModifyPolicyGroupDeviceRedirects             `position:"Query" name:"DeviceRedirects"  type:"Repeated"`
	ScopeValue                    *[]string                                       `position:"Query" name:"ScopeValue"  type:"Repeated"`
	RecordingFps                  requests.Integer                                `position:"Query" name:"RecordingFps"`
	WatermarkFontStyle            string                                          `position:"Query" name:"WatermarkFontStyle"`
	RecordingUserNotifyMessage    string                                          `position:"Query" name:"RecordingUserNotifyMessage"`
	SmoothEnhancement             string                                          `position:"Query" name:"SmoothEnhancement"`
	ColorEnhancement              string                                          `position:"Query" name:"ColorEnhancement"`
	EndUserGroupCoordinate        string                                          `position:"Query" name:"EndUserGroupCoordinate"`
	WatermarkDegree               string                                          `position:"Query" name:"WatermarkDegree"`
	DisplayMode                   string                                          `position:"Query" name:"DisplayMode"`
	RemoteCoordinate              string                                          `position:"Query" name:"RemoteCoordinate"`
	GpuAcceleration               string                                          `position:"Query" name:"GpuAcceleration"`
	Html5FileTransfer             string                                          `position:"Query" name:"Html5FileTransfer"`
	DeviceRules                   *[]ModifyPolicyGroupDeviceRules                 `position:"Query" name:"DeviceRules"  type:"Repeated"`
	VideoEncAvgKbps               requests.Integer                                `position:"Query" name:"VideoEncAvgKbps"`
	MemoryProtectedMode           string                                          `position:"Query" name:"MemoryProtectedMode"`
	VideoEncMaxQP                 requests.Integer                                `position:"Query" name:"VideoEncMaxQP"`
	MemoryDownGradeDuration       requests.Integer                                `position:"Query" name:"MemoryDownGradeDuration"`
	InternetCommunicationProtocol string                                          `position:"Query" name:"InternetCommunicationProtocol"`
	RecordingExpires              requests.Integer                                `position:"Query" name:"RecordingExpires"`
	PreemptLoginUser              *[]string                                       `position:"Query" name:"PreemptLoginUser"  type:"Repeated"`
	DomainList                    string                                          `position:"Query" name:"DomainList"`
	NetRedirect                   string                                          `position:"Query" name:"NetRedirect"`
	WatermarkTransparencyValue    requests.Integer                                `position:"Query" name:"WatermarkTransparencyValue"`
	WatermarkType                 string                                          `position:"Query" name:"WatermarkType"`
	AdminAccess                   string                                          `position:"Query" name:"AdminAccess"`
	RecordingDuration             requests.Integer                                `position:"Query" name:"RecordingDuration"`
	RevokeAccessPolicyRule        *[]ModifyPolicyGroupRevokeAccessPolicyRule      `position:"Query" name:"RevokeAccessPolicyRule"  type:"Repeated"`
	CloudHub                      string                                          `position:"Query" name:"CloudHub"`
	CameraRedirect                string                                          `position:"Query" name:"CameraRedirect"`
	EnableSessionRateLimiting     string                                          `position:"Query" name:"EnableSessionRateLimiting"`
	MemoryProcessors              *[]string                                       `position:"Query" name:"MemoryProcessors"  type:"Repeated"`
	VideoRedirect                 string                                          `position:"Query" name:"VideoRedirect"`
	AppContentProtection          string                                          `position:"Query" name:"AppContentProtection"`
	CpuProtectedMode              string                                          `position:"Query" name:"CpuProtectedMode"`
	VideoEncPeakKbps              requests.Integer                                `position:"Query" name:"VideoEncPeakKbps"`
	AuthorizeAccessPolicyRule     *[]ModifyPolicyGroupAuthorizeAccessPolicyRule   `position:"Query" name:"AuthorizeAccessPolicyRule"  type:"Repeated"`
	MaxReconnectTime              requests.Integer                                `position:"Query" name:"MaxReconnectTime"`
	WatermarkTransparency         string                                          `position:"Query" name:"WatermarkTransparency"`
	SessionMaxRateKbps            requests.Integer                                `position:"Query" name:"SessionMaxRateKbps"`
	Name                          string                                          `position:"Query" name:"Name"`
	MemorySampleDuration          requests.Integer                                `position:"Query" name:"MemorySampleDuration"`
	WatermarkCustomText           string                                          `position:"Query" name:"WatermarkCustomText"`
	PreemptLogin                  string                                          `position:"Query" name:"PreemptLogin"`
	UsbSupplyRedirectRule         *[]ModifyPolicyGroupUsbSupplyRedirectRule       `position:"Query" name:"UsbSupplyRedirectRule"  type:"Repeated"`
	WatermarkFontSize             requests.Integer                                `position:"Query" name:"WatermarkFontSize"`
	WatermarkAntiCam              string                                          `position:"Query" name:"WatermarkAntiCam"`
	Recording                     string                                          `position:"Query" name:"Recording"`
	DomainResolveRuleType         string                                          `position:"Query" name:"DomainResolveRuleType"`
	MemorySingleRateLimit         requests.Integer                                `position:"Query" name:"MemorySingleRateLimit"`
	WatermarkColumnAmount         requests.Integer                                `position:"Query" name:"WatermarkColumnAmount"`
	RecordContent                 string                                          `position:"Query" name:"RecordContent"`
	VideoEncPolicy                string                                          `position:"Query" name:"VideoEncPolicy"`
	Scope                         string                                          `position:"Query" name:"Scope"`
	WatermarkRowAmount            requests.Integer                                `position:"Query" name:"WatermarkRowAmount"`
	CpuProcessors                 *[]string                                       `position:"Query" name:"CpuProcessors"  type:"Repeated"`
	TargetFps                     requests.Integer                                `position:"Query" name:"TargetFps"`
	RecordContentExpires          requests.Integer                                `position:"Query" name:"RecordContentExpires"`
	RecordingAudio                string                                          `position:"Query" name:"RecordingAudio"`
	CpuRateLimit                  requests.Integer                                `position:"Query" name:"CpuRateLimit"`
	CpuSampleDuration             requests.Integer                                `position:"Query" name:"CpuSampleDuration"`
	Html5Access                   string                                          `position:"Query" name:"Html5Access"`
	VideoEncMinQP                 requests.Integer                                `position:"Query" name:"VideoEncMinQP"`
	VisualQuality                 string                                          `position:"Query" name:"VisualQuality"`
	RecordingEndTime              string                                          `position:"Query" name:"RecordingEndTime"`
	StreamingMode                 string                                          `position:"Query" name:"StreamingMode"`
	RecordingUserNotify           string                                          `position:"Query" name:"RecordingUserNotify"`
}

// ModifyPolicyGroupRevokeSecurityPolicyRule is a repeated param struct in ModifyPolicyGroupRequest
type ModifyPolicyGroupRevokeSecurityPolicyRule struct {
	PortRange   string `name:"PortRange"`
	IpProtocol  string `name:"IpProtocol"`
	Description string `name:"Description"`
	Type        string `name:"Type"`
	Priority    string `name:"Priority"`
	CidrIp      string `name:"CidrIp"`
	Policy      string `name:"Policy"`
}

// ModifyPolicyGroupAuthorizeSecurityPolicyRule is a repeated param struct in ModifyPolicyGroupRequest
type ModifyPolicyGroupAuthorizeSecurityPolicyRule struct {
	PortRange   string `name:"PortRange"`
	IpProtocol  string `name:"IpProtocol"`
	Description string `name:"Description"`
	Type        string `name:"Type"`
	Priority    string `name:"Priority"`
	CidrIp      string `name:"CidrIp"`
	Policy      string `name:"Policy"`
}

// ModifyPolicyGroupNetRedirectRule is a repeated param struct in ModifyPolicyGroupRequest
type ModifyPolicyGroupNetRedirectRule struct {
	RuleType string `name:"RuleType"`
	Domain   string `name:"Domain"`
	Policy   string `name:"Policy"`
}

// ModifyPolicyGroupDomainResolveRule is a repeated param struct in ModifyPolicyGroupRequest
type ModifyPolicyGroupDomainResolveRule struct {
	Domain      string `name:"Domain"`
	Description string `name:"Description"`
	Policy      string `name:"Policy"`
}

// ModifyPolicyGroupClientType is a repeated param struct in ModifyPolicyGroupRequest
type ModifyPolicyGroupClientType struct {
	ClientType string `name:"ClientType"`
	Status     string `name:"Status"`
}

// ModifyPolicyGroupDeviceRedirects is a repeated param struct in ModifyPolicyGroupRequest
type ModifyPolicyGroupDeviceRedirects struct {
	RedirectType string `name:"RedirectType"`
	DeviceType   string `name:"DeviceType"`
}

// ModifyPolicyGroupDeviceRules is a repeated param struct in ModifyPolicyGroupRequest
type ModifyPolicyGroupDeviceRules struct {
	DevicePid    string `name:"DevicePid"`
	DeviceName   string `name:"DeviceName"`
	DeviceVid    string `name:"DeviceVid"`
	RedirectType string `name:"RedirectType"`
	DeviceType   string `name:"DeviceType"`
	OptCommand   string `name:"OptCommand"`
}

// ModifyPolicyGroupRevokeAccessPolicyRule is a repeated param struct in ModifyPolicyGroupRequest
type ModifyPolicyGroupRevokeAccessPolicyRule struct {
	Description string `name:"Description"`
	CidrIp      string `name:"CidrIp"`
}

// ModifyPolicyGroupAuthorizeAccessPolicyRule is a repeated param struct in ModifyPolicyGroupRequest
type ModifyPolicyGroupAuthorizeAccessPolicyRule struct {
	Description string `name:"Description"`
	CidrIp      string `name:"CidrIp"`
}

// ModifyPolicyGroupUsbSupplyRedirectRule is a repeated param struct in ModifyPolicyGroupRequest
type ModifyPolicyGroupUsbSupplyRedirectRule struct {
	ProductId       string `name:"ProductId"`
	DeviceSubclass  string `name:"DeviceSubclass"`
	UsbRedirectType string `name:"UsbRedirectType"`
	VendorId        string `name:"VendorId"`
	Description     string `name:"Description"`
	DeviceClass     string `name:"DeviceClass"`
	UsbRuleType     string `name:"UsbRuleType"`
}

// ModifyPolicyGroupResponse is the response struct for api ModifyPolicyGroup
type ModifyPolicyGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyPolicyGroupRequest creates a request to invoke ModifyPolicyGroup API
func CreateModifyPolicyGroupRequest() (request *ModifyPolicyGroupRequest) {
	request = &ModifyPolicyGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "ModifyPolicyGroup", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyPolicyGroupResponse creates a response to parse from ModifyPolicyGroup response
func CreateModifyPolicyGroupResponse() (response *ModifyPolicyGroupResponse) {
	response = &ModifyPolicyGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
