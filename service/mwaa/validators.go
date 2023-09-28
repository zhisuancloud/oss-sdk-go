// Code generated by smithy-go-codegen DO NOT EDIT.

package mwaa

import (
	"context"
	"fmt"
	"oss-sdk-go/service/mwaa/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateCliToken struct {
}

func (*validateOpCreateCliToken) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateCliToken) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateCliTokenInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateCliTokenInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateEnvironment struct {
}

func (*validateOpCreateEnvironment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateEnvironment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateEnvironmentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateEnvironmentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateWebLoginToken struct {
}

func (*validateOpCreateWebLoginToken) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateWebLoginToken) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateWebLoginTokenInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateWebLoginTokenInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteEnvironment struct {
}

func (*validateOpDeleteEnvironment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteEnvironment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteEnvironmentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteEnvironmentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetEnvironment struct {
}

func (*validateOpGetEnvironment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetEnvironment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetEnvironmentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetEnvironmentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPublishMetrics struct {
}

func (*validateOpPublishMetrics) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPublishMetrics) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PublishMetricsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPublishMetricsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateEnvironment struct {
}

func (*validateOpUpdateEnvironment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateEnvironment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateEnvironmentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateEnvironmentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateCliTokenValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateCliToken{}, middleware.After)
}

func addOpCreateEnvironmentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateEnvironment{}, middleware.After)
}

func addOpCreateWebLoginTokenValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateWebLoginToken{}, middleware.After)
}

func addOpDeleteEnvironmentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteEnvironment{}, middleware.After)
}

func addOpGetEnvironmentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetEnvironment{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpPublishMetricsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPublishMetrics{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateEnvironmentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateEnvironment{}, middleware.After)
}

func validateDimension(v *types.Dimension) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Dimension"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateDimensions(v []types.Dimension) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Dimensions"}
	for i := range v {
		if err := validateDimension(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateLoggingConfigurationInput(v *types.LoggingConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "LoggingConfigurationInput"}
	if v.DagProcessingLogs != nil {
		if err := validateModuleLoggingConfigurationInput(v.DagProcessingLogs); err != nil {
			invalidParams.AddNested("DagProcessingLogs", err.(smithy.InvalidParamsError))
		}
	}
	if v.SchedulerLogs != nil {
		if err := validateModuleLoggingConfigurationInput(v.SchedulerLogs); err != nil {
			invalidParams.AddNested("SchedulerLogs", err.(smithy.InvalidParamsError))
		}
	}
	if v.WebserverLogs != nil {
		if err := validateModuleLoggingConfigurationInput(v.WebserverLogs); err != nil {
			invalidParams.AddNested("WebserverLogs", err.(smithy.InvalidParamsError))
		}
	}
	if v.WorkerLogs != nil {
		if err := validateModuleLoggingConfigurationInput(v.WorkerLogs); err != nil {
			invalidParams.AddNested("WorkerLogs", err.(smithy.InvalidParamsError))
		}
	}
	if v.TaskLogs != nil {
		if err := validateModuleLoggingConfigurationInput(v.TaskLogs); err != nil {
			invalidParams.AddNested("TaskLogs", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMetricData(v []types.MetricDatum) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MetricData"}
	for i := range v {
		if err := validateMetricDatum(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMetricDatum(v *types.MetricDatum) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MetricDatum"}
	if v.MetricName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MetricName"))
	}
	if v.Timestamp == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Timestamp"))
	}
	if v.Dimensions != nil {
		if err := validateDimensions(v.Dimensions); err != nil {
			invalidParams.AddNested("Dimensions", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateModuleLoggingConfigurationInput(v *types.ModuleLoggingConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ModuleLoggingConfigurationInput"}
	if v.Enabled == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Enabled"))
	}
	if len(v.LogLevel) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("LogLevel"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUpdateNetworkConfigurationInput(v *types.UpdateNetworkConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateNetworkConfigurationInput"}
	if v.SecurityGroupIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecurityGroupIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateCliTokenInput(v *CreateCliTokenInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateCliTokenInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateEnvironmentInput(v *CreateEnvironmentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateEnvironmentInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.ExecutionRoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ExecutionRoleArn"))
	}
	if v.SourceBucketArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourceBucketArn"))
	}
	if v.DagS3Path == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DagS3Path"))
	}
	if v.NetworkConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkConfiguration"))
	}
	if v.LoggingConfiguration != nil {
		if err := validateLoggingConfigurationInput(v.LoggingConfiguration); err != nil {
			invalidParams.AddNested("LoggingConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateWebLoginTokenInput(v *CreateWebLoginTokenInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateWebLoginTokenInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteEnvironmentInput(v *DeleteEnvironmentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteEnvironmentInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetEnvironmentInput(v *GetEnvironmentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetEnvironmentInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPublishMetricsInput(v *PublishMetricsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PublishMetricsInput"}
	if v.EnvironmentName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentName"))
	}
	if v.MetricData == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MetricData"))
	} else if v.MetricData != nil {
		if err := validateMetricData(v.MetricData); err != nil {
			invalidParams.AddNested("MetricData", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateEnvironmentInput(v *UpdateEnvironmentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateEnvironmentInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.NetworkConfiguration != nil {
		if err := validateUpdateNetworkConfigurationInput(v.NetworkConfiguration); err != nil {
			invalidParams.AddNested("NetworkConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.LoggingConfiguration != nil {
		if err := validateLoggingConfigurationInput(v.LoggingConfiguration); err != nil {
			invalidParams.AddNested("LoggingConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
