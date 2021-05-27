// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type DeleteImageInput struct {
	_ struct{} `type:"structure"`

	// ImageId is a required field
	ImageId *string `locationName:"ImageId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteImageInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteImageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteImageInput"}

	if s.ImageId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ImageId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteImageOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s DeleteImageOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDeleteImage = "DeleteImage"

// DeleteImageRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DeleteImageRequest.
//    req := client.DeleteImageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/DeleteImage.htm
func (c *Client) DeleteImageRequest(input *DeleteImageInput) DeleteImageRequest {
	op := &aws.Operation{
		Name:       opDeleteImage,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DeleteImageInput{}
	}

	req := c.newRequest(op, input, &DeleteImageOutput{})

	return DeleteImageRequest{Request: req, Input: input, Copy: c.DeleteImageRequest}
}

// DeleteImageRequest is the request type for the
// DeleteImage API operation.
type DeleteImageRequest struct {
	*aws.Request
	Input *DeleteImageInput
	Copy  func(*DeleteImageInput) DeleteImageRequest
}

// Send marshals and sends the DeleteImage API request.
func (r DeleteImageRequest) Send(ctx context.Context) (*DeleteImageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteImageResponse{
		DeleteImageOutput: r.Request.Data.(*DeleteImageOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteImageResponse is the response type for the
// DeleteImage API operation.
type DeleteImageResponse struct {
	*DeleteImageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteImage request.
func (r *DeleteImageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
