// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package platform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Datasource for looking up a Gcp connector.
func LookupGcpConnector(ctx *pulumi.Context, args *LookupGcpConnectorArgs, opts ...pulumi.InvokeOption) (*LookupGcpConnectorResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupGcpConnectorResult
	err := ctx.Invoke("harness:platform/getGcpConnector:getGcpConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGcpConnector.
type LookupGcpConnectorArgs struct {
	// Unique identifier of the resource.
	Identifier *string `pulumi:"identifier"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId *string `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId *string `pulumi:"projectId"`
}

// A collection of values returned by getGcpConnector.
type LookupGcpConnectorResult struct {
	// Description of the resource.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Unique identifier of the resource.
	Identifier *string `pulumi:"identifier"`
	// Inherit configuration from delegate.
	InheritFromDelegates []GetGcpConnectorInheritFromDelegate `pulumi:"inheritFromDelegates"`
	// Manual credential configuration.
	Manuals []GetGcpConnectorManual `pulumi:"manuals"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId *string `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId *string `pulumi:"projectId"`
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags []string `pulumi:"tags"`
}

func LookupGcpConnectorOutput(ctx *pulumi.Context, args LookupGcpConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupGcpConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGcpConnectorResult, error) {
			args := v.(LookupGcpConnectorArgs)
			r, err := LookupGcpConnector(ctx, &args, opts...)
			var s LookupGcpConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGcpConnectorResultOutput)
}

// A collection of arguments for invoking getGcpConnector.
type LookupGcpConnectorOutputArgs struct {
	// Unique identifier of the resource.
	Identifier pulumi.StringPtrInput `pulumi:"identifier"`
	// Name of the resource.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
}

func (LookupGcpConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGcpConnectorArgs)(nil)).Elem()
}

// A collection of values returned by getGcpConnector.
type LookupGcpConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupGcpConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGcpConnectorResult)(nil)).Elem()
}

func (o LookupGcpConnectorResultOutput) ToLookupGcpConnectorResultOutput() LookupGcpConnectorResultOutput {
	return o
}

func (o LookupGcpConnectorResultOutput) ToLookupGcpConnectorResultOutputWithContext(ctx context.Context) LookupGcpConnectorResultOutput {
	return o
}

// Description of the resource.
func (o LookupGcpConnectorResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGcpConnectorResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGcpConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGcpConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// Unique identifier of the resource.
func (o LookupGcpConnectorResultOutput) Identifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGcpConnectorResult) *string { return v.Identifier }).(pulumi.StringPtrOutput)
}

// Inherit configuration from delegate.
func (o LookupGcpConnectorResultOutput) InheritFromDelegates() GetGcpConnectorInheritFromDelegateArrayOutput {
	return o.ApplyT(func(v LookupGcpConnectorResult) []GetGcpConnectorInheritFromDelegate { return v.InheritFromDelegates }).(GetGcpConnectorInheritFromDelegateArrayOutput)
}

// Manual credential configuration.
func (o LookupGcpConnectorResultOutput) Manuals() GetGcpConnectorManualArrayOutput {
	return o.ApplyT(func(v LookupGcpConnectorResult) []GetGcpConnectorManual { return v.Manuals }).(GetGcpConnectorManualArrayOutput)
}

// Name of the resource.
func (o LookupGcpConnectorResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGcpConnectorResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Unique identifier of the organization.
func (o LookupGcpConnectorResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGcpConnectorResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

// Unique identifier of the project.
func (o LookupGcpConnectorResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGcpConnectorResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// Tags to associate with the resource. Tags should be in the form `name:value`.
func (o LookupGcpConnectorResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGcpConnectorResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGcpConnectorResultOutput{})
}
