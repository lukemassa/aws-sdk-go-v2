// Code generated by smithy-go-codegen DO NOT EDIT.

package iotdataplane

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotdataplane/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists summary information about the retained messages stored for the account.
// This action returns only the topic names of the retained messages. It doesn't
// return any message payloads. Although this action doesn't return a message
// payload, it can still incur messaging costs. To get the message payload of a
// retained message, call GetRetainedMessage (https://docs.aws.amazon.com/iot/latest/apireference/API_iotdata_GetRetainedMessage.html)
// with the topic name of the retained message. Requires permission to access the
// ListRetainedMessages (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiotfleethubfordevicemanagement.html#awsiotfleethubfordevicemanagement-actions-as-permissions)
// action. For more information about messaging costs, see Amazon Web Services IoT
// Core pricing - Messaging (http://aws.amazon.com/iot-core/pricing/#Messaging) .
func (c *Client) ListRetainedMessages(ctx context.Context, params *ListRetainedMessagesInput, optFns ...func(*Options)) (*ListRetainedMessagesOutput, error) {
	if params == nil {
		params = &ListRetainedMessagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRetainedMessages", params, optFns, c.addOperationListRetainedMessagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRetainedMessagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRetainedMessagesInput struct {

	// The maximum number of results to return at one time.
	MaxResults int32

	// To retrieve the next set of results, the nextToken value from a previous
	// response; otherwise null to receive the first set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListRetainedMessagesOutput struct {

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// A summary list the account's retained messages. The information returned
	// doesn't include the message payloads of the retained messages.
	RetainedTopics []types.RetainedMessageSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRetainedMessagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRetainedMessages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRetainedMessages{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRetainedMessages(options.Region), middleware.Before); err != nil {
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

// ListRetainedMessagesAPIClient is a client that implements the
// ListRetainedMessages operation.
type ListRetainedMessagesAPIClient interface {
	ListRetainedMessages(context.Context, *ListRetainedMessagesInput, ...func(*Options)) (*ListRetainedMessagesOutput, error)
}

var _ ListRetainedMessagesAPIClient = (*Client)(nil)

// ListRetainedMessagesPaginatorOptions is the paginator options for
// ListRetainedMessages
type ListRetainedMessagesPaginatorOptions struct {
	// The maximum number of results to return at one time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRetainedMessagesPaginator is a paginator for ListRetainedMessages
type ListRetainedMessagesPaginator struct {
	options   ListRetainedMessagesPaginatorOptions
	client    ListRetainedMessagesAPIClient
	params    *ListRetainedMessagesInput
	nextToken *string
	firstPage bool
}

// NewListRetainedMessagesPaginator returns a new ListRetainedMessagesPaginator
func NewListRetainedMessagesPaginator(client ListRetainedMessagesAPIClient, params *ListRetainedMessagesInput, optFns ...func(*ListRetainedMessagesPaginatorOptions)) *ListRetainedMessagesPaginator {
	if params == nil {
		params = &ListRetainedMessagesInput{}
	}

	options := ListRetainedMessagesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRetainedMessagesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRetainedMessagesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListRetainedMessages page.
func (p *ListRetainedMessagesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRetainedMessagesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListRetainedMessages(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRetainedMessages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotdata",
		OperationName: "ListRetainedMessages",
	}
}
