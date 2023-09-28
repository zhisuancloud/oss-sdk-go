// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/robomaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes a simulation application.
func (c *Client) DescribeSimulationApplication(ctx context.Context, params *DescribeSimulationApplicationInput, optFns ...func(*Options)) (*DescribeSimulationApplicationOutput, error) {
	if params == nil {
		params = &DescribeSimulationApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSimulationApplication", params, optFns, c.addOperationDescribeSimulationApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSimulationApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSimulationApplicationInput struct {

	// The application information for the simulation application.
	//
	// This member is required.
	Application *string

	// The version of the simulation application to describe.
	ApplicationVersion *string

	noSmithyDocumentSerde
}

type DescribeSimulationApplicationOutput struct {

	// The Amazon Resource Name (ARN) of the robot simulation application.
	Arn *string

	// The object that contains the Docker image URI used to create the simulation
	// application.
	Environment *types.Environment

	// A SHA256 identifier for the Docker image that you use for your simulation
	// application.
	ImageDigest *string

	// The time, in milliseconds since the epoch, when the simulation application was
	// last updated.
	LastUpdatedAt *time.Time

	// The name of the simulation application.
	Name *string

	// The rendering engine for the simulation application.
	RenderingEngine *types.RenderingEngine

	// The revision id of the simulation application.
	RevisionId *string

	// Information about the robot software suite (ROS distribution).
	RobotSoftwareSuite *types.RobotSoftwareSuite

	// The simulation software suite used by the simulation application.
	SimulationSoftwareSuite *types.SimulationSoftwareSuite

	// The sources of the simulation application.
	Sources []types.Source

	// The list of all tags added to the specified simulation application.
	Tags map[string]string

	// The version of the simulation application.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSimulationApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeSimulationApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeSimulationApplication{}, middleware.After)
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
	if err = addOpDescribeSimulationApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSimulationApplication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeSimulationApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "DescribeSimulationApplication",
	}
}
