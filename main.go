package main

import (
	"bufio"
	"fmt"
	"os"
)

// import "strconv"
// import "strings"

func getSingleString() {
	// 文字列１つのみ
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()
	fmt.Println(s)
}

func main() {
}
