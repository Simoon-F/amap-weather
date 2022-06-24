
<h1 align="center">amap-weather</h1>

<p align="center">:rainbow: 基于 高德开放平台 的天气查询扩展包。</p>

## Install

```sh
$ go get -u github.com/Simoon-F/amap-weather
```

## Configure

在使用本扩展之前，你需要去 [高德开放平台](https://console.amap.com/dev/index) 注册账号，然后创建应用，获取应用的 Web APi Key。

## Use

```go
package main

import (
	weather "github.com/Simoon-F/amap-weather"
)

w := weather.NewWeather("key")

```

### GetLiveWeather

```go
resp, _ := w.GetLiveWeather("中山", "json")
```
Example：
```json
{
  "status":"1",
  "count":"2",
  "info":"OK",
  "infocode":"10000",
  "lives":[
    {
      "province":"辽宁",
      "city":"中山区",
      "adcode":"210202",
      "weather":"晴",
      "temperature":"30",
      "winddirection":"西",
      "windpower":"≤3",
      "humidity":"32",
      "reporttime":"2022-06-24 17:32:47"
    },
    {
      "province":"广东",
      "city":"中山市",
      "adcode":"442000",
      "weather":"晴",
      "temperature":"33",
      "winddirection":"西南",
      "windpower":"≤3",
      "humidity":"59",
      "reporttime":"2022-06-24 17:31:06"
    }
  ]
}

```
### GetLiveWeather
```go
resp, err := w.GetForecastsWeather("广州", "json")
```
Example：
```json
{
  "status":"1",
  "count":"1",
  "info":"OK",
  "infocode":"10000",
  "forecasts":[
    {
      "city":"广州市",
      "adcode":"440100",
      "province":"广东",
      "reporttime":"2022-06-24 18:00:26",
      "casts":[
        {
          "date":"2022-06-24",
          "week":"5",
          "dayweather":"多云",
          "nightweather":"多云",
          "daytemp":"35",
          "nighttemp":"25",
          "daywind":"北",
          "nightwind":"北",
          "daypower":"≤3",
          "nightpower":"≤3"
        },
        {
          "date":"2022-06-25",
          "week":"6",
          "dayweather":"多云",
          "nightweather":"多云",
          "daytemp":"35",
          "nighttemp":"25",
          "daywind":"北",
          "nightwind":"北",
          "daypower":"≤3",
          "nightpower":"≤3"
        },
        {
          "date":"2022-06-26",
          "week":"7",
          "dayweather":"多云",
          "nightweather":"多云",
          "daytemp":"35",
          "nighttemp":"26",
          "daywind":"北",
          "nightwind":"北",
          "daypower":"≤3",
          "nightpower":"≤3"
        },
        {
          "date":"2022-06-27",
          "week":"1",
          "dayweather":"雷阵雨",
          "nightweather":"雷阵雨",
          "daytemp":"34",
          "nighttemp":"27",
          "daywind":"北",
          "nightwind":"北",
          "daypower":"≤3",
          "nightpower":"≤3"
        }
      ]
    }
  ]
}
```

## Parameter Description

```
GetLiveWeather(city, format string)
GetForecastsWeather(city, format string)
```

> - `$city` - 城市名/[高德地址位置 adcode](https://lbs.amap.com/api/webservice/download)，比如：“广州” 或者（adcode：440100）；
> - `$format`  - 输出的数据格式，当前只支持 JSON 格式

## Reference

- [高德开放平台天气接口](https://lbs.amap.com/api/webservice/guide/api/weatherinfo/)

## License

MIT