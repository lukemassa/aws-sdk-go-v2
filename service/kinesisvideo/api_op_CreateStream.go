// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new Kinesis video stream. When you create a new stream, Kinesis Video
// Streams assigns it a version number. When you change the stream's metadata,
// Kinesis Video Streams updates the version. CreateStream is an asynchronous
// operation. For information about how the service works, see How it Works (https://docs.aws.amazon.com/kinesisvideostreams/latest/dg/how-it-works.html)
// . You must have permissions for the KinesisVideo:CreateStream action.
func (c *Client) CreateStream(ctx context.Context, params *CreateStreamInput, optFns ...func(*Options)) (*CreateStreamOutput, error) {
	if params == nil {
		params = &CreateStreamInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStream", params, optFns, c.addOperationCreateStreamMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStreamOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateStreamInput struct {

	// A name for the stream that you are creating. The stream name is an identifier
	// for the stream, and must be unique for each account and region.
	//
	// This member is required.
	StreamName *string

	// The number of hours that you want to retain the data in the stream. Kinesis
	// Video Streams retains the data in a data store that is associated with the
	// stream. The default value is 0, indicating that the stream does not persist
	// data. When the DataRetentionInHours value is 0, consumers can still consume the
	// fragments that remain in the service host buffer, which has a retention time
	// limit of 5 minutes and a retention memory limit of 200 MB. Fragments are removed
	// from the buffer when either limit is reached.
	DataRetentionInHours *int32

	// The name of the device that is writing to the stream. In the current
	// implementation, Kinesis Video Streams does not use this name.
	DeviceName *string

	// The ID of the Key Management Service (KMS) key that you want Kinesis Video
	// Streams to use to encrypt stream data. If no key ID is specified, the default,
	// Kinesis Video-managed key ( aws/kinesisvideo ) is used. For more information,
	// see DescribeKey (https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters)
	// .
	KmsKeyId *string

	// The media type of the stream. Consumers of the stream can use this information
	// when processing the stream. For more information about media types, see Media
	// Types (http://www.iana.org/assignments/media-types/media-types.xhtml) . If you
	// choose to specify the MediaType , see Naming Requirements (https://tools.ietf.org/html/rfc6838#section-4.2)
	// for guidelines. Example valid values include "video/h264" and
	// "video/h264,audio/aac". This parameter is optional; the default value is null
	// (or empty in JSON).
	MediaType *string

	// A list of tags to associate with the specified stream. Each tag is a key-value
	// pair (the value is optional).
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateStreamOutput struct {

	// The Amazon Resource Name (ARN) of the stream.
	StreamARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateStreamMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateStream{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateStream{}, middleware.After)
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
	if err = addOpCreateStreamValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStream(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateStream(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "CreateStream",
	}
}
