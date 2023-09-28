// Code generated by smithy-go-codegen DO NOT EDIT.

package query

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/internal/protocoltest/query/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

func (c *Client) XmlEmptyLists(ctx context.Context, params *XmlEmptyListsInput, optFns ...func(*Options)) (*XmlEmptyListsOutput, error) {
	if params == nil {
		params = &XmlEmptyListsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "XmlEmptyLists", params, optFns, c.addOperationXmlEmptyListsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*XmlEmptyListsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type XmlEmptyListsInput struct {
	noSmithyDocumentSerde
}

type XmlEmptyListsOutput struct {
	BooleanList []bool

	EnumList []types.FooEnum

	FlattenedList []string

	FlattenedList2 []string

	FlattenedListWithMemberNamespace []string

	FlattenedListWithNamespace []string

	IntegerList []int32

	// A list of lists of strings.
	NestedStringList [][]string

	RenamedListMembers []string

	StringList []string

	StringSet []string

	StructureList []types.StructureListMember

	TimestampList []time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationXmlEmptyListsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpXmlEmptyLists{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpXmlEmptyLists{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opXmlEmptyLists(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opXmlEmptyLists(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "XmlEmptyLists",
	}
}
