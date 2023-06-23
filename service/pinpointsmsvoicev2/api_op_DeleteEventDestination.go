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

// Deletes an existing event destination. An event destination is a location where
// you send response information about the messages that you send. For example,
// when a message is delivered successfully, you can send information about that
// event to an Amazon CloudWatch destination, or send notifications to endpoints
// that are subscribed to an Amazon SNS topic.
func (c *Client) DeleteEventDestination(ctx context.Context, params *DeleteEventDestinationInput, optFns ...func(*Options)) (*DeleteEventDestinationOutput, error) {
	if params == nil {
		params = &DeleteEventDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteEventDestination", params, optFns, c.addOperationDeleteEventDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteEventDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteEventDestinationInput struct {

	// The name of the configuration set or the configuration set's Amazon Resource
	// Name (ARN) to remove the event destination from. The ConfigurateSetName and
	// ConfigurationSetArn can be found using the DescribeConfigurationSets action.
	//
	// This member is required.
	ConfigurationSetName *string

	// The name of the event destination to delete.
	//
	// This member is required.
	EventDestinationName *string

	noSmithyDocumentSerde
}

type DeleteEventDestinationOutput struct {

	// The Amazon Resource Name (ARN) of the configuration set.
	ConfigurationSetArn *string

	// The name of the configuration set the event destination was deleted from.
	ConfigurationSetName *string

	// The event destination object that was deleted.
	EventDestination *types.EventDestination

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteEventDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteEventDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteEventDestination{}, middleware.After)
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
	if err = addOpDeleteEventDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteEventDestination(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteEventDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "DeleteEventDestination",
	}
}
