// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpConstantQueryString struct {
}

func (*validateOpConstantQueryString) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpConstantQueryString) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ConstantQueryStringInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpConstantQueryStringInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpEndpointWithHostLabelHeaderOperation struct {
}

func (*validateOpEndpointWithHostLabelHeaderOperation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpEndpointWithHostLabelHeaderOperation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*EndpointWithHostLabelHeaderOperationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpEndpointWithHostLabelHeaderOperationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpEndpointWithHostLabelOperation struct {
}

func (*validateOpEndpointWithHostLabelOperation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpEndpointWithHostLabelOperation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*EndpointWithHostLabelOperationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpEndpointWithHostLabelOperationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpHttpRequestWithFloatLabels struct {
}

func (*validateOpHttpRequestWithFloatLabels) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpHttpRequestWithFloatLabels) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*HttpRequestWithFloatLabelsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpHttpRequestWithFloatLabelsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpHttpRequestWithGreedyLabelInPath struct {
}

func (*validateOpHttpRequestWithGreedyLabelInPath) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpHttpRequestWithGreedyLabelInPath) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*HttpRequestWithGreedyLabelInPathInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpHttpRequestWithGreedyLabelInPathInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpHttpRequestWithLabelsAndTimestampFormat struct {
}

func (*validateOpHttpRequestWithLabelsAndTimestampFormat) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpHttpRequestWithLabelsAndTimestampFormat) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*HttpRequestWithLabelsAndTimestampFormatInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpHttpRequestWithLabelsAndTimestampFormatInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpHttpRequestWithLabels struct {
}

func (*validateOpHttpRequestWithLabels) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpHttpRequestWithLabels) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*HttpRequestWithLabelsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpHttpRequestWithLabelsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpConstantQueryStringValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpConstantQueryString{}, middleware.After)
}

func addOpEndpointWithHostLabelHeaderOperationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpEndpointWithHostLabelHeaderOperation{}, middleware.After)
}

func addOpEndpointWithHostLabelOperationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpEndpointWithHostLabelOperation{}, middleware.After)
}

func addOpHttpRequestWithFloatLabelsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpHttpRequestWithFloatLabels{}, middleware.After)
}

func addOpHttpRequestWithGreedyLabelInPathValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpHttpRequestWithGreedyLabelInPath{}, middleware.After)
}

func addOpHttpRequestWithLabelsAndTimestampFormatValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpHttpRequestWithLabelsAndTimestampFormat{}, middleware.After)
}

func addOpHttpRequestWithLabelsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpHttpRequestWithLabels{}, middleware.After)
}

func validateOpConstantQueryStringInput(v *ConstantQueryStringInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ConstantQueryStringInput"}
	if v.Hello == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Hello"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpEndpointWithHostLabelHeaderOperationInput(v *EndpointWithHostLabelHeaderOperationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EndpointWithHostLabelHeaderOperationInput"}
	if v.AccountId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccountId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpEndpointWithHostLabelOperationInput(v *EndpointWithHostLabelOperationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EndpointWithHostLabelOperationInput"}
	if v.Label == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Label"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpHttpRequestWithFloatLabelsInput(v *HttpRequestWithFloatLabelsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "HttpRequestWithFloatLabelsInput"}
	if v.Float == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Float"))
	}
	if v.Double == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Double"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpHttpRequestWithGreedyLabelInPathInput(v *HttpRequestWithGreedyLabelInPathInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "HttpRequestWithGreedyLabelInPathInput"}
	if v.Foo == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Foo"))
	}
	if v.Baz == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Baz"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpHttpRequestWithLabelsAndTimestampFormatInput(v *HttpRequestWithLabelsAndTimestampFormatInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "HttpRequestWithLabelsAndTimestampFormatInput"}
	if v.MemberEpochSeconds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemberEpochSeconds"))
	}
	if v.MemberHttpDate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemberHttpDate"))
	}
	if v.MemberDateTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemberDateTime"))
	}
	if v.DefaultFormat == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DefaultFormat"))
	}
	if v.TargetEpochSeconds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetEpochSeconds"))
	}
	if v.TargetHttpDate == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetHttpDate"))
	}
	if v.TargetDateTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetDateTime"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpHttpRequestWithLabelsInput(v *HttpRequestWithLabelsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "HttpRequestWithLabelsInput"}
	if v.String_ == nil {
		invalidParams.Add(smithy.NewErrParamRequired("String_"))
	}
	if v.Short == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Short"))
	}
	if v.Integer == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Integer"))
	}
	if v.Long == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Long"))
	}
	if v.Float == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Float"))
	}
	if v.Double == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Double"))
	}
	if v.Boolean == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Boolean"))
	}
	if v.Timestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Timestamp"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
