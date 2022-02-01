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

func getMultipleInt() {
	// 1行でスペース区切りの数値を取得
	// 例
	// ==========
	// 123 234 345
	// ==========
	// スペースで区切った数値を配列lに格納する

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	l := []int{}
	for _, s := range strings.Split(sc.Text(), " ") {
		a, _ := strconv.Atoi(s)
		l = append(l, a)
	}
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

func getMultipleRowsInt() {
	// 行数NとN行の数値
	// 例
	// ==========
	// 3
	// 123
	// 234
	// 345
	// ==========
	// N個の数値はスライスmに格納する
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	m := []int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		m = append(m, a)
	}
	fmt.Println(n)
	fmt.Println(m)
}

func getMultipleRowsAndMultipleString() {
	// 行数NとN行の文字列（スペース区切りで複数個）
	// 例
	// ==========
	// 3
	// sample1-1 sample1-2 sample1-3
	// sample2-1 sample2-2 sample2-3
	// sample3-1 sample3-2 sample3-3
	// ==========
	// N行の文字列をスペースで区切り、スライスmに格納する
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	m := []string{}
	for i := 0; i < n; i++ {
		sc.Scan()
		t := strings.Split(sc.Text(), " ")
		m = append(m, t...)
	}
	fmt.Println(n)
	fmt.Println(m)
}

func main() {
	// getSingleString()
	// getSingleInt()
	// getMultipleString()
	// getMultipleInt()
	// getMultipleRowsString()
	// getMultipleRowsInt()
	getMultipleRowsAndMultipleString()
}
