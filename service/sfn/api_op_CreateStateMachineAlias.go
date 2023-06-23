// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates an alias (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-alias.html)
// for a state machine that points to one or two versions (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-state-machine-version.html)
// of the same state machine. You can set your application to call StartExecution
// with an alias and update the version the alias uses without changing the
// client's code. You can also map an alias to split StartExecution requests
// between two versions of a state machine. To do this, add a second RoutingConfig
// object in the routingConfiguration parameter. You must also specify the
// percentage of execution run requests each version should receive in both
// RoutingConfig objects. Step Functions randomly chooses which version runs a
// given execution based on the percentage you specify. To create an alias that
// points to a single version, specify a single RoutingConfig object with a weight
// set to 100. You can create up to 100 aliases for each state machine. You must
// delete unused aliases using the DeleteStateMachineAlias API action.
// CreateStateMachineAlias is an idempotent API. Step Functions bases the
// idempotency check on the stateMachineArn , description , name , and
// routingConfiguration parameters. Requests that contain the same values for these
// parameters return a successful idempotent response without creating a duplicate
// resource. Related operations:
//   - DescribeStateMachineAlias
//   - ListStateMachineAliases
//   - UpdateStateMachineAlias
//   - DeleteStateMachineAlias
func (c *Client) CreateStateMachineAlias(ctx context.Context, params *CreateStateMachineAliasInput, optFns ...func(*Options)) (*CreateStateMachineAliasOutput, error) {
	if params == nil {
		params = &CreateStateMachineAliasInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateStateMachineAlias", params, optFns, c.addOperationCreateStateMachineAliasMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateStateMachineAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateStateMachineAliasInput struct {

	// The name of the state machine alias. To avoid conflict with version ARNs, don't
	// use an integer in the name of the alias.
	//
	// This member is required.
	Name *string

	// The routing configuration of a state machine alias. The routing configuration
	// shifts execution traffic between two state machine versions.
	// routingConfiguration contains an array of RoutingConfig objects that specify up
	// to two state machine versions. Step Functions then randomly choses which version
	// to run an execution with based on the weight assigned to each RoutingConfig .
	//
	// This member is required.
	RoutingConfiguration []types.RoutingConfigurationListItem

	// A description for the state machine alias.
	Description *string

	noSmithyDocumentSerde
}

type CreateStateMachineAliasOutput struct {

	// The date the state machine alias was created.
	//
	// This member is required.
	CreationDate *time.Time

	// The Amazon Resource Name (ARN) that identifies the created state machine alias.
	//
	// This member is required.
	StateMachineAliasArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateStateMachineAliasMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateStateMachineAlias{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateStateMachineAlias{}, middleware.After)
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
	if err = addOpCreateStateMachineAliasValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateStateMachineAlias(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateStateMachineAlias(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "CreateStateMachineAlias",
	}
}
