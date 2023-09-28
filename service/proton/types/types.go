// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// The Proton pipeline service role and repository data shared across the Amazon
// Web Services account.
type AccountSettings struct {

	// The repository configured in the Amazon Web Services account for pipeline
	// provisioning. Required it if you have environments configured for self-managed
	// provisioning with services that include pipelines.
	PipelineProvisioningRepository *RepositoryBranch

	// The Amazon Resource Name (ARN) of the service role you want to use for
	// provisioning pipelines. Assumed by Proton for Amazon Web Services-managed
	// provisioning, and by customer-owned automation for self-managed provisioning.
	PipelineServiceRoleArn *string

	noSmithyDocumentSerde
}

// Compatible environment template data.
type CompatibleEnvironmentTemplate struct {

	// The major version of the compatible environment template.
	//
	// This member is required.
	MajorVersion *string

	// The compatible environment template name.
	//
	// This member is required.
	TemplateName *string

	noSmithyDocumentSerde
}

// Compatible environment template data.
type CompatibleEnvironmentTemplateInput struct {

	// The major version of the compatible environment template.
	//
	// This member is required.
	MajorVersion *string

	// The compatible environment template name.
	//
	// This member is required.
	TemplateName *string

	noSmithyDocumentSerde
}

// Detailed data of an Proton component resource. For more information about
// components, see Proton components
// (https://docs.aws.amazon.com/proton/latest/adminguide/ag-components.html) in the
// Proton Administrator Guide.
type Component struct {

	// The Amazon Resource Name (ARN) of the component.
	//
	// This member is required.
	Arn *string

	// The time when the component was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The component deployment status.
	//
	// This member is required.
	DeploymentStatus DeploymentStatus

	// The name of the Proton environment that this component is associated with.
	//
	// This member is required.
	EnvironmentName *string

	// The time when the component was last modified.
	//
	// This member is required.
	LastModifiedAt *time.Time

	// The name of the component.
	//
	// This member is required.
	Name *string

	// The message associated with the component deployment status.
	DeploymentStatusMessage *string

	// A description of the component.
	Description *string

	// The time when a deployment of the component was last attempted.
	LastDeploymentAttemptedAt *time.Time

	// The time when the component was last deployed successfully.
	LastDeploymentSucceededAt *time.Time

	// The name of the service instance that this component is attached to. Provided
	// when a component is attached to a service instance.
	ServiceInstanceName *string

	// The name of the service that serviceInstanceName is associated with. Provided
	// when a component is attached to a service instance.
	ServiceName *string

	// The service spec that the component uses to access service inputs. Provided when
	// a component is attached to a service instance.
	//
	// This value conforms to the media type: application/yaml
	ServiceSpec *string

	noSmithyDocumentSerde
}

// Summary data of an Proton component resource. For more information about
// components, see Proton components
// (https://docs.aws.amazon.com/proton/latest/adminguide/ag-components.html) in the
// Proton Administrator Guide.
type ComponentSummary struct {

	// The Amazon Resource Name (ARN) of the component.
	//
	// This member is required.
	Arn *string

	// The time when the component was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The component deployment status.
	//
	// This member is required.
	DeploymentStatus DeploymentStatus

	// The name of the Proton environment that this component is associated with.
	//
	// This member is required.
	EnvironmentName *string

	// The time when the component was last modified.
	//
	// This member is required.
	LastModifiedAt *time.Time

	// The name of the component.
	//
	// This member is required.
	Name *string

	// The message associated with the component deployment status.
	DeploymentStatusMessage *string

	// The time when a deployment of the component was last attempted.
	LastDeploymentAttemptedAt *time.Time

	// The time when the component was last deployed successfully.
	LastDeploymentSucceededAt *time.Time

	// The name of the service instance that this component is attached to. Provided
	// when a component is attached to a service instance.
	ServiceInstanceName *string

	// The name of the service that serviceInstanceName is associated with. Provided
	// when a component is attached to a service instance.
	ServiceName *string

	noSmithyDocumentSerde
}

// Detailed data of an Proton environment resource. An Proton environment is a set
// of resources shared across Proton services.
type Environment struct {

	// The Amazon Resource Name (ARN) of the environment.
	//
	// This member is required.
	Arn *string

	// The time when the environment was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The environment deployment status.
	//
	// This member is required.
	DeploymentStatus DeploymentStatus

	// The time when a deployment of the environment was last attempted.
	//
	// This member is required.
	LastDeploymentAttemptedAt *time.Time

	// The time when the environment was last deployed successfully.
	//
	// This member is required.
	LastDeploymentSucceededAt *time.Time

	// The name of the environment.
	//
	// This member is required.
	Name *string

	// The major version of the environment template.
	//
	// This member is required.
	TemplateMajorVersion *string

	// The minor version of the environment template.
	//
	// This member is required.
	TemplateMinorVersion *string

	// The Amazon Resource Name (ARN) of the environment template.
	//
	// This member is required.
	TemplateName *string

	// The Amazon Resource Name (ARN) of the IAM service role that Proton uses when
	// provisioning directly defined components in this environment. It determines the
	// scope of infrastructure that a component can provision. The environment must
	// have a componentRoleArn to allow directly defined components to be associated
	// with the environment. For more information about components, see Proton
	// components
	// (https://docs.aws.amazon.com/proton/latest/adminguide/ag-components.html) in the
	// Proton Administrator Guide.
	ComponentRoleArn *string

	// An environment deployment status message.
	DeploymentStatusMessage *string

	// The description of the environment.
	Description *string

	// The ID of the environment account connection that's used to provision
	// infrastructure resources in an environment account.
	EnvironmentAccountConnectionId *string

	// The ID of the environment account that the environment infrastructure resources
	// are provisioned in.
	EnvironmentAccountId *string

	// The Amazon Resource Name (ARN) of the Proton service role that allows Proton to
	// make calls to other services on your behalf.
	ProtonServiceRoleArn *string

	// When included, indicates that the environment template is for customer
	// provisioned and managed infrastructure.
	Provisioning Provisioning

	// The infrastructure repository that you use to host your rendered infrastructure
	// templates for self-managed provisioning.
	ProvisioningRepository *RepositoryBranch

	// The environment spec.
	//
	// This value conforms to the media type: application/yaml
	Spec *string

	noSmithyDocumentSerde
}

// Detailed data of an Proton environment account connection resource.
type EnvironmentAccountConnection struct {

	// The Amazon Resource Name (ARN) of the environment account connection.
	//
	// This member is required.
	Arn *string

	// The environment account that's connected to the environment account connection.
	//
	// This member is required.
	EnvironmentAccountId *string

	// The name of the environment that's associated with the environment account
	// connection.
	//
	// This member is required.
	EnvironmentName *string

	// The ID of the environment account connection.
	//
	// This member is required.
	Id *string

	// The time when the environment account connection was last modified.
	//
	// This member is required.
	LastModifiedAt *time.Time

	// The ID of the management account that's connected to the environment account
	// connection.
	//
	// This member is required.
	ManagementAccountId *string

	// The time when the environment account connection request was made.
	//
	// This member is required.
	RequestedAt *time.Time

	// The IAM service role that's associated with the environment account connection.
	//
	// This member is required.
	RoleArn *string

	// The status of the environment account connection.
	//
	// This member is required.
	Status EnvironmentAccountConnectionStatus

	// The Amazon Resource Name (ARN) of the IAM service role that Proton uses when
	// provisioning directly defined components in the associated environment account.
	// It determines the scope of infrastructure that a component can provision in the
	// account. The environment account connection must have a componentRoleArn to
	// allow directly defined components to be associated with any environments running
	// in the account. For more information about components, see Proton components
	// (https://docs.aws.amazon.com/proton/latest/adminguide/ag-components.html) in the
	// Proton Administrator Guide.
	ComponentRoleArn *string

	noSmithyDocumentSerde
}

// Summary data of an Proton environment account connection resource.
type EnvironmentAccountConnectionSummary struct {

	// The Amazon Resource Name (ARN) of the environment account connection.
	//
	// This member is required.
	Arn *string

	// The ID of the environment account that's connected to the environment account
	// connection.
	//
	// This member is required.
	EnvironmentAccountId *string

	// The name of the environment that's associated with the environment account
	// connection.
	//
	// This member is required.
	EnvironmentName *string

	// The ID of the environment account connection.
	//
	// This member is required.
	Id *string

	// The time when the environment account connection was last modified.
	//
	// This member is required.
	LastModifiedAt *time.Time

	// The ID of the management account that's connected to the environment account
	// connection.
	//
	// This member is required.
	ManagementAccountId *string

	// The time when the environment account connection request was made.
	//
	// This member is required.
	RequestedAt *time.Time

	// The IAM service role that's associated with the environment account connection.
	//
	// This member is required.
	RoleArn *string

	// The status of the environment account connection.
	//
	// This member is required.
	Status EnvironmentAccountConnectionStatus

	// The Amazon Resource Name (ARN) of the IAM service role that Proton uses when
	// provisioning directly defined components in the associated environment account.
	// It determines the scope of infrastructure that a component can provision in the
	// account. The environment account connection must have a componentRoleArn to
	// allow directly defined components to be associated with any environments running
	// in the account. For more information about components, see Proton components
	// (https://docs.aws.amazon.com/proton/latest/adminguide/ag-components.html) in the
	// Proton Administrator Guide.
	ComponentRoleArn *string

	noSmithyDocumentSerde
}

// Summary data of an Proton environment resource. An Proton environment is a set
// of resources shared across Proton services.
type EnvironmentSummary struct {

	// The Amazon Resource Name (ARN) of the environment.
	//
	// This member is required.
	Arn *string

	// The time when the environment was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The environment deployment status.
	//
	// This member is required.
	DeploymentStatus DeploymentStatus

	// The time when a deployment of the environment was last attempted.
	//
	// This member is required.
	LastDeploymentAttemptedAt *time.Time

	// The time when the environment was last deployed successfully.
	//
	// This member is required.
	LastDeploymentSucceededAt *time.Time

	// The name of the environment.
	//
	// This member is required.
	Name *string

	// The major version of the environment template.
	//
	// This member is required.
	TemplateMajorVersion *string

	// The minor version of the environment template.
	//
	// This member is required.
	TemplateMinorVersion *string

	// The name of the environment template.
	//
	// This member is required.
	TemplateName *string

	// The Amazon Resource Name (ARN) of the IAM service role that Proton uses when
	// provisioning directly defined components in this environment. It determines the
	// scope of infrastructure that a component can provision. The environment must
	// have a componentRoleArn to allow directly defined components to be associated
	// with the environment. For more information about components, see Proton
	// components
	// (https://docs.aws.amazon.com/proton/latest/adminguide/ag-components.html) in the
	// Proton Administrator Guide.
	ComponentRoleArn *string

	// An environment deployment status message.
	DeploymentStatusMessage *string

	// The description of the environment.
	Description *string

	// The ID of the environment account connection that the environment is associated
	// with.
	EnvironmentAccountConnectionId *string

	// The ID of the environment account that the environment infrastructure resources
	// are provisioned in.
	EnvironmentAccountId *string

	// The Amazon Resource Name (ARN) of the Proton service role that allows Proton to
	// make calls to other services on your behalf.
	ProtonServiceRoleArn *string

	// When included, indicates that the environment template is for customer
	// provisioned and managed infrastructure.
	Provisioning Provisioning

	noSmithyDocumentSerde
}

// The environment template data.
type EnvironmentTemplate struct {

	// The Amazon Resource Name (ARN) of the environment template.
	//
	// This member is required.
	Arn *string

	// The time when the environment template was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The time when the environment template was last modified.
	//
	// This member is required.
	LastModifiedAt *time.Time

	// The name of the environment template.
	//
	// This member is required.
	Name *string

	// A description of the environment template.
	Description *string

	// The name of the environment template as displayed in the developer interface.
	DisplayName *string

	// The customer provided encryption key for the environment template.
	EncryptionKey *string

	// When included, indicates that the environment template is for customer
	// provisioned and managed infrastructure.
	Provisioning Provisioning

	// The ID of the recommended version of the environment template.
	RecommendedVersion *string

	noSmithyDocumentSerde
}

// A search filter for environment templates.
type EnvironmentTemplateFilter struct {

	// Include majorVersion to filter search for a major version.
	//
	// This member is required.
	MajorVersion *string

	// Include templateName to filter search for a template name.
	//
	// This member is required.
	TemplateName *string

	noSmithyDocumentSerde
}

// The environment template data.
type EnvironmentTemplateSummary struct {

	// The Amazon Resource Name (ARN) of the environment template.
	//
	// This member is required.
	Arn *string

	// The time when the environment template was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The time when the environment template was last modified.
	//
	// This member is required.
	LastModifiedAt *time.Time

	// The name of the environment template.
	//
	// This member is required.
	Name *string

	// A description of the environment template.
	Description *string

	// The name of the environment template as displayed in the developer interface.
	DisplayName *string

	// When included, indicates that the environment template is for customer
	// provisioned and managed infrastructure.
	Provisioning Provisioning

	// The recommended version of the environment template.
	RecommendedVersion *string

	noSmithyDocumentSerde
}

// The environment template version data.
type EnvironmentTemplateVersion struct {

	// The Amazon Resource Name (ARN) of the version of an environment template.
	//
	// This member is required.
	Arn *string

	// The time when the version of an environment template was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The time when the version of an environment template was last modified.
	//
	// This member is required.
	LastModifiedAt *time.Time

	// The latest major version that's associated with the version of an environment
	// template.
	//
	// This member is required.
	MajorVersion *string

	// The minor version of an environment template.
	//
	// This member is required.
	MinorVersion *string

	// The status of the version of an environment template.
	//
	// This member is required.
	Status TemplateVersionStatus

	// The name of the version of an environment template.
	//
	// This member is required.
	TemplateName *string

	// A description of the minor version of an environment template.
	Description *string

	// The recommended minor version of the environment template.
	RecommendedMinorVersion *string

	// The schema of the version of an environment template.
	//
	// This value conforms to the media type: application/yaml
	Schema *string

	// The status message of the version of an environment template.
	StatusMessage *string

	noSmithyDocumentSerde
}

// A summary of the version of an environment template detail data.
type EnvironmentTemplateVersionSummary struct {

	// The Amazon Resource Name (ARN) of the version of an environment template.
	//
	// This member is required.
	Arn *string

	// The time when the version of an environment template was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The time when the version of an environment template was last modified.
	//
	// This member is required.
	LastModifiedAt *time.Time

	// The latest major version that's associated with the version of an environment
	// template.
	//
	// This member is required.
	MajorVersion *string

	// The version of an environment template.
	//
	// This member is required.
	MinorVersion *string

	// The status of the version of an environment template.
	//
	// This member is required.
	Status TemplateVersionStatus

	// The name of the environment template.
	//
	// This member is required.
	TemplateName *string

	// A description of the version of an environment template.
	Description *string

	// The recommended minor version of the environment template.
	RecommendedMinorVersion *string

	// The status message of the version of an environment template.
	StatusMessage *string

	noSmithyDocumentSerde
}

// An infrastructure as code defined resource output.
type Output struct {

	// The output key.
	Key *string

	// The output value.
	ValueString *string

	noSmithyDocumentSerde
}

// Detail data for a provisioned resource.
type ProvisionedResource struct {

	// The provisioned resource identifier.
	Identifier *string

	// The provisioned resource name.
	Name *string

	// The resource provisioning engine. At this time, CLOUDFORMATION can be used for
	// Amazon Web Services-managed provisioning, and TERRAFORM can be used for
	// self-managed provisioning. For more information, see Self-managed provisioning
	// (https://docs.aws.amazon.com/proton/latest/adminguide/ag-works-prov-methods.html#ag-works-prov-methods-self)
	// in the Proton Administrator Guide.
	ProvisioningEngine ProvisionedResourceEngine

	noSmithyDocumentSerde
}

// Detailed data of a repository that has been registered with Proton.
type Repository struct {

	// The repository Amazon Resource Name (ARN).
	//
	// This member is required.
	Arn *string

	// The repository Amazon Web Services CodeStar connection that connects Proton to
	// your repository.
	//
	// This member is required.
	ConnectionArn *string

	// The repository name.
	//
	// This member is required.
	Name *string

	// The repository provider.
	//
	// This member is required.
	Provider RepositoryProvider

	// Your customer Amazon Web Services KMS encryption key.
	EncryptionKey *string

	noSmithyDocumentSerde
}

// Detail data for a repository branch.
type RepositoryBranch struct {

	// The Amazon Resource Name (ARN) of the repository branch.
	//
	// This member is required.
	Arn *string

	// The repository branch.
	//
	// This member is required.
	Branch *string

	// The repository name.
	//
	// This member is required.
	Name *string

	// The repository provider.
	//
	// This member is required.
	Provider RepositoryProvider

	noSmithyDocumentSerde
}

// Detail input data for a repository branch.
type RepositoryBranchInput struct {

	// The repository branch.
	//
	// This member is required.
	Branch *string

	// The repository name.
	//
	// This member is required.
	Name *string

	// The repository provider.
	//
	// This member is required.
	Provider RepositoryProvider

	noSmithyDocumentSerde
}

// Summary data of a repository that has been registered with Proton.
type RepositorySummary struct {

	// The Amazon Resource Name (ARN) for a repository.
	//
	// This member is required.
	Arn *string

	// The repository name.
	//
	// This member is required.
	Name *string

	// The repository provider.
	//
	// This member is required.
	Provider RepositoryProvider

	noSmithyDocumentSerde
}

// Detail data for a repository sync attempt activated by a push to a repository.
type RepositorySyncAttempt struct {

	// Detail data for sync attempt events.
	//
	// This member is required.
	Events []RepositorySyncEvent

	// The time when the sync attempt started.
	//
	// This member is required.
	StartedAt *time.Time

	// The sync attempt status.
	//
	// This member is required.
	Status RepositorySyncStatus

	noSmithyDocumentSerde
}

// The repository sync definition.
type RepositorySyncDefinition struct {

	// The repository branch.
	//
	// This member is required.
	Branch *string

	// The directory in the repository.
	//
	// This member is required.
	Directory *string

	// The resource that is synced from.
	//
	// This member is required.
	Parent *string

	// The resource that is synced to.
	//
	// This member is required.
	Target *string

	noSmithyDocumentSerde
}

// Repository sync event detail data for a sync attempt.
type RepositorySyncEvent struct {

	// Event detail for a repository sync attempt.
	//
	// This member is required.
	Event *string

	// The time that the sync event occurred.
	//
	// This member is required.
	Time *time.Time

	// The type of event.
	//
	// This member is required.
	Type *string

	// The external ID of the sync event.
	ExternalId *string

	noSmithyDocumentSerde
}

// Detail data for a resource sync attempt activated by a push to a repository.
type ResourceSyncAttempt struct {

	// An array of events with detail data.
	//
	// This member is required.
	Events []ResourceSyncEvent

	// Detail data for the initial repository commit, path and push.
	//
	// This member is required.
	InitialRevision *Revision

	// The time when the sync attempt started.
	//
	// This member is required.
	StartedAt *time.Time

	// The status of the sync attempt.
	//
	// This member is required.
	Status ResourceSyncStatus

	// The resource that is synced to.
	//
	// This member is required.
	Target *string

	// Detail data for the target revision.
	//
	// This member is required.
	TargetRevision *Revision

	noSmithyDocumentSerde
}

// Detail data for a resource sync event.
type ResourceSyncEvent struct {

	// A resource sync event.
	//
	// This member is required.
	Event *string

	// The time when the event occurred.
	//
	// This member is required.
	Time *time.Time

	// The type of event.
	//
	// This member is required.
	Type *string

	// The external ID for the event.
	ExternalId *string

	noSmithyDocumentSerde
}

// Revision detail data for a commit and push that activates a sync attempt
type Revision struct {

	// The repository branch.
	//
	// This member is required.
	Branch *string

	// The repository directory changed by a commit and push that activated the sync
	// attempt.
	//
	// This member is required.
	Directory *string

	// The repository name.
	//
	// This member is required.
	RepositoryName *string

	// The repository provider.
	//
	// This member is required.
	RepositoryProvider RepositoryProvider

	// The secure hash algorithm (SHA) hash for the revision.
	//
	// This member is required.
	Sha *string

	noSmithyDocumentSerde
}

// Template bundle S3 bucket data.
type S3ObjectSource struct {

	// The name of the S3 bucket that contains a template bundle.
	//
	// This member is required.
	Bucket *string

	// The path to the S3 bucket that contains a template bundle.
	//
	// This member is required.
	Key *string

	noSmithyDocumentSerde
}

// Detailed data of an Proton service resource.
type Service struct {

	// The Amazon Resource Name (ARN) of the service.
	//
	// This member is required.
	Arn *string

	// The time when the service was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The time when the service was last modified.
	//
	// This member is required.
	LastModifiedAt *time.Time

	// The name of the service.
	//
	// This member is required.
	Name *string

	// The formatted specification that defines the service.
	//
	// This value conforms to the media type: application/yaml
	//
	// This member is required.
	Spec *string

	// The status of the service.
	//
	// This member is required.
	Status ServiceStatus

	// The name of the service template.
	//
	// This member is required.
	TemplateName *string

	// The name of the code repository branch that holds the code that's deployed in
	// Proton.
	BranchName *string

	// A description of the service.
	Description *string

	// The service pipeline detail data.
	Pipeline *ServicePipeline

	// The Amazon Resource Name (ARN) of the repository connection. For more
	// information, see Set up a repository connection
	// (https://docs.aws.amazon.com/proton/latest/adminguide/setting-up-for-service.html#setting-up-vcontrol)
	// in the Proton Administrator Guide and Setting up with Proton
	// (https://docs.aws.amazon.com/proton/latest/userguide/proton-setup.html#setup-repo-connection)
	// in the Proton User Guide.
	RepositoryConnectionArn *string

	// The ID of the source code repository.
	RepositoryId *string

	// A service status message.
	StatusMessage *string

	noSmithyDocumentSerde
}

// Detailed data of an Proton service instance resource.
type ServiceInstance struct {

	// The Amazon Resource Name (ARN) of the service instance.
	//
	// This member is required.
	Arn *string

	// The time when the service instance was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The service instance deployment status.
	//
	// This member is required.
	DeploymentStatus DeploymentStatus

	// The name of the environment that the service instance was deployed into.
	//
	// This member is required.
	EnvironmentName *string

	// The time when a deployment of the service instance was last attempted.
	//
	// This member is required.
	LastDeploymentAttemptedAt *time.Time

	// The time when the service instance was last deployed successfully.
	//
	// This member is required.
	LastDeploymentSucceededAt *time.Time

	// The name of the service instance.
	//
	// This member is required.
	Name *string

	// The name of the service that the service instance belongs to.
	//
	// This member is required.
	ServiceName *string

	// The major version of the service template that was used to create the service
	// instance.
	//
	// This member is required.
	TemplateMajorVersion *string

	// The minor version of the service template that was used to create the service
	// instance.
	//
	// This member is required.
	TemplateMinorVersion *string

	// The name of the service template that was used to create the service instance.
	//
	// This member is required.
	TemplateName *string

	// The message associated with the service instance deployment status.
	DeploymentStatusMessage *string

	// The service spec that was used to create the service instance.
	//
	// This value conforms to the media type: application/yaml
	Spec *string

	noSmithyDocumentSerde
}

// Summary data of an Proton service instance resource.
type ServiceInstanceSummary struct {

	// The Amazon Resource Name (ARN) of the service instance.
	//
	// This member is required.
	Arn *string

	// The time when the service instance was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The service instance deployment status.
	//
	// This member is required.
	DeploymentStatus DeploymentStatus

	// The name of the environment that the service instance was deployed into.
	//
	// This member is required.
	EnvironmentName *string

	// The time when a deployment of the service was last attempted.
	//
	// This member is required.
	LastDeploymentAttemptedAt *time.Time

	// The time when the service was last deployed successfully.
	//
	// This member is required.
	LastDeploymentSucceededAt *time.Time

	// The name of the service instance.
	//
	// This member is required.
	Name *string

	// The name of the service that the service instance belongs to.
	//
	// This member is required.
	ServiceName *string

	// The service instance template major version.
	//
	// This member is required.
	TemplateMajorVersion *string

	// The service instance template minor version.
	//
	// This member is required.
	TemplateMinorVersion *string

	// The name of the service template.
	//
	// This member is required.
	TemplateName *string

	// A service instance deployment status message.
	DeploymentStatusMessage *string

	noSmithyDocumentSerde
}

// Detailed data of an Proton service instance pipeline resource.
type ServicePipeline struct {

	// The Amazon Resource Name (ARN) of the service pipeline.
	//
	// This member is required.
	Arn *string

	// The time when the service pipeline was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The deployment status of the service pipeline.
	//
	// This member is required.
	DeploymentStatus DeploymentStatus

	// The time when a deployment of the service pipeline was last attempted.
	//
	// This member is required.
	LastDeploymentAttemptedAt *time.Time

	// The time when the service pipeline was last deployed successfully.
	//
	// This member is required.
	LastDeploymentSucceededAt *time.Time

	// The major version of the service template that was used to create the service
	// pipeline.
	//
	// This member is required.
	TemplateMajorVersion *string

	// The minor version of the service template that was used to create the service
	// pipeline.
	//
	// This member is required.
	TemplateMinorVersion *string

	// The name of the service template that was used to create the service pipeline.
	//
	// This member is required.
	TemplateName *string

	// A service pipeline deployment status message.
	DeploymentStatusMessage *string

	// The service spec that was used to create the service pipeline.
	//
	// This value conforms to the media type: application/yaml
	Spec *string

	noSmithyDocumentSerde
}

// Summary data of an Proton service resource.
type ServiceSummary struct {

	// The Amazon Resource Name (ARN) of the service.
	//
	// This member is required.
	Arn *string

	// The time when the service was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The time when the service was last modified.
	//
	// This member is required.
	LastModifiedAt *time.Time

	// The name of the service.
	//
	// This member is required.
	Name *string

	// The status of the service.
	//
	// This member is required.
	Status ServiceStatus

	// The name of the service template.
	//
	// This member is required.
	TemplateName *string

	// A description of the service.
	Description *string

	// A service status message.
	StatusMessage *string

	noSmithyDocumentSerde
}

// Detailed data of an Proton service template resource.
type ServiceTemplate struct {

	// The Amazon Resource Name (ARN) of the service template.
	//
	// This member is required.
	Arn *string

	// The time when the service template was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The time when the service template was last modified.
	//
	// This member is required.
	LastModifiedAt *time.Time

	// The name of the service template.
	//
	// This member is required.
	Name *string

	// A description of the service template.
	Description *string

	// The service template name as displayed in the developer interface.
	DisplayName *string

	// The customer provided service template encryption key that's used to encrypt
	// data.
	EncryptionKey *string

	// If pipelineProvisioning is true, a service pipeline is included in the service
	// template. Otherwise, a service pipeline isn't included in the service template.
	PipelineProvisioning Provisioning

	// The recommended version of the service template.
	RecommendedVersion *string

	noSmithyDocumentSerde
}

// Summary data of an Proton service template resource.
type ServiceTemplateSummary struct {

	// The Amazon Resource Name (ARN) of the service template.
	//
	// This member is required.
	Arn *string

	// The time when the service template was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The time when the service template was last modified.
	//
	// This member is required.
	LastModifiedAt *time.Time

	// The name of the service template.
	//
	// This member is required.
	Name *string

	// A description of the service template.
	Description *string

	// The service template name as displayed in the developer interface.
	DisplayName *string

	// If pipelineProvisioning is true, a service pipeline is included in the service
	// template, otherwise a service pipeline isn't included in the service template.
	PipelineProvisioning Provisioning

	// The recommended version of the service template.
	RecommendedVersion *string

	noSmithyDocumentSerde
}

// Detailed data of an Proton service template version resource.
type ServiceTemplateVersion struct {

	// The Amazon Resource Name (ARN) of the version of a service template.
	//
	// This member is required.
	Arn *string

	// An array of compatible environment template names for the major version of a
	// service template.
	//
	// This member is required.
	CompatibleEnvironmentTemplates []CompatibleEnvironmentTemplate

	// The time when the version of a service template was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The time when the version of a service template was last modified.
	//
	// This member is required.
	LastModifiedAt *time.Time

	// The latest major version that's associated with the version of a service
	// template.
	//
	// This member is required.
	MajorVersion *string

	// The minor version of a service template.
	//
	// This member is required.
	MinorVersion *string

	// The service template version status.
	//
	// This member is required.
	Status TemplateVersionStatus

	// The name of the version of a service template.
	//
	// This member is required.
	TemplateName *string

	// A description of the version of a service template.
	Description *string

	// The recommended minor version of the service template.
	RecommendedMinorVersion *string

	// The schema of the version of a service template.
	//
	// This value conforms to the media type: application/yaml
	Schema *string

	// A service template version status message.
	StatusMessage *string

	// An array of supported component sources. Components with supported sources can
	// be attached to service instances based on this service template version. For
	// more information about components, see Proton components
	// (https://docs.aws.amazon.com/proton/latest/adminguide/ag-components.html) in the
	// Proton Administrator Guide.
	SupportedComponentSources []ServiceTemplateSupportedComponentSourceType

	noSmithyDocumentSerde
}

// Summary data of an Proton service template version resource.
type ServiceTemplateVersionSummary struct {

	// The Amazon Resource Name (ARN) of the version of a service template.
	//
	// This member is required.
	Arn *string

	// The time when the version of a service template was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The time when the version of a service template was last modified.
	//
	// This member is required.
	LastModifiedAt *time.Time

	// The latest major version that's associated with the version of a service
	// template.
	//
	// This member is required.
	MajorVersion *string

	// The minor version of a service template.
	//
	// This member is required.
	MinorVersion *string

	// The service template minor version status.
	//
	// This member is required.
	Status TemplateVersionStatus

	// The name of the service template.
	//
	// This member is required.
	TemplateName *string

	// A description of the version of a service template.
	Description *string

	// The recommended minor version of the service template.
	RecommendedMinorVersion *string

	// A service template minor version status message.
	StatusMessage *string

	noSmithyDocumentSerde
}

// A description of a resource tag.
type Tag struct {

	// The key of the resource tag.
	//
	// This member is required.
	Key *string

	// The value of the resource tag.
	//
	// This member is required.
	Value *string

	noSmithyDocumentSerde
}

// The detail data for a template sync configuration.
type TemplateSyncConfig struct {

	// The repository branch.
	//
	// This member is required.
	Branch *string

	// The name of the repository, for example myrepos/myrepo.
	//
	// This member is required.
	RepositoryName *string

	// The repository provider.
	//
	// This member is required.
	RepositoryProvider RepositoryProvider

	// The template name.
	//
	// This member is required.
	TemplateName *string

	// The template type.
	//
	// This member is required.
	TemplateType TemplateType

	// A subdirectory path to your template bundle version.
	Subdirectory *string

	noSmithyDocumentSerde
}

// Template version source data.
//
// The following types satisfy this interface:
//
//	TemplateVersionSourceInputMemberS3
type TemplateVersionSourceInput interface {
	isTemplateVersionSourceInput()
}

// An S3 source object that includes the template bundle S3 path and name for a
// template minor version.
type TemplateVersionSourceInputMemberS3 struct {
	Value S3ObjectSource

	noSmithyDocumentSerde
}

func (*TemplateVersionSourceInputMemberS3) isTemplateVersionSourceInput() {}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isTemplateVersionSourceInput() {}
