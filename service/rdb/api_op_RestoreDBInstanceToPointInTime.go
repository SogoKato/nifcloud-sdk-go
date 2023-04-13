// Code generated by smithy-go-codegen DO NOT EDIT.

package rdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/nifcloud/nifcloud-sdk-go/service/rdb/types"
	"time"
)

func (c *Client) RestoreDBInstanceToPointInTime(ctx context.Context, params *RestoreDBInstanceToPointInTimeInput, optFns ...func(*Options)) (*RestoreDBInstanceToPointInTimeOutput, error) {
	if params == nil {
		params = &RestoreDBInstanceToPointInTimeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RestoreDBInstanceToPointInTime", params, optFns, c.addOperationRestoreDBInstanceToPointInTimeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RestoreDBInstanceToPointInTimeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreDBInstanceToPointInTimeInput struct {

	// This member is required.
	SourceDBInstanceIdentifier *string

	// This member is required.
	TargetDBInstanceIdentifier *string

	AccountingType types.AccountingTypeOfRestoreDBInstanceToPointInTimeRequest

	AutoMinorVersionUpgrade *bool

	AvailabilityZone *string

	DBInstanceClass types.DBInstanceClassOfRestoreDBInstanceToPointInTimeRequest

	DBName *string

	DBSubnetGroupName *string

	Engine *string

	Iops *int32

	LicenseModel *string

	MultiAZ *bool

	NiftyDBParameterGroupName *string

	NiftyDBSecurityGroups []string

	NiftyMasterPrivateAddress *string

	NiftyMultiAZType *int32

	NiftyNetworkId *string

	NiftySlavePrivateAddress *string

	NiftyStorageType *int32

	NiftyVirtualPrivateAddress *string

	OptionGroupName *string

	Port *int32

	PubliclyAccessible *bool

	RestoreTime *time.Time

	UseLatestRestorableTime *bool

	noSmithyDocumentSerde
}

type RestoreDBInstanceToPointInTimeOutput struct {
	DBInstance *types.DBInstance

	ResponseMetadata *types.ResponseMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRestoreDBInstanceToPointInTimeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRestoreDBInstanceToPointInTime{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRestoreDBInstanceToPointInTime{}, middleware.After)
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
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpRestoreDBInstanceToPointInTimeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreDBInstanceToPointInTime(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRestoreDBInstanceToPointInTime(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rdb",
		OperationName: "RestoreDBInstanceToPointInTime",
	}
}
