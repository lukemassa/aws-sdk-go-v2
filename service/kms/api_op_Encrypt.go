// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Encrypts plaintext of up to 4,096 bytes using a KMS key. You can use a
// symmetric or asymmetric KMS key with a KeyUsage of ENCRYPT_DECRYPT . You can use
// this operation to encrypt small amounts of arbitrary data, such as a personal
// identifier or database password, or other sensitive information. You don't need
// to use the Encrypt operation to encrypt a data key. The GenerateDataKey and
// GenerateDataKeyPair operations return a plaintext data key and an encrypted copy
// of that data key. If you use a symmetric encryption KMS key, you can use an
// encryption context to add additional security to your encryption operation. If
// you specify an EncryptionContext when encrypting data, you must specify the
// same encryption context (a case-sensitive exact match) when decrypting the data.
// Otherwise, the request to decrypt fails with an InvalidCiphertextException . For
// more information, see Encryption Context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
// in the Key Management Service Developer Guide. If you specify an asymmetric KMS
// key, you must also specify the encryption algorithm. The algorithm must be
// compatible with the KMS key spec. When you use an asymmetric KMS key to encrypt
// or reencrypt data, be sure to record the KMS key and encryption algorithm that
// you choose. You will be required to provide the same KMS key and encryption
// algorithm when you decrypt the data. If the KMS key and algorithm do not match
// the values used to encrypt the data, the decrypt operation fails. You are not
// required to supply the key ID and encryption algorithm when you decrypt with
// symmetric encryption KMS keys because KMS stores this information in the
// ciphertext blob. KMS cannot store metadata in ciphertext generated with
// asymmetric keys. The standard format for asymmetric key ciphertext does not
// include configurable fields. The maximum size of the data that you can encrypt
// varies with the type of KMS key and the encryption algorithm that you choose.
//   - Symmetric encryption KMS keys
//   - SYMMETRIC_DEFAULT : 4096 bytes
//   - RSA_2048
//   - RSAES_OAEP_SHA_1 : 214 bytes
//   - RSAES_OAEP_SHA_256 : 190 bytes
//   - RSA_3072
//   - RSAES_OAEP_SHA_1 : 342 bytes
//   - RSAES_OAEP_SHA_256 : 318 bytes
//   - RSA_4096
//   - RSAES_OAEP_SHA_1 : 470 bytes
//   - RSAES_OAEP_SHA_256 : 446 bytes
//   - SM2PKE : 1024 bytes (China Regions only)
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see Key states of KMS keys (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the Key Management Service Developer Guide. Cross-account use: Yes. To
// perform this operation with a KMS key in a different Amazon Web Services
// account, specify the key ARN or alias ARN in the value of the KeyId parameter.
// Required permissions: kms:Encrypt (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
// (key policy) Related operations:
//   - Decrypt
//   - GenerateDataKey
//   - GenerateDataKeyPair
func (c *Client) Encrypt(ctx context.Context, params *EncryptInput, optFns ...func(*Options)) (*EncryptOutput, error) {
	if params == nil {
		params = &EncryptInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "Encrypt", params, optFns, c.addOperationEncryptMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EncryptOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EncryptInput struct {

	// Identifies the KMS key to use in the encryption operation. The KMS key must
	// have a KeyUsage of ENCRYPT_DECRYPT . To find the KeyUsage of a KMS key, use the
	// DescribeKey operation. To specify a KMS key, use its key ID, key ARN, alias
	// name, or alias ARN. When using an alias name, prefix it with "alias/" . To
	// specify a KMS key in a different Amazon Web Services account, you must use the
	// key ARN or alias ARN. For example:
	//   - Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//   - Key ARN:
	//   arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//   - Alias name: alias/ExampleAlias
	//   - Alias ARN: arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	// To get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey . To
	// get the alias name and alias ARN, use ListAliases .
	//
	// This member is required.
	KeyId *string

	// Data to be encrypted.
	//
	// This member is required.
	Plaintext []byte

	// Specifies the encryption algorithm that KMS will use to encrypt the plaintext
	// message. The algorithm must be compatible with the KMS key that you specify.
	// This parameter is required only for asymmetric KMS keys. The default value,
	// SYMMETRIC_DEFAULT , is the algorithm used for symmetric encryption KMS keys. If
	// you are using an asymmetric KMS key, we recommend RSAES_OAEP_SHA_256. The SM2PKE
	// algorithm is only available in China Regions.
	EncryptionAlgorithm types.EncryptionAlgorithmSpec

	// Specifies the encryption context that will be used to encrypt the data. An
	// encryption context is valid only for cryptographic operations (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#cryptographic-operations)
	// with a symmetric encryption KMS key. The standard asymmetric encryption
	// algorithms and HMAC algorithms that KMS uses do not support an encryption
	// context. Do not include confidential or sensitive information in this field.
	// This field may be displayed in plaintext in CloudTrail logs and other output. An
	// encryption context is a collection of non-secret key-value pairs that represent
	// additional authenticated data. When you use an encryption context to encrypt
	// data, you must specify the same (an exact case-sensitive match) encryption
	// context to decrypt the data. An encryption context is supported only on
	// operations with symmetric encryption KMS keys. On operations with symmetric
	// encryption KMS keys, an encryption context is optional, but it is strongly
	// recommended. For more information, see Encryption context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
	// in the Key Management Service Developer Guide.
	EncryptionContext map[string]string

	// A list of grant tokens. Use a grant token when your permission to call this
	// operation comes from a new grant that has not yet achieved eventual consistency.
	// For more information, see Grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant_token)
	// and Using a grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#using-grant-token)
	// in the Key Management Service Developer Guide.
	GrantTokens []string

	noSmithyDocumentSerde
}

type EncryptOutput struct {

	// The encrypted plaintext. When you use the HTTP API or the Amazon Web Services
	// CLI, the value is Base64-encoded. Otherwise, it is not Base64-encoded.
	CiphertextBlob []byte

	// The encryption algorithm that was used to encrypt the plaintext.
	EncryptionAlgorithm types.EncryptionAlgorithmSpec

	// The Amazon Resource Name ( key ARN (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN)
	// ) of the KMS key that was used to encrypt the plaintext.
	KeyId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEncryptMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpEncrypt{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpEncrypt{}, middleware.After)
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
	if err = addOpEncryptValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEncrypt(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEncrypt(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "Encrypt",
	}
}
