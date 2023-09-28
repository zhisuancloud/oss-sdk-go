// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/sqs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets attributes for the specified queue. To determine whether a queue is FIFO
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html),
// you can check whether QueueName ends with the .fifo suffix.
func (c *Client) GetQueueAttributes(ctx context.Context, params *GetQueueAttributesInput, optFns ...func(*Options)) (*GetQueueAttributesOutput, error) {
	if params == nil {
		params = &GetQueueAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetQueueAttributes", params, optFns, c.addOperationGetQueueAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetQueueAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetQueueAttributesInput struct {

	// The URL of the Amazon SQS queue whose attribute information is retrieved. Queue
	// URLs and names are case-sensitive.
	//
	// This member is required.
	QueueUrl *string

	// A list of attributes for which to retrieve information. The AttributeName.N
	// parameter is optional, but if you don't specify values for this parameter, the
	// request returns empty results. In the future, new attributes might be added. If
	// you write code that calls this action, we recommend that you structure your code
	// so that it can handle new attributes gracefully. The following attributes are
	// supported: The ApproximateNumberOfMessagesDelayed,
	// ApproximateNumberOfMessagesNotVisible, and ApproximateNumberOfMessagesVisible
	// metrics may not achieve consistency until at least 1 minute after the producers
	// stop sending messages. This period is required for the queue metadata to reach
	// eventual consistency.
	//
	// * All – Returns all values.
	//
	// *
	// ApproximateNumberOfMessages – Returns the approximate number of messages
	// available for retrieval from the queue.
	//
	// * ApproximateNumberOfMessagesDelayed –
	// Returns the approximate number of messages in the queue that are delayed and not
	// available for reading immediately. This can happen when the queue is configured
	// as a delay queue or when a message has been sent with a delay parameter.
	//
	// *
	// ApproximateNumberOfMessagesNotVisible – Returns the approximate number of
	// messages that are in flight. Messages are considered to be in flight if they
	// have been sent to a client but have not yet been deleted or have not yet reached
	// the end of their visibility window.
	//
	// * CreatedTimestamp – Returns the time when
	// the queue was created in seconds (epoch time
	// (http://en.wikipedia.org/wiki/Unix_time)).
	//
	// * DelaySeconds – Returns the default
	// delay on the queue in seconds.
	//
	// * LastModifiedTimestamp – Returns the time when
	// the queue was last changed in seconds (epoch time
	// (http://en.wikipedia.org/wiki/Unix_time)).
	//
	// * MaximumMessageSize – Returns the
	// limit of how many bytes a message can contain before Amazon SQS rejects it.
	//
	// *
	// MessageRetentionPeriod – Returns the length of time, in seconds, for which
	// Amazon SQS retains a message.
	//
	// * Policy – Returns the policy of the queue.
	//
	// *
	// QueueArn – Returns the Amazon resource name (ARN) of the queue.
	//
	// *
	// ReceiveMessageWaitTimeSeconds – Returns the length of time, in seconds, for
	// which the ReceiveMessage action waits for a message to arrive.
	//
	// * RedrivePolicy
	// – The string that includes the parameters for the dead-letter queue
	// functionality of the source queue as a JSON object. For more information about
	// the redrive policy and dead-letter queues, see Using Amazon SQS Dead-Letter
	// Queues
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html)
	// in the Amazon SQS Developer Guide.
	//
	// * deadLetterTargetArn – The Amazon Resource
	// Name (ARN) of the dead-letter queue to which Amazon SQS moves messages after the
	// value of maxReceiveCount is exceeded.
	//
	// * maxReceiveCount – The number of times a
	// message is delivered to the source queue before being moved to the dead-letter
	// queue. When the ReceiveCount for a message exceeds the maxReceiveCount for a
	// queue, Amazon SQS moves the message to the dead-letter-queue.
	//
	// *
	// VisibilityTimeout – Returns the visibility timeout for the queue. For more
	// information about the visibility timeout, see Visibility Timeout
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html)
	// in the Amazon SQS Developer Guide.
	//
	// The following attributes apply only to
	// server-side-encryption
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html):
	//
	// *
	// KmsMasterKeyId – Returns the ID of an Amazon Web Services managed customer
	// master key (CMK) for Amazon SQS or a custom CMK. For more information, see Key
	// Terms
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-sse-key-terms).
	//
	// *
	// KmsDataKeyReusePeriodSeconds – Returns the length of time, in seconds, for which
	// Amazon SQS can reuse a data key to encrypt or decrypt messages before calling
	// KMS again. For more information, see How Does the Data Key Reuse Period Work?
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-how-does-the-data-key-reuse-period-work).
	//
	// *
	// SqsManagedSseEnabled – Returns information about whether the queue is using
	// SSE-SQS encryption using SQS owned encryption keys. Only one server-side
	// encryption option is supported per queue (e.g. SSE-KMS
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-sse-existing-queue.html)
	// or SSE-SQS
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-configure-sqs-sse-queue.html)).
	//
	// The
	// following attributes apply only to FIFO (first-in-first-out) queues
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html):
	//
	// *
	// FifoQueue – Returns information about whether the queue is FIFO. For more
	// information, see FIFO queue logic
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues-understanding-logic.html)
	// in the Amazon SQS Developer Guide. To determine whether a queue is FIFO
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html),
	// you can check whether QueueName ends with the .fifo suffix.
	//
	// *
	// ContentBasedDeduplication – Returns whether content-based deduplication is
	// enabled for the queue. For more information, see Exactly-once processing
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues-exactly-once-processing.html)
	// in the Amazon SQS Developer Guide.
	//
	// The following attributes apply only to high
	// throughput for FIFO queues
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/high-throughput-fifo.html):
	//
	// *
	// DeduplicationScope – Specifies whether message deduplication occurs at the
	// message group or queue level. Valid values are messageGroup and queue.
	//
	// *
	// FifoThroughputLimit – Specifies whether the FIFO queue throughput quota applies
	// to the entire queue or per message group. Valid values are perQueue and
	// perMessageGroupId. The perMessageGroupId value is allowed only when the value
	// for DeduplicationScope is messageGroup.
	//
	// To enable high throughput for FIFO
	// queues, do the following:
	//
	// * Set DeduplicationScope to messageGroup.
	//
	// * Set
	// FifoThroughputLimit to perMessageGroupId.
	//
	// If you set these attributes to
	// anything other than the values shown for enabling high throughput, normal
	// throughput is in effect and deduplication occurs as specified. For information
	// on throughput quotas, see Quotas related to messages
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/quotas-messages.html)
	// in the Amazon SQS Developer Guide.
	AttributeNames []types.QueueAttributeName

	noSmithyDocumentSerde
}

// A list of returned queue attributes.
type GetQueueAttributesOutput struct {

	// A map of attributes to their respective values.
	Attributes map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetQueueAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetQueueAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetQueueAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetQueueAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetQueueAttributes(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetQueueAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sqs",
		OperationName: "GetQueueAttributes",
	}
}