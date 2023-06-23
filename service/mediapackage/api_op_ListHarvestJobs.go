// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackage

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a collection of HarvestJob records.
func (c *Client) ListHarvestJobs(ctx context.Context, params *ListHarvestJobsInput, optFns ...func(*Options)) (*ListHarvestJobsOutput, error) {
	if params == nil {
		params = &ListHarvestJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListHarvestJobs", params, optFns, c.addOperationListHarvestJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListHarvestJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListHarvestJobsInput struct {

	// When specified, the request will return only HarvestJobs associated with the
	// given Channel ID.
	IncludeChannelId *string

	// When specified, the request will return only HarvestJobs in the given status.
	IncludeStatus *string

	// The upper bound on the number of records to return.
	MaxResults int32

	// A token used to resume pagination from the end of a previous request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListHarvestJobsOutput struct {

	// A list of HarvestJob records.
	HarvestJobs []types.HarvestJob

	// A token that can be used to resume pagination from the end of the collection.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListHarvestJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListHarvestJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListHarvestJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListHarvestJobs(options.Region), middleware.Before); err != nil {
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

// ListHarvestJobsAPIClient is a client that implements the ListHarvestJobs
// operation.
type ListHarvestJobsAPIClient interface {
	ListHarvestJobs(context.Context, *ListHarvestJobsInput, ...func(*Options)) (*ListHarvestJobsOutput, error)
}

var _ ListHarvestJobsAPIClient = (*Client)(nil)

// ListHarvestJobsPaginatorOptions is the paginator options for ListHarvestJobs
type ListHarvestJobsPaginatorOptions struct {
	// The upper bound on the number of records to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListHarvestJobsPaginator is a paginator for ListHarvestJobs
type ListHarvestJobsPaginator struct {
	options   ListHarvestJobsPaginatorOptions
	client    ListHarvestJobsAPIClient
	params    *ListHarvestJobsInput
	nextToken *string
	firstPage bool
}

// NewListHarvestJobsPaginator returns a new ListHarvestJobsPaginator
func NewListHarvestJobsPaginator(client ListHarvestJobsAPIClient, params *ListHarvestJobsInput, optFns ...func(*ListHarvestJobsPaginatorOptions)) *ListHarvestJobsPaginator {
	if params == nil {
		params = &ListHarvestJobsInput{}
	}

	options := ListHarvestJobsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListHarvestJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListHarvestJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListHarvestJobs page.
func (p *ListHarvestJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListHarvestJobsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListHarvestJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListHarvestJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage",
		OperationName: "ListHarvestJobs",
	}
}
