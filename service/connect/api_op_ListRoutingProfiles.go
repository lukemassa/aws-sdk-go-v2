// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides summary information about the routing profiles for the specified
// Amazon Connect instance. For more information about routing profiles, see
// Routing Profiles (https://docs.aws.amazon.com/connect/latest/adminguide/concepts-routing.html)
// and Create a Routing Profile (https://docs.aws.amazon.com/connect/latest/adminguide/routing-profiles.html)
// in the Amazon Connect Administrator Guide.
func (c *Client) ListRoutingProfiles(ctx context.Context, params *ListRoutingProfilesInput, optFns ...func(*Options)) (*ListRoutingProfilesOutput, error) {
	if params == nil {
		params = &ListRoutingProfilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRoutingProfiles", params, optFns, c.addOperationListRoutingProfilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRoutingProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRoutingProfilesInput struct {

	// The identifier of the Amazon Connect instance. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// The maximum number of results to return per page. The default MaxResult size is
	// 100.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRoutingProfilesOutput struct {

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Information about the routing profiles.
	RoutingProfileSummaryList []types.RoutingProfileSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRoutingProfilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRoutingProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRoutingProfiles{}, middleware.After)
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
	if err = addOpListRoutingProfilesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRoutingProfiles(options.Region), middleware.Before); err != nil {
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

// ListRoutingProfilesAPIClient is a client that implements the
// ListRoutingProfiles operation.
type ListRoutingProfilesAPIClient interface {
	ListRoutingProfiles(context.Context, *ListRoutingProfilesInput, ...func(*Options)) (*ListRoutingProfilesOutput, error)
}

var _ ListRoutingProfilesAPIClient = (*Client)(nil)

// ListRoutingProfilesPaginatorOptions is the paginator options for
// ListRoutingProfiles
type ListRoutingProfilesPaginatorOptions struct {
	// The maximum number of results to return per page. The default MaxResult size is
	// 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRoutingProfilesPaginator is a paginator for ListRoutingProfiles
type ListRoutingProfilesPaginator struct {
	options   ListRoutingProfilesPaginatorOptions
	client    ListRoutingProfilesAPIClient
	params    *ListRoutingProfilesInput
	nextToken *string
	firstPage bool
}

// NewListRoutingProfilesPaginator returns a new ListRoutingProfilesPaginator
func NewListRoutingProfilesPaginator(client ListRoutingProfilesAPIClient, params *ListRoutingProfilesInput, optFns ...func(*ListRoutingProfilesPaginatorOptions)) *ListRoutingProfilesPaginator {
	if params == nil {
		params = &ListRoutingProfilesInput{}
	}

	options := ListRoutingProfilesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRoutingProfilesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRoutingProfilesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRoutingProfiles page.
func (p *ListRoutingProfilesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRoutingProfilesOutput, error) {
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

	result, err := p.client.ListRoutingProfiles(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRoutingProfiles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListRoutingProfiles",
	}
}
