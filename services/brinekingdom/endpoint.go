package brinekingdom

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":           "brinekingdom.aliyuncs.com",
			"cn-beijing":            "brinekingdom.aliyuncs.com",
			"ap-south-1":            "brinekingdom.aliyuncs.com",
			"eu-west-1":             "brinekingdom.aliyuncs.com",
			"ap-northeast-1":        "brinekingdom.aliyuncs.com",
			"cn-shenzhen-finance-1": "brinekingdom.aliyuncs.com",
			"me-east-1":             "brinekingdom.aliyuncs.com",
			"cn-chengdu":            "brinekingdom.aliyuncs.com",
			"cn-north-2-gov-1":      "brinekingdom.aliyuncs.com",
			"cn-qingdao":            "brinekingdom.aliyuncs.com",
			"cn-shanghai":           "brinekingdom.aliyuncs.com",
			"cn-shanghai-finance-1": "brinekingdom.aliyuncs.com",
			"cn-hongkong":           "brinekingdom.aliyuncs.com",
			"cn-hangzhou-finance":   "brinekingdom.aliyuncs.com",
			"ap-southeast-1":        "brinekingdom.aliyuncs.com",
			"ap-southeast-2":        "brinekingdom.aliyuncs.com",
			"ap-southeast-3":        "brinekingdom.aliyuncs.com",
			"eu-central-1":          "brinekingdom.aliyuncs.com",
			"cn-huhehaote":          "brinekingdom.aliyuncs.com",
			"ap-southeast-5":        "brinekingdom.aliyuncs.com",
			"us-east-1":             "brinekingdom.aliyuncs.com",
			"cn-zhangjiakou":        "brinekingdom.aliyuncs.com",
			"us-west-1":             "brinekingdom.aliyuncs.com",
			"cn-hangzhou":           "brinekingdom.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
