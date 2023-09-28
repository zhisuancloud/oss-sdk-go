// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes the specified collection. You can use DescribeCollection to get
// information, such as the number of faces indexed into a collection and the
// version of the model used by the collection for face detection. For more
// information, see Describing a Collection in the Amazon Rekognition Developer
// Guide.
func (c *Client) DescribeCollection(ctx context.Context, params *DescribeCollectionInput, optFns ...func(*Options)) (*DescribeCollectionOutput, error) {
	if params == nil {
		params = &DescribeCollectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCollection", params, optFns, c.addOperationDescribeCollectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCollectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCollectionInput struct {

	// The ID of the collection to describe.
	//
	// This member is required.
	CollectionId *string

	noSmithyDocumentSerde
}

type DescribeCollectionOutput struct {

	// The Amazon Resource Name (ARN) of the collection.
	CollectionARN *string

	// The number of milliseconds since the Unix epoch time until the creation of the
	// collection. The Unix epoch time is 00:00:00 Coordinated Universal Time (UTC),
	// Thursday, 1 January 1970.
	CreationTimestamp *time.Time

	// The number of faces that are indexed into the collection. To index faces into a
	// collection, use IndexFaces.
	FaceCount *int64

	// The version of the face model that's used by the collection for face detection.
	// For more information, see Model versioning in the Amazon Rekognition Developer
	// Guide.
	FaceModelVersion *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCollectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCollection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCollection{}, middleware.After)
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
	if err = addOpDescribeCollectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCollection(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeCollection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "DescribeCollection",
	}
}
