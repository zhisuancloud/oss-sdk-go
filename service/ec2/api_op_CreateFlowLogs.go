// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates one or more flow logs to capture information about IP traffic for a
// specific network interface, subnet, or VPC. Flow log data for a monitored
// network interface is recorded as flow log records, which are log events
// consisting of fields that describe the traffic flow. For more information, see
// Flow log records
// (https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html#flow-log-records)
// in the Amazon Virtual Private Cloud User Guide. When publishing to CloudWatch
// Logs, flow log records are published to a log group, and each network interface
// has a unique log stream in the log group. When publishing to Amazon S3, flow log
// records for all of the monitored network interfaces are published to a single
// log file object that is stored in the specified bucket. For more information,
// see VPC Flow Logs
// (https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html) in the Amazon
// Virtual Private Cloud User Guide.
func (c *Client) CreateFlowLogs(ctx context.Context, params *CreateFlowLogsInput, optFns ...func(*Options)) (*CreateFlowLogsOutput, error) {
	if params == nil {
		params = &CreateFlowLogsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFlowLogs", params, optFns, c.addOperationCreateFlowLogsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFlowLogsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFlowLogsInput struct {

	// The IDs of the resources to monitor. For example, if the resource type is VPC,
	// specify the IDs of the VPCs. Constraints: Maximum of 1000 resources
	//
	// This member is required.
	ResourceIds []string

	// The type of resource to monitor.
	//
	// This member is required.
	ResourceType types.FlowLogsResourceType

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. For more information, see How to ensure idempotency
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Run_Instance_Idempotency.html).
	ClientToken *string

	// The ARN of the IAM role that allows Amazon EC2 to publish flow logs across
	// accounts.
	DeliverCrossAccountRole *string

	// The ARN of the IAM role that allows Amazon EC2 to publish flow logs to a
	// CloudWatch Logs log group in your account. This parameter is required if the
	// destination type is cloud-watch-logs and unsupported otherwise.
	DeliverLogsPermissionArn *string

	// The destination options.
	DestinationOptions *types.DestinationOptionsRequest

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The destination for the flow log data. The meaning of this parameter depends on
	// the destination type.
	//
	// * If the destination type is cloud-watch-logs, specify
	// the ARN of a CloudWatch Logs log group. For example:
	// arn:aws:logs:region:account_id:log-group:my_group Alternatively, use the
	// LogGroupName parameter.
	//
	// * If the destination type is s3, specify the ARN of an
	// S3 bucket. For example: arn:aws:s3:::my_bucket/my_subfolder/ The subfolder is
	// optional. Note that you can't use AWSLogs as a subfolder name.
	//
	// * If the
	// destination type is kinesis-data-firehose, specify the ARN of a Kinesis Data
	// Firehose delivery stream. For example:
	// arn:aws:firehose:region:account_id:deliverystream:my_stream
	LogDestination *string

	// The type of destination for the flow log data. Default: cloud-watch-logs
	LogDestinationType types.LogDestinationType

	// The fields to include in the flow log record. List the fields in the order in
	// which they should appear. For more information about the available fields, see
	// Flow log records
	// (https://docs.aws.amazon.com/vpc/latest/userguide/flow-logs.html#flow-log-records).
	// If you omit this parameter, the flow log is created using the default format. If
	// you specify this parameter, you must include at least one field. Specify the
	// fields using the ${field-id} format, separated by spaces. For the CLI, surround
	// this parameter value with single quotes on Linux or double quotes on Windows.
	LogFormat *string

	// The name of a new or existing CloudWatch Logs log group where Amazon EC2
	// publishes your flow logs. This parameter is valid only if the destination type
	// is cloud-watch-logs.
	LogGroupName *string

	// The maximum interval of time during which a flow of packets is captured and
	// aggregated into a flow log record. You can specify 60 seconds (1 minute) or 600
	// seconds (10 minutes). When a network interface is attached to a Nitro-based
	// instance
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances),
	// the aggregation interval is always 60 seconds or less, regardless of the value
	// that you specify. Default: 600
	MaxAggregationInterval *int32

	// The tags to apply to the flow logs.
	TagSpecifications []types.TagSpecification

	// The type of traffic to monitor (accepted traffic, rejected traffic, or all
	// traffic).
	TrafficType types.TrafficType

	noSmithyDocumentSerde
}

type CreateFlowLogsOutput struct {

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	ClientToken *string

	// The IDs of the flow logs.
	FlowLogIds []string

	// Information about the flow logs that could not be created successfully.
	Unsuccessful []types.UnsuccessfulItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFlowLogsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateFlowLogs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateFlowLogs{}, middleware.After)
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
	if err = addOpCreateFlowLogsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFlowLogs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFlowLogs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateFlowLogs",
	}
}
