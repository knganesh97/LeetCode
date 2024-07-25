package problem912

//Counting Sort
func sortArray(nums []int) []int {

	var n = len(nums)
	var count = make(map[int]int, n)
	var minNum, maxNum = nums[0], nums[0]

	for _, num := range nums {
		count[num]++
		minNum = min(num, minNum)
		maxNum = max(num, maxNum)
	}

	var sortedNums []int
	for i := minNum; i <= maxNum; i++ {
		for count[i] > 0 {
			sortedNums = append(sortedNums, i)
			count[i]--
		}
	}

	return sortedNums
}
