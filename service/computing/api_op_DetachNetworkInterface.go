// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type DetachNetworkInterfaceInput struct {
	_ struct{} `type:"structure"`

	// AttachmentId is a required field
	AttachmentId *string `locationName:"AttachmentId" type:"string" required:"true"`

	NiftyReboot NiftyRebootOfDetachNetworkInterfaceRequest `locationName:"NiftyReboot" type:"string" enum:"true"`
}

// String returns the string representation
func (s DetachNetworkInterfaceInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachNetworkInterfaceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DetachNetworkInterfaceInput"}

	if s.AttachmentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AttachmentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DetachNetworkInterfaceOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s DetachNetworkInterfaceOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDetachNetworkInterface = "DetachNetworkInterface"

// DetachNetworkInterfaceRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DetachNetworkInterfaceRequest.
//    req := client.DetachNetworkInterfaceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/DetachNetworkInterface.htm
func (c *Client) DetachNetworkInterfaceRequest(input *DetachNetworkInterfaceInput) DetachNetworkInterfaceRequest {
	op := &aws.Operation{
		Name:       opDetachNetworkInterface,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DetachNetworkInterfaceInput{}
	}

	req := c.newRequest(op, input, &DetachNetworkInterfaceOutput{})

	return DetachNetworkInterfaceRequest{Request: req, Input: input, Copy: c.DetachNetworkInterfaceRequest}
}

// DetachNetworkInterfaceRequest is the request type for the
// DetachNetworkInterface API operation.
type DetachNetworkInterfaceRequest struct {
	*aws.Request
	Input *DetachNetworkInterfaceInput
	Copy  func(*DetachNetworkInterfaceInput) DetachNetworkInterfaceRequest
}

// Send marshals and sends the DetachNetworkInterface API request.
func (r DetachNetworkInterfaceRequest) Send(ctx context.Context) (*DetachNetworkInterfaceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DetachNetworkInterfaceResponse{
		DetachNetworkInterfaceOutput: r.Request.Data.(*DetachNetworkInterfaceOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DetachNetworkInterfaceResponse is the response type for the
// DetachNetworkInterface API operation.
type DetachNetworkInterfaceResponse struct {
	*DetachNetworkInterfaceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DetachNetworkInterface request.
func (r *DetachNetworkInterfaceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
