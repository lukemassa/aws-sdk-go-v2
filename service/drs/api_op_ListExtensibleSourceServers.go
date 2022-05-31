// Code generated by smithy-go-codegen DO NOT EDIT.

package drs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/drs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of source servers on a staging account that are extensible, which
// means that: a. The source server is not already extended into this Account. b.
// The source server on the Account we’re reading from is not an extension of
// another source server.
func (c *Client) ListExtensibleSourceServers(ctx context.Context, params *ListExtensibleSourceServersInput, optFns ...func(*Options)) (*ListExtensibleSourceServersOutput, error) {
	if params == nil {
		params = &ListExtensibleSourceServersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListExtensibleSourceServers", params, optFns, c.addOperationListExtensibleSourceServersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListExtensibleSourceServersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListExtensibleSourceServersInput struct {

	// The Id of the staging Account to retrieve extensible source servers from.
	//
	// This member is required.
	StagingAccountID *string

	// The maximum number of extensible source servers to retrieve.
	MaxResults int32

	// The token of the next extensible source server to retrieve.
	NextToken *string

	noSmithyDocumentSerde
}

type ListExtensibleSourceServersOutput struct {

	// A list of source servers on a staging Account that are extensible.
	Items []types.StagingSourceServer

	// The token of the next extensible source server to retrieve.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListExtensibleSourceServersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListExtensibleSourceServers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListExtensibleSourceServers{}, middleware.After)
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
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListExtensibleSourceServersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListExtensibleSourceServers(options.Region), middleware.Before); err != nil {
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

// ListExtensibleSourceServersAPIClient is a client that implements the
// ListExtensibleSourceServers operation.
type ListExtensibleSourceServersAPIClient interface {
	ListExtensibleSourceServers(context.Context, *ListExtensibleSourceServersInput, ...func(*Options)) (*ListExtensibleSourceServersOutput, error)
}

var _ ListExtensibleSourceServersAPIClient = (*Client)(nil)

// ListExtensibleSourceServersPaginatorOptions is the paginator options for
// ListExtensibleSourceServers
type ListExtensibleSourceServersPaginatorOptions struct {
	// The maximum number of extensible source servers to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListExtensibleSourceServersPaginator is a paginator for
// ListExtensibleSourceServers
type ListExtensibleSourceServersPaginator struct {
	options   ListExtensibleSourceServersPaginatorOptions
	client    ListExtensibleSourceServersAPIClient
	params    *ListExtensibleSourceServersInput
	nextToken *string
	firstPage bool
}

// NewListExtensibleSourceServersPaginator returns a new
// ListExtensibleSourceServersPaginator
func NewListExtensibleSourceServersPaginator(client ListExtensibleSourceServersAPIClient, params *ListExtensibleSourceServersInput, optFns ...func(*ListExtensibleSourceServersPaginatorOptions)) *ListExtensibleSourceServersPaginator {
	if params == nil {
		params = &ListExtensibleSourceServersInput{}
	}

	options := ListExtensibleSourceServersPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListExtensibleSourceServersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListExtensibleSourceServersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListExtensibleSourceServers page.
func (p *ListExtensibleSourceServersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListExtensibleSourceServersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListExtensibleSourceServers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListExtensibleSourceServers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "drs",
		OperationName: "ListExtensibleSourceServers",
	}
}
