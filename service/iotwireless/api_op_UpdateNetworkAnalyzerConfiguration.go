// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update network analyzer configuration.
func (c *Client) UpdateNetworkAnalyzerConfiguration(ctx context.Context, params *UpdateNetworkAnalyzerConfigurationInput, optFns ...func(*Options)) (*UpdateNetworkAnalyzerConfigurationOutput, error) {
	if params == nil {
		params = &UpdateNetworkAnalyzerConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateNetworkAnalyzerConfiguration", params, optFns, c.addOperationUpdateNetworkAnalyzerConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateNetworkAnalyzerConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateNetworkAnalyzerConfigurationInput struct {

	// Name of the network analyzer configuration.
	//
	// This member is required.
	ConfigurationName *string

	// The description of the new resource.
	Description *string

	// Multicast group resources to add to the network analyzer configuration. Provide
	// the MulticastGroupId of the resource to add in the input array.
	MulticastGroupsToAdd []string

	// Multicast group resources to remove from the network analyzer configuration.
	// Provide the MulticastGroupId of the resource to remove in the input array.
	MulticastGroupsToRemove []string

	// Trace content for your wireless gateway and wireless device resources.
	TraceContent *types.TraceContent

	// Wireless device resources to add to the network analyzer configuration. Provide
	// the WirelessDeviceId of the resource to add in the input array.
	WirelessDevicesToAdd []string

	// Wireless device resources to remove from the network analyzer configuration.
	// Provide the WirelessDeviceId of the resources to remove in the input array.
	WirelessDevicesToRemove []string

	// Wireless gateway resources to add to the network analyzer configuration.
	// Provide the WirelessGatewayId of the resource to add in the input array.
	WirelessGatewaysToAdd []string

	// Wireless gateway resources to remove from the network analyzer configuration.
	// Provide the WirelessGatewayId of the resources to remove in the input array.
	WirelessGatewaysToRemove []string

	noSmithyDocumentSerde
}

type UpdateNetworkAnalyzerConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateNetworkAnalyzerConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateNetworkAnalyzerConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateNetworkAnalyzerConfiguration{}, middleware.After)
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
	if err = addOpUpdateNetworkAnalyzerConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateNetworkAnalyzerConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateNetworkAnalyzerConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "UpdateNetworkAnalyzerConfiguration",
	}
}
