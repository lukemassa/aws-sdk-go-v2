// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of aliases for the specified bot.
func (c *Client) ListBotAliases(ctx context.Context, params *ListBotAliasesInput, optFns ...func(*Options)) (*ListBotAliasesOutput, error) {
	if params == nil {
		params = &ListBotAliasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBotAliases", params, optFns, c.addOperationListBotAliasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBotAliasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBotAliasesInput struct {

	// The identifier of the bot to list aliases for.
	//
	// This member is required.
	BotId *string

	// The maximum number of aliases to return in each page of results. If there are
	// fewer results than the max page size, only the actual number of results are
	// returned.
	MaxResults *int32

	// If the response from the ListBotAliases operation contains more results than
	// specified in the maxResults parameter, a token is returned in the response. Use
	// that token in the nextToken parameter to return the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListBotAliasesOutput struct {

	// Summary information for the bot aliases that meet the filter criteria specified
	// in the request. The length of the list is specified in the maxResults parameter
	// of the request. If there are more aliases available, the nextToken field
	// contains a token to get the next page of results.
	BotAliasSummaries []types.BotAliasSummary

	// The identifier of the bot associated with the aliases.
	BotId *string

	// A token that indicates whether there are more results to return in a response
	// to the ListBotAliases operation. If the nextToken field is present, you send
	// the contents as the nextToken parameter of a ListBotAliases operation request
	// to get the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBotAliasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBotAliases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBotAliases{}, middleware.After)
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
	if err = addOpListBotAliasesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBotAliases(options.Region), middleware.Before); err != nil {
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

// ListBotAliasesAPIClient is a client that implements the ListBotAliases
// operation.
type ListBotAliasesAPIClient interface {
	ListBotAliases(context.Context, *ListBotAliasesInput, ...func(*Options)) (*ListBotAliasesOutput, error)
}

var _ ListBotAliasesAPIClient = (*Client)(nil)

// ListBotAliasesPaginatorOptions is the paginator options for ListBotAliases
type ListBotAliasesPaginatorOptions struct {
	// The maximum number of aliases to return in each page of results. If there are
	// fewer results than the max page size, only the actual number of results are
	// returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBotAliasesPaginator is a paginator for ListBotAliases
type ListBotAliasesPaginator struct {
	options   ListBotAliasesPaginatorOptions
	client    ListBotAliasesAPIClient
	params    *ListBotAliasesInput
	nextToken *string
	firstPage bool
}

// NewListBotAliasesPaginator returns a new ListBotAliasesPaginator
func NewListBotAliasesPaginator(client ListBotAliasesAPIClient, params *ListBotAliasesInput, optFns ...func(*ListBotAliasesPaginatorOptions)) *ListBotAliasesPaginator {
	if params == nil {
		params = &ListBotAliasesInput{}
	}

	options := ListBotAliasesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBotAliasesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBotAliasesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBotAliases page.
func (p *ListBotAliasesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBotAliasesOutput, error) {
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

	result, err := p.client.ListBotAliases(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBotAliases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "ListBotAliases",
	}
}
