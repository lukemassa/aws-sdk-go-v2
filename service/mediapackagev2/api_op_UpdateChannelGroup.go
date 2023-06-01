// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackagev2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Update the specified channel group. You can edit the description on a channel
// group for easier identification later from the AWS Elemental MediaPackage
// console. You can't edit the name of the channel group. Any edits you make that
// impact the video output may not be reflected for a few minutes.
func (c *Client) UpdateChannelGroup(ctx context.Context, params *UpdateChannelGroupInput, optFns ...func(*Options)) (*UpdateChannelGroupOutput, error) {
	if params == nil {
		params = &UpdateChannelGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateChannelGroup", params, optFns, c.addOperationUpdateChannelGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateChannelGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateChannelGroupInput struct {

	// The name that describes the channel group. The name is the primary identifier
	// for the channel group, and must be unique for your account in the AWS Region.
	//
	// This member is required.
	ChannelGroupName *string

	// Any descriptive information that you want to add to the channel group for
	// future identification purposes.
	Description *string

	noSmithyDocumentSerde
}

type UpdateChannelGroupOutput struct {

	// The Amazon Resource Name (ARN) associated with the resource.
	//
	// This member is required.
	Arn *string

	// The name that describes the channel group. The name is the primary identifier
	// for the channel group, and must be unique for your account in the AWS Region.
	//
	// This member is required.
	ChannelGroupName *string

	// The date and time the channel group was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The output domain where the source stream is sent. Integrate the domain with a
	// downstream CDN (such as Amazon CloudFront) or playback device.
	//
	// This member is required.
	EgressDomain *string

	// The date and time the channel group was modified.
	//
	// This member is required.
	ModifiedAt *time.Time

	// The description for your channel group.
	Description *string

	// The comma-separated list of tag key:value pairs assigned to the channel group.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateChannelGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateChannelGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateChannelGroup{}, middleware.After)
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
	if err = addOpUpdateChannelGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateChannelGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateChannelGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackagev2",
		OperationName: "UpdateChannelGroup",
	}
}