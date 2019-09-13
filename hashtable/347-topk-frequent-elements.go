package hashmap

func topKFrequent(nums []int, k int) []int {
	// key number
	// value frequency
	record := make(map[int]int)

	for _, num := range nums {
		record[num]++
	}

	// index frequency
	// value number
	bucket := make([][]int, len(nums)+1)

	for num, freq := range record {
		bucket[freq] = append(bucket[freq], num)
	}

	res := []int{}

	for i := len(bucket) - 1; i >= 0; i-- {
		if bucket[i] != nil {
			res = append(res, bucket[i]...)
		}
	}

	return res[:k]

}
