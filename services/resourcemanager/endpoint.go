package resourcemanager

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "central"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":           "resourcemanager.cn-shenzhen.aliyuncs.com",
			"cn-beijing":            "resourcemanager.cn-beijing.aliyuncs.com",
			"cn-wulanchabu":         "resourcemanager.cn-wulanchabu.aliyuncs.com",
			"ap-south-1":            "resourcemanager.ap-south-1.aliyuncs.com",
			"eu-west-1":             "resourcemanager.eu-west-1.aliyuncs.com",
			"ap-northeast-1":        "resourcemanager.ap-northeast-1.aliyuncs.com",
			"cn-shenzhen-finance-1": "resourcemanager.cn-shenzhen-finance-1.aliyuncs.com",
			"me-east-1":             "resourcemanager.me-east-1.aliyuncs.com",
			"cn-chengdu":            "resourcemanager.cn-chengdu.aliyuncs.com",
			"cn-north-2-gov-1":      "resourcemanager.cn-north-2-gov-1.aliyuncs.com",
			"cn-qingdao":            "resourcemanager.cn-qingdao.aliyuncs.com",
			"cn-shanghai-finance-1": "resourcemanager.cn-shanghai-finance-1.aliyuncs.com",
			"cn-hangzhou-finance":   "resourcemanager.cn-hangzhou-finance.aliyuncs.com",
			"cn-hongkong":           "resourcemanager.cn-hongkong.aliyuncs.com",
			"ap-southeast-1":        "resourcemanager.ap-southeast-1.aliyuncs.com",
			"ap-southeast-2":        "resourcemanager.ap-southeast-2.aliyuncs.com",
			"ap-southeast-3":        "resourcemanager.ap-southeast-3.aliyuncs.com",
			"eu-central-1":          "resourcemanager.eu-central-1.aliyuncs.com",
			"cn-huhehaote":          "resourcemanager.cn-huhehaote.aliyuncs.com",
			"ap-southeast-5":        "resourcemanager.ap-southeast-5.aliyuncs.com",
			"us-east-1":             "resourcemanager.us-east-1.aliyuncs.com",
			"cn-zhangjiakou":        "resourcemanager.cn-zhangjiakou.aliyuncs.com",
			"us-west-1":             "resourcemanager.us-west-1.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
