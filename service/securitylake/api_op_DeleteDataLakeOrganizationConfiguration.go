// Code generated by smithy-go-codegen DO NOT EDIT.

package securitylake

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securitylake/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes automatic the enablement of configuration settings for new member
// accounts (but retains the settings for the delegated administrator) from Amazon
// Security Lake. You must run this API using the credentials of the delegated
// administrator. When you run this API, new member accounts that are added after
// the organization enables Security Lake won't contribute to the data lake.
func (c *Client) DeleteDataLakeOrganizationConfiguration(ctx context.Context, params *DeleteDataLakeOrganizationConfigurationInput, optFns ...func(*Options)) (*DeleteDataLakeOrganizationConfigurationOutput, error) {
	if params == nil {
		params = &DeleteDataLakeOrganizationConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteDataLakeOrganizationConfiguration", params, optFns, c.addOperationDeleteDataLakeOrganizationConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteDataLakeOrganizationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDataLakeOrganizationConfigurationInput struct {

	// Removes the automatic enablement of configuration settings for new member
	// accounts in Security Lake.
	//
	// This member is required.
	AutoEnableNewAccount []types.DataLakeAutoEnableNewAccountConfiguration

	noSmithyDocumentSerde
}

type DeleteDataLakeOrganizationConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteDataLakeOrganizationConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteDataLakeOrganizationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteDataLakeOrganizationConfiguration{}, middleware.After)
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
	if err = addOpDeleteDataLakeOrganizationConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDataLakeOrganizationConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteDataLakeOrganizationConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securitylake",
		OperationName: "DeleteDataLakeOrganizationConfiguration",
	}
}
