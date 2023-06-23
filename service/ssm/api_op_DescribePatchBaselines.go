// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the patch baselines in your Amazon Web Services account.
func (c *Client) DescribePatchBaselines(ctx context.Context, params *DescribePatchBaselinesInput, optFns ...func(*Options)) (*DescribePatchBaselinesOutput, error) {
	if params == nil {
		params = &DescribePatchBaselinesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePatchBaselines", params, optFns, c.addOperationDescribePatchBaselinesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePatchBaselinesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePatchBaselinesInput struct {

	// Each element in the array is a structure containing a key-value pair. Supported
	// keys for DescribePatchBaselines include the following:
	//   - NAME_PREFIX Sample values: AWS- | My-
	//   - OWNER Sample values: AWS | Self
	//   - OPERATING_SYSTEM Sample values: AMAZON_LINUX | SUSE | WINDOWS
	Filters []types.PatchOrchestratorFilter

	// The maximum number of patch baselines to return (per page).
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribePatchBaselinesOutput struct {

	// An array of PatchBaselineIdentity elements.
	BaselineIdentities []types.PatchBaselineIdentity

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePatchBaselinesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePatchBaselines{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePatchBaselines{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePatchBaselines(options.Region), middleware.Before); err != nil {
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

// DescribePatchBaselinesAPIClient is a client that implements the
// DescribePatchBaselines operation.
type DescribePatchBaselinesAPIClient interface {
	DescribePatchBaselines(context.Context, *DescribePatchBaselinesInput, ...func(*Options)) (*DescribePatchBaselinesOutput, error)
}

var _ DescribePatchBaselinesAPIClient = (*Client)(nil)

// DescribePatchBaselinesPaginatorOptions is the paginator options for
// DescribePatchBaselines
type DescribePatchBaselinesPaginatorOptions struct {
	// The maximum number of patch baselines to return (per page).
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribePatchBaselinesPaginator is a paginator for DescribePatchBaselines
type DescribePatchBaselinesPaginator struct {
	options   DescribePatchBaselinesPaginatorOptions
	client    DescribePatchBaselinesAPIClient
	params    *DescribePatchBaselinesInput
	nextToken *string
	firstPage bool
}

// NewDescribePatchBaselinesPaginator returns a new DescribePatchBaselinesPaginator
func NewDescribePatchBaselinesPaginator(client DescribePatchBaselinesAPIClient, params *DescribePatchBaselinesInput, optFns ...func(*DescribePatchBaselinesPaginatorOptions)) *DescribePatchBaselinesPaginator {
	if params == nil {
		params = &DescribePatchBaselinesInput{}
	}

	options := DescribePatchBaselinesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribePatchBaselinesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribePatchBaselinesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribePatchBaselines page.
func (p *DescribePatchBaselinesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribePatchBaselinesOutput, error) {
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

	result, err := p.client.DescribePatchBaselines(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribePatchBaselines(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribePatchBaselines",
	}
}
