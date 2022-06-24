package weather

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Weather struct {
	key string
}

func NewWeather(key string) Weather {
	return Weather{key: key}
}

func (w Weather) GetWeather(city, extensions, format string) (string, error) {

	// 对 `format` 与 `extensions` 参数进行检查
	if !(format == "xml" || format == "json") {
		return "", fmt.Errorf("invalid response format: %s", format)
	}

	if !(extensions == "base" || extensions == "all") {
		return "", fmt.Errorf("invalid type value(base/all): %s", format)
	}

	// 组装 query 参数
	data := url.Values{}
	data.Set("key", w.key)
	data.Set("city", city)
	data.Set("output", format)
	data.Set("extensions", extensions)

	// 定义请求 uri
	uri := "https://restapi.amap.com/v3/weather/weatherInfo"

	// 把 string 转换成 url
	u, err := url.ParseRequestURI(uri)

	if err != nil {
		return "", fmt.Errorf("parse url requestUrl failed,err：%s", err)
	}

	// url encode
	u.RawQuery = data.Encode()
	fmt.Println(u.String())

	resp, err := http.Get(u.String())

	if err != nil {
		return "", fmt.Errorf("get failed, err：%s", err)
	}

	defer resp.Body.Close()

	//读取内容
	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", fmt.Errorf("get resp failed, err: %s", err)
	}

	return string(b), nil
}
