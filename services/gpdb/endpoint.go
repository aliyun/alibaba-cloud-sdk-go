package gpdb

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen":           "gpdb.aliyuncs.com",
			"cn-beijing":            "gpdb.aliyuncs.com",
			"cn-shenzhen-finance-1": "gpdb.aliyuncs.com",
			"cn-north-2-gov-1":      "gpdb.aliyuncs.com",
			"cn-qingdao":            "gpdb.aliyuncs.com",
			"cn-shanghai":           "gpdb.aliyuncs.com",
			"cn-shanghai-finance-1": "gpdb.aliyuncs.com",
			"cn-hongkong":           "gpdb.aliyuncs.com",
			"cn-hangzhou-finance":   "gpdb.aliyuncs.com",
			"ap-southeast-1":        "gpdb.aliyuncs.com",
			"us-east-1":             "gpdb.aliyuncs.com",
			"us-west-1":             "gpdb.aliyuncs.com",
			"cn-hangzhou":           "gpdb.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
