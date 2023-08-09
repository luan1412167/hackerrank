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
 * Complete the 'activityNotifications' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY expenditure
 *  2. INTEGER d
 */
 
 func findMedian(countingArr []int32, old int32, new int32, d int32) float32{
     idxCountingArr := make([]int32, 201)
    idxCountingArr[0] = countingArr[0]

    //  for i:=1; i<201; i++{
        //  idxCountingArr[i] = countingArr[i]+idxCountingArr[i-1]
    //  }
     var idx, idx1, idx2, a, b int32
     var ret float32
    //  fmt.Println("here")
     if d%2!=0{
         idx = (d-1)/2
        //  fmt.Println(idx)
         for i:=1; i<201; i++{
             idxCountingArr[i] = countingArr[i]+idxCountingArr[i-1]
            //  fmt.Println(idx, idxCountingArr[i], i)
             if idx < idxCountingArr[i]{
                // fmt.Println(countingArr)
                ret = float32(i)
                break
             }
         }
     }else{
         idx1 = (d-1)/2
         idx2 = d/2
         for i:=int32(1); i<int32(201); i++{
             idxCountingArr[i] = countingArr[i]+idxCountingArr[i-1]
             if idx2 < idxCountingArr[i]{
                 a = i
                //  idxCountingArr[i]--
                 break
             }
         }
         for i:=int32(1); i<int32(201); i++{
             idxCountingArr[i] = countingArr[i]+idxCountingArr[i-1]
             if idx1 < idxCountingArr[i]{
                 b = i
                 ret = float32(a+b)/2
                 break
             }
         }
     }
    countingArr[old]--
    countingArr[new]++
     return ret
 }
 
 

func activityNotifications(expenditure []int32, d int32) int32 {
    // Write your code here
    countingArr := make([]int32, 201)
    for i:=int32(0); i<d; i++{
        countingArr[expenditure[i]]++
    }
    
    var med float32
    var ret int32 = 0
    for i:=d; i<int32(len(expenditure)); i++{   
        // fmt.Println(expenditure[i-d], expenditure[i], d)
        // fmt.Println(countingArr)
        med = findMedian(countingArr, expenditure[i-d], expenditure[i], d)
        fmt.Println(med)
        if float32(expenditure[i]) >= 2*med{
            ret++
        }
    }
    return ret
       

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

    dTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    d := int32(dTemp)

    expenditureTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var expenditure []int32

    for i := 0; i < int(n); i++ {
        expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
        checkError(err)
        expenditureItem := int32(expenditureItemTemp)
        expenditure = append(expenditure, expenditureItem)
    }

    result := activityNotifications(expenditure, d)

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
