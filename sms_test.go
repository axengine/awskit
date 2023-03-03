package awskit

import (
	"context"
	"fmt"
	"testing"
)

func TestSendSMS(t *testing.T) {
	cli, err := New()
	if err != nil {
		t.Fatal(err)
	}
	msgId, err := cli.SendSMS(context.Background(), &SMSMsg{
		SenderID:    "",
		MaxPrice:    0,
		MessageType: Transactional,
		Content:     "Your verification code is:WWER",
		Recipient:   "+11111111111",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(*msgId)
}
