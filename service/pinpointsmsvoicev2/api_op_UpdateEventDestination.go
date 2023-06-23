// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing event destination in a configuration set. You can update
// the IAM role ARN for CloudWatch Logs and Kinesis Data Firehose. You can also
// enable or disable the event destination. You may want to update an event
// destination to change its matching event types or updating the destination
// resource ARN. You can't change an event destination's type between CloudWatch
// Logs, Kinesis Data Firehose, and Amazon SNS.
func (c *Client) UpdateEventDestination(ctx context.Context, params *UpdateEventDestinationInput, optFns ...func(*Options)) (*UpdateEventDestinationOutput, error) {
	if params == nil {
		params = &UpdateEventDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEventDestination", params, optFns, c.addOperationUpdateEventDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEventDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateEventDestinationInput struct {

	// The configuration set to update with the new event destination. Valid values
	// for this can be the ConfigurationSetName or ConfigurationSetArn.
	//
	// This member is required.
	ConfigurationSetName *string

	// The name to use for the event destination.
	//
	// This member is required.
	EventDestinationName *string

	// An object that contains information about an event destination that sends data
	// to CloudWatch Logs.
	CloudWatchLogsDestination *types.CloudWatchLogsDestination

	// When set to true logging is enabled.
	Enabled *bool

	// An object that contains information about an event destination for logging to
	// Kinesis Data Firehose.
	KinesisFirehoseDestination *types.KinesisFirehoseDestination

	// An array of event types that determine which events to log.
	MatchingEventTypes []types.EventType

	// An object that contains information about an event destination that sends data
	// to Amazon SNS.
	SnsDestination *types.SnsDestination

	noSmithyDocumentSerde
}

type UpdateEventDestinationOutput struct {

	// The Amazon Resource Name (ARN) for the ConfigurationSet that was updated.
	ConfigurationSetArn *string

	// The name of the configuration set.
	ConfigurationSetName *string

	// An EventDestination object containing the details of where events will be
	// logged.
	EventDestination *types.EventDestination

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateEventDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateEventDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateEventDestination{}, middleware.After)
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
	if err = addOpUpdateEventDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEventDestination(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateEventDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "UpdateEventDestination",
	}
}
