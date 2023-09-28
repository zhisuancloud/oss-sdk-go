// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/neptune/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The DeleteDBCluster action deletes a previously provisioned DB cluster. When you
// delete a DB cluster, all automated backups for that DB cluster are deleted and
// can't be recovered. Manual DB cluster snapshots of the specified DB cluster are
// not deleted. Note that the DB Cluster cannot be deleted if deletion protection
// is enabled. To delete it, you must first set its DeletionProtection field to
// False.
func (c *Client) DeleteDBCluster(ctx context.Context, params *DeleteDBClusterInput, optFns ...func(*Options)) (*DeleteDBClusterOutput, error) {
	if params == nil {
		params = &DeleteDBClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDBCluster", params, optFns, c.addOperationDeleteDBClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDBClusterInput struct {

	// The DB cluster identifier for the DB cluster to be deleted. This parameter isn't
	// case-sensitive. Constraints:
	//
	// * Must match an existing DBClusterIdentifier.
	//
	// This member is required.
	DBClusterIdentifier *string

	// The DB cluster snapshot identifier of the new DB cluster snapshot created when
	// SkipFinalSnapshot is set to false. Specifying this parameter and also setting
	// the SkipFinalShapshot parameter to true results in an error. Constraints:
	//
	// *
	// Must be 1 to 255 letters, numbers, or hyphens.
	//
	// * First character must be a
	// letter
	//
	// * Cannot end with a hyphen or contain two consecutive hyphens
	FinalDBSnapshotIdentifier *string

	// Determines whether a final DB cluster snapshot is created before the DB cluster
	// is deleted. If true is specified, no DB cluster snapshot is created. If false is
	// specified, a DB cluster snapshot is created before the DB cluster is deleted.
	// You must specify a FinalDBSnapshotIdentifier parameter if SkipFinalSnapshot is
	// false. Default: false
	SkipFinalSnapshot bool

	noSmithyDocumentSerde
}

type DeleteDBClusterOutput struct {

	// Contains the details of an Amazon Neptune DB cluster. This data type is used as
	// a response element in the DescribeDBClusters action.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDBClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteDBCluster{}, middleware.After)
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
	if err = addOpDeleteDBClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDBCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDBCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DeleteDBCluster",
	}
}
