// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type CreateDhcpOptionsInput struct {
	_ struct{} `type:"structure"`

	DhcpConfiguration []RequestDhcpConfiguration `locationName:"DhcpConfiguration" type:"list"`
}

// String returns the string representation
func (s CreateDhcpOptionsInput) String() string {
	return nifcloudutil.Prettify(s)
}

type CreateDhcpOptionsOutput struct {
	_ struct{} `type:"structure"`

	DhcpOptions *DhcpOptions `locationName:"dhcpOptions" type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s CreateDhcpOptionsOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opCreateDhcpOptions = "CreateDhcpOptions"

// CreateDhcpOptionsRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using CreateDhcpOptionsRequest.
//    req := client.CreateDhcpOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/CreateDhcpOptions.htm
func (c *Client) CreateDhcpOptionsRequest(input *CreateDhcpOptionsInput) CreateDhcpOptionsRequest {
	op := &aws.Operation{
		Name:       opCreateDhcpOptions,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &CreateDhcpOptionsInput{}
	}

	req := c.newRequest(op, input, &CreateDhcpOptionsOutput{})

	return CreateDhcpOptionsRequest{Request: req, Input: input, Copy: c.CreateDhcpOptionsRequest}
}

// CreateDhcpOptionsRequest is the request type for the
// CreateDhcpOptions API operation.
type CreateDhcpOptionsRequest struct {
	*aws.Request
	Input *CreateDhcpOptionsInput
	Copy  func(*CreateDhcpOptionsInput) CreateDhcpOptionsRequest
}

// Send marshals and sends the CreateDhcpOptions API request.
func (r CreateDhcpOptionsRequest) Send(ctx context.Context) (*CreateDhcpOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDhcpOptionsResponse{
		CreateDhcpOptionsOutput: r.Request.Data.(*CreateDhcpOptionsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDhcpOptionsResponse is the response type for the
// CreateDhcpOptions API operation.
type CreateDhcpOptionsResponse struct {
	*CreateDhcpOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDhcpOptions request.
func (r *CreateDhcpOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
