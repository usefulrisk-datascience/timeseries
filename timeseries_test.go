package timeseries

import (
	"fmt"
	"testing"
	"time"
)

func TestSeries_Insert(t *testing.T) {
	var sr TimeSeries
	var obs Observation
	obs.Chron = time.Now()
	obs.Meas = NaN()
	sr.Insert(obs)
	fmt.Println(sr)
	sr.Insert(NewObservation(time.Now(), NaN()))
	fmt.Println(sr)
}
func TestNaN(t *testing.T) {
	// TODO: complete non math implementation
	tt := NaN()
	fmt.Println(tt)
}
func TestIsNaN(t *testing.T) {
	nn := NaN()
	if IsNaN(nn) == true {
		fmt.Println("NaN test ok")
	}
}
func TestSeries_InsertDirect(t *testing.T) {
	var sr TimeSeries
	sr.Name = "test"
	sr.InsertDirect(time.Now(), 596.25)
	fmt.Println(sr)
}
func TestBlankObservation(t *testing.T) {
	fmt.Println(BlankObservation())
}
func TestNewObservation(t *testing.T) {
	obs := NewObservation(time.Now(), 458.265)
	fmt.Println(obs)
}
func TestDataSeries_ComputeSimpleStats(t *testing.T) {
	// TODO
}
