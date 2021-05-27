// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
	"github.com/nifcloud/nifcloud-sdk-go/private/protocol/rdb"
)

type DeleteDBParameterGroupInput struct {
	_ struct{} `type:"structure"`

	// DBParameterGroupName is a required field
	DBParameterGroupName *string `locationName:"DBParameterGroupName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDBParameterGroupInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDBParameterGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDBParameterGroupInput"}

	if s.DBParameterGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBParameterGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDBParameterGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteDBParameterGroupOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDeleteDBParameterGroup = "DeleteDBParameterGroup"

// DeleteDBParameterGroupRequest returns a request value for making API operation for
// NIFCLOUD RDB.
//
//    // Example sending a request using DeleteDBParameterGroupRequest.
//    req := client.DeleteDBParameterGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rdb/DeleteDBParameterGroup.htm
func (c *Client) DeleteDBParameterGroupRequest(input *DeleteDBParameterGroupInput) DeleteDBParameterGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteDBParameterGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDBParameterGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteDBParameterGroupOutput{})
	req.Handlers.Unmarshal.Remove(rdb.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteDBParameterGroupRequest{Request: req, Input: input, Copy: c.DeleteDBParameterGroupRequest}
}

// DeleteDBParameterGroupRequest is the request type for the
// DeleteDBParameterGroup API operation.
type DeleteDBParameterGroupRequest struct {
	*aws.Request
	Input *DeleteDBParameterGroupInput
	Copy  func(*DeleteDBParameterGroupInput) DeleteDBParameterGroupRequest
}

// Send marshals and sends the DeleteDBParameterGroup API request.
func (r DeleteDBParameterGroupRequest) Send(ctx context.Context) (*DeleteDBParameterGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDBParameterGroupResponse{
		DeleteDBParameterGroupOutput: r.Request.Data.(*DeleteDBParameterGroupOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDBParameterGroupResponse is the response type for the
// DeleteDBParameterGroup API operation.
type DeleteDBParameterGroupResponse struct {
	*DeleteDBParameterGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDBParameterGroup request.
func (r *DeleteDBParameterGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
