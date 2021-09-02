// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ess

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type GetSendQuotaInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetSendQuotaInput) String() string {
	return nifcloudutil.Prettify(s)
}

type GetSendQuotaOutput struct {
	_ struct{} `type:"structure"`

	Max24HourSend *float64 `locationName:"Max24HourSend" type:"double"`

	MaxSendRate *float64 `locationName:"MaxSendRate" type:"double"`

	ResponseMetadata *ResponseMetadata `locationName:"ResponseMetadata" type:"structure"`

	SentLast24Hours *float64 `locationName:"SentLast24Hours" type:"double"`
}

// String returns the string representation
func (s GetSendQuotaOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opGetSendQuota = "GetSendQuota"

// GetSendQuotaRequest returns a request value for making API operation for
// NIFCLOUD ESS.
//
//    // Example sending a request using GetSendQuotaRequest.
//    req := client.GetSendQuotaRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/ess/GetSendQuota.htm
func (c *Client) GetSendQuotaRequest(input *GetSendQuotaInput) GetSendQuotaRequest {
	op := &aws.Operation{
		Name:       opGetSendQuota,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetSendQuotaInput{}
	}

	req := c.newRequest(op, input, &GetSendQuotaOutput{})

	return GetSendQuotaRequest{Request: req, Input: input, Copy: c.GetSendQuotaRequest}
}

// GetSendQuotaRequest is the request type for the
// GetSendQuota API operation.
type GetSendQuotaRequest struct {
	*aws.Request
	Input *GetSendQuotaInput
	Copy  func(*GetSendQuotaInput) GetSendQuotaRequest
}

// Send marshals and sends the GetSendQuota API request.
func (r GetSendQuotaRequest) Send(ctx context.Context) (*GetSendQuotaResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetSendQuotaResponse{
		GetSendQuotaOutput: r.Request.Data.(*GetSendQuotaOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetSendQuotaResponse is the response type for the
// GetSendQuota API operation.
type GetSendQuotaResponse struct {
	*GetSendQuotaOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetSendQuota request.
func (r *GetSendQuotaResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
