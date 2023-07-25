package main

import "fmt"

type Consumer struct {
	ActiveFlg bool
}
type Consumers []Consumer

func (c Consumers) ActiveCustomer() Consumers {
	resp := make([]Consumer, 0, len(c))
	for _, v := range c {
		if v.ActiveFlg {
			resp = append(resp, v)
		}
	}
	return resp
}

func main() {
	consumers := Consumers{Consumer{
		ActiveFlg: true,
	}, Consumer{
		ActiveFlg: false,
	}}
	resp := consumers.ActiveCustomer()
	fmt.Println(resp)
}
