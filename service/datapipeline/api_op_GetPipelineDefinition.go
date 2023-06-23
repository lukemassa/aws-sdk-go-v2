// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the definition of the specified pipeline. You can call
// GetPipelineDefinition to retrieve the pipeline definition that you provided
// using PutPipelineDefinition . POST / HTTP/1.1 Content-Type:
// application/x-amz-json-1.1 X-Amz-Target: DataPipeline.GetPipelineDefinition
// Content-Length: 40 Host: datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon,
// 12 Nov 2012 17:49:52 GMT Authorization: AuthParams {"pipelineId":
// "df-06372391ZG65EXAMPLE"} x-amzn-RequestId: e28309e5-0776-11e2-8a14-21bb8a1f50ef
// Content-Type: application/x-amz-json-1.1 Content-Length: 890 Date: Mon, 12 Nov
// 2012 17:50:53 GMT {"pipelineObjects": [ {"fields": [ {"key": "workerGroup",
// "stringValue": "workerGroup"} ], "id": "Default", "name": "Default"}, {"fields":
// [ {"key": "startDateTime", "stringValue": "2012-09-25T17:00:00"}, {"key":
// "type", "stringValue": "Schedule"}, {"key": "period", "stringValue": "1 hour"},
// {"key": "endDateTime", "stringValue": "2012-09-25T18:00:00"} ], "id":
// "Schedule", "name": "Schedule"}, {"fields": [ {"key": "schedule", "refValue":
// "Schedule"}, {"key": "command", "stringValue": "echo hello"}, {"key": "parent",
// "refValue": "Default"}, {"key": "type", "stringValue": "ShellCommandActivity"}
// ], "id": "SayHello", "name": "SayHello"} ] }
func (c *Client) GetPipelineDefinition(ctx context.Context, params *GetPipelineDefinitionInput, optFns ...func(*Options)) (*GetPipelineDefinitionOutput, error) {
	if params == nil {
		params = &GetPipelineDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPipelineDefinition", params, optFns, c.addOperationGetPipelineDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPipelineDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for GetPipelineDefinition.
type GetPipelineDefinitionInput struct {

	// The ID of the pipeline.
	//
	// This member is required.
	PipelineId *string

	// The version of the pipeline definition to retrieve. Set this parameter to latest
	// (default) to use the last definition saved to the pipeline or active to use the
	// last definition that was activated.
	Version *string

	noSmithyDocumentSerde
}

// Contains the output of GetPipelineDefinition.
type GetPipelineDefinitionOutput struct {

	// The parameter objects used in the pipeline definition.
	ParameterObjects []types.ParameterObject

	// The parameter values used in the pipeline definition.
	ParameterValues []types.ParameterValue

	// The objects defined in the pipeline.
	PipelineObjects []types.PipelineObject

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPipelineDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetPipelineDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPipelineDefinition{}, middleware.After)
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
	if err = addOpGetPipelineDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPipelineDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetPipelineDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "GetPipelineDefinition",
	}
}
