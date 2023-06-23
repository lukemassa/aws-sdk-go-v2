// Code generated by smithy-go-codegen DO NOT EDIT.

package evidently

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/evidently/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a launch of a given feature. Before you create a launch, you must
// create the feature to use for the launch. You can use a launch to safely
// validate new features by serving them to a specified percentage of your users
// while you roll out the feature. You can monitor the performance of the new
// feature to help you decide when to ramp up traffic to more users. This helps you
// reduce risk and identify unintended consequences before you fully launch the
// feature. Don't use this operation to update an existing launch. Instead, use
// UpdateLaunch (https://docs.aws.amazon.com/cloudwatchevidently/latest/APIReference/API_UpdateLaunch.html)
// .
func (c *Client) CreateLaunch(ctx context.Context, params *CreateLaunchInput, optFns ...func(*Options)) (*CreateLaunchOutput, error) {
	if params == nil {
		params = &CreateLaunchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLaunch", params, optFns, c.addOperationCreateLaunchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLaunchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLaunchInput struct {

	// An array of structures that contains the feature and variations that are to be
	// used for the launch.
	//
	// This member is required.
	Groups []types.LaunchGroupConfig

	// The name for the new launch.
	//
	// This member is required.
	Name *string

	// The name or ARN of the project that you want to create the launch in.
	//
	// This member is required.
	Project *string

	// An optional description for the launch.
	Description *string

	// An array of structures that define the metrics that will be used to monitor the
	// launch performance.
	MetricMonitors []types.MetricMonitorConfig

	// When Evidently assigns a particular user session to a launch, it must use a
	// randomization ID to determine which variation the user session is served. This
	// randomization ID is a combination of the entity ID and randomizationSalt . If
	// you omit randomizationSalt , Evidently uses the launch name as the
	// randomizationSalt .
	RandomizationSalt *string

	// An array of structures that define the traffic allocation percentages among the
	// feature variations during each step of the launch.
	ScheduledSplitsConfig *types.ScheduledSplitsLaunchConfig

	// Assigns one or more tags (key-value pairs) to the launch. Tags can help you
	// organize and categorize your resources. You can also use them to scope user
	// permissions by granting a user permission to access or change only resources
	// with certain tag values. Tags don't have any semantic meaning to Amazon Web
	// Services and are interpreted strictly as strings of characters. You can
	// associate as many as 50 tags with a launch. For more information, see Tagging
	// Amazon Web Services resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// .
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateLaunchOutput struct {

	// A structure that contains the configuration of the launch that was created.
	//
	// This member is required.
	Launch *types.Launch

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLaunchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateLaunch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateLaunch{}, middleware.After)
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
	if err = addOpCreateLaunchValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLaunch(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLaunch(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "evidently",
		OperationName: "CreateLaunch",
	}
}
