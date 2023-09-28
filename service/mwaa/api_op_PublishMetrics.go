// Code generated by smithy-go-codegen DO NOT EDIT.

package mwaa

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/mwaa/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Internal only. Publishes environment health metrics to Amazon CloudWatch.
func (c *Client) PublishMetrics(ctx context.Context, params *PublishMetricsInput, optFns ...func(*Options)) (*PublishMetricsOutput, error) {
	if params == nil {
		params = &PublishMetricsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PublishMetrics", params, optFns, c.addOperationPublishMetricsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PublishMetricsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PublishMetricsInput struct {

	// Internal only. The name of the environment.
	//
	// This member is required.
	EnvironmentName *string

	// Internal only. Publishes metrics to Amazon CloudWatch. To learn more about the
	// metrics published to Amazon CloudWatch, see Amazon MWAA performance metrics in
	// Amazon CloudWatch
	// (https://docs.aws.amazon.com/mwaa/latest/userguide/cw-metrics.html).
	//
	// This member is required.
	MetricData []types.MetricDatum

	noSmithyDocumentSerde
}

type PublishMetricsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPublishMetricsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPublishMetrics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPublishMetrics{}, middleware.After)
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
	if err = addEndpointPrefix_opPublishMetricsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPublishMetricsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPublishMetrics(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opPublishMetricsMiddleware struct {
}

func (*endpointPrefix_opPublishMetricsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opPublishMetricsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "ops." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opPublishMetricsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opPublishMetricsMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opPublishMetrics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "airflow",
		OperationName: "PublishMetrics",
	}
}
