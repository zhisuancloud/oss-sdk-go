// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a response headers policy. A response headers policy contains
// information about a set of HTTP response headers and their values. To create a
// response headers policy, you provide some metadata about the policy, and a set
// of configurations that specify the response headers. After you create a response
// headers policy, you can use its ID to attach it to one or more cache behaviors
// in a CloudFront distribution. When it’s attached to a cache behavior, CloudFront
// adds the headers in the policy to HTTP responses that it sends for requests that
// match the cache behavior.
func (c *Client) CreateResponseHeadersPolicy(ctx context.Context, params *CreateResponseHeadersPolicyInput, optFns ...func(*Options)) (*CreateResponseHeadersPolicyOutput, error) {
	if params == nil {
		params = &CreateResponseHeadersPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateResponseHeadersPolicy", params, optFns, c.addOperationCreateResponseHeadersPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateResponseHeadersPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateResponseHeadersPolicyInput struct {

	// Contains metadata about the response headers policy, and a set of configurations
	// that specify the response headers.
	//
	// This member is required.
	ResponseHeadersPolicyConfig *types.ResponseHeadersPolicyConfig

	noSmithyDocumentSerde
}

type CreateResponseHeadersPolicyOutput struct {

	// The version identifier for the current version of the response headers policy.
	ETag *string

	// The URL of the response headers policy.
	Location *string

	// Contains a response headers policy.
	ResponseHeadersPolicy *types.ResponseHeadersPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateResponseHeadersPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateResponseHeadersPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateResponseHeadersPolicy{}, middleware.After)
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
	if err = addOpCreateResponseHeadersPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateResponseHeadersPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateResponseHeadersPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "CreateResponseHeadersPolicy",
	}
}
