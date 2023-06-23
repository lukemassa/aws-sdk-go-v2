// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchain

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of proposals for the network. Applies only to Hyperledger Fabric.
func (c *Client) ListProposals(ctx context.Context, params *ListProposalsInput, optFns ...func(*Options)) (*ListProposalsOutput, error) {
	if params == nil {
		params = &ListProposalsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProposals", params, optFns, c.addOperationListProposalsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProposalsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProposalsInput struct {

	// The unique identifier of the network.
	//
	// This member is required.
	NetworkId *string

	// The maximum number of proposals to return.
	MaxResults *int32

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	noSmithyDocumentSerde
}

type ListProposalsOutput struct {

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	// The summary of each proposal made on the network.
	Proposals []types.ProposalSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProposalsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListProposals{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListProposals{}, middleware.After)
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
	if err = addOpListProposalsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProposals(options.Region), middleware.Before); err != nil {
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

// ListProposalsAPIClient is a client that implements the ListProposals operation.
type ListProposalsAPIClient interface {
	ListProposals(context.Context, *ListProposalsInput, ...func(*Options)) (*ListProposalsOutput, error)
}

var _ ListProposalsAPIClient = (*Client)(nil)

// ListProposalsPaginatorOptions is the paginator options for ListProposals
type ListProposalsPaginatorOptions struct {
	// The maximum number of proposals to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProposalsPaginator is a paginator for ListProposals
type ListProposalsPaginator struct {
	options   ListProposalsPaginatorOptions
	client    ListProposalsAPIClient
	params    *ListProposalsInput
	nextToken *string
	firstPage bool
}

// NewListProposalsPaginator returns a new ListProposalsPaginator
func NewListProposalsPaginator(client ListProposalsAPIClient, params *ListProposalsInput, optFns ...func(*ListProposalsPaginatorOptions)) *ListProposalsPaginator {
	if params == nil {
		params = &ListProposalsInput{}
	}

	options := ListProposalsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProposalsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProposalsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProposals page.
func (p *ListProposalsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProposalsOutput, error) {
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

	result, err := p.client.ListProposals(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListProposals(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "managedblockchain",
		OperationName: "ListProposals",
	}
}
