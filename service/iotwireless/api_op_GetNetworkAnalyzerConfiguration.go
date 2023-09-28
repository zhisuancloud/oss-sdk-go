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

// Get network analyzer configuration.
func (c *Client) GetNetworkAnalyzerConfiguration(ctx context.Context, params *GetNetworkAnalyzerConfigurationInput, optFns ...func(*Options)) (*GetNetworkAnalyzerConfigurationOutput, error) {
	if params == nil {
		params = &GetNetworkAnalyzerConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetNetworkAnalyzerConfiguration", params, optFns, c.addOperationGetNetworkAnalyzerConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetNetworkAnalyzerConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetNetworkAnalyzerConfigurationInput struct {

	// Name of the network analyzer configuration.
	//
	// This member is required.
	ConfigurationName *string

	noSmithyDocumentSerde
}

type GetNetworkAnalyzerConfigurationOutput struct {

	// The Amazon Resource Name of the new resource.
	Arn *string

	// The description of the new resource.
	Description *string

	// Name of the network analyzer configuration.
	Name *string

	// Trace content for your wireless gateway and wireless device resources.
	TraceContent *types.TraceContent

	// List of wireless gateway resources that have been added to the network analyzer
	// configuration.
	WirelessDevices []string

	// List of wireless gateway resources that have been added to the network analyzer
	// configuration.
	WirelessGateways []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetNetworkAnalyzerConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetNetworkAnalyzerConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetNetworkAnalyzerConfiguration{}, middleware.After)
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
	if err = addOpGetNetworkAnalyzerConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetNetworkAnalyzerConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetNetworkAnalyzerConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "GetNetworkAnalyzerConfiguration",
	}
}