// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a paginated list of rulesets for the specified list of Glue tables.
func (c *Client) ListDataQualityRulesets(ctx context.Context, params *ListDataQualityRulesetsInput, optFns ...func(*Options)) (*ListDataQualityRulesetsOutput, error) {
	if params == nil {
		params = &ListDataQualityRulesetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDataQualityRulesets", params, optFns, c.addOperationListDataQualityRulesetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDataQualityRulesetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDataQualityRulesetsInput struct {

	// The filter criteria.
	Filter *types.DataQualityRulesetFilterCriteria

	// The maximum number of results to return.
	MaxResults *int32

	// A paginated token to offset the results.
	NextToken *string

	// A list of key-value pair tags.
	Tags map[string]string

	noSmithyDocumentSerde
}

type ListDataQualityRulesetsOutput struct {

	// A pagination token, if more results are available.
	NextToken *string

	// A paginated list of rulesets for the specified list of Glue tables.
	Rulesets []types.DataQualityRulesetListDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDataQualityRulesetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDataQualityRulesets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDataQualityRulesets{}, middleware.After)
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
	if err = addOpListDataQualityRulesetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDataQualityRulesets(options.Region), middleware.Before); err != nil {
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

// ListDataQualityRulesetsAPIClient is a client that implements the
// ListDataQualityRulesets operation.
type ListDataQualityRulesetsAPIClient interface {
	ListDataQualityRulesets(context.Context, *ListDataQualityRulesetsInput, ...func(*Options)) (*ListDataQualityRulesetsOutput, error)
}

var _ ListDataQualityRulesetsAPIClient = (*Client)(nil)

// ListDataQualityRulesetsPaginatorOptions is the paginator options for
// ListDataQualityRulesets
type ListDataQualityRulesetsPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDataQualityRulesetsPaginator is a paginator for ListDataQualityRulesets
type ListDataQualityRulesetsPaginator struct {
	options   ListDataQualityRulesetsPaginatorOptions
	client    ListDataQualityRulesetsAPIClient
	params    *ListDataQualityRulesetsInput
	nextToken *string
	firstPage bool
}

// NewListDataQualityRulesetsPaginator returns a new
// ListDataQualityRulesetsPaginator
func NewListDataQualityRulesetsPaginator(client ListDataQualityRulesetsAPIClient, params *ListDataQualityRulesetsInput, optFns ...func(*ListDataQualityRulesetsPaginatorOptions)) *ListDataQualityRulesetsPaginator {
	if params == nil {
		params = &ListDataQualityRulesetsInput{}
	}

	options := ListDataQualityRulesetsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDataQualityRulesetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDataQualityRulesetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDataQualityRulesets page.
func (p *ListDataQualityRulesetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDataQualityRulesetsOutput, error) {
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

	result, err := p.client.ListDataQualityRulesets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDataQualityRulesets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "ListDataQualityRulesets",
	}
}
