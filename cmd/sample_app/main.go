package main

import "fmt"

type SKUCode string

func (c SKUCode) Invalid() bool {
	// 桁数などのチェック
	return len(c) != 9 // SKUコードが9桁でなければ無効とする例
}

func (c SKUCode) ItemCD() string {
	return string(c[0:5])
}

func (c SKUCode) SizeCD() string {
	return string(c[5:7])
}

func (c SKUCode) ColorCD() string {
	return string(c[7:9])
}

func main() {
	sku := SKUCode("123456789")

	if sku.Invalid() {
		fmt.Println("Invalid SKU Code")
	} else {
		fmt.Println("Item Code:", sku.ItemCD())
		fmt.Println("Size Code:", sku.SizeCD())
		fmt.Println("Color Code:", sku.ColorCD())
	}
}
