// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the count, average, sum, minimum, maximum, sum of squares, variance, and
// standard deviation for the specified aggregated field. If the aggregation field
// is of type String, only the count statistic is returned. Requires permission to
// access the GetStatistics
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) GetStatistics(ctx context.Context, params *GetStatisticsInput, optFns ...func(*Options)) (*GetStatisticsOutput, error) {
	if params == nil {
		params = &GetStatisticsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetStatistics", params, optFns, c.addOperationGetStatisticsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetStatisticsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetStatisticsInput struct {

	// The query used to search. You can specify "*" for the query string to get the
	// count of all indexed things in your Amazon Web Services account.
	//
	// This member is required.
	QueryString *string

	// The aggregation field name.
	AggregationField *string

	// The name of the index to search. The default value is AWS_Things.
	IndexName *string

	// The version of the query used to search.
	QueryVersion *string

	noSmithyDocumentSerde
}

type GetStatisticsOutput struct {

	// The statistics returned by the Fleet Indexing service based on the query and
	// aggregation field.
	Statistics *types.Statistics

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetStatisticsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetStatistics{}, middleware.After)
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
	if err = addOpGetStatisticsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetStatistics(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetStatistics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "GetStatistics",
	}
}
