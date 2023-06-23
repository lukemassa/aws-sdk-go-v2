// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns details about the specified delivery channel. If a delivery channel is
// not specified, this action returns the details of all delivery channels
// associated with the account. Currently, you can specify only one delivery
// channel per region in your account.
func (c *Client) DescribeDeliveryChannels(ctx context.Context, params *DescribeDeliveryChannelsInput, optFns ...func(*Options)) (*DescribeDeliveryChannelsOutput, error) {
	if params == nil {
		params = &DescribeDeliveryChannelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDeliveryChannels", params, optFns, c.addOperationDescribeDeliveryChannelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDeliveryChannelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DescribeDeliveryChannels action.
type DescribeDeliveryChannelsInput struct {

	// A list of delivery channel names.
	DeliveryChannelNames []string

	noSmithyDocumentSerde
}

// The output for the DescribeDeliveryChannels action.
type DescribeDeliveryChannelsOutput struct {

	// A list that contains the descriptions of the specified delivery channel.
	DeliveryChannels []types.DeliveryChannel

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDeliveryChannelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDeliveryChannels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDeliveryChannels{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDeliveryChannels(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDeliveryChannels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeDeliveryChannels",
	}
}
