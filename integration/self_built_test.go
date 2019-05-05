package integration

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms/self-built-cms"
)

func TestUploadCustomMetrics(t *testing.T) {
	client := self_cms.NewClient(map[string]string{
		"AccessKeyId":     os.Getenv("ACCESS_KEY_ID"),
		"AccessKeySecret": os.Getenv("ACCESS_KEY_SECRET"),
	})
	client.Endpoint__ = "metrichub-cms-cn-hangzhou.aliyuncs.com"
	metrics := &[]self_cms.CustomMetric{
		{
			GroupId:    101,
			MetricName: "metricName",
			Dimensions: map[string]interface{}{
				"sampleName1": "value1",
				"sampleName2": "value2",
			},
			Time:   "",
			Type:   0,
			Period: 60,
			Values: map[string]interface{}{
				"value": 10.5,
				"Sum":   100,
			},
		},
	}
	response, err := client.UploadCustomMetric(metrics)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.RequestId)
	assert.Equal(t, "success", response.Message)
}

func TestUploadCustomEvents(t *testing.T) {
	client := self_cms.NewClient(map[string]string{
		"AccessKeyId":     os.Getenv("ACCESS_KEY_ID"),
		"AccessKeySecret": os.Getenv("ACCESS_KEY_SECRET"),
	})
	client.Endpoint__ = "metrichub-cms-cn-hangzhou.aliyuncs.com"
	events := &[]self_cms.CustomEvent{
		{
			GroupId: 101,
			Content: "EventContent",
			Time:    "",
			Name:    "EventName",
		},
	}
	response, err := client.UploadCustomEvent(events)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.RequestId)
	assert.Equal(t, "success", response.Message)
}
