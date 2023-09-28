// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about all grants in the Amazon Web Services account and
// Region that have the specified retiring principal. You can specify any principal
// in your Amazon Web Services account. The grants that are returned include grants
// for KMS keys in your Amazon Web Services account and other Amazon Web Services
// accounts. You might use this operation to determine which grants you may retire.
// To retire a grant, use the RetireGrant operation. For detailed information about
// grants, including grant terminology, see Grants in KMS
// (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html) in the Key
// Management Service Developer Guide . For examples of working with grants in
// several programming languages, see Programming grants
// (https://docs.aws.amazon.com/kms/latest/developerguide/programming-grants.html).
// Cross-account use: You must specify a principal in your Amazon Web Services
// account. However, this operation can return grants in any Amazon Web Services
// account. You do not need kms:ListRetirableGrants permission (or any other
// additional permission) in any Amazon Web Services account other than your own.
// Required permissions: kms:ListRetirableGrants
// (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// (IAM policy) in your Amazon Web Services account. Related operations:
//
// *
// CreateGrant
//
// * ListGrants
//
// * RetireGrant
//
// * RevokeGrant
func (c *Client) ListRetirableGrants(ctx context.Context, params *ListRetirableGrantsInput, optFns ...func(*Options)) (*ListRetirableGrantsOutput, error) {
	if params == nil {
		params = &ListRetirableGrantsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRetirableGrants", params, optFns, c.addOperationListRetirableGrantsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRetirableGrantsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRetirableGrantsInput struct {

	// The retiring principal for which to list grants. Enter a principal in your
	// Amazon Web Services account. To specify the retiring principal, use the Amazon
	// Resource Name (ARN)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// an Amazon Web Services principal. Valid Amazon Web Services principals include
	// Amazon Web Services accounts (root), IAM users, federated users, and assumed
	// role users. For examples of the ARN syntax for specifying a principal, see
	// Amazon Web Services Identity and Access Management (IAM)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-iam)
	// in the Example ARNs section of the Amazon Web Services General Reference.
	//
	// This member is required.
	RetiringPrincipal *string

	// Use this parameter to specify the maximum number of items to return. When this
	// value is present, KMS does not return more than the specified number of items,
	// but it might return fewer. This value is optional. If you include a value, it
	// must be between 1 and 100, inclusive. If you do not include a value, it defaults
	// to 50.
	Limit *int32

	// Use this parameter in a subsequent request after you receive a response with
	// truncated results. Set it to the value of NextMarker from the truncated response
	// you just received.
	Marker *string

	noSmithyDocumentSerde
}

type ListRetirableGrantsOutput struct {

	// A list of grants.
	Grants []types.GrantListEntry

	// When Truncated is true, this element is present and contains the value to use
	// for the Marker parameter in a subsequent request.
	NextMarker *string

	// A flag that indicates whether there are more items in the list. When this value
	// is true, the list in this response is truncated. To get more items, pass the
	// value of the NextMarker element in thisresponse to the Marker parameter in a
	// subsequent request.
	Truncated bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRetirableGrantsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListRetirableGrants{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRetirableGrants{}, middleware.After)
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
	if err = addOpListRetirableGrantsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRetirableGrants(options.Region), middleware.Before); err != nil {
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

// ListRetirableGrantsAPIClient is a client that implements the ListRetirableGrants
// operation.
type ListRetirableGrantsAPIClient interface {
	ListRetirableGrants(context.Context, *ListRetirableGrantsInput, ...func(*Options)) (*ListRetirableGrantsOutput, error)
}

var _ ListRetirableGrantsAPIClient = (*Client)(nil)

// ListRetirableGrantsPaginatorOptions is the paginator options for
// ListRetirableGrants
type ListRetirableGrantsPaginatorOptions struct {
	// Use this parameter to specify the maximum number of items to return. When this
	// value is present, KMS does not return more than the specified number of items,
	// but it might return fewer. This value is optional. If you include a value, it
	// must be between 1 and 100, inclusive. If you do not include a value, it defaults
	// to 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRetirableGrantsPaginator is a paginator for ListRetirableGrants
type ListRetirableGrantsPaginator struct {
	options   ListRetirableGrantsPaginatorOptions
	client    ListRetirableGrantsAPIClient
	params    *ListRetirableGrantsInput
	nextToken *string
	firstPage bool
}

// NewListRetirableGrantsPaginator returns a new ListRetirableGrantsPaginator
func NewListRetirableGrantsPaginator(client ListRetirableGrantsAPIClient, params *ListRetirableGrantsInput, optFns ...func(*ListRetirableGrantsPaginatorOptions)) *ListRetirableGrantsPaginator {
	if params == nil {
		params = &ListRetirableGrantsInput{}
	}

	options := ListRetirableGrantsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRetirableGrantsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRetirableGrantsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRetirableGrants page.
func (p *ListRetirableGrantsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRetirableGrantsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListRetirableGrants(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListRetirableGrants(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "ListRetirableGrants",
	}
}
