// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Send a reboot command to the specified input device. The device will begin
// rebooting within a few seconds of sending the command. When the reboot is
// complete, the device’s connection status will change to connected.
func (c *Client) RebootInputDevice(ctx context.Context, params *RebootInputDeviceInput, optFns ...func(*Options)) (*RebootInputDeviceOutput, error) {
	if params == nil {
		params = &RebootInputDeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RebootInputDevice", params, optFns, c.addOperationRebootInputDeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RebootInputDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to reboot an AWS Elemental device.
type RebootInputDeviceInput struct {

	// The unique ID of the input device to reboot. For example, hd-123456789abcdef.
	//
	// This member is required.
	InputDeviceId *string

	// Force a reboot of an input device. If the device is streaming, it will stop
	// streaming and begin rebooting within a few seconds of sending the command. If
	// the device was streaming prior to the reboot, the device will resume streaming
	// when the reboot completes.
	Force types.RebootInputDeviceForce

	noSmithyDocumentSerde
}

// Placeholder documentation for RebootInputDeviceResponse
type RebootInputDeviceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRebootInputDeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRebootInputDevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRebootInputDevice{}, middleware.After)
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
	if err = addOpRebootInputDeviceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRebootInputDevice(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRebootInputDevice(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "RebootInputDevice",
	}
}
