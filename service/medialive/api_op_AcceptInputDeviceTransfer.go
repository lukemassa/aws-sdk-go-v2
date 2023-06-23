// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Accept an incoming input device transfer. The ownership of the device will
// transfer to your AWS account.
func (c *Client) AcceptInputDeviceTransfer(ctx context.Context, params *AcceptInputDeviceTransferInput, optFns ...func(*Options)) (*AcceptInputDeviceTransferOutput, error) {
	if params == nil {
		params = &AcceptInputDeviceTransferInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AcceptInputDeviceTransfer", params, optFns, c.addOperationAcceptInputDeviceTransferMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AcceptInputDeviceTransferOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for AcceptInputDeviceTransferRequest
type AcceptInputDeviceTransferInput struct {

	// The unique ID of the input device to accept. For example, hd-123456789abcdef.
	//
	// This member is required.
	InputDeviceId *string

	noSmithyDocumentSerde
}

// Placeholder documentation for AcceptInputDeviceTransferResponse
type AcceptInputDeviceTransferOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAcceptInputDeviceTransferMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAcceptInputDeviceTransfer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAcceptInputDeviceTransfer{}, middleware.After)
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
	if err = addOpAcceptInputDeviceTransferValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAcceptInputDeviceTransfer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAcceptInputDeviceTransfer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "AcceptInputDeviceTransfer",
	}
}
