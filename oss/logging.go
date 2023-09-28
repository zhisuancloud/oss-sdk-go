// Code generated by aws/logging_generate.go DO NOT EDIT.

package oss

// ClientLogMode represents the logging mode of SDK clients. The client logging mode is a bit-field where
// each bit is a flag that describes the logging behavior for one or more client components.
// The entire 64-bit group is reserved for later expansion by the SDK.
//
// Example: Setting ClientLogMode to enable logging of retries and requests
//
//	clientLogMode := oss.LogRetries | oss.LogRequest
//
// Example: Adding an additional log mode to an existing ClientLogMode value
//
//	clientLogMode |= oss.LogResponse
type ClientLogMode uint64

// Supported ClientLogMode bits that can be configured to toggle logging of specific SDK events.
const (
	LogSigning ClientLogMode = 1 << (64 - 1 - iota)
	LogRetries
	LogRequest
	LogRequestWithBody
	LogResponse
	LogResponseWithBody
	LogDeprecatedUsage
	LogRequestEventMessage
	LogResponseEventMessage
)

// IsSigning returns whether the Signing logging mode bit is set
func (m ClientLogMode) IsSigning() bool {
	return m&LogSigning != 0
}

// IsRetries returns whether the Retries logging mode bit is set
func (m ClientLogMode) IsRetries() bool {
	return m&LogRetries != 0
}

// IsRequest returns whether the Request logging mode bit is set
func (m ClientLogMode) IsRequest() bool {
	return m&LogRequest != 0
}

// IsRequestWithBody returns whether the RequestWithBody logging mode bit is set
func (m ClientLogMode) IsRequestWithBody() bool {
	return m&LogRequestWithBody != 0
}

// IsResponse returns whether the Response logging mode bit is set
func (m ClientLogMode) IsResponse() bool {
	return m&LogResponse != 0
}

// IsResponseWithBody returns whether the ResponseWithBody logging mode bit is set
func (m ClientLogMode) IsResponseWithBody() bool {
	return m&LogResponseWithBody != 0
}

// IsDeprecatedUsage returns whether the DeprecatedUsage logging mode bit is set
func (m ClientLogMode) IsDeprecatedUsage() bool {
	return m&LogDeprecatedUsage != 0
}

// IsRequestEventMessage returns whether the RequestEventMessage logging mode bit is set
func (m ClientLogMode) IsRequestEventMessage() bool {
	return m&LogRequestEventMessage != 0
}

// IsResponseEventMessage returns whether the ResponseEventMessage logging mode bit is set
func (m ClientLogMode) IsResponseEventMessage() bool {
	return m&LogResponseEventMessage != 0
}

// ClearSigning clears the Signing logging mode bit
func (m *ClientLogMode) ClearSigning() {
	*m &^= LogSigning
}

// ClearRetries clears the Retries logging mode bit
func (m *ClientLogMode) ClearRetries() {
	*m &^= LogRetries
}

// ClearRequest clears the Request logging mode bit
func (m *ClientLogMode) ClearRequest() {
	*m &^= LogRequest
}

// ClearRequestWithBody clears the RequestWithBody logging mode bit
func (m *ClientLogMode) ClearRequestWithBody() {
	*m &^= LogRequestWithBody
}

// ClearResponse clears the Response logging mode bit
func (m *ClientLogMode) ClearResponse() {
	*m &^= LogResponse
}

// ClearResponseWithBody clears the ResponseWithBody logging mode bit
func (m *ClientLogMode) ClearResponseWithBody() {
	*m &^= LogResponseWithBody
}

// ClearDeprecatedUsage clears the DeprecatedUsage logging mode bit
func (m *ClientLogMode) ClearDeprecatedUsage() {
	*m &^= LogDeprecatedUsage
}

// ClearRequestEventMessage clears the RequestEventMessage logging mode bit
func (m *ClientLogMode) ClearRequestEventMessage() {
	*m &^= LogRequestEventMessage
}

// ClearResponseEventMessage clears the ResponseEventMessage logging mode bit
func (m *ClientLogMode) ClearResponseEventMessage() {
	*m &^= LogResponseEventMessage
}
