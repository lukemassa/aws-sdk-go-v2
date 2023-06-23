// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the resource policy that is set on a repository. After a resource
// policy is deleted, the permissions allowed and denied by the deleted policy are
// removed. The effect of deleting a resource policy might not be immediate. Use
// DeleteRepositoryPermissionsPolicy with caution. After a policy is deleted,
// Amazon Web Services users, roles, and accounts lose permissions to perform the
// repository actions granted by the deleted policy.
func (c *Client) DeleteRepositoryPermissionsPolicy(ctx context.Context, params *DeleteRepositoryPermissionsPolicyInput, optFns ...func(*Options)) (*DeleteRepositoryPermissionsPolicyOutput, error) {
	if params == nil {
		params = &DeleteRepositoryPermissionsPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteRepositoryPermissionsPolicy", params, optFns, c.addOperationDeleteRepositoryPermissionsPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteRepositoryPermissionsPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRepositoryPermissionsPolicyInput struct {

	// The name of the domain that contains the repository associated with the
	// resource policy to be deleted.
	//
	// This member is required.
	Domain *string

	// The name of the repository that is associated with the resource policy to be
	// deleted
	//
	// This member is required.
	Repository *string

	// The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	DomainOwner *string

	// The revision of the repository's resource policy to be deleted. This revision
	// is used for optimistic locking, which prevents others from accidentally
	// overwriting your changes to the repository's resource policy.
	PolicyRevision *string

	noSmithyDocumentSerde
}

type DeleteRepositoryPermissionsPolicyOutput struct {

	// Information about the deleted policy after processing the request.
	Policy *types.ResourcePolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteRepositoryPermissionsPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteRepositoryPermissionsPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteRepositoryPermissionsPolicy{}, middleware.After)
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
	if err = addOpDeleteRepositoryPermissionsPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRepositoryPermissionsPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteRepositoryPermissionsPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "DeleteRepositoryPermissionsPolicy",
	}
}
