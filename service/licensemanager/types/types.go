// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Describes automated discovery.
type AutomatedDiscoveryInformation struct {

	// Time that automated discovery last ran.
	LastRunTime *time.Time

	noSmithyDocumentSerde
}

// Details about a borrow configuration.
type BorrowConfiguration struct {

	// Indicates whether early check-ins are allowed.
	//
	// This member is required.
	AllowEarlyCheckIn *bool

	// Maximum time for the borrow configuration, in minutes.
	//
	// This member is required.
	MaxTimeToLiveInMinutes *int32

	noSmithyDocumentSerde
}

// Details about license consumption.
type ConsumedLicenseSummary struct {

	// Number of licenses consumed by the resource.
	ConsumedLicenses *int64

	// Resource type of the resource consuming a license.
	ResourceType ResourceType

	noSmithyDocumentSerde
}

// Details about a consumption configuration.
type ConsumptionConfiguration struct {

	// Details about a borrow configuration.
	BorrowConfiguration *BorrowConfiguration

	// Details about a provisional configuration.
	ProvisionalConfiguration *ProvisionalConfiguration

	// Renewal frequency.
	RenewType RenewType

	noSmithyDocumentSerde
}

// Describes a time range, in ISO8601-UTC format.
type DatetimeRange struct {

	// Start of the time range.
	//
	// This member is required.
	Begin *string

	// End of the time range.
	End *string

	noSmithyDocumentSerde
}

// Describes a resource entitled for use with a license.
type Entitlement struct {

	// Entitlement name.
	//
	// This member is required.
	Name *string

	// Entitlement unit.
	//
	// This member is required.
	Unit EntitlementUnit

	// Indicates whether check-ins are allowed.
	AllowCheckIn *bool

	// Maximum entitlement count. Use if the unit is not None.
	MaxCount *int64

	// Indicates whether overages are allowed.
	Overage *bool

	// Entitlement resource. Use only if the unit is None.
	Value *string

	noSmithyDocumentSerde
}

// Data associated with an entitlement resource.
type EntitlementData struct {

	// Entitlement data name.
	//
	// This member is required.
	Name *string

	// Entitlement data unit.
	//
	// This member is required.
	Unit EntitlementDataUnit

	// Entitlement data value.
	Value *string

	noSmithyDocumentSerde
}

// Usage associated with an entitlement resource.
type EntitlementUsage struct {

	// Resource usage consumed.
	//
	// This member is required.
	ConsumedValue *string

	// Entitlement usage name.
	//
	// This member is required.
	Name *string

	// Entitlement usage unit.
	//
	// This member is required.
	Unit EntitlementDataUnit

	// Maximum entitlement usage count.
	MaxCount *string

	noSmithyDocumentSerde
}

// A filter name and value pair that is used to return more specific results from a
// describe operation. Filters can be used to match a set of resources by specific
// criteria, such as tags, attributes, or IDs.
type Filter struct {

	// Name of the filter. Filter names are case-sensitive.
	Name *string

	// Filter values. Filter values are case-sensitive.
	Values []string

	noSmithyDocumentSerde
}

// Describes a grant.
type Grant struct {

	// Amazon Resource Name (ARN) of the grant.
	//
	// This member is required.
	GrantArn *string

	// Grant name.
	//
	// This member is required.
	GrantName *string

	// Grant status.
	//
	// This member is required.
	GrantStatus GrantStatus

	// Granted operations.
	//
	// This member is required.
	GrantedOperations []AllowedOperation

	// The grantee principal ARN.
	//
	// This member is required.
	GranteePrincipalArn *string

	// Home Region of the grant.
	//
	// This member is required.
	HomeRegion *string

	// License ARN.
	//
	// This member is required.
	LicenseArn *string

	// Parent ARN.
	//
	// This member is required.
	ParentArn *string

	// Grant version.
	//
	// This member is required.
	Version *string

	// Grant status reason.
	StatusReason *string

	noSmithyDocumentSerde
}

// Describes a license that is granted to a grantee.
type GrantedLicense struct {

	// Granted license beneficiary.
	Beneficiary *string

	// Configuration for consumption of the license.
	ConsumptionConfiguration *ConsumptionConfiguration

	// Creation time of the granted license.
	CreateTime *string

	// License entitlements.
	Entitlements []Entitlement

	// Home Region of the granted license.
	HomeRegion *string

	// Granted license issuer.
	Issuer *IssuerDetails

	// Amazon Resource Name (ARN) of the license.
	LicenseArn *string

	// Granted license metadata.
	LicenseMetadata []Metadata

	// License name.
	LicenseName *string

	// Product name.
	ProductName *string

	// Product SKU.
	ProductSKU *string

	// Granted license received metadata.
	ReceivedMetadata *ReceivedMetadata

	// Granted license status.
	Status LicenseStatus

	// Date and time range during which the granted license is valid, in ISO8601-UTC
	// format.
	Validity *DatetimeRange

	// Version of the granted license.
	Version *string

	noSmithyDocumentSerde
}

// An inventory filter.
type InventoryFilter struct {

	// Condition of the filter.
	//
	// This member is required.
	Condition InventoryFilterCondition

	// Name of the filter.
	//
	// This member is required.
	Name *string

	// Value of the filter.
	Value *string

	noSmithyDocumentSerde
}

// Details about the issuer of a license.
type Issuer struct {

	// Issuer name.
	//
	// This member is required.
	Name *string

	// Asymmetric KMS key from Key Management Service. The KMS key must have a key
	// usage of sign and verify, and support the RSASSA-PSS SHA-256 signing algorithm.
	SignKey *string

	noSmithyDocumentSerde
}

// Details associated with the issuer of a license.
type IssuerDetails struct {

	// Issuer key fingerprint.
	KeyFingerprint *string

	// Issuer name.
	Name *string

	// Asymmetric KMS key from Key Management Service. The KMS key must have a key
	// usage of sign and verify, and support the RSASSA-PSS SHA-256 signing algorithm.
	SignKey *string

	noSmithyDocumentSerde
}

// Software license that is managed in License Manager.
type License struct {

	// License beneficiary.
	Beneficiary *string

	// Configuration for consumption of the license.
	ConsumptionConfiguration *ConsumptionConfiguration

	// License creation time.
	CreateTime *string

	// License entitlements.
	Entitlements []Entitlement

	// Home Region of the license.
	HomeRegion *string

	// License issuer.
	Issuer *IssuerDetails

	// Amazon Resource Name (ARN) of the license.
	LicenseArn *string

	// License metadata.
	LicenseMetadata []Metadata

	// License name.
	LicenseName *string

	// Product name.
	ProductName *string

	// Product SKU.
	ProductSKU *string

	// License status.
	Status LicenseStatus

	// Date and time range during which the license is valid, in ISO8601-UTC format.
	Validity *DatetimeRange

	// License version.
	Version *string

	noSmithyDocumentSerde
}

// A license configuration is an abstraction of a customer license agreement that
// can be consumed and enforced by License Manager. Components include
// specifications for the license type (licensing by instance, socket, CPU, or
// vCPU), allowed tenancy (shared tenancy, Dedicated Instance, Dedicated Host, or
// all of these), host affinity (how long a VM must be associated with a host), and
// the number of licenses purchased and used.
type LicenseConfiguration struct {

	// Automated discovery information.
	AutomatedDiscoveryInformation *AutomatedDiscoveryInformation

	// Summaries for licenses consumed by various resources.
	ConsumedLicenseSummaryList []ConsumedLicenseSummary

	// Number of licenses consumed.
	ConsumedLicenses *int64

	// Description of the license configuration.
	Description *string

	// When true, disassociates a resource when software is uninstalled.
	DisassociateWhenNotFound *bool

	// Amazon Resource Name (ARN) of the license configuration.
	LicenseConfigurationArn *string

	// Unique ID of the license configuration.
	LicenseConfigurationId *string

	// Number of licenses managed by the license configuration.
	LicenseCount *int64

	// Number of available licenses as a hard limit.
	LicenseCountHardLimit *bool

	// Dimension to use to track the license inventory.
	LicenseCountingType LicenseCountingType

	// License rules.
	LicenseRules []string

	// Summaries for managed resources.
	ManagedResourceSummaryList []ManagedResourceSummary

	// Name of the license configuration.
	Name *string

	// Account ID of the license configuration's owner.
	OwnerAccountId *string

	// Product information.
	ProductInformationList []ProductInformation

	// Status of the license configuration.
	Status *string

	noSmithyDocumentSerde
}

// Describes an association with a license configuration.
type LicenseConfigurationAssociation struct {

	// Scope of AMI associations. The possible value is cross-account.
	AmiAssociationScope *string

	// Time when the license configuration was associated with the resource.
	AssociationTime *time.Time

	// Amazon Resource Name (ARN) of the resource.
	ResourceArn *string

	// ID of the Amazon Web Services account that owns the resource consuming licenses.
	ResourceOwnerId *string

	// Type of server resource.
	ResourceType ResourceType

	noSmithyDocumentSerde
}

// Details about the usage of a resource associated with a license configuration.
type LicenseConfigurationUsage struct {

	// Time when the license configuration was initially associated with the resource.
	AssociationTime *time.Time

	// Number of licenses consumed by the resource.
	ConsumedLicenses *int64

	// Amazon Resource Name (ARN) of the resource.
	ResourceArn *string

	// ID of the account that owns the resource.
	ResourceOwnerId *string

	// Status of the resource.
	ResourceStatus *string

	// Type of resource.
	ResourceType ResourceType

	noSmithyDocumentSerde
}

// Information about a license type conversion task.
type LicenseConversionContext struct {

	// The Usage operation value that corresponds to the license type you are
	// converting your resource from. For more information about which platforms
	// correspond to which usage operation values see Sample data: usage operation by
	// platform
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/billing-info-fields.html#billing-info)
	UsageOperation *string

	noSmithyDocumentSerde
}

// Information about a license type conversion task.
type LicenseConversionTask struct {

	// Information about the license type this conversion task converted to.
	DestinationLicenseContext *LicenseConversionContext

	// The time the conversion task was completed.
	EndTime *time.Time

	// The ID of the license type conversion task.
	LicenseConversionTaskId *string

	// The time the usage operation value of the resource was changed.
	LicenseConversionTime *time.Time

	// The Amazon Resource Name (ARN) of the resource associated with the license type
	// conversion task.
	ResourceArn *string

	// Information about the license type this conversion task converted from.
	SourceLicenseContext *LicenseConversionContext

	// The time the conversion task was started at.
	StartTime *time.Time

	// The status of the conversion task.
	Status LicenseConversionTaskStatus

	// The status message for the conversion task.
	StatusMessage *string

	noSmithyDocumentSerde
}

// Describes the failure of a license operation.
type LicenseOperationFailure struct {

	// Error message.
	ErrorMessage *string

	// Failure time.
	FailureTime *time.Time

	// Reserved.
	MetadataList []Metadata

	// Name of the operation.
	OperationName *string

	// The requester is "License Manager Automated Discovery".
	OperationRequestedBy *string

	// Amazon Resource Name (ARN) of the resource.
	ResourceArn *string

	// ID of the Amazon Web Services account that owns the resource.
	ResourceOwnerId *string

	// Resource type.
	ResourceType ResourceType

	noSmithyDocumentSerde
}

// Details for associating a license configuration with a resource.
type LicenseSpecification struct {

	// Amazon Resource Name (ARN) of the license configuration.
	//
	// This member is required.
	LicenseConfigurationArn *string

	// Scope of AMI associations. The possible value is cross-account.
	AmiAssociationScope *string

	noSmithyDocumentSerde
}

// Describes the entitlement usage associated with a license.
type LicenseUsage struct {

	// License entitlement usages.
	EntitlementUsages []EntitlementUsage

	noSmithyDocumentSerde
}

// Summary information about a managed resource.
type ManagedResourceSummary struct {

	// Number of resources associated with licenses.
	AssociationCount *int64

	// Type of resource associated with a license.
	ResourceType ResourceType

	noSmithyDocumentSerde
}

// Describes key/value pairs.
type Metadata struct {

	// The key name.
	Name *string

	// The value.
	Value *string

	noSmithyDocumentSerde
}

// Configuration information for Organizations.
type OrganizationConfiguration struct {

	// Enables Organizations integration.
	//
	// This member is required.
	EnableIntegration bool

	noSmithyDocumentSerde
}

// Describes product information for a license configuration.
type ProductInformation struct {

	// A Product information filter consists of a ProductInformationFilterComparator
	// which is a logical operator, a ProductInformationFilterName which specifies the
	// type of filter being declared, and a ProductInformationFilterValue that
	// specifies the value to filter on. Accepted values for
	// ProductInformationFilterName are listed here along with descriptions and valid
	// options for ProductInformationFilterComparator. The following filters and are
	// supported when the resource type is SSM_MANAGED:
	//
	// * Application Name - The name
	// of the application. Logical operator is EQUALS.
	//
	// * Application Publisher - The
	// publisher of the application. Logical operator is EQUALS.
	//
	// * Application Version
	// - The version of the application. Logical operator is EQUALS.
	//
	// * Platform Name -
	// The name of the platform. Logical operator is EQUALS.
	//
	// * Platform Type - The
	// platform type. Logical operator is EQUALS.
	//
	// * Tag:key - The key of a tag
	// attached to an Amazon Web Services resource you wish to exclude from automated
	// discovery. Logical operator is NOT_EQUALS. The key for your tag must be appended
	// to Tag: following the example: Tag:name-of-your-key.
	// ProductInformationFilterValue is optional if you are not using values for the
	// key.
	//
	// * AccountId - The 12-digit ID of an Amazon Web Services account you wish
	// to exclude from automated discovery. Logical operator is NOT_EQUALS.
	//
	// * License
	// Included - The type of license included. Logical operators are EQUALS and
	// NOT_EQUALS. Possible values are: sql-server-enterprise | sql-server-standard |
	// sql-server-web | windows-server-datacenter.
	//
	// The following filters and logical
	// operators are supported when the resource type is RDS:
	//
	// * Engine Edition - The
	// edition of the database engine. Logical operator is EQUALS. Possible values are:
	// oracle-ee | oracle-se | oracle-se1 | oracle-se2.
	//
	// * License Pack - The license
	// pack. Logical operator is EQUALS. Possible values are: data guard | diagnostic
	// pack sqlt | tuning pack sqlt | ols | olap.
	//
	// This member is required.
	ProductInformationFilterList []ProductInformationFilter

	// Resource type. The possible values are SSM_MANAGED | RDS.
	//
	// This member is required.
	ResourceType *string

	noSmithyDocumentSerde
}

// Describes product information filters.
type ProductInformationFilter struct {

	// Logical operator.
	//
	// This member is required.
	ProductInformationFilterComparator *string

	// Filter name.
	//
	// This member is required.
	ProductInformationFilterName *string

	// Filter value.
	ProductInformationFilterValue []string

	noSmithyDocumentSerde
}

// Details about a provisional configuration.
type ProvisionalConfiguration struct {

	// Maximum time for the provisional configuration, in minutes.
	//
	// This member is required.
	MaxTimeToLiveInMinutes *int32

	noSmithyDocumentSerde
}

// Metadata associated with received licenses and grants.
type ReceivedMetadata struct {

	// Allowed operations.
	AllowedOperations []AllowedOperation

	// Received status.
	ReceivedStatus ReceivedStatus

	// Received status reason.
	ReceivedStatusReason *string

	noSmithyDocumentSerde
}

// Details of the license configuration that this generator reports on.
type ReportContext struct {

	// Amazon Resource Name (ARN) of the license configuration that this generator
	// reports on.
	//
	// This member is required.
	LicenseConfigurationArns []string

	noSmithyDocumentSerde
}

// Details about how frequently reports are generated.
type ReportFrequency struct {

	// Time period between each report. The period can be daily, weekly, or monthly.
	Period ReportFrequencyType

	// Number of times within the frequency period that a report is generated. The only
	// supported value is 1.
	Value *int32

	noSmithyDocumentSerde
}

// Describe the details of a report generator.
type ReportGenerator struct {

	// Time the report was created.
	CreateTime *string

	// Description of the report generator.
	Description *string

	// Time the last report was generated at.
	LastReportGenerationTime *string

	// Failure message for the last report generation attempt.
	LastRunFailureReason *string

	// Status of the last report generation attempt.
	LastRunStatus *string

	// Amazon Resource Name (ARN) of the report generator.
	LicenseManagerReportGeneratorArn *string

	// License configuration type for this generator.
	ReportContext *ReportContext

	// The Amazon Web Services account ID used to create the report generator.
	ReportCreatorAccount *string

	// Details about how frequently reports are generated.
	ReportFrequency *ReportFrequency

	// Name of the report generator.
	ReportGeneratorName *string

	// Type of reports that are generated.
	ReportType []ReportType

	// Details of the S3 bucket that report generator reports are published to.
	S3Location *S3Location

	// Tags associated with the report generator.
	Tags []Tag

	noSmithyDocumentSerde
}

// Details about a resource.
type ResourceInventory struct {

	// Platform of the resource.
	Platform *string

	// Platform version of the resource in the inventory.
	PlatformVersion *string

	// Amazon Resource Name (ARN) of the resource.
	ResourceArn *string

	// ID of the resource.
	ResourceId *string

	// ID of the account that owns the resource.
	ResourceOwningAccountId *string

	// Type of resource.
	ResourceType ResourceType

	noSmithyDocumentSerde
}

// Details of the S3 bucket that report generator reports are published to.
type S3Location struct {

	// Name of the S3 bucket reports are published to.
	Bucket *string

	// Prefix of the S3 bucket reports are published to.
	KeyPrefix *string

	noSmithyDocumentSerde
}

// Details about a tag for a license configuration.
type Tag struct {

	// Tag key.
	Key *string

	// Tag value.
	Value *string

	noSmithyDocumentSerde
}

// Describes a token.
type TokenData struct {

	// Token expiration time, in ISO8601-UTC format.
	ExpirationTime *string

	// Amazon Resource Name (ARN) of the license.
	LicenseArn *string

	// Amazon Resource Names (ARN) of the roles included in the token.
	RoleArns []string

	// Token status. The possible values are AVAILABLE and DELETED.
	Status *string

	// Token ID.
	TokenId *string

	// Data specified by the caller.
	TokenProperties []string

	// Type of token generated. The supported value is REFRESH_TOKEN.
	TokenType *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
