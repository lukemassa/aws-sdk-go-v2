// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Schedules a query of a log group using CloudWatch Logs Insights. You specify
// the log group and time range to query and the query string to use. For more
// information, see CloudWatch Logs Insights Query Syntax (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html)
// . Queries time out after 60 minutes of runtime. If your queries are timing out,
// reduce the time range being searched or partition your query into a number of
// queries. If you are using CloudWatch cross-account observability, you can use
// this operation in a monitoring account to start a query in a linked source
// account. For more information, see CloudWatch cross-account observability (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-Unified-Cross-Account.html)
// . For a cross-account StartQuery operation, the query definition must be
// defined in the monitoring account. You can have up to 30 concurrent CloudWatch
// Logs insights queries, including queries that have been added to dashboards.
func (c *Client) StartQuery(ctx context.Context, params *StartQueryInput, optFns ...func(*Options)) (*StartQueryOutput, error) {
	if params == nil {
		params = &StartQueryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartQuery", params, optFns, c.addOperationStartQueryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartQueryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartQueryInput struct {

	// The end of the time range to query. The range is inclusive, so the specified
	// end time is included in the query. Specified as epoch time, the number of
	// seconds since January 1, 1970, 00:00:00 UTC .
	//
	// This member is required.
	EndTime *int64

	// The query string to use. For more information, see CloudWatch Logs Insights
	// Query Syntax (https://docs.aws.amazon.com/AmazonCloudWatch/latest/logs/CWL_QuerySyntax.html)
	// .
	//
	// This member is required.
	QueryString *string

	// The beginning of the time range to query. The range is inclusive, so the
	// specified start time is included in the query. Specified as epoch time, the
	// number of seconds since January 1, 1970, 00:00:00 UTC .
	//
	// This member is required.
	StartTime *int64

	// The maximum number of log events to return in the query. If the query string
	// uses the fields command, only the specified fields and their values are
	// returned. The default is 1000.
	Limit *int32

	// The list of log groups to query. You can include up to 50 log groups. You can
	// specify them by the log group name or ARN. If a log group that you're querying
	// is in a source account and you're using a monitoring account, you must specify
	// the ARN of the log group here. The query definition must also be defined in the
	// monitoring account. If you specify an ARN, the ARN can't end with an asterisk
	// (*). A StartQuery operation must include exactly one of the following
	// parameters: logGroupName , logGroupNames or logGroupIdentifiers .
	LogGroupIdentifiers []string

	// The log group on which to perform the query. A StartQuery operation must
	// include exactly one of the following parameters: logGroupName , logGroupNames
	// or logGroupIdentifiers .
	LogGroupName *string

	// The list of log groups to be queried. You can include up to 50 log groups. A
	// StartQuery operation must include exactly one of the following parameters:
	// logGroupName , logGroupNames or logGroupIdentifiers .
	LogGroupNames []string

	noSmithyDocumentSerde
}

type StartQueryOutput struct {

	// The unique ID of the query.
	QueryId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartQueryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartQuery{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartQuery{}, middleware.After)
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
	if err = addOpStartQueryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartQuery(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartQuery(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "StartQuery",
	}
}
