// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List import errors.
func (c *Client) ListImportErrors(ctx context.Context, params *ListImportErrorsInput, optFns ...func(*Options)) (*ListImportErrorsOutput, error) {
	if params == nil {
		params = &ListImportErrorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListImportErrors", params, optFns, c.addOperationListImportErrorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListImportErrorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// List import errors request.
type ListImportErrorsInput struct {

	// List import errors request import id.
	//
	// This member is required.
	ImportID *string

	// List import errors request max results.
	MaxResults int32

	// List import errors request next token.
	NextToken *string

	noSmithyDocumentSerde
}

// List imports errors response.
type ListImportErrorsOutput struct {

	// List imports errors response items.
	Items []types.ImportTaskError

	// List imports errors response next token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListImportErrorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListImportErrors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListImportErrors{}, middleware.After)
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
	if err = addOpListImportErrorsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListImportErrors(options.Region), middleware.Before); err != nil {
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

// ListImportErrorsAPIClient is a client that implements the ListImportErrors
// operation.
type ListImportErrorsAPIClient interface {
	ListImportErrors(context.Context, *ListImportErrorsInput, ...func(*Options)) (*ListImportErrorsOutput, error)
}

var _ ListImportErrorsAPIClient = (*Client)(nil)

// ListImportErrorsPaginatorOptions is the paginator options for ListImportErrors
type ListImportErrorsPaginatorOptions struct {
	// List import errors request max results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListImportErrorsPaginator is a paginator for ListImportErrors
type ListImportErrorsPaginator struct {
	options   ListImportErrorsPaginatorOptions
	client    ListImportErrorsAPIClient
	params    *ListImportErrorsInput
	nextToken *string
	firstPage bool
}

// NewListImportErrorsPaginator returns a new ListImportErrorsPaginator
func NewListImportErrorsPaginator(client ListImportErrorsAPIClient, params *ListImportErrorsInput, optFns ...func(*ListImportErrorsPaginatorOptions)) *ListImportErrorsPaginator {
	if params == nil {
		params = &ListImportErrorsInput{}
	}

	options := ListImportErrorsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListImportErrorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListImportErrorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListImportErrors page.
func (p *ListImportErrorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListImportErrorsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListImportErrors(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListImportErrors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgn",
		OperationName: "ListImportErrors",
	}
}
