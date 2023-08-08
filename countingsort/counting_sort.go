package main
import "fmt"

func countingSort(orgArr [7]int32) []int32 {
	maxOfOrgArr := orgArr[0]
	for i := 0; i < len(orgArr); i++ {
		if orgArr[i] > maxOfOrgArr {
			maxOfOrgArr = orgArr[i]
		}
	}

	idxArr := make([]int32, maxOfOrgArr+1)

	for i := 0; i < len(orgArr); i++ {
		idxArr[orgArr[i]]++
	}

	for i := 1; i < len(idxArr); i++ {
		idxArr[i] += idxArr[i-1]
	}

	//var sortedArr [len(orgArr)]int32
	sortedArr := make([]int32, len(orgArr))
	for i := 0; i < len(orgArr); i++ {
		sortedArr[idxArr[orgArr[i]]-1] = orgArr[i]
		idxArr[orgArr[i]]--
	}

	return sortedArr

}

func main() {
	orgArr := [7]int32{4, 2, 2, 8, 3, 3, 1}
	var sortedArr []int32
	sortedArr = countingSort(orgArr)
	fmt.Print("sorted Arr", sortedArr)
}
