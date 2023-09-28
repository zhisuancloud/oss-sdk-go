// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Puts the specified workflow run properties for the given workflow run. If a
// property already exists for the specified run, then it overrides the value
// otherwise adds the property to existing properties.
func (c *Client) PutWorkflowRunProperties(ctx context.Context, params *PutWorkflowRunPropertiesInput, optFns ...func(*Options)) (*PutWorkflowRunPropertiesOutput, error) {
	if params == nil {
		params = &PutWorkflowRunPropertiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutWorkflowRunProperties", params, optFns, c.addOperationPutWorkflowRunPropertiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutWorkflowRunPropertiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutWorkflowRunPropertiesInput struct {

	// Name of the workflow which was run.
	//
	// This member is required.
	Name *string

	// The ID of the workflow run for which the run properties should be updated.
	//
	// This member is required.
	RunId *string

	// The properties to put for the specified run.
	//
	// This member is required.
	RunProperties map[string]string

	noSmithyDocumentSerde
}

type PutWorkflowRunPropertiesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutWorkflowRunPropertiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutWorkflowRunProperties{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutWorkflowRunProperties{}, middleware.After)
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
	if err = addOpPutWorkflowRunPropertiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutWorkflowRunProperties(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutWorkflowRunProperties(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "PutWorkflowRunProperties",
	}
}
