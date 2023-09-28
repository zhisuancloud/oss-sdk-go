// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describe details for Windows AMIs that are configured for faster launching.
func (c *Client) DescribeFastLaunchImages(ctx context.Context, params *DescribeFastLaunchImagesInput, optFns ...func(*Options)) (*DescribeFastLaunchImagesOutput, error) {
	if params == nil {
		params = &DescribeFastLaunchImagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFastLaunchImages", params, optFns, c.addOperationDescribeFastLaunchImagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFastLaunchImagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFastLaunchImagesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// Use the following filters to streamline results.
	//
	// * resource-type - The resource
	// type for pre-provisioning.
	//
	// * launch-template - The launch template that is
	// associated with the pre-provisioned Windows AMI.
	//
	// * owner-id - The owner ID for
	// the pre-provisioning resource.
	//
	// * state - The current state of fast launching
	// for the Windows AMI.
	Filters []types.Filter

	// Details for one or more Windows AMI image IDs.
	ImageIds []string

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another request with the returned NextToken value. If
	// this parameter is not specified, then all results are returned.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeFastLaunchImagesOutput struct {

	// A collection of details about the fast-launch enabled Windows images that meet
	// the requested criteria.
	FastLaunchImages []types.DescribeFastLaunchImagesSuccessItem

	// The token to use for the next set of results. This value is null when there are
	// no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFastLaunchImagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeFastLaunchImages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeFastLaunchImages{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFastLaunchImages(options.Region), middleware.Before); err != nil {
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

// DescribeFastLaunchImagesAPIClient is a client that implements the
// DescribeFastLaunchImages operation.
type DescribeFastLaunchImagesAPIClient interface {
	DescribeFastLaunchImages(context.Context, *DescribeFastLaunchImagesInput, ...func(*Options)) (*DescribeFastLaunchImagesOutput, error)
}

var _ DescribeFastLaunchImagesAPIClient = (*Client)(nil)

// DescribeFastLaunchImagesPaginatorOptions is the paginator options for
// DescribeFastLaunchImages
type DescribeFastLaunchImagesPaginatorOptions struct {
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another request with the returned NextToken value. If
	// this parameter is not specified, then all results are returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeFastLaunchImagesPaginator is a paginator for DescribeFastLaunchImages
type DescribeFastLaunchImagesPaginator struct {
	options   DescribeFastLaunchImagesPaginatorOptions
	client    DescribeFastLaunchImagesAPIClient
	params    *DescribeFastLaunchImagesInput
	nextToken *string
	firstPage bool
}

// NewDescribeFastLaunchImagesPaginator returns a new
// DescribeFastLaunchImagesPaginator
func NewDescribeFastLaunchImagesPaginator(client DescribeFastLaunchImagesAPIClient, params *DescribeFastLaunchImagesInput, optFns ...func(*DescribeFastLaunchImagesPaginatorOptions)) *DescribeFastLaunchImagesPaginator {
	if params == nil {
		params = &DescribeFastLaunchImagesInput{}
	}

	options := DescribeFastLaunchImagesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeFastLaunchImagesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeFastLaunchImagesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeFastLaunchImages page.
func (p *DescribeFastLaunchImagesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeFastLaunchImagesOutput, error) {
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

	result, err := p.client.DescribeFastLaunchImages(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeFastLaunchImages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeFastLaunchImages",
	}
}
