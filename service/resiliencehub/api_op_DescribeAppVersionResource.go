// Code generated by smithy-go-codegen DO NOT EDIT.

package resiliencehub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes a resource of the AWS Resilience Hub application. This API accepts
// only one of the following parameters to descibe the resource:
//
// * resourceName
//
// *
// logicalResourceId
//
// * physicalResourceId (Along with physicalResourceId, you can
// also provide awsAccountId, and awsRegion)
func (c *Client) DescribeAppVersionResource(ctx context.Context, params *DescribeAppVersionResourceInput, optFns ...func(*Options)) (*DescribeAppVersionResourceOutput, error) {
	if params == nil {
		params = &DescribeAppVersionResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAppVersionResource", params, optFns, c.addOperationDescribeAppVersionResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAppVersionResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAppVersionResourceInput struct {

	// The Amazon Resource Name (ARN) of the AWS Resilience Hub application. The format
	// for this ARN is: arn:partition:resiliencehub:region:account:app/app-id. For more
	// information about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference guide.
	//
	// This member is required.
	AppArn *string

	// The AWS Resilience Hub application version.
	//
	// This member is required.
	AppVersion *string

	// The Amazon Web Services account that owns the physical resource.
	AwsAccountId *string

	// The Amazon Web Services region that owns the physical resource.
	AwsRegion *string

	// The logical identifier of the resource.
	LogicalResourceId *types.LogicalResourceId

	// The physical identifier of the resource.
	PhysicalResourceId *string

	// The name of the resource.
	ResourceName *string

	noSmithyDocumentSerde
}

type DescribeAppVersionResourceOutput struct {

	// The Amazon Resource Name (ARN) of the AWS Resilience Hub application. The format
	// for this ARN is: arn:partition:resiliencehub:region:account:app/app-id. For more
	// information about ARNs, see  Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the AWS General Reference guide.
	//
	// This member is required.
	AppArn *string

	// The AWS Resilience Hub application version.
	//
	// This member is required.
	AppVersion *string

	// Defines a physical resource. A physical resource is a resource that exists in
	// your account. It can be identified using an Amazon Resource Name (ARN) or a
	// Resilience Hub-native identifier.
	PhysicalResource *types.PhysicalResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAppVersionResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAppVersionResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAppVersionResource{}, middleware.After)
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
	if err = addOpDescribeAppVersionResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAppVersionResource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAppVersionResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resiliencehub",
		OperationName: "DescribeAppVersionResource",
	}
}
