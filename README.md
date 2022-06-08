# timeseries #
Go Package timeseries provides functionality for creating, manipulating and processing computations on Time Series.

_Copyright 2022 Frederic Flament. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file._

 A Time Series is a series of time-ordered data points. Commonly a time series is a sequence taken at successive equally spaced points in time (https:en.wikipedia.org/wiki/Time_series).
 A point is called Observation. An Observation is a two-dimensional point. The first dimension, called "Chron",
 is of type time.Time (under the hood int64). The second dimension, called "Meas" is of type float64:

```` 
type Observation struct{

 	Chron time.time
 	Meas float64
  }
````
 A TimeSeries struct has a "Name" field and a "slice of Observations" field, called DataSeries.

````
  type DataSeries []Observation

  type TimeSeries struct{
 	Name string
 	Data DataSeries
  }
````
 Most of computations are processed on a pointer to DataSeries.
 The Chron dimension makes extensive use of the standard "time" package, which hides an int64 type behind human readable dates.

 Conversion of human readable dates and durations into integers are neither straightforward nor trivial; a thorough examination of the time package documentation is advised.

 ## Regularization ##

 It is frequent to create or change the human readable periodicity of a Time Series.

 First example: converting daily rain precipitation into annual precipitation.

 Second example: regularize random arrival of IoT sensors into regularized time series.

 This package provides Method to regularize a Time Series into regularly distanced Second ("sec"), Minute ("min"), Hour ("hou"), Day ("day"), Month ("mon"), Year ("yea") time periods,
 as well as multiple methods for selecting Observations within periods (Minimum, Maximum, Last Valid Observation, Last Observation, First Observation, First Valid Observation, Mean, Standard Deviation).
 Using function as parameter, any other function can be implemented by user.

 ## Missing Data

A distinctive feature of this package is the use of a "Missing Data" value for measure.
 This feature is a necessity because "Missing Data" appears as soon as simple first difference computation is performed,
 and the regularization process frequently needs a Missing Data feature if there is no Measure for a specific period. Furthermore, it is mandatory to be able
 to differentiate between a valid measure of 0 (freezing water in C, or "no-earthquake" information for instance) and a missing value.

 ~~We use a drop-in implementation of math.NaN() (so that import of math package is not required - math.Nan() works as well)~~ (_TODO_). Using NaN() avoids costly alternative
 solutions in terms of memory management (like adding variable like bool "IsValid" or using extensively pointers to float64 measures).
 The IEEE 754 floating-point standard implements a "Not a Number" value in float64 type. Using NaN() specification has several advantages:

 - Although implemented initially in order to manage results of some operations (like square root of negative number or division by zero), there is no counter indication to use "Not a Number" as "Missing Data"

 - Most interestingly in Go language, a NaN() belongs to the float64 type and is accepted by compiler

- It is implemented in all most known languages implementing IEEE 754

Using NaN() as drop-in for Missing Data requires nevertheless some work:

 - math package and statistical package do not accept occurence of NaN() in computation; an Error is produced. We want to be able to perform a calculation on a Time Series in which there are some Missing Data without blocking. We implement our own version of main statistical non-blocking calculations, like Mean, Standard Deviation, Minimum, Maximum and so on. We use the possibility of multiple returns in Go language to return the number of valid measures used in computation.

 - Implementing NaN() as non-blocking requires to take care of propagation rule; that is any calculus chain should be consistent with the same implementation.

 - JSON.Marshall() does not accept NaN(). Presence of NaN() is fatal to the process. Although some language accept presence of null value within a floating array (Javascript to convert to a Javascript Object for example), the JSON specification does not. Specific routine to marshal to compatible "Javascript Object" can be implemented.

 ## A note on alternatives: sql.NullFloat

 When dealing with SQL-type ORM, one meets null value. Usual sql packages implement a NullFloat type:
````
  type NullFloat64 struct {
    Float64 float64
    Valid   bool  Valid is true if Float64 is not NULL
  }
````

 We do not favor this implementation because:

 - Having a two-field struct requires most complex programs and do not avoid dedicated stat computations

 - NullFloat64 are comparable even if non Valid. This is potentially errors prone.