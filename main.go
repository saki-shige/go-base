package main

import (
	"fmt"
)

func main() {
	// fmt.Println
	// テキストを一行で出力
	fmt.Println("Hello, Golang!") // -> "Hello, Golang!"
	// 引数を複数取れる。スペースで区切られて連結
	fmt.Println("My", "Name", "Is", "Taro") // -> "My Name Is Taro"

	// fmt.Printf
	// 引数を埋め込む
	fmt.Printf("10進数=%d 2進数=%b 8進数=%o 16進数=%x\n", 17, 17, 17, 17)
	// 10進数=17 2進数=10001 8進数=21 16進数=11
}
