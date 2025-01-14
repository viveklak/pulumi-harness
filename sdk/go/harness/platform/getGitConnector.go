// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package platform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Datasource for looking up a Git connector.
func LookupGitConnector(ctx *pulumi.Context, args *LookupGitConnectorArgs, opts ...pulumi.InvokeOption) (*LookupGitConnectorResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupGitConnectorResult
	err := ctx.Invoke("harness:platform/getGitConnector:getGitConnector", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGitConnector.
type LookupGitConnectorArgs struct {
	// Unique identifier of the resource.
	Identifier *string `pulumi:"identifier"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId *string `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId *string `pulumi:"projectId"`
}

// A collection of values returned by getGitConnector.
type LookupGitConnectorResult struct {
	// Whether the connection we're making is to a git repository or a git account. Valid values are Account, Repo.
	ConnectionType string `pulumi:"connectionType"`
	// Credentials to use for the connection.
	Credentials []GetGitConnectorCredential `pulumi:"credentials"`
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
	// Url of the git repository or account.
	Url string `pulumi:"url"`
	// Repository to test the connection with. This is only used when `connectionType` is `Account`.
	ValidationRepo string `pulumi:"validationRepo"`
}

func LookupGitConnectorOutput(ctx *pulumi.Context, args LookupGitConnectorOutputArgs, opts ...pulumi.InvokeOption) LookupGitConnectorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGitConnectorResult, error) {
			args := v.(LookupGitConnectorArgs)
			r, err := LookupGitConnector(ctx, &args, opts...)
			var s LookupGitConnectorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGitConnectorResultOutput)
}

// A collection of arguments for invoking getGitConnector.
type LookupGitConnectorOutputArgs struct {
	// Unique identifier of the resource.
	Identifier pulumi.StringPtrInput `pulumi:"identifier"`
	// Name of the resource.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrInput `pulumi:"projectId"`
}

func (LookupGitConnectorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGitConnectorArgs)(nil)).Elem()
}

// A collection of values returned by getGitConnector.
type LookupGitConnectorResultOutput struct{ *pulumi.OutputState }

func (LookupGitConnectorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGitConnectorResult)(nil)).Elem()
}

func (o LookupGitConnectorResultOutput) ToLookupGitConnectorResultOutput() LookupGitConnectorResultOutput {
	return o
}

func (o LookupGitConnectorResultOutput) ToLookupGitConnectorResultOutputWithContext(ctx context.Context) LookupGitConnectorResultOutput {
	return o
}

// Whether the connection we're making is to a git repository or a git account. Valid values are Account, Repo.
func (o LookupGitConnectorResultOutput) ConnectionType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGitConnectorResult) string { return v.ConnectionType }).(pulumi.StringOutput)
}

// Credentials to use for the connection.
func (o LookupGitConnectorResultOutput) Credentials() GetGitConnectorCredentialArrayOutput {
	return o.ApplyT(func(v LookupGitConnectorResult) []GetGitConnectorCredential { return v.Credentials }).(GetGitConnectorCredentialArrayOutput)
}

// Connect using only the delegates which have these tags.
func (o LookupGitConnectorResultOutput) DelegateSelectors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGitConnectorResult) []string { return v.DelegateSelectors }).(pulumi.StringArrayOutput)
}

// Description of the resource.
func (o LookupGitConnectorResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGitConnectorResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupGitConnectorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGitConnectorResult) string { return v.Id }).(pulumi.StringOutput)
}

// Unique identifier of the resource.
func (o LookupGitConnectorResultOutput) Identifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGitConnectorResult) *string { return v.Identifier }).(pulumi.StringPtrOutput)
}

// Name of the resource.
func (o LookupGitConnectorResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGitConnectorResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Unique identifier of the organization.
func (o LookupGitConnectorResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGitConnectorResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

// Unique identifier of the project.
func (o LookupGitConnectorResultOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGitConnectorResult) *string { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// Tags to associate with the resource. Tags should be in the form `name:value`.
func (o LookupGitConnectorResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupGitConnectorResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// Url of the git repository or account.
func (o LookupGitConnectorResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGitConnectorResult) string { return v.Url }).(pulumi.StringOutput)
}

// Repository to test the connection with. This is only used when `connectionType` is `Account`.
func (o LookupGitConnectorResultOutput) ValidationRepo() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGitConnectorResult) string { return v.ValidationRepo }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGitConnectorResultOutput{})
}
