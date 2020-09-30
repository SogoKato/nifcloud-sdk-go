// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package hatoba

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type DeleteSnapshotsInput struct {
	_ struct{} `type:"structure"`

	Names *string `location:"querystring" locationName:"names" type:"string"`
}

// String returns the string representation
func (s DeleteSnapshotsInput) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteSnapshotsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Names != nil {
		v := *s.Names

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "names", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteSnapshotsOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Snapshots []Snapshot `locationName:"snapshots" type:"list"`
}

// String returns the string representation
func (s DeleteSnapshotsOutput) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteSnapshotsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Snapshots != nil {
		v := s.Snapshots

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "snapshots", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opDeleteSnapshots = "DeleteSnapshots"

// DeleteSnapshotsRequest returns a request value for making API operation for
// NIFCLOUD Hatoba beta.
//
//    // Example sending a request using DeleteSnapshotsRequest.
//    req := client.DeleteSnapshotsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/hatoba/DeleteSnapshots.htm
func (c *Client) DeleteSnapshotsRequest(input *DeleteSnapshotsInput) DeleteSnapshotsRequest {
	op := &aws.Operation{
		Name:       opDeleteSnapshots,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/snapshots",
	}

	if input == nil {
		input = &DeleteSnapshotsInput{}
	}

	req := c.newRequest(op, input, &DeleteSnapshotsOutput{})

	return DeleteSnapshotsRequest{Request: req, Input: input, Copy: c.DeleteSnapshotsRequest}
}

// DeleteSnapshotsRequest is the request type for the
// DeleteSnapshots API operation.
type DeleteSnapshotsRequest struct {
	*aws.Request
	Input *DeleteSnapshotsInput
	Copy  func(*DeleteSnapshotsInput) DeleteSnapshotsRequest
}

// Send marshals and sends the DeleteSnapshots API request.
func (r DeleteSnapshotsRequest) Send(ctx context.Context) (*DeleteSnapshotsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteSnapshotsResponse{
		DeleteSnapshotsOutput: r.Request.Data.(*DeleteSnapshotsOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteSnapshotsResponse is the response type for the
// DeleteSnapshots API operation.
type DeleteSnapshotsResponse struct {
	*DeleteSnapshotsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteSnapshots request.
func (r *DeleteSnapshotsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
