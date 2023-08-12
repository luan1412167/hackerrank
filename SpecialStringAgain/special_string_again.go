package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the substrCount function below.
func substrCount(n int32, s string) int64 {
    var count int32 = n
    for i:=int32(0); i<n; i++{
        // string with 1 char
        var j = int32(i+1)
        for(j<n && s[i]==s[j]){
            // var lenSubs = j-i
            count += 1 
            j++
        }
        
        // string with middle one
        
        var k = int32(1)
        // if(i-k>=0 && i+k<n){
        //     fmt.Println(s, i+k, i+1, i-k)
        for(i-k>=0 && i+k<n && s[i]!=s[i+k] && s[i-k]==s[i+k] && s[i+1]==s[i+k]){
        count++
        k++
        }
        // }

    }
    return int64(count)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    s := readLine(reader)

    result := substrCount(n, s)

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
