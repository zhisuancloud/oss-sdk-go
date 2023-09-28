// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// A structure that encapsulates a signaling channel's metadata and properties.
type ChannelInfo struct {

	// The Amazon Resource Name (ARN) of the signaling channel.
	ChannelARN *string

	// The name of the signaling channel.
	ChannelName *string

	// Current status of the signaling channel.
	ChannelStatus Status

	// The type of the signaling channel.
	ChannelType ChannelType

	// The time at which the signaling channel was created.
	CreationTime *time.Time

	// A structure that contains the configuration for the SINGLE_MASTER channel type.
	SingleMasterConfiguration *SingleMasterConfiguration

	// The current version of the signaling channel.
	Version *string

	noSmithyDocumentSerde
}

// An optional input parameter for the ListSignalingChannels API. When this
// parameter is specified while invoking ListSignalingChannels, the API returns
// only the channels that satisfy a condition specified in ChannelNameCondition.
type ChannelNameCondition struct {

	// A comparison operator. Currently, you can only specify the BEGINS_WITH operator,
	// which finds signaling channels whose names begin with a given prefix.
	ComparisonOperator ComparisonOperator

	// A value to compare.
	ComparisonValue *string

	noSmithyDocumentSerde
}

// The structure that contains the information required for the KVS images
// delivery. If null, the configuration will be deleted from the stream.
type ImageGenerationConfiguration struct {

	// The structure that contains the information required to deliver images to a
	// customer.
	//
	// This member is required.
	DestinationConfig *ImageGenerationDestinationConfig

	// The accepted image format.
	//
	// This member is required.
	Format Format

	// The origin of the Server or Producer timestamps to use to generate the images.
	//
	// This member is required.
	ImageSelectorType ImageSelectorType

	// The time interval in milliseconds (ms) at which the images need to be generated
	// from the stream. The minimum value that can be provided is 33 ms, because a
	// camera that generates content at 30 FPS would create a frame every 33.3 ms. If
	// the timestamp range is less than the sampling interval, the Image from the
	// StartTimestamp will be returned if available.
	//
	// This member is required.
	SamplingInterval *int32

	// Indicates whether the ContinuousImageGenerationConfigurations API is enabled or
	// disabled.
	//
	// This member is required.
	Status ConfigurationStatus

	// The list of a key-value pair structure that contains extra parameters that can
	// be applied when the image is generated. The FormatConfig key is the JPEGQuality,
	// which indicates the JPEG quality key to be used to generate the image. The
	// FormatConfig value accepts ints from 1 to 100. If the value is 1, the image will
	// be generated with less quality and the best compression. If the value is 100,
	// the image will be generated with the best quality and less compression. If no
	// value is provided, the default value of the JPEGQuality key will be set to 80.
	FormatConfig map[string]string

	// The height of the output image that is used in conjunction with the WidthPixels
	// parameter. When both HeightPixels and WidthPixels parameters are provided, the
	// image will be stretched to fit the specified aspect ratio. If only the
	// HeightPixels parameter is provided, its original aspect ratio will be used to
	// calculate the WidthPixels ratio. If neither parameter is provided, the original
	// image size will be returned.
	HeightPixels *int32

	// The width of the output image that is used in conjunction with the HeightPixels
	// parameter. When both WidthPixels and HeightPixels parameters are provided, the
	// image will be stretched to fit the specified aspect ratio. If only the
	// WidthPixels parameter is provided, its original aspect ratio will be used to
	// calculate the HeightPixels ratio. If neither parameter is provided, the original
	// image size will be returned.
	WidthPixels *int32

	noSmithyDocumentSerde
}

// The structure that contains the information required to deliver images to a
// customer.
type ImageGenerationDestinationConfig struct {

	// The AWS Region of the S3 bucket where images will be delivered. This
	// DestinationRegion must match the Region where the stream is located.
	//
	// This member is required.
	DestinationRegion *string

	// The Uniform Resource Idenifier (URI) that identifies where the images will be
	// delivered.
	//
	// This member is required.
	Uri *string

	noSmithyDocumentSerde
}

// The structure that contains the notification information for the KVS images
// delivery. If this parameter is null, the configuration will be deleted from the
// stream.
type NotificationConfiguration struct {

	// The destination information required to deliver a notification to a customer.
	//
	// This member is required.
	DestinationConfig *NotificationDestinationConfig

	// Indicates if a notification configuration is enabled or disabled.
	//
	// This member is required.
	Status ConfigurationStatus

	noSmithyDocumentSerde
}

// The structure that contains the information required to deliver a notification
// to a customer.
type NotificationDestinationConfig struct {

	// The Uniform Resource Idenifier (URI) that identifies where the images will be
	// delivered.
	//
	// This member is required.
	Uri *string

	noSmithyDocumentSerde
}

// An object that describes the endpoint of the signaling channel returned by the
// GetSignalingChannelEndpoint API.
type ResourceEndpointListItem struct {

	// The protocol of the signaling channel returned by the
	// GetSignalingChannelEndpoint API.
	Protocol ChannelProtocol

	// The endpoint of the signaling channel returned by the
	// GetSignalingChannelEndpoint API.
	ResourceEndpoint *string

	noSmithyDocumentSerde
}

// An object that contains the endpoint configuration for the SINGLE_MASTER channel
// type.
type SingleMasterChannelEndpointConfiguration struct {

	// This property is used to determine the nature of communication over this
	// SINGLE_MASTER signaling channel. If WSS is specified, this API returns a
	// websocket endpoint. If HTTPS is specified, this API returns an HTTPS endpoint.
	Protocols []ChannelProtocol

	// This property is used to determine messaging permissions in this SINGLE_MASTER
	// signaling channel. If MASTER is specified, this API returns an endpoint that a
	// client can use to receive offers from and send answers to any of the viewers on
	// this signaling channel. If VIEWER is specified, this API returns an endpoint
	// that a client can use only to send offers to another MASTER client on this
	// signaling channel.
	Role ChannelRole

	noSmithyDocumentSerde
}

// A structure that contains the configuration for the SINGLE_MASTER channel type.
type SingleMasterConfiguration struct {

	// The period of time a signaling channel retains undelivered messages before they
	// are discarded.
	MessageTtlSeconds *int32

	noSmithyDocumentSerde
}

// An object describing a Kinesis video stream.
type StreamInfo struct {

	// A time stamp that indicates when the stream was created.
	CreationTime *time.Time

	// How long the stream retains data, in hours.
	DataRetentionInHours *int32

	// The name of the device that is associated with the stream.
	DeviceName *string

	// The ID of the Key Management Service (KMS) key that Kinesis Video Streams uses
	// to encrypt data on the stream.
	KmsKeyId *string

	// The MediaType of the stream.
	MediaType *string

	// The status of the stream.
	Status Status

	// The Amazon Resource Name (ARN) of the stream.
	StreamARN *string

	// The name of the stream.
	StreamName *string

	// The version of the stream.
	Version *string

	noSmithyDocumentSerde
}

// Specifies the condition that streams must satisfy to be returned when you list
// streams (see the ListStreams API). A condition has a comparison operation and a
// value. Currently, you can specify only the BEGINS_WITH operator, which finds
// streams whose names start with a given prefix.
type StreamNameCondition struct {

	// A comparison operator. Currently, you can specify only the BEGINS_WITH operator,
	// which finds streams whose names start with a given prefix.
	ComparisonOperator ComparisonOperator

	// A value to compare.
	ComparisonValue *string

	noSmithyDocumentSerde
}

// A key and value pair that is associated with the specified signaling channel.
type Tag struct {

	// The key of the tag that is associated with the specified signaling channel.
	//
	// This member is required.
	Key *string

	// The value of the tag that is associated with the specified signaling channel.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde