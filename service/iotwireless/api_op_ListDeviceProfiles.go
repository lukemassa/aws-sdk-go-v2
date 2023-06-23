// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the device profiles registered to your AWS account.
func (c *Client) ListDeviceProfiles(ctx context.Context, params *ListDeviceProfilesInput, optFns ...func(*Options)) (*ListDeviceProfilesOutput, error) {
	if params == nil {
		params = &ListDeviceProfilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDeviceProfiles", params, optFns, c.addOperationListDeviceProfilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDeviceProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDeviceProfilesInput struct {

	// A filter to list only device profiles that use this type, which can be LoRaWAN
	// or Sidewalk .
	DeviceProfileType types.DeviceProfileType

	// The maximum number of results to return in this operation.
	MaxResults int32

	// To retrieve the next set of results, the nextToken value from a previous
	// response; otherwise null to receive the first set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDeviceProfilesOutput struct {

	// The list of device profiles.
	DeviceProfileList []types.DeviceProfile

	// The token to use to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDeviceProfilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDeviceProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDeviceProfiles{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDeviceProfiles(options.Region), middleware.Before); err != nil {
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

// ListDeviceProfilesAPIClient is a client that implements the ListDeviceProfiles
// operation.
type ListDeviceProfilesAPIClient interface {
	ListDeviceProfiles(context.Context, *ListDeviceProfilesInput, ...func(*Options)) (*ListDeviceProfilesOutput, error)
}

var _ ListDeviceProfilesAPIClient = (*Client)(nil)

// ListDeviceProfilesPaginatorOptions is the paginator options for
// ListDeviceProfiles
type ListDeviceProfilesPaginatorOptions struct {
	// The maximum number of results to return in this operation.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDeviceProfilesPaginator is a paginator for ListDeviceProfiles
type ListDeviceProfilesPaginator struct {
	options   ListDeviceProfilesPaginatorOptions
	client    ListDeviceProfilesAPIClient
	params    *ListDeviceProfilesInput
	nextToken *string
	firstPage bool
}

// NewListDeviceProfilesPaginator returns a new ListDeviceProfilesPaginator
func NewListDeviceProfilesPaginator(client ListDeviceProfilesAPIClient, params *ListDeviceProfilesInput, optFns ...func(*ListDeviceProfilesPaginatorOptions)) *ListDeviceProfilesPaginator {
	if params == nil {
		params = &ListDeviceProfilesInput{}
	}

	options := ListDeviceProfilesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDeviceProfilesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDeviceProfilesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDeviceProfiles page.
func (p *ListDeviceProfilesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDeviceProfilesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListDeviceProfiles(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDeviceProfiles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "ListDeviceProfiles",
	}
}
