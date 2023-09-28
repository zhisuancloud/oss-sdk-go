// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the default vocabularies for the specified Amazon Connect instance.
func (c *Client) ListDefaultVocabularies(ctx context.Context, params *ListDefaultVocabulariesInput, optFns ...func(*Options)) (*ListDefaultVocabulariesOutput, error) {
	if params == nil {
		params = &ListDefaultVocabulariesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDefaultVocabularies", params, optFns, c.addOperationListDefaultVocabulariesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDefaultVocabulariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDefaultVocabulariesInput struct {

	// The identifier of the Amazon Connect instance. You can find the instanceId in
	// the ARN of the instance.
	//
	// This member is required.
	InstanceId *string

	// The language code of the vocabulary entries. For a list of languages and their
	// corresponding language codes, see What is Amazon Transcribe?
	// (https://docs.aws.amazon.com/transcribe/latest/dg/transcribe-whatis.html)
	LanguageCode types.VocabularyLanguageCode

	// The maximum number of results to return per page.
	MaxResults int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDefaultVocabulariesOutput struct {

	// A list of default vocabularies.
	//
	// This member is required.
	DefaultVocabularyList []types.DefaultVocabulary

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDefaultVocabulariesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDefaultVocabularies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDefaultVocabularies{}, middleware.After)
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
	if err = addOpListDefaultVocabulariesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDefaultVocabularies(options.Region), middleware.Before); err != nil {
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

// ListDefaultVocabulariesAPIClient is a client that implements the
// ListDefaultVocabularies operation.
type ListDefaultVocabulariesAPIClient interface {
	ListDefaultVocabularies(context.Context, *ListDefaultVocabulariesInput, ...func(*Options)) (*ListDefaultVocabulariesOutput, error)
}

var _ ListDefaultVocabulariesAPIClient = (*Client)(nil)

// ListDefaultVocabulariesPaginatorOptions is the paginator options for
// ListDefaultVocabularies
type ListDefaultVocabulariesPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDefaultVocabulariesPaginator is a paginator for ListDefaultVocabularies
type ListDefaultVocabulariesPaginator struct {
	options   ListDefaultVocabulariesPaginatorOptions
	client    ListDefaultVocabulariesAPIClient
	params    *ListDefaultVocabulariesInput
	nextToken *string
	firstPage bool
}

// NewListDefaultVocabulariesPaginator returns a new
// ListDefaultVocabulariesPaginator
func NewListDefaultVocabulariesPaginator(client ListDefaultVocabulariesAPIClient, params *ListDefaultVocabulariesInput, optFns ...func(*ListDefaultVocabulariesPaginatorOptions)) *ListDefaultVocabulariesPaginator {
	if params == nil {
		params = &ListDefaultVocabulariesInput{}
	}

	options := ListDefaultVocabulariesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDefaultVocabulariesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDefaultVocabulariesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDefaultVocabularies page.
func (p *ListDefaultVocabulariesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDefaultVocabulariesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListDefaultVocabularies(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDefaultVocabularies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListDefaultVocabularies",
	}
}
