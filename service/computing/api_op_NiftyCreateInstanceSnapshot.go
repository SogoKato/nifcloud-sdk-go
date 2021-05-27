// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type NiftyCreateInstanceSnapshotInput struct {
	_ struct{} `type:"structure"`

	Description *string `locationName:"Description" type:"string"`

	// InstanceId is a required field
	InstanceId *string `locationName:"InstanceId" type:"string" required:"true"`

	// SnapshotName is a required field
	SnapshotName *string `locationName:"SnapshotName" type:"string" required:"true"`
}

// String returns the string representation
func (s NiftyCreateInstanceSnapshotInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *NiftyCreateInstanceSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "NiftyCreateInstanceSnapshotInput"}

	if s.InstanceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceId"))
	}

	if s.SnapshotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnapshotName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type NiftyCreateInstanceSnapshotOutput struct {
	_ struct{} `type:"structure"`

	InstanceSet []InstanceSet `locationName:"instanceSet" locationNameList:"item" type:"list"`

	RequestId *string `locationName:"requestId" type:"string"`

	SnapshotName *string `locationName:"snapshotName" type:"string"`
}

// String returns the string representation
func (s NiftyCreateInstanceSnapshotOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyCreateInstanceSnapshot = "NiftyCreateInstanceSnapshot"

// NiftyCreateInstanceSnapshotRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyCreateInstanceSnapshotRequest.
//    req := client.NiftyCreateInstanceSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/NiftyCreateInstanceSnapshot.htm
func (c *Client) NiftyCreateInstanceSnapshotRequest(input *NiftyCreateInstanceSnapshotInput) NiftyCreateInstanceSnapshotRequest {
	op := &aws.Operation{
		Name:       opNiftyCreateInstanceSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyCreateInstanceSnapshotInput{}
	}

	req := c.newRequest(op, input, &NiftyCreateInstanceSnapshotOutput{})

	return NiftyCreateInstanceSnapshotRequest{Request: req, Input: input, Copy: c.NiftyCreateInstanceSnapshotRequest}
}

// NiftyCreateInstanceSnapshotRequest is the request type for the
// NiftyCreateInstanceSnapshot API operation.
type NiftyCreateInstanceSnapshotRequest struct {
	*aws.Request
	Input *NiftyCreateInstanceSnapshotInput
	Copy  func(*NiftyCreateInstanceSnapshotInput) NiftyCreateInstanceSnapshotRequest
}

// Send marshals and sends the NiftyCreateInstanceSnapshot API request.
func (r NiftyCreateInstanceSnapshotRequest) Send(ctx context.Context) (*NiftyCreateInstanceSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyCreateInstanceSnapshotResponse{
		NiftyCreateInstanceSnapshotOutput: r.Request.Data.(*NiftyCreateInstanceSnapshotOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyCreateInstanceSnapshotResponse is the response type for the
// NiftyCreateInstanceSnapshot API operation.
type NiftyCreateInstanceSnapshotResponse struct {
	*NiftyCreateInstanceSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyCreateInstanceSnapshot request.
func (r *NiftyCreateInstanceSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
