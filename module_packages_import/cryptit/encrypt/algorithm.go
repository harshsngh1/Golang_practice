package encrypt

func Nimbus(str string) string {
	encrypted := ""
	for _, c := range str {
		ascii := int(c)
		character := string(ascii + 3)
		encrypted += character
	}
	return encrypted
}
