// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type NiftyDescribeAlarmsPartitionsInput struct {
	_ struct{} `type:"structure"`

	InstanceId []string `locationName:"InstanceId" type:"list"`
}

// String returns the string representation
func (s NiftyDescribeAlarmsPartitionsInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyDescribeAlarmsPartitionsOutput struct {
	_ struct{} `type:"structure"`

	AlarmTargetSet []AlarmTargetSet `locationName:"alarmTargetSet" locationNameList:"item" type:"list"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s NiftyDescribeAlarmsPartitionsOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyDescribeAlarmsPartitions = "NiftyDescribeAlarmsPartitions"

// NiftyDescribeAlarmsPartitionsRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyDescribeAlarmsPartitionsRequest.
//    req := client.NiftyDescribeAlarmsPartitionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/NiftyDescribeAlarmsPartitions.htm
func (c *Client) NiftyDescribeAlarmsPartitionsRequest(input *NiftyDescribeAlarmsPartitionsInput) NiftyDescribeAlarmsPartitionsRequest {
	op := &aws.Operation{
		Name:       opNiftyDescribeAlarmsPartitions,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyDescribeAlarmsPartitionsInput{}
	}

	req := c.newRequest(op, input, &NiftyDescribeAlarmsPartitionsOutput{})

	return NiftyDescribeAlarmsPartitionsRequest{Request: req, Input: input, Copy: c.NiftyDescribeAlarmsPartitionsRequest}
}

// NiftyDescribeAlarmsPartitionsRequest is the request type for the
// NiftyDescribeAlarmsPartitions API operation.
type NiftyDescribeAlarmsPartitionsRequest struct {
	*aws.Request
	Input *NiftyDescribeAlarmsPartitionsInput
	Copy  func(*NiftyDescribeAlarmsPartitionsInput) NiftyDescribeAlarmsPartitionsRequest
}

// Send marshals and sends the NiftyDescribeAlarmsPartitions API request.
func (r NiftyDescribeAlarmsPartitionsRequest) Send(ctx context.Context) (*NiftyDescribeAlarmsPartitionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyDescribeAlarmsPartitionsResponse{
		NiftyDescribeAlarmsPartitionsOutput: r.Request.Data.(*NiftyDescribeAlarmsPartitionsOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyDescribeAlarmsPartitionsResponse is the response type for the
// NiftyDescribeAlarmsPartitions API operation.
type NiftyDescribeAlarmsPartitionsResponse struct {
	*NiftyDescribeAlarmsPartitionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyDescribeAlarmsPartitions request.
func (r *NiftyDescribeAlarmsPartitionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
