package awskit

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type AWSKit struct {
	s3Client  *s3.Client
	sesClient *ses.Client
	snsClient *sns.Client
}

// New ASWKit
// profile:credentials profile in .aws dir
// region
func New(profile, region string) (*AWSKit, error) {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(profile),
		config.WithRegion(region))
	if err != nil {
		return nil, err
	}
	return &AWSKit{
		s3Client:  s3.NewFromConfig(sdkConfig),
		sesClient: ses.NewFromConfig(sdkConfig),
		snsClient: sns.NewFromConfig(sdkConfig),
	}, nil
}
