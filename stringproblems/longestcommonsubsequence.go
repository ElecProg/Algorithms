package stringproblems

// LongestCommonSubsequence implements a dynamic programming solution for the LCS problem.
func LongestCommonSubsequence(seqA, seqB string) int {
	L := make([][]int, len(seqA))
	for i := range L {
		L[i] = make([]int, len(seqB))
	}

	for i := 1; i < len(seqA)+1; i++ {
		for j := 1; j < len(seqB)+1; j++ {
			if seqA[i-1] == seqB[j-1] {
				L[i][j] = get(L, i-1, j-1) + 1

			} else {
				L[i][j] = max(get(L, i, j-1), get(L, i-1, j))
			}
		}
	}

	return L[len(seqA)-1][len(seqB)-1]
}

func get(matrix [][]int, i, j int) int {
	if i < 0 || j < 0 {
		return 0
	}

	return matrix[i][j]
}
