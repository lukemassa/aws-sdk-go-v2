// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets aside (overrides) all approval rule requirements for a specified pull
// request.
func (c *Client) OverridePullRequestApprovalRules(ctx context.Context, params *OverridePullRequestApprovalRulesInput, optFns ...func(*Options)) (*OverridePullRequestApprovalRulesOutput, error) {
	if params == nil {
		params = &OverridePullRequestApprovalRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "OverridePullRequestApprovalRules", params, optFns, c.addOperationOverridePullRequestApprovalRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*OverridePullRequestApprovalRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type OverridePullRequestApprovalRulesInput struct {

	// Whether you want to set aside approval rule requirements for the pull request
	// (OVERRIDE) or revoke a previous override and apply approval rule requirements
	// (REVOKE). REVOKE status is not stored.
	//
	// This member is required.
	OverrideStatus types.OverrideStatus

	// The system-generated ID of the pull request for which you want to override all
	// approval rule requirements. To get this information, use GetPullRequest .
	//
	// This member is required.
	PullRequestId *string

	// The system-generated ID of the most recent revision of the pull request. You
	// cannot override approval rules for anything but the most recent revision of a
	// pull request. To get the revision ID, use GetPullRequest.
	//
	// This member is required.
	RevisionId *string

	noSmithyDocumentSerde
}

type OverridePullRequestApprovalRulesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationOverridePullRequestApprovalRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpOverridePullRequestApprovalRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpOverridePullRequestApprovalRules{}, middleware.After)
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
	if err = addOpOverridePullRequestApprovalRulesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opOverridePullRequestApprovalRules(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opOverridePullRequestApprovalRules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "OverridePullRequestApprovalRules",
	}
}
