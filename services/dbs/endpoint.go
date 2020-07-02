package dbs

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":           "dbs-api.cn-hangzhou.aliyuncs.com",
			"cn-beijing":            "dbs-api.cn-hangzhou.aliyuncs.com",
			"ap-south-1":            "dbs-api.ap-south-1.aliyuncs.com",
			"eu-west-1":             "dbs-api.eu-west-1.aliyuncs.com",
			"ap-northeast-1":        "dbs-api.ap-northeast-1.aliyuncs.com",
			"cn-shenzhen-finance-1": "dbs-api.cn-hangzhou.aliyuncs.com",
			"me-east-1":             "dbs-api.me-east-1.aliyuncs.com",
			"cn-chengdu":            "dbs-api.cn-chengdu.aliyuncs.com",
			"cn-qingdao":            "dbs-api.cn-hangzhou.aliyuncs.com",
			"cn-shanghai":           "dbs-api.cn-hangzhou.aliyuncs.com",
			"cn-shanghai-finance-1": "dbs-api.cn-hangzhou.aliyuncs.com",
			"cn-hongkong":           "dbs-api.cn-hangzhou.aliyuncs.com",
			"cn-hangzhou-finance":   "dbs-api.cn-hangzhou.aliyuncs.com",
			"ap-southeast-1":        "dbs-api.ap-southeast-1.aliyuncs.com",
			"ap-southeast-2":        "dbs-api.ap-southeast-2.aliyuncs.com",
			"ap-southeast-3":        "dbs-api.ap-southeast-3.aliyuncs.com",
			"eu-central-1":          "dbs-api.eu-central-1.aliyuncs.com",
			"cn-huhehaote":          "dbs-api.cn-huhehaote.aliyuncs.com",
			"ap-southeast-5":        "dbs-api.ap-southeast-5.aliyuncs.com",
			"us-east-1":             "dbs-api.cn-hangzhou.aliyuncs.com",
			"cn-zhangjiakou":        "dbs-api.cn-hangzhou.aliyuncs.com",
			"us-west-1":             "dbs-api.cn-hangzhou.aliyuncs.com",
			"cn-hangzhou":           "dbs-api.cn-hangzhou.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
