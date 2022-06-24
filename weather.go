package weather

import (
	"encoding/json"
	"fmt"
	"github.com/Simoon-F/amap-weather/types"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Weather struct {
	key string
}

func NewWeather(key string) *Weather {
	return &Weather{key: key}
}

// GetLiveWeather 获取实时天气
func (w *Weather) GetLiveWeather(city, format string) (types.WeatherResponse, error) {
	return w.getWeather(city, "base", format)
}

// GetForecastsWeather 获取天气预报
func (w *Weather) GetForecastsWeather(city, format string) (types.WeatherResponse, error) {
	return w.getWeather(city, "all", format)
}

func (w *Weather) getWeather(city, extensions, format string) (types.WeatherResponse, error) {

	// 对 `format` 与 `extensions` 参数进行检查
	if !(format == "xml" || format == "json") {
		return types.WeatherResponse{}, fmt.Errorf("invalid response format: %s", format)
	}

	if !(extensions == "base" || extensions == "all") {
		return types.WeatherResponse{}, fmt.Errorf("Invalid type value(live/forecast): %s", extensions)
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
		return types.WeatherResponse{}, fmt.Errorf("parse url requestUrl failed,err：%s", err)
	}

	// url encode
	u.RawQuery = data.Encode()
	resp, err := http.Get(u.String())
	defer resp.Body.Close()
	if err != nil {
		return types.WeatherResponse{}, fmt.Errorf("get failed, err：%s", err)
	}

	//读取内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return types.WeatherResponse{}, fmt.Errorf("get resp failed, err: %s", err)
	}

	// json 转换为 结构体
	var wr types.WeatherResponse
	err = json.Unmarshal([]byte(body), &wr)
	if err != nil {
		return types.WeatherResponse{}, fmt.Errorf("json Unmarshal, err: %s", err)
	}

	return wr, nil
}
