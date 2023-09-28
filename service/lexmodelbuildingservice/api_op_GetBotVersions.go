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

// Gets information about all of the versions of a bot. The GetBotVersions
// operation returns a BotMetadata object for each version of a bot. For example,
// if a bot has three numbered versions, the GetBotVersions operation returns four
// BotMetadata objects in the response, one for each numbered version and one for
// the $LATEST version. The GetBotVersions operation always returns at least one
// version, the $LATEST version. This operation requires permissions for the
// lex:GetBotVersions action.
func (c *Client) GetBotVersions(ctx context.Context, params *GetBotVersionsInput, optFns ...func(*Options)) (*GetBotVersionsOutput, error) {
	if params == nil {
		params = &GetBotVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBotVersions", params, optFns, c.addOperationGetBotVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBotVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBotVersionsInput struct {

	// The name of the bot for which versions should be returned.
	//
	// This member is required.
	Name *string

	// The maximum number of bot versions to return in the response. The default is 10.
	MaxResults *int32

	// A pagination token for fetching the next page of bot versions. If the response
	// to this call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of versions, specify the pagination token in
	// the next request.
	NextToken *string

	noSmithyDocumentSerde
}

type GetBotVersionsOutput struct {

	// An array of BotMetadata objects, one for each numbered version of the bot plus
	// one for the $LATEST version.
	Bots []types.BotMetadata

	// A pagination token for fetching the next page of bot versions. If the response
	// to this call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of versions, specify the pagination token in
	// the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBotVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBotVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBotVersions{}, middleware.After)
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
	if err = addOpGetBotVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBotVersions(options.Region), middleware.Before); err != nil {
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

// GetBotVersionsAPIClient is a client that implements the GetBotVersions
// operation.
type GetBotVersionsAPIClient interface {
	GetBotVersions(context.Context, *GetBotVersionsInput, ...func(*Options)) (*GetBotVersionsOutput, error)
}

var _ GetBotVersionsAPIClient = (*Client)(nil)

// GetBotVersionsPaginatorOptions is the paginator options for GetBotVersions
type GetBotVersionsPaginatorOptions struct {
	// The maximum number of bot versions to return in the response. The default is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetBotVersionsPaginator is a paginator for GetBotVersions
type GetBotVersionsPaginator struct {
	options   GetBotVersionsPaginatorOptions
	client    GetBotVersionsAPIClient
	params    *GetBotVersionsInput
	nextToken *string
	firstPage bool
}

// NewGetBotVersionsPaginator returns a new GetBotVersionsPaginator
func NewGetBotVersionsPaginator(client GetBotVersionsAPIClient, params *GetBotVersionsInput, optFns ...func(*GetBotVersionsPaginatorOptions)) *GetBotVersionsPaginator {
	if params == nil {
		params = &GetBotVersionsInput{}
	}

	options := GetBotVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetBotVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetBotVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetBotVersions page.
func (p *GetBotVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetBotVersionsOutput, error) {
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

	result, err := p.client.GetBotVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetBotVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetBotVersions",
	}
}
