package main

import (
	"fmt"
	"strconv"
	"strings"
)

const testString = "Suppose we want to add these two binary strings \"111\" and \"1011\" whose numeric values are 7 and 11 whose addition is 18 with binary representation \"10010\". Now we will do the addition of these strings step by step below.variables with the respective values you want to multiply.\n\nAs you can see the length of string \"111\" is less than \"1011\" so we have to make them equal for which we can add \"0\" in a string that is shorter in length due to which the value will also remain the same. The length of \"111\" is shorter than \"1011\" so we will add one zero."

func main() {

	var originalBits string
	for _, letter := range testString {
		originalBits += fmt.Sprintf("%.8b", letter)
	}

	compressedBits := compress(originalBits)

	uncompressed := uncompress(compressedBits, string(originalBits[0]))

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
		}
	}

	fmt.Println("==== received ====")
	fmt.Println(received)
	fmt.Println("==== received ====")

	fmt.Println("==== len(out2) ====")
	fmt.Println(len(originalBits))
	fmt.Println("==== len(out2) ====")

	fmt.Println("==== len(compressedVersion) ====")
	fmt.Println(len(compressedBits))
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

	for _, count := range compressed {
		countAsInt, err := strconv.Atoi(string(count))
		if err != nil {
			panic(err)
		}

		bitList := make([]string, countAsInt)

		for i := range bitList {
			bitList[i] = currBin
		}

		out += strings.Join(bitList, "")

		if currBin == "1" {
			currBin = "0"
		} else {
			currBin = "1"
		}
	}

	return out
}
