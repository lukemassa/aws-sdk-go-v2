// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/acmpca/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an audit report that lists every time that your CA private key is used.
// The report is saved in the Amazon S3 bucket that you specify on input. The
// IssueCertificate (https://docs.aws.amazon.com/privateca/latest/APIReference/API_IssueCertificate.html)
// and RevokeCertificate (https://docs.aws.amazon.com/privateca/latest/APIReference/API_RevokeCertificate.html)
// actions use the private key. Both Amazon Web Services Private CA and the IAM
// principal must have permission to write to the S3 bucket that you specify. If
// the IAM principal making the call does not have permission to write to the
// bucket, then an exception is thrown. For more information, see Access policies
// for CRLs in Amazon S3 (https://docs.aws.amazon.com/privateca/latest/userguide/crl-planning.html#s3-policies)
// . Amazon Web Services Private CA assets that are stored in Amazon S3 can be
// protected with encryption. For more information, see Encrypting Your Audit
// Reports (https://docs.aws.amazon.com/privateca/latest/userguide/PcaAuditReport.html#audit-report-encryption)
// . You can generate a maximum of one report every 30 minutes.
func (c *Client) CreateCertificateAuthorityAuditReport(ctx context.Context, params *CreateCertificateAuthorityAuditReportInput, optFns ...func(*Options)) (*CreateCertificateAuthorityAuditReportOutput, error) {
	if params == nil {
		params = &CreateCertificateAuthorityAuditReportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCertificateAuthorityAuditReport", params, optFns, c.addOperationCreateCertificateAuthorityAuditReportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCertificateAuthorityAuditReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCertificateAuthorityAuditReportInput struct {

	// The format in which to create the report. This can be either JSON or CSV.
	//
	// This member is required.
	AuditReportResponseFormat types.AuditReportResponseFormat

	// The Amazon Resource Name (ARN) of the CA to be audited. This is of the form:
	// arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012
	// .
	//
	// This member is required.
	CertificateAuthorityArn *string

	// The name of the S3 bucket that will contain the audit report.
	//
	// This member is required.
	S3BucketName *string

	noSmithyDocumentSerde
}

type CreateCertificateAuthorityAuditReportOutput struct {

	// An alphanumeric string that contains a report identifier.
	AuditReportId *string

	// The key that uniquely identifies the report file in your S3 bucket.
	S3Key *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateCertificateAuthorityAuditReportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCertificateAuthorityAuditReport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCertificateAuthorityAuditReport{}, middleware.After)
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
	if err = addOpCreateCertificateAuthorityAuditReportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCertificateAuthorityAuditReport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCertificateAuthorityAuditReport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "CreateCertificateAuthorityAuditReport",
	}
}
