// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes a context.
func (c *Client) DescribeContext(ctx context.Context, params *DescribeContextInput, optFns ...func(*Options)) (*DescribeContextOutput, error) {
	if params == nil {
		params = &DescribeContextInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeContext", params, optFns, c.addOperationDescribeContextMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeContextOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeContextInput struct {

	// The name of the context to describe.
	//
	// This member is required.
	ContextName *string

	noSmithyDocumentSerde
}

type DescribeContextOutput struct {

	// The Amazon Resource Name (ARN) of the context.
	ContextArn *string

	// The name of the context.
	ContextName *string

	// The type of the context.
	ContextType *string

	// Information about the user who created or modified an experiment, trial, trial
	// component, lineage group, or project.
	CreatedBy *types.UserContext

	// When the context was created.
	CreationTime *time.Time

	// The description of the context.
	Description *string

	// Information about the user who created or modified an experiment, trial, trial
	// component, lineage group, or project.
	LastModifiedBy *types.UserContext

	// When the context was last modified.
	LastModifiedTime *time.Time

	// The Amazon Resource Name (ARN) of the lineage group.
	LineageGroupArn *string

	// A list of the context's properties.
	Properties map[string]string

	// The source of the context.
	Source *types.ContextSource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeContextMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeContext{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeContext{}, middleware.After)
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
	if err = addOpDescribeContextValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeContext(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeContext(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeContext",
	}
}
