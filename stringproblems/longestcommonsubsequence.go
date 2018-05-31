package stringproblems

// LongestCommonSubsequence implements a dynamic programming solution for the LCS problem.
func LongestCommonSubsequence(seqA, seqB string) int {
	L := make([][]int, len(seqA)+1)
	for i := range L {
		L[i] = make([]int, len(seqB)+1)
	}

	for i := 1; i < len(seqA)+1; i++ {
		for j := 1; j < len(seqB)+1; j++ {
			if seqA[i-1] == seqB[j-1] {
				L[i][j] = L[i-1][j-1] + 1

			} else {
				L[i][j] = L[i][j-1]
			}
		}
	}

	return L[len(seqA)][len(seqB)]
}
