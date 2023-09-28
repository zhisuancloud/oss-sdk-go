// Code generated by smithy-go-codegen DO NOT EDIT.

package devopsguru

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the number of open reactive insights, the number of open proactive
// insights, and the number of metrics analyzed in your Amazon Web Services
// account. Use these numbers to gauge the health of operations in your Amazon Web
// Services account.
func (c *Client) DescribeAccountHealth(ctx context.Context, params *DescribeAccountHealthInput, optFns ...func(*Options)) (*DescribeAccountHealthOutput, error) {
	if params == nil {
		params = &DescribeAccountHealthInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAccountHealth", params, optFns, c.addOperationDescribeAccountHealthMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAccountHealthOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccountHealthInput struct {
	noSmithyDocumentSerde
}

type DescribeAccountHealthOutput struct {

	// An integer that specifies the number of metrics that have been analyzed in your
	// Amazon Web Services account.
	//
	// This member is required.
	MetricsAnalyzed int32

	// An integer that specifies the number of open proactive insights in your Amazon
	// Web Services account.
	//
	// This member is required.
	OpenProactiveInsights int32

	// An integer that specifies the number of open reactive insights in your Amazon
	// Web Services account.
	//
	// This member is required.
	OpenReactiveInsights int32

	// The number of Amazon DevOps Guru resource analysis hours billed to the current
	// Amazon Web Services account in the last hour.
	//
	// This member is required.
	ResourceHours *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAccountHealthMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAccountHealth{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAccountHealth{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccountHealth(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAccountHealth(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devops-guru",
		OperationName: "DescribeAccountHealth",
	}
}