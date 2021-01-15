package iteration

const repeatCount = 5

func Repeat(character string) string {
	str := ""

	for i := 0; i < repeatCount; i++ {
		str += character
	}
	return str
}
