// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html)
// . With the latest version, AWS WAF has a single set of endpoints for regional
// and global use. Returns the status of a ChangeToken that you got by calling
// GetChangeToken . ChangeTokenStatus is one of the following values:
//   - PROVISIONED : You requested the change token by calling GetChangeToken , but
//     you haven't used it yet in a call to create, update, or delete an AWS WAF
//     object.
//   - PENDING : AWS WAF is propagating the create, update, or delete request to
//     all AWS WAF servers.
//   - INSYNC : Propagation is complete.
func (c *Client) GetChangeTokenStatus(ctx context.Context, params *GetChangeTokenStatusInput, optFns ...func(*Options)) (*GetChangeTokenStatusOutput, error) {
	if params == nil {
		params = &GetChangeTokenStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetChangeTokenStatus", params, optFns, c.addOperationGetChangeTokenStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetChangeTokenStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetChangeTokenStatusInput struct {

	// The change token for which you want to get the status. This change token was
	// previously returned in the GetChangeToken response.
	//
	// This member is required.
	ChangeToken *string

	noSmithyDocumentSerde
}

type GetChangeTokenStatusOutput struct {

	// The status of the change token.
	ChangeTokenStatus types.ChangeTokenStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetChangeTokenStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetChangeTokenStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetChangeTokenStatus{}, middleware.After)
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
	if err = addOpGetChangeTokenStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetChangeTokenStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetChangeTokenStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "GetChangeTokenStatus",
	}
}
