// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmediapipelines

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmediapipelines/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A structure that contains the static configurations for a media insights
// pipeline.
func (c *Client) CreateMediaInsightsPipelineConfiguration(ctx context.Context, params *CreateMediaInsightsPipelineConfigurationInput, optFns ...func(*Options)) (*CreateMediaInsightsPipelineConfigurationOutput, error) {
	if params == nil {
		params = &CreateMediaInsightsPipelineConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMediaInsightsPipelineConfiguration", params, optFns, c.addOperationCreateMediaInsightsPipelineConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMediaInsightsPipelineConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMediaInsightsPipelineConfigurationInput struct {

	// The elements in the request, such as a processor for Amazon Transcribe or a
	// sink for a Kinesis Data Stream.
	//
	// This member is required.
	Elements []types.MediaInsightsPipelineConfigurationElement

	// The name of the media insights pipeline configuration.
	//
	// This member is required.
	MediaInsightsPipelineConfigurationName *string

	// The ARN of the role used by the service to access Amazon Web Services
	// resources, including Transcribe and Transcribe Call Analytics , on the caller’s
	// behalf.
	//
	// This member is required.
	ResourceAccessRoleArn *string

	// The unique identifier for the media insights pipeline configuration request.
	ClientRequestToken *string

	// The configuration settings for the real-time alerts in a media insights
	// pipeline configuration.
	RealTimeAlertConfiguration *types.RealTimeAlertConfiguration

	// The tags assigned to the media insights pipeline configuration.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateMediaInsightsPipelineConfigurationOutput struct {

	// The configuration settings for the media insights pipeline.
	MediaInsightsPipelineConfiguration *types.MediaInsightsPipelineConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateMediaInsightsPipelineConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMediaInsightsPipelineConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMediaInsightsPipelineConfiguration{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateMediaInsightsPipelineConfigurationMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateMediaInsightsPipelineConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMediaInsightsPipelineConfiguration(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateMediaInsightsPipelineConfiguration struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateMediaInsightsPipelineConfiguration) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateMediaInsightsPipelineConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateMediaInsightsPipelineConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateMediaInsightsPipelineConfigurationInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateMediaInsightsPipelineConfigurationMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateMediaInsightsPipelineConfiguration{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateMediaInsightsPipelineConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "CreateMediaInsightsPipelineConfiguration",
	}
}
