// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package platform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Data source for retrieving a Harness environment.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-harness/sdk/go/harness/platform"
//	"github.com/pulumi/pulumi-harness/sdk/go/harness/platform"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := platform.LookupEnvironment(ctx, &platform.LookupEnvironmentArgs{
//				Name:      pulumi.StringRef("name"),
//				OrgId:     "org_id",
//				ProjectId: "project_id",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupEnvironment(ctx *pulumi.Context, args *LookupEnvironmentArgs, opts ...pulumi.InvokeOption) (*LookupEnvironmentResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupEnvironmentResult
	err := ctx.Invoke("harness:platform/getEnvironment:getEnvironment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEnvironment.
type LookupEnvironmentArgs struct {
	// Unique identifier of the resource.
	Identifier *string `pulumi:"identifier"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId string `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getEnvironment.
type LookupEnvironmentResult struct {
	// Color of the environment.
	Color string `pulumi:"color"`
	// Description of the resource.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Unique identifier of the resource.
	Identifier *string `pulumi:"identifier"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId string `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId string `pulumi:"projectId"`
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags []string `pulumi:"tags"`
	// The type of environment.
	Type string `pulumi:"type"`
}

func LookupEnvironmentOutput(ctx *pulumi.Context, args LookupEnvironmentOutputArgs, opts ...pulumi.InvokeOption) LookupEnvironmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEnvironmentResult, error) {
			args := v.(LookupEnvironmentArgs)
			r, err := LookupEnvironment(ctx, &args, opts...)
			var s LookupEnvironmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEnvironmentResultOutput)
}

// A collection of arguments for invoking getEnvironment.
type LookupEnvironmentOutputArgs struct {
	// Unique identifier of the resource.
	Identifier pulumi.StringPtrInput `pulumi:"identifier"`
	// Name of the resource.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId pulumi.StringInput `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (LookupEnvironmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentArgs)(nil)).Elem()
}

// A collection of values returned by getEnvironment.
type LookupEnvironmentResultOutput struct{ *pulumi.OutputState }

func (LookupEnvironmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEnvironmentResult)(nil)).Elem()
}

func (o LookupEnvironmentResultOutput) ToLookupEnvironmentResultOutput() LookupEnvironmentResultOutput {
	return o
}

func (o LookupEnvironmentResultOutput) ToLookupEnvironmentResultOutputWithContext(ctx context.Context) LookupEnvironmentResultOutput {
	return o
}

// Color of the environment.
func (o LookupEnvironmentResultOutput) Color() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Color }).(pulumi.StringOutput)
}

// Description of the resource.
func (o LookupEnvironmentResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupEnvironmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// Unique identifier of the resource.
func (o LookupEnvironmentResultOutput) Identifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.Identifier }).(pulumi.StringPtrOutput)
}

// Name of the resource.
func (o LookupEnvironmentResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Unique identifier of the organization.
func (o LookupEnvironmentResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.OrgId }).(pulumi.StringOutput)
}

// Unique identifier of the project.
func (o LookupEnvironmentResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Tags to associate with the resource. Tags should be in the form `name:value`.
func (o LookupEnvironmentResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The type of environment.
func (o LookupEnvironmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEnvironmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEnvironmentResultOutput{})
}
