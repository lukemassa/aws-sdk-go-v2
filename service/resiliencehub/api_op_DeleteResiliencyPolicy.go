// Code generated by smithy-go-codegen DO NOT EDIT.

package resiliencehub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a resiliency policy. This is a destructive action that can't be undone.
func (c *Client) DeleteResiliencyPolicy(ctx context.Context, params *DeleteResiliencyPolicyInput, optFns ...func(*Options)) (*DeleteResiliencyPolicyOutput, error) {
	if params == nil {
		params = &DeleteResiliencyPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteResiliencyPolicy", params, optFns, c.addOperationDeleteResiliencyPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteResiliencyPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteResiliencyPolicyInput struct {

	// The Amazon Resource Name (ARN) of the resiliency policy. The format for this
	// ARN is: arn: partition :resiliencehub: region : account :resiliency-policy/
	// policy-id . For more information about ARNs, see  Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference guide.
	//
	// This member is required.
	PolicyArn *string

	// Used for an idempotency token. A client token is a unique, case-sensitive
	// string of up to 64 ASCII characters. You should not reuse the same client token
	// for other API requests.
	ClientToken *string

	noSmithyDocumentSerde
}

type DeleteResiliencyPolicyOutput struct {

	// The Amazon Resource Name (ARN) of the resiliency policy. The format for this
	// ARN is: arn: partition :resiliencehub: region : account :resiliency-policy/
	// policy-id . For more information about ARNs, see  Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference guide.
	//
	// This member is required.
	PolicyArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteResiliencyPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteResiliencyPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteResiliencyPolicy{}, middleware.After)
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
	if err = addIdempotencyToken_opDeleteResiliencyPolicyMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteResiliencyPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteResiliencyPolicy(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpDeleteResiliencyPolicy struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDeleteResiliencyPolicy) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDeleteResiliencyPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DeleteResiliencyPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DeleteResiliencyPolicyInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opDeleteResiliencyPolicyMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpDeleteResiliencyPolicy{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteResiliencyPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resiliencehub",
		OperationName: "DeleteResiliencyPolicy",
	}
}
