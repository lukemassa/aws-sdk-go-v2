// Code generated by smithy-go-codegen DO NOT EDIT.

package ivs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ivs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a stream key, used to initiate a stream, for the specified channel ARN.
// Note that CreateChannel creates a stream key. If you subsequently use
// CreateStreamKey on the same channel, it will fail because a stream key already
// exists and there is a limit of 1 stream key per channel. To reset the stream key
// on a channel, use DeleteStreamKey and then CreateStreamKey.
func (c *Client) CreateStreamKey(ctx context.Context, params *CreateStreamKeyInput, optFns ...func(*Options)) (*CreateStreamKeyOutput, error) {
	if params == nil {
		params = &CreateStreamKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStreamKey", params, optFns, c.addOperationCreateStreamKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStreamKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateStreamKeyInput struct {

	// ARN of the channel for which to create the stream key.
	//
	// This member is required.
	ChannelArn *string

	// Array of 1-50 maps, each of the form string:string (key:value) . See Tagging
	// Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// for more information, including restrictions that apply to tags and "Tag naming
	// limits and requirements"; Amazon IVS has no service-specific constraints beyond
	// what is documented there.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateStreamKeyOutput struct {

	// Stream key used to authenticate an RTMPS stream for ingestion.
	StreamKey *types.StreamKey

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateStreamKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateStreamKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateStreamKey{}, middleware.After)
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
	if err = addOpCreateStreamKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStreamKey(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateStreamKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ivs",
		OperationName: "CreateStreamKey",
	}
}
