// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all ReplicationConfigurations, filtered by Source Server ID.
func (c *Client) GetReplicationConfiguration(ctx context.Context, params *GetReplicationConfigurationInput, optFns ...func(*Options)) (*GetReplicationConfigurationOutput, error) {
	if params == nil {
		params = &GetReplicationConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetReplicationConfiguration", params, optFns, c.addOperationGetReplicationConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetReplicationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetReplicationConfigurationInput struct {

	// Request to get Replication Configuration by Source Server ID.
	//
	// This member is required.
	SourceServerID *string

	noSmithyDocumentSerde
}

type GetReplicationConfigurationOutput struct {

	// Replication Configuration associate default Application Migration Service
	// Security Group.
	AssociateDefaultSecurityGroup *bool

	// Replication Configuration set bandwidth throttling.
	BandwidthThrottling int64

	// Replication Configuration create Public IP.
	CreatePublicIP *bool

	// Replication Configuration data plane routing.
	DataPlaneRouting types.ReplicationConfigurationDataPlaneRouting

	// Replication Configuration use default large Staging Disks.
	DefaultLargeStagingDiskType types.ReplicationConfigurationDefaultLargeStagingDiskType

	// Replication Configuration EBS encryption.
	EbsEncryption types.ReplicationConfigurationEbsEncryption

	// Replication Configuration EBS encryption key ARN.
	EbsEncryptionKeyArn *string

	// Replication Configuration name.
	Name *string

	// Replication Configuration replicated disks.
	ReplicatedDisks []types.ReplicationConfigurationReplicatedDisk

	// Replication Configuration Replication Server instance type.
	ReplicationServerInstanceType *string

	// Replication Configuration Replication Server Security Group IDs.
	ReplicationServersSecurityGroupsIDs []string

	// Replication Configuration Source Server ID.
	SourceServerID *string

	// Replication Configuration Staging Area subnet ID.
	StagingAreaSubnetId *string

	// Replication Configuration Staging Area tags.
	StagingAreaTags map[string]string

	// Replication Configuration use Dedicated Replication Server.
	UseDedicatedReplicationServer *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetReplicationConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetReplicationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetReplicationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetReplicationConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetReplicationConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetReplicationConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgn",
		OperationName: "GetReplicationConfiguration",
	}
}
