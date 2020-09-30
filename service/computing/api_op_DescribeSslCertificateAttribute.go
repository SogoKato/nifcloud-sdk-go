// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type DescribeSslCertificateAttributeInput struct {
	_ struct{} `type:"structure"`

	Attribute AttributeOfDescribeSslCertificateAttributeRequest `locationName:"Attribute" type:"string" enum:"true"`

	FqdnId *string `locationName:"FqdnId" type:"string"`
}

// String returns the string representation
func (s DescribeSslCertificateAttributeInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DescribeSslCertificateAttributeOutput struct {
	_ struct{} `type:"structure"`

	CaState *CaState `locationName:"caState" type:"structure"`

	CertAuthority *CertAuthority `locationName:"certAuthority" type:"structure"`

	CertInfo *CertInfoOfDescribeSslCertificateAttribute `locationName:"certInfo" type:"structure"`

	CertState *CertState `locationName:"certState" type:"structure"`

	Count *Count `locationName:"count" type:"structure"`

	Description *DescriptionOfDescribeSslCertificateAttribute `locationName:"description" type:"structure"`

	Fqdn *string `locationName:"fqdn" type:"string"`

	FqdnId *string `locationName:"fqdnId" type:"string"`

	KeyLength *KeyLength `locationName:"keyLength" type:"structure"`

	Period *PeriodOfDescribeSslCertificateAttribute `locationName:"period" type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	UploadState *UploadState `locationName:"uploadState" type:"structure"`
}

// String returns the string representation
func (s DescribeSslCertificateAttributeOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDescribeSslCertificateAttribute = "DescribeSslCertificateAttribute"

// DescribeSslCertificateAttributeRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DescribeSslCertificateAttributeRequest.
//    req := client.DescribeSslCertificateAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/DescribeSslCertificateAttribute.htm
func (c *Client) DescribeSslCertificateAttributeRequest(input *DescribeSslCertificateAttributeInput) DescribeSslCertificateAttributeRequest {
	op := &aws.Operation{
		Name:       opDescribeSslCertificateAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DescribeSslCertificateAttributeInput{}
	}

	req := c.newRequest(op, input, &DescribeSslCertificateAttributeOutput{})

	return DescribeSslCertificateAttributeRequest{Request: req, Input: input, Copy: c.DescribeSslCertificateAttributeRequest}
}

// DescribeSslCertificateAttributeRequest is the request type for the
// DescribeSslCertificateAttribute API operation.
type DescribeSslCertificateAttributeRequest struct {
	*aws.Request
	Input *DescribeSslCertificateAttributeInput
	Copy  func(*DescribeSslCertificateAttributeInput) DescribeSslCertificateAttributeRequest
}

// Send marshals and sends the DescribeSslCertificateAttribute API request.
func (r DescribeSslCertificateAttributeRequest) Send(ctx context.Context) (*DescribeSslCertificateAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSslCertificateAttributeResponse{
		DescribeSslCertificateAttributeOutput: r.Request.Data.(*DescribeSslCertificateAttributeOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSslCertificateAttributeResponse is the response type for the
// DescribeSslCertificateAttribute API operation.
type DescribeSslCertificateAttributeResponse struct {
	*DescribeSslCertificateAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSslCertificateAttribute request.
func (r *DescribeSslCertificateAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
