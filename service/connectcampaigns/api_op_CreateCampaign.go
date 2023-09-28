// Code generated by smithy-go-codegen DO NOT EDIT.

package connectcampaigns

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/connectcampaigns/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a campaign for the specified Amazon Connect account. This API is
// idempotent.
func (c *Client) CreateCampaign(ctx context.Context, params *CreateCampaignInput, optFns ...func(*Options)) (*CreateCampaignOutput, error) {
	if params == nil {
		params = &CreateCampaignInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCampaign", params, optFns, c.addOperationCreateCampaignMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCampaignOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request for Create Campaign API.
type CreateCampaignInput struct {

	// Amazon Connect Instance Id
	//
	// This member is required.
	ConnectInstanceId *string

	// The possible types of dialer config parameters
	//
	// This member is required.
	DialerConfig types.DialerConfig

	// The name of an Amazon Connect Campaign name.
	//
	// This member is required.
	Name *string

	// The configuration used for outbound calls.
	//
	// This member is required.
	OutboundCallConfig *types.OutboundCallConfig

	// Tag map with key and value.
	Tags map[string]string

	noSmithyDocumentSerde
}

// The response for Create Campaign API
type CreateCampaignOutput struct {

	// The resource name of an Amazon Connect campaign.
	Arn *string

	// Identifier representing a Campaign
	Id *string

	// Tag map with key and value.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCampaignMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateCampaign{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateCampaign{}, middleware.After)
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
	if err = addOpCreateCampaignValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCampaign(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCampaign(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect-campaigns",
		OperationName: "CreateCampaign",
	}
}