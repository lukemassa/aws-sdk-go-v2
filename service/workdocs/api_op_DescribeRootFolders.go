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

// Describes the current user's special folders; the RootFolder and the RecycleBin
// . RootFolder is the root of user's files and folders and RecycleBin is the root
// of recycled items. This is not a valid action for SigV4 (administrative API)
// clients. This action requires an authentication token. To get an authentication
// token, register an application with Amazon WorkDocs. For more information, see
// Authentication and Access Control for User Applications (https://docs.aws.amazon.com/workdocs/latest/developerguide/wd-auth-user.html)
// in the Amazon WorkDocs Developer Guide.
func (c *Client) DescribeRootFolders(ctx context.Context, params *DescribeRootFoldersInput, optFns ...func(*Options)) (*DescribeRootFoldersOutput, error) {
	if params == nil {
		params = &DescribeRootFoldersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeRootFolders", params, optFns, c.addOperationDescribeRootFoldersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeRootFoldersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRootFoldersInput struct {

	// Amazon WorkDocs authentication token.
	//
	// This member is required.
	AuthenticationToken *string

	// The maximum number of items to return.
	Limit *int32

	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string

	noSmithyDocumentSerde
}

type DescribeRootFoldersOutput struct {

	// The user's special folders.
	Folders []types.FolderMetadata

	// The marker for the next set of results.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeRootFoldersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeRootFolders{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeRootFolders{}, middleware.After)
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
	if err = addOpDescribeRootFoldersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRootFolders(options.Region), middleware.Before); err != nil {
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

// DescribeRootFoldersAPIClient is a client that implements the
// DescribeRootFolders operation.
type DescribeRootFoldersAPIClient interface {
	DescribeRootFolders(context.Context, *DescribeRootFoldersInput, ...func(*Options)) (*DescribeRootFoldersOutput, error)
}

var _ DescribeRootFoldersAPIClient = (*Client)(nil)

// DescribeRootFoldersPaginatorOptions is the paginator options for
// DescribeRootFolders
type DescribeRootFoldersPaginatorOptions struct {
	// The maximum number of items to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeRootFoldersPaginator is a paginator for DescribeRootFolders
type DescribeRootFoldersPaginator struct {
	options   DescribeRootFoldersPaginatorOptions
	client    DescribeRootFoldersAPIClient
	params    *DescribeRootFoldersInput
	nextToken *string
	firstPage bool
}

// NewDescribeRootFoldersPaginator returns a new DescribeRootFoldersPaginator
func NewDescribeRootFoldersPaginator(client DescribeRootFoldersAPIClient, params *DescribeRootFoldersInput, optFns ...func(*DescribeRootFoldersPaginatorOptions)) *DescribeRootFoldersPaginator {
	if params == nil {
		params = &DescribeRootFoldersInput{}
	}

	options := DescribeRootFoldersPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeRootFoldersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeRootFoldersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeRootFolders page.
func (p *DescribeRootFoldersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeRootFoldersOutput, error) {
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

	result, err := p.client.DescribeRootFolders(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeRootFolders(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workdocs",
		OperationName: "DescribeRootFolders",
	}
}
