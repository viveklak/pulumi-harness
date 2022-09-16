// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package platformconnector

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for creating a Sumologic connector.
type Sumologic struct {
	pulumi.CustomResourceState

	// Reference to the Harness secret containing the access id.
	AccessIdRef pulumi.StringOutput `pulumi:"accessIdRef"`
	// Reference to the Harness secret containing the access key.
	AccessKeyRef pulumi.StringOutput `pulumi:"accessKeyRef"`
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
	// Url of the SumoLogic server.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewSumologic registers a new resource with the given unique name, arguments, and options.
func NewSumologic(ctx *pulumi.Context,
	name string, args *SumologicArgs, opts ...pulumi.ResourceOption) (*Sumologic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessIdRef == nil {
		return nil, errors.New("invalid value for required argument 'AccessIdRef'")
	}
	if args.AccessKeyRef == nil {
		return nil, errors.New("invalid value for required argument 'AccessKeyRef'")
	}
	if args.Identifier == nil {
		return nil, errors.New("invalid value for required argument 'Identifier'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Sumologic
	err := ctx.RegisterResource("harness:PlatformConnector/sumologic:Sumologic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSumologic gets an existing Sumologic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSumologic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SumologicState, opts ...pulumi.ResourceOption) (*Sumologic, error) {
	var resource Sumologic
	err := ctx.ReadResource("harness:PlatformConnector/sumologic:Sumologic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Sumologic resources.
type sumologicState struct {
	// Reference to the Harness secret containing the access id.
	AccessIdRef *string `pulumi:"accessIdRef"`
	// Reference to the Harness secret containing the access key.
	AccessKeyRef *string `pulumi:"accessKeyRef"`
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
	// Url of the SumoLogic server.
	Url *string `pulumi:"url"`
}

type SumologicState struct {
	// Reference to the Harness secret containing the access id.
	AccessIdRef pulumi.StringPtrInput
	// Reference to the Harness secret containing the access key.
	AccessKeyRef pulumi.StringPtrInput
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
	// Url of the SumoLogic server.
	Url pulumi.StringPtrInput
}

func (SumologicState) ElementType() reflect.Type {
	return reflect.TypeOf((*sumologicState)(nil)).Elem()
}

type sumologicArgs struct {
	// Reference to the Harness secret containing the access id.
	AccessIdRef string `pulumi:"accessIdRef"`
	// Reference to the Harness secret containing the access key.
	AccessKeyRef string `pulumi:"accessKeyRef"`
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
	// Url of the SumoLogic server.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a Sumologic resource.
type SumologicArgs struct {
	// Reference to the Harness secret containing the access id.
	AccessIdRef pulumi.StringInput
	// Reference to the Harness secret containing the access key.
	AccessKeyRef pulumi.StringInput
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
	// Url of the SumoLogic server.
	Url pulumi.StringInput
}

func (SumologicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sumologicArgs)(nil)).Elem()
}

type SumologicInput interface {
	pulumi.Input

	ToSumologicOutput() SumologicOutput
	ToSumologicOutputWithContext(ctx context.Context) SumologicOutput
}

func (*Sumologic) ElementType() reflect.Type {
	return reflect.TypeOf((**Sumologic)(nil)).Elem()
}

func (i *Sumologic) ToSumologicOutput() SumologicOutput {
	return i.ToSumologicOutputWithContext(context.Background())
}

func (i *Sumologic) ToSumologicOutputWithContext(ctx context.Context) SumologicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SumologicOutput)
}

// SumologicArrayInput is an input type that accepts SumologicArray and SumologicArrayOutput values.
// You can construct a concrete instance of `SumologicArrayInput` via:
//
//	SumologicArray{ SumologicArgs{...} }
type SumologicArrayInput interface {
	pulumi.Input

	ToSumologicArrayOutput() SumologicArrayOutput
	ToSumologicArrayOutputWithContext(context.Context) SumologicArrayOutput
}

type SumologicArray []SumologicInput

func (SumologicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Sumologic)(nil)).Elem()
}

func (i SumologicArray) ToSumologicArrayOutput() SumologicArrayOutput {
	return i.ToSumologicArrayOutputWithContext(context.Background())
}

func (i SumologicArray) ToSumologicArrayOutputWithContext(ctx context.Context) SumologicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SumologicArrayOutput)
}

// SumologicMapInput is an input type that accepts SumologicMap and SumologicMapOutput values.
// You can construct a concrete instance of `SumologicMapInput` via:
//
//	SumologicMap{ "key": SumologicArgs{...} }
type SumologicMapInput interface {
	pulumi.Input

	ToSumologicMapOutput() SumologicMapOutput
	ToSumologicMapOutputWithContext(context.Context) SumologicMapOutput
}

type SumologicMap map[string]SumologicInput

func (SumologicMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Sumologic)(nil)).Elem()
}

func (i SumologicMap) ToSumologicMapOutput() SumologicMapOutput {
	return i.ToSumologicMapOutputWithContext(context.Background())
}

func (i SumologicMap) ToSumologicMapOutputWithContext(ctx context.Context) SumologicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SumologicMapOutput)
}

type SumologicOutput struct{ *pulumi.OutputState }

func (SumologicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sumologic)(nil)).Elem()
}

func (o SumologicOutput) ToSumologicOutput() SumologicOutput {
	return o
}

func (o SumologicOutput) ToSumologicOutputWithContext(ctx context.Context) SumologicOutput {
	return o
}

// Reference to the Harness secret containing the access id.
func (o SumologicOutput) AccessIdRef() pulumi.StringOutput {
	return o.ApplyT(func(v *Sumologic) pulumi.StringOutput { return v.AccessIdRef }).(pulumi.StringOutput)
}

// Reference to the Harness secret containing the access key.
func (o SumologicOutput) AccessKeyRef() pulumi.StringOutput {
	return o.ApplyT(func(v *Sumologic) pulumi.StringOutput { return v.AccessKeyRef }).(pulumi.StringOutput)
}

// Connect using only the delegates which have these tags.
func (o SumologicOutput) DelegateSelectors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Sumologic) pulumi.StringArrayOutput { return v.DelegateSelectors }).(pulumi.StringArrayOutput)
}

// Description of the resource.
func (o SumologicOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sumologic) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Unique identifier of the resource.
func (o SumologicOutput) Identifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Sumologic) pulumi.StringOutput { return v.Identifier }).(pulumi.StringOutput)
}

// Name of the resource.
func (o SumologicOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Sumologic) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Unique identifier of the organization.
func (o SumologicOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sumologic) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// Unique identifier of the project.
func (o SumologicOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sumologic) pulumi.StringPtrOutput { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// Tags to associate with the resource. Tags should be in the form `name:value`.
func (o SumologicOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Sumologic) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Url of the SumoLogic server.
func (o SumologicOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Sumologic) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type SumologicArrayOutput struct{ *pulumi.OutputState }

func (SumologicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Sumologic)(nil)).Elem()
}

func (o SumologicArrayOutput) ToSumologicArrayOutput() SumologicArrayOutput {
	return o
}

func (o SumologicArrayOutput) ToSumologicArrayOutputWithContext(ctx context.Context) SumologicArrayOutput {
	return o
}

func (o SumologicArrayOutput) Index(i pulumi.IntInput) SumologicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Sumologic {
		return vs[0].([]*Sumologic)[vs[1].(int)]
	}).(SumologicOutput)
}

type SumologicMapOutput struct{ *pulumi.OutputState }

func (SumologicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Sumologic)(nil)).Elem()
}

func (o SumologicMapOutput) ToSumologicMapOutput() SumologicMapOutput {
	return o
}

func (o SumologicMapOutput) ToSumologicMapOutputWithContext(ctx context.Context) SumologicMapOutput {
	return o
}

func (o SumologicMapOutput) MapIndex(k pulumi.StringInput) SumologicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Sumologic {
		return vs[0].(map[string]*Sumologic)[vs[1].(string)]
	}).(SumologicOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SumologicInput)(nil)).Elem(), &Sumologic{})
	pulumi.RegisterInputType(reflect.TypeOf((*SumologicArrayInput)(nil)).Elem(), SumologicArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SumologicMapInput)(nil)).Elem(), SumologicMap{})
	pulumi.RegisterOutputType(SumologicOutput{})
	pulumi.RegisterOutputType(SumologicArrayOutput{})
	pulumi.RegisterOutputType(SumologicMapOutput{})
}