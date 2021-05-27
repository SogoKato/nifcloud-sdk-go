// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type NiftyCreateWebProxyInput struct {
	_ struct{} `type:"structure"`

	Agreement *bool `locationName:"Agreement" type:"boolean"`

	BypassInterface *RequestBypassInterface `locationName:"BypassInterface" type:"structure"`

	Description *string `locationName:"Description" type:"string"`

	ListenInterface *RequestListenInterface `locationName:"ListenInterface" type:"structure"`

	// ListenPort is a required field
	ListenPort *string `locationName:"ListenPort" type:"string" required:"true"`

	Option *RequestOption `locationName:"Option" type:"structure"`

	RouterId *string `locationName:"RouterId" type:"string"`

	RouterName *string `locationName:"RouterName" type:"string"`
}

// String returns the string representation
func (s NiftyCreateWebProxyInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *NiftyCreateWebProxyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "NiftyCreateWebProxyInput"}

	if s.ListenPort == nil {
		invalidParams.Add(aws.NewErrParamRequired("ListenPort"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type NiftyCreateWebProxyOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	WebProxy *WebProxy `locationName:"webProxy" type:"structure"`
}

// String returns the string representation
func (s NiftyCreateWebProxyOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyCreateWebProxy = "NiftyCreateWebProxy"

// NiftyCreateWebProxyRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyCreateWebProxyRequest.
//    req := client.NiftyCreateWebProxyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/NiftyCreateWebProxy.htm
func (c *Client) NiftyCreateWebProxyRequest(input *NiftyCreateWebProxyInput) NiftyCreateWebProxyRequest {
	op := &aws.Operation{
		Name:       opNiftyCreateWebProxy,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyCreateWebProxyInput{}
	}

	req := c.newRequest(op, input, &NiftyCreateWebProxyOutput{})

	return NiftyCreateWebProxyRequest{Request: req, Input: input, Copy: c.NiftyCreateWebProxyRequest}
}

// NiftyCreateWebProxyRequest is the request type for the
// NiftyCreateWebProxy API operation.
type NiftyCreateWebProxyRequest struct {
	*aws.Request
	Input *NiftyCreateWebProxyInput
	Copy  func(*NiftyCreateWebProxyInput) NiftyCreateWebProxyRequest
}

// Send marshals and sends the NiftyCreateWebProxy API request.
func (r NiftyCreateWebProxyRequest) Send(ctx context.Context) (*NiftyCreateWebProxyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyCreateWebProxyResponse{
		NiftyCreateWebProxyOutput: r.Request.Data.(*NiftyCreateWebProxyOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyCreateWebProxyResponse is the response type for the
// NiftyCreateWebProxy API operation.
type NiftyCreateWebProxyResponse struct {
	*NiftyCreateWebProxyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyCreateWebProxy request.
func (r *NiftyCreateWebProxyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
