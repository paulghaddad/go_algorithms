package main

// IsValidSubsequence determines whether sequence is a subsequence of array
// Time: O(n)
// Space: O(1)
func IsValidSubsequence(array []int, sequence []int) bool {
	seqID := 0
	arrID := 0

	for arrID < len(array) && seqID < len(sequence) {
		if array[arrID] == sequence[seqID] {
			seqID++
		}
		arrID++
	}

	return len(sequence) == seqID
}
