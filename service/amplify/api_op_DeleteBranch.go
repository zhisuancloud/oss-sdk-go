// Code generated by smithy-go-codegen DO NOT EDIT.

package amplify

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/amplify/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a branch for an Amplify app.
func (c *Client) DeleteBranch(ctx context.Context, params *DeleteBranchInput, optFns ...func(*Options)) (*DeleteBranchOutput, error) {
	if params == nil {
		params = &DeleteBranchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBranch", params, optFns, c.addOperationDeleteBranchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBranchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request structure for the delete branch request.
type DeleteBranchInput struct {

	// The unique ID for an Amplify app.
	//
	// This member is required.
	AppId *string

	// The name for the branch.
	//
	// This member is required.
	BranchName *string

	noSmithyDocumentSerde
}

// The result structure for the delete branch request.
type DeleteBranchOutput struct {

	// The branch for an Amplify app, which maps to a third-party repository branch.
	//
	// This member is required.
	Branch *types.Branch

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteBranchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteBranch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteBranch{}, middleware.After)
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
	if err = addOpDeleteBranchValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBranch(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteBranch(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplify",
		OperationName: "DeleteBranch",
	}
}
