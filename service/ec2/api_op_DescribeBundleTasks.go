// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Describes the specified bundle tasks or all of your bundle tasks. Completed
// bundle tasks are listed for only a limited time. If your bundle task is no
// longer in the list, you can still register an AMI from it. Just use
// RegisterImage with the Amazon S3 bucket name and image manifest name you
// provided to the bundle task.
func (c *Client) DescribeBundleTasks(ctx context.Context, params *DescribeBundleTasksInput, optFns ...func(*Options)) (*DescribeBundleTasksOutput, error) {
	if params == nil {
		params = &DescribeBundleTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBundleTasks", params, optFns, c.addOperationDescribeBundleTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBundleTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBundleTasksInput struct {

	// The bundle task IDs. Default: Describes all your bundle tasks.
	BundleIds []string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The filters.
	//
	// * bundle-id - The ID of the bundle task.
	//
	// * error-code - If the
	// task failed, the error code returned.
	//
	// * error-message - If the task failed, the
	// error message returned.
	//
	// * instance-id - The ID of the instance.
	//
	// * progress -
	// The level of task completion, as a percentage (for example, 20%).
	//
	// * s3-bucket -
	// The Amazon S3 bucket to store the AMI.
	//
	// * s3-prefix - The beginning of the AMI
	// name.
	//
	// * start-time - The time the task started (for example,
	// 2013-09-15T17:15:20.000Z).
	//
	// * state - The state of the task (pending |
	// waiting-for-shutdown | bundling | storing | cancelling | complete | failed).
	//
	// *
	// update-time - The time of the most recent update for the task.
	Filters []types.Filter

	noSmithyDocumentSerde
}

type DescribeBundleTasksOutput struct {

	// Information about the bundle tasks.
	BundleTasks []types.BundleTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBundleTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeBundleTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeBundleTasks{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBundleTasks(options.Region), middleware.Before); err != nil {
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

// DescribeBundleTasksAPIClient is a client that implements the DescribeBundleTasks
// operation.
type DescribeBundleTasksAPIClient interface {
	DescribeBundleTasks(context.Context, *DescribeBundleTasksInput, ...func(*Options)) (*DescribeBundleTasksOutput, error)
}

var _ DescribeBundleTasksAPIClient = (*Client)(nil)

// BundleTaskCompleteWaiterOptions are waiter options for BundleTaskCompleteWaiter
type BundleTaskCompleteWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// BundleTaskCompleteWaiter will use default minimum delay of 15 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, BundleTaskCompleteWaiter will use default max delay of 120 seconds.
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
	Retryable func(context.Context, *DescribeBundleTasksInput, *DescribeBundleTasksOutput, error) (bool, error)
}

// BundleTaskCompleteWaiter defines the waiters for BundleTaskComplete
type BundleTaskCompleteWaiter struct {
	client DescribeBundleTasksAPIClient

	options BundleTaskCompleteWaiterOptions
}

// NewBundleTaskCompleteWaiter constructs a BundleTaskCompleteWaiter.
func NewBundleTaskCompleteWaiter(client DescribeBundleTasksAPIClient, optFns ...func(*BundleTaskCompleteWaiterOptions)) *BundleTaskCompleteWaiter {
	options := BundleTaskCompleteWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = bundleTaskCompleteStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &BundleTaskCompleteWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for BundleTaskComplete waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *BundleTaskCompleteWaiter) Wait(ctx context.Context, params *DescribeBundleTasksInput, maxWaitDur time.Duration, optFns ...func(*BundleTaskCompleteWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for BundleTaskComplete waiter and
// returns the output of the successful operation. The maxWaitDur is the maximum
// wait duration the waiter will wait. The maxWaitDur is required and must be
// greater than zero.
func (w *BundleTaskCompleteWaiter) WaitForOutput(ctx context.Context, params *DescribeBundleTasksInput, maxWaitDur time.Duration, optFns ...func(*BundleTaskCompleteWaiterOptions)) (*DescribeBundleTasksOutput, error) {
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

		out, err := w.client.DescribeBundleTasks(ctx, params, func(o *Options) {
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
	return nil, fmt.Errorf("exceeded max wait time for BundleTaskComplete waiter")
}

func bundleTaskCompleteStateRetryable(ctx context.Context, input *DescribeBundleTasksInput, output *DescribeBundleTasksOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("BundleTasks[].State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "complete"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(types.BundleTaskState)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.BundleTaskState value, got %T", pathValue)
			}

			if string(value) != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("BundleTasks[].State", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "failed"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(types.BundleTaskState)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.BundleTaskState value, got %T", pathValue)
			}

			if string(value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeBundleTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeBundleTasks",
	}
}
