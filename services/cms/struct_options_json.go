package cms

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

// OptionsJson is a nested struct in cms response
type OptionsJson struct {
	Password               string  `json:"password" xml:"password"`
	RequestFormat          string  `json:"request_format" xml:"request_format"`
	ExpectValue            string  `json:"expect_value" xml:"expect_value"`
	ResponseContent        string  `json:"response_content" xml:"response_content"`
	FailureRate            float64 `json:"failure_rate" xml:"failure_rate"`
	TimeOut                int64   `json:"time_out" xml:"time_out"`
	Header                 string  `json:"header" xml:"header"`
	Cookie                 string  `json:"cookie" xml:"cookie"`
	Port                   int     `json:"port" xml:"port"`
	PingNum                int     `json:"ping_num" xml:"ping_num"`
	Authentication         int     `json:"authentication" xml:"authentication"`
	HttpMethod             string  `json:"http_method" xml:"http_method"`
	MatchRule              int     `json:"match_rule" xml:"match_rule"`
	DnsMatchRule           string  `json:"dns_match_rule" xml:"dns_match_rule"`
	RequestContent         string  `json:"request_content" xml:"request_content"`
	Username               string  `json:"username" xml:"username"`
	Traceroute             int64   `json:"traceroute" xml:"traceroute"`
	ResponseFormat         string  `json:"response_format" xml:"response_format"`
	DnsType                string  `json:"dns_type" xml:"dns_type"`
	GroupId                string  `json:"group_id" xml:"group_id"`
	DnsServer              string  `json:"dns_server" xml:"dns_server"`
	EnableOperatorDns      bool    `json:"enable_operator_dns" xml:"enable_operator_dns"`
	Attempts               int64   `json:"attempts" xml:"attempts"`
	Protocol               string  `json:"protocol" xml:"protocol"`
	ProxyProtocol          bool    `json:"proxy_protocol" xml:"proxy_protocol"`
	AcceptableResponseCode string  `json:"acceptable_response_code" xml:"acceptable_response_code"`
	IsBase64Encode         string  `json:"isBase64Encode" xml:"isBase64Encode"`
	CertVerify             bool    `json:"cert_verify" xml:"cert_verify"`
	UnfollowRedirect       bool    `json:"unfollow_redirect" xml:"unfollow_redirect"`
	DiagnosisMtr           bool    `json:"diagnosis_mtr" xml:"diagnosis_mtr"`
	DiagnosisPing          bool    `json:"diagnosis_ping" xml:"diagnosis_ping"`
	RetryDelay             int     `json:"retry_delay" xml:"retry_delay"`
}
