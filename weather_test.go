package weather

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLiveWeather(t *testing.T) {
	w := NewWeather("ab24eb3025e1bfa830b3790b1f721037")

	resp, err := w.GetLiveWeather("中山", "json")

	assert.Equal(t, nil, err)
	assert.Equal(t, "10000", resp.Infocode)
	assert.Equal(t, "中山市", resp.Lives[1].City)
	assert.Equal(t, "广东", resp.Lives[1].Province)
}

func TestGetForecastsWeather(t *testing.T) {
	w := NewWeather("ab24eb3025e1bfa830b3790b1f721037")

	resp, err := w.GetForecastsWeather("广州", "json")

	assert.Equal(t, nil, err)
	assert.Equal(t, "10000", resp.Infocode)
	assert.Equal(t, "广州市", resp.Forecasts[0].City)
	assert.Equal(t, "广东", resp.Forecasts[0].Province)
	assert.Equal(t, "440100", resp.Forecasts[0].Adcode)
}
