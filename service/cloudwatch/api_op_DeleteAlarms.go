// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified alarms. You can delete up to 100 alarms in one operation.
// However, this total can include no more than one composite alarm. For example,
// you could delete 99 metric alarms and one composite alarms with one operation,
// but you can't delete two composite alarms with one operation. In the event of an
// error, no alarms are deleted. It is possible to create a loop or cycle of
// composite alarms, where composite alarm A depends on composite alarm B, and
// composite alarm B also depends on composite alarm A. In this scenario, you can't
// delete any composite alarm that is part of the cycle because there is always
// still a composite alarm that depends on that alarm that you want to delete. To
// get out of such a situation, you must break the cycle by changing the rule of
// one of the composite alarms in the cycle to remove a dependency that creates the
// cycle. The simplest change to make to break a cycle is to change the AlarmRule
// of one of the alarms to False. Additionally, the evaluation of composite alarms
// stops if CloudWatch detects a cycle in the evaluation path.
func (c *Client) DeleteAlarms(ctx context.Context, params *DeleteAlarmsInput, optFns ...func(*Options)) (*DeleteAlarmsOutput, error) {
	if params == nil {
		params = &DeleteAlarmsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAlarms", params, optFns, c.addOperationDeleteAlarmsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAlarmsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAlarmsInput struct {

	// The alarms to be deleted.
	//
	// This member is required.
	AlarmNames []string

	noSmithyDocumentSerde
}

type DeleteAlarmsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteAlarmsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteAlarms{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteAlarms{}, middleware.After)
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
	if err = addOpDeleteAlarmsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAlarms(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteAlarms(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "DeleteAlarms",
	}
}
