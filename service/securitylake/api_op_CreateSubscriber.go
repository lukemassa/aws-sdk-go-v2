// Code generated by smithy-go-codegen DO NOT EDIT.

package securitylake

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securitylake/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a subscription permission for accounts that are already enabled in
// Amazon Security Lake. You can create a subscriber with access to data in the
// current Amazon Web Services Region.
func (c *Client) CreateSubscriber(ctx context.Context, params *CreateSubscriberInput, optFns ...func(*Options)) (*CreateSubscriberOutput, error) {
	if params == nil {
		params = &CreateSubscriberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSubscriber", params, optFns, c.addOperationCreateSubscriberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSubscriberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSubscriberInput struct {

	// The supported Amazon Web Services from which logs and events are collected.
	// Security Lake supports log and event collection for natively supported Amazon
	// Web Services.
	//
	// This member is required.
	Sources []types.LogSourceResource

	// The AWS identity used to access your data.
	//
	// This member is required.
	SubscriberIdentity *types.AwsIdentity

	// The name of your Security Lake subscriber account.
	//
	// This member is required.
	SubscriberName *string

	// The Amazon S3 or Lake Formation access type.
	AccessTypes []types.AccessType

	// The description for your subscriber account in Security Lake.
	SubscriberDescription *string

	// An array of objects, one for each tag to associate with the subscriber. For
	// each tag, you must specify both a tag key and a tag value. A tag value cannot be
	// null, but it can be an empty string.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateSubscriberOutput struct {

	// Retrieve information about the subscriber created using the CreateSubscriber
	// API.
	Subscriber *types.SubscriberResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSubscriberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSubscriber{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSubscriber{}, middleware.After)
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
	if err = addOpCreateSubscriberValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSubscriber(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSubscriber(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securitylake",
		OperationName: "CreateSubscriber",
	}
}
