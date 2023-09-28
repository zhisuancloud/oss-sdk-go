// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/servicequotas/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about the specified quota increase request in your quota
// request template.
func (c *Client) GetServiceQuotaIncreaseRequestFromTemplate(ctx context.Context, params *GetServiceQuotaIncreaseRequestFromTemplateInput, optFns ...func(*Options)) (*GetServiceQuotaIncreaseRequestFromTemplateOutput, error) {
	if params == nil {
		params = &GetServiceQuotaIncreaseRequestFromTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetServiceQuotaIncreaseRequestFromTemplate", params, optFns, c.addOperationGetServiceQuotaIncreaseRequestFromTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetServiceQuotaIncreaseRequestFromTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServiceQuotaIncreaseRequestFromTemplateInput struct {

	// The AWS Region.
	//
	// This member is required.
	AwsRegion *string

	// The quota identifier.
	//
	// This member is required.
	QuotaCode *string

	// The service identifier.
	//
	// This member is required.
	ServiceCode *string

	noSmithyDocumentSerde
}

type GetServiceQuotaIncreaseRequestFromTemplateOutput struct {

	// Information about the quota increase request.
	ServiceQuotaIncreaseRequestInTemplate *types.ServiceQuotaIncreaseRequestInTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetServiceQuotaIncreaseRequestFromTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetServiceQuotaIncreaseRequestFromTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetServiceQuotaIncreaseRequestFromTemplate{}, middleware.After)
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
	if err = addOpGetServiceQuotaIncreaseRequestFromTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetServiceQuotaIncreaseRequestFromTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetServiceQuotaIncreaseRequestFromTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "GetServiceQuotaIncreaseRequestFromTemplate",
	}
}
