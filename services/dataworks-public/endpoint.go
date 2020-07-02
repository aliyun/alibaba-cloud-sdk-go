package dataworks_public

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":           "dataworks.cn-shenzhen.aliyuncs.com",
			"cn-beijing":            "dataworks.cn-beijing.aliyuncs.com",
			"ap-south-1":            "dataworks.ap-south-1.aliyuncs.com",
			"eu-west-1":             "dataworks.eu-west-1.aliyuncs.com",
			"ap-northeast-1":        "dataworks.ap-northeast-1.aliyuncs.com",
			"cn-shenzhen-finance-1": "dataworks.aliyuncs.com",
			"me-east-1":             "dataworks.me-east-1.aliyuncs.com",
			"cn-chengdu":            "dataworks.cn-chengdu.aliyuncs.com",
			"cn-north-2-gov-1":      "dataworks.aliyuncs.com",
			"cn-qingdao":            "dataworks.aliyuncs.com",
			"cn-shanghai":           "dataworks.cn-shanghai.aliyuncs.com",
			"cn-shanghai-finance-1": "dataworks.aliyuncs.com",
			"cn-hongkong":           "dataworks.cn-hongkong.aliyuncs.com",
			"cn-hangzhou-finance":   "dataworks.aliyuncs.com",
			"ap-southeast-1":        "dataworks.ap-southeast-1.aliyuncs.com",
			"ap-southeast-2":        "dataworks.ap-southeast-2.aliyuncs.com",
			"ap-southeast-3":        "dataworks.ap-southeast-3.aliyuncs.com",
			"eu-central-1":          "dataworks.eu-central-1.aliyuncs.com",
			"cn-huhehaote":          "dataworks.aliyuncs.com",
			"ap-southeast-5":        "dataworks.ap-southeast-5.aliyuncs.com",
			"us-east-1":             "dataworks.us-east-1.aliyuncs.com",
			"cn-zhangjiakou":        "dataworks.aliyuncs.com",
			"us-west-1":             "dataworks.us-west-1.aliyuncs.com",
			"cn-hangzhou":           "dataworks.cn-hangzhou.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
