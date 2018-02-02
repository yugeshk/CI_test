package mypack

import (
	"strconv"
	"strings"
)

func checkip(ip string) bool {
	var i = 0
	var ipSlice = strings.Split(ip, ".")
	if len(ipSlice) == 4 {
		for i < 4 {
			var ip_piece, _ = strconv.Atoi(ipSlice[i])
			if !(ip_piece >= 0 && ip_piece <= 255) {
				// fmt.Println("Err")
				return false
			}
			i++
		}
	} else {
		// fmt.Println("Err")
		return false
	}
	return true
}

func clean(value string) string {
	var stripped string
	var asci = 0
	for i := 0; i < len(value); i++ {
		asci = int(value[i])
		if (asci >= 48 && asci <= 57) || (asci >= 65 && asci <= 90) || (asci >= 97 && asci <= 122) {
			stripped = stripped + string(value[i])
		}
	}
	return stripped
}

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter text: ")
// 	text, _ := reader.ReadString('\n')
// 	fmt.Println(checkip(text))
// }
