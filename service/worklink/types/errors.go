// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The service is temporarily unavailable.
type InternalServerErrorException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InternalServerErrorException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerErrorException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerErrorException) ErrorCode() string             { return "InternalServerErrorException" }
func (e *InternalServerErrorException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// The request is not valid.
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

// The resource already exists.
type ResourceAlreadyExistsException struct {
	Message *string

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

// The requested resource was not found.
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

// The number of requests exceeds the limit.
type TooManyRequestsException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *TooManyRequestsException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *TooManyRequestsException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *TooManyRequestsException) ErrorCode() string             { return "TooManyRequestsException" }
func (e *TooManyRequestsException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// You are not authorized to perform this action.
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
