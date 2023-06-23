// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakergeospatial

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemakergeospatial/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Use this operation to copy results of a Vector Enrichment job to an Amazon S3
// location.
func (c *Client) ExportVectorEnrichmentJob(ctx context.Context, params *ExportVectorEnrichmentJobInput, optFns ...func(*Options)) (*ExportVectorEnrichmentJobOutput, error) {
	if params == nil {
		params = &ExportVectorEnrichmentJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportVectorEnrichmentJob", params, optFns, c.addOperationExportVectorEnrichmentJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportVectorEnrichmentJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportVectorEnrichmentJobInput struct {

	// The Amazon Resource Name (ARN) of the Vector Enrichment job.
	//
	// This member is required.
	Arn *string

	// The Amazon Resource Name (ARN) of the IAM rolewith permission to upload to the
	// location in OutputConfig.
	//
	// This member is required.
	ExecutionRoleArn *string

	// Output location information for exporting Vector Enrichment Job results.
	//
	// This member is required.
	OutputConfig *types.ExportVectorEnrichmentJobOutputConfig

	// A unique token that guarantees that the call to this API is idempotent.
	ClientToken *string

	noSmithyDocumentSerde
}

type ExportVectorEnrichmentJobOutput struct {

	// The Amazon Resource Name (ARN) of the Vector Enrichment job being exported.
	//
	// This member is required.
	Arn *string

	// The creation time.
	//
	// This member is required.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the IAM role with permission to upload to the
	// location in OutputConfig.
	//
	// This member is required.
	ExecutionRoleArn *string

	// The status of the results the Vector Enrichment job being exported.
	//
	// This member is required.
	ExportStatus types.VectorEnrichmentJobExportStatus

	// Output location information for exporting Vector Enrichment Job results.
	//
	// This member is required.
	OutputConfig *types.ExportVectorEnrichmentJobOutputConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExportVectorEnrichmentJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpExportVectorEnrichmentJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpExportVectorEnrichmentJob{}, middleware.After)
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
	if err = addIdempotencyToken_opExportVectorEnrichmentJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpExportVectorEnrichmentJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExportVectorEnrichmentJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpExportVectorEnrichmentJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpExportVectorEnrichmentJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpExportVectorEnrichmentJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*ExportVectorEnrichmentJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *ExportVectorEnrichmentJobInput ")
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
func addIdempotencyToken_opExportVectorEnrichmentJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpExportVectorEnrichmentJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opExportVectorEnrichmentJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker-geospatial",
		OperationName: "ExportVectorEnrichmentJob",
	}
}
