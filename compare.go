package piscine

func Compare(s, toFind string) int {
	if s == toFind{
		return 0
	}else if s < toFind{
		return -1
	}else{
		return 1
	}
}