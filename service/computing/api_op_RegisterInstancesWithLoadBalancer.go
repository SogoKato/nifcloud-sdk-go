// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type RegisterInstancesWithLoadBalancerInput struct {
	_ struct{} `type:"structure"`

	// InstancePort is a required field
	InstancePort *int64 `locationName:"InstancePort" type:"integer" required:"true"`

	// Instances is a required field
	Instances []RequestInstances `locationName:"Instances" locationNameList:"member" type:"list" required:"true"`

	// LoadBalancerName is a required field
	LoadBalancerName *string `locationName:"LoadBalancerName" type:"string" required:"true"`

	// LoadBalancerPort is a required field
	LoadBalancerPort *int64 `locationName:"LoadBalancerPort" type:"integer" required:"true"`
}

// String returns the string representation
func (s RegisterInstancesWithLoadBalancerInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterInstancesWithLoadBalancerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterInstancesWithLoadBalancerInput"}

	if s.InstancePort == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstancePort"))
	}

	if s.Instances == nil {
		invalidParams.Add(aws.NewErrParamRequired("Instances"))
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

type RegisterInstancesWithLoadBalancerOutput struct {
	_ struct{} `type:"structure"`

	Instances []Instances `locationName:"Instances" locationNameList:"member" type:"list"`

	ResponseMetadata *ResponseMetadata `locationName:"ResponseMetadata" type:"structure"`
}

// String returns the string representation
func (s RegisterInstancesWithLoadBalancerOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opRegisterInstancesWithLoadBalancer = "RegisterInstancesWithLoadBalancer"

// RegisterInstancesWithLoadBalancerRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using RegisterInstancesWithLoadBalancerRequest.
//    req := client.RegisterInstancesWithLoadBalancerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/RegisterInstancesWithLoadBalancer.htm
func (c *Client) RegisterInstancesWithLoadBalancerRequest(input *RegisterInstancesWithLoadBalancerInput) RegisterInstancesWithLoadBalancerRequest {
	op := &aws.Operation{
		Name:       opRegisterInstancesWithLoadBalancer,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &RegisterInstancesWithLoadBalancerInput{}
	}

	req := c.newRequest(op, input, &RegisterInstancesWithLoadBalancerOutput{})

	return RegisterInstancesWithLoadBalancerRequest{Request: req, Input: input, Copy: c.RegisterInstancesWithLoadBalancerRequest}
}

// RegisterInstancesWithLoadBalancerRequest is the request type for the
// RegisterInstancesWithLoadBalancer API operation.
type RegisterInstancesWithLoadBalancerRequest struct {
	*aws.Request
	Input *RegisterInstancesWithLoadBalancerInput
	Copy  func(*RegisterInstancesWithLoadBalancerInput) RegisterInstancesWithLoadBalancerRequest
}

// Send marshals and sends the RegisterInstancesWithLoadBalancer API request.
func (r RegisterInstancesWithLoadBalancerRequest) Send(ctx context.Context) (*RegisterInstancesWithLoadBalancerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RegisterInstancesWithLoadBalancerResponse{
		RegisterInstancesWithLoadBalancerOutput: r.Request.Data.(*RegisterInstancesWithLoadBalancerOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RegisterInstancesWithLoadBalancerResponse is the response type for the
// RegisterInstancesWithLoadBalancer API operation.
type RegisterInstancesWithLoadBalancerResponse struct {
	*RegisterInstancesWithLoadBalancerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RegisterInstancesWithLoadBalancer request.
func (r *RegisterInstancesWithLoadBalancerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
