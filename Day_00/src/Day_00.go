package main

import (
	"flag"
	"fmt"
	"math"
	"sort"
	"strconv"
)

var f_mean, f_median, f_mode, f_sd *bool

func init() {
	f_mean = flag.Bool("mean", false, "Finde mean")
	f_median = flag.Bool("median", false, "Finde median")
	f_mode = flag.Bool("mode", false, "Finde mode")
	f_sd = flag.Bool("sd", false, "Finde SD")
	flag.Parse()
}

type Stat struct {
	numbers          []int
	frequency        map[int]int
	mean, median, SD float32
	mode             int
}

func main() {

	var s Stat
	s.InputNumbers()

	if len(s.numbers) != 0 {
		fmt.Println("____________________")
		fmt.Println("NUMS:\t", s.numbers)
		fmt.Println("____________________")

		if !(*f_mean || *f_median || *f_mode || *f_sd) {
			s.CalcMean()
			fmt.Println("MEAN:\t", s.mean)
			s.CalcMedian()
			s.CalcMode()
			s.CalcSD()
		} else {
			if *f_mean {
				s.CalcMean()
				fmt.Println("MEAN:\t", s.mean)
			}
			if *f_median {
				s.CalcMedian()
			}
			if *f_mode {
				s.CalcMode()
			}
			if *f_sd {
				s.CalcSD()
			}
		}
	} else {
		fmt.Println("\n¯\\_(ツ)_/¯")
	}

	fmt.Println()
}

func (s *Stat) InputNumbers() {
	s.numbers = make([]int, 0, 10)
	s.frequency = make(map[int]int, 10)
	flag := true
	var inputString string
	for flag {
		fmt.Scanln(&inputString)
		number, err := strconv.Atoi(inputString)
		if err != nil {
			flag = false
		} else {
			s.numbers = append(s.numbers, number)
			_, numberExists := s.frequency[number]
			if numberExists {
				s.frequency[number] += 1
			} else {
				s.frequency[number] = 1
			}
		}
	}
}

func (s *Stat) CalcMean() {
	s.mean = 0
	for _, value := range s.numbers {
		s.mean += float32(value)
	}
	s.mean /= float32(len(s.numbers))
}

func (s *Stat) CalcMedian() {
	sort.Ints(s.numbers)
	if len(s.numbers)%2 != 0 {
		s.median = float32(s.numbers[len(s.numbers)/2])
	} else {
		s.median = (float32(s.numbers[len(s.numbers)/2]) + float32(s.numbers[len(s.numbers)/2-1])) / 2
	}
	fmt.Println("MEDIAN:\t", s.median)
}

func (s *Stat) CalcMode() {
	var quantity int
	for key, value := range s.frequency {
		if value > quantity {
			quantity = value
			s.mode = key
		}
	}
	fmt.Println("MODE:\t", s.mode, "| quantity:", quantity)
}

func (s *Stat) CalcSD() {
	s.SD = 0
	s.CalcMean()
	for _, value := range s.numbers {
		s.SD += float32(math.Pow(float64(value)-float64(s.mean), 2))
	}
	s.SD = float32(math.Sqrt(float64(s.SD) / float64(len(s.numbers))))
	fmt.Println("SD:\t", s.SD)
}

/*
Check that project only requires running `go build` to produce an executable
Check that submission includes *go.mod* and *go.sum* (in case external dependencies are used)
Check that application works when being passed a sequence of integer numbers, separated by newlines
Check that results are correct both when input is sorted and when it isn't
Check that cases like an empty string with newline, value out of [-100000, 100000] bounds,
characters don't crash the program
Check that application can be run in such way it prints only a specified subset of metrics
Check that mean is calculated correctly for both odd and even number of inputs
Check that the output for mean is rounded to two decimal points
Check that median is calculated correctly for both odd and even number of inputs
Check that the median output is rounded to two decimal points
Check that mode is always equal to the most frequent number in the input
Check that if there are multiple most frequent numbers, then mode is equal to the smallest one among those
Check that SD calculation works even if mean calculation is disabled
Check that SD is calculated correctly, being equal to either regular SD or a population SD
Check that the SD output is rounded to two decimal points
*/
