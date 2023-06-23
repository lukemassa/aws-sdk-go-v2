// Code generated by smithy-go-codegen DO NOT EDIT.

package devopsguru

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devopsguru/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns lists Amazon Web Services resources that are of the specified resource
// collection type. The two types of Amazon Web Services resource collections
// supported are Amazon Web Services CloudFormation stacks and Amazon Web Services
// resources that contain the same Amazon Web Services tag. DevOps Guru can be
// configured to analyze the Amazon Web Services resources that are defined in the
// stacks or that are tagged using the same tag key. You can specify up to 500
// Amazon Web Services CloudFormation stacks.
func (c *Client) GetResourceCollection(ctx context.Context, params *GetResourceCollectionInput, optFns ...func(*Options)) (*GetResourceCollectionOutput, error) {
	if params == nil {
		params = &GetResourceCollectionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResourceCollection", params, optFns, c.addOperationGetResourceCollectionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResourceCollectionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResourceCollectionInput struct {

	// The type of Amazon Web Services resource collections to return. The one valid
	// value is CLOUD_FORMATION for Amazon Web Services CloudFormation stacks.
	//
	// This member is required.
	ResourceCollectionType types.ResourceCollectionType

	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string

	noSmithyDocumentSerde
}

type GetResourceCollectionOutput struct {

	// The pagination token to use to retrieve the next page of results for this
	// operation. If there are no more pages, this value is null.
	NextToken *string

	// The requested list of Amazon Web Services resource collections. The two types
	// of Amazon Web Services resource collections supported are Amazon Web Services
	// CloudFormation stacks and Amazon Web Services resources that contain the same
	// Amazon Web Services tag. DevOps Guru can be configured to analyze the Amazon Web
	// Services resources that are defined in the stacks or that are tagged using the
	// same tag key. You can specify up to 500 Amazon Web Services CloudFormation
	// stacks.
	ResourceCollection *types.ResourceCollectionFilter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetResourceCollectionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetResourceCollection{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetResourceCollection{}, middleware.After)
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
	if err = addOpGetResourceCollectionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetResourceCollection(options.Region), middleware.Before); err != nil {
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

// GetResourceCollectionAPIClient is a client that implements the
// GetResourceCollection operation.
type GetResourceCollectionAPIClient interface {
	GetResourceCollection(context.Context, *GetResourceCollectionInput, ...func(*Options)) (*GetResourceCollectionOutput, error)
}

var _ GetResourceCollectionAPIClient = (*Client)(nil)

// GetResourceCollectionPaginatorOptions is the paginator options for
// GetResourceCollection
type GetResourceCollectionPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetResourceCollectionPaginator is a paginator for GetResourceCollection
type GetResourceCollectionPaginator struct {
	options   GetResourceCollectionPaginatorOptions
	client    GetResourceCollectionAPIClient
	params    *GetResourceCollectionInput
	nextToken *string
	firstPage bool
}

// NewGetResourceCollectionPaginator returns a new GetResourceCollectionPaginator
func NewGetResourceCollectionPaginator(client GetResourceCollectionAPIClient, params *GetResourceCollectionInput, optFns ...func(*GetResourceCollectionPaginatorOptions)) *GetResourceCollectionPaginator {
	if params == nil {
		params = &GetResourceCollectionInput{}
	}

	options := GetResourceCollectionPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetResourceCollectionPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetResourceCollectionPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetResourceCollection page.
func (p *GetResourceCollectionPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetResourceCollectionOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.GetResourceCollection(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetResourceCollection(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devops-guru",
		OperationName: "GetResourceCollection",
	}
}
