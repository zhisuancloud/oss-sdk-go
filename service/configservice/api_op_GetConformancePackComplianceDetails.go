// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns compliance details of a conformance pack for all Amazon Web Services
// resources that are monitered by conformance pack.
func (c *Client) GetConformancePackComplianceDetails(ctx context.Context, params *GetConformancePackComplianceDetailsInput, optFns ...func(*Options)) (*GetConformancePackComplianceDetailsOutput, error) {
	if params == nil {
		params = &GetConformancePackComplianceDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetConformancePackComplianceDetails", params, optFns, c.addOperationGetConformancePackComplianceDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetConformancePackComplianceDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetConformancePackComplianceDetailsInput struct {

	// Name of the conformance pack.
	//
	// This member is required.
	ConformancePackName *string

	// A ConformancePackEvaluationFilters object.
	Filters *types.ConformancePackEvaluationFilters

	// The maximum number of evaluation results returned on each page. If you do no
	// specify a number, Config uses the default. The default is 100.
	Limit int32

	// The nextToken string returned in a previous request that you use to request the
	// next page of results in a paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

type GetConformancePackComplianceDetailsOutput struct {

	// Name of the conformance pack.
	//
	// This member is required.
	ConformancePackName *string

	// Returns a list of ConformancePackEvaluationResult objects.
	ConformancePackRuleEvaluationResults []types.ConformancePackEvaluationResult

	// The nextToken string returned in a previous request that you use to request the
	// next page of results in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetConformancePackComplianceDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetConformancePackComplianceDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetConformancePackComplianceDetails{}, middleware.After)
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
	if err = addOpGetConformancePackComplianceDetailsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetConformancePackComplianceDetails(options.Region), middleware.Before); err != nil {
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

// GetConformancePackComplianceDetailsAPIClient is a client that implements the
// GetConformancePackComplianceDetails operation.
type GetConformancePackComplianceDetailsAPIClient interface {
	GetConformancePackComplianceDetails(context.Context, *GetConformancePackComplianceDetailsInput, ...func(*Options)) (*GetConformancePackComplianceDetailsOutput, error)
}

var _ GetConformancePackComplianceDetailsAPIClient = (*Client)(nil)

// GetConformancePackComplianceDetailsPaginatorOptions is the paginator options for
// GetConformancePackComplianceDetails
type GetConformancePackComplianceDetailsPaginatorOptions struct {
	// The maximum number of evaluation results returned on each page. If you do no
	// specify a number, Config uses the default. The default is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetConformancePackComplianceDetailsPaginator is a paginator for
// GetConformancePackComplianceDetails
type GetConformancePackComplianceDetailsPaginator struct {
	options   GetConformancePackComplianceDetailsPaginatorOptions
	client    GetConformancePackComplianceDetailsAPIClient
	params    *GetConformancePackComplianceDetailsInput
	nextToken *string
	firstPage bool
}

// NewGetConformancePackComplianceDetailsPaginator returns a new
// GetConformancePackComplianceDetailsPaginator
func NewGetConformancePackComplianceDetailsPaginator(client GetConformancePackComplianceDetailsAPIClient, params *GetConformancePackComplianceDetailsInput, optFns ...func(*GetConformancePackComplianceDetailsPaginatorOptions)) *GetConformancePackComplianceDetailsPaginator {
	if params == nil {
		params = &GetConformancePackComplianceDetailsInput{}
	}

	options := GetConformancePackComplianceDetailsPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetConformancePackComplianceDetailsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetConformancePackComplianceDetailsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetConformancePackComplianceDetails page.
func (p *GetConformancePackComplianceDetailsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetConformancePackComplianceDetailsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.GetConformancePackComplianceDetails(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetConformancePackComplianceDetails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "GetConformancePackComplianceDetails",
	}
}
