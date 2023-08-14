package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type HttpError struct {
	StatusCode int
	URL        string
	err        error
}

func (he HttpError) Error() string {
	return fmt.Sprintf("http status code = %d, URL = %s", he.StatusCode, he.URL)
}

func (he HttpError) Unwrap() error {
	return he.err
}

func ReadContents(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, &HttpError{
			StatusCode: resp.StatusCode,
			URL:        url,
			err:        errors.New("http get error"),
		}
	}
	return io.ReadAll(resp.Body)
}

func main() {
	body, err := ReadContents("https://google.com/hogehoge")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(body)
}
