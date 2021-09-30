// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the existing custom verification email templates for your account in the
// current Amazon Web Services Region. For more information about custom
// verification email templates, see Using Custom Verification Email Templates
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-email-verify-address-custom.html)
// in the Amazon SES Developer Guide. You can execute this operation no more than
// once per second.
func (c *Client) ListCustomVerificationEmailTemplates(ctx context.Context, params *ListCustomVerificationEmailTemplatesInput, optFns ...func(*Options)) (*ListCustomVerificationEmailTemplatesOutput, error) {
	if params == nil {
		params = &ListCustomVerificationEmailTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCustomVerificationEmailTemplates", params, optFns, c.addOperationListCustomVerificationEmailTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCustomVerificationEmailTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to list the existing custom verification email templates
// for your account.
type ListCustomVerificationEmailTemplatesInput struct {

	// A token returned from a previous call to ListCustomVerificationEmailTemplates to
	// indicate the position in the list of custom verification email templates.
	NextToken *string

	// The number of results to show in a single call to
	// ListCustomVerificationEmailTemplates. If the number of results is larger than
	// the number you specified in this parameter, then the response includes a
	// NextToken element, which you can use to obtain additional results. The value you
	// specify has to be at least 1, and can be no more than 50.
	PageSize *int32

	noSmithyDocumentSerde
}

// The following elements are returned by the service.
type ListCustomVerificationEmailTemplatesOutput struct {

	// A list of the custom verification email templates that exist in your account.
	CustomVerificationEmailTemplates []types.CustomVerificationEmailTemplateMetadata

	// A token indicating that there are additional custom verification email templates
	// available to be listed. Pass this token to a subsequent call to
	// ListCustomVerificationEmailTemplates to retrieve the next 50 custom verification
	// email templates.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCustomVerificationEmailTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListCustomVerificationEmailTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListCustomVerificationEmailTemplates{}, middleware.After)
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
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCustomVerificationEmailTemplates(options.Region), middleware.Before); err != nil {
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

// ListCustomVerificationEmailTemplatesAPIClient is a client that implements the
// ListCustomVerificationEmailTemplates operation.
type ListCustomVerificationEmailTemplatesAPIClient interface {
	ListCustomVerificationEmailTemplates(context.Context, *ListCustomVerificationEmailTemplatesInput, ...func(*Options)) (*ListCustomVerificationEmailTemplatesOutput, error)
}

var _ ListCustomVerificationEmailTemplatesAPIClient = (*Client)(nil)

// ListCustomVerificationEmailTemplatesPaginatorOptions is the paginator options
// for ListCustomVerificationEmailTemplates
type ListCustomVerificationEmailTemplatesPaginatorOptions struct {
	// The number of results to show in a single call to
	// ListCustomVerificationEmailTemplates. If the number of results is larger than
	// the number you specified in this parameter, then the response includes a
	// NextToken element, which you can use to obtain additional results. The value you
	// specify has to be at least 1, and can be no more than 50.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCustomVerificationEmailTemplatesPaginator is a paginator for
// ListCustomVerificationEmailTemplates
type ListCustomVerificationEmailTemplatesPaginator struct {
	options   ListCustomVerificationEmailTemplatesPaginatorOptions
	client    ListCustomVerificationEmailTemplatesAPIClient
	params    *ListCustomVerificationEmailTemplatesInput
	nextToken *string
	firstPage bool
}

// NewListCustomVerificationEmailTemplatesPaginator returns a new
// ListCustomVerificationEmailTemplatesPaginator
func NewListCustomVerificationEmailTemplatesPaginator(client ListCustomVerificationEmailTemplatesAPIClient, params *ListCustomVerificationEmailTemplatesInput, optFns ...func(*ListCustomVerificationEmailTemplatesPaginatorOptions)) *ListCustomVerificationEmailTemplatesPaginator {
	if params == nil {
		params = &ListCustomVerificationEmailTemplatesInput{}
	}

	options := ListCustomVerificationEmailTemplatesPaginatorOptions{}
	if params.PageSize != nil {
		options.Limit = *params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCustomVerificationEmailTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCustomVerificationEmailTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListCustomVerificationEmailTemplates page.
func (p *ListCustomVerificationEmailTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCustomVerificationEmailTemplatesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.PageSize = limit

	result, err := p.client.ListCustomVerificationEmailTemplates(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListCustomVerificationEmailTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListCustomVerificationEmailTemplates",
	}
}
