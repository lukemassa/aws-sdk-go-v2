// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified product.
func (c *Client) UpdateProduct(ctx context.Context, params *UpdateProductInput, optFns ...func(*Options)) (*UpdateProductOutput, error) {
	if params == nil {
		params = &UpdateProductInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateProduct", params, optFns, c.addOperationUpdateProductMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateProductOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateProductInput struct {

	// The product identifier.
	//
	// This member is required.
	Id *string

	// The language code.
	//   - jp - Japanese
	//   - zh - Chinese
	AcceptLanguage *string

	// The tags to add to the product.
	AddTags []types.Tag

	// The updated description of the product.
	Description *string

	// The updated distributor of the product.
	Distributor *string

	// The updated product name.
	Name *string

	// The updated owner of the product.
	Owner *string

	// The tags to remove from the product.
	RemoveTags []string

	// Specifies connection details for the updated product and syncs the product to
	// the connection source artifact. This automatically manages the product's
	// artifacts based on changes to the source. The SourceConnection parameter
	// consists of the following sub-fields.
	//   - Type
	//   - ConnectionParamters
	SourceConnection *types.SourceConnection

	// The updated support description for the product.
	SupportDescription *string

	// The updated support email for the product.
	SupportEmail *string

	// The updated support URL for the product.
	SupportUrl *string

	noSmithyDocumentSerde
}

type UpdateProductOutput struct {

	// Information about the product view.
	ProductViewDetail *types.ProductViewDetail

	// Information about the tags associated with the product.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateProductMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateProduct{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateProduct{}, middleware.After)
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
	if err = addOpUpdateProductValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateProduct(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateProduct(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "UpdateProduct",
	}
}
