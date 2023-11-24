package main

func BinarySearch() int {
	haystack := [10]int{5, 6, 7, 8, 9, 10, 11, 12, 19, 20}
	neddle := 20
	lo := 0
	hi := len(haystack) - 1

	if haystack[lo] > neddle || neddle > haystack[hi] {
		return -1
	} else if haystack[hi] == neddle {
		return hi
	} else if haystack[lo] == neddle {
		return lo
	}

	for lo < hi {
		middle := lo + (hi-lo)/2
		if haystack[middle] == neddle {
			return middle
		} else if haystack[middle] > neddle {
			hi = middle + 1
		} else if haystack[middle] < neddle {
			lo = middle
		}
	}

	return -1
}
