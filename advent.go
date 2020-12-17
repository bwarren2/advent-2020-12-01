package advent

import (
	"bufio"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// NumsFromFile gets a list of numbers from a file
func NumsFromFile(filename string) (nums []int64) {
	file, err := os.Open(filename)
	check(err)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		parsedInt, err := strconv.ParseInt(scanner.Text(), 10, 64)
		check(err)
		nums = append(nums, parsedInt)
	}
	return
}

// GetSumTo returns a pair of numbers that sum to a target given a list of nums and a target, return a pair that
func GetSumTo(nums []int64, target int64) (a, b int64) {
	existenceSet := make(map[int64]bool)
	for _, value := range nums {
		existenceSet[value] = true
	}
	for _, value := range nums {
		want := target - value
		if _, ok := existenceSet[want]; ok {
			return want, value
		}
	}
	return 0, 0 //Because the given problems have solutions, we never get here.
}

// SumTo satisfied advent problem 1
func SumTo(filename string, target int64) int64 {

	nums := NumsFromFile(filename)
	a, b := GetSumTo(nums, target)
	return a * b
}

// TriSumTo finds three number that sum to a value
func TriSumTo(filename string, target int64) int64 {
	nums := NumsFromFile(filename)
	for idx, value := range nums {
		want := target - value
		var sublist []int64
		for innerIdx, innerValue := range nums {
			if innerIdx != idx {
				sublist = append(sublist, innerValue)
			}
		}
		a, b := GetSumTo(sublist, want)
		if a != 0 && b != 0 {

			return value * a * b
		}
	}
	return 0
}
