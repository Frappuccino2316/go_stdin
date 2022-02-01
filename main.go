package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// import "strings"

func getSingleString() {
	// 文字列１つのみ
	// 例
	// ==========
	// sample
	// ==========
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()
	fmt.Println(s)
}

func getSingleInt() {
	// 文字列１つのみ
	// 例
	// ==========
	// sample
	// ==========
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	fmt.Println(i)
}

func main() {
	// getSingleString()
	getSingleInt()
}
