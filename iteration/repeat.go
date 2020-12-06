package iteration

const repeatedTimes = 5

func Repeat(ch string) string {
	var repeatedStr string
	for i := 0; i < repeatedTimes; i++ {
		repeatedStr += ch
	}
	return repeatedStr
}
