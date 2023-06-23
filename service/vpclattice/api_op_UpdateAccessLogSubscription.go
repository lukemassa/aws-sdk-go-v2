// Code generated by smithy-go-codegen DO NOT EDIT.

package vpclattice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified access log subscription.
func (c *Client) UpdateAccessLogSubscription(ctx context.Context, params *UpdateAccessLogSubscriptionInput, optFns ...func(*Options)) (*UpdateAccessLogSubscriptionOutput, error) {
	if params == nil {
		params = &UpdateAccessLogSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAccessLogSubscription", params, optFns, c.addOperationUpdateAccessLogSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAccessLogSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAccessLogSubscriptionInput struct {

	// The ID or Amazon Resource Name (ARN) of the access log subscription.
	//
	// This member is required.
	AccessLogSubscriptionIdentifier *string

	// The Amazon Resource Name (ARN) of the access log destination.
	//
	// This member is required.
	DestinationArn *string

	noSmithyDocumentSerde
}

type UpdateAccessLogSubscriptionOutput struct {

	// The Amazon Resource Name (ARN) of the access log subscription.
	//
	// This member is required.
	Arn *string

	// The Amazon Resource Name (ARN) of the access log destination.
	//
	// This member is required.
	DestinationArn *string

	// The ID of the access log subscription.
	//
	// This member is required.
	Id *string

	// The Amazon Resource Name (ARN) of the access log subscription.
	//
	// This member is required.
	ResourceArn *string

	// The ID of the resource.
	//
	// This member is required.
	ResourceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAccessLogSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAccessLogSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAccessLogSubscription{}, middleware.After)
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
	if err = addOpUpdateAccessLogSubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAccessLogSubscription(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAccessLogSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "vpc-lattice",
		OperationName: "UpdateAccessLogSubscription",
	}
}
