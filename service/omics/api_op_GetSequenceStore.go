// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets information about a sequence store.
func (c *Client) GetSequenceStore(ctx context.Context, params *GetSequenceStoreInput, optFns ...func(*Options)) (*GetSequenceStoreOutput, error) {
	if params == nil {
		params = &GetSequenceStoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSequenceStore", params, optFns, c.addOperationGetSequenceStoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSequenceStoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSequenceStoreInput struct {

	// The store's ID.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type GetSequenceStoreOutput struct {

	// The store's ARN.
	//
	// This member is required.
	Arn *string

	// When the store was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The store's ID.
	//
	// This member is required.
	Id *string

	// The store's description.
	Description *string

	// An S3 location that is used to store files that have failed a direct upload.
	FallbackLocation *string

	// The store's name.
	Name *string

	// The store's server-side encryption (SSE) settings.
	SseConfig *types.SseConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSequenceStoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSequenceStore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSequenceStore{}, middleware.After)
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
	if err = addEndpointPrefix_opGetSequenceStoreMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetSequenceStoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSequenceStore(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetSequenceStoreMiddleware struct {
}

func (*endpointPrefix_opGetSequenceStoreMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetSequenceStoreMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "control-storage-" + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opGetSequenceStoreMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetSequenceStoreMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opGetSequenceStore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "omics",
		OperationName: "GetSequenceStore",
	}
}
