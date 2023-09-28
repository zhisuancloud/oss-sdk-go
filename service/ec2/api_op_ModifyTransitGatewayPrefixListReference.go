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

// Modifies a reference (route) to a prefix list in a specified transit gateway
// route table.
func (c *Client) ModifyTransitGatewayPrefixListReference(ctx context.Context, params *ModifyTransitGatewayPrefixListReferenceInput, optFns ...func(*Options)) (*ModifyTransitGatewayPrefixListReferenceOutput, error) {
	if params == nil {
		params = &ModifyTransitGatewayPrefixListReferenceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyTransitGatewayPrefixListReference", params, optFns, c.addOperationModifyTransitGatewayPrefixListReferenceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyTransitGatewayPrefixListReferenceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyTransitGatewayPrefixListReferenceInput struct {

	// The ID of the prefix list.
	//
	// This member is required.
	PrefixListId *string

	// The ID of the transit gateway route table.
	//
	// This member is required.
	TransitGatewayRouteTableId *string

	// Indicates whether to drop traffic that matches this route.
	Blackhole *bool

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The ID of the attachment to which traffic is routed.
	TransitGatewayAttachmentId *string

	noSmithyDocumentSerde
}

type ModifyTransitGatewayPrefixListReferenceOutput struct {

	// Information about the prefix list reference.
	TransitGatewayPrefixListReference *types.TransitGatewayPrefixListReference

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyTransitGatewayPrefixListReferenceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyTransitGatewayPrefixListReference{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyTransitGatewayPrefixListReference{}, middleware.After)
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
	if err = addOpModifyTransitGatewayPrefixListReferenceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyTransitGatewayPrefixListReference(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyTransitGatewayPrefixListReference(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyTransitGatewayPrefixListReference",
	}
}
