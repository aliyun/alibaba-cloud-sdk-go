package live

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

// DomainDetail is a nested struct in live response
type DomainDetail struct {
	CertName        string `json:"CertName" xml:"CertName"`
	Cname           string `json:"Cname" xml:"Cname"`
	Description     string `json:"Description" xml:"Description"`
	DomainName      string `json:"DomainName" xml:"DomainName"`
	DomainStatus    string `json:"DomainStatus" xml:"DomainStatus"`
	GmtCreated      string `json:"GmtCreated" xml:"GmtCreated"`
	GmtModified     string `json:"GmtModified" xml:"GmtModified"`
	LiveDomainType  string `json:"LiveDomainType" xml:"LiveDomainType"`
	Region          string `json:"Region" xml:"Region"`
	SSLProtocol     string `json:"SSLProtocol" xml:"SSLProtocol"`
	SSLPub          string `json:"SSLPub" xml:"SSLPub"`
	Scope           string `json:"Scope" xml:"Scope"`
	ResourceGroupId string `json:"ResourceGroupId" xml:"ResourceGroupId"`
}
