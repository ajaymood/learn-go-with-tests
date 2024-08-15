package iteration

func Iterator(inputChar string, times int) string {
	// repeatedBytes := make([]byte, 5)
	// for i := 0; i < 5; i++ {
	// 	repeatedBytes = append(repeatedBytes, byte(" "[0])+byte(inputChar[0]))
	// }
	// return string(repeatedBytes)
	repeatedChar := ""
	for i := 0; i < times-1; i++ {
		repeatedChar = repeatedChar + inputChar + " "
	}
	return repeatedChar + inputChar
}
