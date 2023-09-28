// Code generated by smithy-go-codegen DO NOT EDIT.

package drs

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allows you to update the failback replication configuration of a Recovery
// Instance by ID.
func (c *Client) UpdateFailbackReplicationConfiguration(ctx context.Context, params *UpdateFailbackReplicationConfigurationInput, optFns ...func(*Options)) (*UpdateFailbackReplicationConfigurationOutput, error) {
	if params == nil {
		params = &UpdateFailbackReplicationConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateFailbackReplicationConfiguration", params, optFns, c.addOperationUpdateFailbackReplicationConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateFailbackReplicationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateFailbackReplicationConfigurationInput struct {

	// The ID of the Recovery Instance.
	//
	// This member is required.
	RecoveryInstanceID *string

	// Configure bandwidth throttling for the outbound data transfer rate of the
	// Recovery Instance in Mbps.
	BandwidthThrottling int64

	// The name of the Failback Replication Configuration.
	Name *string

	// Whether to use Private IP for the failback replication of the Recovery Instance.
	UsePrivateIP *bool

	noSmithyDocumentSerde
}

type UpdateFailbackReplicationConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateFailbackReplicationConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateFailbackReplicationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateFailbackReplicationConfiguration{}, middleware.After)
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
	if err = addOpUpdateFailbackReplicationConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFailbackReplicationConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateFailbackReplicationConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "drs",
		OperationName: "UpdateFailbackReplicationConfiguration",
	}
}
