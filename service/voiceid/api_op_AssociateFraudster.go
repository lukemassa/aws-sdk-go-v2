// Code generated by smithy-go-codegen DO NOT EDIT.

package voiceid

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/voiceid/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates the fraudsters with the watchlist specified in the same domain.
func (c *Client) AssociateFraudster(ctx context.Context, params *AssociateFraudsterInput, optFns ...func(*Options)) (*AssociateFraudsterOutput, error) {
	if params == nil {
		params = &AssociateFraudsterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateFraudster", params, optFns, c.addOperationAssociateFraudsterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateFraudsterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateFraudsterInput struct {

	// The identifier of the domain that contains the fraudster.
	//
	// This member is required.
	DomainId *string

	// The identifier of the fraudster to be associated with the watchlist.
	//
	// This member is required.
	FraudsterId *string

	// The identifier of the watchlist you want to associate with the fraudster.
	//
	// This member is required.
	WatchlistId *string

	noSmithyDocumentSerde
}

type AssociateFraudsterOutput struct {

	// Contains all the information about a fraudster.
	Fraudster *types.Fraudster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateFraudsterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpAssociateFraudster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpAssociateFraudster{}, middleware.After)
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
	if err = addOpAssociateFraudsterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateFraudster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateFraudster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "voiceid",
		OperationName: "AssociateFraudster",
	}
}
