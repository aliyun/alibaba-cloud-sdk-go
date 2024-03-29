package qualitycheck

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

// ParamInGetRuleDetail is a nested struct in qualitycheck response
type ParamInGetRuleDetail struct {
	DifferentRole         bool                    `json:"DifferentRole" xml:"DifferentRole"`
	Regex                 string                  `json:"Regex" xml:"Regex"`
	TargetRole            string                  `json:"TargetRole" xml:"TargetRole"`
	VelocityInMint        int                     `json:"VelocityInMint" xml:"VelocityInMint"`
	Average               bool                    `json:"Average" xml:"Average"`
	KeywordExtension      bool                    `json:"KeywordExtension" xml:"KeywordExtension"`
	Score                 int                     `json:"Score" xml:"Score"`
	NotRegex              string                  `json:"NotRegex" xml:"NotRegex"`
	CompareOperator       string                  `json:"CompareOperator" xml:"CompareOperator"`
	DelayTime             int                     `json:"DelayTime" xml:"DelayTime"`
	KeywordMatchSize      int                     `json:"KeywordMatchSize" xml:"KeywordMatchSize"`
	HitTime               int                     `json:"HitTime" xml:"HitTime"`
	BeginType             string                  `json:"BeginType" xml:"BeginType"`
	Target                int                     `json:"Target" xml:"Target"`
	MaxEmotionChangeValue int                     `json:"MaxEmotionChangeValue" xml:"MaxEmotionChangeValue"`
	Threshold             float64                 `json:"Threshold" xml:"Threshold"`
	From                  int                     `json:"From" xml:"From"`
	FromEnd               bool                    `json:"FromEnd" xml:"FromEnd"`
	MinWordSize           int                     `json:"MinWordSize" xml:"MinWordSize"`
	InSentence            bool                    `json:"InSentence" xml:"InSentence"`
	Phrase                string                  `json:"Phrase" xml:"Phrase"`
	SimilarityThreshold   float64                 `json:"Similarity_threshold" xml:"Similarity_threshold"`
	CheckType             int                     `json:"CheckType" xml:"CheckType"`
	Interval              int                     `json:"Interval" xml:"Interval"`
	ContextChatMatch      bool                    `json:"ContextChatMatch" xml:"ContextChatMatch"`
	Excludes              ExcludesInGetRuleDetail `json:"Excludes" xml:"Excludes"`
	AntModelInfo          AntModelInfo            `json:"AntModelInfo" xml:"AntModelInfo"`
	Pvalues               Pvalues                 `json:"Pvalues" xml:"Pvalues"`
	References            References              `json:"References" xml:"References"`
	SimilarlySentences    SimilarlySentences      `json:"SimilarlySentences" xml:"SimilarlySentences"`
	OperKeyWords          OperKeyWords            `json:"OperKeyWords" xml:"OperKeyWords"`
}
