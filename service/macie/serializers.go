// Code generated by smithy-go-codegen DO NOT EDIT.

package macie

import (
	"bytes"
	"context"
	"fmt"
	"oss-sdk-go/service/macie/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/encoding/httpbinding"
	smithyjson "github.com/aws/smithy-go/encoding/json"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"path"
)

type awsAwsjson11_serializeOpAssociateMemberAccount struct {
}

func (*awsAwsjson11_serializeOpAssociateMemberAccount) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpAssociateMemberAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*AssociateMemberAccountInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("MacieService.AssociateMemberAccount")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentAssociateMemberAccountInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpAssociateS3Resources struct {
}

func (*awsAwsjson11_serializeOpAssociateS3Resources) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpAssociateS3Resources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*AssociateS3ResourcesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("MacieService.AssociateS3Resources")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentAssociateS3ResourcesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDisassociateMemberAccount struct {
}

func (*awsAwsjson11_serializeOpDisassociateMemberAccount) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDisassociateMemberAccount) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DisassociateMemberAccountInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("MacieService.DisassociateMemberAccount")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentDisassociateMemberAccountInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDisassociateS3Resources struct {
}

func (*awsAwsjson11_serializeOpDisassociateS3Resources) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDisassociateS3Resources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DisassociateS3ResourcesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("MacieService.DisassociateS3Resources")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentDisassociateS3ResourcesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpListMemberAccounts struct {
}

func (*awsAwsjson11_serializeOpListMemberAccounts) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpListMemberAccounts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListMemberAccountsInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("MacieService.ListMemberAccounts")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentListMemberAccountsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpListS3Resources struct {
}

func (*awsAwsjson11_serializeOpListS3Resources) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpListS3Resources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ListS3ResourcesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("MacieService.ListS3Resources")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentListS3ResourcesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpUpdateS3Resources struct {
}

func (*awsAwsjson11_serializeOpUpdateS3Resources) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpUpdateS3Resources) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*UpdateS3ResourcesInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	operationPath := "/"
	if len(request.Request.URL.Path) == 0 {
		request.Request.URL.Path = operationPath
	} else {
		request.Request.URL.Path = path.Join(request.Request.URL.Path, operationPath)
		if request.Request.URL.Path != "/" && operationPath[len(operationPath)-1] == '/' {
			request.Request.URL.Path += "/"
		}
	}
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("MacieService.UpdateS3Resources")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeOpDocumentUpdateS3ResourcesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsAwsjson11_serializeDocumentClassificationType(v *types.ClassificationType, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.Continuous) > 0 {
		ok := object.Key("continuous")
		ok.String(string(v.Continuous))
	}

	if len(v.OneTime) > 0 {
		ok := object.Key("oneTime")
		ok.String(string(v.OneTime))
	}

	return nil
}

func awsAwsjson11_serializeDocumentClassificationTypeUpdate(v *types.ClassificationTypeUpdate, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.Continuous) > 0 {
		ok := object.Key("continuous")
		ok.String(string(v.Continuous))
	}

	if len(v.OneTime) > 0 {
		ok := object.Key("oneTime")
		ok.String(string(v.OneTime))
	}

	return nil
}

func awsAwsjson11_serializeDocumentS3Resource(v *types.S3Resource, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.BucketName != nil {
		ok := object.Key("bucketName")
		ok.String(*v.BucketName)
	}

	if v.Prefix != nil {
		ok := object.Key("prefix")
		ok.String(*v.Prefix)
	}

	return nil
}

func awsAwsjson11_serializeDocumentS3ResourceClassification(v *types.S3ResourceClassification, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.BucketName != nil {
		ok := object.Key("bucketName")
		ok.String(*v.BucketName)
	}

	if v.ClassificationType != nil {
		ok := object.Key("classificationType")
		if err := awsAwsjson11_serializeDocumentClassificationType(v.ClassificationType, ok); err != nil {
			return err
		}
	}

	if v.Prefix != nil {
		ok := object.Key("prefix")
		ok.String(*v.Prefix)
	}

	return nil
}

func awsAwsjson11_serializeDocumentS3ResourceClassificationUpdate(v *types.S3ResourceClassificationUpdate, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.BucketName != nil {
		ok := object.Key("bucketName")
		ok.String(*v.BucketName)
	}

	if v.ClassificationTypeUpdate != nil {
		ok := object.Key("classificationTypeUpdate")
		if err := awsAwsjson11_serializeDocumentClassificationTypeUpdate(v.ClassificationTypeUpdate, ok); err != nil {
			return err
		}
	}

	if v.Prefix != nil {
		ok := object.Key("prefix")
		ok.String(*v.Prefix)
	}

	return nil
}

func awsAwsjson11_serializeDocumentS3Resources(v []types.S3Resource, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson11_serializeDocumentS3Resource(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeDocumentS3ResourcesClassification(v []types.S3ResourceClassification, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson11_serializeDocumentS3ResourceClassification(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeDocumentS3ResourcesClassificationUpdate(v []types.S3ResourceClassificationUpdate, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if err := awsAwsjson11_serializeDocumentS3ResourceClassificationUpdate(&v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeOpDocumentAssociateMemberAccountInput(v *AssociateMemberAccountInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MemberAccountId != nil {
		ok := object.Key("memberAccountId")
		ok.String(*v.MemberAccountId)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentAssociateS3ResourcesInput(v *AssociateS3ResourcesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MemberAccountId != nil {
		ok := object.Key("memberAccountId")
		ok.String(*v.MemberAccountId)
	}

	if v.S3Resources != nil {
		ok := object.Key("s3Resources")
		if err := awsAwsjson11_serializeDocumentS3ResourcesClassification(v.S3Resources, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentDisassociateMemberAccountInput(v *DisassociateMemberAccountInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MemberAccountId != nil {
		ok := object.Key("memberAccountId")
		ok.String(*v.MemberAccountId)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentDisassociateS3ResourcesInput(v *DisassociateS3ResourcesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AssociatedS3Resources != nil {
		ok := object.Key("associatedS3Resources")
		if err := awsAwsjson11_serializeDocumentS3Resources(v.AssociatedS3Resources, ok); err != nil {
			return err
		}
	}

	if v.MemberAccountId != nil {
		ok := object.Key("memberAccountId")
		ok.String(*v.MemberAccountId)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentListMemberAccountsInput(v *ListMemberAccountsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxResults != nil {
		ok := object.Key("maxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("nextToken")
		ok.String(*v.NextToken)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentListS3ResourcesInput(v *ListS3ResourcesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MaxResults != nil {
		ok := object.Key("maxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.MemberAccountId != nil {
		ok := object.Key("memberAccountId")
		ok.String(*v.MemberAccountId)
	}

	if v.NextToken != nil {
		ok := object.Key("nextToken")
		ok.String(*v.NextToken)
	}

	return nil
}

func awsAwsjson11_serializeOpDocumentUpdateS3ResourcesInput(v *UpdateS3ResourcesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.MemberAccountId != nil {
		ok := object.Key("memberAccountId")
		ok.String(*v.MemberAccountId)
	}

	if v.S3ResourcesUpdate != nil {
		ok := object.Key("s3ResourcesUpdate")
		if err := awsAwsjson11_serializeDocumentS3ResourcesClassificationUpdate(v.S3ResourcesUpdate, ok); err != nil {
			return err
		}
	}

	return nil
}
