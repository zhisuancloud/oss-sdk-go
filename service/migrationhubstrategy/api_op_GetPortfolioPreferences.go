// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhubstrategy

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/migrationhubstrategy/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves your migration and modernization preferences.
func (c *Client) GetPortfolioPreferences(ctx context.Context, params *GetPortfolioPreferencesInput, optFns ...func(*Options)) (*GetPortfolioPreferencesOutput, error) {
	if params == nil {
		params = &GetPortfolioPreferencesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPortfolioPreferences", params, optFns, c.addOperationGetPortfolioPreferencesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPortfolioPreferencesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPortfolioPreferencesInput struct {
	noSmithyDocumentSerde
}

type GetPortfolioPreferencesOutput struct {

	// The transformation preferences for non-database applications.
	ApplicationPreferences *types.ApplicationPreferences

	// The transformation preferences for database applications.
	DatabasePreferences *types.DatabasePreferences

	// The rank of business goals based on priority.
	PrioritizeBusinessGoals *types.PrioritizeBusinessGoals

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPortfolioPreferencesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetPortfolioPreferences{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPortfolioPreferences{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPortfolioPreferences(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetPortfolioPreferences(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "migrationhub-strategy",
		OperationName: "GetPortfolioPreferences",
	}
}
