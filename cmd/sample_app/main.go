package main

import (
	"fmt"
	"os"
)

type SlackServiceInterface interface {
	Send(message string) error
}

type SlackService struct{}

func (s *SlackService) Send(message string) error {
	// 実際にSlackへ送る処理
	fmt.Println("Sending message to Slack:", message)
	return nil
}

type SlackDummyService struct{}

func (d *SlackDummyService) Send(message string) error {
	// ログへ出力する処理
	fmt.Println("Logging message (dummy):", message)
	return nil
}

func NewSlackService() SlackServiceInterface {
	if os.Getenv("ENV") == "local" {
		return &SlackDummyService{}
	}
	return &SlackService{}
}

func main() {
	service := NewSlackService()
	err := service.Send("Dummy message")
	if err != nil {
		fmt.Println(err)
	}
}
