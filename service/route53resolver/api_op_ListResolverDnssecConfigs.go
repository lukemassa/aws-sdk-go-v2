// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the configurations for DNSSEC validation that are associated with the
// current Amazon Web Services account.
func (c *Client) ListResolverDnssecConfigs(ctx context.Context, params *ListResolverDnssecConfigsInput, optFns ...func(*Options)) (*ListResolverDnssecConfigsOutput, error) {
	if params == nil {
		params = &ListResolverDnssecConfigsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResolverDnssecConfigs", params, optFns, c.addOperationListResolverDnssecConfigsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResolverDnssecConfigsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResolverDnssecConfigsInput struct {

	// An optional specification to return a subset of objects.
	Filters []types.Filter

	// Optional: An integer that specifies the maximum number of DNSSEC configuration
	// results that you want Amazon Route 53 to return. If you don't specify a value
	// for MaxResults , Route 53 returns up to 100 configuration per page.
	MaxResults *int32

	// (Optional) If the current Amazon Web Services account has more than MaxResults
	// DNSSEC configurations, use NextToken to get the second and subsequent pages of
	// results. For the first ListResolverDnssecConfigs request, omit this value. For
	// the second and subsequent requests, get the value of NextToken from the
	// previous response and specify that value for NextToken in the request.
	NextToken *string

	noSmithyDocumentSerde
}

type ListResolverDnssecConfigsOutput struct {

	// If a response includes the last of the DNSSEC configurations that are
	// associated with the current Amazon Web Services account, NextToken doesn't
	// appear in the response. If a response doesn't include the last of the
	// configurations, you can get more configurations by submitting another
	// ListResolverDnssecConfigs (https://docs.aws.amazon.com/Route53/latest/APIReference/API_ListResolverDnssecConfigs.html)
	// request. Get the value of NextToken that Amazon Route 53 returned in the
	// previous response and include it in NextToken in the next request.
	NextToken *string

	// An array that contains one ResolverDnssecConfig (https://docs.aws.amazon.com/Route53/latest/APIReference/API_ResolverDnssecConfig.html)
	// element for each configuration for DNSSEC validation that is associated with the
	// current Amazon Web Services account. It doesn't contain disabled DNSSEC
	// configurations for the resource.
	ResolverDnssecConfigs []types.ResolverDnssecConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListResolverDnssecConfigsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListResolverDnssecConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListResolverDnssecConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListResolverDnssecConfigs"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListResolverDnssecConfigs(options.Region), middleware.Before); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

// ListResolverDnssecConfigsAPIClient is a client that implements the
// ListResolverDnssecConfigs operation.
type ListResolverDnssecConfigsAPIClient interface {
	ListResolverDnssecConfigs(context.Context, *ListResolverDnssecConfigsInput, ...func(*Options)) (*ListResolverDnssecConfigsOutput, error)
}

var _ ListResolverDnssecConfigsAPIClient = (*Client)(nil)

// ListResolverDnssecConfigsPaginatorOptions is the paginator options for
// ListResolverDnssecConfigs
type ListResolverDnssecConfigsPaginatorOptions struct {
	// Optional: An integer that specifies the maximum number of DNSSEC configuration
	// results that you want Amazon Route 53 to return. If you don't specify a value
	// for MaxResults , Route 53 returns up to 100 configuration per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListResolverDnssecConfigsPaginator is a paginator for ListResolverDnssecConfigs
type ListResolverDnssecConfigsPaginator struct {
	options   ListResolverDnssecConfigsPaginatorOptions
	client    ListResolverDnssecConfigsAPIClient
	params    *ListResolverDnssecConfigsInput
	nextToken *string
	firstPage bool
}

// NewListResolverDnssecConfigsPaginator returns a new
// ListResolverDnssecConfigsPaginator
func NewListResolverDnssecConfigsPaginator(client ListResolverDnssecConfigsAPIClient, params *ListResolverDnssecConfigsInput, optFns ...func(*ListResolverDnssecConfigsPaginatorOptions)) *ListResolverDnssecConfigsPaginator {
	if params == nil {
		params = &ListResolverDnssecConfigsInput{}
	}

	options := ListResolverDnssecConfigsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListResolverDnssecConfigsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListResolverDnssecConfigsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListResolverDnssecConfigs page.
func (p *ListResolverDnssecConfigsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListResolverDnssecConfigsOutput, error) {
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

	result, err := p.client.ListResolverDnssecConfigs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListResolverDnssecConfigs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListResolverDnssecConfigs",
	}
}
