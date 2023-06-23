// Code generated by smithy-go-codegen DO NOT EDIT.

package accessanalyzer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an archive rule for the specified analyzer. Archive rules automatically
// archive new findings that meet the criteria you define when you create the rule.
// To learn about filter keys that you can use to create an archive rule, see IAM
// Access Analyzer filter keys (https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-reference-filter-keys.html)
// in the IAM User Guide.
func (c *Client) CreateArchiveRule(ctx context.Context, params *CreateArchiveRuleInput, optFns ...func(*Options)) (*CreateArchiveRuleOutput, error) {
	if params == nil {
		params = &CreateArchiveRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateArchiveRule", params, optFns, c.addOperationCreateArchiveRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateArchiveRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates an archive rule.
type CreateArchiveRuleInput struct {

	// The name of the created analyzer.
	//
	// This member is required.
	AnalyzerName *string

	// The criteria for the rule.
	//
	// This member is required.
	Filter map[string]types.Criterion

	// The name of the rule to create.
	//
	// This member is required.
	RuleName *string

	// A client token.
	ClientToken *string

	noSmithyDocumentSerde
}

type CreateArchiveRuleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateArchiveRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateArchiveRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateArchiveRule{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateArchiveRuleMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateArchiveRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateArchiveRule(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateArchiveRule struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateArchiveRule) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateArchiveRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateArchiveRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateArchiveRuleInput ")
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
func addIdempotencyToken_opCreateArchiveRuleMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateArchiveRule{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateArchiveRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "access-analyzer",
		OperationName: "CreateArchiveRule",
	}
}
