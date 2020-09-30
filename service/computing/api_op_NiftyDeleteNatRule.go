// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type NiftyDeleteNatRuleInput struct {
	_ struct{} `type:"structure"`

	NatTableId *string `locationName:"NatTableId" type:"string"`

	NatType NatTypeOfNiftyDeleteNatRuleRequest `locationName:"NatType" type:"string" enum:"true"`

	RuleNumber *string `locationName:"RuleNumber" type:"string"`
}

// String returns the string representation
func (s NiftyDeleteNatRuleInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyDeleteNatRuleOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s NiftyDeleteNatRuleOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyDeleteNatRule = "NiftyDeleteNatRule"

// NiftyDeleteNatRuleRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyDeleteNatRuleRequest.
//    req := client.NiftyDeleteNatRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/NiftyDeleteNatRule.htm
func (c *Client) NiftyDeleteNatRuleRequest(input *NiftyDeleteNatRuleInput) NiftyDeleteNatRuleRequest {
	op := &aws.Operation{
		Name:       opNiftyDeleteNatRule,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyDeleteNatRuleInput{}
	}

	req := c.newRequest(op, input, &NiftyDeleteNatRuleOutput{})

	return NiftyDeleteNatRuleRequest{Request: req, Input: input, Copy: c.NiftyDeleteNatRuleRequest}
}

// NiftyDeleteNatRuleRequest is the request type for the
// NiftyDeleteNatRule API operation.
type NiftyDeleteNatRuleRequest struct {
	*aws.Request
	Input *NiftyDeleteNatRuleInput
	Copy  func(*NiftyDeleteNatRuleInput) NiftyDeleteNatRuleRequest
}

// Send marshals and sends the NiftyDeleteNatRule API request.
func (r NiftyDeleteNatRuleRequest) Send(ctx context.Context) (*NiftyDeleteNatRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyDeleteNatRuleResponse{
		NiftyDeleteNatRuleOutput: r.Request.Data.(*NiftyDeleteNatRuleOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyDeleteNatRuleResponse is the response type for the
// NiftyDeleteNatRule API operation.
type NiftyDeleteNatRuleResponse struct {
	*NiftyDeleteNatRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyDeleteNatRule request.
func (r *NiftyDeleteNatRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
