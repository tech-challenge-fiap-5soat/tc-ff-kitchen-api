package interfaces

type PublisherGateway interface {
	PublishMessage(queueName, message string) error
	GetQueueUrl() string
}
