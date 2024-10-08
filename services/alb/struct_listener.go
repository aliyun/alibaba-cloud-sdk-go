package alb

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

// Listener is a nested struct in alb response
type Listener struct {
	GzipEnabled           bool                 `json:"GzipEnabled" xml:"GzipEnabled"`
	Http2Enabled          bool                 `json:"Http2Enabled" xml:"Http2Enabled"`
	IdleTimeout           int                  `json:"IdleTimeout" xml:"IdleTimeout"`
	ListenerDescription   string               `json:"ListenerDescription" xml:"ListenerDescription"`
	ListenerId            string               `json:"ListenerId" xml:"ListenerId"`
	ListenerPort          int                  `json:"ListenerPort" xml:"ListenerPort"`
	ListenerProtocol      string               `json:"ListenerProtocol" xml:"ListenerProtocol"`
	ListenerStatus        string               `json:"ListenerStatus" xml:"ListenerStatus"`
	LoadBalancerId        string               `json:"LoadBalancerId" xml:"LoadBalancerId"`
	ServiceManagedEnabled bool                 `json:"ServiceManagedEnabled" xml:"ServiceManagedEnabled"`
	ServiceManagedMode    string               `json:"ServiceManagedMode" xml:"ServiceManagedMode"`
	RequestTimeout        int                  `json:"RequestTimeout" xml:"RequestTimeout"`
	SecurityPolicyId      string               `json:"SecurityPolicyId" xml:"SecurityPolicyId"`
	LogConfig             LogConfig            `json:"LogConfig" xml:"LogConfig"`
	QuicConfig            QuicConfig           `json:"QuicConfig" xml:"QuicConfig"`
	XForwardedForConfig   XForwardedForConfig  `json:"XForwardedForConfig" xml:"XForwardedForConfig"`
	DefaultActions        []DefaultAction      `json:"DefaultActions" xml:"DefaultActions"`
	Tags                  []Tag                `json:"Tags" xml:"Tags"`
	AssociatedResources   []AssociatedResource `json:"AssociatedResources" xml:"AssociatedResources"`
}
