// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Restarts a failed batch build. Only batch builds that have failed can be
// retried.
func (c *Client) RetryBuildBatch(ctx context.Context, params *RetryBuildBatchInput, optFns ...func(*Options)) (*RetryBuildBatchOutput, error) {
	if params == nil {
		params = &RetryBuildBatchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RetryBuildBatch", params, optFns, c.addOperationRetryBuildBatchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RetryBuildBatchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RetryBuildBatchInput struct {

	// Specifies the identifier of the batch build to restart.
	Id *string

	// A unique, case sensitive identifier you provide to ensure the idempotency of
	// the RetryBuildBatch request. The token is included in the RetryBuildBatch
	// request and is valid for five minutes. If you repeat the RetryBuildBatch
	// request with the same token, but change a parameter, CodeBuild returns a
	// parameter mismatch error.
	IdempotencyToken *string

	// Specifies the type of retry to perform.
	RetryType types.RetryBuildBatchType

	noSmithyDocumentSerde
}

type RetryBuildBatchOutput struct {

	// Contains information about a batch build.
	BuildBatch *types.BuildBatch

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRetryBuildBatchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRetryBuildBatch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRetryBuildBatch{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRetryBuildBatch(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRetryBuildBatch(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "RetryBuildBatch",
	}
}
