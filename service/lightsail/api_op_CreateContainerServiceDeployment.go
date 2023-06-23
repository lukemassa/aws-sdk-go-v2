// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a deployment for your Amazon Lightsail container service. A deployment
// specifies the containers that will be launched on the container service and
// their settings, such as the ports to open, the environment variables to apply,
// and the launch command to run. It also specifies the container that will serve
// as the public endpoint of the deployment and its settings, such as the HTTP or
// HTTPS port to use, and the health check configuration. You can deploy containers
// to your container service using container images from a public registry such as
// Amazon ECR Public, or from your local machine. For more information, see
// Creating container images for your Amazon Lightsail container services (https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-creating-container-images)
// in the Amazon Lightsail Developer Guide.
func (c *Client) CreateContainerServiceDeployment(ctx context.Context, params *CreateContainerServiceDeploymentInput, optFns ...func(*Options)) (*CreateContainerServiceDeploymentOutput, error) {
	if params == nil {
		params = &CreateContainerServiceDeploymentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateContainerServiceDeployment", params, optFns, c.addOperationCreateContainerServiceDeploymentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateContainerServiceDeploymentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateContainerServiceDeploymentInput struct {

	// The name of the container service for which to create the deployment.
	//
	// This member is required.
	ServiceName *string

	// An object that describes the settings of the containers that will be launched
	// on the container service.
	Containers map[string]types.Container

	// An object that describes the settings of the public endpoint for the container
	// service.
	PublicEndpoint *types.EndpointRequest

	noSmithyDocumentSerde
}

type CreateContainerServiceDeploymentOutput struct {

	// An object that describes a container service.
	ContainerService *types.ContainerService

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateContainerServiceDeploymentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateContainerServiceDeployment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateContainerServiceDeployment{}, middleware.After)
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
	if err = addOpCreateContainerServiceDeploymentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateContainerServiceDeployment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateContainerServiceDeployment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CreateContainerServiceDeployment",
	}
}
