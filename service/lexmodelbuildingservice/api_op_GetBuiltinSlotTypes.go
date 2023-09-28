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

// Gets a list of built-in slot types that meet the specified criteria. For a list
// of built-in slot types, see Slot Type Reference
// (https://developer.amazon.com/public/solutions/alexa/alexa-skills-kit/docs/built-in-intent-ref/slot-type-reference)
// in the Alexa Skills Kit. This operation requires permission for the
// lex:GetBuiltInSlotTypes action.
func (c *Client) GetBuiltinSlotTypes(ctx context.Context, params *GetBuiltinSlotTypesInput, optFns ...func(*Options)) (*GetBuiltinSlotTypesOutput, error) {
	if params == nil {
		params = &GetBuiltinSlotTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBuiltinSlotTypes", params, optFns, c.addOperationGetBuiltinSlotTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBuiltinSlotTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBuiltinSlotTypesInput struct {

	// A list of locales that the slot type supports.
	Locale types.Locale

	// The maximum number of slot types to return in the response. The default is 10.
	MaxResults *int32

	// A pagination token that fetches the next page of slot types. If the response to
	// this API call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of slot types, specify the pagination token in
	// the next request.
	NextToken *string

	// Substring to match in built-in slot type signatures. A slot type will be
	// returned if any part of its signature matches the substring. For example, "xyz"
	// matches both "xyzabc" and "abcxyz."
	SignatureContains *string

	noSmithyDocumentSerde
}

type GetBuiltinSlotTypesOutput struct {

	// If the response is truncated, the response includes a pagination token that you
	// can use in your next request to fetch the next page of slot types.
	NextToken *string

	// An array of BuiltInSlotTypeMetadata objects, one entry for each slot type
	// returned.
	SlotTypes []types.BuiltinSlotTypeMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBuiltinSlotTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBuiltinSlotTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBuiltinSlotTypes{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBuiltinSlotTypes(options.Region), middleware.Before); err != nil {
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

// GetBuiltinSlotTypesAPIClient is a client that implements the GetBuiltinSlotTypes
// operation.
type GetBuiltinSlotTypesAPIClient interface {
	GetBuiltinSlotTypes(context.Context, *GetBuiltinSlotTypesInput, ...func(*Options)) (*GetBuiltinSlotTypesOutput, error)
}

var _ GetBuiltinSlotTypesAPIClient = (*Client)(nil)

// GetBuiltinSlotTypesPaginatorOptions is the paginator options for
// GetBuiltinSlotTypes
type GetBuiltinSlotTypesPaginatorOptions struct {
	// The maximum number of slot types to return in the response. The default is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetBuiltinSlotTypesPaginator is a paginator for GetBuiltinSlotTypes
type GetBuiltinSlotTypesPaginator struct {
	options   GetBuiltinSlotTypesPaginatorOptions
	client    GetBuiltinSlotTypesAPIClient
	params    *GetBuiltinSlotTypesInput
	nextToken *string
	firstPage bool
}

// NewGetBuiltinSlotTypesPaginator returns a new GetBuiltinSlotTypesPaginator
func NewGetBuiltinSlotTypesPaginator(client GetBuiltinSlotTypesAPIClient, params *GetBuiltinSlotTypesInput, optFns ...func(*GetBuiltinSlotTypesPaginatorOptions)) *GetBuiltinSlotTypesPaginator {
	if params == nil {
		params = &GetBuiltinSlotTypesInput{}
	}

	options := GetBuiltinSlotTypesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetBuiltinSlotTypesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetBuiltinSlotTypesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetBuiltinSlotTypes page.
func (p *GetBuiltinSlotTypesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetBuiltinSlotTypesOutput, error) {
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

	result, err := p.client.GetBuiltinSlotTypes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetBuiltinSlotTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetBuiltinSlotTypes",
	}
}
