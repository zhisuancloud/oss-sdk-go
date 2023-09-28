// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the maintenance windows in an Amazon Web Services account.
func (c *Client) DescribeMaintenanceWindows(ctx context.Context, params *DescribeMaintenanceWindowsInput, optFns ...func(*Options)) (*DescribeMaintenanceWindowsOutput, error) {
	if params == nil {
		params = &DescribeMaintenanceWindowsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMaintenanceWindows", params, optFns, c.addOperationDescribeMaintenanceWindowsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMaintenanceWindowsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMaintenanceWindowsInput struct {

	// Optional filters used to narrow down the scope of the returned maintenance
	// windows. Supported filter keys are Name and Enabled. For example,
	// Name=MyMaintenanceWindow and Enabled=True.
	Filters []types.MaintenanceWindowFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeMaintenanceWindowsOutput struct {

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Information about the maintenance windows.
	WindowIdentities []types.MaintenanceWindowIdentity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMaintenanceWindowsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMaintenanceWindows{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMaintenanceWindows{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMaintenanceWindows(options.Region), middleware.Before); err != nil {
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

// DescribeMaintenanceWindowsAPIClient is a client that implements the
// DescribeMaintenanceWindows operation.
type DescribeMaintenanceWindowsAPIClient interface {
	DescribeMaintenanceWindows(context.Context, *DescribeMaintenanceWindowsInput, ...func(*Options)) (*DescribeMaintenanceWindowsOutput, error)
}

var _ DescribeMaintenanceWindowsAPIClient = (*Client)(nil)

// DescribeMaintenanceWindowsPaginatorOptions is the paginator options for
// DescribeMaintenanceWindows
type DescribeMaintenanceWindowsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeMaintenanceWindowsPaginator is a paginator for
// DescribeMaintenanceWindows
type DescribeMaintenanceWindowsPaginator struct {
	options   DescribeMaintenanceWindowsPaginatorOptions
	client    DescribeMaintenanceWindowsAPIClient
	params    *DescribeMaintenanceWindowsInput
	nextToken *string
	firstPage bool
}

// NewDescribeMaintenanceWindowsPaginator returns a new
// DescribeMaintenanceWindowsPaginator
func NewDescribeMaintenanceWindowsPaginator(client DescribeMaintenanceWindowsAPIClient, params *DescribeMaintenanceWindowsInput, optFns ...func(*DescribeMaintenanceWindowsPaginatorOptions)) *DescribeMaintenanceWindowsPaginator {
	if params == nil {
		params = &DescribeMaintenanceWindowsInput{}
	}

	options := DescribeMaintenanceWindowsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeMaintenanceWindowsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeMaintenanceWindowsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeMaintenanceWindows page.
func (p *DescribeMaintenanceWindowsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeMaintenanceWindowsOutput, error) {
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

	result, err := p.client.DescribeMaintenanceWindows(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeMaintenanceWindows(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeMaintenanceWindows",
	}
}
