// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sends a signal to an Automation execution to change the current behavior or
// status of the execution.
func (c *Client) SendAutomationSignal(ctx context.Context, params *SendAutomationSignalInput, optFns ...func(*Options)) (*SendAutomationSignalOutput, error) {
	if params == nil {
		params = &SendAutomationSignalInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SendAutomationSignal", params, optFns, c.addOperationSendAutomationSignalMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SendAutomationSignalOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SendAutomationSignalInput struct {

	// The unique identifier for an existing Automation execution that you want to
	// send the signal to.
	//
	// This member is required.
	AutomationExecutionId *string

	// The type of signal to send to an Automation execution.
	//
	// This member is required.
	SignalType types.SignalType

	// The data sent with the signal. The data schema depends on the type of signal
	// used in the request. For Approve and Reject signal types, the payload is an
	// optional comment that you can send with the signal type. For example:
	// Comment="Looks good" For StartStep and Resume signal types, you must send the
	// name of the Automation step to start or resume as the payload. For example:
	// StepName="step1" For the StopStep signal type, you must send the step execution
	// ID as the payload. For example:
	// StepExecutionId="97fff367-fc5a-4299-aed8-0123456789ab"
	Payload map[string][]string

	noSmithyDocumentSerde
}

type SendAutomationSignalOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSendAutomationSignalMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSendAutomationSignal{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSendAutomationSignal{}, middleware.After)
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
	if err = addOpSendAutomationSignalValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSendAutomationSignal(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSendAutomationSignal(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "SendAutomationSignal",
	}
}
