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

// Query a set of OpsItems. You must have permission in Identity and Access
// Management (IAM) to query a list of OpsItems. For more information, see Getting
// started with OpsCenter (https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-getting-started.html)
// in the Amazon Web Services Systems Manager User Guide. Operations engineers and
// IT professionals use Amazon Web Services Systems Manager OpsCenter to view,
// investigate, and remediate operational issues impacting the performance and
// health of their Amazon Web Services resources. For more information, see
// OpsCenter (https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter.html)
// in the Amazon Web Services Systems Manager User Guide.
func (c *Client) DescribeOpsItems(ctx context.Context, params *DescribeOpsItemsInput, optFns ...func(*Options)) (*DescribeOpsItemsOutput, error) {
	if params == nil {
		params = &DescribeOpsItemsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOpsItems", params, optFns, c.addOperationDescribeOpsItemsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOpsItemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOpsItemsInput struct {

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string

	// One or more filters to limit the response.
	//   - Key: CreatedTime Operations: GreaterThan, LessThan
	//   - Key: LastModifiedBy Operations: Contains, Equals
	//   - Key: LastModifiedTime Operations: GreaterThan, LessThan
	//   - Key: Priority Operations: Equals
	//   - Key: Source Operations: Contains, Equals
	//   - Key: Status Operations: Equals
	//   - Key: Title* Operations: Equals,Contains
	//   - Key: OperationalData** Operations: Equals
	//   - Key: OperationalDataKey Operations: Equals
	//   - Key: OperationalDataValue Operations: Equals, Contains
	//   - Key: OpsItemId Operations: Equals
	//   - Key: ResourceId Operations: Contains
	//   - Key: AutomationId Operations: Equals
	// *The Equals operator for Title matches the first 100 characters. If you specify
	// more than 100 characters, they system returns an error that the filter value
	// exceeds the length limit. **If you filter the response by using the
	// OperationalData operator, specify a key-value pair by using the following JSON
	// format: {"key":"key_name","value":"a_value"}
	OpsItemFilters []types.OpsItemFilter

	noSmithyDocumentSerde
}

type DescribeOpsItemsOutput struct {

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// A list of OpsItems.
	OpsItemSummaries []types.OpsItemSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeOpsItemsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeOpsItems{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeOpsItems{}, middleware.After)
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
	if err = addOpDescribeOpsItemsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOpsItems(options.Region), middleware.Before); err != nil {
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

// DescribeOpsItemsAPIClient is a client that implements the DescribeOpsItems
// operation.
type DescribeOpsItemsAPIClient interface {
	DescribeOpsItems(context.Context, *DescribeOpsItemsInput, ...func(*Options)) (*DescribeOpsItemsOutput, error)
}

var _ DescribeOpsItemsAPIClient = (*Client)(nil)

// DescribeOpsItemsPaginatorOptions is the paginator options for DescribeOpsItems
type DescribeOpsItemsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeOpsItemsPaginator is a paginator for DescribeOpsItems
type DescribeOpsItemsPaginator struct {
	options   DescribeOpsItemsPaginatorOptions
	client    DescribeOpsItemsAPIClient
	params    *DescribeOpsItemsInput
	nextToken *string
	firstPage bool
}

// NewDescribeOpsItemsPaginator returns a new DescribeOpsItemsPaginator
func NewDescribeOpsItemsPaginator(client DescribeOpsItemsAPIClient, params *DescribeOpsItemsInput, optFns ...func(*DescribeOpsItemsPaginatorOptions)) *DescribeOpsItemsPaginator {
	if params == nil {
		params = &DescribeOpsItemsInput{}
	}

	options := DescribeOpsItemsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeOpsItemsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeOpsItemsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeOpsItems page.
func (p *DescribeOpsItemsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeOpsItemsOutput, error) {
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

	result, err := p.client.DescribeOpsItems(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeOpsItems(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeOpsItems",
	}
}
