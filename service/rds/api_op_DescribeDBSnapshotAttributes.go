// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of DB snapshot attribute names and values for a manual DB
// snapshot. When sharing snapshots with other Amazon Web Services accounts,
// DescribeDBSnapshotAttributes returns the restore attribute and a list of IDs for
// the Amazon Web Services accounts that are authorized to copy or restore the
// manual DB snapshot. If all is included in the list of values for the restore
// attribute, then the manual DB snapshot is public and can be copied or restored
// by all Amazon Web Services accounts. To add or remove access for an Amazon Web
// Services account to copy or restore a manual DB snapshot, or to make the manual
// DB snapshot public or private, use the ModifyDBSnapshotAttribute API action.
func (c *Client) DescribeDBSnapshotAttributes(ctx context.Context, params *DescribeDBSnapshotAttributesInput, optFns ...func(*Options)) (*DescribeDBSnapshotAttributesOutput, error) {
	if params == nil {
		params = &DescribeDBSnapshotAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBSnapshotAttributes", params, optFns, c.addOperationDescribeDBSnapshotAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBSnapshotAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDBSnapshotAttributesInput struct {

	// The identifier for the DB snapshot to describe the attributes for.
	//
	// This member is required.
	DBSnapshotIdentifier *string

	noSmithyDocumentSerde
}

type DescribeDBSnapshotAttributesOutput struct {

	// Contains the results of a successful call to the DescribeDBSnapshotAttributes
	// API action. Manual DB snapshot attributes are used to authorize other Amazon Web
	// Services accounts to copy or restore a manual DB snapshot. For more information,
	// see the ModifyDBSnapshotAttribute API action.
	DBSnapshotAttributesResult *types.DBSnapshotAttributesResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDBSnapshotAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBSnapshotAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBSnapshotAttributes{}, middleware.After)
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
	if err = addOpDescribeDBSnapshotAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBSnapshotAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDBSnapshotAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBSnapshotAttributes",
	}
}
