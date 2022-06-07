package timeseries

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func TestFirstObs(t *testing.T) {
	var sr1 TimeSeries
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 52, 548561489, time.UTC), 54.56)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 54, 548561489, time.UTC), 90.2)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 54, 548561489, time.UTC), 47.5)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 55, 548561489, time.UTC), 98.56)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 55, 999999999, time.UTC), NaN())
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 1, 648561489, time.UTC), 999.55) //Last Observation
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 58, 548561489, time.UTC), 1.5)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 58, 548561489, time.UTC), 2.89)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 59, 225688654, time.UTC), 34.156)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 59, 999999999, time.UTC), NaN())
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 0, 548561489, time.UTC), 90005)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 1, 589, time.UTC), 45.98)
	obs, err := FirstObs(&sr1.Data)
	wanted := Observation{
		Chron: time.Date(1983, 1, 15, 10, 15, 52, 548561489, time.UTC),
		Meas:  54.56,
	}
	if obs != wanted || err != nil {
		t.Errorf("Wanted: %v - Obtained: %v\n", wanted, obs)
	}
}
func TestFirst(t *testing.T) {
	var sr1 TimeSeries
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 52, 548561489, time.UTC), 54.56)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 54, 548561489, time.UTC), 90.2)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 54, 548561489, time.UTC), 47.5)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 55, 548561489, time.UTC), 98.56)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 55, 999999999, time.UTC), NaN())
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 1, 648561489, time.UTC), 999.55) //Last Observation
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 58, 548561489, time.UTC), 1.5)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 58, 548561489, time.UTC), 2.89)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 59, 225688654, time.UTC), 34.156)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 59, 999999999, time.UTC), NaN())
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 0, 548561489, time.UTC), 90005)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 1, 589, time.UTC), 45.98)
	obs := First(&sr1.Data)
	wanted := 54.56
	if obs != wanted {
		t.Errorf("Wanted: %v - Obtained: %v\n", wanted, obs)
	}
}
func TestFirstValidObs(t *testing.T) {
	var sr1 TimeSeries
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 52, 548561489, time.UTC), NaN())
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 54, 548561489, time.UTC), 90.2)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 54, 548561489, time.UTC), 47.5)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 55, 548561489, time.UTC), 98.56)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 55, 999999999, time.UTC), NaN())
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 1, 648561489, time.UTC), NaN()) //Last Observation
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 58, 548561489, time.UTC), 1.5)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 58, 548561489, time.UTC), 2.89)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 59, 225688654, time.UTC), 34.156)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 59, 999999999, time.UTC), NaN())
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 0, 548561489, time.UTC), 90005)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 1, 589, time.UTC), 45.98)
	obs, _, err := FirstValidObs(&sr1.Data)
	wanted := Observation{
		Chron: time.Date(1983, 1, 15, 10, 15, 54, 548561489, time.UTC),
		Meas:  90.2,
	}
	fmt.Printf("%v\n", obs)
	sr1.PrettyPrint()
	if obs != wanted || err != nil {
		t.Errorf("Wanted: %v - Obtained: %v\n", wanted, obs)
	}
}
func TestFirstValid(t *testing.T) {
	var sr1 TimeSeries
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 52, 548561489, time.UTC), NaN())
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 54, 548561489, time.UTC), 90.2)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 54, 548561489, time.UTC), 47.5)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 55, 548561489, time.UTC), 98.56)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 55, 999999999, time.UTC), NaN())
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 1, 648561489, time.UTC), NaN()) //Last Observation
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 58, 548561489, time.UTC), 1.5)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 58, 548561489, time.UTC), 2.89)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 59, 225688654, time.UTC), 34.156)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 59, 999999999, time.UTC), NaN())
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 0, 548561489, time.UTC), 90005)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 1, 589, time.UTC), 45.98)
	obs := FirstValid(&sr1.Data)
	wanted := 90.2
	fmt.Printf("%v\n", obs)
	sr1.PrettyPrint()
	if obs != wanted {
		t.Errorf("Wanted: %v - Obtained: %v\n", wanted, obs)
	}
}
func TestLastObs(t *testing.T) {
	var sr1 TimeSeries
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 52, 548561489, time.UTC), 54.56)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 54, 548561489, time.UTC), 90.2)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 54, 548561489, time.UTC), 47.5)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 55, 548561489, time.UTC), 98.56)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 55, 999999999, time.UTC), NaN())
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 1, 648561489, time.UTC), 999.55) //Last Observation
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 58, 548561489, time.UTC), 1.5)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 58, 548561489, time.UTC), 2.89)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 59, 225688654, time.UTC), 34.156)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 59, 999999999, time.UTC), NaN())
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 0, 548561489, time.UTC), 90005)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 1, 589, time.UTC), 45.98)
	obs, err := LastObs(&sr1.Data)
	wanted := Observation{
		Chron: time.Date(1983, 1, 15, 10, 16, 1, 648561489, time.UTC),
		Meas:  999.55,
	}
	if obs != wanted || err != nil {
		t.Errorf("Wanted: %v - Obtained: %v\n", wanted, obs)
	}
}
func TestLast(t *testing.T) {
	var sr1 TimeSeries
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 52, 548561489, time.UTC), 54.56)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 54, 548561489, time.UTC), 90.2)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 54, 548561489, time.UTC), 47.5)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 55, 548561489, time.UTC), 98.56)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 55, 999999999, time.UTC), NaN())
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 1, 648561489, time.UTC), 999.55) //Last Observation
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 58, 548561489, time.UTC), 1.5)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 58, 548561489, time.UTC), 2.89)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 59, 225688654, time.UTC), 34.156)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 15, 59, 999999999, time.UTC), NaN())
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 0, 548561489, time.UTC), 90005)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 1, 589, time.UTC), 45.98)
	obs := Last(&sr1.Data)
	wanted := 999.55
	if obs != wanted {
		t.Errorf("Wanted: %v - Obtained: %v\n", wanted, obs)
	}
}
func TestLastValidObs(t *testing.T) {
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
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 0, 548561489, time.UTC), 90005)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 1, 589, time.UTC), 45.98)
	obs, _, err := LastValidObs(&sr1.Data)
	wanted := Observation{
		Chron: time.Date(1983, 1, 15, 10, 16, 1, 589, time.UTC),
		Meas:  45.98,
	}
	fmt.Printf("%v\n", obs)
	sr1.PrettyPrint()
	if obs != wanted || err != nil {
		t.Errorf("Wanted: %v - Obtained: %v\n", wanted, obs)
	}
}
func TestLastValid(t *testing.T) {
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
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 0, 548561489, time.UTC), 90005)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 1, 589, time.UTC), 45.98)
	obs := LastValid(&sr1.Data)
	wanted := 45.98
	fmt.Printf("%v\n", obs)
	sr1.PrettyPrint()
	if obs != wanted {
		t.Errorf("Wanted: %v - Obtained: %v\n", wanted, obs)
	}
}
func TestMaxObs(t *testing.T) {
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
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 0, 548561489, time.UTC), 90005)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 1, 589, time.UTC), 45.98)
	obs, _, err := MaxObs(&sr1.Data)
	wanted := Observation{
		Chron: time.Date(1983, 1, 15, 10, 16, 0, 548561489, time.UTC),
		Meas:  90005,
	}
	fmt.Printf("%v\n", obs)
	sr1.PrettyPrint()
	if obs != wanted || err != nil {
		t.Errorf("Wanted: %v - Obtained: %v\n", wanted, obs)
	}
}
func TestMax(t *testing.T) {
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
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 0, 548561489, time.UTC), 90005)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 1, 589, time.UTC), 45.98)
	obs := Max(&sr1.Data)
	wanted := 90005.0
	fmt.Printf("%v\n", obs)
	sr1.PrettyPrint()
	if obs != wanted {
		t.Errorf("Wanted: %v - Obtained: %v\n", wanted, obs)
	}
}
func TestMinObs(t *testing.T) {
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
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 0, 548561489, time.UTC), 90005)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 1, 589, time.UTC), 45.98)
	obs, _, err := MinObs(&sr1.Data)
	wanted := Observation{
		Chron: time.Date(1983, 1, 15, 10, 15, 58, 548561489, time.UTC),
		Meas:  1.5,
	}
	fmt.Printf("%v\n", obs)
	sr1.PrettyPrint()
	if obs != wanted || err != nil {
		t.Errorf("Wanted: %v - Obtained: %v\n", wanted, obs)
	}
}
func TestMin(t *testing.T) {
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
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 0, 548561489, time.UTC), 90005)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 1, 589, time.UTC), 45.98)
	obs := Min(&sr1.Data)
	wanted := 1.5
	fmt.Printf("%v\n", obs)
	sr1.PrettyPrint()
	if obs != wanted {
		t.Errorf("Wanted: %v - Obtained: %v\n", wanted, obs)
	}
}
func TestMeanObs(t *testing.T) {
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
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 0, 548561489, time.UTC), 90005)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 1, 589, time.UTC), 45.98)
	avg, pop, err := MeanObs(&sr1.Data)
	wantedavg := 10042.260666666665
	wantedpop := 9
	fmt.Printf("%v\n", avg)
	sr1.PrettyPrint()
	if avg != wantedavg || wantedpop != pop || err != nil {
		t.Errorf("Wanted: %v - Obtained: %v\n", wantedavg, avg)
	}
}
func TestMean(t *testing.T) {
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
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 0, 548561489, time.UTC), 90005)
	sr1.InsertDirect(time.Date(1983, 1, 15, 10, 16, 1, 589, time.UTC), 45.98)
	avg := Mean(&sr1.Data)
	wantedavg := 10042.260666666665
	fmt.Printf("%v\n", avg)
	sr1.PrettyPrint()
	if avg != wantedavg {
		t.Errorf("Wanted: %v - Obtained: %v\n", wantedavg, avg)
	}
}
func TestStdDevObs(t *testing.T) {
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
	std, pop, err := StdDevObs(&sr1.Data)
	wantedstd := 33.69763959092685
	wantedpop := 8
	fmt.Printf("%v\n", std)
	sr1.PrettyPrint()
	if math.Abs(std-wantedstd) > 0.01 || wantedpop != pop-1 || err != nil {
		t.Errorf("Wanted: %v - Obtained: %v - Difference: %v\n", wantedstd, std, wantedstd-std)
	}
}
func TestEstimatedGauss(t *testing.T) {
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
	_, std, err := EstimatedGauss(&sr1.Data)
	wantedstd := 33.69763959092685
	fmt.Printf("%v\n", std)
	sr1.PrettyPrint()
	if math.Abs(std-wantedstd) > 0.01 || err != nil {
		t.Errorf("Wanted: %v - Obtained: %v - Difference: %v\n", wantedstd, std, wantedstd-std)
	}
}
func TestMedianObs(t *testing.T) {
	// TODO
}
