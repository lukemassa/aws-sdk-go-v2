// Code generated by smithy-go-codegen DO NOT EDIT.

package athena

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists available workgroups for the account.
func (c *Client) ListWorkGroups(ctx context.Context, params *ListWorkGroupsInput, optFns ...func(*Options)) (*ListWorkGroupsOutput, error) {
	if params == nil {
		params = &ListWorkGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkGroups", params, optFns, c.addOperationListWorkGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkGroupsInput struct {

	// The maximum number of workgroups to return in this request.
	MaxResults *int32

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListWorkGroupsOutput struct {

	// A token generated by the Athena service that specifies where to continue
	// pagination if a previous request was truncated. To obtain the next set of pages,
	// pass in the NextToken from the response object of the previous page call.
	NextToken *string

	// A list of WorkGroupSummary objects that include the names, descriptions,
	// creation times, and states for each workgroup.
	WorkGroups []types.WorkGroupSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWorkGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListWorkGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListWorkGroups{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkGroups(options.Region), middleware.Before); err != nil {
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

// ListWorkGroupsAPIClient is a client that implements the ListWorkGroups
// operation.
type ListWorkGroupsAPIClient interface {
	ListWorkGroups(context.Context, *ListWorkGroupsInput, ...func(*Options)) (*ListWorkGroupsOutput, error)
}

var _ ListWorkGroupsAPIClient = (*Client)(nil)

// ListWorkGroupsPaginatorOptions is the paginator options for ListWorkGroups
type ListWorkGroupsPaginatorOptions struct {
	// The maximum number of workgroups to return in this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorkGroupsPaginator is a paginator for ListWorkGroups
type ListWorkGroupsPaginator struct {
	options   ListWorkGroupsPaginatorOptions
	client    ListWorkGroupsAPIClient
	params    *ListWorkGroupsInput
	nextToken *string
	firstPage bool
}

// NewListWorkGroupsPaginator returns a new ListWorkGroupsPaginator
func NewListWorkGroupsPaginator(client ListWorkGroupsAPIClient, params *ListWorkGroupsInput, optFns ...func(*ListWorkGroupsPaginatorOptions)) *ListWorkGroupsPaginator {
	if params == nil {
		params = &ListWorkGroupsInput{}
	}

	options := ListWorkGroupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWorkGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorkGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWorkGroups page.
func (p *ListWorkGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorkGroupsOutput, error) {
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

	result, err := p.client.ListWorkGroups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWorkGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "athena",
		OperationName: "ListWorkGroups",
	}
}
