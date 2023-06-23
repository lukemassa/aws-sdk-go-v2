// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about a specified Resolver query logging configuration, such
// as the number of VPCs that the configuration is logging queries for and the
// location that logs are sent to.
func (c *Client) GetResolverQueryLogConfig(ctx context.Context, params *GetResolverQueryLogConfigInput, optFns ...func(*Options)) (*GetResolverQueryLogConfigOutput, error) {
	if params == nil {
		params = &GetResolverQueryLogConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResolverQueryLogConfig", params, optFns, c.addOperationGetResolverQueryLogConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResolverQueryLogConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResolverQueryLogConfigInput struct {

	// The ID of the Resolver query logging configuration that you want to get
	// information about.
	//
	// This member is required.
	ResolverQueryLogConfigId *string

	noSmithyDocumentSerde
}

type GetResolverQueryLogConfigOutput struct {

	// Information about the Resolver query logging configuration that you specified
	// in a GetQueryLogConfig request.
	ResolverQueryLogConfig *types.ResolverQueryLogConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetResolverQueryLogConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetResolverQueryLogConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetResolverQueryLogConfig{}, middleware.After)
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
	if err = addOpGetResolverQueryLogConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetResolverQueryLogConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetResolverQueryLogConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "GetResolverQueryLogConfig",
	}
}
