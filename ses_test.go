package awskit

import (
	"context"
	"fmt"
	"testing"
)

func TestAWSKit_SendEMail(t *testing.T) {
	cli, err := New()
	if err != nil {
		t.Fatal(err)
	}
	msgId, err := cli.SendEMail(context.Background(), &EMailMsg{
		Sender:    "notify@juncachain.com",
		Recipient: "------@gmail.com",
		Subject:   "E-mail verification code",
		HTMLBody: fmt.Sprintf(`<p>Hello users!</p>
<p>Your email verification code: %s</p>
<p>The verification code is valid within 15 minutes, please verify in time~</p>`, "12345"),
		TextBody: "",
		CharSet:  "UTF-8",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(*msgId)
}
