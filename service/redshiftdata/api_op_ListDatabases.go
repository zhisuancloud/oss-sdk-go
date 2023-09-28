// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftdata

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the databases in a cluster. A token is returned to page through the
// database list. Depending on the authorization method, use one of the following
// combinations of request parameters:
//
// * Secrets Manager - when connecting to a
// cluster, specify the Amazon Resource Name (ARN) of the secret, the database
// name, and the cluster identifier that matches the cluster in the secret. When
// connecting to a serverless workgroup, specify the Amazon Resource Name (ARN) of
// the secret and the database name.
//
// * Temporary credentials - when connecting to
// a cluster, specify the cluster identifier, the database name, and the database
// user name. Also, permission to call the redshift:GetClusterCredentials operation
// is required. When connecting to a serverless workgroup, specify the workgroup
// name and database name. Also, permission to call the
// redshift-serverless:GetCredentials operation is required.
func (c *Client) ListDatabases(ctx context.Context, params *ListDatabasesInput, optFns ...func(*Options)) (*ListDatabasesOutput, error) {
	if params == nil {
		params = &ListDatabasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDatabases", params, optFns, c.addOperationListDatabasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDatabasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDatabasesInput struct {

	// The name of the database. This parameter is required when authenticating using
	// either Secrets Manager or temporary credentials.
	//
	// This member is required.
	Database *string

	// The cluster identifier. This parameter is required when connecting to a cluster
	// and authenticating using either Secrets Manager or temporary credentials.
	ClusterIdentifier *string

	// The database user name. This parameter is required when connecting to a cluster
	// and authenticating using temporary credentials.
	DbUser *string

	// The maximum number of databases to return in the response. If more databases
	// exist than fit in one response, then NextToken is returned to page through the
	// results.
	MaxResults int32

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned NextToken value in the next
	// NextToken parameter and retrying the command. If the NextToken field is empty,
	// all response records have been retrieved for the request.
	NextToken *string

	// The name or ARN of the secret that enables access to the database. This
	// parameter is required when authenticating using Secrets Manager.
	SecretArn *string

	// The serverless workgroup name. This parameter is required when connecting to a
	// serverless workgroup and authenticating using either Secrets Manager or
	// temporary credentials.
	WorkgroupName *string

	noSmithyDocumentSerde
}

type ListDatabasesOutput struct {

	// The names of databases.
	Databases []string

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned NextToken value in the next
	// NextToken parameter and retrying the command. If the NextToken field is empty,
	// all response records have been retrieved for the request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDatabasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDatabases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDatabases{}, middleware.After)
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
	if err = addOpListDatabasesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDatabases(options.Region), middleware.Before); err != nil {
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

// ListDatabasesAPIClient is a client that implements the ListDatabases operation.
type ListDatabasesAPIClient interface {
	ListDatabases(context.Context, *ListDatabasesInput, ...func(*Options)) (*ListDatabasesOutput, error)
}

var _ ListDatabasesAPIClient = (*Client)(nil)

// ListDatabasesPaginatorOptions is the paginator options for ListDatabases
type ListDatabasesPaginatorOptions struct {
	// The maximum number of databases to return in the response. If more databases
	// exist than fit in one response, then NextToken is returned to page through the
	// results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDatabasesPaginator is a paginator for ListDatabases
type ListDatabasesPaginator struct {
	options   ListDatabasesPaginatorOptions
	client    ListDatabasesAPIClient
	params    *ListDatabasesInput
	nextToken *string
	firstPage bool
}

// NewListDatabasesPaginator returns a new ListDatabasesPaginator
func NewListDatabasesPaginator(client ListDatabasesAPIClient, params *ListDatabasesInput, optFns ...func(*ListDatabasesPaginatorOptions)) *ListDatabasesPaginator {
	if params == nil {
		params = &ListDatabasesInput{}
	}

	options := ListDatabasesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDatabasesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDatabasesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDatabases page.
func (p *ListDatabasesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDatabasesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListDatabases(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDatabases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift-data",
		OperationName: "ListDatabases",
	}
}
