package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/text/unicode/norm"
)

func Normalize(w io.Writer, r io.Reader) error {
	br := bufio.NewReader(r)
	for {
		s, err := br.ReadString('\n')
		if s != "" {
			_, err := io.WriteString(w, norm.NFKC.String(s))
			if err != nil {
				return err
			}
		}
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
	}
}

func NormalizeFile(input, output string) error {
	r, err := os.Open(input)
	if err != nil {
		return err
	}
	w, err := os.Create(output)
	if err != nil {
		return err
	}
	return Normalize(w, r)
}

func NormalizeString(i string) (string, error) {
	r := strings.NewReader(i)
	var w strings.Builder
	err := Normalize(&w, r)
	if err != nil {
		return "", err
	}
	return w.String(), nil
}

func main() {
	str, err := NormalizeString("ホゲ 　　　ｌホゲ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)

	err = NormalizeFile("input.txt", "output.txt")
	if err != nil {
		fmt.Println(err)
	}
}
