// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehendmedical

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehendmedical/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Inspects the clinical text for protected health information (PHI) entities and
// returns the entity category, location, and confidence score for each entity.
// Amazon Comprehend Medical only detects entities in English language texts.
func (c *Client) DetectPHI(ctx context.Context, params *DetectPHIInput, optFns ...func(*Options)) (*DetectPHIOutput, error) {
	if params == nil {
		params = &DetectPHIInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetectPHI", params, optFns, c.addOperationDetectPHIMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetectPHIOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectPHIInput struct {

	// A UTF-8 text string containing the clinical content being examined for PHI
	// entities. Each string must contain fewer than 20,000 bytes of characters.
	//
	// This member is required.
	Text *string

	noSmithyDocumentSerde
}

type DetectPHIOutput struct {

	// The collection of PHI entities extracted from the input text and their
	// associated information. For each entity, the response provides the entity text,
	// the entity category, where the entity text begins and ends, and the level of
	// confidence that Comprehend Medical; has in its detection.
	//
	// This member is required.
	Entities []types.Entity

	// The version of the model used to analyze the documents. The version number
	// looks like X.X.X. You can use this information to track the model used for a
	// particular batch of documents.
	//
	// This member is required.
	ModelVersion *string

	// If the result of the previous request to DetectPHI was truncated, include the
	// PaginationToken to fetch the next page of PHI entities.
	PaginationToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetectPHIMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetectPHI{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectPHI{}, middleware.After)
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
	if err = addOpDetectPHIValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetectPHI(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetectPHI(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehendmedical",
		OperationName: "DetectPHI",
	}
}
