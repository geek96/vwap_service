package vwap

import (
	"math/big"
)

// VWAP formulae
// Cumulative ( price * volume) / Cumulative ( volume)
// Cumulative ( 41519.31 * 0.10) / Cumulative (0.10)

type VWAPDataPoint struct {
	Volume *big.Float
	Price *big.Float
}

type VWAPData struct {
	DataPoints []VWAPDataPoint
}

func (v *VWAPData) CalcAvg() *big.Float {
	totalVol := big.NewFloat(0)
	totalPriceVol := big.NewFloat(0)
	// Need a better approach to calculate avg for the sliding window
	// This will increase the complexity linearly as number of inputs increased by O(N)
	// Just a random thought, keep the 1st data element and totalPriceVol and Total Volume
	// Subtract the 1st data element from the total PriceVol and Total Volume
	// Add the last added item price and volume in total Price Vol and Total Vol
	// Update the
	for _, dt :=  range v.DataPoints {
		pv := new(big.Float).Mul(dt.Price, dt.Volume)
		totalPriceVol.Add(totalPriceVol, pv)
		totalVol.Add(totalVol, dt.Volume)
	}
	return new(big.Float).Quo(totalPriceVol, totalVol)
}
