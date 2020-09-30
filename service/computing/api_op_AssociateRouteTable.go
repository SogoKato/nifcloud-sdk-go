// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type AssociateRouteTableInput struct {
	_ struct{} `type:"structure"`

	Agreement *bool `locationName:"Agreement" type:"boolean"`

	RouteTableId *string `locationName:"RouteTableId" type:"string"`

	RouterId *string `locationName:"RouterId" type:"string"`

	RouterName *string `locationName:"RouterName" type:"string"`

	SubnetId *string `locationName:"SubnetId" type:"string"`
}

// String returns the string representation
func (s AssociateRouteTableInput) String() string {
	return nifcloudutil.Prettify(s)
}

type AssociateRouteTableOutput struct {
	_ struct{} `type:"structure"`

	AssociationId *string `locationName:"associationId" type:"string"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s AssociateRouteTableOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opAssociateRouteTable = "AssociateRouteTable"

// AssociateRouteTableRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using AssociateRouteTableRequest.
//    req := client.AssociateRouteTableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rest/AssociateRouteTable.htm
func (c *Client) AssociateRouteTableRequest(input *AssociateRouteTableInput) AssociateRouteTableRequest {
	op := &aws.Operation{
		Name:       opAssociateRouteTable,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &AssociateRouteTableInput{}
	}

	req := c.newRequest(op, input, &AssociateRouteTableOutput{})

	return AssociateRouteTableRequest{Request: req, Input: input, Copy: c.AssociateRouteTableRequest}
}

// AssociateRouteTableRequest is the request type for the
// AssociateRouteTable API operation.
type AssociateRouteTableRequest struct {
	*aws.Request
	Input *AssociateRouteTableInput
	Copy  func(*AssociateRouteTableInput) AssociateRouteTableRequest
}

// Send marshals and sends the AssociateRouteTable API request.
func (r AssociateRouteTableRequest) Send(ctx context.Context) (*AssociateRouteTableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateRouteTableResponse{
		AssociateRouteTableOutput: r.Request.Data.(*AssociateRouteTableOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateRouteTableResponse is the response type for the
// AssociateRouteTable API operation.
type AssociateRouteTableResponse struct {
	*AssociateRouteTableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateRouteTable request.
func (r *AssociateRouteTableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
