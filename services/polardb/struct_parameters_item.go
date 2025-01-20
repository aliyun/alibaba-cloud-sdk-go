package polardb

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

// ParametersItem is a nested struct in polardb response
type ParametersItem struct {
	RdsParameterName         string `json:"rdsParameterName" xml:"rdsParameterName"`
	RdsParameterValue        string `json:"rdsParameterValue" xml:"rdsParameterValue"`
	RdsParameterOptional     string `json:"rdsParameterOptional" xml:"rdsParameterOptional"`
	DistParameterName        string `json:"distParameterName" xml:"distParameterName"`
	DistParameterValue       string `json:"distParameterValue" xml:"distParameterValue"`
	DistParameterOptional    string `json:"distParameterOptional" xml:"distParameterOptional"`
	IsEqual                  string `json:"IsEqual" xml:"IsEqual"`
	DistParameterDescription string `json:"distParameterDescription" xml:"distParameterDescription"`
	RdsParameterDescription  string `json:"rdsParameterDescription" xml:"rdsParameterDescription"`
	IsRdsKey                 string `json:"IsRdsKey" xml:"IsRdsKey"`
	IsPolarDBKey             string `json:"IsPolarDBKey" xml:"IsPolarDBKey"`
	IsInstancePolarDBKey     string `json:"IsInstancePolarDBKey" xml:"IsInstancePolarDBKey"`
	IsInstanceRdsKey         string `json:"IsInstanceRdsKey" xml:"IsInstanceRdsKey"`
}
