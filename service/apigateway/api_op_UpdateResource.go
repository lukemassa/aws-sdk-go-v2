// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Changes information about a Resource resource.
func (c *Client) UpdateResource(ctx context.Context, params *UpdateResourceInput, optFns ...func(*Options)) (*UpdateResourceOutput, error) {
	if params == nil {
		params = &UpdateResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateResource", params, optFns, c.addOperationUpdateResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to change information about a Resource resource.
type UpdateResourceInput struct {

	// The identifier of the Resource resource.
	//
	// This member is required.
	ResourceId *string

	// The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// For more information about supported patch operations, see Patch Operations (https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html)
	// .
	PatchOperations []types.PatchOperation

	noSmithyDocumentSerde
}

// Represents an API resource.
type UpdateResourceOutput struct {

	// The resource's identifier.
	Id *string

	// The parent resource's identifier.
	ParentId *string

	// The full path for this resource.
	Path *string

	// The last path segment for this resource.
	PathPart *string

	// Gets an API resource's method of a given HTTP verb.
	ResourceMethods map[string]types.Method

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateResource{}, middleware.After)
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
	if err = addOpUpdateResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateResource(options.Region), middleware.Before); err != nil {
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
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opUpdateResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "UpdateResource",
	}
}
