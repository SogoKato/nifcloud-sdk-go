// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type NiftyDeleteAlarmInput struct {
	_ struct{} `type:"structure"`

	FunctionName FunctionNameOfNiftyDeleteAlarmRequest `locationName:"FunctionName" type:"string" enum:"true"`

	RuleName *string `locationName:"RuleName" type:"string"`
}

// String returns the string representation
func (s NiftyDeleteAlarmInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyDeleteAlarmOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s NiftyDeleteAlarmOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyDeleteAlarm = "NiftyDeleteAlarm"

// NiftyDeleteAlarmRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyDeleteAlarmRequest.
//    req := client.NiftyDeleteAlarmRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/NiftyDeleteAlarm.htm
func (c *Client) NiftyDeleteAlarmRequest(input *NiftyDeleteAlarmInput) NiftyDeleteAlarmRequest {
	op := &aws.Operation{
		Name:       opNiftyDeleteAlarm,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyDeleteAlarmInput{}
	}

	req := c.newRequest(op, input, &NiftyDeleteAlarmOutput{})

	return NiftyDeleteAlarmRequest{Request: req, Input: input, Copy: c.NiftyDeleteAlarmRequest}
}

// NiftyDeleteAlarmRequest is the request type for the
// NiftyDeleteAlarm API operation.
type NiftyDeleteAlarmRequest struct {
	*aws.Request
	Input *NiftyDeleteAlarmInput
	Copy  func(*NiftyDeleteAlarmInput) NiftyDeleteAlarmRequest
}

// Send marshals and sends the NiftyDeleteAlarm API request.
func (r NiftyDeleteAlarmRequest) Send(ctx context.Context) (*NiftyDeleteAlarmResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyDeleteAlarmResponse{
		NiftyDeleteAlarmOutput: r.Request.Data.(*NiftyDeleteAlarmOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyDeleteAlarmResponse is the response type for the
// NiftyDeleteAlarm API operation.
type NiftyDeleteAlarmResponse struct {
	*NiftyDeleteAlarmOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyDeleteAlarm request.
func (r *NiftyDeleteAlarmResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
