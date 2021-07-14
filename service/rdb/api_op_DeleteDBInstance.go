// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type DeleteDBInstanceInput struct {
	_ struct{} `type:"structure"`

	// DBInstanceIdentifier is a required field
	DBInstanceIdentifier *string `locationName:"DBInstanceIdentifier" type:"string" required:"true"`

	FinalDBSnapshotIdentifier *string `locationName:"FinalDBSnapshotIdentifier" type:"string"`

	SkipFinalSnapshot *bool `locationName:"SkipFinalSnapshot" type:"boolean"`
}

// String returns the string representation
func (s DeleteDBInstanceInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDBInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteDBInstanceInput"}

	if s.DBInstanceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBInstanceIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteDBInstanceOutput struct {
	_ struct{} `type:"structure"`

	DBInstance *DBInstance `locationName:"DBInstance" type:"structure"`

	ResponseMetadata *ResponseMetadata `locationName:"ResponseMetadata" type:"structure"`
}

// String returns the string representation
func (s DeleteDBInstanceOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDeleteDBInstance = "DeleteDBInstance"

// DeleteDBInstanceRequest returns a request value for making API operation for
// NIFCLOUD RDB.
//
//    // Example sending a request using DeleteDBInstanceRequest.
//    req := client.DeleteDBInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rdb/DeleteDBInstance.htm
func (c *Client) DeleteDBInstanceRequest(input *DeleteDBInstanceInput) DeleteDBInstanceRequest {
	op := &aws.Operation{
		Name:       opDeleteDBInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDBInstanceInput{}
	}

	req := c.newRequest(op, input, &DeleteDBInstanceOutput{})

	return DeleteDBInstanceRequest{Request: req, Input: input, Copy: c.DeleteDBInstanceRequest}
}

// DeleteDBInstanceRequest is the request type for the
// DeleteDBInstance API operation.
type DeleteDBInstanceRequest struct {
	*aws.Request
	Input *DeleteDBInstanceInput
	Copy  func(*DeleteDBInstanceInput) DeleteDBInstanceRequest
}

// Send marshals and sends the DeleteDBInstance API request.
func (r DeleteDBInstanceRequest) Send(ctx context.Context) (*DeleteDBInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDBInstanceResponse{
		DeleteDBInstanceOutput: r.Request.Data.(*DeleteDBInstanceOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDBInstanceResponse is the response type for the
// DeleteDBInstance API operation.
type DeleteDBInstanceResponse struct {
	*DeleteDBInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDBInstance request.
func (r *DeleteDBInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
