// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides information about the bootstrap actions associated with a cluster.
func (c *Client) ListBootstrapActions(ctx context.Context, params *ListBootstrapActionsInput, optFns ...func(*Options)) (*ListBootstrapActionsOutput, error) {
	if params == nil {
		params = &ListBootstrapActionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBootstrapActions", params, optFns, c.addOperationListBootstrapActionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBootstrapActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// This input determines which bootstrap actions to retrieve.
type ListBootstrapActionsInput struct {

	// The cluster identifier for the bootstrap actions to list.
	//
	// This member is required.
	ClusterId *string

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string

	noSmithyDocumentSerde
}

// This output contains the bootstrap actions detail.
type ListBootstrapActionsOutput struct {

	// The bootstrap actions associated with the cluster.
	BootstrapActions []types.Command

	// The pagination token that indicates the next set of results to retrieve.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBootstrapActionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListBootstrapActions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListBootstrapActions{}, middleware.After)
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
	if err = addOpListBootstrapActionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBootstrapActions(options.Region), middleware.Before); err != nil {
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

// ListBootstrapActionsAPIClient is a client that implements the
// ListBootstrapActions operation.
type ListBootstrapActionsAPIClient interface {
	ListBootstrapActions(context.Context, *ListBootstrapActionsInput, ...func(*Options)) (*ListBootstrapActionsOutput, error)
}

var _ ListBootstrapActionsAPIClient = (*Client)(nil)

// ListBootstrapActionsPaginatorOptions is the paginator options for
// ListBootstrapActions
type ListBootstrapActionsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBootstrapActionsPaginator is a paginator for ListBootstrapActions
type ListBootstrapActionsPaginator struct {
	options   ListBootstrapActionsPaginatorOptions
	client    ListBootstrapActionsAPIClient
	params    *ListBootstrapActionsInput
	nextToken *string
	firstPage bool
}

// NewListBootstrapActionsPaginator returns a new ListBootstrapActionsPaginator
func NewListBootstrapActionsPaginator(client ListBootstrapActionsAPIClient, params *ListBootstrapActionsInput, optFns ...func(*ListBootstrapActionsPaginatorOptions)) *ListBootstrapActionsPaginator {
	if params == nil {
		params = &ListBootstrapActionsInput{}
	}

	options := ListBootstrapActionsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBootstrapActionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBootstrapActionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBootstrapActions page.
func (p *ListBootstrapActionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBootstrapActionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	result, err := p.client.ListBootstrapActions(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListBootstrapActions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "ListBootstrapActions",
	}
}
