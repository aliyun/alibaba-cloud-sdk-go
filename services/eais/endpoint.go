package eais

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "eais.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "eais.aliyuncs.com",
			"cn-shanghai-inner":           "eais.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "eais.aliyuncs.com",
			"cn-north-2-gov-1":            "eais.aliyuncs.com",
			"cn-yushanfang":               "eais.aliyuncs.com",
			"cn-qingdao-nebula":           "eais.aliyuncs.com",
			"cn-beijing-finance-pop":      "eais.aliyuncs.com",
			"cn-wuhan":                    "eais.aliyuncs.com",
			"cn-zhangjiakou":              "eais.aliyuncs.com",
			"cn-zhangbei":                 "eais.aliyuncs.com",
			"rus-west-1-pop":              "eais.aliyuncs.com",
			"cn-shanghai-et15-b01":        "eais.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "eais.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "eais.aliyuncs.com",
			"ap-northeast-1":              "eais.aliyuncs.com",
			"cn-huhehaote-nebula-1":       "eais.aliyuncs.com",
			"cn-shanghai-et2-b01":         "eais.aliyuncs.com",
			"ap-southeast-1":              "eais.aliyuncs.com",
			"ap-southeast-2":              "eais.aliyuncs.com",
			"ap-southeast-3":              "eais.aliyuncs.com",
			"ap-southeast-5":              "eais.aliyuncs.com",
			"us-east-1":                   "eais.aliyuncs.com",
			"cn-shenzhen-inner":           "eais.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "eais.aliyuncs.com",
			"cn-beijing-gov-1":            "eais.aliyuncs.com",
			"cn-wulanchabu":               "eais.aliyuncs.com",
			"ap-south-1":                  "eais.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "eais.aliyuncs.com",
			"cn-haidian-cm12-c01":         "eais.aliyuncs.com",
			"cn-qingdao":                  "eais.aliyuncs.com",
			"cn-hongkong-finance-pop":     "eais.aliyuncs.com",
			"cn-shanghai-finance-1":       "eais.aliyuncs.com",
			"cn-hongkong":                 "eais.aliyuncs.com",
			"eu-central-1":                "eais.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "eais.aliyuncs.com",
			"eu-west-1":                   "eais.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "eais.aliyuncs.com",
			"eu-west-1-oxs":               "eais.aliyuncs.com",
			"cn-beijing-finance-1":        "eais.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "eais.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "eais.aliyuncs.com",
			"cn-shenzhen-finance-1":       "eais.aliyuncs.com",
			"me-east-1":                   "eais.aliyuncs.com",
			"cn-hangzhou-test-306":        "eais.aliyuncs.com",
			"cn-hangzhou-finance":         "eais.aliyuncs.com",
			"cn-beijing-nu16-b01":         "eais.aliyuncs.com",
			"cn-edge-1":                   "eais.aliyuncs.com",
			"cn-huhehaote":                "eais.aliyuncs.com",
			"cn-fujian":                   "eais.aliyuncs.com",
			"ap-northeast-2-pop":          "eais.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
