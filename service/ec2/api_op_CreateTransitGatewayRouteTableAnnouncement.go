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

// Advertises a new transit gateway route table.
func (c *Client) CreateTransitGatewayRouteTableAnnouncement(ctx context.Context, params *CreateTransitGatewayRouteTableAnnouncementInput, optFns ...func(*Options)) (*CreateTransitGatewayRouteTableAnnouncementOutput, error) {
	if params == nil {
		params = &CreateTransitGatewayRouteTableAnnouncementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTransitGatewayRouteTableAnnouncement", params, optFns, c.addOperationCreateTransitGatewayRouteTableAnnouncementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTransitGatewayRouteTableAnnouncementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTransitGatewayRouteTableAnnouncementInput struct {

	// The ID of the peering attachment.
	//
	// This member is required.
	PeeringAttachmentId *string

	// The ID of the transit gateway route table.
	//
	// This member is required.
	TransitGatewayRouteTableId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The tags specifications applied to the transit gateway route table announcement.
	TagSpecifications []types.TagSpecification

	noSmithyDocumentSerde
}

type CreateTransitGatewayRouteTableAnnouncementOutput struct {

	// Provides details about the transit gateway route table announcement.
	TransitGatewayRouteTableAnnouncement *types.TransitGatewayRouteTableAnnouncement

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTransitGatewayRouteTableAnnouncementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateTransitGatewayRouteTableAnnouncement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateTransitGatewayRouteTableAnnouncement{}, middleware.After)
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
	if err = addOpCreateTransitGatewayRouteTableAnnouncementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTransitGatewayRouteTableAnnouncement(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTransitGatewayRouteTableAnnouncement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateTransitGatewayRouteTableAnnouncement",
	}
}
