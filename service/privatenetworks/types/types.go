// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Information about an address.
type Address struct {

	// The city for this address.
	//
	// This member is required.
	City *string

	// The country for this address.
	//
	// This member is required.
	Country *string

	// The recipient's name for this address.
	//
	// This member is required.
	Name *string

	// The postal code for this address.
	//
	// This member is required.
	PostalCode *string

	// The state or province for this address.
	//
	// This member is required.
	StateOrProvince *string

	// The first line of the street address.
	//
	// This member is required.
	Street1 *string

	// The company name for this address.
	Company *string

	// The phone number for this address.
	PhoneNumber *string

	// The second line of the street address.
	Street2 *string

	// The third line of the street address.
	Street3 *string

	noSmithyDocumentSerde
}

// Information about a subscriber of a device that can use a network.
type DeviceIdentifier struct {

	// The creation time of this device identifier.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of the device identifier.
	DeviceIdentifierArn *string

	// The Integrated Circuit Card Identifier of the device identifier.
	Iccid *string

	// The International Mobile Subscriber Identity of the device identifier.
	Imsi *string

	// The Amazon Resource Name (ARN) of the network on which the device identifier
	// appears.
	NetworkArn *string

	// The Amazon Resource Name (ARN) of the order used to purchase the device
	// identifier.
	OrderArn *string

	// The status of the device identifier.
	Status DeviceIdentifierStatus

	// The Amazon Resource Name (ARN) of the traffic group to which the device
	// identifier belongs.
	TrafficGroupArn *string

	// The vendor of the device identifier.
	Vendor *string

	noSmithyDocumentSerde
}

// Information about a name/value pair.
type NameValuePair struct {

	// The name of the pair.
	//
	// This member is required.
	Name *string

	// The value of the pair.
	Value *string

	noSmithyDocumentSerde
}

// Information about a network.
type Network struct {

	// The Amazon Resource Name (ARN) of the network.
	//
	// This member is required.
	NetworkArn *string

	// The name of the network.
	//
	// This member is required.
	NetworkName *string

	// The status of the network.
	//
	// This member is required.
	Status NetworkStatus

	// The creation time of the network.
	CreatedAt *time.Time

	// The description of the network.
	Description *string

	// The status reason of the network.
	StatusReason *string

	noSmithyDocumentSerde
}

// Information about a network resource.
type NetworkResource struct {

	// The attributes of the network resource.
	Attributes []NameValuePair

	// The creation time of the network resource.
	CreatedAt *time.Time

	// The description of the network resource.
	Description *string

	// The health of the network resource.
	Health HealthStatus

	// The model of the network resource.
	Model *string

	// The Amazon Resource Name (ARN) of the network on which this network resource
	// appears.
	NetworkArn *string

	// The Amazon Resource Name (ARN) of the network resource.
	NetworkResourceArn *string

	// The Amazon Resource Name (ARN) of the network site on which this network
	// resource appears.
	NetworkSiteArn *string

	// The Amazon Resource Name (ARN) of the order used to purchase this network
	// resource.
	OrderArn *string

	// The position of the network resource.
	Position *Position

	// The serial number of the network resource.
	SerialNumber *string

	// The status of the network resource.
	Status NetworkResourceStatus

	// The status reason of the network resource.
	StatusReason *string

	// The type of the network resource.
	Type NetworkResourceType

	// The vendor of the network resource.
	Vendor *string

	noSmithyDocumentSerde
}

// Information about a network resource definition.
type NetworkResourceDefinition struct {

	// The count in the network resource definition.
	//
	// This member is required.
	Count *int32

	// The type in the network resource definition.
	//
	// This member is required.
	Type NetworkResourceDefinitionType

	// The options in the network resource definition.
	Options []NameValuePair

	noSmithyDocumentSerde
}

// Information about a network site.
type NetworkSite struct {

	// The Amazon Resource Name (ARN) of the network to which the network site belongs.
	//
	// This member is required.
	NetworkArn *string

	// The Amazon Resource Name (ARN) of the network site.
	//
	// This member is required.
	NetworkSiteArn *string

	// The name of the network site.
	//
	// This member is required.
	NetworkSiteName *string

	// The status of the network site.
	//
	// This member is required.
	Status NetworkSiteStatus

	// The parent Availability Zone for the network site.
	AvailabilityZone *string

	// The parent Availability Zone ID for the network site.
	AvailabilityZoneId *string

	// The creation time of the network site.
	CreatedAt *time.Time

	// The current plan of the network site.
	CurrentPlan *SitePlan

	// The description of the network site.
	Description *string

	// The pending plan of the network site.
	PendingPlan *SitePlan

	// The status reason of the network site.
	StatusReason *string

	noSmithyDocumentSerde
}

// Information about an order.
type Order struct {

	// The acknowledgement status of the order.
	AcknowledgmentStatus AcknowledgmentStatus

	// The creation time of the order.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of the network associated with this order.
	NetworkArn *string

	// The Amazon Resource Name (ARN) of the network site associated with this order.
	NetworkSiteArn *string

	// The Amazon Resource Name (ARN) of the order.
	OrderArn *string

	// The shipping address of the order.
	ShippingAddress *Address

	// The tracking information of the order.
	TrackingInformation []TrackingInformation

	noSmithyDocumentSerde
}

// Information about a position.
type Position struct {

	// The elevation of the equipment at this position.
	Elevation *float64

	// The reference point from which elevation is reported.
	ElevationReference ElevationReference

	// The units used to measure the elevation of the position.
	ElevationUnit ElevationUnit

	// The latitude of the position.
	Latitude *float64

	// The longitude of the position.
	Longitude *float64

	noSmithyDocumentSerde
}

// Information about a site plan.
type SitePlan struct {

	// The options of the plan.
	Options []NameValuePair

	// The resource definitions of the plan.
	ResourceDefinitions []NetworkResourceDefinition

	noSmithyDocumentSerde
}

// Information about tracking a shipment.
type TrackingInformation struct {

	// The tracking number of the shipment.
	TrackingNumber *string

	noSmithyDocumentSerde
}

// Information about a field that failed validation.
type ValidationExceptionField struct {

	// The message about the validation failure.
	//
	// This member is required.
	Message *string

	// The field name that failed validation.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
