package main

import "fmt"

/*
結果:
len=6 cap=6 [2 3 5 7 11 13]
len=6 cap=6 [2 3 5 7 11 13]
len=0 cap=6 []
len=4 cap=6 [2 3 5 7]
len=2 cap=4 [5 7]

s[2:] は cap が変化し、s[2:] は cap が変化しない理由は以下の通り。

> スライス式の左が0もしくは記述しない場合、capが変わることはありません。
> 元の配列分の容量を確保してくれます。
> https://qiita.com/Kashiwara/items/e621a4ad8ec00974f025 より

また、スライスの cap を超えて値を追加する場合、cap は元の2倍の容量で確保される。

> これはGo言語の仕様によるもので、容量オーバーしたら基本的に２倍の値を確保します。
> appendするたびに一々メモリ領域を確保してコピーする手順を踏むのは効率が悪いため
> このような動作になっています。
*/
func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	printSlice(s)

	s = s[:]
	printSlice(s)

	s = s[:0]
	printSlice(s)

	s = s[:4]
	printSlice(s)

	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
