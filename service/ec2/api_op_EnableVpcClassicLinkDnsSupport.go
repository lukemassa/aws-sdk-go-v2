// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// We are retiring EC2-Classic. We recommend that you migrate from EC2-Classic to
// a VPC. For more information, see Migrate from EC2-Classic to a VPC (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-migrate.html)
// in the Amazon Elastic Compute Cloud User Guide. Enables a VPC to support DNS
// hostname resolution for ClassicLink. If enabled, the DNS hostname of a linked
// EC2-Classic instance resolves to its private IP address when addressed from an
// instance in the VPC to which it's linked. Similarly, the DNS hostname of an
// instance in a VPC resolves to its private IP address when addressed from a
// linked EC2-Classic instance. For more information, see ClassicLink (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html)
// in the Amazon Elastic Compute Cloud User Guide. You must specify a VPC ID in the
// request.
func (c *Client) EnableVpcClassicLinkDnsSupport(ctx context.Context, params *EnableVpcClassicLinkDnsSupportInput, optFns ...func(*Options)) (*EnableVpcClassicLinkDnsSupportOutput, error) {
	if params == nil {
		params = &EnableVpcClassicLinkDnsSupportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableVpcClassicLinkDnsSupport", params, optFns, c.addOperationEnableVpcClassicLinkDnsSupportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableVpcClassicLinkDnsSupportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableVpcClassicLinkDnsSupportInput struct {

	// The ID of the VPC.
	VpcId *string

	noSmithyDocumentSerde
}

type EnableVpcClassicLinkDnsSupportOutput struct {

	// Returns true if the request succeeds; otherwise, it returns an error.
	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableVpcClassicLinkDnsSupportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpEnableVpcClassicLinkDnsSupport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpEnableVpcClassicLinkDnsSupport{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableVpcClassicLinkDnsSupport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableVpcClassicLinkDnsSupport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "EnableVpcClassicLinkDnsSupport",
	}
}
