// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package nas

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type DescribeNASSecurityGroupsInput struct {
	_ struct{} `type:"structure"`

	NASSecurityGroupName *string `locationName:"NASSecurityGroupName" type:"string"`
}

// String returns the string representation
func (s DescribeNASSecurityGroupsInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DescribeNASSecurityGroupsOutput struct {
	_ struct{} `type:"structure"`

	NASSecurityGroups []NASSecurityGroup `locationNameList:"NASSecurityGroup" type:"list"`
}

// String returns the string representation
func (s DescribeNASSecurityGroupsOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDescribeNASSecurityGroups = "DescribeNASSecurityGroups"

// DescribeNASSecurityGroupsRequest returns a request value for making API operation for
// NIFCLOUD NAS.
//
//    // Example sending a request using DescribeNASSecurityGroupsRequest.
//    req := client.DescribeNASSecurityGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/nas/DescribeNASSecurityGroups.htm
func (c *Client) DescribeNASSecurityGroupsRequest(input *DescribeNASSecurityGroupsInput) DescribeNASSecurityGroupsRequest {
	op := &aws.Operation{
		Name:       opDescribeNASSecurityGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeNASSecurityGroupsInput{}
	}

	req := c.newRequest(op, input, &DescribeNASSecurityGroupsOutput{})

	return DescribeNASSecurityGroupsRequest{Request: req, Input: input, Copy: c.DescribeNASSecurityGroupsRequest}
}

// DescribeNASSecurityGroupsRequest is the request type for the
// DescribeNASSecurityGroups API operation.
type DescribeNASSecurityGroupsRequest struct {
	*aws.Request
	Input *DescribeNASSecurityGroupsInput
	Copy  func(*DescribeNASSecurityGroupsInput) DescribeNASSecurityGroupsRequest
}

// Send marshals and sends the DescribeNASSecurityGroups API request.
func (r DescribeNASSecurityGroupsRequest) Send(ctx context.Context) (*DescribeNASSecurityGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeNASSecurityGroupsResponse{
		DescribeNASSecurityGroupsOutput: r.Request.Data.(*DescribeNASSecurityGroupsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeNASSecurityGroupsResponse is the response type for the
// DescribeNASSecurityGroups API operation.
type DescribeNASSecurityGroupsResponse struct {
	*DescribeNASSecurityGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeNASSecurityGroups request.
func (r *DescribeNASSecurityGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
