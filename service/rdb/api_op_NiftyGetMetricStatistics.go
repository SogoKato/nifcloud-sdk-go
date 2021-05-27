// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdb

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nifcloud/nifcloud-sdk-go/internal/nifcloudutil"
)

type NiftyGetMetricStatisticsInput struct {
	_ struct{} `type:"structure"`

	// Dimensions is a required field
	Dimensions []RequestDimensions `locationName:"Dimensions" locationNameList:"member" type:"list" required:"true"`

	EndTime *time.Time `locationName:"EndTime" type:"timestamp"`

	// MetricName is a required field
	MetricName *string `locationName:"MetricName" type:"string" required:"true"`

	StartTime *time.Time `locationName:"StartTime" type:"timestamp"`
}

// String returns the string representation
func (s NiftyGetMetricStatisticsInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *NiftyGetMetricStatisticsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "NiftyGetMetricStatisticsInput"}

	if s.Dimensions == nil {
		invalidParams.Add(aws.NewErrParamRequired("Dimensions"))
	}

	if s.MetricName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MetricName"))
	}
	if s.Dimensions != nil {
		for i, v := range s.Dimensions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Dimensions", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type NiftyGetMetricStatisticsOutput struct {
	_ struct{} `type:"structure"`

	Datapoints []Member `locationNameList:"member" type:"list"`

	Label *string `type:"string"`
}

// String returns the string representation
func (s NiftyGetMetricStatisticsOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyGetMetricStatistics = "NiftyGetMetricStatistics"

// NiftyGetMetricStatisticsRequest returns a request value for making API operation for
// NIFCLOUD RDB.
//
//    // Example sending a request using NiftyGetMetricStatisticsRequest.
//    req := client.NiftyGetMetricStatisticsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://pfs.nifcloud.com/api/rdb/NiftyGetMetricStatistics.htm
func (c *Client) NiftyGetMetricStatisticsRequest(input *NiftyGetMetricStatisticsInput) NiftyGetMetricStatisticsRequest {
	op := &aws.Operation{
		Name:       opNiftyGetMetricStatistics,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &NiftyGetMetricStatisticsInput{}
	}

	req := c.newRequest(op, input, &NiftyGetMetricStatisticsOutput{})

	return NiftyGetMetricStatisticsRequest{Request: req, Input: input, Copy: c.NiftyGetMetricStatisticsRequest}
}

// NiftyGetMetricStatisticsRequest is the request type for the
// NiftyGetMetricStatistics API operation.
type NiftyGetMetricStatisticsRequest struct {
	*aws.Request
	Input *NiftyGetMetricStatisticsInput
	Copy  func(*NiftyGetMetricStatisticsInput) NiftyGetMetricStatisticsRequest
}

// Send marshals and sends the NiftyGetMetricStatistics API request.
func (r NiftyGetMetricStatisticsRequest) Send(ctx context.Context) (*NiftyGetMetricStatisticsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyGetMetricStatisticsResponse{
		NiftyGetMetricStatisticsOutput: r.Request.Data.(*NiftyGetMetricStatisticsOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyGetMetricStatisticsResponse is the response type for the
// NiftyGetMetricStatistics API operation.
type NiftyGetMetricStatisticsResponse struct {
	*NiftyGetMetricStatisticsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyGetMetricStatistics request.
func (r *NiftyGetMetricStatisticsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
