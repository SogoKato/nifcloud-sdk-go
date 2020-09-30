// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type NiftyRegisterPortWithElasticLoadBalancerInput struct {
	_ struct{} `type:"structure"`

	ElasticLoadBalancerId *string `locationName:"ElasticLoadBalancerId" type:"string"`

	ElasticLoadBalancerName *string `locationName:"ElasticLoadBalancerName" type:"string"`

	Listeners []RequestListenersOfNiftyRegisterPortWithElasticLoadBalancer `locationName:"Listeners" locationNameList:"member" type:"list"`
}

// String returns the string representation
func (s NiftyRegisterPortWithElasticLoadBalancerInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyRegisterPortWithElasticLoadBalancerOutput struct {
	_ struct{} `type:"structure"`

	Listeners []ListenersOfNiftyRegisterPortWithElasticLoadBalancer `locationName:"Listeners" locationNameList:"member" type:"list"`

	NiftyRegisterPortWithElasticLoadBalancerResult *NiftyRegisterPortWithElasticLoadBalancerResult `locationName:"NiftyRegisterPortWithElasticLoadBalancerResult" type:"structure"`

	ResponseMetadata *ResponseMetadataOfNiftyRegisterPortWithElasticLoadBalancer `locationName:"ResponseMetadata" type:"structure"`
}

// String returns the string representation
func (s NiftyRegisterPortWithElasticLoadBalancerOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyRegisterPortWithElasticLoadBalancer = "NiftyRegisterPortWithElasticLoadBalancer"

// NiftyRegisterPortWithElasticLoadBalancerRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyRegisterPortWithElasticLoadBalancerRequest.
//    req := client.NiftyRegisterPortWithElasticLoadBalancerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/NiftyRegisterPortWithElasticLoadBalancer.htm
func (c *Client) NiftyRegisterPortWithElasticLoadBalancerRequest(input *NiftyRegisterPortWithElasticLoadBalancerInput) NiftyRegisterPortWithElasticLoadBalancerRequest {
	op := &aws.Operation{
		Name:       opNiftyRegisterPortWithElasticLoadBalancer,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyRegisterPortWithElasticLoadBalancerInput{}
	}

	req := c.newRequest(op, input, &NiftyRegisterPortWithElasticLoadBalancerOutput{})

	return NiftyRegisterPortWithElasticLoadBalancerRequest{Request: req, Input: input, Copy: c.NiftyRegisterPortWithElasticLoadBalancerRequest}
}

// NiftyRegisterPortWithElasticLoadBalancerRequest is the request type for the
// NiftyRegisterPortWithElasticLoadBalancer API operation.
type NiftyRegisterPortWithElasticLoadBalancerRequest struct {
	*aws.Request
	Input *NiftyRegisterPortWithElasticLoadBalancerInput
	Copy  func(*NiftyRegisterPortWithElasticLoadBalancerInput) NiftyRegisterPortWithElasticLoadBalancerRequest
}

// Send marshals and sends the NiftyRegisterPortWithElasticLoadBalancer API request.
func (r NiftyRegisterPortWithElasticLoadBalancerRequest) Send(ctx context.Context) (*NiftyRegisterPortWithElasticLoadBalancerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyRegisterPortWithElasticLoadBalancerResponse{
		NiftyRegisterPortWithElasticLoadBalancerOutput: r.Request.Data.(*NiftyRegisterPortWithElasticLoadBalancerOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyRegisterPortWithElasticLoadBalancerResponse is the response type for the
// NiftyRegisterPortWithElasticLoadBalancer API operation.
type NiftyRegisterPortWithElasticLoadBalancerResponse struct {
	*NiftyRegisterPortWithElasticLoadBalancerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyRegisterPortWithElasticLoadBalancer request.
func (r *NiftyRegisterPortWithElasticLoadBalancerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
