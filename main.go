package main

import (
	"fmt"
	"strconv"
	"strings"
)

const testString = ""

func main() {

	var out2 string
	for _, letter := range testString {
		out2 += fmt.Sprintf("%.8b", letter)
	}

	compressedVersion := compress(out2)
	uncompressed := uncompress(compressedVersion, string(out2[0]))

	var (
		currBuf  string
		received string
	)

	for _, bit := range uncompressed {
		currBuf += string(bit)
		if len(currBuf) == 8 {
			parsed, err := strconv.ParseInt(currBuf, 2, 64)
			if err != nil {
				panic(err)
			}

			currBuf = ""

			received += string(int32(parsed))
			continue
		}
	}

	fmt.Println("==== received ====")
	fmt.Println(received)
	fmt.Println("==== received ====")

	fmt.Println("==== len(out2) ====")
	fmt.Println(len(out2))
	fmt.Println("==== len(out2) ====")

	fmt.Println("==== len(compressedVersion) ====")
	fmt.Println(len(compressedVersion))
	fmt.Println("==== len(compressedVersion) ====")
}

// compress takes a binary representation and returns a shorter version
func compress(binaryStr string) string {
	var (
		out       string
		currBin   int32 = 0
		currCount       = 0
	)

	for i, bin := range binaryStr {
		if bin != currBin {
			if i != 0 {
				out += strconv.Itoa(currCount)
			}
			currBin = bin
			currCount = 1
			continue
		}

		currCount++
	}

	// get last count
	out += strconv.Itoa(currCount)

	return out
}

func uncompress(compressed string, first string) string {
	if len(compressed) < 1 {
		panic("need length")
	}

	var out string

	// first letter is our start
	currBin := first

	for _, bin := range compressed {
		asInt, err := strconv.Atoi(string(bin))
		if err != nil {
			panic(err)
		}

		sl := make([]string, asInt)

		for i := range sl {
			sl[i] = currBin
		}

		out += strings.Join(sl, "")

		if currBin == "1" {
			currBin = "0"
		} else {
			currBin = "1"
		}
	}

	return out
}