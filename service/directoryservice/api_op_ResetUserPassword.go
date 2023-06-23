// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Resets the password for any user in your Managed Microsoft AD or Simple AD
// directory. You can reset the password for any user in your directory with the
// following exceptions:
//   - For Simple AD, you cannot reset the password for any user that is a member
//     of either the Domain Admins or Enterprise Admins group except for the
//     administrator user.
//   - For Managed Microsoft AD, you can only reset the password for a user that
//     is in an OU based off of the NetBIOS name that you typed when you created your
//     directory. For example, you cannot reset the password for a user in the Amazon
//     Web Services Reserved OU. For more information about the OU structure for an
//     Managed Microsoft AD directory, see What Gets Created (https://docs.aws.amazon.com/directoryservice/latest/admin-guide/ms_ad_getting_started_what_gets_created.html)
//     in the Directory Service Administration Guide.
func (c *Client) ResetUserPassword(ctx context.Context, params *ResetUserPasswordInput, optFns ...func(*Options)) (*ResetUserPasswordOutput, error) {
	if params == nil {
		params = &ResetUserPasswordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResetUserPassword", params, optFns, c.addOperationResetUserPasswordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResetUserPasswordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResetUserPasswordInput struct {

	// Identifier of the Managed Microsoft AD or Simple AD directory in which the user
	// resides.
	//
	// This member is required.
	DirectoryId *string

	// The new password that will be reset.
	//
	// This member is required.
	NewPassword *string

	// The user name of the user whose password will be reset.
	//
	// This member is required.
	UserName *string

	noSmithyDocumentSerde
}

type ResetUserPasswordOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationResetUserPasswordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpResetUserPassword{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpResetUserPassword{}, middleware.After)
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
	if err = addOpResetUserPasswordValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opResetUserPassword(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opResetUserPassword(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "ResetUserPassword",
	}
}
