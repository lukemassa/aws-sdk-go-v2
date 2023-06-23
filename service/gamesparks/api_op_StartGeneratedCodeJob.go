// Code generated by smithy-go-codegen DO NOT EDIT.

package gamesparks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamesparks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an asynchronous process that generates client code for system-defined
// and custom messages. The resulting code is collected as a .zip file and uploaded
// to a pre-signed Amazon S3 URL.
func (c *Client) StartGeneratedCodeJob(ctx context.Context, params *StartGeneratedCodeJobInput, optFns ...func(*Options)) (*StartGeneratedCodeJobOutput, error) {
	if params == nil {
		params = &StartGeneratedCodeJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartGeneratedCodeJob", params, optFns, c.addOperationStartGeneratedCodeJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartGeneratedCodeJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartGeneratedCodeJobInput struct {

	// The name of the game.
	//
	// This member is required.
	GameName *string

	// Properties of the generator to use for the job.
	//
	// This member is required.
	Generator *types.Generator

	// The identifier of the snapshot for which to generate code.
	//
	// This member is required.
	SnapshotId *string

	noSmithyDocumentSerde
}

type StartGeneratedCodeJobOutput struct {

	// The identifier of the code generation job. You can use this identifier in the
	// GetGeneratedCodeJob operation.
	GeneratedCodeJobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartGeneratedCodeJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartGeneratedCodeJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartGeneratedCodeJob{}, middleware.After)
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
	if err = addOpStartGeneratedCodeJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartGeneratedCodeJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartGeneratedCodeJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamesparks",
		OperationName: "StartGeneratedCodeJob",
	}
}
