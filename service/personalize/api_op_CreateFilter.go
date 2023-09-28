// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a recommendation filter. For more information, see Filtering
// recommendations and user segments
// (https://docs.aws.amazon.com/personalize/latest/dg/filter.html).
func (c *Client) CreateFilter(ctx context.Context, params *CreateFilterInput, optFns ...func(*Options)) (*CreateFilterOutput, error) {
	if params == nil {
		params = &CreateFilterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFilter", params, optFns, c.addOperationCreateFilterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFilterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFilterInput struct {

	// The ARN of the dataset group that the filter will belong to.
	//
	// This member is required.
	DatasetGroupArn *string

	// The filter expression defines which items are included or excluded from
	// recommendations. Filter expression must follow specific format rules. For
	// information about filter expression structure and syntax, see Filter expressions
	// (https://docs.aws.amazon.com/personalize/latest/dg/filter-expressions.html).
	//
	// This member is required.
	FilterExpression *string

	// The name of the filter to create.
	//
	// This member is required.
	Name *string

	// A list of tags
	// (https://docs.aws.amazon.com/personalize/latest/dev/tagging-resources.html) to
	// apply to the filter.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateFilterOutput struct {

	// The ARN of the new filter.
	FilterArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFilterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateFilter{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateFilter{}, middleware.After)
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
	if err = addOpCreateFilterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFilter(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFilter(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "CreateFilter",
	}
}
