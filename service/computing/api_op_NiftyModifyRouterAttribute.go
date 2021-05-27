// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type NiftyModifyRouterAttributeInput struct {
	_ struct{} `type:"structure"`

	Agreement *bool `locationName:"Agreement" type:"boolean"`

	// Attribute is a required field
	Attribute AttributeOfNiftyModifyRouterAttributeRequest `locationName:"Attribute" type:"string" required:"true" enum:"true"`

	RouterId *string `locationName:"RouterId" type:"string"`

	RouterName *string `locationName:"RouterName" type:"string"`

	// Value is a required field
	Value *string `locationName:"Value" type:"string" required:"true"`
}

// String returns the string representation
func (s NiftyModifyRouterAttributeInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *NiftyModifyRouterAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "NiftyModifyRouterAttributeInput"}
	if len(s.Attribute) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Attribute"))
	}

	if s.Value == nil {
		invalidParams.Add(aws.NewErrParamRequired("Value"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type NiftyModifyRouterAttributeOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s NiftyModifyRouterAttributeOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyModifyRouterAttribute = "NiftyModifyRouterAttribute"

// NiftyModifyRouterAttributeRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyModifyRouterAttributeRequest.
//    req := client.NiftyModifyRouterAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/NiftyModifyRouterAttribute.htm
func (c *Client) NiftyModifyRouterAttributeRequest(input *NiftyModifyRouterAttributeInput) NiftyModifyRouterAttributeRequest {
	op := &aws.Operation{
		Name:       opNiftyModifyRouterAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyModifyRouterAttributeInput{}
	}

	req := c.newRequest(op, input, &NiftyModifyRouterAttributeOutput{})

	return NiftyModifyRouterAttributeRequest{Request: req, Input: input, Copy: c.NiftyModifyRouterAttributeRequest}
}

// NiftyModifyRouterAttributeRequest is the request type for the
// NiftyModifyRouterAttribute API operation.
type NiftyModifyRouterAttributeRequest struct {
	*aws.Request
	Input *NiftyModifyRouterAttributeInput
	Copy  func(*NiftyModifyRouterAttributeInput) NiftyModifyRouterAttributeRequest
}

// Send marshals and sends the NiftyModifyRouterAttribute API request.
func (r NiftyModifyRouterAttributeRequest) Send(ctx context.Context) (*NiftyModifyRouterAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyModifyRouterAttributeResponse{
		NiftyModifyRouterAttributeOutput: r.Request.Data.(*NiftyModifyRouterAttributeOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyModifyRouterAttributeResponse is the response type for the
// NiftyModifyRouterAttribute API operation.
type NiftyModifyRouterAttributeResponse struct {
	*NiftyModifyRouterAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyModifyRouterAttribute request.
func (r *NiftyModifyRouterAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
