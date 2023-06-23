// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a presigned Amazon S3 URL that can be used to upload a file as manual
// evidence. For instructions on how to use this operation, see Upload a file from
// your browser  (https://docs.aws.amazon.com/audit-manager/latest/userguide/upload-evidence.html#how-to-upload-manual-evidence-files)
// in the Audit Manager User Guide. The following restrictions apply to this
// operation:
//   - Maximum size of an individual evidence file: 100 MB
//   - Number of daily manual evidence uploads per control: 100
//   - Supported file formats: See Supported file types for manual evidence (https://docs.aws.amazon.com/audit-manager/latest/userguide/upload-evidence.html#supported-manual-evidence-files)
//     in the Audit Manager User Guide
//
// For more information about Audit Manager service restrictions, see Quotas and
// restrictions for Audit Manager (https://docs.aws.amazon.com/audit-manager/latest/userguide/service-quotas.html)
// .
func (c *Client) GetEvidenceFileUploadUrl(ctx context.Context, params *GetEvidenceFileUploadUrlInput, optFns ...func(*Options)) (*GetEvidenceFileUploadUrlOutput, error) {
	if params == nil {
		params = &GetEvidenceFileUploadUrlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEvidenceFileUploadUrl", params, optFns, c.addOperationGetEvidenceFileUploadUrlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEvidenceFileUploadUrlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEvidenceFileUploadUrlInput struct {

	// The file that you want to upload. For a list of supported file formats, see
	// Supported file types for manual evidence (https://docs.aws.amazon.com/audit-manager/latest/userguide/upload-evidence.html#supported-manual-evidence-files)
	// in the Audit Manager User Guide.
	//
	// This member is required.
	FileName *string

	noSmithyDocumentSerde
}

type GetEvidenceFileUploadUrlOutput struct {

	// The name of the uploaded manual evidence file that the presigned URL was
	// generated for.
	EvidenceFileName *string

	// The presigned URL that was generated.
	UploadUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEvidenceFileUploadUrlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetEvidenceFileUploadUrl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEvidenceFileUploadUrl{}, middleware.After)
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
	if err = addOpGetEvidenceFileUploadUrlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEvidenceFileUploadUrl(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEvidenceFileUploadUrl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "GetEvidenceFileUploadUrl",
	}
}
