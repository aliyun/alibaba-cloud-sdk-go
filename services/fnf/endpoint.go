package fnf

// EndpointMap Endpoint Data
var EndpointMap map[string]string

// EndpointType regional or central
var EndpointType = "regional"

// GetEndpointMap Get Endpoint Data Map
func GetEndpointMap() map[string]string {
	if EndpointMap == nil {
		EndpointMap = map[string]string{
			"cn-shenzhen": "cn-shenzhen.fnf.aliyuncs.com",
			"cn-beijing":  "cn-beijing.fnf.aliyuncs.com",
			"cn-shanghai": "cn-shanghai.fnf.aliyuncs.com",
			"cn-hangzhou": "cn-hangzhou.fnf.aliyuncs.com",
		}
	}
	return EndpointMap
}

// GetEndpointType Get Endpoint Type Value
func GetEndpointType() string {
	return EndpointType
}
