// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Accepts a structured query language (SQL) SELECT command and an aggregator to
// query configuration state of Amazon Web Services resources across multiple
// accounts and regions, performs the corresponding search, and returns resource
// configurations matching the properties. For more information about query
// components, see the Query Components  (https://docs.aws.amazon.com/config/latest/developerguide/query-components.html)
// section in the Config Developer Guide. If you run an aggregation query (i.e.,
// using GROUP BY or using aggregate functions such as COUNT ; e.g., SELECT
// resourceId, COUNT(*) WHERE resourceType = 'AWS::IAM::Role' GROUP BY resourceId )
// and do not specify the MaxResults or the Limit query parameters, the default
// page size is set to 500. If you run a non-aggregation query (i.e., not using
// GROUP BY or aggregate function; e.g., SELECT * WHERE resourceType =
// 'AWS::IAM::Role' ) and do not specify the MaxResults or the Limit query
// parameters, the default page size is set to 25.
func (c *Client) SelectAggregateResourceConfig(ctx context.Context, params *SelectAggregateResourceConfigInput, optFns ...func(*Options)) (*SelectAggregateResourceConfigOutput, error) {
	if params == nil {
		params = &SelectAggregateResourceConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SelectAggregateResourceConfig", params, optFns, c.addOperationSelectAggregateResourceConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SelectAggregateResourceConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SelectAggregateResourceConfigInput struct {

	// The name of the configuration aggregator.
	//
	// This member is required.
	ConfigurationAggregatorName *string

	// The SQL query SELECT command.
	//
	// This member is required.
	Expression *string

	// The maximum number of query results returned on each page.
	Limit int32

	// The maximum number of query results returned on each page. Config also allows
	// the Limit request parameter.
	MaxResults int32

	// The nextToken string returned in a previous request that you use to request the
	// next page of results in a paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

type SelectAggregateResourceConfigOutput struct {

	// The nextToken string returned in a previous request that you use to request the
	// next page of results in a paginated response.
	NextToken *string

	// Details about the query.
	QueryInfo *types.QueryInfo

	// Returns the results for the SQL query.
	Results []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSelectAggregateResourceConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSelectAggregateResourceConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSelectAggregateResourceConfig{}, middleware.After)
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
	if err = addOpSelectAggregateResourceConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSelectAggregateResourceConfig(options.Region), middleware.Before); err != nil {
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

// SelectAggregateResourceConfigAPIClient is a client that implements the
// SelectAggregateResourceConfig operation.
type SelectAggregateResourceConfigAPIClient interface {
	SelectAggregateResourceConfig(context.Context, *SelectAggregateResourceConfigInput, ...func(*Options)) (*SelectAggregateResourceConfigOutput, error)
}

var _ SelectAggregateResourceConfigAPIClient = (*Client)(nil)

// SelectAggregateResourceConfigPaginatorOptions is the paginator options for
// SelectAggregateResourceConfig
type SelectAggregateResourceConfigPaginatorOptions struct {
	// The maximum number of query results returned on each page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SelectAggregateResourceConfigPaginator is a paginator for
// SelectAggregateResourceConfig
type SelectAggregateResourceConfigPaginator struct {
	options   SelectAggregateResourceConfigPaginatorOptions
	client    SelectAggregateResourceConfigAPIClient
	params    *SelectAggregateResourceConfigInput
	nextToken *string
	firstPage bool
}

// NewSelectAggregateResourceConfigPaginator returns a new
// SelectAggregateResourceConfigPaginator
func NewSelectAggregateResourceConfigPaginator(client SelectAggregateResourceConfigAPIClient, params *SelectAggregateResourceConfigInput, optFns ...func(*SelectAggregateResourceConfigPaginatorOptions)) *SelectAggregateResourceConfigPaginator {
	if params == nil {
		params = &SelectAggregateResourceConfigInput{}
	}

	options := SelectAggregateResourceConfigPaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SelectAggregateResourceConfigPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SelectAggregateResourceConfigPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SelectAggregateResourceConfig page.
func (p *SelectAggregateResourceConfigPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SelectAggregateResourceConfigOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.SelectAggregateResourceConfig(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSelectAggregateResourceConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "SelectAggregateResourceConfig",
	}
}
