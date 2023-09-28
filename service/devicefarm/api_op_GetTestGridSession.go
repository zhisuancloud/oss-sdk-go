// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/devicefarm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A session is an instance of a browser created through a RemoteWebDriver with the
// URL from CreateTestGridUrlResult$url. You can use the following to look up
// sessions:
//
// * The session ARN (GetTestGridSessionRequest$sessionArn).
//
// * The
// project ARN and a session ID (GetTestGridSessionRequest$projectArn and
// GetTestGridSessionRequest$sessionId).
func (c *Client) GetTestGridSession(ctx context.Context, params *GetTestGridSessionInput, optFns ...func(*Options)) (*GetTestGridSessionOutput, error) {
	if params == nil {
		params = &GetTestGridSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTestGridSession", params, optFns, c.addOperationGetTestGridSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTestGridSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTestGridSessionInput struct {

	// The ARN for the project that this session belongs to. See CreateTestGridProject
	// and ListTestGridProjects.
	ProjectArn *string

	// An ARN that uniquely identifies a TestGridSession.
	SessionArn *string

	// An ID associated with this session.
	SessionId *string

	noSmithyDocumentSerde
}

type GetTestGridSessionOutput struct {

	// The TestGridSession that was requested.
	TestGridSession *types.TestGridSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTestGridSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetTestGridSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetTestGridSession{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTestGridSession(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTestGridSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "GetTestGridSession",
	}
}
