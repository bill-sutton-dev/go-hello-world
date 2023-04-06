package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) {
	var positiveCount float64
	var negativeCount float64
	var zeroCount float64
	arraySize := len(arr)

	for i := 0; i < arraySize; i++ {
		if arr[i] > 0 {
			positiveCount += 1
		} else if arr[i] < 0 {
			negativeCount += 1
		} else {
			zeroCount += 1
		}
	}
	fmt.Printf("%.6f \n", doMath(positiveCount, arraySize))
	fmt.Printf("%.6f \n", doMath(negativeCount, arraySize))
	fmt.Printf("%.6f \n", doMath(zeroCount, arraySize))

}

func doMath(count float64, arraySize int) float64 {
	const prec = 6
	// numCount, _ := new(big.Float).SetPrec(prec).SetString(fmt.Sprint(count))
	// aSize, _ := new(big.Float).SetPrec(prec).SetString(fmt.Sprint(arraySize))
	// output := numCount/aSize
	processedNum := count / float64(arraySize)
	return processedNum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
