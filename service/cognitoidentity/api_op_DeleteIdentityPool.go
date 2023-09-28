// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentity

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes an identity pool. Once a pool is deleted, users will not be able to
// authenticate with the pool. You must use AWS Developer credentials to call this
// API.
func (c *Client) DeleteIdentityPool(ctx context.Context, params *DeleteIdentityPoolInput, optFns ...func(*Options)) (*DeleteIdentityPoolOutput, error) {
	if params == nil {
		params = &DeleteIdentityPoolInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteIdentityPool", params, optFns, c.addOperationDeleteIdentityPoolMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteIdentityPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to the DeleteIdentityPool action.
type DeleteIdentityPoolInput struct {

	// An identity pool ID in the format REGION:GUID.
	//
	// This member is required.
	IdentityPoolId *string

	noSmithyDocumentSerde
}

type DeleteIdentityPoolOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteIdentityPoolMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteIdentityPool{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteIdentityPool{}, middleware.After)
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
	if err = addOpDeleteIdentityPoolValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteIdentityPool(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteIdentityPool(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-identity",
		OperationName: "DeleteIdentityPool",
	}
}
