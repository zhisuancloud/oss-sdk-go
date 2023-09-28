// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/athena/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Streams the results of a single query execution specified by QueryExecutionId
// from the Athena query results location in Amazon S3. For more information, see
// Query Results (https://docs.aws.amazon.com/athena/latest/ug/querying.html) in
// the Amazon Athena User Guide. This request does not execute the query but
// returns results. Use StartQueryExecution to run a query. To stream query results
// successfully, the IAM principal with permission to call GetQueryResults also
// must have permissions to the Amazon S3 GetObject action for the Athena query
// results location. IAM principals with permission to the Amazon S3 GetObject
// action for the query results location are able to retrieve query results from
// Amazon S3 even if permission to the GetQueryResults action is denied. To
// restrict user or role access, ensure that Amazon S3 permissions to the Athena
// query location are denied.
func (c *Client) GetQueryResults(ctx context.Context, params *GetQueryResultsInput, optFns ...func(*Options)) (*GetQueryResultsOutput, error) {
	if params == nil {
		params = &GetQueryResultsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetQueryResults", params, optFns, c.addOperationGetQueryResultsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetQueryResultsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetQueryResultsInput struct {

	// The unique ID of the query execution.
	//
	// This member is required.
	QueryExecutionId *string

	// The maximum number of results (rows) to return in this request.
	MaxResults *int32

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string

	noSmithyDocumentSerde
}

type GetQueryResultsOutput struct {

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string

	// The results of the query execution.
	ResultSet *types.ResultSet

	// The number of rows inserted with a CREATE TABLE AS SELECT statement.
	UpdateCount *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetQueryResultsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetQueryResults{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetQueryResults{}, middleware.After)
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
	if err = addOpGetQueryResultsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetQueryResults(options.Region), middleware.Before); err != nil {
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

// GetQueryResultsAPIClient is a client that implements the GetQueryResults
// operation.
type GetQueryResultsAPIClient interface {
	GetQueryResults(context.Context, *GetQueryResultsInput, ...func(*Options)) (*GetQueryResultsOutput, error)
}

var _ GetQueryResultsAPIClient = (*Client)(nil)

// GetQueryResultsPaginatorOptions is the paginator options for GetQueryResults
type GetQueryResultsPaginatorOptions struct {
	// The maximum number of results (rows) to return in this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetQueryResultsPaginator is a paginator for GetQueryResults
type GetQueryResultsPaginator struct {
	options   GetQueryResultsPaginatorOptions
	client    GetQueryResultsAPIClient
	params    *GetQueryResultsInput
	nextToken *string
	firstPage bool
}

// NewGetQueryResultsPaginator returns a new GetQueryResultsPaginator
func NewGetQueryResultsPaginator(client GetQueryResultsAPIClient, params *GetQueryResultsInput, optFns ...func(*GetQueryResultsPaginatorOptions)) *GetQueryResultsPaginator {
	if params == nil {
		params = &GetQueryResultsInput{}
	}

	options := GetQueryResultsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetQueryResultsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetQueryResultsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetQueryResults page.
func (p *GetQueryResultsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetQueryResultsOutput, error) {
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

	result, err := p.client.GetQueryResults(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetQueryResults(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "GetQueryResults",
	}
}
