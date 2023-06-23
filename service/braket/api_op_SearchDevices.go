// Code generated by smithy-go-codegen DO NOT EDIT.

package braket

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/braket/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches for devices using the specified filters.
func (c *Client) SearchDevices(ctx context.Context, params *SearchDevicesInput, optFns ...func(*Options)) (*SearchDevicesOutput, error) {
	if params == nil {
		params = &SearchDevicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchDevices", params, optFns, c.addOperationSearchDevicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchDevicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchDevicesInput struct {

	// The filter values to use to search for a device.
	//
	// This member is required.
	Filters []types.SearchDevicesFilter

	// The maximum number of results to return in the response.
	MaxResults *int32

	// A token used for pagination of results returned in the response. Use the token
	// returned from the previous request continue results where the previous request
	// ended.
	NextToken *string

	noSmithyDocumentSerde
}

type SearchDevicesOutput struct {

	// An array of DeviceSummary objects for devices that match the specified filter
	// values.
	//
	// This member is required.
	Devices []types.DeviceSummary

	// A token used for pagination of results, or null if there are no additional
	// results. Use the token value in a subsequent request to continue results where
	// the previous request ended.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchDevicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchDevices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchDevices{}, middleware.After)
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
	if err = addOpSearchDevicesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchDevices(options.Region), middleware.Before); err != nil {
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

// SearchDevicesAPIClient is a client that implements the SearchDevices operation.
type SearchDevicesAPIClient interface {
	SearchDevices(context.Context, *SearchDevicesInput, ...func(*Options)) (*SearchDevicesOutput, error)
}

var _ SearchDevicesAPIClient = (*Client)(nil)

// SearchDevicesPaginatorOptions is the paginator options for SearchDevices
type SearchDevicesPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchDevicesPaginator is a paginator for SearchDevices
type SearchDevicesPaginator struct {
	options   SearchDevicesPaginatorOptions
	client    SearchDevicesAPIClient
	params    *SearchDevicesInput
	nextToken *string
	firstPage bool
}

// NewSearchDevicesPaginator returns a new SearchDevicesPaginator
func NewSearchDevicesPaginator(client SearchDevicesAPIClient, params *SearchDevicesInput, optFns ...func(*SearchDevicesPaginatorOptions)) *SearchDevicesPaginator {
	if params == nil {
		params = &SearchDevicesInput{}
	}

	options := SearchDevicesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchDevicesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchDevicesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchDevices page.
func (p *SearchDevicesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchDevicesOutput, error) {
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

	result, err := p.client.SearchDevices(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchDevices(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "braket",
		OperationName: "SearchDevices",
	}
}
