package accessoryPkg

import (
	"fmt"
	"math"
)

type AccessoryStats struct {
	Min int
	Max int
}

type Accessory struct {
	Name  string
	Stats map[int]AccessoryStats
}

type OptionValues struct {
	하 float64
	중 float64
	상 float64
}

type Config struct {
	PolishingEffect map[string]int
	OptionValues    map[int]OptionValues
}

func NewConfig() *Config {
	return &Config{
		PolishingEffect: map[string]int{
			"공격력 %":        45,
			"공격력 +":        53,
			"무기 공격력 %":     46,
			"무기 공격력 +":     54,
			"적에게 주는 피해 증가": 42,
			"추가 피해":        41,
			"치명타 적중률":      49,
			"치명타 피해":       50,
			"아군 공격력 강화 효과": 51,
			"아군 피해량 강화 효과": 52,
			"낙인력":          44,
		},
		OptionValues: map[int]OptionValues{
			45: {하: 0.4, 중: 0.95, 상: 1.55},
			53: {하: 80, 중: 195, 상: 390},
			46: {하: 0.8, 중: 1.8, 상: 3},
			54: {하: 195, 중: 480, 상: 960},
			42: {하: 0.55, 중: 1.2, 상: 2},
			41: {하: 0.7, 중: 1.6, 상: 2.6},
			49: {하: 0.4, 중: 0.95, 상: 1.55},
			50: {하: 1.1, 중: 2.4, 상: 4},
			51: {하: 1.35, 중: 3, 상: 5},
			52: {하: 2, 중: 4.5, 상: 7.5},
			44: {하: 2.15, 중: 4.8, 상: 8},
		},
	}
}

var accessories = map[string]Accessory{
	"반지": {
		Name: "반지",
		Stats: map[int]AccessoryStats{
			0: {Min: 9156, Max: 11091},
			1: {Min: 9414, Max: 11349},
			2: {Min: 9930, Max: 11865},
			3: {Min: 10962, Max: 12897},
		},
	},
	"귀걸이": {
		Name: "귀걸이",
		Stats: map[int]AccessoryStats{
			0: {Min: 9861, Max: 11944},
			1: {Min: 10139, Max: 12222},
			2: {Min: 10695, Max: 12778},
			3: {Min: 11806, Max: 13889},
		},
	},
	"목걸이": {
		Name: "목걸이",
		Stats: map[int]AccessoryStats{
			0: {Min: 12678, Max: 15357},
			1: {Min: 13035, Max: 15714},
			2: {Min: 13749, Max: 16428},
			3: {Min: 15178, Max: 17857},
		},
	},
}

func (c *Config) GetMinValue(polishingEffectCode int, grade string) int {
	if polishingEffectCode == 42 { // 적에게 주는 피해 증가
		return 621000000
	}

	optionValues, exists := c.OptionValues[polishingEffectCode]
	if !exists {
		return 0
	}

	var value float64
	switch grade {
	case "하":
		value = optionValues.하
	case "중":
		value = optionValues.중
	case "상":
		value = optionValues.상
	default:
		return 0
	}

	return int(value * 100)
}

func (c *Config) GetMaxValue(polishingEffect int) int {
	if polishingEffect == c.PolishingEffect["적에게 주는 피해 증가"] {
		return 621000002
	}

	optionValues, exists := c.OptionValues[polishingEffect]
	if !exists {
		return 0
	}

	return int(math.Max(math.Max(optionValues.하, optionValues.중), optionValues.상) * 100)
}

/*
ex) percent, grade, err := EvaluateAccessory("반지", 2, 10500)
*/
func EvaluateAccessory(name string, level int, value int) (float64, string, error) {
	accessory, exists := accessories[name]
	if !exists {
		return 0, "", fmt.Errorf("악세서리 '%s'를 찾을 수 없습니다", name)
	}

	stats, exists := accessory.Stats[level]
	if !exists {
		return 0, "", fmt.Errorf("악세서리 '%s'의 연마단계 %d를 찾을 수 없습니다", name, level)
	}

	if value < stats.Min || value > stats.Max {
		return 0, "", fmt.Errorf("값 %d가 유효 범위를 벗어났습니다 (최소: %d, 최대: %d)", value, stats.Min, stats.Max)
	}

	percent := float64(value-stats.Min) / float64(stats.Max-stats.Min) * 100

	var grade string
	switch {
	case percent >= 90:
		grade = "상"
	case percent >= 70:
		grade = "중"
	default:
		grade = "하"
	}

	return percent, grade, nil
}
