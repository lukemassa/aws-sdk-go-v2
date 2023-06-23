// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all PrincipalARN s and corresponding PrincipalType s associated with the
// specified portfolio.
func (c *Client) ListPrincipalsForPortfolio(ctx context.Context, params *ListPrincipalsForPortfolioInput, optFns ...func(*Options)) (*ListPrincipalsForPortfolioOutput, error) {
	if params == nil {
		params = &ListPrincipalsForPortfolioInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPrincipalsForPortfolio", params, optFns, c.addOperationListPrincipalsForPortfolioMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPrincipalsForPortfolioOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPrincipalsForPortfolioInput struct {

	// The portfolio identifier.
	//
	// This member is required.
	PortfolioId *string

	// The language code.
	//   - jp - Japanese
	//   - zh - Chinese
	AcceptLanguage *string

	// The maximum number of items to return with this call.
	PageSize int32

	// The page token for the next set of results. To retrieve the first set of
	// results, use null.
	PageToken *string

	noSmithyDocumentSerde
}

type ListPrincipalsForPortfolioOutput struct {

	// The page token to use to retrieve the next set of results. If there are no
	// additional results, this value is null.
	NextPageToken *string

	// The PrincipalARN s and corresponding PrincipalType s associated with the
	// portfolio.
	Principals []types.Principal

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPrincipalsForPortfolioMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPrincipalsForPortfolio{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPrincipalsForPortfolio{}, middleware.After)
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
	if err = addOpListPrincipalsForPortfolioValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPrincipalsForPortfolio(options.Region), middleware.Before); err != nil {
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

// ListPrincipalsForPortfolioAPIClient is a client that implements the
// ListPrincipalsForPortfolio operation.
type ListPrincipalsForPortfolioAPIClient interface {
	ListPrincipalsForPortfolio(context.Context, *ListPrincipalsForPortfolioInput, ...func(*Options)) (*ListPrincipalsForPortfolioOutput, error)
}

var _ ListPrincipalsForPortfolioAPIClient = (*Client)(nil)

// ListPrincipalsForPortfolioPaginatorOptions is the paginator options for
// ListPrincipalsForPortfolio
type ListPrincipalsForPortfolioPaginatorOptions struct {
	// The maximum number of items to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPrincipalsForPortfolioPaginator is a paginator for
// ListPrincipalsForPortfolio
type ListPrincipalsForPortfolioPaginator struct {
	options   ListPrincipalsForPortfolioPaginatorOptions
	client    ListPrincipalsForPortfolioAPIClient
	params    *ListPrincipalsForPortfolioInput
	nextToken *string
	firstPage bool
}

// NewListPrincipalsForPortfolioPaginator returns a new
// ListPrincipalsForPortfolioPaginator
func NewListPrincipalsForPortfolioPaginator(client ListPrincipalsForPortfolioAPIClient, params *ListPrincipalsForPortfolioInput, optFns ...func(*ListPrincipalsForPortfolioPaginatorOptions)) *ListPrincipalsForPortfolioPaginator {
	if params == nil {
		params = &ListPrincipalsForPortfolioInput{}
	}

	options := ListPrincipalsForPortfolioPaginatorOptions{}
	if params.PageSize != 0 {
		options.Limit = params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPrincipalsForPortfolioPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.PageToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPrincipalsForPortfolioPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPrincipalsForPortfolio page.
func (p *ListPrincipalsForPortfolioPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPrincipalsForPortfolioOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PageToken = p.nextToken

	params.PageSize = p.options.Limit

	result, err := p.client.ListPrincipalsForPortfolio(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextPageToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListPrincipalsForPortfolio(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "ListPrincipalsForPortfolio",
	}
}
