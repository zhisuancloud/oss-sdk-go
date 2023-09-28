// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/autoscaling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes one or more instances from the specified Auto Scaling group. After the
// instances are detached, you can manage them independent of the Auto Scaling
// group. If you do not specify the option to decrement the desired capacity,
// Amazon EC2 Auto Scaling launches instances to replace the ones that are
// detached. If there is a Classic Load Balancer attached to the Auto Scaling
// group, the instances are deregistered from the load balancer. If there are
// target groups attached to the Auto Scaling group, the instances are deregistered
// from the target groups. For more information, see Detach EC2 instances from your
// Auto Scaling group
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/detach-instance-asg.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) DetachInstances(ctx context.Context, params *DetachInstancesInput, optFns ...func(*Options)) (*DetachInstancesOutput, error) {
	if params == nil {
		params = &DetachInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetachInstances", params, optFns, c.addOperationDetachInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetachInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetachInstancesInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// Indicates whether the Auto Scaling group decrements the desired capacity value
	// by the number of instances detached.
	//
	// This member is required.
	ShouldDecrementDesiredCapacity *bool

	// The IDs of the instances. You can specify up to 20 instances.
	InstanceIds []string

	noSmithyDocumentSerde
}

type DetachInstancesOutput struct {

	// The activities related to detaching the instances from the Auto Scaling group.
	Activities []types.Activity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetachInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDetachInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDetachInstances{}, middleware.After)
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
	if err = addOpDetachInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetachInstances(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetachInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DetachInstances",
	}
}
