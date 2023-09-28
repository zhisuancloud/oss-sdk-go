// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type FixedPosition string

// Enum values for FixedPosition
const (
	FixedPositionFirst FixedPosition = "first"
)

// Values returns all known values for FixedPosition. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FixedPosition) Values() []FixedPosition {
	return []FixedPosition{
		"first",
	}
}

type FormActionType string

// Enum values for FormActionType
const (
	FormActionTypeCreate FormActionType = "create"
	FormActionTypeUpdate FormActionType = "update"
)

// Values returns all known values for FormActionType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FormActionType) Values() []FormActionType {
	return []FormActionType{
		"create",
		"update",
	}
}

type FormButtonsPosition string

// Enum values for FormButtonsPosition
const (
	FormButtonsPositionTop          FormButtonsPosition = "top"
	FormButtonsPositionBottom       FormButtonsPosition = "bottom"
	FormButtonsPositionTopAndBottom FormButtonsPosition = "top_and_bottom"
)

// Values returns all known values for FormButtonsPosition. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FormButtonsPosition) Values() []FormButtonsPosition {
	return []FormButtonsPosition{
		"top",
		"bottom",
		"top_and_bottom",
	}
}

type FormDataSourceType string

// Enum values for FormDataSourceType
const (
	// Will use a provided Amplify DataStore enabled API
	FormDataSourceTypeDatastore FormDataSourceType = "DataStore"
	// Will use passed in hooks to use when creating a form from scratch
	FormDataSourceTypeCustom FormDataSourceType = "Custom"
)

// Values returns all known values for FormDataSourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FormDataSourceType) Values() []FormDataSourceType {
	return []FormDataSourceType{
		"DataStore",
		"Custom",
	}
}

type SortDirection string

// Values returns all known values for SortDirection. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SortDirection) Values() []SortDirection {
	return []SortDirection{
		"ASC",
		"DESC",
	}
}

type TokenProviders string

// Enum values for TokenProviders
const (
	// The figma token provider.
	TokenProvidersFigma TokenProviders = "figma"
)

// Values returns all known values for TokenProviders. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TokenProviders) Values() []TokenProviders {
	return []TokenProviders{
		"figma",
	}
}