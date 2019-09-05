package bssopenapi

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":           "business.aliyuncs.com",
			"cn-beijing":            "business.aliyuncs.com",
			"ap-south-1":            "business.ap-southeast-1.aliyuncs.com",
			"eu-west-1":             "business.ap-southeast-1.aliyuncs.com",
			"ap-northeast-1":        "business.ap-southeast-1.aliyuncs.com",
			"cn-shenzhen-finance-1": "business.aliyuncs.com",
			"me-east-1":             "business.ap-southeast-1.aliyuncs.com",
			"cn-chengdu":            "business.aliyuncs.com",
			"cn-north-2-gov-1":      "business.aliyuncs.com",
			"cn-qingdao":            "business.aliyuncs.com",
			"cn-shanghai":           "business.aliyuncs.com",
			"cn-shanghai-finance-1": "business.aliyuncs.com",
			"cn-hongkong":           "business.aliyuncs.com",
			"cn-hangzhou-finance":   "business.aliyuncs.com",
			"ap-southeast-1":        "business.ap-southeast-1.aliyuncs.com",
			"ap-southeast-2":        "business.ap-southeast-1.aliyuncs.com",
			"ap-southeast-3":        "business.ap-southeast-1.aliyuncs.com",
			"eu-central-1":          "business.ap-southeast-1.aliyuncs.com",
			"cn-huhehaote":          "business.aliyuncs.com",
			"ap-southeast-5":        "business.ap-southeast-1.aliyuncs.com",
			"us-east-1":             "business.ap-southeast-1.aliyuncs.com",
			"cn-zhangjiakou":        "business.aliyuncs.com",
			"us-west-1":             "business.ap-southeast-1.aliyuncs.com",
			"cn-hangzhou":           "business.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
