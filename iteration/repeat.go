package iteration

func Repeat(ch string) string {
	var repeatedStr string
	for i := 0; i < 5; i++ {
		repeatedStr = repeatedStr + ch
	}
	return repeatedStr
}
