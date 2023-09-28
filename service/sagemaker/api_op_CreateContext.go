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

// Creates a context. A context is a lineage tracking entity that represents a
// logical grouping of other tracking or experiment entities. Some examples are an
// endpoint and a model package. For more information, see Amazon SageMaker ML
// Lineage Tracking
// (https://docs.aws.amazon.com/sagemaker/latest/dg/lineage-tracking.html).
func (c *Client) CreateContext(ctx context.Context, params *CreateContextInput, optFns ...func(*Options)) (*CreateContextOutput, error) {
	if params == nil {
		params = &CreateContextInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateContext", params, optFns, c.addOperationCreateContextMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateContextOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateContextInput struct {

	// The name of the context. Must be unique to your account in an Amazon Web
	// Services Region.
	//
	// This member is required.
	ContextName *string

	// The context type.
	//
	// This member is required.
	ContextType *string

	// The source type, ID, and URI.
	//
	// This member is required.
	Source *types.ContextSource

	// The description of the context.
	Description *string

	// A list of properties to add to the context.
	Properties map[string]string

	// A list of tags to apply to the context.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateContextOutput struct {

	// The Amazon Resource Name (ARN) of the context.
	ContextArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateContextMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateContext{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateContext{}, middleware.After)
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
	if err = addOpCreateContextValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateContext(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateContext(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateContext",
	}
}
