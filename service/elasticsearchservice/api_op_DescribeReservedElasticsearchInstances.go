// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about reserved Elasticsearch instances for this account.
func (c *Client) DescribeReservedElasticsearchInstances(ctx context.Context, params *DescribeReservedElasticsearchInstancesInput, optFns ...func(*Options)) (*DescribeReservedElasticsearchInstancesOutput, error) {
	if params == nil {
		params = &DescribeReservedElasticsearchInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReservedElasticsearchInstances", params, optFns, c.addOperationDescribeReservedElasticsearchInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReservedElasticsearchInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for parameters to DescribeReservedElasticsearchInstances
type DescribeReservedElasticsearchInstancesInput struct {

	// Set this value to limit the number of results returned. If not specified,
	// defaults to 100.
	MaxResults int32

	// NextToken should be sent in case if earlier API call produced result containing
	// NextToken. It is used for pagination.
	NextToken *string

	// The reserved instance identifier filter value. Use this parameter to show only
	// the reservation that matches the specified reserved Elasticsearch instance ID.
	ReservedElasticsearchInstanceId *string

	noSmithyDocumentSerde
}

// Container for results from DescribeReservedElasticsearchInstances
type DescribeReservedElasticsearchInstancesOutput struct {

	// Provides an identifier to allow retrieval of paginated results.
	NextToken *string

	// List of reserved Elasticsearch instances.
	ReservedElasticsearchInstances []types.ReservedElasticsearchInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReservedElasticsearchInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeReservedElasticsearchInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeReservedElasticsearchInstances{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReservedElasticsearchInstances(options.Region), middleware.Before); err != nil {
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

// DescribeReservedElasticsearchInstancesAPIClient is a client that implements the
// DescribeReservedElasticsearchInstances operation.
type DescribeReservedElasticsearchInstancesAPIClient interface {
	DescribeReservedElasticsearchInstances(context.Context, *DescribeReservedElasticsearchInstancesInput, ...func(*Options)) (*DescribeReservedElasticsearchInstancesOutput, error)
}

var _ DescribeReservedElasticsearchInstancesAPIClient = (*Client)(nil)

// DescribeReservedElasticsearchInstancesPaginatorOptions is the paginator options
// for DescribeReservedElasticsearchInstances
type DescribeReservedElasticsearchInstancesPaginatorOptions struct {
	// Set this value to limit the number of results returned. If not specified,
	// defaults to 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeReservedElasticsearchInstancesPaginator is a paginator for
// DescribeReservedElasticsearchInstances
type DescribeReservedElasticsearchInstancesPaginator struct {
	options   DescribeReservedElasticsearchInstancesPaginatorOptions
	client    DescribeReservedElasticsearchInstancesAPIClient
	params    *DescribeReservedElasticsearchInstancesInput
	nextToken *string
	firstPage bool
}

// NewDescribeReservedElasticsearchInstancesPaginator returns a new
// DescribeReservedElasticsearchInstancesPaginator
func NewDescribeReservedElasticsearchInstancesPaginator(client DescribeReservedElasticsearchInstancesAPIClient, params *DescribeReservedElasticsearchInstancesInput, optFns ...func(*DescribeReservedElasticsearchInstancesPaginatorOptions)) *DescribeReservedElasticsearchInstancesPaginator {
	if params == nil {
		params = &DescribeReservedElasticsearchInstancesInput{}
	}

	options := DescribeReservedElasticsearchInstancesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeReservedElasticsearchInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeReservedElasticsearchInstancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeReservedElasticsearchInstances page.
func (p *DescribeReservedElasticsearchInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeReservedElasticsearchInstancesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeReservedElasticsearchInstances(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeReservedElasticsearchInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DescribeReservedElasticsearchInstances",
	}
}
