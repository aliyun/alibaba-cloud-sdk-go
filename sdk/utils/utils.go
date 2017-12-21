/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"net/url"
	"reflect"
	"strconv"
	"time"
)

func GetUUIDV4() (uuidHex string) {
	uuidV4 := uuid.NewV4()
	uuidHex = hex.EncodeToString(uuidV4.Bytes())
	return
}

func GetMD5Base64(bytes []byte) (base64Value string) {
	md5Ctx := md5.New()
	md5Ctx.Write(bytes)
	md5Value := md5Ctx.Sum(nil)
	base64Value = base64.StdEncoding.EncodeToString(md5Value)
	return
}

func GetTimeInFormatISO8601() (timeStr string) {
	gmt, err := time.LoadLocation("GMT")
	if err != nil {
		panic(err)
	}
	return time.Now().In(gmt).Format("2006-01-02T15:04:05Z")
}

func GetTimeInFormatRFC2616() (timeStr string) {
	gmt, err := time.LoadLocation("GMT")
	if err != nil {
		panic(err)
	}
	return time.Now().In(gmt).Format("Mon, 02 Jan 2006 15:04:05 GMT")
}

func GetUrlFormedMap(source map[string]string) (urlEncoded string) {
	urlEncoder := url.Values{}
	for key, value := range source {
		urlEncoder.Add(key, value)
	}
	urlEncoded = urlEncoder.Encode()
	return
}

func GetFromJsonString(jsonString, key string) (result string, err error) {
	var responseMap map[string]*json.RawMessage
	err = json.Unmarshal([]byte(jsonString), &responseMap)
	if err != nil {
		return
	}
	fmt.Println(string(*responseMap[key]))
	err = json.Unmarshal(*responseMap[key], &result)
	return
}

func InitStructWithDefaultTag(bean interface{}) {
	configType := reflect.TypeOf(bean)
	for i := 0; i < configType.Elem().NumField(); i++ {
		field := configType.Elem().Field(i)
		defaultValue := field.Tag.Get("default")
		if defaultValue == "" {
			continue
		}
		setter := reflect.ValueOf(bean).Elem().Field(i)
		switch field.Type.String() {
		case "int":
			intValue, _ := strconv.ParseInt(defaultValue, 10, 64)
			setter.SetInt(intValue)
		case "string":
			setter.SetString(defaultValue)
		case "bool":
			boolValue, _ := strconv.ParseBool(defaultValue)
			setter.SetBool(boolValue)
		}
	}
}

type IZero interface {
	IsZero() bool
}

func IsEmptyValue(v reflect.Value) bool {
	if v.CanInterface() {
		if zeroChecker, ok := v.Interface().(IZero); ok {
			return zeroChecker.IsZero()
		}
	}

	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

func ConvertToString(v interface{}) (string, error) {
	if stringer, ok := v.(fmt.Stringer); ok {
		return stringer.String(), nil
	}
	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.String:
		return v.(string), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8:
		return fmt.Sprintf("%d", v), nil
	case reflect.Bool:
		return strconv.FormatBool(v.(bool)), nil
	case reflect.Float32:
		return strconv.FormatFloat(float64(v.(float32)), 'f', -1, 32), nil
	case reflect.Float64:
		return strconv.FormatFloat(v.(float64), 'f', -1, 64), nil
	default:
		return "", fmt.Errorf("unsupport type")
	}
}
