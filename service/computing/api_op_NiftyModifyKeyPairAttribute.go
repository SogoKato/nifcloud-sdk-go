// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type NiftyModifyKeyPairAttributeInput struct {
	_ struct{} `type:"structure"`

	Attribute AttributeOfNiftyModifyKeyPairAttributeRequest `locationName:"Attribute" type:"string" enum:"true"`

	KeyName *string `locationName:"KeyName" type:"string"`

	Value *string `locationName:"Value" type:"string"`
}

// String returns the string representation
func (s NiftyModifyKeyPairAttributeInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyModifyKeyPairAttributeOutput struct {
	_ struct{} `type:"structure"`

	Attribute *string `locationName:"attribute" type:"string"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`

	Value *string `locationName:"value" type:"string"`
}

// String returns the string representation
func (s NiftyModifyKeyPairAttributeOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyModifyKeyPairAttribute = "NiftyModifyKeyPairAttribute"

// NiftyModifyKeyPairAttributeRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyModifyKeyPairAttributeRequest.
//    req := client.NiftyModifyKeyPairAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/NiftyModifyKeyPairAttribute.htm
func (c *Client) NiftyModifyKeyPairAttributeRequest(input *NiftyModifyKeyPairAttributeInput) NiftyModifyKeyPairAttributeRequest {
	op := &aws.Operation{
		Name:       opNiftyModifyKeyPairAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyModifyKeyPairAttributeInput{}
	}

	req := c.newRequest(op, input, &NiftyModifyKeyPairAttributeOutput{})

	return NiftyModifyKeyPairAttributeRequest{Request: req, Input: input, Copy: c.NiftyModifyKeyPairAttributeRequest}
}

// NiftyModifyKeyPairAttributeRequest is the request type for the
// NiftyModifyKeyPairAttribute API operation.
type NiftyModifyKeyPairAttributeRequest struct {
	*aws.Request
	Input *NiftyModifyKeyPairAttributeInput
	Copy  func(*NiftyModifyKeyPairAttributeInput) NiftyModifyKeyPairAttributeRequest
}

// Send marshals and sends the NiftyModifyKeyPairAttribute API request.
func (r NiftyModifyKeyPairAttributeRequest) Send(ctx context.Context) (*NiftyModifyKeyPairAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyModifyKeyPairAttributeResponse{
		NiftyModifyKeyPairAttributeOutput: r.Request.Data.(*NiftyModifyKeyPairAttributeOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyModifyKeyPairAttributeResponse is the response type for the
// NiftyModifyKeyPairAttribute API operation.
type NiftyModifyKeyPairAttributeResponse struct {
	*NiftyModifyKeyPairAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyModifyKeyPairAttribute request.
func (r *NiftyModifyKeyPairAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
