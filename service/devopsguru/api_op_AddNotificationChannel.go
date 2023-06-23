// Code generated by smithy-go-codegen DO NOT EDIT.

package devopsguru

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devopsguru/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a notification channel to DevOps Guru. A notification channel is used to
// notify you about important DevOps Guru events, such as when an insight is
// generated. If you use an Amazon SNS topic in another account, you must attach a
// policy to it that grants DevOps Guru permission to it notifications. DevOps Guru
// adds the required policy on your behalf to send notifications using Amazon SNS
// in your account. DevOps Guru only supports standard SNS topics. For more
// information, see Permissions for cross account Amazon SNS topics (https://docs.aws.amazon.com/devops-guru/latest/userguide/sns-required-permissions.html)
// . If you use an Amazon SNS topic in another account, you must attach a policy to
// it that grants DevOps Guru permission to it notifications. DevOps Guru adds the
// required policy on your behalf to send notifications using Amazon SNS in your
// account. For more information, see Permissions for cross account Amazon SNS
// topics. If you use an Amazon SNS topic that is encrypted by an Amazon Web
// Services Key Management Service customer-managed key (CMK), then you must add
// permissions to the CMK. For more information, see Permissions for Amazon Web
// Services KMS–encrypted Amazon SNS topics (https://docs.aws.amazon.com/devops-guru/latest/userguide/sns-kms-permissions.html)
// .
func (c *Client) AddNotificationChannel(ctx context.Context, params *AddNotificationChannelInput, optFns ...func(*Options)) (*AddNotificationChannelOutput, error) {
	if params == nil {
		params = &AddNotificationChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AddNotificationChannel", params, optFns, c.addOperationAddNotificationChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AddNotificationChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddNotificationChannelInput struct {

	// A NotificationChannelConfig object that specifies what type of notification
	// channel to add. The one supported notification channel is Amazon Simple
	// Notification Service (Amazon SNS).
	//
	// This member is required.
	Config *types.NotificationChannelConfig

	noSmithyDocumentSerde
}

type AddNotificationChannelOutput struct {

	// The ID of the added notification channel.
	//
	// This member is required.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAddNotificationChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAddNotificationChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAddNotificationChannel{}, middleware.After)
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
	if err = addOpAddNotificationChannelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAddNotificationChannel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAddNotificationChannel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devops-guru",
		OperationName: "AddNotificationChannel",
	}
}
