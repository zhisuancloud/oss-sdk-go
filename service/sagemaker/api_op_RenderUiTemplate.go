// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Renders the UI template so that you can preview the worker's experience.
func (c *Client) RenderUiTemplate(ctx context.Context, params *RenderUiTemplateInput, optFns ...func(*Options)) (*RenderUiTemplateOutput, error) {
	if params == nil {
		params = &RenderUiTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RenderUiTemplate", params, optFns, c.addOperationRenderUiTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RenderUiTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RenderUiTemplateInput struct {

	// The Amazon Resource Name (ARN) that has access to the S3 objects that are used
	// by the template.
	//
	// This member is required.
	RoleArn *string

	// A RenderableTask object containing a representative task to render.
	//
	// This member is required.
	Task *types.RenderableTask

	// The HumanTaskUiArn of the worker UI that you want to render. Do not provide a
	// HumanTaskUiArn if you use the UiTemplate parameter. See a list of available
	// Human Ui Amazon Resource Names (ARNs) in UiConfig.
	HumanTaskUiArn *string

	// A Template object containing the worker UI template to render.
	UiTemplate *types.UiTemplate

	noSmithyDocumentSerde
}

type RenderUiTemplateOutput struct {

	// A list of one or more RenderingError objects if any were encountered while
	// rendering the template. If there were no errors, the list is empty.
	//
	// This member is required.
	Errors []types.RenderingError

	// A Liquid template that renders the HTML for the worker UI.
	//
	// This member is required.
	RenderedContent *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRenderUiTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRenderUiTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRenderUiTemplate{}, middleware.After)
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
	if err = addOpRenderUiTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRenderUiTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRenderUiTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "RenderUiTemplate",
	}
}
