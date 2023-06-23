// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Gets information about a workflow run task.
func (c *Client) GetRunTask(ctx context.Context, params *GetRunTaskInput, optFns ...func(*Options)) (*GetRunTaskOutput, error) {
	if params == nil {
		params = &GetRunTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRunTask", params, optFns, c.addOperationGetRunTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRunTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRunTaskInput struct {

	// The task's ID.
	//
	// This member is required.
	Id *string

	// The task's ID.
	//
	// This member is required.
	TaskId *string

	noSmithyDocumentSerde
}

type GetRunTaskOutput struct {

	// The task's CPU usage.
	Cpus *int32

	// When the task was created.
	CreationTime *time.Time

	// The number of Graphics Processing Units (GPU) specified in the task.
	Gpus *int32

	// The task's log stream.
	LogStream *string

	// The task's memory use in gigabytes.
	Memory *int32

	// The task's name.
	Name *string

	// The task's start time.
	StartTime *time.Time

	// The task's status.
	Status types.TaskStatus

	// The task's status message.
	StatusMessage *string

	// The task's stop time.
	StopTime *time.Time

	// The task's ID.
	TaskId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRunTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetRunTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRunTask{}, middleware.After)
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
	if err = addEndpointPrefix_opGetRunTaskMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetRunTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRunTask(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetRunTaskMiddleware struct {
}

func (*endpointPrefix_opGetRunTaskMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetRunTaskMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "workflows-" + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opGetRunTaskMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetRunTaskMiddleware{}, `OperationSerializer`, middleware.After)
}

// GetRunTaskAPIClient is a client that implements the GetRunTask operation.
type GetRunTaskAPIClient interface {
	GetRunTask(context.Context, *GetRunTaskInput, ...func(*Options)) (*GetRunTaskOutput, error)
}

var _ GetRunTaskAPIClient = (*Client)(nil)

// TaskRunningWaiterOptions are waiter options for TaskRunningWaiter
type TaskRunningWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// TaskRunningWaiter will use default minimum delay of 30 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, TaskRunningWaiter will use default max delay of 600 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *GetRunTaskInput, *GetRunTaskOutput, error) (bool, error)
}

// TaskRunningWaiter defines the waiters for TaskRunning
type TaskRunningWaiter struct {
	client GetRunTaskAPIClient

	options TaskRunningWaiterOptions
}

// NewTaskRunningWaiter constructs a TaskRunningWaiter.
func NewTaskRunningWaiter(client GetRunTaskAPIClient, optFns ...func(*TaskRunningWaiterOptions)) *TaskRunningWaiter {
	options := TaskRunningWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 600 * time.Second
	options.Retryable = taskRunningStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &TaskRunningWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for TaskRunning waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *TaskRunningWaiter) Wait(ctx context.Context, params *GetRunTaskInput, maxWaitDur time.Duration, optFns ...func(*TaskRunningWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for TaskRunning waiter and returns the
// output of the successful operation. The maxWaitDur is the maximum wait duration
// the waiter will wait. The maxWaitDur is required and must be greater than zero.
func (w *TaskRunningWaiter) WaitForOutput(ctx context.Context, params *GetRunTaskInput, maxWaitDur time.Duration, optFns ...func(*TaskRunningWaiterOptions)) (*GetRunTaskOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 600 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.GetRunTask(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for TaskRunning waiter")
}

func taskRunningStateRetryable(ctx context.Context, input *GetRunTaskInput, output *GetRunTaskOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "RUNNING"
		value, ok := pathValue.(types.TaskStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.TaskStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "PENDING"
		value, ok := pathValue.(types.TaskStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.TaskStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "STARTING"
		value, ok := pathValue.(types.TaskStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.TaskStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "FAILED"
		value, ok := pathValue.(types.TaskStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.TaskStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "CANCELLED"
		value, ok := pathValue.(types.TaskStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.TaskStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

// TaskCompletedWaiterOptions are waiter options for TaskCompletedWaiter
type TaskCompletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// TaskCompletedWaiter will use default minimum delay of 30 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or
	// set to zero, TaskCompletedWaiter will use default max delay of 600 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *GetRunTaskInput, *GetRunTaskOutput, error) (bool, error)
}

// TaskCompletedWaiter defines the waiters for TaskCompleted
type TaskCompletedWaiter struct {
	client GetRunTaskAPIClient

	options TaskCompletedWaiterOptions
}

// NewTaskCompletedWaiter constructs a TaskCompletedWaiter.
func NewTaskCompletedWaiter(client GetRunTaskAPIClient, optFns ...func(*TaskCompletedWaiterOptions)) *TaskCompletedWaiter {
	options := TaskCompletedWaiterOptions{}
	options.MinDelay = 30 * time.Second
	options.MaxDelay = 600 * time.Second
	options.Retryable = taskCompletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &TaskCompletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for TaskCompleted waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *TaskCompletedWaiter) Wait(ctx context.Context, params *GetRunTaskInput, maxWaitDur time.Duration, optFns ...func(*TaskCompletedWaiterOptions)) error {
	_, err := w.WaitForOutput(ctx, params, maxWaitDur, optFns...)
	return err
}

// WaitForOutput calls the waiter function for TaskCompleted waiter and returns
// the output of the successful operation. The maxWaitDur is the maximum wait
// duration the waiter will wait. The maxWaitDur is required and must be greater
// than zero.
func (w *TaskCompletedWaiter) WaitForOutput(ctx context.Context, params *GetRunTaskInput, maxWaitDur time.Duration, optFns ...func(*TaskCompletedWaiterOptions)) (*GetRunTaskOutput, error) {
	if maxWaitDur <= 0 {
		return nil, fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 600 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return nil, fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.GetRunTask(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return nil, err
		}
		if !retryable {
			return out, nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return nil, fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return nil, fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return nil, fmt.Errorf("exceeded max wait time for TaskCompleted waiter")
}

func taskCompletedStateRetryable(ctx context.Context, input *GetRunTaskInput, output *GetRunTaskOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "COMPLETED"
		value, ok := pathValue.(types.TaskStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.TaskStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "PENDING"
		value, ok := pathValue.(types.TaskStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.TaskStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "STARTING"
		value, ok := pathValue.(types.TaskStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.TaskStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "RUNNING"
		value, ok := pathValue.(types.TaskStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.TaskStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "STOPPING"
		value, ok := pathValue.(types.TaskStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.TaskStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return true, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "FAILED"
		value, ok := pathValue.(types.TaskStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.TaskStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opGetRunTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "omics",
		OperationName: "GetRunTask",
	}
}
