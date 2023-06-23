// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates some of the parameters for an existing agreement. Provide the
// AgreementId and the ServerId for the agreement that you want to update, along
// with the new values for the parameters to update.
func (c *Client) UpdateAgreement(ctx context.Context, params *UpdateAgreementInput, optFns ...func(*Options)) (*UpdateAgreementOutput, error) {
	if params == nil {
		params = &UpdateAgreementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAgreement", params, optFns, c.addOperationUpdateAgreementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAgreementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAgreementInput struct {

	// A unique identifier for the agreement. This identifier is returned when you
	// create an agreement.
	//
	// This member is required.
	AgreementId *string

	// A system-assigned unique identifier for a server instance. This is the specific
	// server that the agreement uses.
	//
	// This member is required.
	ServerId *string

	// With AS2, you can send files by calling StartFileTransfer and specifying the
	// file paths in the request parameter, SendFilePaths . We use the file’s parent
	// directory (for example, for --send-file-paths /bucket/dir/file.txt , parent
	// directory is /bucket/dir/ ) to temporarily store a processed AS2 message file,
	// store the MDN when we receive them from the partner, and write a final JSON file
	// containing relevant metadata of the transmission. So, the AccessRole needs to
	// provide read and write access to the parent directory of the file location used
	// in the StartFileTransfer request. Additionally, you need to provide read and
	// write access to the parent directory of the files that you intend to send with
	// StartFileTransfer .
	AccessRole *string

	// To change the landing directory (folder) for files that are transferred,
	// provide the bucket folder that you want to use; for example,
	// /DOC-EXAMPLE-BUCKET/home/mydirectory .
	BaseDirectory *string

	// To replace the existing description, provide a short description for the
	// agreement.
	Description *string

	// A unique identifier for the AS2 local profile. To change the local profile
	// identifier, provide a new value here.
	LocalProfileId *string

	// A unique identifier for the partner profile. To change the partner profile
	// identifier, provide a new value here.
	PartnerProfileId *string

	// You can update the status for the agreement, either activating an inactive
	// agreement or the reverse.
	Status types.AgreementStatusType

	noSmithyDocumentSerde
}

type UpdateAgreementOutput struct {

	// A unique identifier for the agreement. This identifier is returned when you
	// create an agreement.
	//
	// This member is required.
	AgreementId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAgreementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateAgreement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateAgreement{}, middleware.After)
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
	if err = addOpUpdateAgreementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAgreement(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateAgreement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "UpdateAgreement",
	}
}
