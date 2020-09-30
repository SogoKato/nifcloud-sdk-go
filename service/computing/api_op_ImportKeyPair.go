// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type ImportKeyPairInput struct {
	_ struct{} `type:"structure"`

	Description *string `locationName:"Description" type:"string"`

	KeyName *string `locationName:"KeyName" type:"string"`

	PublicKeyMaterial *string `locationName:"PublicKeyMaterial" type:"string"`
}

// String returns the string representation
func (s ImportKeyPairInput) String() string {
	return nifcloudutil.Prettify(s)
}

type ImportKeyPairOutput struct {
	_ struct{} `type:"structure"`

	KeyFingerprint *string `locationName:"keyFingerprint" type:"string"`

	KeyName *string `locationName:"keyName" type:"string"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s ImportKeyPairOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opImportKeyPair = "ImportKeyPair"

// ImportKeyPairRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using ImportKeyPairRequest.
//    req := client.ImportKeyPairRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/ImportKeyPair.htm
func (c *Client) ImportKeyPairRequest(input *ImportKeyPairInput) ImportKeyPairRequest {
	op := &aws.Operation{
		Name:       opImportKeyPair,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &ImportKeyPairInput{}
	}

	req := c.newRequest(op, input, &ImportKeyPairOutput{})

	return ImportKeyPairRequest{Request: req, Input: input, Copy: c.ImportKeyPairRequest}
}

// ImportKeyPairRequest is the request type for the
// ImportKeyPair API operation.
type ImportKeyPairRequest struct {
	*aws.Request
	Input *ImportKeyPairInput
	Copy  func(*ImportKeyPairInput) ImportKeyPairRequest
}

// Send marshals and sends the ImportKeyPair API request.
func (r ImportKeyPairRequest) Send(ctx context.Context) (*ImportKeyPairResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ImportKeyPairResponse{
		ImportKeyPairOutput: r.Request.Data.(*ImportKeyPairOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ImportKeyPairResponse is the response type for the
// ImportKeyPair API operation.
type ImportKeyPairResponse struct {
	*ImportKeyPairOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ImportKeyPair request.
func (r *ImportKeyPairResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
