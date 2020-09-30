// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type NiftyDescribeNatTablesInput struct {
	_ struct{} `type:"structure"`

	Filter []RequestFilterOfNiftyDescribeNatTables `locationName:"Filter" type:"list"`

	NatTableId []string `locationName:"NatTableId" type:"list"`
}

// String returns the string representation
func (s NiftyDescribeNatTablesInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyDescribeNatTablesOutput struct {
	_ struct{} `type:"structure"`

	NatTableSet []NatTableSet `locationName:"natTableSet" locationNameList:"item" type:"list"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s NiftyDescribeNatTablesOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyDescribeNatTables = "NiftyDescribeNatTables"

// NiftyDescribeNatTablesRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyDescribeNatTablesRequest.
//    req := client.NiftyDescribeNatTablesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/NiftyDescribeNatTables.htm
func (c *Client) NiftyDescribeNatTablesRequest(input *NiftyDescribeNatTablesInput) NiftyDescribeNatTablesRequest {
	op := &aws.Operation{
		Name:       opNiftyDescribeNatTables,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyDescribeNatTablesInput{}
	}

	req := c.newRequest(op, input, &NiftyDescribeNatTablesOutput{})

	return NiftyDescribeNatTablesRequest{Request: req, Input: input, Copy: c.NiftyDescribeNatTablesRequest}
}

// NiftyDescribeNatTablesRequest is the request type for the
// NiftyDescribeNatTables API operation.
type NiftyDescribeNatTablesRequest struct {
	*aws.Request
	Input *NiftyDescribeNatTablesInput
	Copy  func(*NiftyDescribeNatTablesInput) NiftyDescribeNatTablesRequest
}

// Send marshals and sends the NiftyDescribeNatTables API request.
func (r NiftyDescribeNatTablesRequest) Send(ctx context.Context) (*NiftyDescribeNatTablesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyDescribeNatTablesResponse{
		NiftyDescribeNatTablesOutput: r.Request.Data.(*NiftyDescribeNatTablesOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyDescribeNatTablesResponse is the response type for the
// NiftyDescribeNatTables API operation.
type NiftyDescribeNatTablesResponse struct {
	*NiftyDescribeNatTablesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyDescribeNatTables request.
func (r *NiftyDescribeNatTablesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
