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

// Creates a default subnet with a size /20 IPv4 CIDR block in the specified
// Availability Zone in your default VPC. You can have only one default subnet per
// Availability Zone. For more information, see Creating a default subnet
// (https://docs.aws.amazon.com/vpc/latest/userguide/default-vpc.html#create-default-subnet)
// in the Amazon Virtual Private Cloud User Guide.
func (c *Client) CreateDefaultSubnet(ctx context.Context, params *CreateDefaultSubnetInput, optFns ...func(*Options)) (*CreateDefaultSubnetOutput, error) {
	if params == nil {
		params = &CreateDefaultSubnetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDefaultSubnet", params, optFns, c.addOperationCreateDefaultSubnetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDefaultSubnetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDefaultSubnetInput struct {

	// The Availability Zone in which to create the default subnet.
	//
	// This member is required.
	AvailabilityZone *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// Indicates whether to create an IPv6 only subnet. If you already have a default
	// subnet for this Availability Zone, you must delete it before you can create an
	// IPv6 only subnet.
	Ipv6Native *bool

	noSmithyDocumentSerde
}

type CreateDefaultSubnetOutput struct {

	// Information about the subnet.
	Subnet *types.Subnet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDefaultSubnetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateDefaultSubnet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateDefaultSubnet{}, middleware.After)
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
	if err = addOpCreateDefaultSubnetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDefaultSubnet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDefaultSubnet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateDefaultSubnet",
	}
}
