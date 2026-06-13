package controls

var Ways map[string]int

func InitWays() {
	Ways = make(map[string]int)
	Ways["left"] = 1
	Ways["up"] = 2
	Ways["down"] = 3
	Ways["right"] = 4
}

func GetWays(way string) (int, bool) {
	key, ok := Ways[way]
	return key, ok

}
