// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Extends the expiration date for license consumption.
func (c *Client) ExtendLicenseConsumption(ctx context.Context, params *ExtendLicenseConsumptionInput, optFns ...func(*Options)) (*ExtendLicenseConsumptionOutput, error) {
	if params == nil {
		params = &ExtendLicenseConsumptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExtendLicenseConsumption", params, optFns, c.addOperationExtendLicenseConsumptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExtendLicenseConsumptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExtendLicenseConsumptionInput struct {

	// License consumption token.
	//
	// This member is required.
	LicenseConsumptionToken *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request. Provides an error response if you do not have the
	// required permissions.
	DryRun bool

	noSmithyDocumentSerde
}

type ExtendLicenseConsumptionOutput struct {

	// Date and time at which the license consumption expires.
	Expiration *string

	// License consumption token.
	LicenseConsumptionToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExtendLicenseConsumptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpExtendLicenseConsumption{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpExtendLicenseConsumption{}, middleware.After)
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
	if err = addOpExtendLicenseConsumptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExtendLicenseConsumption(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opExtendLicenseConsumption(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager",
		OperationName: "ExtendLicenseConsumption",
	}
}
