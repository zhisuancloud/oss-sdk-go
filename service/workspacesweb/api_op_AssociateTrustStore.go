// Code generated by smithy-go-codegen DO NOT EDIT.

package workspacesweb

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a trust store with a web portal.
func (c *Client) AssociateTrustStore(ctx context.Context, params *AssociateTrustStoreInput, optFns ...func(*Options)) (*AssociateTrustStoreOutput, error) {
	if params == nil {
		params = &AssociateTrustStoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateTrustStore", params, optFns, c.addOperationAssociateTrustStoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateTrustStoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateTrustStoreInput struct {

	// The ARN of the web portal.
	//
	// This member is required.
	PortalArn *string

	// The ARN of the trust store.
	//
	// This member is required.
	TrustStoreArn *string

	noSmithyDocumentSerde
}

type AssociateTrustStoreOutput struct {

	// The ARN of the web portal.
	//
	// This member is required.
	PortalArn *string

	// The ARN of the trust store.
	//
	// This member is required.
	TrustStoreArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateTrustStoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateTrustStore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateTrustStore{}, middleware.After)
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
	if err = addOpAssociateTrustStoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateTrustStore(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateTrustStore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces-web",
		OperationName: "AssociateTrustStore",
	}
}
