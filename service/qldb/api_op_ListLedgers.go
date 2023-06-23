// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns all ledgers that are associated with the current Amazon Web Services
// account and Region. This action returns a maximum of MaxResults items and is
// paginated so that you can retrieve all the items by calling ListLedgers
// multiple times.
func (c *Client) ListLedgers(ctx context.Context, params *ListLedgersInput, optFns ...func(*Options)) (*ListLedgersOutput, error) {
	if params == nil {
		params = &ListLedgersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListLedgers", params, optFns, c.addOperationListLedgersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListLedgersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLedgersInput struct {

	// The maximum number of results to return in a single ListLedgers request. (The
	// actual number of results returned might be fewer.)
	MaxResults *int32

	// A pagination token, indicating that you want to retrieve the next page of
	// results. If you received a value for NextToken in the response from a previous
	// ListLedgers call, then you should use that value as input here.
	NextToken *string

	noSmithyDocumentSerde
}

type ListLedgersOutput struct {

	// The ledgers that are associated with the current Amazon Web Services account
	// and Region.
	Ledgers []types.LedgerSummary

	// A pagination token, indicating whether there are more results available:
	//   - If NextToken is empty, then the last page of results has been processed and
	//   there are no more results to be retrieved.
	//   - If NextToken is not empty, then there are more results available. To
	//   retrieve the next page of results, use the value of NextToken in a subsequent
	//   ListLedgers call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListLedgersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListLedgers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListLedgers{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListLedgers(options.Region), middleware.Before); err != nil {
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

// ListLedgersAPIClient is a client that implements the ListLedgers operation.
type ListLedgersAPIClient interface {
	ListLedgers(context.Context, *ListLedgersInput, ...func(*Options)) (*ListLedgersOutput, error)
}

var _ ListLedgersAPIClient = (*Client)(nil)

// ListLedgersPaginatorOptions is the paginator options for ListLedgers
type ListLedgersPaginatorOptions struct {
	// The maximum number of results to return in a single ListLedgers request. (The
	// actual number of results returned might be fewer.)
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListLedgersPaginator is a paginator for ListLedgers
type ListLedgersPaginator struct {
	options   ListLedgersPaginatorOptions
	client    ListLedgersAPIClient
	params    *ListLedgersInput
	nextToken *string
	firstPage bool
}

// NewListLedgersPaginator returns a new ListLedgersPaginator
func NewListLedgersPaginator(client ListLedgersAPIClient, params *ListLedgersInput, optFns ...func(*ListLedgersPaginatorOptions)) *ListLedgersPaginator {
	if params == nil {
		params = &ListLedgersInput{}
	}

	options := ListLedgersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListLedgersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListLedgersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListLedgers page.
func (p *ListLedgersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListLedgersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListLedgers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListLedgers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "ListLedgers",
	}
}
