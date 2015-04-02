package caesar

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func rotate(s string, rot int) string {
	rot %= 26
	b := []byte(s)
	for i, c := range b {
		c |= 0x20
		if 'a' <= c && c <= 'z' {
			b[i] = alphabet[(int(('z'-'a'+1)+(c-'a'))+rot)%26]
		}
	}
	return string(b)
}

func Decode(cipher string, rot int) (text string) {
	return rotate(cipher, -rot)
}

func Encode(text string, rot int) (cipher string) {
	return rotate(text, rot)
}
