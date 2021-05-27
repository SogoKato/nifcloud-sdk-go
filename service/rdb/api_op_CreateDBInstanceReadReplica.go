// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type CreateDBInstanceReadReplicaInput struct {
	_ struct{} `type:"structure"`

	AccountingType *string `locationName:"AccountingType" type:"string"`

	DBInstanceClass *string `locationName:"DBInstanceClass" type:"string"`

	// DBInstanceIdentifier is a required field
	DBInstanceIdentifier *string `locationName:"DBInstanceIdentifier" type:"string" required:"true"`

	NiftyReadReplicaPrivateAddress *string `locationName:"NiftyReadReplicaPrivateAddress" type:"string"`

	NiftyStorageType *int64 `locationName:"NiftyStorageType" type:"integer"`

	// SourceDBInstanceIdentifier is a required field
	SourceDBInstanceIdentifier *string `locationName:"SourceDBInstanceIdentifier" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateDBInstanceReadReplicaInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDBInstanceReadReplicaInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDBInstanceReadReplicaInput"}

	if s.DBInstanceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("DBInstanceIdentifier"))
	}

	if s.SourceDBInstanceIdentifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceDBInstanceIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateDBInstanceReadReplicaOutput struct {
	_ struct{} `type:"structure"`

	DBInstance *DBInstance `type:"structure"`
}

// String returns the string representation
func (s CreateDBInstanceReadReplicaOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opCreateDBInstanceReadReplica = "CreateDBInstanceReadReplica"

// CreateDBInstanceReadReplicaRequest returns a request value for making API operation for
// NIFCLOUD RDB.
//
//    // Example sending a request using CreateDBInstanceReadReplicaRequest.
//    req := client.CreateDBInstanceReadReplicaRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rdb/CreateDBInstanceReadReplica.htm
func (c *Client) CreateDBInstanceReadReplicaRequest(input *CreateDBInstanceReadReplicaInput) CreateDBInstanceReadReplicaRequest {
	op := &aws.Operation{
		Name:       opCreateDBInstanceReadReplica,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDBInstanceReadReplicaInput{}
	}

	req := c.newRequest(op, input, &CreateDBInstanceReadReplicaOutput{})

	return CreateDBInstanceReadReplicaRequest{Request: req, Input: input, Copy: c.CreateDBInstanceReadReplicaRequest}
}

// CreateDBInstanceReadReplicaRequest is the request type for the
// CreateDBInstanceReadReplica API operation.
type CreateDBInstanceReadReplicaRequest struct {
	*aws.Request
	Input *CreateDBInstanceReadReplicaInput
	Copy  func(*CreateDBInstanceReadReplicaInput) CreateDBInstanceReadReplicaRequest
}

// Send marshals and sends the CreateDBInstanceReadReplica API request.
func (r CreateDBInstanceReadReplicaRequest) Send(ctx context.Context) (*CreateDBInstanceReadReplicaResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDBInstanceReadReplicaResponse{
		CreateDBInstanceReadReplicaOutput: r.Request.Data.(*CreateDBInstanceReadReplicaOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDBInstanceReadReplicaResponse is the response type for the
// CreateDBInstanceReadReplica API operation.
type CreateDBInstanceReadReplicaResponse struct {
	*CreateDBInstanceReadReplicaOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDBInstanceReadReplica request.
func (r *CreateDBInstanceReadReplicaResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
