// Code generated by smithy-go-codegen DO NOT EDIT.

package billingconductor

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/billingconductor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the custom line item identified by the given ARN in the current, or
// previous billing period.
func (c *Client) DeleteCustomLineItem(ctx context.Context, params *DeleteCustomLineItemInput, optFns ...func(*Options)) (*DeleteCustomLineItemOutput, error) {
	if params == nil {
		params = &DeleteCustomLineItemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCustomLineItem", params, optFns, c.addOperationDeleteCustomLineItemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteCustomLineItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteCustomLineItemInput struct {

	// The ARN of the custom line item to be deleted.
	//
	// This member is required.
	Arn *string

	// The billing period range in which the custom line item request will be applied.
	BillingPeriodRange *types.CustomLineItemBillingPeriodRange

	noSmithyDocumentSerde
}

type DeleteCustomLineItemOutput struct {

	// Then ARN of the deleted custom line item.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteCustomLineItemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteCustomLineItem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteCustomLineItem{}, middleware.After)
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
	if err = addOpDeleteCustomLineItemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCustomLineItem(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteCustomLineItem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "billingconductor",
		OperationName: "DeleteCustomLineItem",
	}
}