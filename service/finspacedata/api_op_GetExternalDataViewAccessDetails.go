// Code generated by smithy-go-codegen DO NOT EDIT.

package finspacedata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/finspacedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the credentials to access the external Dataview from an S3 location. To
// call this API:
//   - You must retrieve the programmatic credentials.
//   - You must be a member of a FinSpace user group, where the dataset that you
//     want to access has Read Dataset Data permissions.
func (c *Client) GetExternalDataViewAccessDetails(ctx context.Context, params *GetExternalDataViewAccessDetailsInput, optFns ...func(*Options)) (*GetExternalDataViewAccessDetailsOutput, error) {
	if params == nil {
		params = &GetExternalDataViewAccessDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetExternalDataViewAccessDetails", params, optFns, c.addOperationGetExternalDataViewAccessDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetExternalDataViewAccessDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetExternalDataViewAccessDetailsInput struct {

	// The unique identifier for the Dataview that you want to access.
	//
	// This member is required.
	DataViewId *string

	// The unique identifier for the Dataset.
	//
	// This member is required.
	DatasetId *string

	noSmithyDocumentSerde
}

type GetExternalDataViewAccessDetailsOutput struct {

	// The credentials required to access the external Dataview from the S3 location.
	Credentials *types.AwsCredentials

	// The location where the external Dataview is stored.
	S3Location *types.S3Location

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetExternalDataViewAccessDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetExternalDataViewAccessDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetExternalDataViewAccessDetails{}, middleware.After)
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetExternalDataViewAccessDetailsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetExternalDataViewAccessDetails(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetExternalDataViewAccessDetails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "finspace-api",
		OperationName: "GetExternalDataViewAccessDetails",
	}
}
