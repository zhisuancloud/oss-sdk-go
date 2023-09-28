// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets operating information about a wireless device.
func (c *Client) GetWirelessDeviceStatistics(ctx context.Context, params *GetWirelessDeviceStatisticsInput, optFns ...func(*Options)) (*GetWirelessDeviceStatisticsOutput, error) {
	if params == nil {
		params = &GetWirelessDeviceStatisticsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWirelessDeviceStatistics", params, optFns, c.addOperationGetWirelessDeviceStatisticsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWirelessDeviceStatisticsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWirelessDeviceStatisticsInput struct {

	// The ID of the wireless device for which to get the data.
	//
	// This member is required.
	WirelessDeviceId *string

	noSmithyDocumentSerde
}

type GetWirelessDeviceStatisticsOutput struct {

	// The date and time when the most recent uplink was received.
	LastUplinkReceivedAt *string

	// Information about the wireless device's operations.
	LoRaWAN *types.LoRaWANDeviceMetadata

	// MetaData for Sidewalk device.
	Sidewalk *types.SidewalkDeviceMetadata

	// The ID of the wireless device.
	WirelessDeviceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetWirelessDeviceStatisticsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetWirelessDeviceStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetWirelessDeviceStatistics{}, middleware.After)
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
	if err = addOpGetWirelessDeviceStatisticsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetWirelessDeviceStatistics(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetWirelessDeviceStatistics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "GetWirelessDeviceStatistics",
	}
}
