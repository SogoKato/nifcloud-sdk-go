// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type DescribeInstanceHealthInput struct {
	_ struct{} `type:"structure"`

	// InstancePort is a required field
	InstancePort *int64 `locationName:"InstancePort" type:"integer" required:"true"`

	Instances []RequestInstances `locationName:"Instances" locationNameList:"member" type:"list"`

	// LoadBalancerName is a required field
	LoadBalancerName *string `locationName:"LoadBalancerName" type:"string" required:"true"`

	// LoadBalancerPort is a required field
	LoadBalancerPort *int64 `locationName:"LoadBalancerPort" type:"integer" required:"true"`
}

// String returns the string representation
func (s DescribeInstanceHealthInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeInstanceHealthInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeInstanceHealthInput"}

	if s.InstancePort == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstancePort"))
	}

	if s.LoadBalancerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerName"))
	}

	if s.LoadBalancerPort == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerPort"))
	}
	if s.Instances != nil {
		for i, v := range s.Instances {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Instances", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeInstanceHealthOutput struct {
	_ struct{} `type:"structure"`

	InstanceStates []InstanceStates `locationName:"InstanceStates" locationNameList:"member" type:"list"`

	ResponseMetadata *ResponseMetadata `locationName:"ResponseMetadata" type:"structure"`
}

// String returns the string representation
func (s DescribeInstanceHealthOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDescribeInstanceHealth = "DescribeInstanceHealth"

// DescribeInstanceHealthRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DescribeInstanceHealthRequest.
//    req := client.DescribeInstanceHealthRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/DescribeInstanceHealth.htm
func (c *Client) DescribeInstanceHealthRequest(input *DescribeInstanceHealthInput) DescribeInstanceHealthRequest {
	op := &aws.Operation{
		Name:       opDescribeInstanceHealth,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DescribeInstanceHealthInput{}
	}

	req := c.newRequest(op, input, &DescribeInstanceHealthOutput{})

	return DescribeInstanceHealthRequest{Request: req, Input: input, Copy: c.DescribeInstanceHealthRequest}
}

// DescribeInstanceHealthRequest is the request type for the
// DescribeInstanceHealth API operation.
type DescribeInstanceHealthRequest struct {
	*aws.Request
	Input *DescribeInstanceHealthInput
	Copy  func(*DescribeInstanceHealthInput) DescribeInstanceHealthRequest
}

// Send marshals and sends the DescribeInstanceHealth API request.
func (r DescribeInstanceHealthRequest) Send(ctx context.Context) (*DescribeInstanceHealthResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeInstanceHealthResponse{
		DescribeInstanceHealthOutput: r.Request.Data.(*DescribeInstanceHealthOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeInstanceHealthResponse is the response type for the
// DescribeInstanceHealth API operation.
type DescribeInstanceHealthResponse struct {
	*DescribeInstanceHealthOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeInstanceHealth request.
func (r *DescribeInstanceHealthResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
