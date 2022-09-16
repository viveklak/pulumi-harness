// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package platformconnector

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for creating a Dynatrace connector.
type Dynatrace struct {
	pulumi.CustomResourceState

	// The reference to the Harness secret containing the api token.
	ApiTokenRef pulumi.StringOutput `pulumi:"apiTokenRef"`
	// Connect using only the delegates which have these tags.
	DelegateSelectors pulumi.StringArrayOutput `pulumi:"delegateSelectors"`
	// Description of the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Unique identifier of the resource.
	Identifier pulumi.StringOutput `pulumi:"identifier"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrOutput `pulumi:"projectId"`
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Url of the Dynatrace server.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewDynatrace registers a new resource with the given unique name, arguments, and options.
func NewDynatrace(ctx *pulumi.Context,
	name string, args *DynatraceArgs, opts ...pulumi.ResourceOption) (*Dynatrace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiTokenRef == nil {
		return nil, errors.New("invalid value for required argument 'ApiTokenRef'")
	}
	if args.Identifier == nil {
		return nil, errors.New("invalid value for required argument 'Identifier'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Dynatrace
	err := ctx.RegisterResource("harness:PlatformConnector/dynatrace:Dynatrace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDynatrace gets an existing Dynatrace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDynatrace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DynatraceState, opts ...pulumi.ResourceOption) (*Dynatrace, error) {
	var resource Dynatrace
	err := ctx.ReadResource("harness:PlatformConnector/dynatrace:Dynatrace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dynatrace resources.
type dynatraceState struct {
	// The reference to the Harness secret containing the api token.
	ApiTokenRef *string `pulumi:"apiTokenRef"`
	// Connect using only the delegates which have these tags.
	DelegateSelectors []string `pulumi:"delegateSelectors"`
	// Description of the resource.
	Description *string `pulumi:"description"`
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
	// Url of the Dynatrace server.
	Url *string `pulumi:"url"`
}

type DynatraceState struct {
	// The reference to the Harness secret containing the api token.
	ApiTokenRef pulumi.StringPtrInput
	// Connect using only the delegates which have these tags.
	DelegateSelectors pulumi.StringArrayInput
	// Description of the resource.
	Description pulumi.StringPtrInput
	// Unique identifier of the resource.
	Identifier pulumi.StringPtrInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// Unique identifier of the organization.
	OrgId pulumi.StringPtrInput
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrInput
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags pulumi.StringArrayInput
	// Url of the Dynatrace server.
	Url pulumi.StringPtrInput
}

func (DynatraceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dynatraceState)(nil)).Elem()
}

type dynatraceArgs struct {
	// The reference to the Harness secret containing the api token.
	ApiTokenRef string `pulumi:"apiTokenRef"`
	// Connect using only the delegates which have these tags.
	DelegateSelectors []string `pulumi:"delegateSelectors"`
	// Description of the resource.
	Description *string `pulumi:"description"`
	// Unique identifier of the resource.
	Identifier string `pulumi:"identifier"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId *string `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId *string `pulumi:"projectId"`
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags []string `pulumi:"tags"`
	// Url of the Dynatrace server.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a Dynatrace resource.
type DynatraceArgs struct {
	// The reference to the Harness secret containing the api token.
	ApiTokenRef pulumi.StringInput
	// Connect using only the delegates which have these tags.
	DelegateSelectors pulumi.StringArrayInput
	// Description of the resource.
	Description pulumi.StringPtrInput
	// Unique identifier of the resource.
	Identifier pulumi.StringInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// Unique identifier of the organization.
	OrgId pulumi.StringPtrInput
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrInput
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags pulumi.StringArrayInput
	// Url of the Dynatrace server.
	Url pulumi.StringInput
}

func (DynatraceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dynatraceArgs)(nil)).Elem()
}

type DynatraceInput interface {
	pulumi.Input

	ToDynatraceOutput() DynatraceOutput
	ToDynatraceOutputWithContext(ctx context.Context) DynatraceOutput
}

func (*Dynatrace) ElementType() reflect.Type {
	return reflect.TypeOf((**Dynatrace)(nil)).Elem()
}

func (i *Dynatrace) ToDynatraceOutput() DynatraceOutput {
	return i.ToDynatraceOutputWithContext(context.Background())
}

func (i *Dynatrace) ToDynatraceOutputWithContext(ctx context.Context) DynatraceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DynatraceOutput)
}

// DynatraceArrayInput is an input type that accepts DynatraceArray and DynatraceArrayOutput values.
// You can construct a concrete instance of `DynatraceArrayInput` via:
//
//	DynatraceArray{ DynatraceArgs{...} }
type DynatraceArrayInput interface {
	pulumi.Input

	ToDynatraceArrayOutput() DynatraceArrayOutput
	ToDynatraceArrayOutputWithContext(context.Context) DynatraceArrayOutput
}

type DynatraceArray []DynatraceInput

func (DynatraceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dynatrace)(nil)).Elem()
}

func (i DynatraceArray) ToDynatraceArrayOutput() DynatraceArrayOutput {
	return i.ToDynatraceArrayOutputWithContext(context.Background())
}

func (i DynatraceArray) ToDynatraceArrayOutputWithContext(ctx context.Context) DynatraceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DynatraceArrayOutput)
}

// DynatraceMapInput is an input type that accepts DynatraceMap and DynatraceMapOutput values.
// You can construct a concrete instance of `DynatraceMapInput` via:
//
//	DynatraceMap{ "key": DynatraceArgs{...} }
type DynatraceMapInput interface {
	pulumi.Input

	ToDynatraceMapOutput() DynatraceMapOutput
	ToDynatraceMapOutputWithContext(context.Context) DynatraceMapOutput
}

type DynatraceMap map[string]DynatraceInput

func (DynatraceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dynatrace)(nil)).Elem()
}

func (i DynatraceMap) ToDynatraceMapOutput() DynatraceMapOutput {
	return i.ToDynatraceMapOutputWithContext(context.Background())
}

func (i DynatraceMap) ToDynatraceMapOutputWithContext(ctx context.Context) DynatraceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DynatraceMapOutput)
}

type DynatraceOutput struct{ *pulumi.OutputState }

func (DynatraceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dynatrace)(nil)).Elem()
}

func (o DynatraceOutput) ToDynatraceOutput() DynatraceOutput {
	return o
}

func (o DynatraceOutput) ToDynatraceOutputWithContext(ctx context.Context) DynatraceOutput {
	return o
}

// The reference to the Harness secret containing the api token.
func (o DynatraceOutput) ApiTokenRef() pulumi.StringOutput {
	return o.ApplyT(func(v *Dynatrace) pulumi.StringOutput { return v.ApiTokenRef }).(pulumi.StringOutput)
}

// Connect using only the delegates which have these tags.
func (o DynatraceOutput) DelegateSelectors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Dynatrace) pulumi.StringArrayOutput { return v.DelegateSelectors }).(pulumi.StringArrayOutput)
}

// Description of the resource.
func (o DynatraceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dynatrace) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Unique identifier of the resource.
func (o DynatraceOutput) Identifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Dynatrace) pulumi.StringOutput { return v.Identifier }).(pulumi.StringOutput)
}

// Name of the resource.
func (o DynatraceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Dynatrace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Unique identifier of the organization.
func (o DynatraceOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dynatrace) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// Unique identifier of the project.
func (o DynatraceOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dynatrace) pulumi.StringPtrOutput { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// Tags to associate with the resource. Tags should be in the form `name:value`.
func (o DynatraceOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Dynatrace) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Url of the Dynatrace server.
func (o DynatraceOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Dynatrace) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type DynatraceArrayOutput struct{ *pulumi.OutputState }

func (DynatraceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dynatrace)(nil)).Elem()
}

func (o DynatraceArrayOutput) ToDynatraceArrayOutput() DynatraceArrayOutput {
	return o
}

func (o DynatraceArrayOutput) ToDynatraceArrayOutputWithContext(ctx context.Context) DynatraceArrayOutput {
	return o
}

func (o DynatraceArrayOutput) Index(i pulumi.IntInput) DynatraceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Dynatrace {
		return vs[0].([]*Dynatrace)[vs[1].(int)]
	}).(DynatraceOutput)
}

type DynatraceMapOutput struct{ *pulumi.OutputState }

func (DynatraceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dynatrace)(nil)).Elem()
}

func (o DynatraceMapOutput) ToDynatraceMapOutput() DynatraceMapOutput {
	return o
}

func (o DynatraceMapOutput) ToDynatraceMapOutputWithContext(ctx context.Context) DynatraceMapOutput {
	return o
}

func (o DynatraceMapOutput) MapIndex(k pulumi.StringInput) DynatraceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Dynatrace {
		return vs[0].(map[string]*Dynatrace)[vs[1].(string)]
	}).(DynatraceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DynatraceInput)(nil)).Elem(), &Dynatrace{})
	pulumi.RegisterInputType(reflect.TypeOf((*DynatraceArrayInput)(nil)).Elem(), DynatraceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DynatraceMapInput)(nil)).Elem(), DynatraceMap{})
	pulumi.RegisterOutputType(DynatraceOutput{})
	pulumi.RegisterOutputType(DynatraceArrayOutput{})
	pulumi.RegisterOutputType(DynatraceMapOutput{})
}