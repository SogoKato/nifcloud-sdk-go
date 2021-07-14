// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type AuthorizeDBSecurityGroupIngressInput struct {
	_ struct{} `type:"structure"`

	CIDRIP *string `locationName:"CIDRIP" type:"string"`

	// DBSecurityGroupName is a required field
	DBSecurityGroupName *string `locationName:"DBSecurityGroupName" type:"string" required:"true"`

	EC2SecurityGroupId *string `locationName:"EC2SecurityGroupId" type:"string"`

	EC2SecurityGroupName *string `locationName:"EC2SecurityGroupName" type:"string"`

	EC2SecurityGroupOwnerId *string `locationName:"EC2SecurityGroupOwnerId" type:"string"`
}

// String returns the string representation
func (s AuthorizeDBSecurityGroupIngressInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AuthorizeDBSecurityGroupIngressInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AuthorizeDBSecurityGroupIngressInput"}

	if s.DBSecurityGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBSecurityGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AuthorizeDBSecurityGroupIngressOutput struct {
	_ struct{} `type:"structure"`

	DBSecurityGroup *DBSecurityGroup `locationName:"DBSecurityGroup" type:"structure"`

	ResponseMetadata *ResponseMetadata `locationName:"ResponseMetadata" type:"structure"`
}

// String returns the string representation
func (s AuthorizeDBSecurityGroupIngressOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opAuthorizeDBSecurityGroupIngress = "AuthorizeDBSecurityGroupIngress"

// AuthorizeDBSecurityGroupIngressRequest returns a request value for making API operation for
// NIFCLOUD RDB.
//
//    // Example sending a request using AuthorizeDBSecurityGroupIngressRequest.
//    req := client.AuthorizeDBSecurityGroupIngressRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rdb/AuthorizeDBSecurityGroupIngress.htm
func (c *Client) AuthorizeDBSecurityGroupIngressRequest(input *AuthorizeDBSecurityGroupIngressInput) AuthorizeDBSecurityGroupIngressRequest {
	op := &aws.Operation{
		Name:       opAuthorizeDBSecurityGroupIngress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AuthorizeDBSecurityGroupIngressInput{}
	}

	req := c.newRequest(op, input, &AuthorizeDBSecurityGroupIngressOutput{})

	return AuthorizeDBSecurityGroupIngressRequest{Request: req, Input: input, Copy: c.AuthorizeDBSecurityGroupIngressRequest}
}

// AuthorizeDBSecurityGroupIngressRequest is the request type for the
// AuthorizeDBSecurityGroupIngress API operation.
type AuthorizeDBSecurityGroupIngressRequest struct {
	*aws.Request
	Input *AuthorizeDBSecurityGroupIngressInput
	Copy  func(*AuthorizeDBSecurityGroupIngressInput) AuthorizeDBSecurityGroupIngressRequest
}

// Send marshals and sends the AuthorizeDBSecurityGroupIngress API request.
func (r AuthorizeDBSecurityGroupIngressRequest) Send(ctx context.Context) (*AuthorizeDBSecurityGroupIngressResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AuthorizeDBSecurityGroupIngressResponse{
		AuthorizeDBSecurityGroupIngressOutput: r.Request.Data.(*AuthorizeDBSecurityGroupIngressOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AuthorizeDBSecurityGroupIngressResponse is the response type for the
// AuthorizeDBSecurityGroupIngress API operation.
type AuthorizeDBSecurityGroupIngressResponse struct {
	*AuthorizeDBSecurityGroupIngressOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AuthorizeDBSecurityGroupIngress request.
func (r *AuthorizeDBSecurityGroupIngressResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
