// Code generated by smithy-go-codegen DO NOT EDIT.

package workmailmessageflow

import (
	"bytes"
	"context"
	"fmt"
	"oss-sdk-go/service/workmailmessageflow/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

type awsRestjson1_serializeOpGetRawMessageContent struct {
}

func (*awsRestjson1_serializeOpGetRawMessageContent) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetRawMessageContent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetRawMessageContentInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/messages/{messageId}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "GET"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsGetRawMessageContentInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsGetRawMessageContentInput(v *GetRawMessageContentInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.MessageId == nil || len(*v.MessageId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member messageId must not be empty")}
	}
	if v.MessageId != nil {
		if err := encoder.SetURI("messageId").String(*v.MessageId); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpPutRawMessageContent struct {
}

func (*awsRestjson1_serializeOpPutRawMessageContent) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpPutRawMessageContent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*PutRawMessageContentInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/messages/{messageId}")
	request.URL.Path = smithyhttp.JoinPath(request.URL.Path, opPath)
	request.URL.RawQuery = smithyhttp.JoinRawQuery(request.URL.RawQuery, opQuery)
	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeOpHttpBindingsPutRawMessageContentInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeOpDocumentPutRawMessageContentInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeOpHttpBindingsPutRawMessageContentInput(v *PutRawMessageContentInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.MessageId == nil || len(*v.MessageId) == 0 {
		return &smithy.SerializationError{Err: fmt.Errorf("input member messageId must not be empty")}
	}
	if v.MessageId != nil {
		if err := encoder.SetURI("messageId").String(*v.MessageId); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeOpDocumentPutRawMessageContentInput(v *PutRawMessageContentInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Content != nil {
		ok := object.Key("content")
		if err := awsRestjson1_serializeDocumentRawMessageContent(v.Content, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentRawMessageContent(v *types.RawMessageContent, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.S3Reference != nil {
		ok := object.Key("s3Reference")
		if err := awsRestjson1_serializeDocumentS3Reference(v.S3Reference, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentS3Reference(v *types.S3Reference, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Bucket != nil {
		ok := object.Key("bucket")
		ok.String(*v.Bucket)
	}

	if v.Key != nil {
		ok := object.Key("key")
		ok.String(*v.Key)
	}

	if v.ObjectVersion != nil {
		ok := object.Key("objectVersion")
		ok.String(*v.ObjectVersion)
	}

	return nil
}