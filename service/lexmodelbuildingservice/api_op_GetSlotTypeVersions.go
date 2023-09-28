// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/lexmodelbuildingservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about all versions of a slot type. The GetSlotTypeVersions
// operation returns a SlotTypeMetadata object for each version of a slot type. For
// example, if a slot type has three numbered versions, the GetSlotTypeVersions
// operation returns four SlotTypeMetadata objects in the response, one for each
// numbered version and one for the $LATEST version. The GetSlotTypeVersions
// operation always returns at least one version, the $LATEST version. This
// operation requires permissions for the lex:GetSlotTypeVersions action.
func (c *Client) GetSlotTypeVersions(ctx context.Context, params *GetSlotTypeVersionsInput, optFns ...func(*Options)) (*GetSlotTypeVersionsOutput, error) {
	if params == nil {
		params = &GetSlotTypeVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSlotTypeVersions", params, optFns, c.addOperationGetSlotTypeVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSlotTypeVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSlotTypeVersionsInput struct {

	// The name of the slot type for which versions should be returned.
	//
	// This member is required.
	Name *string

	// The maximum number of slot type versions to return in the response. The default
	// is 10.
	MaxResults *int32

	// A pagination token for fetching the next page of slot type versions. If the
	// response to this call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of versions, specify the pagination token in
	// the next request.
	NextToken *string

	noSmithyDocumentSerde
}

type GetSlotTypeVersionsOutput struct {

	// A pagination token for fetching the next page of slot type versions. If the
	// response to this call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of versions, specify the pagination token in
	// the next request.
	NextToken *string

	// An array of SlotTypeMetadata objects, one for each numbered version of the slot
	// type plus one for the $LATEST version.
	SlotTypes []types.SlotTypeMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSlotTypeVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSlotTypeVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSlotTypeVersions{}, middleware.After)
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
	if err = addOpGetSlotTypeVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSlotTypeVersions(options.Region), middleware.Before); err != nil {
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

// GetSlotTypeVersionsAPIClient is a client that implements the GetSlotTypeVersions
// operation.
type GetSlotTypeVersionsAPIClient interface {
	GetSlotTypeVersions(context.Context, *GetSlotTypeVersionsInput, ...func(*Options)) (*GetSlotTypeVersionsOutput, error)
}

var _ GetSlotTypeVersionsAPIClient = (*Client)(nil)

// GetSlotTypeVersionsPaginatorOptions is the paginator options for
// GetSlotTypeVersions
type GetSlotTypeVersionsPaginatorOptions struct {
	// The maximum number of slot type versions to return in the response. The default
	// is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetSlotTypeVersionsPaginator is a paginator for GetSlotTypeVersions
type GetSlotTypeVersionsPaginator struct {
	options   GetSlotTypeVersionsPaginatorOptions
	client    GetSlotTypeVersionsAPIClient
	params    *GetSlotTypeVersionsInput
	nextToken *string
	firstPage bool
}

// NewGetSlotTypeVersionsPaginator returns a new GetSlotTypeVersionsPaginator
func NewGetSlotTypeVersionsPaginator(client GetSlotTypeVersionsAPIClient, params *GetSlotTypeVersionsInput, optFns ...func(*GetSlotTypeVersionsPaginatorOptions)) *GetSlotTypeVersionsPaginator {
	if params == nil {
		params = &GetSlotTypeVersionsInput{}
	}

	options := GetSlotTypeVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetSlotTypeVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetSlotTypeVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetSlotTypeVersions page.
func (p *GetSlotTypeVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetSlotTypeVersionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.GetSlotTypeVersions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetSlotTypeVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetSlotTypeVersions",
	}
}
