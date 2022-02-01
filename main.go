package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	// 数値１つのみ
	// 例
	// ==========
	// 123
	// ==========
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	fmt.Println(i)
}

func getMultipleString() {
	// 1行でスペース区切りの文字列を取得
	// 例
	// ==========
	// sample sample1 sample2
	// ==========
	// スペースで区切った文字列を配列lに格納する

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	l := strings.Split(sc.Text(), " ")
	fmt.Println(l)
}

func getMultipleRowsString() {
	// 行数NとN行の文字列
	// 例
	// ==========
	// 3
	// sample
	// sample2
	// sample3
	// ==========
	// N個の文字列はスライスmに格納する
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	m := []string{}
	for i := 0; i < n; i++ {
		sc.Scan()
		m = append(m, sc.Text())
	}
	fmt.Println(n)
	fmt.Println(m)
}

func main() {
	// getSingleString()
	// getSingleInt()
	getMultipleString()
	// getMultipleRowsString()
}
