package timeseries

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

func (du *Observation) PrettyPrint() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 5, 0, 3, ' ', tabwriter.AlignRight)
	printingstring := "%v|\t%v|\t%v|\n"
	fmt.Fprintf(w, printingstring, "Chron", "Measure")
	fmt.Fprintln(w, "--------|\t------------------------|\t--------------------------------------|")
	fmt.Fprintf(w, printingstring, du.Chron.Round(0), du.Meas)
	fmt.Fprintln(w)
	w.Flush()
}

func (sr *TimeSeries) PrettyPrint(what ...int) {
	fmt.Printf("TimeSeries: %v\n", sr.Name)
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------")
	var j, k int
	switch len(what) {
	case 0:
		{
			j = 0
			k = len(sr.Data)
		}
	case 1:
		{
			j = 0
			k = what[1]
		}
	case 2:
		{
			j = what[0]
			k = what[1]
		}
	default:
		{
			j = 0
			k = len(sr.Data)
		}
	}
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 5, 0, 3, ' ', tabwriter.AlignRight)
	printingstring := "%v|\t%s|\t\t%v\n"
	fmt.Fprintf(w, printingstring, "Index", "Chron", "Measure")
	fmt.Fprintln(w, "------|\t--------------------------------|\t--------------------------------------|")
	for i := j; i < k && i < len(sr.Data); i++ {
		if i == 0 {
			fmt.Fprintf(w, printingstring, i, sr.Data[i].Chron.Format(time.RFC3339Nano), sr.Data[i].Meas)
		} else {
			fmt.Fprintf(w, printingstring, i, sr.Data[i].Chron.Format(time.RFC3339Nano), sr.Data[i].Meas)
		}
	}
	fmt.Fprintln(w)
	fmt.Println("-----------------------------------------------------------------------------------------------------------------------")
	w.Flush()
}

// PrintTsStats prints TsStats struct in a readable way in output terminal
/*
func (ts *TimeSeries) PrintTsStats() {
	fmt.Println("------------------------------------------")
	fmt.Println(ts.Name)
	fmt.Printf("Warning: %v Missing Data\n", ts.UniStats.NbreOfNaN)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 5, 0, 3, ' ', tabwriter.AlignRight)
	fmt.Fprintf(w, "Length| %v|\t\n", ts.UniStats.Len)
	fmt.Fprintln(w, "\tChron|\tMeasure|\tDChron|\tDMeas|\t")
	fmt.Fprintln(w, "-\t-----------------\t------------\t------------\t------------\t")
	fmt.Fprintf(w, "MinObs|\t %v|\t%v|\t%v|\t%v|\t\n", ts.UniStats.First.Round(0), ts.UniStats.Msmin, ts.UniStats.DChmin, ts.UniStats.DMsmin)
	fmt.Fprintf(w, "MaxObs|\t %v|\t%v|\t%v|\t%v|\t\n", ts.UniStats.LastObs.Round(0), ts.UniStats.Msmax, ts.UniStats.DChmax, ts.UniStats.DMsmax)
	fmt.Fprintf(w, "MeanObs|\t %v|\t%v|\t%v|\t%v|\t\n", ts.UniStats.Chmean, ts.UniStats.Msmean, ts.UniStats.DChmean, ts.UniStats.DMsmean)
	fmt.Fprintf(w, "Median|\t %v|\t%v|\t%v|\t%v|\t\n", ts.UniStats.Chmed, ts.UniStats.Msmean, ts.UniStats.DChmed, ts.UniStats.DMsmed)
	fmt.Fprintf(w, "StdDev|\t %v|\t%v|\t%v|\t%v|\t\n", " ", ts.UniStats.Msstd, ts.UniStats.DChstd, ts.UniStats.DMsstd)

	fmt.Fprintln(w)
	w.Flush()

}
func (ts *TimeSeries) PrettyPrintAll() {
	ts.PrettyPrint()
	ts.PrintTsStats()
}


*/
