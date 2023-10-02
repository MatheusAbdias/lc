package colorsneithbors

func winnerOfGame(colors string) bool {
	mvP1 := 0
	mvP2 := 0
	for index := 1; index < len(colors)-1; index++ {
		if colors[index] == 'A' && colors[index-1] == 'A' && colors[index+1] == 'A' {
			mvP1++
		} else {
			if colors[index] == 'B' && colors[index-1] == 'B' && colors[index+1] == 'B' {
				mvP2++
			}
		}
	}

	return max(mvP1, mvP2)
}

func max(mvP1 int, mvP2 int) bool {
	return mvP1 > mvP2
}
