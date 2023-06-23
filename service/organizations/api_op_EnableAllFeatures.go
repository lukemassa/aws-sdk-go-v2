// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables all features in an organization. This enables the use of organization
// policies that can restrict the services and actions that can be called in each
// account. Until you enable all features, you have access only to consolidated
// billing, and you can't use any of the advanced account administration features
// that Organizations supports. For more information, see Enabling All Features in
// Your Organization (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html)
// in the Organizations User Guide. This operation is required only for
// organizations that were created explicitly with only the consolidated billing
// features enabled. Calling this operation sends a handshake to every invited
// account in the organization. The feature set change can be finalized and the
// additional features enabled only after all administrators in the invited
// accounts approve the change by accepting the handshake. After you enable all
// features, you can separately enable or disable individual policy types in a root
// using EnablePolicyType and DisablePolicyType . To see the status of policy types
// in a root, use ListRoots . After all invited member accounts accept the
// handshake, you finalize the feature set change by accepting the handshake that
// contains "Action": "ENABLE_ALL_FEATURES" . This completes the change. After you
// enable all features in your organization, the management account in the
// organization can apply policies on all member accounts. These policies can
// restrict what users and even administrators in those accounts can do. The
// management account can apply policies that prevent accounts from leaving the
// organization. Ensure that your account administrators are aware of this. This
// operation can be called only from the organization's management account.
func (c *Client) EnableAllFeatures(ctx context.Context, params *EnableAllFeaturesInput, optFns ...func(*Options)) (*EnableAllFeaturesOutput, error) {
	if params == nil {
		params = &EnableAllFeaturesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableAllFeatures", params, optFns, c.addOperationEnableAllFeaturesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableAllFeaturesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableAllFeaturesInput struct {
	noSmithyDocumentSerde
}

type EnableAllFeaturesOutput struct {

	// A structure that contains details about the handshake created to support this
	// request to enable all features in the organization.
	Handshake *types.Handshake

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableAllFeaturesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpEnableAllFeatures{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableAllFeatures{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableAllFeatures(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableAllFeatures(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "EnableAllFeatures",
	}
}
