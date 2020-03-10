package acms_open

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":           "acm.cn-shenzhen.aliyuncs.com",
			"cn-beijing":            "acm.cn-beijing.aliyuncs.com",
			"cn-shenzhen-finance-1": "acm.cn-shenzhen-finance-1.aliyuncs.com",
			"cn-north-2-gov-1":      "acm.cn-north-2-gov-1.aliyuncs.com",
			"cn-qingdao":            "acm.cn-qingdao.aliyuncs.com",
			"cn-shanghai":           "acm.cn-shanghai.aliyuncs.com",
			"cn-shanghai-finance-1": "acm.cn-shanghai-finance-1.aliyuncs.com",
			"cn-hongkong":           "acm.cn-hongkong.aliyuncs.com",
			"cn-hangzhou-finance":   "acm.cn-hangzhou-finance.aliyuncs.com",
			"ap-southeast-1":        "acm.ap-southeast-1.aliyuncs.com",
			"ap-southeast-2":        "acm.ap-southeast-2.aliyuncs.com",
			"eu-central-1":          "acm.eu-central-1.aliyuncs.com",
			"us-east-1":             "acm.us-east-1.aliyuncs.com",
			"cn-zhangjiakou":        "acm.cn-zhangjiakou.aliyuncs.com",
			"us-west-1":             "acm.us-west-1.aliyuncs.com",
			"cn-hangzhou":           "acm.cn-hangzhou.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
