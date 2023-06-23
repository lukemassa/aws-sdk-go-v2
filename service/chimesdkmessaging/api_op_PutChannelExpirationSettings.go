// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmessaging

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkmessaging/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the number of days before the channel is automatically deleted.
//   - A background process deletes expired channels within 6 hours of expiration.
//     Actual deletion times may vary.
//   - Expired channels that have not yet been deleted appear as active, and you
//     can update their expiration settings. The system honors the new settings.
//   - The x-amz-chime-bearer request header is mandatory. Use the ARN of the
//     AppInstanceUser or AppInstanceBot that makes the API call as the value in the
//     header.
func (c *Client) PutChannelExpirationSettings(ctx context.Context, params *PutChannelExpirationSettingsInput, optFns ...func(*Options)) (*PutChannelExpirationSettingsOutput, error) {
	if params == nil {
		params = &PutChannelExpirationSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutChannelExpirationSettings", params, optFns, c.addOperationPutChannelExpirationSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutChannelExpirationSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutChannelExpirationSettingsInput struct {

	// The ARN of the channel.
	//
	// This member is required.
	ChannelArn *string

	// The ARN of the AppInstanceUser or AppInstanceBot that makes the API call.
	ChimeBearer *string

	// Settings that control the interval after which a channel is deleted.
	ExpirationSettings *types.ExpirationSettings

	noSmithyDocumentSerde
}

type PutChannelExpirationSettingsOutput struct {

	// The channel ARN.
	ChannelArn *string

	// Settings that control the interval after which a channel is deleted.
	ExpirationSettings *types.ExpirationSettings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutChannelExpirationSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutChannelExpirationSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutChannelExpirationSettings{}, middleware.After)
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
	if err = addOpPutChannelExpirationSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutChannelExpirationSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutChannelExpirationSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "PutChannelExpirationSettings",
	}
}
