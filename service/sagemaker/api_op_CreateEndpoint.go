// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an endpoint using the endpoint configuration specified in the request.
// SageMaker uses the endpoint to provision resources and deploy models. You create
// the endpoint configuration with the CreateEndpointConfig (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateEndpointConfig.html)
// API. Use this API to deploy models using SageMaker hosting services. For an
// example that calls this method when deploying a model to SageMaker hosting
// services, see the Create Endpoint example notebook. (https://github.com/aws/amazon-sagemaker-examples/blob/master/sagemaker-fundamentals/create-endpoint/create_endpoint.ipynb)
// You must not delete an EndpointConfig that is in use by an endpoint that is
// live or while the UpdateEndpoint or CreateEndpoint operations are being
// performed on the endpoint. To update an endpoint, you must create a new
// EndpointConfig . The endpoint name must be unique within an Amazon Web Services
// Region in your Amazon Web Services account. When it receives the request,
// SageMaker creates the endpoint, launches the resources (ML compute instances),
// and deploys the model(s) on them. When you call CreateEndpoint (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateEndpoint.html)
// , a load call is made to DynamoDB to verify that your endpoint configuration
// exists. When you read data from a DynamoDB table supporting Eventually
// Consistent Reads (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadConsistency.html)
// , the response might not reflect the results of a recently completed write
// operation. The response might include some stale data. If the dependent entities
// are not yet in DynamoDB, this causes a validation error. If you repeat your read
// request after a short time, the response should return the latest data. So retry
// logic is recommended to handle these possible issues. We also recommend that
// customers call DescribeEndpointConfig (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeEndpointConfig.html)
// before calling CreateEndpoint (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateEndpoint.html)
// to minimize the potential impact of a DynamoDB eventually consistent read. When
// SageMaker receives the request, it sets the endpoint status to Creating . After
// it creates the endpoint, it sets the status to InService . SageMaker can then
// process incoming requests for inferences. To check the status of an endpoint,
// use the DescribeEndpoint (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeEndpoint.html)
// API. If any of the models hosted at this endpoint get model data from an Amazon
// S3 location, SageMaker uses Amazon Web Services Security Token Service to
// download model artifacts from the S3 path you provided. Amazon Web Services STS
// is activated in your Amazon Web Services account by default. If you previously
// deactivated Amazon Web Services STS for a region, you need to reactivate Amazon
// Web Services STS for that region. For more information, see Activating and
// Deactivating Amazon Web Services STS in an Amazon Web Services Region (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html)
// in the Amazon Web Services Identity and Access Management User Guide. To add the
// IAM role policies for using this API operation, go to the IAM console (https://console.aws.amazon.com/iam/)
// , and choose Roles in the left navigation pane. Search the IAM role that you
// want to grant access to use the CreateEndpoint (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateEndpoint.html)
// and CreateEndpointConfig (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateEndpointConfig.html)
// API operations, add the following policies to the role.
//   - Option 1: For a full SageMaker access, search and attach the
//     AmazonSageMakerFullAccess policy.
//   - Option 2: For granting a limited access to an IAM role, paste the following
//     Action elements manually into the JSON file of the IAM role: "Action":
//     ["sagemaker:CreateEndpoint", "sagemaker:CreateEndpointConfig"] "Resource": [
//     "arn:aws:sagemaker:region:account-id:endpoint/endpointName"
//     "arn:aws:sagemaker:region:account-id:endpoint-config/endpointConfigName" ] For
//     more information, see SageMaker API Permissions: Actions, Permissions, and
//     Resources Reference (https://docs.aws.amazon.com/sagemaker/latest/dg/api-permissions-reference.html)
//     .
func (c *Client) CreateEndpoint(ctx context.Context, params *CreateEndpointInput, optFns ...func(*Options)) (*CreateEndpointOutput, error) {
	if params == nil {
		params = &CreateEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEndpoint", params, optFns, c.addOperationCreateEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEndpointInput struct {

	// The name of an endpoint configuration. For more information, see
	// CreateEndpointConfig (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateEndpointConfig.html)
	// .
	//
	// This member is required.
	EndpointConfigName *string

	// The name of the endpoint.The name must be unique within an Amazon Web Services
	// Region in your Amazon Web Services account. The name is case-insensitive in
	// CreateEndpoint , but the case is preserved and must be matched in InvokeEndpoint (https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_runtime_InvokeEndpoint.html)
	// .
	//
	// This member is required.
	EndpointName *string

	// The deployment configuration for an endpoint, which contains the desired
	// deployment strategy and rollback configurations.
	DeploymentConfig *types.DeploymentConfig

	// An array of key-value pairs. You can use tags to categorize your Amazon Web
	// Services resources in different ways, for example, by purpose, owner, or
	// environment. For more information, see Tagging Amazon Web Services Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// .
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateEndpointOutput struct {

	// The Amazon Resource Name (ARN) of the endpoint.
	//
	// This member is required.
	EndpointArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEndpoint{}, middleware.After)
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
	if err = addOpCreateEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateEndpoint",
	}
}
