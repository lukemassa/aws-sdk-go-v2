// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of supported languages for a specified categoryCode , issueType
// and serviceCode . The returned supported languages will include a ISO 639-1 code
// for the language , and the language display name.
//   - You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to
//     use the Amazon Web Services Support API.
//   - If you call the Amazon Web Services Support API from an account that
//     doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the
//     SubscriptionRequiredException error message appears. For information about
//     changing your support plan, see Amazon Web Services Support (http://aws.amazon.com/premiumsupport/)
//     .
func (c *Client) DescribeSupportedLanguages(ctx context.Context, params *DescribeSupportedLanguagesInput, optFns ...func(*Options)) (*DescribeSupportedLanguagesOutput, error) {
	if params == nil {
		params = &DescribeSupportedLanguagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSupportedLanguages", params, optFns, c.addOperationDescribeSupportedLanguagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSupportedLanguagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSupportedLanguagesInput struct {

	// The category of problem for the support case. You also use the DescribeServices
	// operation to get the category code for a service. Each Amazon Web Services
	// service defines its own set of category codes.
	//
	// This member is required.
	CategoryCode *string

	// The type of issue for the case. You can specify customer-service or technical .
	//
	// This member is required.
	IssueType *string

	// The code for the Amazon Web Services service. You can use the DescribeServices
	// operation to get the possible serviceCode values.
	//
	// This member is required.
	ServiceCode *string

	noSmithyDocumentSerde
}

type DescribeSupportedLanguagesOutput struct {

	// A JSON-formatted array that contains the available ISO 639-1 language codes.
	SupportedLanguages []types.SupportedLanguage

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSupportedLanguagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSupportedLanguages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSupportedLanguages{}, middleware.After)
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
	if err = addOpDescribeSupportedLanguagesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSupportedLanguages(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeSupportedLanguages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "support",
		OperationName: "DescribeSupportedLanguages",
	}
}
