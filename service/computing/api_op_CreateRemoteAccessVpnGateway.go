// Code generated by smithy-go-codegen DO NOT EDIT.

package computing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/nifcloud/nifcloud-sdk-go/service/computing/types"
)

func (c *Client) CreateRemoteAccessVpnGateway(ctx context.Context, params *CreateRemoteAccessVpnGatewayInput, optFns ...func(*Options)) (*CreateRemoteAccessVpnGatewayOutput, error) {
	if params == nil {
		params = &CreateRemoteAccessVpnGatewayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRemoteAccessVpnGateway", params, optFns, c.addOperationCreateRemoteAccessVpnGatewayMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRemoteAccessVpnGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRemoteAccessVpnGatewayInput struct {

	// This member is required.
	CipherSuite []string

	// This member is required.
	NetworkInterface []types.RequestNetworkInterfaceOfCreateRemoteAccessVpnGateway

	// This member is required.
	PoolNetworkCidr *string

	// This member is required.
	SSLCertificateId *string

	AccountingType *int32

	CACertificateId *string

	Description *string

	Placement *types.RequestPlacementOfCreateRemoteAccessVpnGateway

	RemoteAccessVpnGatewayName *string

	RemoteAccessVpnGatewayType types.RemoteAccessVpnGatewayTypeOfCreateRemoteAccessVpnGatewayRequest

	noSmithyDocumentSerde
}

type CreateRemoteAccessVpnGatewayOutput struct {
	RemoteAccessVpnGateway *types.RemoteAccessVpnGateway

	RequestId *string

	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRemoteAccessVpnGatewayMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateRemoteAccessVpnGateway{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateRemoteAccessVpnGateway{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV2Middleware(stack, options); err != nil {
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
	if err = addOpCreateRemoteAccessVpnGatewayValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRemoteAccessVpnGateway(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRemoteAccessVpnGateway(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateRemoteAccessVpnGateway",
	}
}
