package main

import (
	"fmt"
	"strconv"
	"strings"
)

const testString = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis aliquam ante hendrerit gravida viverra. Nam rutrum ornare purus non mollis. Pellentesque at commodo orci, eu varius neque. Sed dapibus sed ipsum vitae volutpat. Nam laoreet accumsan elit id aliquam. Quisque lacus elit, convallis in mollis et, mollis a nisl. Proin non ultrices neque. Sed interdum pharetra diam dignissim ultricies. Suspendisse ac nibh lectus. Nulla facilisi. Curabitur pharetra leo vel viverra cursus. Phasellus malesuada turpis sed nisi ultrices pulvinar.

Praesent placerat id augue et blandit. Aliquam elementum tempor velit, in rhoncus turpis scelerisque nec. Nam lobortis sem eu magna dapibus, non euismod lorem eleifend. Vestibulum ut ante augue. Aenean quis cursus dui. Sed quis volutpat orci, quis lobortis dolor. In sed venenatis odio.

Nunc dictum vulputate erat nec dignissim. In libero dolor, pharetra at euismod nec, interdum sit amet leo. Sed imperdiet nibh a arcu hendrerit laoreet. Integer lacinia elit quis nibh efficitur, a scelerisque tellus sodales. Cras eget mi magna. Phasellus ut arcu in nisi pellentesque vestibulum ut non risus. Vestibulum tortor sem, efficitur et maximus at, posuere nec lorem.

Fusce massa velit, porttitor at sem et, commodo malesuada ante. Etiam a purus faucibus, lacinia ante ac, scelerisque felis. Quisque sit amet risus nisi. Duis id ipsum congue lorem lacinia vestibulum. Nam in enim diam. Vivamus quis rutrum odio. Integer ullamcorper ultricies molestie. Cras ex metus, elementum non purus vel, dictum imperdiet sem. Cras suscipit ullamcorper vulputate. Quisque porttitor nisi felis, a tempor dolor consequat in. Suspendisse non eros eget ex maximus auctor vitae id nulla. Vivamus ut risus arcu. Duis venenatis, purus sit amet eleifend fringilla, risus diam gravida arcu, id tempor lacus ex nec odio. Morbi nulla dui, facilisis at massa sit amet, aliquam porttitor ex. Mauris et pharetra nunc. Mauris et facilisis elit, id luctus urna.

Morbi eu nisi et ante molestie suscipit. Duis vitae libero dapibus, sollicitudin risus ut, condimentum lorem. Fusce at lorem non nisl pharetra ultricies. Phasellus neque orci, hendrerit a urna in, porta imperdiet neque. Pellentesque condimentum eros id tempor fringilla. Cras lectus purus, dictum eu fringilla ac, convallis interdum lectus. Integer quis luctus odio. Donec magna nisi, mollis ac venenatis nec, pharetra et erat. Nulla et ligula a leo semper dignissim eu a leo. Cras varius fringilla tincidunt. Donec ultricies vel eros in gravida. In ipsum nunc, rhoncus sit amet aliquam eget, vulputate eget mi. Quisque eget ex finibus felis vehicula rutrum.`

func main() {

	var originalBits string
	for _, letter := range testString {
		originalBits += fmt.Sprintf("%.8b", letter)
	}

	compressedBits := compress(originalBits)

	uncompressedBits := uncompress(compressedBits, string(originalBits[0]))

	var (
		currBuf  string
		received string
	)

	for _, bit := range uncompressedBits {
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

	fmt.Println("==== received == testString ====")
	fmt.Println(received == testString)
	fmt.Println("==== received == testString ====")

	fmt.Println("==== len(out2) ====")
	fmt.Println(len(originalBits))
	fmt.Println("==== len(out2) ====")

	fmt.Println("==== len(compressedVersion) ====")
	fmt.Println(len(compressedBits))
	fmt.Println("==== len(compressedVersion) ====")

	// check compression %
	fmt.Println("==== len(compressedBits) / len(originalBits) ====")
	fmt.Println(float32(len(compressedBits)) / float32(len(originalBits)))
	fmt.Println("==== len(compressedBits) / len(originalBits) ====")
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
