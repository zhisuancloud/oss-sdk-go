// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new web distribution. You create a CloudFront distribution to tell
// CloudFront where you want content to be delivered from, and the details about
// how to track and manage content delivery. Send a POST request to the /CloudFront
// API version/distribution/distribution ID resource. When you update a
// distribution, there are more required fields than when you create a
// distribution. When you update your distribution by using UpdateDistribution
// (https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_UpdateDistribution.html),
// follow the steps included in the documentation to get the current configuration
// and then make your updates. This helps to make sure that you include all of the
// required fields. To view a summary, see Required Fields for Create Distribution
// and Update Distribution
// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/distribution-overview-required-fields.html)
// in the Amazon CloudFront Developer Guide.
func (c *Client) CreateDistribution(ctx context.Context, params *CreateDistributionInput, optFns ...func(*Options)) (*CreateDistributionOutput, error) {
	if params == nil {
		params = &CreateDistributionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDistribution", params, optFns, c.addOperationCreateDistributionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDistributionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to create a new distribution.
type CreateDistributionInput struct {

	// The distribution's configuration information.
	//
	// This member is required.
	DistributionConfig *types.DistributionConfig

	noSmithyDocumentSerde
}

// The returned result of the corresponding request.
type CreateDistributionOutput struct {

	// The distribution's information.
	Distribution *types.Distribution

	// The current version of the distribution created.
	ETag *string

	// The fully qualified URI of the new distribution resource just created.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDistributionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateDistribution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateDistribution{}, middleware.After)
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
	if err = addOpCreateDistributionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDistribution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDistribution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "CreateDistribution",
	}
}