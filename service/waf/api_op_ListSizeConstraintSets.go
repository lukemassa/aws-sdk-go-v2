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
// and global use. Returns an array of SizeConstraintSetSummary objects.
func (c *Client) ListSizeConstraintSets(ctx context.Context, params *ListSizeConstraintSetsInput, optFns ...func(*Options)) (*ListSizeConstraintSetsOutput, error) {
	if params == nil {
		params = &ListSizeConstraintSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSizeConstraintSets", params, optFns, c.addOperationListSizeConstraintSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSizeConstraintSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSizeConstraintSetsInput struct {

	// Specifies the number of SizeConstraintSet objects that you want AWS WAF to
	// return for this request. If you have more SizeConstraintSets objects than the
	// number you specify for Limit , the response includes a NextMarker value that
	// you can use to get another batch of SizeConstraintSet objects.
	Limit int32

	// If you specify a value for Limit and you have more SizeConstraintSets than the
	// value of Limit , AWS WAF returns a NextMarker value in the response that allows
	// you to list another group of SizeConstraintSets . For the second and subsequent
	// ListSizeConstraintSets requests, specify the value of NextMarker from the
	// previous response to get information about another batch of SizeConstraintSets .
	NextMarker *string

	noSmithyDocumentSerde
}

type ListSizeConstraintSetsOutput struct {

	// If you have more SizeConstraintSet objects than the number that you specified
	// for Limit in the request, the response includes a NextMarker value. To list
	// more SizeConstraintSet objects, submit another ListSizeConstraintSets request,
	// and specify the NextMarker value from the response in the NextMarker value in
	// the next request.
	NextMarker *string

	// An array of SizeConstraintSetSummary objects.
	SizeConstraintSets []types.SizeConstraintSetSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSizeConstraintSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSizeConstraintSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSizeConstraintSets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSizeConstraintSets(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListSizeConstraintSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "ListSizeConstraintSets",
	}
}
