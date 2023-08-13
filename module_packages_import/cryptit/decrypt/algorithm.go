package decrypt

func Nimbus(str string) string {
	decryptstr := ""

	for _, c := range str {
		asciiCode := int(c)
		charecter := string(asciiCode - 3)
		decryptstr += charecter
	}
	return decryptstr
}
