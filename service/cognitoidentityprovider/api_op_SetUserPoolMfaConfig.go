// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the user pool multi-factor authentication (MFA) configuration. This action
// might generate an SMS text message. Starting June 1, 2021, US telecom carriers
// require you to register an origination phone number before you can send SMS
// messages to US phone numbers. If you use SMS text messages in Amazon Cognito,
// you must register a phone number with Amazon Pinpoint (https://console.aws.amazon.com/pinpoint/home/)
// . Amazon Cognito uses the registered number automatically. Otherwise, Amazon
// Cognito users who must receive SMS messages might not be able to sign up,
// activate their accounts, or sign in. If you have never used SMS text messages
// with Amazon Cognito or any other Amazon Web Service, Amazon Simple Notification
// Service might place your account in the SMS sandbox. In sandbox mode (https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html)
// , you can send messages only to verified phone numbers. After you test your app
// while in the sandbox environment, you can move out of the sandbox and into
// production. For more information, see SMS message settings for Amazon Cognito
// user pools (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-sms-userpool-settings.html)
// in the Amazon Cognito Developer Guide.
func (c *Client) SetUserPoolMfaConfig(ctx context.Context, params *SetUserPoolMfaConfigInput, optFns ...func(*Options)) (*SetUserPoolMfaConfigOutput, error) {
	if params == nil {
		params = &SetUserPoolMfaConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetUserPoolMfaConfig", params, optFns, c.addOperationSetUserPoolMfaConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetUserPoolMfaConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetUserPoolMfaConfigInput struct {

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// The MFA configuration. If you set the MfaConfiguration value to ‘ON’, only
	// users who have set up an MFA factor can sign in. To learn more, see Adding
	// Multi-Factor Authentication (MFA) to a user pool (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-mfa.html)
	// . Valid values include:
	//   - OFF MFA won't be used for any users.
	//   - ON MFA is required for all users to sign in.
	//   - OPTIONAL MFA will be required only for individual users who have an MFA
	//   factor activated.
	MfaConfiguration types.UserPoolMfaType

	// The SMS text message MFA configuration.
	SmsMfaConfiguration *types.SmsMfaConfigType

	// The software token MFA configuration.
	SoftwareTokenMfaConfiguration *types.SoftwareTokenMfaConfigType

	noSmithyDocumentSerde
}

type SetUserPoolMfaConfigOutput struct {

	// The MFA configuration. Valid values include:
	//   - OFF MFA won't be used for any users.
	//   - ON MFA is required for all users to sign in.
	//   - OPTIONAL MFA will be required only for individual users who have an MFA
	//   factor enabled.
	MfaConfiguration types.UserPoolMfaType

	// The SMS text message MFA configuration.
	SmsMfaConfiguration *types.SmsMfaConfigType

	// The software token MFA configuration.
	SoftwareTokenMfaConfiguration *types.SoftwareTokenMfaConfigType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetUserPoolMfaConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetUserPoolMfaConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetUserPoolMfaConfig{}, middleware.After)
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
	if err = addOpSetUserPoolMfaConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetUserPoolMfaConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetUserPoolMfaConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "SetUserPoolMfaConfig",
	}
}
