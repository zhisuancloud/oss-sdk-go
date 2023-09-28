// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associate a multicast group with a FUOTA task.
func (c *Client) AssociateMulticastGroupWithFuotaTask(ctx context.Context, params *AssociateMulticastGroupWithFuotaTaskInput, optFns ...func(*Options)) (*AssociateMulticastGroupWithFuotaTaskOutput, error) {
	if params == nil {
		params = &AssociateMulticastGroupWithFuotaTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateMulticastGroupWithFuotaTask", params, optFns, c.addOperationAssociateMulticastGroupWithFuotaTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateMulticastGroupWithFuotaTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateMulticastGroupWithFuotaTaskInput struct {

	// The ID of a FUOTA task.
	//
	// This member is required.
	Id *string

	// The ID of the multicast group.
	//
	// This member is required.
	MulticastGroupId *string

	noSmithyDocumentSerde
}

type AssociateMulticastGroupWithFuotaTaskOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateMulticastGroupWithFuotaTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateMulticastGroupWithFuotaTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateMulticastGroupWithFuotaTask{}, middleware.After)
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
	if err = addOpAssociateMulticastGroupWithFuotaTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateMulticastGroupWithFuotaTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateMulticastGroupWithFuotaTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "AssociateMulticastGroupWithFuotaTask",
	}
}
