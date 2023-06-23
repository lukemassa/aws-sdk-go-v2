// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a usage limit for a specified Amazon Redshift feature on a cluster. The
// usage limit is identified by the returned usage limit identifier.
func (c *Client) CreateUsageLimit(ctx context.Context, params *CreateUsageLimitInput, optFns ...func(*Options)) (*CreateUsageLimitOutput, error) {
	if params == nil {
		params = &CreateUsageLimitInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUsageLimit", params, optFns, c.addOperationCreateUsageLimitMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUsageLimitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateUsageLimitInput struct {

	// The limit amount. If time-based, this amount is in minutes. If data-based, this
	// amount is in terabytes (TB). The value must be a positive number.
	//
	// This member is required.
	Amount int64

	// The identifier of the cluster that you want to limit usage.
	//
	// This member is required.
	ClusterIdentifier *string

	// The Amazon Redshift feature that you want to limit.
	//
	// This member is required.
	FeatureType types.UsageLimitFeatureType

	// The type of limit. Depending on the feature type, this can be based on a time
	// duration or data size. If FeatureType is spectrum , then LimitType must be
	// data-scanned . If FeatureType is concurrency-scaling , then LimitType must be
	// time . If FeatureType is cross-region-datasharing , then LimitType must be
	// data-scanned .
	//
	// This member is required.
	LimitType types.UsageLimitLimitType

	// The action that Amazon Redshift takes when the limit is reached. The default is
	// log. For more information about this parameter, see UsageLimit .
	BreachAction types.UsageLimitBreachAction

	// The time period that the amount applies to. A weekly period begins on Sunday.
	// The default is monthly .
	Period types.UsageLimitPeriod

	// A list of tag instances.
	Tags []types.Tag

	noSmithyDocumentSerde
}

// Describes a usage limit object for a cluster.
type CreateUsageLimitOutput struct {

	// The limit amount. If time-based, this amount is in minutes. If data-based, this
	// amount is in terabytes (TB).
	Amount int64

	// The action that Amazon Redshift takes when the limit is reached. Possible
	// values are:
	//   - log - To log an event in a system table. The default is log.
	//   - emit-metric - To emit CloudWatch metrics.
	//   - disable - To disable the feature until the next usage period begins.
	BreachAction types.UsageLimitBreachAction

	// The identifier of the cluster with a usage limit.
	ClusterIdentifier *string

	// The Amazon Redshift feature to which the limit applies.
	FeatureType types.UsageLimitFeatureType

	// The type of limit. Depending on the feature type, this can be based on a time
	// duration or data size.
	LimitType types.UsageLimitLimitType

	// The time period that the amount applies to. A weekly period begins on Sunday.
	// The default is monthly .
	Period types.UsageLimitPeriod

	// A list of tag instances.
	Tags []types.Tag

	// The identifier of the usage limit.
	UsageLimitId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateUsageLimitMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateUsageLimit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateUsageLimit{}, middleware.After)
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
	if err = addOpCreateUsageLimitValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUsageLimit(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateUsageLimit(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "CreateUsageLimit",
	}
}
