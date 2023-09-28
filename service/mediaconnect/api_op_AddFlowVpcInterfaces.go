// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/mediaconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds VPC interfaces to flow
func (c *Client) AddFlowVpcInterfaces(ctx context.Context, params *AddFlowVpcInterfacesInput, optFns ...func(*Options)) (*AddFlowVpcInterfacesOutput, error) {
	if params == nil {
		params = &AddFlowVpcInterfacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddFlowVpcInterfaces", params, optFns, c.addOperationAddFlowVpcInterfacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddFlowVpcInterfacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to add VPC interfaces to the flow.
type AddFlowVpcInterfacesInput struct {

	// The flow that you want to mutate.
	//
	// This member is required.
	FlowArn *string

	// A list of VPC interfaces that you want to add.
	//
	// This member is required.
	VpcInterfaces []types.VpcInterfaceRequest

	noSmithyDocumentSerde
}

type AddFlowVpcInterfacesOutput struct {

	// The ARN of the flow that these VPC interfaces were added to.
	FlowArn *string

	// The details of the newly added VPC interfaces.
	VpcInterfaces []types.VpcInterface

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddFlowVpcInterfacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAddFlowVpcInterfaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAddFlowVpcInterfaces{}, middleware.After)
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
	if err = addOpAddFlowVpcInterfacesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddFlowVpcInterfaces(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddFlowVpcInterfaces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "AddFlowVpcInterfaces",
	}
}
