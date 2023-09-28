// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/elasticsearchservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Allows the source domain owner to delete an existing outbound cross-cluster
// search connection.
func (c *Client) DeleteOutboundCrossClusterSearchConnection(ctx context.Context, params *DeleteOutboundCrossClusterSearchConnectionInput, optFns ...func(*Options)) (*DeleteOutboundCrossClusterSearchConnectionOutput, error) {
	if params == nil {
		params = &DeleteOutboundCrossClusterSearchConnectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteOutboundCrossClusterSearchConnection", params, optFns, c.addOperationDeleteOutboundCrossClusterSearchConnectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteOutboundCrossClusterSearchConnectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the DeleteOutboundCrossClusterSearchConnection
// operation.
type DeleteOutboundCrossClusterSearchConnectionInput struct {

	// The id of the outbound connection that you want to permanently delete.
	//
	// This member is required.
	CrossClusterSearchConnectionId *string

	noSmithyDocumentSerde
}

// The result of a DeleteOutboundCrossClusterSearchConnection operation. Contains
// details of deleted outbound connection.
type DeleteOutboundCrossClusterSearchConnectionOutput struct {

	// Specifies the OutboundCrossClusterSearchConnection of deleted outbound
	// connection.
	CrossClusterSearchConnection *types.OutboundCrossClusterSearchConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteOutboundCrossClusterSearchConnectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteOutboundCrossClusterSearchConnection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteOutboundCrossClusterSearchConnection{}, middleware.After)
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
	if err = addOpDeleteOutboundCrossClusterSearchConnectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteOutboundCrossClusterSearchConnection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteOutboundCrossClusterSearchConnection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DeleteOutboundCrossClusterSearchConnection",
	}
}
