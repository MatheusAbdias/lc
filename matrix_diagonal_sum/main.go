package matrixdiagonalsum

func diagonalSum(mat [][]int) int {
	sum := 0
	for ip, is := 0, len(mat)-1; ip < len(mat) && is >= 0; ip, is = ip+1, is-1 {
		if ip != is {
			sum += mat[ip][ip]
			sum += mat[ip][is]
		} else {
			sum += mat[ip][ip]
		}
	}

	return sum
}
