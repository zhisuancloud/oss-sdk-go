// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/directconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a transit virtual interface. A transit virtual interface should be used
// to access one or more transit gateways associated with Direct Connect gateways.
// A transit virtual interface enables the connection of multiple VPCs attached to
// a transit gateway to a Direct Connect gateway. If you associate your transit
// gateway with one or more Direct Connect gateways, the Autonomous System Number
// (ASN) used by the transit gateway and the Direct Connect gateway must be
// different. For example, if you use the default ASN 64512 for both your the
// transit gateway and Direct Connect gateway, the association request fails.
// Setting the MTU of a virtual interface to 8500 (jumbo frames) can cause an
// update to the underlying physical connection if it wasn't updated to support
// jumbo frames. Updating the connection disrupts network connectivity for all
// virtual interfaces associated with the connection for up to 30 seconds. To check
// whether your connection supports jumbo frames, call DescribeConnections. To
// check whether your virtual interface supports jumbo frames, call
// DescribeVirtualInterfaces.
func (c *Client) CreateTransitVirtualInterface(ctx context.Context, params *CreateTransitVirtualInterfaceInput, optFns ...func(*Options)) (*CreateTransitVirtualInterfaceOutput, error) {
	if params == nil {
		params = &CreateTransitVirtualInterfaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTransitVirtualInterface", params, optFns, c.addOperationCreateTransitVirtualInterfaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTransitVirtualInterfaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTransitVirtualInterfaceInput struct {

	// The ID of the connection.
	//
	// This member is required.
	ConnectionId *string

	// Information about the transit virtual interface.
	//
	// This member is required.
	NewTransitVirtualInterface *types.NewTransitVirtualInterface

	noSmithyDocumentSerde
}

type CreateTransitVirtualInterfaceOutput struct {

	// Information about a virtual interface.
	VirtualInterface *types.VirtualInterface

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTransitVirtualInterfaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTransitVirtualInterface{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTransitVirtualInterface{}, middleware.After)
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
	if err = addOpCreateTransitVirtualInterfaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTransitVirtualInterface(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTransitVirtualInterface(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "CreateTransitVirtualInterface",
	}
}
