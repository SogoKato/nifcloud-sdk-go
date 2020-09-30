// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type ResetExternalMasterInput struct {
	_ struct{} `type:"structure"`

	DBInstanceIdentifier *string `locationName:"DBInstanceIdentifier" type:"string"`
}

// String returns the string representation
func (s ResetExternalMasterInput) String() string {
	return nifcloudutil.Prettify(s)
}

type ResetExternalMasterOutput struct {
	_ struct{} `type:"structure"`

	DBInstance *DBInstance `type:"structure"`
}

// String returns the string representation
func (s ResetExternalMasterOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opResetExternalMaster = "ResetExternalMaster"

// ResetExternalMasterRequest returns a request value for making API operation for
// NIFCLOUD RDB.
//
//    // Example sending a request using ResetExternalMasterRequest.
//    req := client.ResetExternalMasterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rdb/ResetExternalMaster.htm
func (c *Client) ResetExternalMasterRequest(input *ResetExternalMasterInput) ResetExternalMasterRequest {
	op := &aws.Operation{
		Name:       opResetExternalMaster,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResetExternalMasterInput{}
	}

	req := c.newRequest(op, input, &ResetExternalMasterOutput{})

	return ResetExternalMasterRequest{Request: req, Input: input, Copy: c.ResetExternalMasterRequest}
}

// ResetExternalMasterRequest is the request type for the
// ResetExternalMaster API operation.
type ResetExternalMasterRequest struct {
	*aws.Request
	Input *ResetExternalMasterInput
	Copy  func(*ResetExternalMasterInput) ResetExternalMasterRequest
}

// Send marshals and sends the ResetExternalMaster API request.
func (r ResetExternalMasterRequest) Send(ctx context.Context) (*ResetExternalMasterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResetExternalMasterResponse{
		ResetExternalMasterOutput: r.Request.Data.(*ResetExternalMasterOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResetExternalMasterResponse is the response type for the
// ResetExternalMaster API operation.
type ResetExternalMasterResponse struct {
	*ResetExternalMasterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResetExternalMaster request.
func (r *ResetExternalMasterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
