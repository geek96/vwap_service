package vwap

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestSWAPDataCalculate(t *testing.T) {
	vwapData := VWAPData{
		DataPoints: []VWAPDataPoint{
			{
				Price:  big.NewFloat(41519.31),
				Volume: big.NewFloat(0.10),
			},
			{
				Price:  big.NewFloat(41525.54),
				Volume: big.NewFloat(0.4),
			},
		},
	}
	got := vwapData.CalcAvg()
	assert.NotNil(t, got)
	assert.Equal(t, "41524.294", got.String())
}

func TestSWAPDataCalculateWithZeroVol(t *testing.T) {
	vwapData := VWAPData{
		DataPoints: []VWAPDataPoint{
			{
				Price:  big.NewFloat(41519.01),
				Volume: big.NewFloat(0.10),
			},
			{
				Price:  big.NewFloat(41525.54),
				Volume: big.NewFloat(0),
			},
		},
	}
	got := vwapData.CalcAvg()
	assert.NotNil(t, got)
	assert.Equal(t, "41519.01", got.String())
}

func TestVWAPData_CalcAvgShouldPanic(t *testing.T) {
	defer func(t *testing.T) {
		r := recover();
		assert.NotNil(t, r)
	}(t)
	vwapData := VWAPData{
		DataPoints: []VWAPDataPoint{
			{
				Price:  big.NewFloat(41525.54),
				Volume: big.NewFloat(0),
			},
		},
	}
	vwapData.CalcAvg()
}

func BenchmarkVWAPData_CalcAvg(b *testing.B) {
	vwapData := VWAPData{
		DataPoints: []VWAPDataPoint{
			{
				Price:  big.NewFloat(41519.01),
				Volume: big.NewFloat(0.10),
			},
			{
				Price:  big.NewFloat(41525.54),
				Volume: big.NewFloat(0.20),
			},
		},
	}
	vwapData.CalcAvg()

}
