package help

func Run() {
	helpMsg := `
ccwc - word, line, character, and byte count

Input:
- relative path to a file

Args:
- help: display this help message

Flags:
-c: outputs the number of bytes in a file
-l: outputs the number of lines in a file
-w: outputs the number of words in a file
-m: outputs the number of characters in a file
	`

	println(helpMsg)
}
