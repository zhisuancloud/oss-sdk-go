// Code generated by smithy-go-codegen DO NOT EDIT.

package codegurureviewer

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/codegurureviewer/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Returns the metadata associated with the code review along with its status.
func (c *Client) DescribeCodeReview(ctx context.Context, params *DescribeCodeReviewInput, optFns ...func(*Options)) (*DescribeCodeReviewOutput, error) {
	if params == nil {
		params = &DescribeCodeReviewInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCodeReview", params, optFns, c.addOperationDescribeCodeReviewMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCodeReviewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCodeReviewInput struct {

	// The Amazon Resource Name (ARN) of the CodeReview
	// (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_CodeReview.html)
	// object.
	//
	// This member is required.
	CodeReviewArn *string

	noSmithyDocumentSerde
}

type DescribeCodeReviewOutput struct {

	// Information about the code review.
	CodeReview *types.CodeReview

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCodeReviewMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeCodeReview{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeCodeReview{}, middleware.After)
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
	if err = addOpDescribeCodeReviewValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCodeReview(options.Region), middleware.Before); err != nil {
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

// DescribeCodeReviewAPIClient is a client that implements the DescribeCodeReview
// operation.
type DescribeCodeReviewAPIClient interface {
	DescribeCodeReview(context.Context, *DescribeCodeReviewInput, ...func(*Options)) (*DescribeCodeReviewOutput, error)
}

var _ DescribeCodeReviewAPIClient = (*Client)(nil)

// CodeReviewCompletedWaiterOptions are waiter options for
// CodeReviewCompletedWaiter
type CodeReviewCompletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// CodeReviewCompletedWaiter will use default minimum delay of 10 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, CodeReviewCompletedWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeCodeReviewInput, *DescribeCodeReviewOutput, error) (bool, error)
}

// CodeReviewCompletedWaiter defines the waiters for CodeReviewCompleted
type CodeReviewCompletedWaiter struct {
	client DescribeCodeReviewAPIClient

	options CodeReviewCompletedWaiterOptions
}

// NewCodeReviewCompletedWaiter constructs a CodeReviewCompletedWaiter.
func NewCodeReviewCompletedWaiter(client DescribeCodeReviewAPIClient, optFns ...func(*CodeReviewCompletedWaiterOptions)) *CodeReviewCompletedWaiter {
	options := CodeReviewCompletedWaiterOptions{}
	options.MinDelay = 10 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = codeReviewCompletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &CodeReviewCompletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for CodeReviewCompleted waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *CodeReviewCompletedWaiter) Wait(ctx context.Context, params *DescribeCodeReviewInput, maxWaitDur time.Duration, optFns ...func(*CodeReviewCompletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for CodeReviewCompleted waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *CodeReviewCompletedWaiter) WaitForOutput(ctx context.Context, params *DescribeCodeReviewInput, maxWaitDur time.Duration, optFns ...func(*CodeReviewCompletedWaiterOptions)) (*DescribeCodeReviewOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeCodeReview(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for CodeReviewCompleted waiter")
}

func codeReviewCompletedStateRetryable(ctx context.Context, input *DescribeCodeReviewInput, output *DescribeCodeReviewOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("CodeReview.State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Completed"
		value, ok := pathValue.(types.JobState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.JobState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("CodeReview.State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Failed"
		value, ok := pathValue.(types.JobState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.JobState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("CodeReview.State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "Pending"
		value, ok := pathValue.(types.JobState)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.JobState value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeCodeReview(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-reviewer",
		OperationName: "DescribeCodeReview",
	}
}
