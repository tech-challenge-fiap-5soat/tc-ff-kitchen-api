package gateway

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/tech-challenge-fiap-5soat/tc-ff-kitchen-api/src/common/interfaces"
)

type PublisherGatewayConfig struct {
	SQSQueueUrl        string
	SQSEndpoint        string
	AWSRegion          string
	AWSAccessKeyID     string
	AWSSecretAccessKey string
}

type PublisherGateway struct {
	sqsClient *sqs.Client
	QueueUrl  string
}

func NewPublisherGateway(pgconfig PublisherGatewayConfig) interfaces.PublisherGateway {
	var cfg aws.Config
	var err error

	if pgconfig.SQSEndpoint != "" {
		cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(pgconfig.AWSRegion),
			config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(
					pgconfig.AWSAccessKeyID, pgconfig.AWSSecretAccessKey, "")),
			config.WithEndpointResolver(
				aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
					return aws.Endpoint{
						URL:           pgconfig.SQSEndpoint,
						SigningRegion: region,
					}, nil
				}),
			),
		)
	} else {
		cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithRegion(pgconfig.AWSRegion),
			config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(
					pgconfig.AWSAccessKeyID, pgconfig.AWSSecretAccessKey, "")),
		)
	}

	if err != nil {
		fmt.Printf("unable to load SDK config, %v", err)
	}

	sqsClient := sqs.NewFromConfig(cfg)

	publisher := &PublisherGateway{
		sqsClient: sqsClient,
		QueueUrl:  pgconfig.SQSQueueUrl,
	}
	return publisher
}

func (pg *PublisherGateway) GetQueueUrl() string {
	return pg.QueueUrl
}

func (pg *PublisherGateway) PublishMessage(queueName, message string) error {
	_, err := pg.sqsClient.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:     &queueName,
		MessageBody:  aws.String(message),
		DelaySeconds: 5,
	})
	if err != nil {
		return fmt.Errorf("error occurred while sending message to SQS: %s", err.Error())
	}
	return nil
}
