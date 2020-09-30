// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type ReleaseMultiIpAddressesInput struct {
	_ struct{} `type:"structure"`

	IpAddress []string `locationName:"IpAddress" type:"list"`

	MultiIpAddressGroupId *string `locationName:"MultiIpAddressGroupId" type:"string"`
}

// String returns the string representation
func (s ReleaseMultiIpAddressesInput) String() string {
	return nifcloudutil.Prettify(s)
}

type ReleaseMultiIpAddressesOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s ReleaseMultiIpAddressesOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opReleaseMultiIpAddresses = "ReleaseMultiIpAddresses"

// ReleaseMultiIpAddressesRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using ReleaseMultiIpAddressesRequest.
//    req := client.ReleaseMultiIpAddressesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/ReleaseMultiIpAddresses.htm
func (c *Client) ReleaseMultiIpAddressesRequest(input *ReleaseMultiIpAddressesInput) ReleaseMultiIpAddressesRequest {
	op := &aws.Operation{
		Name:       opReleaseMultiIpAddresses,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &ReleaseMultiIpAddressesInput{}
	}

	req := c.newRequest(op, input, &ReleaseMultiIpAddressesOutput{})

	return ReleaseMultiIpAddressesRequest{Request: req, Input: input, Copy: c.ReleaseMultiIpAddressesRequest}
}

// ReleaseMultiIpAddressesRequest is the request type for the
// ReleaseMultiIpAddresses API operation.
type ReleaseMultiIpAddressesRequest struct {
	*aws.Request
	Input *ReleaseMultiIpAddressesInput
	Copy  func(*ReleaseMultiIpAddressesInput) ReleaseMultiIpAddressesRequest
}

// Send marshals and sends the ReleaseMultiIpAddresses API request.
func (r ReleaseMultiIpAddressesRequest) Send(ctx context.Context) (*ReleaseMultiIpAddressesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ReleaseMultiIpAddressesResponse{
		ReleaseMultiIpAddressesOutput: r.Request.Data.(*ReleaseMultiIpAddressesOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ReleaseMultiIpAddressesResponse is the response type for the
// ReleaseMultiIpAddresses API operation.
type ReleaseMultiIpAddressesResponse struct {
	*ReleaseMultiIpAddressesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ReleaseMultiIpAddresses request.
func (r *ReleaseMultiIpAddressesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
