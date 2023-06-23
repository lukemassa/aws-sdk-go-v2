// Code generated by smithy-go-codegen DO NOT EDIT.

package workdocs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workdocs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the document versions for the specified document. By default, only
// active versions are returned.
func (c *Client) DescribeDocumentVersions(ctx context.Context, params *DescribeDocumentVersionsInput, optFns ...func(*Options)) (*DescribeDocumentVersionsOutput, error) {
	if params == nil {
		params = &DescribeDocumentVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDocumentVersions", params, optFns, c.addOperationDescribeDocumentVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDocumentVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDocumentVersionsInput struct {

	// The ID of the document.
	//
	// This member is required.
	DocumentId *string

	// Amazon WorkDocs authentication token. Not required when using Amazon Web
	// Services administrator credentials to access the API.
	AuthenticationToken *string

	// Specify "SOURCE" to include initialized versions and a URL for the source
	// document.
	Fields *string

	// A comma-separated list of values. Specify "INITIALIZED" to include incomplete
	// versions.
	Include *string

	// The maximum number of versions to return with this call.
	Limit *int32

	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string

	noSmithyDocumentSerde
}

type DescribeDocumentVersionsOutput struct {

	// The document versions.
	DocumentVersions []types.DocumentVersionMetadata

	// The marker to use when requesting the next set of results. If there are no
	// additional results, the string is empty.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDocumentVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDocumentVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDocumentVersions{}, middleware.After)
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
	if err = addOpDescribeDocumentVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDocumentVersions(options.Region), middleware.Before); err != nil {
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

// DescribeDocumentVersionsAPIClient is a client that implements the
// DescribeDocumentVersions operation.
type DescribeDocumentVersionsAPIClient interface {
	DescribeDocumentVersions(context.Context, *DescribeDocumentVersionsInput, ...func(*Options)) (*DescribeDocumentVersionsOutput, error)
}

var _ DescribeDocumentVersionsAPIClient = (*Client)(nil)

// DescribeDocumentVersionsPaginatorOptions is the paginator options for
// DescribeDocumentVersions
type DescribeDocumentVersionsPaginatorOptions struct {
	// The maximum number of versions to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDocumentVersionsPaginator is a paginator for DescribeDocumentVersions
type DescribeDocumentVersionsPaginator struct {
	options   DescribeDocumentVersionsPaginatorOptions
	client    DescribeDocumentVersionsAPIClient
	params    *DescribeDocumentVersionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeDocumentVersionsPaginator returns a new
// DescribeDocumentVersionsPaginator
func NewDescribeDocumentVersionsPaginator(client DescribeDocumentVersionsAPIClient, params *DescribeDocumentVersionsInput, optFns ...func(*DescribeDocumentVersionsPaginatorOptions)) *DescribeDocumentVersionsPaginator {
	if params == nil {
		params = &DescribeDocumentVersionsInput{}
	}

	options := DescribeDocumentVersionsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDocumentVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDocumentVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDocumentVersions page.
func (p *DescribeDocumentVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDocumentVersionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeDocumentVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeDocumentVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "DescribeDocumentVersions",
	}
}
