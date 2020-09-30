// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type CreateSslCertificateInput struct {
	_ struct{} `type:"structure"`

	ApproverEmailAddress *string `locationName:"ApproverEmailAddress" type:"string"`

	CertAuthority *int64 `locationName:"CertAuthority" type:"integer"`

	CertInfo *RequestCertInfo `locationName:"CertInfo" type:"structure"`

	Count *int64 `locationName:"Count" type:"integer"`

	Fqdn *string `locationName:"Fqdn" type:"string"`

	FqdnId *string `locationName:"FqdnId" type:"string"`

	KeyLength *int64 `locationName:"KeyLength" type:"integer"`

	ValidityTerm *int64 `locationName:"ValidityTerm" type:"integer"`
}

// String returns the string representation
func (s CreateSslCertificateInput) String() string {
	return nifcloudutil.Prettify(s)
}

type CreateSslCertificateOutput struct {
	_ struct{} `type:"structure"`

	ApproverEmailAddress *string `locationName:"approverEmailAddress" type:"string"`

	CertAuthority *string `locationName:"certAuthority" type:"string"`

	CertState *string `locationName:"certState" type:"string"`

	Fqdn *string `locationName:"fqdn" type:"string"`

	FqdnId *string `locationName:"fqdnId" type:"string"`

	RequestId *string `locationName:"requestId" type:"string"`

	ValidityTerm *int64 `locationName:"validityTerm" type:"integer"`
}

// String returns the string representation
func (s CreateSslCertificateOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opCreateSslCertificate = "CreateSslCertificate"

// CreateSslCertificateRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using CreateSslCertificateRequest.
//    req := client.CreateSslCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/CreateSslCertificate.htm
func (c *Client) CreateSslCertificateRequest(input *CreateSslCertificateInput) CreateSslCertificateRequest {
	op := &aws.Operation{
		Name:       opCreateSslCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &CreateSslCertificateInput{}
	}

	req := c.newRequest(op, input, &CreateSslCertificateOutput{})

	return CreateSslCertificateRequest{Request: req, Input: input, Copy: c.CreateSslCertificateRequest}
}

// CreateSslCertificateRequest is the request type for the
// CreateSslCertificate API operation.
type CreateSslCertificateRequest struct {
	*aws.Request
	Input *CreateSslCertificateInput
	Copy  func(*CreateSslCertificateInput) CreateSslCertificateRequest
}

// Send marshals and sends the CreateSslCertificate API request.
func (r CreateSslCertificateRequest) Send(ctx context.Context) (*CreateSslCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSslCertificateResponse{
		CreateSslCertificateOutput: r.Request.Data.(*CreateSslCertificateOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSslCertificateResponse is the response type for the
// CreateSslCertificate API operation.
type CreateSslCertificateResponse struct {
	*CreateSslCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSslCertificate request.
func (r *CreateSslCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
