package timeseries

import (
	"sort"
)

// SortChronAsc sort a TimeSeries in chronological ascending order in-place
func (ds DataSeries) SortChronAsc() {
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].Chron.Before(ds[j].Chron)
	})
}

// SortChronDesc sort a TimeSeries in chronological descending order in-place
func (ds DataSeries) SortChronDesc() {
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].Chron.After(ds[j].Chron)
	})
}

// Sort a TimeSeries in ascending order of measure in-place
func (ds DataSeries) SortMeasAsc() {
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].Meas < ds[j].Meas
	})
}

// Sort a TimeSeries in descending order of measure in-place
func (ds DataSeries) SortMeasDesc() {
	sort.Slice(ds, func(i, j int) bool {
		return ds[i].Meas > ds[j].Meas
	})
}

// Reset a TimeSeries to a zero length TimeSeries but keep metadata information at the TimeSeries level
func (ds DataSeries) Reset() {
	ds = ds[0:0]
}

// Deep copy of a time series
func (sr *TimeSeries) Copy() TimeSeries {
	var copyofts TimeSeries
	copyofts.Name = sr.Name

	for _, value := range sr.Data {
		copyofts.Data = append(copyofts.Data, value)
	}
	return copyofts
}
