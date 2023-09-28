// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists tags for one health check or hosted zone. For information about using tags
// for cost allocation, see Using Cost Allocation Tags
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)
// in the Billing and Cost Management User Guide.
func (c *Client) ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) {
	if params == nil {
		params = &ListTagsForResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTagsForResource", params, optFns, c.addOperationListTagsForResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTagsForResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type containing information about a request for a list of the tags
// that are associated with an individual resource.
type ListTagsForResourceInput struct {

	// The ID of the resource for which you want to retrieve tags.
	//
	// This member is required.
	ResourceId *string

	// The type of the resource.
	//
	// * The resource type for health checks is
	// healthcheck.
	//
	// * The resource type for hosted zones is hostedzone.
	//
	// This member is required.
	ResourceType types.TagResourceType

	noSmithyDocumentSerde
}

// A complex type that contains information about the health checks or hosted zones
// for which you want to list tags.
type ListTagsForResourceOutput struct {

	// A ResourceTagSet containing tags associated with the specified resource.
	//
	// This member is required.
	ResourceTagSet *types.ResourceTagSet

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTagsForResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListTagsForResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListTagsForResource{}, middleware.After)
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
	if err = addOpListTagsForResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTagsForResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListTagsForResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "ListTagsForResource",
	}
}
