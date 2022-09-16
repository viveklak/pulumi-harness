// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package platformconnector

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Datasource for looking up a PagerDuty connector.
func LookupPagerduty(ctx *pulumi.Context, args *LookupPagerdutyArgs, opts ...pulumi.InvokeOption) (*LookupPagerdutyResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupPagerdutyResult
	err := ctx.Invoke("harness:PlatformConnector/getPagerduty:getPagerduty", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPagerduty.
type LookupPagerdutyArgs struct {
	// Unique identifier of the resource.
	Identifier *string `pulumi:"identifier"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId *string `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId *string `pulumi:"projectId"`
}

// A collection of values returned by getPagerduty.
type LookupPagerdutyResult struct {
	// Reference to the Harness secret containing the api token.
	ApiTokenRef string `pulumi:"apiTokenRef"`
	// Connect using only the delegates which have these tags.
	DelegateSelectors []string `pulumi:"delegateSelectors"`
	// Description of the resource.
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Unique identifier of the resource.
	Identifier *string `pulumi:"identifier"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId *string `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId *string `pulumi:"projectId"`
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags []string `pulumi:"tags"`
}

func LookupPagerdutyOutput(ctx *pulumi.Context, args LookupPagerdutyOutputArgs, opts ...pulumi.InvokeOption) LookupPagerdutyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPagerdutyResult, error) {
			args := v.(LookupPagerdutyArgs)
			r, err := LookupPagerduty(ctx, &args, opts...)
			var s LookupPagerdutyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPagerdutyResultOutput)
}

// A collection of arguments for invoking getPagerduty.
type LookupPagerdutyOutputArgs struct {
	// Unique identifier of the resource.
	Identifier pulumi.StringPtrInput `pulumi:"identifier"`
	// Name of the resource.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
}

func (LookupPagerdutyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPagerdutyArgs)(nil)).Elem()
}

// A collection of values returned by getPagerduty.
type LookupPagerdutyResultOutput struct{ *pulumi.OutputState }

func (LookupPagerdutyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPagerdutyResult)(nil)).Elem()
}

func (o LookupPagerdutyResultOutput) ToLookupPagerdutyResultOutput() LookupPagerdutyResultOutput {
	return o
}

func (o LookupPagerdutyResultOutput) ToLookupPagerdutyResultOutputWithContext(ctx context.Context) LookupPagerdutyResultOutput {
	return o
}

// Reference to the Harness secret containing the api token.
func (o LookupPagerdutyResultOutput) ApiTokenRef() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPagerdutyResult) string { return v.ApiTokenRef }).(pulumi.StringOutput)
}

// Connect using only the delegates which have these tags.
func (o LookupPagerdutyResultOutput) DelegateSelectors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPagerdutyResult) []string { return v.DelegateSelectors }).(pulumi.StringArrayOutput)
}

// Description of the resource.
func (o LookupPagerdutyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPagerdutyResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupPagerdutyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPagerdutyResult) string { return v.Id }).(pulumi.StringOutput)
}

// Unique identifier of the resource.
func (o LookupPagerdutyResultOutput) Identifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPagerdutyResult) *string { return v.Identifier }).(pulumi.StringPtrOutput)
}

// Name of the resource.
func (o LookupPagerdutyResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPagerdutyResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Unique identifier of the organization.
func (o LookupPagerdutyResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPagerdutyResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

// Unique identifier of the project.
func (o LookupPagerdutyResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPagerdutyResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// Tags to associate with the resource. Tags should be in the form `name:value`.
func (o LookupPagerdutyResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPagerdutyResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPagerdutyResultOutput{})
}