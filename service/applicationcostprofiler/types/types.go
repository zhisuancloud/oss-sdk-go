// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// The configuration of a report in AWS Application Cost Profiler.
type ReportDefinition struct {

	// Timestamp (milliseconds) when this report definition was created.
	CreatedAt *time.Time

	// The location in Amazon Simple Storage Service (Amazon S3) the reports should be
	// saved to.
	DestinationS3Location *S3Location

	// The format used for the generated reports.
	Format Format

	// Timestamp (milliseconds) when this report definition was last updated.
	LastUpdatedAt *time.Time

	// Description of the report
	ReportDescription *string

	// The cadence at which the report is generated.
	ReportFrequency ReportFrequency

	// The ID of the report.
	ReportId *string

	noSmithyDocumentSerde
}

// Represents the Amazon Simple Storage Service (Amazon S3) location where AWS
// Application Cost Profiler reports are generated and then written to.
type S3Location struct {

	// Name of the S3 bucket.
	//
	// This member is required.
	Bucket *string

	// Prefix for the location to write to.
	//
	// This member is required.
	Prefix *string

	noSmithyDocumentSerde
}

// Represents the Amazon Simple Storage Service (Amazon S3) location where usage
// data is read from.
type SourceS3Location struct {

	// Name of the bucket.
	//
	// This member is required.
	Bucket *string

	// Key of the object.
	//
	// This member is required.
	Key *string

	// Region of the bucket. Only required for Regions that are disabled by default.
	// For more infomration about Regions that are disabled by default, see  Enabling a
	// Region
	// (https://docs.aws.amazon.com/general/latest/gr/rande-manage.html#rande-manage-enable)
	// in the AWS General Reference guide.
	Region S3BucketRegion

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde