// Code generated by aws/endpoints/v3model_codegen.go. DO NOT EDIT.

package endpoints

import (
	"regexp"
)

// NewDefaultResolver returns an Endpoint resolver that will be able
// to resolve endpoints for: NIFCLOUD Standard.
func NewDefaultResolver() *Resolver {
	return &Resolver{
		partitions: defaultPartitions,
	}
}

var defaultPartitions = partitions{
	nifcloudPartition,
}

var nifcloudPartition = partition{
	ID:        "nifcloud",
	Name:      "NIFCLOUD Standard",
	DNSSuffix: "api.nifcloud.com",
	RegionRegex: regionRegex{
		Regexp: func() *regexp.Regexp {
			reg, _ := regexp.Compile("^(us|jp)\\-\\w+\\-\\d+$")
			return reg
		}(),
	},
	Defaults: endpoint{
		Hostname:          "{region}.{service}.{dnsSuffix}",
		Protocols:         []string{"https"},
		SignatureVersions: []string{"v4", "v3", "v2"},
	},
	Regions: regions{
		"jp-east-1": region{
			Description: "jp-east-1",
		},
		"jp-east-2": region{
			Description: "jp-east-2",
		},
		"jp-east-3": region{
			Description: "jp-east-3",
		},
		"jp-east-4": region{
			Description: "jp-east-4",
		},
		"jp-west-1": region{
			Description: "jp-west-1",
		},
		"us-east-1": region{
			Description: "us-east-1",
		},
	},
	Services: services{
		"cloudsearchdomain": service{
			Defaults: endpoint{
				Unresolveable: boxedTrue,
			},
		},
		"computing": service{
			Defaults: endpoint{
				SignatureVersions: []string{"v2"},
			},
			Endpoints: endpoints{
				"jp-east-1": endpoint{},
				"jp-east-2": endpoint{},
				"jp-east-3": endpoint{},
				"jp-east-4": endpoint{},
				"jp-west-1": endpoint{},
				"us-east-1": endpoint{},
			},
		},
		"data.iot": service{
			Defaults: endpoint{
				Unresolveable: boxedTrue,
			},
		},
		"ec2metadata": service{
			PartitionEndpoint: "aws-global",
			IsRegionalized:    boxedFalse,

			Endpoints: endpoints{
				"aws-global": endpoint{
					Hostname:  "169.254.169.254/latest",
					Protocols: []string{"http"},
				},
			},
		},
		"hatoba": service{

			Endpoints: endpoints{
				"jp-east-1": endpoint{},
			},
		},
		"nas": service{

			Endpoints: endpoints{
				"jp-east-1": endpoint{},
				"jp-east-2": endpoint{},
				"jp-east-3": endpoint{},
				"jp-east-4": endpoint{},
				"jp-west-1": endpoint{},
			},
		},
		"rdb": service{

			Endpoints: endpoints{
				"jp-east-1": endpoint{},
				"jp-east-2": endpoint{},
				"jp-east-3": endpoint{},
				"jp-east-4": endpoint{},
				"jp-west-1": endpoint{},
			},
		},
		"script": service{
			PartitionEndpoint: "aws-global",
			IsRegionalized:    boxedFalse,

			Endpoints: endpoints{
				"aws-global": endpoint{
					Hostname: "script.api.nifcloud.com",
					CredentialScope: credentialScope{
						Region: "jp-east-1",
					},
				},
			},
		},
	},
}
