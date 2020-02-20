package retailcloud

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "retailcloud.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "retailcloud.aliyuncs.com",
			"cn-beijing":                  "retailcloud.aliyuncs.com",
			"cn-shanghai-inner":           "retailcloud.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "retailcloud.aliyuncs.com",
			"cn-north-2-gov-1":            "retailcloud.aliyuncs.com",
			"cn-yushanfang":               "retailcloud.aliyuncs.com",
			"cn-qingdao-nebula":           "retailcloud.aliyuncs.com",
			"cn-beijing-finance-pop":      "retailcloud.aliyuncs.com",
			"cn-wuhan":                    "retailcloud.aliyuncs.com",
			"cn-zhangjiakou":              "retailcloud.aliyuncs.com",
			"us-west-1":                   "retailcloud.aliyuncs.com",
			"rus-west-1-pop":              "retailcloud.aliyuncs.com",
			"cn-shanghai-et15-b01":        "retailcloud.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "retailcloud.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "retailcloud.aliyuncs.com",
			"ap-northeast-1":              "retailcloud.aliyuncs.com",
			"cn-shanghai-et2-b01":         "retailcloud.aliyuncs.com",
			"ap-southeast-1":              "retailcloud.aliyuncs.com",
			"ap-southeast-2":              "retailcloud.aliyuncs.com",
			"ap-southeast-3":              "retailcloud.aliyuncs.com",
			"ap-southeast-5":              "retailcloud.aliyuncs.com",
			"us-east-1":                   "retailcloud.aliyuncs.com",
			"cn-shenzhen-inner":           "retailcloud.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "retailcloud.aliyuncs.com",
			"cn-beijing-gov-1":            "retailcloud.aliyuncs.com",
			"ap-south-1":                  "retailcloud.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "retailcloud.aliyuncs.com",
			"cn-haidian-cm12-c01":         "retailcloud.aliyuncs.com",
			"cn-qingdao":                  "retailcloud.aliyuncs.com",
			"cn-hongkong-finance-pop":     "retailcloud.aliyuncs.com",
			"cn-shanghai":                 "retailcloud.aliyuncs.com",
			"cn-shanghai-finance-1":       "retailcloud.aliyuncs.com",
			"cn-hongkong":                 "retailcloud.aliyuncs.com",
			"eu-central-1":                "retailcloud.aliyuncs.com",
			"cn-shenzhen":                 "retailcloud.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "retailcloud.aliyuncs.com",
			"eu-west-1":                   "retailcloud.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "retailcloud.aliyuncs.com",
			"eu-west-1-oxs":               "retailcloud.aliyuncs.com",
			"cn-beijing-finance-1":        "retailcloud.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "retailcloud.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "retailcloud.aliyuncs.com",
			"cn-shenzhen-finance-1":       "retailcloud.aliyuncs.com",
			"me-east-1":                   "retailcloud.aliyuncs.com",
			"cn-chengdu":                  "retailcloud.aliyuncs.com",
			"cn-hangzhou-test-306":        "retailcloud.aliyuncs.com",
			"cn-hangzhou-finance":         "retailcloud.aliyuncs.com",
			"cn-beijing-nu16-b01":         "retailcloud.aliyuncs.com",
			"cn-edge-1":                   "retailcloud.aliyuncs.com",
			"cn-huhehaote":                "retailcloud.aliyuncs.com",
			"cn-fujian":                   "retailcloud.aliyuncs.com",
			"ap-northeast-2-pop":          "retailcloud.aliyuncs.com",
			"cn-hangzhou":                 "retailcloud.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
