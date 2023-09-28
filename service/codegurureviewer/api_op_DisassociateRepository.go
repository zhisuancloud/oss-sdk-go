// Code generated by smithy-go-codegen DO NOT EDIT.

package codegurureviewer

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/codegurureviewer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the association between Amazon CodeGuru Reviewer and a repository.
func (c *Client) DisassociateRepository(ctx context.Context, params *DisassociateRepositoryInput, optFns ...func(*Options)) (*DisassociateRepositoryOutput, error) {
	if params == nil {
		params = &DisassociateRepositoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateRepository", params, optFns, c.addOperationDisassociateRepositoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateRepositoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateRepositoryInput struct {

	// The Amazon Resource Name (ARN) of the RepositoryAssociation
	// (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_RepositoryAssociation.html)
	// object. You can retrieve this ARN by calling ListRepositoryAssociations
	// (https://docs.aws.amazon.com/codeguru/latest/reviewer-api/API_ListRepositoryAssociations.html).
	//
	// This member is required.
	AssociationArn *string

	noSmithyDocumentSerde
}

type DisassociateRepositoryOutput struct {

	// Information about the disassociated repository.
	RepositoryAssociation *types.RepositoryAssociation

	// An array of key-value pairs used to tag an associated repository. A tag is a
	// custom attribute label with two parts:
	//
	// * A tag key (for example, CostCenter,
	// Environment, Project, or Secret). Tag keys are case sensitive.
	//
	// * An optional
	// field known as a tag value (for example, 111122223333, Production, or a team
	// name). Omitting the tag value is the same as using an empty string. Like tag
	// keys, tag values are case sensitive.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateRepositoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateRepository{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateRepository{}, middleware.After)
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
	if err = addOpDisassociateRepositoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateRepository(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateRepository(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-reviewer",
		OperationName: "DisassociateRepository",
	}
}
