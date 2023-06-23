// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrassv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrassv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a list of components that meet the component, version, and platform
// requirements of a deployment. Greengrass core devices call this operation when
// they receive a deployment to identify the components to install. This operation
// identifies components that meet all dependency requirements for a deployment. If
// the requirements conflict, then this operation returns an error and the
// deployment fails. For example, this occurs if component A requires version
// >2.0.0 and component B requires version <2.0.0 of a component dependency. When
// you specify the component candidates to resolve, IoT Greengrass compares each
// component's digest from the core device with the component's digest in the
// Amazon Web Services Cloud. If the digests don't match, then IoT Greengrass
// specifies to use the version from the Amazon Web Services Cloud. To use this
// operation, you must use the data plane API endpoint and authenticate with an IoT
// device certificate. For more information, see IoT Greengrass endpoints and
// quotas (https://docs.aws.amazon.com/general/latest/gr/greengrass.html) .
func (c *Client) ResolveComponentCandidates(ctx context.Context, params *ResolveComponentCandidatesInput, optFns ...func(*Options)) (*ResolveComponentCandidatesOutput, error) {
	if params == nil {
		params = &ResolveComponentCandidatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResolveComponentCandidates", params, optFns, c.addOperationResolveComponentCandidatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResolveComponentCandidatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResolveComponentCandidatesInput struct {

	// The list of components to resolve.
	ComponentCandidates []types.ComponentCandidate

	// The platform to use to resolve compatible components.
	Platform *types.ComponentPlatform

	noSmithyDocumentSerde
}

type ResolveComponentCandidatesOutput struct {

	// A list of components that meet the requirements that you specify in the
	// request. This list includes each component's recipe that you can use to install
	// the component.
	ResolvedComponentVersions []types.ResolvedComponentVersion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationResolveComponentCandidatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpResolveComponentCandidates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpResolveComponentCandidates{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opResolveComponentCandidates(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opResolveComponentCandidates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "ResolveComponentCandidates",
	}
}
