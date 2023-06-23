// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates the associated KMS key from the specified log group. After the
// KMS key is disassociated from the log group, CloudWatch Logs stops encrypting
// newly ingested data for the log group. All previously ingested data remains
// encrypted, and CloudWatch Logs requires permissions for the KMS key whenever the
// encrypted data is requested. Note that it can take up to 5 minutes for this
// operation to take effect.
func (c *Client) DisassociateKmsKey(ctx context.Context, params *DisassociateKmsKeyInput, optFns ...func(*Options)) (*DisassociateKmsKeyOutput, error) {
	if params == nil {
		params = &DisassociateKmsKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateKmsKey", params, optFns, c.addOperationDisassociateKmsKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateKmsKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateKmsKeyInput struct {

	// The name of the log group.
	//
	// This member is required.
	LogGroupName *string

	noSmithyDocumentSerde
}

type DisassociateKmsKeyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateKmsKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateKmsKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateKmsKey{}, middleware.After)
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
	if err = addOpDisassociateKmsKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateKmsKey(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateKmsKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "DisassociateKmsKey",
	}
}
