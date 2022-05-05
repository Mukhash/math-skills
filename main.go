package main

import "os"
import "bufio"
import "fmt"
import "io"
import "log"
import "strconv"
import "errors"

func main() {
	if len(os.Args) != 2 {
		log.Fatal("invalid number of args...\n")
	}

	fileName := os.Args[1]
	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	nums, err := readFile(f)
	if err != nil {
		log.Fatal(err)
	}
	if nums == nil {
		log.Fatal("empty file...")
	}

	fmt.Printf("Average: %d\n", average(nums))
}

func average(nums []int) int {
	var sum int
	for n := range nums {
		sum += n
	}
	fmt.Println(sum)
	fmt.Println(len(nums))
	return sum / len(nums)
}

func readFile(f io.Reader) ([]int, error) {
	rd := bufio.NewReader(f)

	var res []int

	for {
		numRaw, err := rd.ReadBytes('\n')
		var EOF bool
		if err == io.EOF {
			EOF = true
		}

		if !EOF && err != nil {
			return nil, err
		}

		if len(numRaw) <= 1 && EOF { // in case of empty line with EOF
			break
		} else if len(numRaw) <= 1 { // in case of empty line in the middle of file
			return nil, errors.New("invalid file...")
		}

		numStr := string(numRaw[:len(numRaw)-1])
		if EOF && numRaw[len(numRaw)-1] != '\n' { // in case of no \n for the last line
			numStr = string(numRaw)
		}

		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, err
		}

		res = append(res, num)
		if EOF {
			break
		}
	}
	return res, nil
}
