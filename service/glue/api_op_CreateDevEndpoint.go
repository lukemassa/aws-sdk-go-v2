// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new development endpoint.
func (c *Client) CreateDevEndpoint(ctx context.Context, params *CreateDevEndpointInput, optFns ...func(*Options)) (*CreateDevEndpointOutput, error) {
	if params == nil {
		params = &CreateDevEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDevEndpoint", params, optFns, c.addOperationCreateDevEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDevEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDevEndpointInput struct {

	// The name to be assigned to the new DevEndpoint .
	//
	// This member is required.
	EndpointName *string

	// The IAM role for the DevEndpoint .
	//
	// This member is required.
	RoleArn *string

	// A map of arguments used to configure the DevEndpoint .
	Arguments map[string]string

	// The path to one or more Java .jar files in an S3 bucket that should be loaded
	// in your DevEndpoint .
	ExtraJarsS3Path *string

	// The paths to one or more Python libraries in an Amazon S3 bucket that should be
	// loaded in your DevEndpoint . Multiple values must be complete paths separated by
	// a comma. You can only use pure Python libraries with a DevEndpoint . Libraries
	// that rely on C extensions, such as the pandas (http://pandas.pydata.org/)
	// Python data analysis library, are not yet supported.
	ExtraPythonLibsS3Path *string

	// Glue version determines the versions of Apache Spark and Python that Glue
	// supports. The Python version indicates the version supported for running your
	// ETL scripts on development endpoints. For more information about the available
	// Glue versions and corresponding Spark and Python versions, see Glue version (https://docs.aws.amazon.com/glue/latest/dg/add-job.html)
	// in the developer guide. Development endpoints that are created without
	// specifying a Glue version default to Glue 0.9. You can specify a version of
	// Python support for development endpoints by using the Arguments parameter in
	// the CreateDevEndpoint or UpdateDevEndpoint APIs. If no arguments are provided,
	// the version defaults to Python 2.
	GlueVersion *string

	// The number of Glue Data Processing Units (DPUs) to allocate to this DevEndpoint .
	NumberOfNodes int32

	// The number of workers of a defined workerType that are allocated to the
	// development endpoint. The maximum number of workers you can define are 299 for
	// G.1X , and 149 for G.2X .
	NumberOfWorkers *int32

	// The public key to be used by this DevEndpoint for authentication. This
	// attribute is provided for backward compatibility because the recommended
	// attribute to use is public keys.
	PublicKey *string

	// A list of public keys to be used by the development endpoints for
	// authentication. The use of this attribute is preferred over a single public key
	// because the public keys allow you to have a different private key per client. If
	// you previously created an endpoint with a public key, you must remove that key
	// to be able to set a list of public keys. Call the UpdateDevEndpoint API with
	// the public key content in the deletePublicKeys attribute, and the list of new
	// keys in the addPublicKeys attribute.
	PublicKeys []string

	// The name of the SecurityConfiguration structure to be used with this DevEndpoint
	// .
	SecurityConfiguration *string

	// Security group IDs for the security groups to be used by the new DevEndpoint .
	SecurityGroupIds []string

	// The subnet ID for the new DevEndpoint to use.
	SubnetId *string

	// The tags to use with this DevEndpoint. You may use tags to limit access to the
	// DevEndpoint. For more information about tags in Glue, see Amazon Web Services
	// Tags in Glue (https://docs.aws.amazon.com/glue/latest/dg/monitor-tags.html) in
	// the developer guide.
	Tags map[string]string

	// The type of predefined worker that is allocated to the development endpoint.
	// Accepts a value of Standard, G.1X, or G.2X.
	//   - For the Standard worker type, each worker provides 4 vCPU, 16 GB of memory
	//   and a 50GB disk, and 2 executors per worker.
	//   - For the G.1X worker type, each worker maps to 1 DPU (4 vCPU, 16 GB of
	//   memory, 64 GB disk), and provides 1 executor per worker. We recommend this
	//   worker type for memory-intensive jobs.
	//   - For the G.2X worker type, each worker maps to 2 DPU (8 vCPU, 32 GB of
	//   memory, 128 GB disk), and provides 1 executor per worker. We recommend this
	//   worker type for memory-intensive jobs.
	// Known issue: when a development endpoint is created with the G.2X WorkerType
	// configuration, the Spark drivers for the development endpoint will run on 4
	// vCPU, 16 GB of memory, and a 64 GB disk.
	WorkerType types.WorkerType

	noSmithyDocumentSerde
}

type CreateDevEndpointOutput struct {

	// The map of arguments used to configure this DevEndpoint . Valid arguments are:
	//   - "--enable-glue-datacatalog": ""
	// You can specify a version of Python support for development endpoints by using
	// the Arguments parameter in the CreateDevEndpoint or UpdateDevEndpoint APIs. If
	// no arguments are provided, the version defaults to Python 2.
	Arguments map[string]string

	// The Amazon Web Services Availability Zone where this DevEndpoint is located.
	AvailabilityZone *string

	// The point in time at which this DevEndpoint was created.
	CreatedTimestamp *time.Time

	// The name assigned to the new DevEndpoint .
	EndpointName *string

	// Path to one or more Java .jar files in an S3 bucket that will be loaded in your
	// DevEndpoint .
	ExtraJarsS3Path *string

	// The paths to one or more Python libraries in an S3 bucket that will be loaded
	// in your DevEndpoint .
	ExtraPythonLibsS3Path *string

	// The reason for a current failure in this DevEndpoint .
	FailureReason *string

	// Glue version determines the versions of Apache Spark and Python that Glue
	// supports. The Python version indicates the version supported for running your
	// ETL scripts on development endpoints. For more information about the available
	// Glue versions and corresponding Spark and Python versions, see Glue version (https://docs.aws.amazon.com/glue/latest/dg/add-job.html)
	// in the developer guide.
	GlueVersion *string

	// The number of Glue Data Processing Units (DPUs) allocated to this DevEndpoint.
	NumberOfNodes int32

	// The number of workers of a defined workerType that are allocated to the
	// development endpoint.
	NumberOfWorkers *int32

	// The Amazon Resource Name (ARN) of the role assigned to the new DevEndpoint .
	RoleArn *string

	// The name of the SecurityConfiguration structure being used with this DevEndpoint
	// .
	SecurityConfiguration *string

	// The security groups assigned to the new DevEndpoint .
	SecurityGroupIds []string

	// The current status of the new DevEndpoint .
	Status *string

	// The subnet ID assigned to the new DevEndpoint .
	SubnetId *string

	// The ID of the virtual private cloud (VPC) used by this DevEndpoint .
	VpcId *string

	// The type of predefined worker that is allocated to the development endpoint.
	// May be a value of Standard, G.1X, or G.2X.
	WorkerType types.WorkerType

	// The address of the YARN endpoint used by this DevEndpoint .
	YarnEndpointAddress *string

	// The Apache Zeppelin port for the remote Apache Spark interpreter.
	ZeppelinRemoteSparkInterpreterPort int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDevEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDevEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDevEndpoint{}, middleware.After)
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
	if err = addOpCreateDevEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDevEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDevEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "CreateDevEndpoint",
	}
}
