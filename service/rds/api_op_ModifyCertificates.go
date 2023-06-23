// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Override the system-default Secure Sockets Layer/Transport Layer Security
// (SSL/TLS) certificate for Amazon RDS for new DB instances, or remove the
// override. By using this operation, you can specify an RDS-approved SSL/TLS
// certificate for new DB instances that is different from the default certificate
// provided by RDS. You can also use this operation to remove the override, so that
// new DB instances use the default certificate provided by RDS. You might need to
// override the default certificate in the following situations:
//   - You already migrated your applications to support the latest certificate
//     authority (CA) certificate, but the new CA certificate is not yet the RDS
//     default CA certificate for the specified Amazon Web Services Region.
//   - RDS has already moved to a new default CA certificate for the specified
//     Amazon Web Services Region, but you are still in the process of supporting the
//     new CA certificate. In this case, you temporarily need additional time to finish
//     your application changes.
//
// For more information about rotating your SSL/TLS certificate for RDS DB
// engines, see Rotating Your SSL/TLS Certificate (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL-certificate-rotation.html)
// in the Amazon RDS User Guide. For more information about rotating your SSL/TLS
// certificate for Aurora DB engines, see Rotating Your SSL/TLS Certificate (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL-certificate-rotation.html)
// in the Amazon Aurora User Guide.
func (c *Client) ModifyCertificates(ctx context.Context, params *ModifyCertificatesInput, optFns ...func(*Options)) (*ModifyCertificatesOutput, error) {
	if params == nil {
		params = &ModifyCertificatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyCertificates", params, optFns, c.addOperationModifyCertificatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyCertificatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyCertificatesInput struct {

	// The new default certificate identifier to override the current one with. To
	// determine the valid values, use the describe-certificates CLI command or the
	// DescribeCertificates API operation.
	CertificateIdentifier *string

	// A value that indicates whether to remove the override for the default
	// certificate. If the override is removed, the default certificate is the system
	// default.
	RemoveCustomerOverride *bool

	noSmithyDocumentSerde
}

type ModifyCertificatesOutput struct {

	// A CA certificate for an Amazon Web Services account. For more information, see
	// Using SSL/TLS to encrypt a connection to a DB instance (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/UsingWithRDS.SSL.html)
	// in the Amazon RDS User Guide and Using SSL/TLS to encrypt a connection to a DB
	// cluster (https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/UsingWithRDS.SSL.html)
	// in the Amazon Aurora User Guide.
	Certificate *types.Certificate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyCertificatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyCertificates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyCertificates{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyCertificates(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyCertificates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyCertificates",
	}
}
