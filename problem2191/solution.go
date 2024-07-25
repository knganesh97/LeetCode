package problem2191

import (
	"sort"
	"strconv"
	"strings"
)

func sortJumbled(mapping []int, nums []int) []int {

	n := len(nums)

	type number struct {
		original  string
		converted string
		index     int
	}

	var numbers = make([]number, n)

	var convert = func(num string) string {
		var converted strings.Builder
		for i := range num {
			digit := int(num[i] - '0')
			converted.WriteString(strconv.Itoa(mapping[digit]))
		}
		num = converted.String()
		for len(num) > 1 && num[0] == '0' {
			num = num[1:]
		}
		return num
	}

	for i, num := range nums {
		numbers[i].index = i
		numbers[i].original = strconv.Itoa(num)
		numbers[i].converted = convert(numbers[i].original)
	}

	sort.Slice(numbers, func(i, j int) bool {
		if numbers[i].converted == numbers[j].converted {
			return numbers[i].index < numbers[j].index
		}
		if len(numbers[i].converted) != len(numbers[j].converted) {
			return len(numbers[i].converted) < len(numbers[j].converted)
		}
		return numbers[i].converted < numbers[j].converted
	})

	var jumbled []int
	for _, num := range numbers {
		jumbled = append(jumbled, nums[num.index])
	}

	return jumbled
}
