// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type DetachIsoImageInput struct {
	_ struct{} `type:"structure"`

	// InstanceUniqueId is a required field
	InstanceUniqueId *string `locationName:"InstanceUniqueId" type:"string" required:"true"`

	// IsoImageId is a required field
	IsoImageId *string `locationName:"IsoImageId" type:"string" required:"true"`
}

// String returns the string representation
func (s DetachIsoImageInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachIsoImageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetachIsoImageInput"}

	if s.InstanceUniqueId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceUniqueId"))
	}

	if s.IsoImageId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IsoImageId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DetachIsoImageOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s DetachIsoImageOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDetachIsoImage = "DetachIsoImage"

// DetachIsoImageRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DetachIsoImageRequest.
//    req := client.DetachIsoImageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/DetachIsoImage.htm
func (c *Client) DetachIsoImageRequest(input *DetachIsoImageInput) DetachIsoImageRequest {
	op := &aws.Operation{
		Name:       opDetachIsoImage,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DetachIsoImageInput{}
	}

	req := c.newRequest(op, input, &DetachIsoImageOutput{})

	return DetachIsoImageRequest{Request: req, Input: input, Copy: c.DetachIsoImageRequest}
}

// DetachIsoImageRequest is the request type for the
// DetachIsoImage API operation.
type DetachIsoImageRequest struct {
	*aws.Request
	Input *DetachIsoImageInput
	Copy  func(*DetachIsoImageInput) DetachIsoImageRequest
}

// Send marshals and sends the DetachIsoImage API request.
func (r DetachIsoImageRequest) Send(ctx context.Context) (*DetachIsoImageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetachIsoImageResponse{
		DetachIsoImageOutput: r.Request.Data.(*DetachIsoImageOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetachIsoImageResponse is the response type for the
// DetachIsoImage API operation.
type DetachIsoImageResponse struct {
	*DetachIsoImageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetachIsoImage request.
func (r *DetachIsoImageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
