// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Your request has conflicting operations. This can occur if you're trying to
// perform more than one operation on the same resource at the same time.
type ConflictingOperationException struct {
	Message *string

	ResourceId  *string
	ResourceArn *string

	noSmithyDocumentSerde
}

func (e *ConflictingOperationException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConflictingOperationException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConflictingOperationException) ErrorCode() string             { return "ConflictingOperationException" }
func (e *ConflictingOperationException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// IoT SiteWise can't process your request right now. Try again later.
type InternalFailureException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InternalFailureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalFailureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalFailureException) ErrorCode() string             { return "InternalFailureException" }
func (e *InternalFailureException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The request isn't valid. This can occur if your request contains malformed JSON
// or unsupported characters. Check your request and try again.
type InvalidRequestException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string             { return "InvalidRequestException" }
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You've reached the limit for a resource. For example, this can occur if you're
// trying to associate more than the allowed number of child assets or attempting
// to create more than the allowed number of properties for an asset model. For
// more information, see Quotas
// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the
// IoT SiteWise User Guide.
type LimitExceededException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The resource already exists.
type ResourceAlreadyExistsException struct {
	Message *string

	ResourceId  *string
	ResourceArn *string

	noSmithyDocumentSerde
}

func (e *ResourceAlreadyExistsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceAlreadyExistsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceAlreadyExistsException) ErrorCode() string             { return "ResourceAlreadyExistsException" }
func (e *ResourceAlreadyExistsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested resource can't be found.
type ResourceNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested service is unavailable.
type ServiceUnavailableException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ServiceUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceUnavailableException) ErrorCode() string             { return "ServiceUnavailableException" }
func (e *ServiceUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Your request exceeded a rate limit. For example, you might have exceeded the
// number of IoT SiteWise assets that can be created per second, the allowed number
// of messages per second, and so on. For more information, see Quotas
// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/quotas.html) in the
// IoT SiteWise User Guide.
type ThrottlingException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottlingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottlingException) ErrorCode() string             { return "ThrottlingException" }
func (e *ThrottlingException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You've reached the limit for the number of tags allowed for a resource. For more
// information, see Tag naming limits and requirements
// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html#tag-conventions)
// in the Amazon Web Services General Reference.
type TooManyTagsException struct {
	Message *string

	ResourceName *string

	noSmithyDocumentSerde
}

func (e *TooManyTagsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyTagsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyTagsException) ErrorCode() string             { return "TooManyTagsException" }
func (e *TooManyTagsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You are not authorized.
type UnauthorizedException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *UnauthorizedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnauthorizedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnauthorizedException) ErrorCode() string             { return "UnauthorizedException" }
func (e *UnauthorizedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
