// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Invokes a Lambda function. You can invoke a function synchronously (and wait for
// the response), or asynchronously. To invoke a function asynchronously, set
// InvocationType to Event. For synchronous invocation
// (https://docs.aws.amazon.com/lambda/latest/dg/invocation-sync.html), details
// about the function response, including errors, are included in the response body
// and headers. For either invocation type, you can find more information in the
// execution log
// (https://docs.aws.amazon.com/lambda/latest/dg/monitoring-functions.html) and
// trace (https://docs.aws.amazon.com/lambda/latest/dg/lambda-x-ray.html). When an
// error occurs, your function may be invoked multiple times. Retry behavior varies
// by error type, client, event source, and invocation type. For example, if you
// invoke a function asynchronously and it returns an error, Lambda executes the
// function up to two more times. For more information, see Retry Behavior
// (https://docs.aws.amazon.com/lambda/latest/dg/retries-on-errors.html). For
// asynchronous invocation
// (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html), Lambda
// adds events to a queue before sending them to your function. If your function
// does not have enough capacity to keep up with the queue, events may be lost.
// Occasionally, your function may receive the same event multiple times, even if
// no error occurs. To retain events that were not processed, configure your
// function with a dead-letter queue
// (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#dlq). The
// status code in the API response doesn't reflect function errors. Error codes are
// reserved for errors that prevent your function from executing, such as
// permissions errors, limit errors
// (https://docs.aws.amazon.com/lambda/latest/dg/limits.html), or issues with your
// function's code and configuration. For example, Lambda returns
// TooManyRequestsException if executing the function would cause you to exceed a
// concurrency limit at either the account level
// (ConcurrentInvocationLimitExceeded) or function level
// (ReservedFunctionConcurrentInvocationLimitExceeded). For functions with a long
// timeout, your client might be disconnected during synchronous invocation while
// it waits for a response. Configure your HTTP client, SDK, firewall, proxy, or
// operating system to allow for long connections with timeout or keep-alive
// settings. This operation requires permission for the lambda:InvokeFunction
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/list_awslambda.html) action.
func (c *Client) Invoke(ctx context.Context, params *InvokeInput, optFns ...func(*Options)) (*InvokeOutput, error) {
	if params == nil {
		params = &InvokeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Invoke", params, optFns, c.addOperationInvokeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InvokeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InvokeInput struct {

	// The name of the Lambda function, version, or alias. Name formats
	//
	// * Function
	// name - my-function (name-only), my-function:v1 (with alias).
	//
	// * Function ARN -
	// arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	// * Partial ARN -
	// 123456789012:function:my-function.
	//
	// You can append a version number or alias to
	// any of the formats. The length constraint applies only to the full ARN. If you
	// specify only the function name, it is limited to 64 characters in length.
	//
	// This member is required.
	FunctionName *string

	// Up to 3583 bytes of base64-encoded data about the invoking client to pass to the
	// function in the context object.
	ClientContext *string

	// Choose from the following options.
	//
	// * RequestResponse (default) - Invoke the
	// function synchronously. Keep the connection open until the function returns a
	// response or times out. The API response includes the function response and
	// additional data.
	//
	// * Event - Invoke the function asynchronously. Send events that
	// fail multiple times to the function's dead-letter queue (if it's configured).
	// The API response only includes a status code.
	//
	// * DryRun - Validate parameter
	// values and verify that the user or role has permission to invoke the function.
	InvocationType types.InvocationType

	// Set to Tail to include the execution log in the response. Applies to
	// synchronously invoked functions only.
	LogType types.LogType

	// The JSON that you want to provide to your Lambda function as input.
	Payload []byte

	// Specify a version or alias to invoke a published version of the function.
	Qualifier *string

	noSmithyDocumentSerde
}

type InvokeOutput struct {

	// The version of the function that executed. When you invoke a function with an
	// alias, this indicates which version the alias resolved to.
	ExecutedVersion *string

	// If present, indicates that an error occurred during function execution. Details
	// about the error are included in the response payload.
	FunctionError *string

	// The last 4 KB of the execution log, which is base64 encoded.
	LogResult *string

	// The response from the function, or an error object.
	Payload []byte

	// The HTTP status code is in the 200 range for a successful request. For the
	// RequestResponse invocation type, this status code is 200. For the Event
	// invocation type, this status code is 202. For the DryRun invocation type, the
	// status code is 204.
	StatusCode int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationInvokeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpInvoke{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpInvoke{}, middleware.After)
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
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpInvokeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInvoke(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opInvoke(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "Invoke",
	}
}
