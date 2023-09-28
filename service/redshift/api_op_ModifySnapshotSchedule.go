// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Modifies a snapshot schedule. Any schedule associated with a cluster is modified
// asynchronously.
func (c *Client) ModifySnapshotSchedule(ctx context.Context, params *ModifySnapshotScheduleInput, optFns ...func(*Options)) (*ModifySnapshotScheduleOutput, error) {
	if params == nil {
		params = &ModifySnapshotScheduleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifySnapshotSchedule", params, optFns, c.addOperationModifySnapshotScheduleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifySnapshotScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifySnapshotScheduleInput struct {

	// An updated list of schedule definitions. A schedule definition is made up of
	// schedule expressions, for example, "cron(30 12 *)" or "rate(12 hours)".
	//
	// This member is required.
	ScheduleDefinitions []string

	// A unique alphanumeric identifier of the schedule to modify.
	//
	// This member is required.
	ScheduleIdentifier *string

	noSmithyDocumentSerde
}

// Describes a snapshot schedule. You can set a regular interval for creating
// snapshots of a cluster. You can also schedule snapshots for specific dates.
type ModifySnapshotScheduleOutput struct {

	// The number of clusters associated with the schedule.
	AssociatedClusterCount *int32

	// A list of clusters associated with the schedule. A maximum of 100 clusters is
	// returned.
	AssociatedClusters []types.ClusterAssociatedToSchedule

	//
	NextInvocations []time.Time

	// A list of ScheduleDefinitions.
	ScheduleDefinitions []string

	// The description of the schedule.
	ScheduleDescription *string

	// A unique identifier for the schedule.
	ScheduleIdentifier *string

	// An optional set of tags describing the schedule.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifySnapshotScheduleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifySnapshotSchedule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifySnapshotSchedule{}, middleware.After)
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
	if err = addOpModifySnapshotScheduleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifySnapshotSchedule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifySnapshotSchedule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "ModifySnapshotSchedule",
	}
}
