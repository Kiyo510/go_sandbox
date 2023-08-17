package main

import (
	"errors"
	"fmt"
)

// ConvertSlices []Tを[]Uへ変換
func ConvertSlices[T, U any](srcList []T, convertFunc func(T) (U, error)) ([]U, error) {
	result := make([]U, len(srcList))
	for _, v := range srcList {
		u, err := convertFunc(v)
		if err != nil {
			return nil, err
		}
		result = append(result, u)
	}
	return result, nil
}

func main() {
	var fishList = []any{"鯖", "鰤", "鮪"}
	strs, err := ConvertSlices(fishList, func(f any) (string, error) {
		if fn, ok := f.(string); ok {
			return fn, nil
		}
		return "", errors.New("cannot assertion")
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(strs)
}
