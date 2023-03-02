package awskit

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/pkg/errors"
)

const (
	// SenderIDSMSAttribute a custom name that's displayed as the message sender on the receiving device.
	SenderIDSMSAttribute = "SenderID"
	// MaxPriceSMSAttribute a maximum price in USD that you are willing to pay to send the message.
	MaxPriceSMSAttribute = "MaxPrice"
	// MessageTypeSMSAttribute a SMS type, can be either Promotional or Transactional.
	MessageTypeSMSAttribute = "MessageType"

	// Promotional message type used for promotional purposes that are noncritical, won't be delivered
	// to DND (Do Not Disturb) numbers.
	Promotional MessageType = "Promotional"
	// Transactional message type used for transactional purposes which includes critical messages as multi-factor
	// authentication. This message type might be more expensive than Promotional message type. Will be delivered to
	// to DND numbers.
	Transactional MessageType = "Transactional"
)

// MessageType SNS SMS type
type MessageType string

// IsValid returns true if the message type is either
// transactional or promotional used to validate that
// a valid message type is being provided.
func (mt MessageType) IsValid() bool {
	return mt == Transactional || mt == Promotional
}

// String returns message type string value
func (mt MessageType) String() string {
	return string(mt)
}

type SMSMsg struct {
	SenderID    string      // Default ""
	MaxPrice    float32     // Default 0.05
	MessageType MessageType // Default Transactional
	Content     string      //
	Recipient   string      // E.164
	attr        map[string]types.MessageAttributeValue
}

func (sm *SMSMsg) attribute() map[string]types.MessageAttributeValue {
	sm.attr = make(map[string]types.MessageAttributeValue)

	if len(sm.SenderID) != 0 {
		sm.attr[SenderIDSMSAttribute] = types.MessageAttributeValue{
			DataType:    aws.String("String"),
			StringValue: &sm.SenderID,
		}
	}

	// 忽略浮点数比较问题
	if sm.MaxPrice == 0 {
		sm.MaxPrice = 0.05
	}
	sm.attr[MaxPriceSMSAttribute] = types.MessageAttributeValue{
		DataType:    aws.String("String"),
		StringValue: aws.String(fmt.Sprintf("%.2f", sm.MaxPrice)),
	}

	if len(sm.MessageType) == 0 {
		sm.MessageType = Transactional
	}
	sm.attr[MessageTypeSMSAttribute] = types.MessageAttributeValue{
		DataType:    aws.String("String"),
		StringValue: aws.String(string(sm.MessageType)),
	}

	return sm.attr
}

// SendSMS message to the given message to the receiver mobile phone number
func (ak *AWSKit) SendSMS(ctx context.Context, msg *SMSMsg) (*string, error) {
	if !msg.MessageType.IsValid() {
		return nil, errors.New("invalid msg type")
	}
	params := &sns.PublishInput{
		Message:           aws.String(msg.Content),
		PhoneNumber:       aws.String(msg.Recipient),
		MessageAttributes: msg.attribute(),
	}
	resp, err := ak.snsClient.Publish(ctx, params)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to publish a text message to %s", msg.Recipient)
	}
	return resp.MessageId, nil
}
