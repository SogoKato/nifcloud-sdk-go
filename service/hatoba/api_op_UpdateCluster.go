// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package hatoba

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type UpdateClusterInput struct {
	_ struct{} `type:"structure"`

	// Cluster is a required field
	Cluster *UpdateClusterRequestCluster `locationName:"cluster" type:"structure" required:"true"`

	// ClusterName is a required field
	ClusterName *string `location:"uri" locationName:"ClusterName" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateClusterInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateClusterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateClusterInput"}

	if s.Cluster == nil {
		invalidParams.Add(aws.NewErrParamRequired("Cluster"))
	}

	if s.ClusterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterName"))
	}
	if s.Cluster != nil {
		if err := s.Cluster.Validate(); err != nil {
			invalidParams.AddNested("Cluster", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateClusterInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Cluster != nil {
		v := s.Cluster

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "cluster", v, metadata)
	}
	if s.ClusterName != nil {
		v := *s.ClusterName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ClusterName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateClusterOutput struct {
	_ struct{} `type:"structure"`

	Cluster *Cluster `locationName:"cluster" type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s UpdateClusterOutput) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateClusterOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Cluster != nil {
		v := s.Cluster

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "cluster", v, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateCluster = "UpdateCluster"

// UpdateClusterRequest returns a request value for making API operation for
// NIFCLOUD Hatoba beta.
//
//    // Example sending a request using UpdateClusterRequest.
//    req := client.UpdateClusterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/hatoba/UpdateCluster.htm
func (c *Client) UpdateClusterRequest(input *UpdateClusterInput) UpdateClusterRequest {
	op := &aws.Operation{
		Name:       opUpdateCluster,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/clusters/{ClusterName}",
	}

	if input == nil {
		input = &UpdateClusterInput{}
	}

	req := c.newRequest(op, input, &UpdateClusterOutput{})

	return UpdateClusterRequest{Request: req, Input: input, Copy: c.UpdateClusterRequest}
}

// UpdateClusterRequest is the request type for the
// UpdateCluster API operation.
type UpdateClusterRequest struct {
	*aws.Request
	Input *UpdateClusterInput
	Copy  func(*UpdateClusterInput) UpdateClusterRequest
}

// Send marshals and sends the UpdateCluster API request.
func (r UpdateClusterRequest) Send(ctx context.Context) (*UpdateClusterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateClusterResponse{
		UpdateClusterOutput: r.Request.Data.(*UpdateClusterOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateClusterResponse is the response type for the
// UpdateCluster API operation.
type UpdateClusterResponse struct {
	*UpdateClusterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateCluster request.
func (r *UpdateClusterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
