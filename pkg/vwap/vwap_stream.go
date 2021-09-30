package vwap

import (
	"github.com/geek96/vwap_service/pkg/coinbase"
	"log"
	"sync"
)

type VolWeightedAvgPrice struct {
	DataPointsLen int
	sync.Map
}

func NewVolWeightAvgPrice(datapointLen int) *VolWeightedAvgPrice {
	return &VolWeightedAvgPrice{
		DataPointsLen: datapointLen,
	}
}

func (vw *VolWeightedAvgPrice) Process(ch chan coinbase.MatchResponse, done chan bool) {
	for {
		select {
		case <- done:
			return
		case resp := <-ch:
			dt := VWAPDataPoint{
				Volume: resp.Volume,
				Price:  resp.Price,
			}
			prodId := resp.ProductId
			p, ok := vw.Load(prodId)
			// If the product is not present add in the map with datapoint
			if !ok {
				vwData := VWAPData{DataPoints: []VWAPDataPoint{dt}}
				vw.Store(prodId, vwData)
				continue
			}
			product := p.(VWAPData)
			dataPoints := product.DataPoints
			// If the datapoints
			if len(dataPoints) == vw.DataPointsLen {
				dataPoints = dataPoints[1:]
			}
			product.DataPoints = append(dataPoints, dt)
			vw.Store(prodId, product)
			avg := product.CalcAvg()
			log.Printf("Product [%s], VWAP [%v]", prodId, avg)
		}
	}
}
