package main

func cmpFeed(a, b itemFeed) int {
	if a.afinidad != b.afinidad {
		return b.afinidad - a.afinidad
	}
	return b.idPost - a.idPost
}

func cmpString(a, b string) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
