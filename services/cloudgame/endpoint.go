package cloudgame

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shanghai-internal-test-1": "cloudgame.aliyuncs.com",
			"cn-beijing-gov-1":            "cloudgame.aliyuncs.com",
			"cn-shenzhen-su18-b01":        "cloudgame.aliyuncs.com",
			"cn-wulanchabu":               "cloudgame.aliyuncs.com",
			"ap-south-1":                  "cloudgame.aliyuncs.com",
			"cn-shanghai-inner":           "cloudgame.aliyuncs.com",
			"cn-shenzhen-st4-d01":         "cloudgame.aliyuncs.com",
			"cn-haidian-cm12-c01":         "cloudgame.aliyuncs.com",
			"cn-hangzhou-internal-prod-1": "cloudgame.aliyuncs.com",
			"cn-north-2-gov-1":            "cloudgame.aliyuncs.com",
			"cn-yushanfang":               "cloudgame.aliyuncs.com",
			"cn-hongkong-finance-pop":     "cloudgame.aliyuncs.com",
			"cn-qingdao-nebula":           "cloudgame.aliyuncs.com",
			"cn-shanghai-finance-1":       "cloudgame.aliyuncs.com",
			"cn-hongkong":                 "cloudgame.aliyuncs.com",
			"cn-beijing-finance-pop":      "cloudgame.aliyuncs.com",
			"cn-wuhan":                    "cloudgame.aliyuncs.com",
			"eu-central-1":                "cloudgame.aliyuncs.com",
			"us-west-1":                   "cloudgame.aliyuncs.com",
			"cn-zhengzhou-nebula-1":       "cloudgame.aliyuncs.com",
			"rus-west-1-pop":              "cloudgame.aliyuncs.com",
			"cn-shanghai-et15-b01":        "cloudgame.aliyuncs.com",
			"cn-hangzhou-bj-b01":          "cloudgame.aliyuncs.com",
			"eu-west-1":                   "cloudgame.aliyuncs.com",
			"cn-hangzhou-internal-test-1": "cloudgame.aliyuncs.com",
			"eu-west-1-oxs":               "cloudgame.aliyuncs.com",
			"cn-zhangbei-na61-b01":        "cloudgame.aliyuncs.com",
			"cn-beijing-finance-1":        "cloudgame.aliyuncs.com",
			"cn-hangzhou-internal-test-3": "cloudgame.aliyuncs.com",
			"ap-northeast-1":              "cloudgame.aliyuncs.com",
			"cn-hangzhou-internal-test-2": "cloudgame.aliyuncs.com",
			"cn-shenzhen-finance-1":       "cloudgame.aliyuncs.com",
			"me-east-1":                   "cloudgame.aliyuncs.com",
			"cn-hangzhou-test-306":        "cloudgame.aliyuncs.com",
			"cn-huhehaote-nebula-1":       "cloudgame.aliyuncs.com",
			"cn-shanghai-et2-b01":         "cloudgame.aliyuncs.com",
			"cn-hangzhou-finance":         "cloudgame.aliyuncs.com",
			"ap-southeast-1":              "cloudgame.aliyuncs.com",
			"cn-beijing-nu16-b01":         "cloudgame.aliyuncs.com",
			"cn-edge-1":                   "cloudgame.aliyuncs.com",
			"ap-southeast-2":              "cloudgame.aliyuncs.com",
			"ap-southeast-3":              "cloudgame.aliyuncs.com",
			"ap-southeast-5":              "cloudgame.aliyuncs.com",
			"cn-fujian":                   "cloudgame.aliyuncs.com",
			"us-east-1":                   "cloudgame.aliyuncs.com",
			"ap-northeast-2-pop":          "cloudgame.aliyuncs.com",
			"cn-shenzhen-inner":           "cloudgame.aliyuncs.com",
			"cn-zhangjiakou-na62-a01":     "cloudgame.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
