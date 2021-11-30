// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package hatoba

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type ListNodePoolsInput struct {
	_ struct{} `type:"structure"`

	// ClusterName is a required field
	ClusterName *string `location:"uri" locationName:"ClusterName" type:"string" required:"true"`
}

// String returns the string representation
func (s ListNodePoolsInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListNodePoolsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListNodePoolsInput"}

	if s.ClusterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListNodePoolsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ClusterName != nil {
		v := *s.ClusterName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ClusterName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type ListNodePoolsOutput struct {
	_ struct{} `type:"structure"`

	NodePools []NodePools `locationName:"nodePools" type:"list"`
}

// String returns the string representation
func (s ListNodePoolsOutput) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListNodePoolsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NodePools != nil {
		v := s.NodePools

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "nodePools", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opListNodePools = "ListNodePools"

// ListNodePoolsRequest returns a request value for making API operation for
// NIFCLOUD Kubernetes Service Hatoba.
//
//    // Example sending a request using ListNodePoolsRequest.
//    req := client.ListNodePoolsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/hatoba/ListNodePools.htm
func (c *Client) ListNodePoolsRequest(input *ListNodePoolsInput) ListNodePoolsRequest {
	op := &aws.Operation{
		Name:       opListNodePools,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/clusters/{ClusterName}/nodePools",
	}

	if input == nil {
		input = &ListNodePoolsInput{}
	}

	req := c.newRequest(op, input, &ListNodePoolsOutput{})

	return ListNodePoolsRequest{Request: req, Input: input, Copy: c.ListNodePoolsRequest}
}

// ListNodePoolsRequest is the request type for the
// ListNodePools API operation.
type ListNodePoolsRequest struct {
	*aws.Request
	Input *ListNodePoolsInput
	Copy  func(*ListNodePoolsInput) ListNodePoolsRequest
}

// Send marshals and sends the ListNodePools API request.
func (r ListNodePoolsRequest) Send(ctx context.Context) (*ListNodePoolsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListNodePoolsResponse{
		ListNodePoolsOutput: r.Request.Data.(*ListNodePoolsOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ListNodePoolsResponse is the response type for the
// ListNodePools API operation.
type ListNodePoolsResponse struct {
	*ListNodePoolsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListNodePools request.
func (r *ListNodePoolsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
