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

// Decrypts ciphertext and then reencrypts it entirely within KMS. You can use
// this operation to change the KMS key under which data is encrypted, such as when
// you manually rotate (https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html#rotate-keys-manually)
// a KMS key or change the KMS key that protects a ciphertext. You can also use it
// to reencrypt ciphertext under the same KMS key, such as to change the
// encryption context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
// of a ciphertext. The ReEncrypt operation can decrypt ciphertext that was
// encrypted by using a KMS key in an KMS operation, such as Encrypt or
// GenerateDataKey . It can also decrypt ciphertext that was encrypted by using the
// public key of an asymmetric KMS key (https://docs.aws.amazon.com/kms/latest/developerguide/symm-asymm-concepts.html#asymmetric-cmks)
// outside of KMS. However, it cannot decrypt ciphertext produced by other
// libraries, such as the Amazon Web Services Encryption SDK (https://docs.aws.amazon.com/encryption-sdk/latest/developer-guide/)
// or Amazon S3 client-side encryption (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingClientSideEncryption.html)
// . These libraries return a ciphertext format that is incompatible with KMS. When
// you use the ReEncrypt operation, you need to provide information for the
// decrypt operation and the subsequent encrypt operation.
//   - If your ciphertext was encrypted under an asymmetric KMS key, you must use
//     the SourceKeyId parameter to identify the KMS key that encrypted the
//     ciphertext. You must also supply the encryption algorithm that was used. This
//     information is required to decrypt the data.
//   - If your ciphertext was encrypted under a symmetric encryption KMS key, the
//     SourceKeyId parameter is optional. KMS can get this information from metadata
//     that it adds to the symmetric ciphertext blob. This feature adds durability to
//     your implementation by ensuring that authorized users can decrypt ciphertext
//     decades after it was encrypted, even if they've lost track of the key ID.
//     However, specifying the source KMS key is always recommended as a best practice.
//     When you use the SourceKeyId parameter to specify a KMS key, KMS uses only the
//     KMS key you specify. If the ciphertext was encrypted under a different KMS key,
//     the ReEncrypt operation fails. This practice ensures that you use the KMS key
//     that you intend.
//   - To reencrypt the data, you must use the DestinationKeyId parameter to
//     specify the KMS key that re-encrypts the data after it is decrypted. If the
//     destination KMS key is an asymmetric KMS key, you must also provide the
//     encryption algorithm. The algorithm that you choose must be compatible with the
//     KMS key. When you use an asymmetric KMS key to encrypt or reencrypt data, be
//     sure to record the KMS key and encryption algorithm that you choose. You will be
//     required to provide the same KMS key and encryption algorithm when you decrypt
//     the data. If the KMS key and algorithm do not match the values used to encrypt
//     the data, the decrypt operation fails. You are not required to supply the key ID
//     and encryption algorithm when you decrypt with symmetric encryption KMS keys
//     because KMS stores this information in the ciphertext blob. KMS cannot store
//     metadata in ciphertext generated with asymmetric keys. The standard format for
//     asymmetric key ciphertext does not include configurable fields.
//
// The KMS key that you use for this operation must be in a compatible key state.
// For details, see Key states of KMS keys (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html)
// in the Key Management Service Developer Guide. Cross-account use: Yes. The
// source KMS key and destination KMS key can be in different Amazon Web Services
// accounts. Either or both KMS keys can be in a different account than the caller.
// To specify a KMS key in a different account, you must use its key ARN or alias
// ARN. Required permissions:
//   - kms:ReEncryptFrom (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
//     permission on the source KMS key (key policy)
//   - kms:ReEncryptTo (https://docs.aws.amazon.com/kms/latest/developerguide/kms-api-permissions-reference.html)
//     permission on the destination KMS key (key policy)
//
// To permit reencryption from or to a KMS key, include the "kms:ReEncrypt*"
// permission in your key policy (https://docs.aws.amazon.com/kms/latest/developerguide/key-policies.html)
// . This permission is automatically included in the key policy when you use the
// console to create a KMS key. But you must include it manually when you create a
// KMS key programmatically or when you use the PutKeyPolicy operation to set a
// key policy. Related operations:
//   - Decrypt
//   - Encrypt
//   - GenerateDataKey
//   - GenerateDataKeyPair
func (c *Client) ReEncrypt(ctx context.Context, params *ReEncryptInput, optFns ...func(*Options)) (*ReEncryptOutput, error) {
	if params == nil {
		params = &ReEncryptInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReEncrypt", params, optFns, c.addOperationReEncryptMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReEncryptOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReEncryptInput struct {

	// Ciphertext of the data to reencrypt.
	//
	// This member is required.
	CiphertextBlob []byte

	// A unique identifier for the KMS key that is used to reencrypt the data. Specify
	// a symmetric encryption KMS key or an asymmetric KMS key with a KeyUsage value
	// of ENCRYPT_DECRYPT . To find the KeyUsage value of a KMS key, use the
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
	DestinationKeyId *string

	// Specifies the encryption algorithm that KMS will use to reecrypt the data after
	// it has decrypted it. The default value, SYMMETRIC_DEFAULT , represents the
	// encryption algorithm used for symmetric encryption KMS keys. This parameter is
	// required only when the destination KMS key is an asymmetric KMS key.
	DestinationEncryptionAlgorithm types.EncryptionAlgorithmSpec

	// Specifies that encryption context to use when the reencrypting the data. Do not
	// include confidential or sensitive information in this field. This field may be
	// displayed in plaintext in CloudTrail logs and other output. A destination
	// encryption context is valid only when the destination KMS key is a symmetric
	// encryption KMS key. The standard ciphertext format for asymmetric KMS keys does
	// not include fields for metadata. An encryption context is a collection of
	// non-secret key-value pairs that represent additional authenticated data. When
	// you use an encryption context to encrypt data, you must specify the same (an
	// exact case-sensitive match) encryption context to decrypt the data. An
	// encryption context is supported only on operations with symmetric encryption KMS
	// keys. On operations with symmetric encryption KMS keys, an encryption context is
	// optional, but it is strongly recommended. For more information, see Encryption
	// context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
	// in the Key Management Service Developer Guide.
	DestinationEncryptionContext map[string]string

	// A list of grant tokens. Use a grant token when your permission to call this
	// operation comes from a new grant that has not yet achieved eventual consistency.
	// For more information, see Grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grants.html#grant_token)
	// and Using a grant token (https://docs.aws.amazon.com/kms/latest/developerguide/grant-manage.html#using-grant-token)
	// in the Key Management Service Developer Guide.
	GrantTokens []string

	// Specifies the encryption algorithm that KMS will use to decrypt the ciphertext
	// before it is reencrypted. The default value, SYMMETRIC_DEFAULT , represents the
	// algorithm used for symmetric encryption KMS keys. Specify the same algorithm
	// that was used to encrypt the ciphertext. If you specify a different algorithm,
	// the decrypt attempt fails. This parameter is required only when the ciphertext
	// was encrypted under an asymmetric KMS key.
	SourceEncryptionAlgorithm types.EncryptionAlgorithmSpec

	// Specifies the encryption context to use to decrypt the ciphertext. Enter the
	// same encryption context that was used to encrypt the ciphertext. An encryption
	// context is a collection of non-secret key-value pairs that represent additional
	// authenticated data. When you use an encryption context to encrypt data, you must
	// specify the same (an exact case-sensitive match) encryption context to decrypt
	// the data. An encryption context is supported only on operations with symmetric
	// encryption KMS keys. On operations with symmetric encryption KMS keys, an
	// encryption context is optional, but it is strongly recommended. For more
	// information, see Encryption context (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#encrypt_context)
	// in the Key Management Service Developer Guide.
	SourceEncryptionContext map[string]string

	// Specifies the KMS key that KMS will use to decrypt the ciphertext before it is
	// re-encrypted. Enter a key ID of the KMS key that was used to encrypt the
	// ciphertext. If you identify a different KMS key, the ReEncrypt operation throws
	// an IncorrectKeyException . This parameter is required only when the ciphertext
	// was encrypted under an asymmetric KMS key. If you used a symmetric encryption
	// KMS key, KMS can get the KMS key from metadata that it adds to the symmetric
	// ciphertext blob. However, it is always recommended as a best practice. This
	// practice ensures that you use the KMS key that you intend. To specify a KMS key,
	// use its key ID, key ARN, alias name, or alias ARN. When using an alias name,
	// prefix it with "alias/" . To specify a KMS key in a different Amazon Web
	// Services account, you must use the key ARN or alias ARN. For example:
	//   - Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//   - Key ARN:
	//   arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//   - Alias name: alias/ExampleAlias
	//   - Alias ARN: arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	// To get the key ID and key ARN for a KMS key, use ListKeys or DescribeKey . To
	// get the alias name and alias ARN, use ListAliases .
	SourceKeyId *string

	noSmithyDocumentSerde
}

type ReEncryptOutput struct {

	// The reencrypted data. When you use the HTTP API or the Amazon Web Services CLI,
	// the value is Base64-encoded. Otherwise, it is not Base64-encoded.
	CiphertextBlob []byte

	// The encryption algorithm that was used to reencrypt the data.
	DestinationEncryptionAlgorithm types.EncryptionAlgorithmSpec

	// The Amazon Resource Name ( key ARN (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#key-id-key-ARN)
	// ) of the KMS key that was used to reencrypt the data.
	KeyId *string

	// The encryption algorithm that was used to decrypt the ciphertext before it was
	// reencrypted.
	SourceEncryptionAlgorithm types.EncryptionAlgorithmSpec

	// Unique identifier of the KMS key used to originally encrypt the data.
	SourceKeyId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationReEncryptMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpReEncrypt{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpReEncrypt{}, middleware.After)
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
	if err = addOpReEncryptValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opReEncrypt(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opReEncrypt(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "ReEncrypt",
	}
}
