// Code generated by smithy-go-codegen DO NOT EDIT.

package resourceexplorer2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the specified view as the default for the Amazon Web Services Region in
// which you call this operation. When a user performs a Search that doesn't
// explicitly specify which view to use, then Amazon Web Services Resource Explorer
// automatically chooses this default view for searches performed in this Amazon
// Web Services Region. If an Amazon Web Services Region doesn't have a default
// view configured, then users must explicitly specify a view with every Search
// operation performed in that Region.
func (c *Client) AssociateDefaultView(ctx context.Context, params *AssociateDefaultViewInput, optFns ...func(*Options)) (*AssociateDefaultViewOutput, error) {
	if params == nil {
		params = &AssociateDefaultViewInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateDefaultView", params, optFns, c.addOperationAssociateDefaultViewMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateDefaultViewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateDefaultViewInput struct {

	// The Amazon resource name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the view to set as the default for the Amazon Web Services Region and Amazon
	// Web Services account in which you call this operation. The specified view must
	// already exist in the called Region.
	//
	// This member is required.
	ViewArn *string

	noSmithyDocumentSerde
}

type AssociateDefaultViewOutput struct {

	// The Amazon resource name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the view that the operation set as the default for queries made in the Amazon
	// Web Services Region and Amazon Web Services account in which you called this
	// operation.
	ViewArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateDefaultViewMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateDefaultView{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateDefaultView{}, middleware.After)
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
	if err = addOpAssociateDefaultViewValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateDefaultView(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateDefaultView(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "resource-explorer-2",
		OperationName: "AssociateDefaultView",
	}
}
