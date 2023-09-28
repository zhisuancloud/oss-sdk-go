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

// Describes the specified attribute of the specified AMI. You can specify only one
// attribute at a time.
func (c *Client) DescribeImageAttribute(ctx context.Context, params *DescribeImageAttributeInput, optFns ...func(*Options)) (*DescribeImageAttributeOutput, error) {
	if params == nil {
		params = &DescribeImageAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeImageAttribute", params, optFns, c.addOperationDescribeImageAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeImageAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeImageAttribute.
type DescribeImageAttributeInput struct {

	// The AMI attribute. Note: The blockDeviceMapping attribute is deprecated. Using
	// this attribute returns the Client.AuthFailure error. To get information about
	// the block device mappings for an AMI, use the DescribeImages action.
	//
	// This member is required.
	Attribute types.ImageAttributeName

	// The ID of the AMI.
	//
	// This member is required.
	ImageId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	noSmithyDocumentSerde
}

// Describes an image attribute.
type DescribeImageAttributeOutput struct {

	// The block device mapping entries.
	BlockDeviceMappings []types.BlockDeviceMapping

	// The boot mode.
	BootMode *types.AttributeValue

	// A description for the AMI.
	Description *types.AttributeValue

	// The ID of the AMI.
	ImageId *string

	// The kernel ID.
	KernelId *types.AttributeValue

	// The date and time, in ISO 8601 date-time format
	// (http://www.iso.org/iso/iso8601), when the AMI was last used to launch an EC2
	// instance. When the AMI is used to launch an instance, there is a 24-hour delay
	// before that usage is reported. lastLaunchedTime data is available starting April
	// 2017.
	LastLaunchedTime *types.AttributeValue

	// The launch permissions.
	LaunchPermissions []types.LaunchPermission

	// The product codes.
	ProductCodes []types.ProductCode

	// The RAM disk ID.
	RamdiskId *types.AttributeValue

	// Indicates whether enhanced networking with the Intel 82599 Virtual Function
	// interface is enabled.
	SriovNetSupport *types.AttributeValue

	// If the image is configured for NitroTPM support, the value is v2.0.
	TpmSupport *types.AttributeValue

	// Base64 representation of the non-volatile UEFI variable store. To retrieve the
	// UEFI data, use the GetInstanceUefiData
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetInstanceUefiData)
	// command. You can inspect and modify the UEFI data by using the python-uefivars
	// tool (https://github.com/awslabs/python-uefivars) on GitHub. For more
	// information, see UEFI Secure Boot
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/uefi-secure-boot.html) in
	// the Amazon Elastic Compute Cloud User Guide.
	UefiData *types.AttributeValue

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeImageAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeImageAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeImageAttribute{}, middleware.After)
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
	if err = addOpDescribeImageAttributeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeImageAttribute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeImageAttribute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeImageAttribute",
	}
}