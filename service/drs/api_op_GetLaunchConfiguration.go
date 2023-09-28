// Code generated by smithy-go-codegen DO NOT EDIT.

package drs

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/drs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a LaunchConfiguration, filtered by Source Server IDs.
func (c *Client) GetLaunchConfiguration(ctx context.Context, params *GetLaunchConfigurationInput, optFns ...func(*Options)) (*GetLaunchConfigurationOutput, error) {
	if params == nil {
		params = &GetLaunchConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLaunchConfiguration", params, optFns, c.addOperationGetLaunchConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLaunchConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLaunchConfigurationInput struct {

	// The ID of the Source Server that we want to retrieve a Launch Configuration for.
	//
	// This member is required.
	SourceServerID *string

	noSmithyDocumentSerde
}

type GetLaunchConfigurationOutput struct {

	// Whether we should copy the Private IP of the Source Server to the Recovery
	// Instance.
	CopyPrivateIp *bool

	// Whether we want to copy the tags of the Source Server to the EC2 machine of the
	// Recovery Instance.
	CopyTags *bool

	// The EC2 launch template ID of this launch configuration.
	Ec2LaunchTemplateID *string

	// The state of the Recovery Instance in EC2 after the recovery operation.
	LaunchDisposition types.LaunchDisposition

	// The licensing configuration to be used for this launch configuration.
	Licensing *types.Licensing

	// The name of the launch configuration.
	Name *string

	// The ID of the Source Server for this launch configuration.
	SourceServerID *string

	// Whether Elastic Disaster Recovery should try to automatically choose the
	// instance type that best matches the OS, CPU, and RAM of your Source Server.
	TargetInstanceTypeRightSizingMethod types.TargetInstanceTypeRightSizingMethod

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLaunchConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetLaunchConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetLaunchConfiguration{}, middleware.After)
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
	if err = addOpGetLaunchConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLaunchConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLaunchConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "drs",
		OperationName: "GetLaunchConfiguration",
	}
}
