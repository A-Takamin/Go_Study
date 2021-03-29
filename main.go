package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	// アドレスとポインタについて
	var hoge int = 5
	var hoge2 *int = &hoge
	var hoge3 int = *hoge2
	// &をつけるとアドレスを取得。
	// &をつけてアドレスを取得した変数の型は「*」をつけて表す。この型をポインタ型という。
	fmt.Println(hoge2)
	fmt.Println(hoge3)

	// 最初
	p := Person{
		Name: "太郎",
		Age:  20,
	}
	fmt.Printf("最初のp :%+v\n", p)

	// 二番目
	p2 := p // こいつ…値のコピーをしてやがる（Javaとは違う）
	p2.Name = "次郎"
	p2.Age = 21
	fmt.Printf("次のp2 :%+v\n", p2)

	fmt.Printf("　最初のp :%+v\n", p)

	// 3番目
	p3 := &p //変数のアドレスを取得
	p3.Name = "三郎"
	p3.Age = 22
	fmt.Printf("3つめp3 :%+v\n", p3)

	fmt.Printf("　最初のp :%+v\n", p)

}
