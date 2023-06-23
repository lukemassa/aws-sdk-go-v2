// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	internalEndpointDiscovery "github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new item, or replaces an old item with a new item. If an item that
// has the same primary key as the new item already exists in the specified table,
// the new item completely replaces the existing item. You can perform a
// conditional put operation (add a new item if one with the specified primary key
// doesn't exist), or replace an existing item if it has certain attribute values.
// You can return the item's attribute values in the same operation, using the
// ReturnValues parameter. When you add an item, the primary key attributes are the
// only required attributes. Empty String and Binary attribute values are allowed.
// Attribute values of type String and Binary must have a length greater than zero
// if the attribute is used as a key attribute for a table or index. Set type
// attributes cannot be empty. Invalid Requests with empty values will be rejected
// with a ValidationException exception. To prevent a new item from replacing an
// existing item, use a conditional expression that contains the
// attribute_not_exists function with the name of the attribute being used as the
// partition key for the table. Since every record must contain that attribute, the
// attribute_not_exists function will only succeed if no matching item exists. For
// more information about PutItem , see Working with Items (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithItems.html)
// in the Amazon DynamoDB Developer Guide.
func (c *Client) PutItem(ctx context.Context, params *PutItemInput, optFns ...func(*Options)) (*PutItemOutput, error) {
	if params == nil {
		params = &PutItemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutItem", params, optFns, c.addOperationPutItemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a PutItem operation.
type PutItemInput struct {

	// A map of attribute name/value pairs, one for each attribute. Only the primary
	// key attributes are required; you can optionally provide other attribute
	// name-value pairs for the item. You must provide all of the attributes for the
	// primary key. For example, with a simple primary key, you only need to provide a
	// value for the partition key. For a composite primary key, you must provide both
	// values for both the partition key and the sort key. If you specify any
	// attributes that are part of an index key, then the data types for those
	// attributes must match those of the schema in the table's attribute definition.
	// Empty String and Binary attribute values are allowed. Attribute values of type
	// String and Binary must have a length greater than zero if the attribute is used
	// as a key attribute for a table or index. For more information about primary
	// keys, see Primary Key (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.CoreComponents.html#HowItWorks.CoreComponents.PrimaryKey)
	// in the Amazon DynamoDB Developer Guide. Each element in the Item map is an
	// AttributeValue object.
	//
	// This member is required.
	Item map[string]types.AttributeValue

	// The name of the table to contain the item.
	//
	// This member is required.
	TableName *string

	// A condition that must be satisfied in order for a conditional PutItem operation
	// to succeed. An expression can contain any of the following:
	//   - Functions: attribute_exists | attribute_not_exists | attribute_type |
	//   contains | begins_with | size These function names are case-sensitive.
	//   - Comparison operators: = | <> | < | > | <= | >= | BETWEEN | IN
	//   - Logical operators: AND | OR | NOT
	// For more information on condition expressions, see Condition Expressions (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html)
	// in the Amazon DynamoDB Developer Guide.
	ConditionExpression *string

	// This is a legacy parameter. Use ConditionExpression instead. For more
	// information, see ConditionalOperator (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.ConditionalOperator.html)
	// in the Amazon DynamoDB Developer Guide.
	ConditionalOperator types.ConditionalOperator

	// This is a legacy parameter. Use ConditionExpression instead. For more
	// information, see Expected (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.Expected.html)
	// in the Amazon DynamoDB Developer Guide.
	Expected map[string]types.ExpectedAttributeValue

	// One or more substitution tokens for attribute names in an expression. The
	// following are some use cases for using ExpressionAttributeNames :
	//   - To access an attribute whose name conflicts with a DynamoDB reserved word.
	//   - To create a placeholder for repeating occurrences of an attribute name in
	//   an expression.
	//   - To prevent special characters in an attribute name from being
	//   misinterpreted in an expression.
	// Use the # character in an expression to dereference an attribute name. For
	// example, consider the following attribute name:
	//   - Percentile
	// The name of this attribute conflicts with a reserved word, so it cannot be used
	// directly in an expression. (For the complete list of reserved words, see
	// Reserved Words (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html)
	// in the Amazon DynamoDB Developer Guide). To work around this, you could specify
	// the following for ExpressionAttributeNames :
	//   - {"#P":"Percentile"}
	// You could then use this substitution in an expression, as in this example:
	//   - #P = :val
	// Tokens that begin with the : character are expression attribute values, which
	// are placeholders for the actual value at runtime. For more information on
	// expression attribute names, see Specifying Item Attributes (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html)
	// in the Amazon DynamoDB Developer Guide.
	ExpressionAttributeNames map[string]string

	// One or more values that can be substituted in an expression. Use the : (colon)
	// character in an expression to dereference an attribute value. For example,
	// suppose that you wanted to check whether the value of the ProductStatus
	// attribute was one of the following: Available | Backordered | Discontinued You
	// would first need to specify ExpressionAttributeValues as follows: {
	// ":avail":{"S":"Available"}, ":back":{"S":"Backordered"},
	// ":disc":{"S":"Discontinued"} } You could then use these values in an expression,
	// such as this: ProductStatus IN (:avail, :back, :disc) For more information on
	// expression attribute values, see Condition Expressions (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html)
	// in the Amazon DynamoDB Developer Guide.
	ExpressionAttributeValues map[string]types.AttributeValue

	// Determines the level of detail about either provisioned or on-demand throughput
	// consumption that is returned in the response:
	//   - INDEXES - The response includes the aggregate ConsumedCapacity for the
	//   operation, together with ConsumedCapacity for each table and secondary index
	//   that was accessed. Note that some operations, such as GetItem and BatchGetItem
	//   , do not access any indexes at all. In these cases, specifying INDEXES will
	//   only return ConsumedCapacity information for table(s).
	//   - TOTAL - The response includes only the aggregate ConsumedCapacity for the
	//   operation.
	//   - NONE - No ConsumedCapacity details are included in the response.
	ReturnConsumedCapacity types.ReturnConsumedCapacity

	// Determines whether item collection metrics are returned. If set to SIZE , the
	// response includes statistics about item collections, if any, that were modified
	// during the operation are returned in the response. If set to NONE (the
	// default), no statistics are returned.
	ReturnItemCollectionMetrics types.ReturnItemCollectionMetrics

	// Use ReturnValues if you want to get the item attributes as they appeared before
	// they were updated with the PutItem request. For PutItem , the valid values are:
	//   - NONE - If ReturnValues is not specified, or if its value is NONE , then
	//   nothing is returned. (This setting is the default for ReturnValues .)
	//   - ALL_OLD - If PutItem overwrote an attribute name-value pair, then the
	//   content of the old item is returned.
	// The values returned are strongly consistent. There is no additional cost
	// associated with requesting a return value aside from the small network and
	// processing overhead of receiving a larger response. No read capacity units are
	// consumed. The ReturnValues parameter is used by several DynamoDB operations;
	// however, PutItem does not recognize any values other than NONE or ALL_OLD .
	ReturnValues types.ReturnValue

	noSmithyDocumentSerde
}

// Represents the output of a PutItem operation.
type PutItemOutput struct {

	// The attribute values as they appeared before the PutItem operation, but only if
	// ReturnValues is specified as ALL_OLD in the request. Each element consists of
	// an attribute name and an attribute value.
	Attributes map[string]types.AttributeValue

	// The capacity units consumed by the PutItem operation. The data returned
	// includes the total provisioned throughput consumed, along with statistics for
	// the table and any indexes involved in the operation. ConsumedCapacity is only
	// returned if the ReturnConsumedCapacity parameter was specified. For more
	// information, see Provisioned Throughput (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughputIntro.html)
	// in the Amazon DynamoDB Developer Guide.
	ConsumedCapacity *types.ConsumedCapacity

	// Information about item collections, if any, that were affected by the PutItem
	// operation. ItemCollectionMetrics is only returned if the
	// ReturnItemCollectionMetrics parameter was specified. If the table does not have
	// any local secondary indexes, this information is not returned in the response.
	// Each ItemCollectionMetrics element consists of:
	//   - ItemCollectionKey - The partition key value of the item collection. This is
	//   the same as the partition key value of the item itself.
	//   - SizeEstimateRangeGB - An estimate of item collection size, in gigabytes.
	//   This value is a two-element array containing a lower bound and an upper bound
	//   for the estimate. The estimate includes the size of all the items in the table,
	//   plus the size of all attributes projected into all of the local secondary
	//   indexes on that table. Use this estimate to measure whether a local secondary
	//   index is approaching its size limit. The estimate is subject to change over
	//   time; therefore, do not rely on the precision or accuracy of the estimate.
	ItemCollectionMetrics *types.ItemCollectionMetrics

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutItemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpPutItem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpPutItem{}, middleware.After)
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
	if err = addOpPutItemDiscoverEndpointMiddleware(stack, options, c); err != nil {
		return err
	}
	if err = addOpPutItemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutItem(options.Region), middleware.Before); err != nil {
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
	if err = addValidateResponseChecksum(stack, options); err != nil {
		return err
	}
	if err = addAcceptEncodingGzip(stack, options); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func addOpPutItemDiscoverEndpointMiddleware(stack *middleware.Stack, o Options, c *Client) error {
	return stack.Serialize.Insert(&internalEndpointDiscovery.DiscoverEndpoint{
		Options: []func(*internalEndpointDiscovery.DiscoverEndpointOptions){
			func(opt *internalEndpointDiscovery.DiscoverEndpointOptions) {
				opt.DisableHTTPS = o.EndpointOptions.DisableHTTPS
				opt.Logger = o.Logger
			},
		},
		DiscoverOperation:            c.fetchOpPutItemDiscoverEndpoint,
		EndpointDiscoveryEnableState: o.EndpointDiscovery.EnableEndpointDiscovery,
		EndpointDiscoveryRequired:    false,
	}, "ResolveEndpoint", middleware.After)
}

func (c *Client) fetchOpPutItemDiscoverEndpoint(ctx context.Context, input interface{}, optFns ...func(*internalEndpointDiscovery.DiscoverEndpointOptions)) (internalEndpointDiscovery.WeightedAddress, error) {
	in, ok := input.(*PutItemInput)
	if !ok {
		return internalEndpointDiscovery.WeightedAddress{}, fmt.Errorf("unknown input type %T", input)
	}
	_ = in

	identifierMap := make(map[string]string, 0)

	key := fmt.Sprintf("DynamoDB.%v", identifierMap)

	if v, ok := c.endpointCache.Get(key); ok {
		return v, nil
	}

	discoveryOperationInput := &DescribeEndpointsInput{}

	opt := internalEndpointDiscovery.DiscoverEndpointOptions{}
	for _, fn := range optFns {
		fn(&opt)
	}

	go c.handleEndpointDiscoveryFromService(ctx, discoveryOperationInput, key, opt)
	return internalEndpointDiscovery.WeightedAddress{}, nil
}

func newServiceMetadataMiddleware_opPutItem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "PutItem",
	}
}
