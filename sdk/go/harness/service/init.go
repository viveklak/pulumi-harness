// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package service

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/lbrlabs/pulumi-harness/sdk/go/harness"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "harness:service/ami:Ami":
		r = &Ami{}
	case "harness:service/codedeploy:Codedeploy":
		r = &Codedeploy{}
	case "harness:service/ecs:Ecs":
		r = &Ecs{}
	case "harness:service/helm:Helm":
		r = &Helm{}
	case "harness:service/kubernetes:Kubernetes":
		r = &Kubernetes{}
	case "harness:service/lambda:Lambda":
		r = &Lambda{}
	case "harness:service/ssh:Ssh":
		r = &Ssh{}
	case "harness:service/tanzu:Tanzu":
		r = &Tanzu{}
	case "harness:service/winrm:Winrm":
		r = &Winrm{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := harness.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"harness",
		"service/ami",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harness",
		"service/codedeploy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harness",
		"service/ecs",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harness",
		"service/helm",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harness",
		"service/kubernetes",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harness",
		"service/lambda",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harness",
		"service/ssh",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harness",
		"service/tanzu",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"harness",
		"service/winrm",
		&module{version},
	)
}