// Code generated by smithy-go-codegen DO NOT EDIT.

package codestar

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/codestar/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a project, including project resources. This action creates a project
// based on a submitted project request. A set of source code files and a toolchain
// template file can be included with the project request. If these are not
// provided, an empty project is created.
func (c *Client) CreateProject(ctx context.Context, params *CreateProjectInput, optFns ...func(*Options)) (*CreateProjectOutput, error) {
	if params == nil {
		params = &CreateProjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProject", params, optFns, c.addOperationCreateProjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProjectInput struct {

	// The ID of the project to be created in AWS CodeStar.
	//
	// This member is required.
	Id *string

	// The display name for the project to be created in AWS CodeStar.
	//
	// This member is required.
	Name *string

	// A user- or system-generated token that identifies the entity that requested
	// project creation. This token can be used to repeat the request.
	ClientRequestToken *string

	// The description of the project, if any.
	Description *string

	// A list of the Code objects submitted with the project request. If this parameter
	// is specified, the request must also include the toolchain parameter.
	SourceCode []types.Code

	// The tags created for the project.
	Tags map[string]string

	// The name of the toolchain template file submitted with the project request. If
	// this parameter is specified, the request must also include the sourceCode
	// parameter.
	Toolchain *types.Toolchain

	noSmithyDocumentSerde
}

type CreateProjectOutput struct {

	// The Amazon Resource Name (ARN) of the created project.
	//
	// This member is required.
	Arn *string

	// The ID of the project.
	//
	// This member is required.
	Id *string

	// A user- or system-generated token that identifies the entity that requested
	// project creation.
	ClientRequestToken *string

	// Reserved for future use.
	ProjectTemplateId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateProjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateProject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateProject{}, middleware.After)
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
	if err = addOpCreateProjectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProject(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateProject(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar",
		OperationName: "CreateProject",
	}
}
