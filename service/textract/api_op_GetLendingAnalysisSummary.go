// Code generated by smithy-go-codegen DO NOT EDIT.

package textract

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets summarized results for the StartLendingAnalysis operation, which analyzes
// text in a lending document. The returned summary consists of information about
// documents grouped together by a common document type. Information like detected
// signatures, page numbers, and split documents is returned with respect to the
// type of grouped document. You start asynchronous text analysis by calling
// StartLendingAnalysis , which returns a job identifier ( JobId ). When the text
// analysis operation finishes, Amazon Textract publishes a completion status to
// the Amazon Simple Notification Service (Amazon SNS) topic that's registered in
// the initial call to StartLendingAnalysis . To get the results of the text
// analysis operation, first check that the status value published to the Amazon
// SNS topic is SUCCEEDED. If so, call GetLendingAnalysisSummary , and pass the job
// identifier ( JobId ) from the initial call to StartLendingAnalysis .
func (c *Client) GetLendingAnalysisSummary(ctx context.Context, params *GetLendingAnalysisSummaryInput, optFns ...func(*Options)) (*GetLendingAnalysisSummaryOutput, error) {
	if params == nil {
		params = &GetLendingAnalysisSummaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLendingAnalysisSummary", params, optFns, c.addOperationGetLendingAnalysisSummaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLendingAnalysisSummaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLendingAnalysisSummaryInput struct {

	// A unique identifier for the lending or text-detection job. The JobId is
	// returned from StartLendingAnalysis. A JobId value is only valid for 7 days.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

type GetLendingAnalysisSummaryOutput struct {

	// The current model version of the Analyze Lending API.
	AnalyzeLendingModelVersion *string

	// Information about the input document.
	DocumentMetadata *types.DocumentMetadata

	// The current status of the lending analysis job.
	JobStatus types.JobStatus

	// Returns if the lending analysis could not be completed. Contains explanation
	// for what error occurred.
	StatusMessage *string

	// Contains summary information for documents grouped by type.
	Summary *types.LendingSummary

	// A list of warnings that occurred during the lending analysis operation.
	Warnings []types.Warning

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLendingAnalysisSummaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetLendingAnalysisSummary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetLendingAnalysisSummary{}, middleware.After)
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
	if err = addOpGetLendingAnalysisSummaryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLendingAnalysisSummary(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLendingAnalysisSummary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "textract",
		OperationName: "GetLendingAnalysisSummary",
	}
}
