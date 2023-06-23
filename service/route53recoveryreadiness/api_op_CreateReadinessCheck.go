// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoveryreadiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a readiness check in an account. A readiness check monitors a resource
// set in your application, such as a set of Amazon Aurora instances, that
// Application Recovery Controller is auditing recovery readiness for. The audits
// run once every minute on every resource that's associated with a readiness
// check.
func (c *Client) CreateReadinessCheck(ctx context.Context, params *CreateReadinessCheckInput, optFns ...func(*Options)) (*CreateReadinessCheckOutput, error) {
	if params == nil {
		params = &CreateReadinessCheckInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateReadinessCheck", params, optFns, c.addOperationCreateReadinessCheckMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateReadinessCheckOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateReadinessCheckInput struct {

	// The name of the readiness check to create.
	//
	// This member is required.
	ReadinessCheckName *string

	// The name of the resource set to check.
	//
	// This member is required.
	ResourceSetName *string

	// A collection of tags associated with a resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateReadinessCheckOutput struct {

	// The Amazon Resource Name (ARN) associated with a readiness check.
	ReadinessCheckArn *string

	// Name of a readiness check.
	ReadinessCheckName *string

	// Name of the resource set to be checked.
	ResourceSet *string

	// A collection of tags associated with a resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateReadinessCheckMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateReadinessCheck{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateReadinessCheck{}, middleware.After)
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
	if err = addOpCreateReadinessCheckValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReadinessCheck(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateReadinessCheck(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53-recovery-readiness",
		OperationName: "CreateReadinessCheck",
	}
}
