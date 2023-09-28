// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/devicefarm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates information about an Amazon Virtual Private Cloud (VPC) endpoint
// configuration.
func (c *Client) UpdateVPCEConfiguration(ctx context.Context, params *UpdateVPCEConfigurationInput, optFns ...func(*Options)) (*UpdateVPCEConfigurationOutput, error) {
	if params == nil {
		params = &UpdateVPCEConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateVPCEConfiguration", params, optFns, c.addOperationUpdateVPCEConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateVPCEConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateVPCEConfigurationInput struct {

	// The Amazon Resource Name (ARN) of the VPC endpoint configuration you want to
	// update.
	//
	// This member is required.
	Arn *string

	// The DNS (domain) name used to connect to your private service in your VPC. The
	// DNS name must not already be in use on the internet.
	ServiceDnsName *string

	// An optional description that provides details about your VPC endpoint
	// configuration.
	VpceConfigurationDescription *string

	// The friendly name you give to your VPC endpoint configuration to manage your
	// configurations more easily.
	VpceConfigurationName *string

	// The name of the VPC endpoint service running in your AWS account that you want
	// Device Farm to test.
	VpceServiceName *string

	noSmithyDocumentSerde
}

type UpdateVPCEConfigurationOutput struct {

	// An object that contains information about your VPC endpoint configuration.
	VpceConfiguration *types.VPCEConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateVPCEConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateVPCEConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateVPCEConfiguration{}, middleware.After)
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
	if err = addOpUpdateVPCEConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateVPCEConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateVPCEConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "UpdateVPCEConfiguration",
	}
}
