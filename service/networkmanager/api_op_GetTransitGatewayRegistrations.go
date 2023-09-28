// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/networkmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about the transit gateway registrations in a specified global
// network.
func (c *Client) GetTransitGatewayRegistrations(ctx context.Context, params *GetTransitGatewayRegistrationsInput, optFns ...func(*Options)) (*GetTransitGatewayRegistrationsOutput, error) {
	if params == nil {
		params = &GetTransitGatewayRegistrationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTransitGatewayRegistrations", params, optFns, c.addOperationGetTransitGatewayRegistrationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTransitGatewayRegistrationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTransitGatewayRegistrationsInput struct {

	// The ID of the global network.
	//
	// This member is required.
	GlobalNetworkId *string

	// The maximum number of results to return.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The Amazon Resource Names (ARNs) of one or more transit gateways. The maximum is
	// 10.
	TransitGatewayArns []string

	noSmithyDocumentSerde
}

type GetTransitGatewayRegistrationsOutput struct {

	// The token for the next page of results.
	NextToken *string

	// The transit gateway registrations.
	TransitGatewayRegistrations []types.TransitGatewayRegistration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTransitGatewayRegistrationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetTransitGatewayRegistrations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTransitGatewayRegistrations{}, middleware.After)
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
	if err = addOpGetTransitGatewayRegistrationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTransitGatewayRegistrations(options.Region), middleware.Before); err != nil {
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

// GetTransitGatewayRegistrationsAPIClient is a client that implements the
// GetTransitGatewayRegistrations operation.
type GetTransitGatewayRegistrationsAPIClient interface {
	GetTransitGatewayRegistrations(context.Context, *GetTransitGatewayRegistrationsInput, ...func(*Options)) (*GetTransitGatewayRegistrationsOutput, error)
}

var _ GetTransitGatewayRegistrationsAPIClient = (*Client)(nil)

// GetTransitGatewayRegistrationsPaginatorOptions is the paginator options for
// GetTransitGatewayRegistrations
type GetTransitGatewayRegistrationsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetTransitGatewayRegistrationsPaginator is a paginator for
// GetTransitGatewayRegistrations
type GetTransitGatewayRegistrationsPaginator struct {
	options   GetTransitGatewayRegistrationsPaginatorOptions
	client    GetTransitGatewayRegistrationsAPIClient
	params    *GetTransitGatewayRegistrationsInput
	nextToken *string
	firstPage bool
}

// NewGetTransitGatewayRegistrationsPaginator returns a new
// GetTransitGatewayRegistrationsPaginator
func NewGetTransitGatewayRegistrationsPaginator(client GetTransitGatewayRegistrationsAPIClient, params *GetTransitGatewayRegistrationsInput, optFns ...func(*GetTransitGatewayRegistrationsPaginatorOptions)) *GetTransitGatewayRegistrationsPaginator {
	if params == nil {
		params = &GetTransitGatewayRegistrationsInput{}
	}

	options := GetTransitGatewayRegistrationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetTransitGatewayRegistrationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetTransitGatewayRegistrationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetTransitGatewayRegistrations page.
func (p *GetTransitGatewayRegistrationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetTransitGatewayRegistrationsOutput, error) {
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

	result, err := p.client.GetTransitGatewayRegistrations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetTransitGatewayRegistrations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "GetTransitGatewayRegistrations",
	}
}
