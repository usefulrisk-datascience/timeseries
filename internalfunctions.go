package timeseries

import (
	"errors"
	"math"
)

// FirstObs returns the first Observation of a DataSeries, whether it is NaN() or not. See FirstValidObs
func FirstObs(data *DataSeries) (lst Observation, err error) {
	data.SortChronAsc()
	if len(*data) > 0 {
		lst = (*data)[0]
	} else {
		//err = errors.New("DataSeries is of Length 0")
	}
	return lst, err
}

// First uses FirstObs and returns a float64
func First(data *DataSeries) float64 {
	last, _ := FirstObs(data)
	return last.Meas
}

// FirstValidObs returns the first Valid Observation (IsNaN()==false) of a DataSeries. See FirstObs
func FirstValidObs(data *DataSeries) (lst Observation, idx int, err error) {
	data.SortChronAsc()
	if len(*data) > 0 {
		for i := 0; i < len((*data)); i++ {
			if IsNaN((*data)[i].Meas) == false {
				lst = (*data)[i]
				idx = i
				break
			} else if i == len((*data))-1 && IsNaN((*data)[0].Meas) == true {
				lst.Chron = (*data)[len((*data))-1].Chron
				lst.Meas = NaN()
			}
		}
	} else {
		err = errors.New("No Valid Observation in DataSeries")
	}
	return lst, idx, err
}

// FirstValid uses FirstValidObs and returns a float64
func FirstValid(data *DataSeries) float64 {
	last, _, _ := FirstValidObs(data)
	return last.Meas
}

// LastObs returns the first Observation of a DataSeries, whether it is NaN() or not. See FirstValidObs
func LastObs(data *DataSeries) (lst Observation, err error) {
	data.SortChronAsc()
	if len(*data) > 0 {
		lst = (*data)[len(*data)-1]
	} else {
		//err = errors.New("DataSeries is of Length 0")
	}
	return lst, err
}

// Last uses LastObs and returns a float64
func Last(data *DataSeries) float64 {
	last, _ := LastObs(data)
	return last.Meas
}

// LastValidObs returns the last Valid Observation (IsNaN()==false) of a DataSeries. See LastObs
func LastValidObs(data *DataSeries) (lst Observation, idx int, err error) {
	data.SortChronAsc()
	if len(*data) > 0 {
		for i := len(*data) - 1; i >= 0; i-- {
			if IsNaN((*data)[i].Meas) == false {
				lst = (*data)[i]
				idx = i
				break
			} else if i == 0 && IsNaN((*data)[0].Meas) == true {
				lst.Chron = (*data)[0].Chron
				lst.Meas = NaN()
			}
		}
	} else {
		err = errors.New("No Valid Observation in DataSeries")
	}
	return lst, idx, err
}

// LastValid uses FirstValidObs and returns a float64
func LastValid(data *DataSeries) float64 {
	last, _, _ := LastValidObs(data)
	return last.Meas
}

// MaxObs returns the Observation showing the Maximum  value of Measures in a DataSeries. If no Valid Observation is available, returns NaN().
// MaxObs returns also the place of that Observation in DataSeries for possible further computations.
func MaxObs(data *DataSeries) (max Observation, idx int, err error) {
	if len(*data) > 0 {
		var flag bool
		for i := 0; i < len(*data); i++ {
			if IsNaN((*data)[i].Meas) == false {
				max = (*data)[i]
				flag = true
				break
			}
		}
		if flag == false {
			max = BlankObservation()
			//err = errors.New("No Valid Maximum in DataSeries")
		} else {
			for i := 0; i < len(*data); i++ {
				if math.IsNaN((*data)[i].Meas) == false {
					if (*data)[i].Meas > max.Meas {
						max = (*data)[i]
					}
				}
			}
		}
	} else {
		err = errors.New("DataSeries is of Length 0")
	}
	return
}

// Max returns the float64 value showing the Maximum value of Measures in a DataSeries. If no Valid Observation is available, returns NaN().
// See MaxObs
func Max(data *DataSeries) (max float64) {
	obs, _, _ := MaxObs(data)
	return obs.Meas
}

// MinObs returns the Observation showing the Minimum  value of Measures in a DataSeries. If no Valid Observation is available, returns NaN().
// MinObs returns also the place of that Observation in DataSeries for possible further computations.
func MinObs(data *DataSeries) (min Observation, idx int, err error) {
	if len((*data)) > 0 {
		var flag bool
		for i := 0; i < len((*data)); i++ {
			if IsNaN((*data)[i].Meas) == false {
				min = (*data)[i]
				flag = true
				break
			}
		}
		if flag == false {
			min = BlankObservation()
			err = errors.New("No Valid Minimum in DataSeries")
		} else {
			for i := 0; i < len((*data)); i++ {
				if math.IsNaN((*data)[i].Meas) == false {
					if (*data)[i].Meas < min.Meas {
						min = (*data)[i]
					}
				}
			}
		}
	} else {
		err = errors.New("DataSeries is of Length 0")
	}
	return
}

// Min returns the float64 value showing the Minimum value of Measures in a DataSeries. If no Valid Observation is available, returns NaN().
// See MinObs
func Min(data *DataSeries) (min float64) {
	obs, _, _ := MinObs(data)
	return obs.Meas
}

// MeanObs returns the Estimated Mean value of Measures in a DataSeries. If no Valid Observation is available, returns NaN().
// MeanObs also returns the amount of Valid Observations in DataSeries for possible further computations.
func MeanObs(data *DataSeries) (avg float64, pop int, err error) {
	if len(*data) > 0 {
		for _, v := range *data {
			if IsNaN(v.Meas) == false {
				avg += v.Meas
				pop += 1
			}
		}
		avg = avg / float64(pop)
	} else {
		avg = NaN()
		pop = 0
		err = errors.New("DataSeries is of Length 0")
	}
	return
}

// Mean returns the float64 value showing the Estimated Mean value of Measures in a DataSeries. If no Valid Observation is available, returns NaN().
// See MeanObs
func Mean(data *DataSeries) (mean float64) {
	mean, _, _ = MeanObs(data)
	return mean

}

// StdDevObs returns the unbiased Estimated Standard Deviation value of Measures in a DataSeries. If no Valid Observation is available, returns NaN().
// StdDevObs also returns the amount of Valid Observations in DataSeries for possible further computations.
func StdDevObs(data *DataSeries) (std float64, pop int, err error) {
	var zxc float64
	if len(*data) > 1 {
		avg, _, err := MeanObs(data)
		avgsq:=avg*avg
		if err != nil {
			err = errors.New("Non Valid MeanObs. Standard Deviation Computation Impossible.")
		}
		var nn int
		for _, v := range *data {
			if IsNaN(v.Meas) == false {
				zxc += (v.Meas*v.Meas - avgsq)
				nn += 1
			}
		}
		std = math.Sqrt(zxc / (float64(nn) - 1))
		pop=nn
	} else {
		std = NaN()
		pop = 0
		err = errors.New("DataSeries is of Length 0")
	}
	return
}

// StdDev returns the float64 value showing the unbiased Estimated Standard Deviation value of Measures in a DataSeries. If no Valid Observation is available, returns NaN().
// See StdDevObs
func StdDev(data *DataSeries) (std float64) {
	std, _, _ = StdDevObs(data)
	return std
}

// EstimatedGauss returns the Estimated Gauss function underlying a sample. It is simply a container for Mean and StdDev.
func EstimatedGauss(data *DataSeries) (mean float64, std float64, err error) {
	if len(*data) > 1 {
		mean, _, err = MeanObs(data)
		if err != nil {
			err = errors.New("No MeanObs to DataSeries")
		}
		std, _, err = StdDevObs(data)
		if err != nil {
			err = errors.New("No Standard Error to DataSeries")
		}
	} else if len(*data) == 1 {
		_, _, err := MeanObs(data)
		if err != nil {
			err = errors.New("No MeanObs to DataSeries")
		}
		err = errors.New("DataSeries of length 1")
	} else {
		err = errors.New("DataSeries is of Length 0")
	}
	return
}

// MedianObs returns the Observation closest to the Median Observation (percentile 50) in a DataSeries.
func MedianObs(data *DataSeries) (obs Observation, err error) {
	// TODO
	return obs, err
}
