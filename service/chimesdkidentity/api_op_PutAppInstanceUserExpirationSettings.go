// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkidentity

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkidentity/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the number of days before the AppInstanceUser is automatically deleted. A
// background process deletes expired AppInstanceUsers within 6 hours of
// expiration. Actual deletion times may vary. Expired AppInstanceUsers that have
// not yet been deleted appear as active, and you can update their expiration
// settings. The system honors the new settings.
func (c *Client) PutAppInstanceUserExpirationSettings(ctx context.Context, params *PutAppInstanceUserExpirationSettingsInput, optFns ...func(*Options)) (*PutAppInstanceUserExpirationSettingsOutput, error) {
	if params == nil {
		params = &PutAppInstanceUserExpirationSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAppInstanceUserExpirationSettings", params, optFns, c.addOperationPutAppInstanceUserExpirationSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAppInstanceUserExpirationSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAppInstanceUserExpirationSettingsInput struct {

	// The ARN of the AppInstanceUser .
	//
	// This member is required.
	AppInstanceUserArn *string

	// Settings that control the interval after which an AppInstanceUser is
	// automatically deleted.
	ExpirationSettings *types.ExpirationSettings

	noSmithyDocumentSerde
}

type PutAppInstanceUserExpirationSettingsOutput struct {

	// The ARN of the AppInstanceUser .
	AppInstanceUserArn *string

	// Settings that control the interval after which an AppInstanceUser is
	// automatically deleted.
	ExpirationSettings *types.ExpirationSettings

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutAppInstanceUserExpirationSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutAppInstanceUserExpirationSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutAppInstanceUserExpirationSettings{}, middleware.After)
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
	if err = addOpPutAppInstanceUserExpirationSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAppInstanceUserExpirationSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutAppInstanceUserExpirationSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "PutAppInstanceUserExpirationSettings",
	}
}
