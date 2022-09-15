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

// TranscodeData is a nested struct in live response
type TranscodeData struct {
	TanscodeType string `json:"TanscodeType" xml:"TanscodeType"`
	Domain       string `json:"Domain" xml:"Domain"`
	Region       string `json:"Region" xml:"Region"`
	Duration     int    `json:"Duration" xml:"Duration"`
	TimeStamp    string `json:"TimeStamp" xml:"TimeStamp"`
	Fps          string `json:"Fps" xml:"Fps"`
	Resolution   string `json:"Resolution" xml:"Resolution"`
}
