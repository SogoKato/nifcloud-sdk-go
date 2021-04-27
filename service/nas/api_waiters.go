// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package nas

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

// WaitUntilNASInstanceAvailable uses the nas API operation
// DescribeNASInstances to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilNASInstanceAvailable(ctx context.Context, input *DescribeNASInstancesInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilNASInstanceAvailable",
		MaxAttempts: 80,
		Delay:       aws.ConstantWaiterDelay(40 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "NASInstances[].NASInstanceStatus",
				Expected: "available",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeNASInstancesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeNASInstancesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilNASInstanceDeleted uses the nas API operation
// DescribeNASInstances to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilNASInstanceDeleted(ctx context.Context, input *DescribeNASInstancesInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilNASInstanceDeleted",
		MaxAttempts: 80,
		Delay:       aws.ConstantWaiterDelay(40 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:    aws.SuccessWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "Client.InvalidParameter.NotFound.NASInstanceIdentifier",
			},
			{
				State:   aws.RetryWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "length(NASInstances[]) > `0`",
				Expected: true,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeNASInstancesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeNASInstancesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilNASInstanceExists uses the nas API operation
// DescribeNASInstances to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilNASInstanceExists(ctx context.Context, input *DescribeNASInstancesInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilNASInstanceExists",
		MaxAttempts: 80,
		Delay:       aws.ConstantWaiterDelay(40 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "length(NASInstances[]) > `0`",
				Expected: true,
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "Client.InvalidParameter.NotFound.NASInstanceIdentifier",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeNASInstancesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeNASInstancesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilNASInstanceFailed uses the nas API operation
// DescribeNASInstances to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilNASInstanceFailed(ctx context.Context, input *DescribeNASInstancesInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilNASInstanceFailed",
		MaxAttempts: 80,
		Delay:       aws.ConstantWaiterDelay(40 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "NASInstances[].NASInstanceStatus",
				Expected: "failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeNASInstancesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeNASInstancesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilNASInstanceStorageFull uses the nas API operation
// DescribeNASInstances to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilNASInstanceStorageFull(ctx context.Context, input *DescribeNASInstancesInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilNASInstanceStorageFull",
		MaxAttempts: 80,
		Delay:       aws.ConstantWaiterDelay(40 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "NASInstances[].DNASInstanceStatus",
				Expected: "storage-full",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeNASInstancesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeNASInstancesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilNASSecurityGroupDeleted uses the nas API operation
// DescribeNASSecurityGroups to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilNASSecurityGroupDeleted(ctx context.Context, input *DescribeNASSecurityGroupsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilNASSecurityGroupDeleted",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:    aws.SuccessWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "Client.InvalidParameter.NotFound.NASSecurityGroupName",
			},
			{
				State:   aws.RetryWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "length(NASSecurityGroups[]) > `0`",
				Expected: true,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeNASSecurityGroupsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeNASSecurityGroupsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilNASSecurityGroupExists uses the nas API operation
// DescribeNASSecurityGroups to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilNASSecurityGroupExists(ctx context.Context, input *DescribeNASSecurityGroupsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilNASSecurityGroupExists",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "length(NASSecurityGroups[]) > `0`",
				Expected: true,
			},
			{
				State:    aws.RetryWaiterState,
				Matcher:  aws.ErrorWaiterMatch,
				Expected: "Client.InvalidParameter.NotFound.NASSecurityGroupName",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeNASSecurityGroupsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeNASSecurityGroupsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilNASSecurityGroupIPRangesAuthFailed uses the nas API operation
// DescribeNASSecurityGroups to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilNASSecurityGroupIPRangesAuthFailed(ctx context.Context, input *DescribeNASSecurityGroupsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilNASSecurityGroupIPRangesAuthFailed",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "NASSecurityGroups[].IPRanges[].Status",
				Expected: "auth-failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeNASSecurityGroupsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeNASSecurityGroupsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilNASSecurityGroupIPRangesAuthorized uses the nas API operation
// DescribeNASSecurityGroups to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilNASSecurityGroupIPRangesAuthorized(ctx context.Context, input *DescribeNASSecurityGroupsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilNASSecurityGroupIPRangesAuthorized",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "NASSecurityGroups[].IPRanges[].Status",
				Expected: "authorized",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeNASSecurityGroupsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeNASSecurityGroupsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilNASSecurityGroupIPRangesEmptied uses the nas API operation
// DescribeNASSecurityGroups to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilNASSecurityGroupIPRangesEmptied(ctx context.Context, input *DescribeNASSecurityGroupsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilNASSecurityGroupIPRangesEmptied",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "length(NASSecurityGroups[0].IPRanges[]) == `0`",
				Expected: true,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeNASSecurityGroupsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeNASSecurityGroupsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilNASSecurityGroupIPRangesRevokeFailed uses the nas API operation
// DescribeNASSecurityGroups to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilNASSecurityGroupIPRangesRevokeFailed(ctx context.Context, input *DescribeNASSecurityGroupsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilNASSecurityGroupIPRangesRevokeFailed",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "NASSecurityGroups[].IPRanges[].Status",
				Expected: "revoke-failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeNASSecurityGroupsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeNASSecurityGroupsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilNASSecurityGroupSecurityGroupsAuthFailed uses the nas API operation
// DescribeNASSecurityGroups to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilNASSecurityGroupSecurityGroupsAuthFailed(ctx context.Context, input *DescribeNASSecurityGroupsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilNASSecurityGroupSecurityGroupsAuthFailed",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "NASSecurityGroups[].SecurityGroups[].Status",
				Expected: "auth-failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeNASSecurityGroupsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeNASSecurityGroupsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilNASSecurityGroupSecurityGroupsAuthorized uses the nas API operation
// DescribeNASSecurityGroups to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilNASSecurityGroupSecurityGroupsAuthorized(ctx context.Context, input *DescribeNASSecurityGroupsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilNASSecurityGroupSecurityGroupsAuthorized",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "NASSecurityGroups[].SecurityGroups[].Status",
				Expected: "authorized",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeNASSecurityGroupsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeNASSecurityGroupsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilNASSecurityGroupSecurityGroupsEmptied uses the nas API operation
// DescribeNASSecurityGroups to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilNASSecurityGroupSecurityGroupsEmptied(ctx context.Context, input *DescribeNASSecurityGroupsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilNASSecurityGroupSecurityGroupsEmptied",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathWaiterMatch, Argument: "length(NASSecurityGroups[0].SecurityGroups[]) == `0`",
				Expected: true,
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeNASSecurityGroupsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeNASSecurityGroupsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}

// WaitUntilNASSecurityGroupSecurityGroupsRevokeFailed uses the nas API operation
// DescribeNASSecurityGroups to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Client) WaitUntilNASSecurityGroupSecurityGroupsRevokeFailed(ctx context.Context, input *DescribeNASSecurityGroupsInput, opts ...aws.WaiterOption) error {
	w := aws.Waiter{
		Name:        "WaitUntilNASSecurityGroupSecurityGroupsRevokeFailed",
		MaxAttempts: 40,
		Delay:       aws.ConstantWaiterDelay(20 * time.Second),
		Acceptors: []aws.WaiterAcceptor{
			{
				State:   aws.SuccessWaiterState,
				Matcher: aws.PathAllWaiterMatch, Argument: "NASSecurityGroups[].SecurityGroups[].Status",
				Expected: "revoke-failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []aws.Option) (*aws.Request, error) {
			var inCpy *DescribeNASSecurityGroupsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req := c.DescribeNASSecurityGroupsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req.Request, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.Wait(ctx)
}
