// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/cloudwatch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Temporarily sets the state of an alarm for testing purposes. When the updated
// state differs from the previous value, the action configured for the appropriate
// state is invoked. For example, if your alarm is configured to send an Amazon SNS
// message when an alarm is triggered, temporarily changing the alarm state to
// ALARM sends an SNS message. Metric alarms returns to their actual state quickly,
// often within seconds. Because the metric alarm state change happens quickly, it
// is typically only visible in the alarm's History tab in the Amazon CloudWatch
// console or through DescribeAlarmHistory
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_DescribeAlarmHistory.html).
// If you use SetAlarmState on a composite alarm, the composite alarm is not
// guaranteed to return to its actual state. It returns to its actual state only
// once any of its children alarms change state. It is also reevaluated if you
// update its configuration. If an alarm triggers EC2 Auto Scaling policies or
// application Auto Scaling policies, you must include information in the
// StateReasonData parameter to enable the policy to take the correct action.
func (c *Client) SetAlarmState(ctx context.Context, params *SetAlarmStateInput, optFns ...func(*Options)) (*SetAlarmStateOutput, error) {
	if params == nil {
		params = &SetAlarmStateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetAlarmState", params, optFns, c.addOperationSetAlarmStateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetAlarmStateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetAlarmStateInput struct {

	// The name of the alarm.
	//
	// This member is required.
	AlarmName *string

	// The reason that this alarm is set to this specific state, in text format.
	//
	// This member is required.
	StateReason *string

	// The value of the state.
	//
	// This member is required.
	StateValue types.StateValue

	// The reason that this alarm is set to this specific state, in JSON format. For
	// SNS or EC2 alarm actions, this is just informational. But for EC2 Auto Scaling
	// or application Auto Scaling alarm actions, the Auto Scaling policy uses the
	// information in this field to take the correct action.
	StateReasonData *string

	noSmithyDocumentSerde
}

type SetAlarmStateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetAlarmStateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpSetAlarmState{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpSetAlarmState{}, middleware.After)
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
	if err = addOpSetAlarmStateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetAlarmState(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetAlarmState(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "SetAlarmState",
	}
}
