// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the custom MAIL FROM attributes for a list of identities (email
// addresses : domains). This operation is throttled at one request per second and
// can only get custom MAIL FROM attributes for up to 100 identities at a time.
func (c *Client) GetIdentityMailFromDomainAttributes(ctx context.Context, params *GetIdentityMailFromDomainAttributesInput, optFns ...func(*Options)) (*GetIdentityMailFromDomainAttributesOutput, error) {
	if params == nil {
		params = &GetIdentityMailFromDomainAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIdentityMailFromDomainAttributes", params, optFns, c.addOperationGetIdentityMailFromDomainAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIdentityMailFromDomainAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to return the Amazon SES custom MAIL FROM attributes for a
// list of identities. For information about using a custom MAIL FROM domain, see
// the Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/mail-from.html)
// .
type GetIdentityMailFromDomainAttributesInput struct {

	// A list of one or more identities.
	//
	// This member is required.
	Identities []string

	noSmithyDocumentSerde
}

// Represents the custom MAIL FROM attributes for a list of identities.
type GetIdentityMailFromDomainAttributesOutput struct {

	// A map of identities to custom MAIL FROM attributes.
	//
	// This member is required.
	MailFromDomainAttributes map[string]types.IdentityMailFromDomainAttributes

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetIdentityMailFromDomainAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetIdentityMailFromDomainAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetIdentityMailFromDomainAttributes{}, middleware.After)
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
	if err = addOpGetIdentityMailFromDomainAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIdentityMailFromDomainAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetIdentityMailFromDomainAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetIdentityMailFromDomainAttributes",
	}
}
