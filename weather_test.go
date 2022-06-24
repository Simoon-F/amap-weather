package weather

import (
	"github.com/stretchr/testify/assert"
	"github.com/thedevsaddam/gojsonq/v2"
	"testing"
)

func TestGetWeather(t *testing.T) {
	w := NewWeather("ab24eb3025e1bfa830b3790b1f721037")

	req, err := w.GetWeather("中山", "base", "json")

	infocode := gojsonq.New().FromString(req).Find("infocode")

	assert.Equal(t, nil, err)
	assert.Equal(t, "10000", infocode)
}
