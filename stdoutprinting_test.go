package timeseries

import (
	"testing"
	"time"
)

func TestObservation_PrettyPrint(t *testing.T) {
	var sr1 TimeSeries
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 52, 548561489, time.UTC), 54.56)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 54, 548561489, time.UTC), 90.2)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 54, 548561489, time.UTC), 47.5)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 55, 548561489, time.UTC), 98.56)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 55, 999999999, time.UTC), NaN())
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 1, 648561489, time.UTC), NaN()) //Last Observation
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 58, 548561489, time.UTC), 1.5)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 58, 548561489, time.UTC), 2.89)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 59, 225688654, time.UTC), 34.156)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 59, 999999999, time.UTC), NaN())
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 0, 548561489, time.UTC), 65.9)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 1, 589, time.UTC), 45.98)
	sr1.Data[0].PrettyPrint()
}
func TestTimeSeries_PrettyPrint(t *testing.T) {
	var sr1 TimeSeries
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 52, 548561489, time.UTC), 54.56)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 54, 548561489, time.UTC), 90.2)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 54, 548561489, time.UTC), 47.5)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 55, 548561489, time.UTC), 98.56)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 55, 999999999, time.UTC), NaN())
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 1, 648561489, time.UTC), NaN()) //Last Observation
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 58, 548561489, time.UTC), 1.5)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 58, 548561489, time.UTC), 2.89)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 59, 225688654, time.UTC), 34.156)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 59, 999999999, time.UTC), NaN())
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 0, 548561489, time.UTC), 65.9)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 1, 589, time.UTC), 45.98)
	sr1.Data.SortChronAsc()
	sr1.PrettyPrint()
}
