// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the version of the theme that the specified theme alias points to. If
// you provide a specific alias, you delete the version of the theme that the alias
// points to.
func (c *Client) DeleteThemeAlias(ctx context.Context, params *DeleteThemeAliasInput, optFns ...func(*Options)) (*DeleteThemeAliasOutput, error) {
	if params == nil {
		params = &DeleteThemeAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteThemeAlias", params, optFns, c.addOperationDeleteThemeAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteThemeAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteThemeAliasInput struct {

	// The unique name for the theme alias to delete.
	//
	// This member is required.
	AliasName *string

	// The ID of the Amazon Web Services account that contains the theme alias to
	// delete.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the theme that the specified alias is for.
	//
	// This member is required.
	ThemeId *string

	noSmithyDocumentSerde
}

type DeleteThemeAliasOutput struct {

	// The name for the theme alias.
	AliasName *string

	// The Amazon Resource Name (ARN) of the theme resource using the deleted alias.
	Arn *string

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// An ID for the theme associated with the deletion.
	ThemeId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteThemeAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteThemeAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteThemeAlias{}, middleware.After)
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
	if err = addOpDeleteThemeAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteThemeAlias(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteThemeAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DeleteThemeAlias",
	}
}
