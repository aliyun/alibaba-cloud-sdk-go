package ens

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

// Ipv4sItem is a nested struct in ens response
type Ipv4sItem struct {
	Display             string `json:"Display" xml:"Display"`
	Isp                 string `json:"Isp" xml:"Isp"`
	OversellRatio       int64  `json:"OversellRatio" xml:"OversellRatio"`
	Remain              int64  `json:"Remain" xml:"Remain"`
	Reserved            int64  `json:"Reserved" xml:"Reserved"`
	ReserveDisable      bool   `json:"ReserveDisable" xml:"ReserveDisable"`
	ReserveDisableTotal int64  `json:"ReserveDisableTotal" xml:"ReserveDisableTotal"`
	StatusDisable       bool   `json:"StatusDisable" xml:"StatusDisable"`
	StatusDisableTotal  int64  `json:"StatusDisableTotal" xml:"StatusDisableTotal"`
	Total               int64  `json:"Total" xml:"Total"`
	Type                string `json:"Type" xml:"Type"`
	Used                int64  `json:"Used" xml:"Used"`
	UsedRatio           int64  `json:"UsedRatio" xml:"UsedRatio"`
	Vlan                string `json:"Vlan" xml:"Vlan"`
}
