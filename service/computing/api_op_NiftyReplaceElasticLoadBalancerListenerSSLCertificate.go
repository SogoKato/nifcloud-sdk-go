// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type NiftyReplaceElasticLoadBalancerListenerSSLCertificateInput struct {
	_ struct{} `type:"structure"`

	ElasticLoadBalancerId *string `locationName:"ElasticLoadBalancerId" type:"string"`

	ElasticLoadBalancerName *string `locationName:"ElasticLoadBalancerName" type:"string"`

	// ElasticLoadBalancerPort is a required field
	ElasticLoadBalancerPort *int64 `locationName:"ElasticLoadBalancerPort" type:"integer" required:"true"`

	// InstancePort is a required field
	InstancePort *int64 `locationName:"InstancePort" type:"integer" required:"true"`

	// Protocol is a required field
	Protocol ProtocolOfNiftyReplaceElasticLoadBalancerListenerSSLCertificateRequest `locationName:"Protocol" type:"string" required:"true" enum:"true"`

	// SSLCertificateId is a required field
	SSLCertificateId *string `locationName:"SSLCertificateId" type:"string" required:"true"`
}

// String returns the string representation
func (s NiftyReplaceElasticLoadBalancerListenerSSLCertificateInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *NiftyReplaceElasticLoadBalancerListenerSSLCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "NiftyReplaceElasticLoadBalancerListenerSSLCertificateInput"}

	if s.ElasticLoadBalancerPort == nil {
		invalidParams.Add(aws.NewErrParamRequired("ElasticLoadBalancerPort"))
	}

	if s.InstancePort == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstancePort"))
	}
	if len(s.Protocol) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Protocol"))
	}

	if s.SSLCertificateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SSLCertificateId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type NiftyReplaceElasticLoadBalancerListenerSSLCertificateOutput struct {
	_ struct{} `type:"structure"`

	NiftyReplaceElasticLoadBalancerListenerSSLCertificateResult *string `locationName:"NiftyReplaceElasticLoadBalancerListenerSSLCertificateResult" type:"string"`

	ResponseMetadata *ResponseMetadata `locationName:"ResponseMetadata" type:"structure"`
}

// String returns the string representation
func (s NiftyReplaceElasticLoadBalancerListenerSSLCertificateOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyReplaceElasticLoadBalancerListenerSSLCertificate = "NiftyReplaceElasticLoadBalancerListenerSSLCertificate"

// NiftyReplaceElasticLoadBalancerListenerSSLCertificateRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyReplaceElasticLoadBalancerListenerSSLCertificateRequest.
//    req := client.NiftyReplaceElasticLoadBalancerListenerSSLCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/NiftyReplaceElasticLoadBalancerListenerSSLCertificate.htm
func (c *Client) NiftyReplaceElasticLoadBalancerListenerSSLCertificateRequest(input *NiftyReplaceElasticLoadBalancerListenerSSLCertificateInput) NiftyReplaceElasticLoadBalancerListenerSSLCertificateRequest {
	op := &aws.Operation{
		Name:       opNiftyReplaceElasticLoadBalancerListenerSSLCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyReplaceElasticLoadBalancerListenerSSLCertificateInput{}
	}

	req := c.newRequest(op, input, &NiftyReplaceElasticLoadBalancerListenerSSLCertificateOutput{})

	return NiftyReplaceElasticLoadBalancerListenerSSLCertificateRequest{Request: req, Input: input, Copy: c.NiftyReplaceElasticLoadBalancerListenerSSLCertificateRequest}
}

// NiftyReplaceElasticLoadBalancerListenerSSLCertificateRequest is the request type for the
// NiftyReplaceElasticLoadBalancerListenerSSLCertificate API operation.
type NiftyReplaceElasticLoadBalancerListenerSSLCertificateRequest struct {
	*aws.Request
	Input *NiftyReplaceElasticLoadBalancerListenerSSLCertificateInput
	Copy  func(*NiftyReplaceElasticLoadBalancerListenerSSLCertificateInput) NiftyReplaceElasticLoadBalancerListenerSSLCertificateRequest
}

// Send marshals and sends the NiftyReplaceElasticLoadBalancerListenerSSLCertificate API request.
func (r NiftyReplaceElasticLoadBalancerListenerSSLCertificateRequest) Send(ctx context.Context) (*NiftyReplaceElasticLoadBalancerListenerSSLCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyReplaceElasticLoadBalancerListenerSSLCertificateResponse{
		NiftyReplaceElasticLoadBalancerListenerSSLCertificateOutput: r.Request.Data.(*NiftyReplaceElasticLoadBalancerListenerSSLCertificateOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyReplaceElasticLoadBalancerListenerSSLCertificateResponse is the response type for the
// NiftyReplaceElasticLoadBalancerListenerSSLCertificate API operation.
type NiftyReplaceElasticLoadBalancerListenerSSLCertificateResponse struct {
	*NiftyReplaceElasticLoadBalancerListenerSSLCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyReplaceElasticLoadBalancerListenerSSLCertificate request.
func (r *NiftyReplaceElasticLoadBalancerListenerSSLCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
