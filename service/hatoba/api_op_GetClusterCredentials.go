// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package hatoba

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type GetClusterCredentialsInput struct {
	_ struct{} `type:"structure"`

	// ClusterName is a required field
	ClusterName *string `location:"uri" locationName:"ClusterName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetClusterCredentialsInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetClusterCredentialsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetClusterCredentialsInput"}

	if s.ClusterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetClusterCredentialsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ClusterName != nil {
		v := *s.ClusterName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ClusterName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetClusterCredentialsOutput struct {
	_ struct{} `type:"structure"`

	Credentials *string `locationName:"credentials" type:"string"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s GetClusterCredentialsOutput) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetClusterCredentialsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Credentials != nil {
		v := *s.Credentials

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "credentials", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opGetClusterCredentials = "GetClusterCredentials"

// GetClusterCredentialsRequest returns a request value for making API operation for
// NIFCLOUD Hatoba beta.
//
//    // Example sending a request using GetClusterCredentialsRequest.
//    req := client.GetClusterCredentialsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/hatoba/GetClusterCredentials.htm
func (c *Client) GetClusterCredentialsRequest(input *GetClusterCredentialsInput) GetClusterCredentialsRequest {
	op := &aws.Operation{
		Name:       opGetClusterCredentials,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/clusters/{ClusterName}/credentials",
	}

	if input == nil {
		input = &GetClusterCredentialsInput{}
	}

	req := c.newRequest(op, input, &GetClusterCredentialsOutput{})

	return GetClusterCredentialsRequest{Request: req, Input: input, Copy: c.GetClusterCredentialsRequest}
}

// GetClusterCredentialsRequest is the request type for the
// GetClusterCredentials API operation.
type GetClusterCredentialsRequest struct {
	*aws.Request
	Input *GetClusterCredentialsInput
	Copy  func(*GetClusterCredentialsInput) GetClusterCredentialsRequest
}

// Send marshals and sends the GetClusterCredentials API request.
func (r GetClusterCredentialsRequest) Send(ctx context.Context) (*GetClusterCredentialsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetClusterCredentialsResponse{
		GetClusterCredentialsOutput: r.Request.Data.(*GetClusterCredentialsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetClusterCredentialsResponse is the response type for the
// GetClusterCredentials API operation.
type GetClusterCredentialsResponse struct {
	*GetClusterCredentialsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetClusterCredentials request.
func (r *GetClusterCredentialsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
