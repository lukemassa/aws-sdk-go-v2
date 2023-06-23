// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get the field-level encryption profile information.
func (c *Client) GetFieldLevelEncryptionProfile(ctx context.Context, params *GetFieldLevelEncryptionProfileInput, optFns ...func(*Options)) (*GetFieldLevelEncryptionProfileOutput, error) {
	if params == nil {
		params = &GetFieldLevelEncryptionProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFieldLevelEncryptionProfile", params, optFns, c.addOperationGetFieldLevelEncryptionProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFieldLevelEncryptionProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFieldLevelEncryptionProfileInput struct {

	// Get the ID for the field-level encryption profile information.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type GetFieldLevelEncryptionProfileOutput struct {

	// The current version of the field level encryption profile. For example:
	// E2QWRUHAPOMQZL .
	ETag *string

	// Return the field-level encryption profile information.
	FieldLevelEncryptionProfile *types.FieldLevelEncryptionProfile

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFieldLevelEncryptionProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetFieldLevelEncryptionProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetFieldLevelEncryptionProfile{}, middleware.After)
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
	if err = addOpGetFieldLevelEncryptionProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFieldLevelEncryptionProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetFieldLevelEncryptionProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "GetFieldLevelEncryptionProfile",
	}
}
