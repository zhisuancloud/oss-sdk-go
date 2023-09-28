// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the phone number settings for the administrator's AWS account, such as
// the default outbound calling name.
func (c *Client) GetPhoneNumberSettings(ctx context.Context, params *GetPhoneNumberSettingsInput, optFns ...func(*Options)) (*GetPhoneNumberSettingsOutput, error) {
	if params == nil {
		params = &GetPhoneNumberSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPhoneNumberSettings", params, optFns, c.addOperationGetPhoneNumberSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPhoneNumberSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPhoneNumberSettingsInput struct {
	noSmithyDocumentSerde
}

type GetPhoneNumberSettingsOutput struct {

	// The default outbound calling name for the account.
	CallingName *string

	// The updated outbound calling name timestamp, in ISO 8601 format.
	CallingNameUpdatedTimestamp *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPhoneNumberSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetPhoneNumberSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPhoneNumberSettings{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPhoneNumberSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetPhoneNumberSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "GetPhoneNumberSettings",
	}
}
