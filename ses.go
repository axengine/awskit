package awskit

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/pkg/errors"
)

const (
	// Subject is the subject line for the email
	Subject = "Amazon SES Test (AWS SDK for Go)"

	// HTMLBody is the HTML body for the email
	HTMLBody = "<h1>Amazon SES Test Email (AWS SDK for Go)</h1><p>This email was sent with " +
		"<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
		"<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"

	// TextBody is the email body for recipients with non-HTML email clients
	TextBody = "This email was sent with Amazon SES using the AWS SDK for Go."

	// CharSet is the character encoding for the email
	CharSet = "UTF-8"
)

type EMailMsg struct {
	Sender    string
	Recipient string
	Subject   string
	HTMLBody  string
	TextBody  string
	CharSet   string // Default is UTF-8
}

func (em *EMailMsg) IsValid() bool {
	if len(em.Sender) == 0 ||
		len(em.Recipient) == 0 ||
		len(em.Subject) == 0 {
		return false
	}
	return true
}

func (ak *AWSKit) SendEMail(ctx context.Context, msg *EMailMsg) (*string, error) {
	if !msg.IsValid() {
		return nil, errors.New("invalid msg")
	}
	if len(msg.CharSet) == 0 {
		msg.CharSet = "UTF-8"
	}
	out, err := ak.sesClient.SendEmail(ctx, &ses.SendEmailInput{
		Destination: &types.Destination{
			CcAddresses: []string{},
			ToAddresses: []string{
				msg.Recipient,
			},
		},
		Message: &types.Message{
			Body: &types.Body{
				Html: &types.Content{
					Charset: aws.String(msg.CharSet),
					Data:    aws.String(msg.HTMLBody),
				},
				Text: &types.Content{
					Charset: aws.String(msg.CharSet),
					Data:    aws.String(msg.TextBody),
				},
			},
			Subject: &types.Content{
				Charset: aws.String(msg.CharSet),
				Data:    aws.String(msg.Subject),
			},
		},
		Source: aws.String(msg.Sender),
	})

	return out.MessageId, err
}
