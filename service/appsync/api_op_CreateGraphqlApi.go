// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a GraphqlApi object.
func (c *Client) CreateGraphqlApi(ctx context.Context, params *CreateGraphqlApiInput, optFns ...func(*Options)) (*CreateGraphqlApiOutput, error) {
	if params == nil {
		params = &CreateGraphqlApiInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateGraphqlApi", params, optFns, c.addOperationCreateGraphqlApiMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateGraphqlApiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGraphqlApiInput struct {

	// The authentication type: API key, Identity and Access Management (IAM), OpenID
	// Connect (OIDC), Amazon Cognito user pools, or Lambda.
	//
	// This member is required.
	AuthenticationType types.AuthenticationType

	// A user-supplied name for the GraphqlApi .
	//
	// This member is required.
	Name *string

	// A list of additional authentication providers for the GraphqlApi API.
	AdditionalAuthenticationProviders []types.AdditionalAuthenticationProvider

	// The value that indicates whether the GraphQL API is a standard API ( GRAPHQL )
	// or merged API ( MERGED ).
	ApiType types.GraphQLApiType

	// Configuration for Lambda function authorization.
	LambdaAuthorizerConfig *types.LambdaAuthorizerConfig

	// The Amazon CloudWatch Logs configuration.
	LogConfig *types.LogConfig

	// The Identity and Access Management service role ARN for a merged API. The
	// AppSync service assumes this role on behalf of the Merged API to validate access
	// to source APIs at runtime and to prompt the AUTO_MERGE to update the merged API
	// endpoint with the source API changes automatically.
	MergedApiExecutionRoleArn *string

	// The OIDC configuration.
	OpenIDConnectConfig *types.OpenIDConnectConfig

	// The owner contact information for an API resource. This field accepts any
	// string input with a length of 0 - 256 characters.
	OwnerContact *string

	// A TagMap object.
	Tags map[string]string

	// The Amazon Cognito user pool configuration.
	UserPoolConfig *types.UserPoolConfig

	// Sets the value of the GraphQL API to public ( GLOBAL ) or private ( PRIVATE ).
	// If no value is provided, the visibility will be set to GLOBAL by default. This
	// value cannot be changed once the API has been created.
	Visibility types.GraphQLApiVisibility

	// A flag indicating whether to use X-Ray tracing for the GraphqlApi .
	XrayEnabled bool

	noSmithyDocumentSerde
}

type CreateGraphqlApiOutput struct {

	// The GraphqlApi .
	GraphqlApi *types.GraphqlApi

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateGraphqlApiMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateGraphqlApi{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateGraphqlApi{}, middleware.After)
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
	if err = addOpCreateGraphqlApiValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGraphqlApi(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateGraphqlApi(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appsync",
		OperationName: "CreateGraphqlApi",
	}
}
