// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Defines a file system on a Network File System (NFS) server that can be read
// from or written to.
func (c *Client) CreateLocationNfs(ctx context.Context, params *CreateLocationNfsInput, optFns ...func(*Options)) (*CreateLocationNfsOutput, error) {
	if params == nil {
		params = &CreateLocationNfsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLocationNfs", params, optFns, c.addOperationCreateLocationNfsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLocationNfsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CreateLocationNfsRequest
type CreateLocationNfsInput struct {

	// Contains a list of Amazon Resource Names (ARNs) of agents that are used to
	// connect to an NFS server. If you are copying data to or from your Snowcone
	// device, see NFS Server on Snowcone (https://docs.aws.amazon.com/datasync/latest/userguide/create-nfs-location.html#nfs-on-snowcone)
	// for more information.
	//
	// This member is required.
	OnPremConfig *types.OnPremConfig

	// The name of the NFS server. This value is the IP address or Domain Name Service
	// (DNS) name of the NFS server. An agent that is installed on-premises uses this
	// hostname to mount the NFS server in a network. If you are copying data to or
	// from your Snowcone device, see NFS Server on Snowcone (https://docs.aws.amazon.com/datasync/latest/userguide/create-nfs-location.html#nfs-on-snowcone)
	// for more information. This name must either be DNS-compliant or must be an IP
	// version 4 (IPv4) address.
	//
	// This member is required.
	ServerHostname *string

	// The subdirectory in the NFS file system that is used to read data from the NFS
	// source location or write data to the NFS destination. The NFS path should be a
	// path that's exported by the NFS server, or a subdirectory of that path. The path
	// should be such that it can be mounted by other NFS clients in your network. To
	// see all the paths exported by your NFS server, run " showmount -e nfs-server-name
	// " from an NFS client that has access to your server. You can specify any
	// directory that appears in the results, and any subdirectory of that directory.
	// Ensure that the NFS export is accessible without Kerberos authentication. To
	// transfer all the data in the folder you specified, DataSync needs to have
	// permissions to read all the data. To ensure this, either configure the NFS
	// export with no_root_squash, or ensure that the permissions for all of the files
	// that you want DataSync allow read access for all users. Doing either enables the
	// agent to read the files. For the agent to access directories, you must
	// additionally enable all execute access. If you are copying data to or from your
	// Snowcone device, see NFS Server on Snowcone (https://docs.aws.amazon.com/datasync/latest/userguide/create-nfs-location.html#nfs-on-snowcone)
	// for more information. For information about NFS export configuration, see 18.7.
	// The /etc/exports Configuration File in the Red Hat Enterprise Linux
	// documentation.
	//
	// This member is required.
	Subdirectory *string

	// The NFS mount options that DataSync can use to mount your NFS share.
	MountOptions *types.NfsMountOptions

	// The key-value pair that represents the tag that you want to add to the
	// location. The value can be an empty string. We recommend using tags to name your
	// resources.
	Tags []types.TagListEntry

	noSmithyDocumentSerde
}

// CreateLocationNfsResponse
type CreateLocationNfsOutput struct {

	// The Amazon Resource Name (ARN) of the source NFS file system location that is
	// created.
	LocationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLocationNfsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLocationNfs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLocationNfs{}, middleware.After)
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
	if err = addOpCreateLocationNfsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLocationNfs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLocationNfs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "CreateLocationNfs",
	}
}
