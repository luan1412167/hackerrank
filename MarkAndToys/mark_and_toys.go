package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
)

/*
 * Complete the 'maximumToys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY prices
 *  2. INTEGER k
 */

func maximumToys(prices []int32, k int32) int32 {
    // Write your code here
    sort.Slice(prices, func(i, j int) bool { return prices[i] < prices[j] })
    var spend int32 = 0
    for i:=0; i<len(prices); i++{
        spend += prices[i]
        if spend >= k{
            return int32(i)
        }
    }
    return int32(0)

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    k := int32(kTemp)

    pricesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var prices []int32

    for i := 0; i < int(n); i++ {
        pricesItemTemp, err := strconv.ParseInt(pricesTemp[i], 10, 64)
        checkError(err)
        pricesItem := int32(pricesItemTemp)
        prices = append(prices, pricesItem)
    }

    result := maximumToys(prices, k)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
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
