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
 * Complete the 'lilysHomework' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func min(x, y int32) int32 {
    if x < y {
        return x
    }
    return y
}
func lilysHomework(arr []int32) int32 {
    // Write your code here
    arr_org := make([]int32, len(arr))
    copy(arr_org, arr)
    // fmt.Println(arr_org)
    sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
    arr_rev := make([]int32, len(arr))
    copy(arr_rev, arr)
    sort.Slice(arr_rev, func(i, j int) bool {
      return arr_rev[j] < arr_rev[i]
   })
    fmt.Println("arr_org, arr, arr_rev", arr_org, arr, arr_rev)
    mapArr := make(map[int32]int32)
    for i, mapAr := range arr_org{
        mapArr[mapAr] = int32(i)
    }
    mapArrDec := make(map[int32]int32)
    for k, v := range mapArr{
         mapArrDec[k] = v       
    }
    fmt.Println("mapArr", mapArr, mapArrDec)
    
    arr_org_inc := make([]int32, len(arr_org))
    copy(arr_org_inc, arr_org)
    arr_org_dec := make([]int32, len(arr_org))
    copy(arr_org_dec, arr_org)
    var ret_inc int32 = 0
    var ret_dec int32 = 0
    for i:=0; i<len(arr_org_inc); i++ {
        if arr[i] != arr_org_inc[i]{
            ret_inc++
            need_val := arr[i]
            swap_val := arr_org_inc[i]
            arr_org_inc[i] = need_val
            arr_org_inc[mapArr[need_val]] = swap_val
            idx_need_val := mapArr[need_val]
            mapArr[need_val] = int32(i)
            mapArr[swap_val] = idx_need_val   
        }
    }
    
    for j:=0; j<len(arr_org_dec); j++ {
        if arr_rev[j] != arr_org_dec[j]{
            ret_dec++
            need_val := arr_rev[j]
            swap_val := arr_org_dec[j]
            arr_org_dec[j] = need_val
            arr_org_dec[mapArrDec[need_val]] = swap_val
            idx_need_val := mapArrDec[need_val]
            mapArrDec[need_val] = int32(j)
            mapArrDec[swap_val] = idx_need_val   
        }
    } 
    
    return min(ret_inc, ret_dec)

}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

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

    result := lilysHomework(arr)

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
