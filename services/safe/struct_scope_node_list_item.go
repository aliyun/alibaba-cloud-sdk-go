package safe

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

// ScopeNodeListItem is a nested struct in safe response
type ScopeNodeListItem struct {
	RuleId    int64  `json:"RuleId" xml:"RuleId"`
	Type      string `json:"Type" xml:"Type"`
	Level1    string `json:"Level1" xml:"Level1"`
	Level2    string `json:"Level2" xml:"Level2"`
	Level3    string `json:"Level3" xml:"Level3"`
	Level4    string `json:"Level4" xml:"Level4"`
	Level5    string `json:"Level5" xml:"Level5"`
	LeafLevel string `json:"LeafLevel" xml:"LeafLevel"`
	Path      string `json:"Path" xml:"Path"`
}
