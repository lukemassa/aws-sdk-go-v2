// Code generated by smithy-go-codegen DO NOT EDIT.

package tnb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a network instance. A network instance is a single network created in
// Amazon Web Services TNB that can be deployed and on which life-cycle operations
// (like terminate, update, and delete) can be performed. Creating a network
// instance is the third step after creating a network package. For more
// information about network instances, Network instances
// (https://docs.aws.amazon.com/tnb/latest/ug/network-instances.html) in the Amazon
// Web Services Telco Network Builder User Guide. Once you create a network
// instance, you can instantiate it. To instantiate a network, see
// InstantiateSolNetworkInstance
// (https://docs.aws.amazon.com/TNB/latest/APIReference/API_InstantiateSolNetworkInstance.html).
func (c *Client) CreateSolNetworkInstance(ctx context.Context, params *CreateSolNetworkInstanceInput, optFns ...func(*Options)) (*CreateSolNetworkInstanceOutput, error) {
	if params == nil {
		params = &CreateSolNetworkInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSolNetworkInstance", params, optFns, c.addOperationCreateSolNetworkInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSolNetworkInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSolNetworkInstanceInput struct {

	// Network instance name.
	//
	// This member is required.
	NsName *string

	// ID for network service descriptor.
	//
	// This member is required.
	NsdInfoId *string

	// Network instance description.
	NsDescription *string

	// A tag is a label that you assign to an Amazon Web Services resource. Each tag
	// consists of a key and an optional value. You can use tags to search and filter
	// your resources or track your Amazon Web Services costs.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateSolNetworkInstanceOutput struct {

	// Network instance ARN.
	//
	// This member is required.
	Arn *string

	// Network instance ID.
	//
	// This member is required.
	Id *string

	// Network instance name.
	//
	// This member is required.
	NsInstanceName *string

	// Network service descriptor ID.
	//
	// This member is required.
	NsdInfoId *string

	// A tag is a label that you assign to an Amazon Web Services resource. Each tag
	// consists of a key and an optional value. You can use tags to search and filter
	// your resources or track your Amazon Web Services costs.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSolNetworkInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSolNetworkInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSolNetworkInstance{}, middleware.After)
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
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateSolNetworkInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSolNetworkInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSolNetworkInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "tnb",
		OperationName: "CreateSolNetworkInstance",
	}
}
