// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates properties of a wireless gateway.
func (c *Client) UpdateWirelessGateway(ctx context.Context, params *UpdateWirelessGatewayInput, optFns ...func(*Options)) (*UpdateWirelessGatewayOutput, error) {
	if params == nil {
		params = &UpdateWirelessGatewayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateWirelessGateway", params, optFns, c.addOperationUpdateWirelessGatewayMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateWirelessGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateWirelessGatewayInput struct {

	// The ID of the resource to update.
	//
	// This member is required.
	Id *string

	// A new description of the resource.
	Description *string

	// A list of JoinEuiRange used by LoRa gateways to filter LoRa frames.
	JoinEuiFilters [][]string

	// The new name of the resource.
	Name *string

	// A list of NetId values that are used by LoRa gateways to filter the uplink
	// frames.
	NetIdFilters []string

	noSmithyDocumentSerde
}

type UpdateWirelessGatewayOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateWirelessGatewayMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateWirelessGateway{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateWirelessGateway{}, middleware.After)
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
	if err = addOpUpdateWirelessGatewayValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateWirelessGateway(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateWirelessGateway(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "UpdateWirelessGateway",
	}
}