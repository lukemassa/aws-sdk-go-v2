// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of upgrade compatible Elastisearch versions. You can optionally
// pass a DomainName to get all upgrade compatible Elasticsearch versions for that
// specific domain.
func (c *Client) GetCompatibleElasticsearchVersions(ctx context.Context, params *GetCompatibleElasticsearchVersionsInput, optFns ...func(*Options)) (*GetCompatibleElasticsearchVersionsOutput, error) {
	if params == nil {
		params = &GetCompatibleElasticsearchVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCompatibleElasticsearchVersions", params, optFns, c.addOperationGetCompatibleElasticsearchVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCompatibleElasticsearchVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for request parameters to GetCompatibleElasticsearchVersions
// operation.
type GetCompatibleElasticsearchVersionsInput struct {

	// The name of an Elasticsearch domain. Domain names are unique across the domains
	// owned by an account within an AWS region. Domain names start with a letter or
	// number and can contain the following characters: a-z (lowercase), 0-9, and -
	// (hyphen).
	DomainName *string

	noSmithyDocumentSerde
}

// Container for response returned by GetCompatibleElasticsearchVersions operation.
type GetCompatibleElasticsearchVersionsOutput struct {

	// A map of compatible Elasticsearch versions returned as part of the
	// GetCompatibleElasticsearchVersions operation.
	CompatibleElasticsearchVersions []types.CompatibleVersionsMap

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetCompatibleElasticsearchVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCompatibleElasticsearchVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCompatibleElasticsearchVersions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCompatibleElasticsearchVersions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCompatibleElasticsearchVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "GetCompatibleElasticsearchVersions",
	}
}
