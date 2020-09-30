// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type DescribeLoadBalancersInput struct {
	_ struct{} `type:"structure"`

	LoadBalancerNames []RequestLoadBalancerNames `locationName:"LoadBalancerNames" locationNameList:"member" type:"list"`

	Owner OwnerOfDescribeLoadBalancersRequest `locationName:"Owner" type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeLoadBalancersInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DescribeLoadBalancersOutput struct {
	_ struct{} `type:"structure"`

	DescribeLoadBalancersResult *DescribeLoadBalancersResult `locationName:"DescribeLoadBalancersResult" type:"structure"`

	LoadBalancerDescriptions []LoadBalancerDescriptions `locationName:"LoadBalancerDescriptions" locationNameList:"member" type:"list"`

	ResponseMetadata *ResponseMetadataOfDescribeLoadBalancers `locationName:"ResponseMetadata" type:"structure"`
}

// String returns the string representation
func (s DescribeLoadBalancersOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDescribeLoadBalancers = "DescribeLoadBalancers"

// DescribeLoadBalancersRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DescribeLoadBalancersRequest.
//    req := client.DescribeLoadBalancersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/DescribeLoadBalancers.htm
func (c *Client) DescribeLoadBalancersRequest(input *DescribeLoadBalancersInput) DescribeLoadBalancersRequest {
	op := &aws.Operation{
		Name:       opDescribeLoadBalancers,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DescribeLoadBalancersInput{}
	}

	req := c.newRequest(op, input, &DescribeLoadBalancersOutput{})

	return DescribeLoadBalancersRequest{Request: req, Input: input, Copy: c.DescribeLoadBalancersRequest}
}

// DescribeLoadBalancersRequest is the request type for the
// DescribeLoadBalancers API operation.
type DescribeLoadBalancersRequest struct {
	*aws.Request
	Input *DescribeLoadBalancersInput
	Copy  func(*DescribeLoadBalancersInput) DescribeLoadBalancersRequest
}

// Send marshals and sends the DescribeLoadBalancers API request.
func (r DescribeLoadBalancersRequest) Send(ctx context.Context) (*DescribeLoadBalancersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeLoadBalancersResponse{
		DescribeLoadBalancersOutput: r.Request.Data.(*DescribeLoadBalancersOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeLoadBalancersResponse is the response type for the
// DescribeLoadBalancers API operation.
type DescribeLoadBalancersResponse struct {
	*DescribeLoadBalancersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeLoadBalancers request.
func (r *DescribeLoadBalancersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
