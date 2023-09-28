// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/costexplorer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a forecast for how much Amazon Web Services predicts that you will
// spend over the forecast time period that you select, based on your past costs.
func (c *Client) GetCostForecast(ctx context.Context, params *GetCostForecastInput, optFns ...func(*Options)) (*GetCostForecastOutput, error) {
	if params == nil {
		params = &GetCostForecastInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCostForecast", params, optFns, c.addOperationGetCostForecastMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCostForecastOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCostForecastInput struct {

	// How granular you want the forecast to be. You can get 3 months of DAILY
	// forecasts or 12 months of MONTHLY forecasts. The GetCostForecast operation
	// supports only DAILY and MONTHLY granularities.
	//
	// This member is required.
	Granularity types.Granularity

	// Which metric Cost Explorer uses to create your forecast. For more information
	// about blended and unblended rates, see Why does the "blended" annotation appear
	// on some line items in my bill?
	// (http://aws.amazon.com/premiumsupport/knowledge-center/blended-rates-intro/).
	// Valid values for a GetCostForecast call are the following:
	//
	// * AMORTIZED_COST
	//
	// *
	// BLENDED_COST
	//
	// * NET_AMORTIZED_COST
	//
	// * NET_UNBLENDED_COST
	//
	// * UNBLENDED_COST
	//
	// This member is required.
	Metric types.Metric

	// The period of time that you want the forecast to cover. The start date must be
	// equal to or no later than the current date to avoid a validation error.
	//
	// This member is required.
	TimePeriod *types.DateInterval

	// The filters that you want to use to filter your forecast. The GetCostForecast
	// API supports filtering by the following dimensions:
	//
	// * AZ
	//
	// * INSTANCE_TYPE
	//
	// *
	// LINKED_ACCOUNT
	//
	// * LINKED_ACCOUNT_NAME
	//
	// * OPERATION
	//
	// * PURCHASE_TYPE
	//
	// * REGION
	//
	// *
	// SERVICE
	//
	// * USAGE_TYPE
	//
	// * USAGE_TYPE_GROUP
	//
	// * RECORD_TYPE
	//
	// * OPERATING_SYSTEM
	//
	// *
	// TENANCY
	//
	// * SCOPE
	//
	// * PLATFORM
	//
	// * SUBSCRIPTION_ID
	//
	// * LEGAL_ENTITY_NAME
	//
	// *
	// DEPLOYMENT_OPTION
	//
	// * DATABASE_ENGINE
	//
	// * INSTANCE_TYPE_FAMILY
	//
	// *
	// BILLING_ENTITY
	//
	// * RESERVATION_ID
	//
	// * SAVINGS_PLAN_ARN
	Filter *types.Expression

	// Cost Explorer always returns the mean forecast as a single point. You can
	// request a prediction interval around the mean by specifying a confidence level.
	// The higher the confidence level, the more confident Cost Explorer is about the
	// actual value falling in the prediction interval. Higher confidence levels result
	// in wider prediction intervals.
	PredictionIntervalLevel *int32

	noSmithyDocumentSerde
}

type GetCostForecastOutput struct {

	// The forecasts for your query, in order. For DAILY forecasts, this is a list of
	// days. For MONTHLY forecasts, this is a list of months.
	ForecastResultsByTime []types.ForecastResult

	// How much you are forecasted to spend over the forecast period, in USD.
	Total *types.MetricValue

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCostForecastMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCostForecast{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCostForecast{}, middleware.After)
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
	if err = addOpGetCostForecastValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCostForecast(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCostForecast(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "GetCostForecast",
	}
}
