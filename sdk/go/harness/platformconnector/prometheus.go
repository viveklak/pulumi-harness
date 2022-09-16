// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package platformconnector

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for creating a Prometheus connector.
type Prometheus struct {
	pulumi.CustomResourceState

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
	// Url of the Prometheus server.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewPrometheus registers a new resource with the given unique name, arguments, and options.
func NewPrometheus(ctx *pulumi.Context,
	name string, args *PrometheusArgs, opts ...pulumi.ResourceOption) (*Prometheus, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Identifier == nil {
		return nil, errors.New("invalid value for required argument 'Identifier'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Prometheus
	err := ctx.RegisterResource("harness:PlatformConnector/prometheus:Prometheus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrometheus gets an existing Prometheus resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrometheus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrometheusState, opts ...pulumi.ResourceOption) (*Prometheus, error) {
	var resource Prometheus
	err := ctx.ReadResource("harness:PlatformConnector/prometheus:Prometheus", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Prometheus resources.
type prometheusState struct {
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
	// Url of the Prometheus server.
	Url *string `pulumi:"url"`
}

type PrometheusState struct {
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
	// Url of the Prometheus server.
	Url pulumi.StringPtrInput
}

func (PrometheusState) ElementType() reflect.Type {
	return reflect.TypeOf((*prometheusState)(nil)).Elem()
}

type prometheusArgs struct {
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
	// Url of the Prometheus server.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a Prometheus resource.
type PrometheusArgs struct {
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
	// Url of the Prometheus server.
	Url pulumi.StringInput
}

func (PrometheusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*prometheusArgs)(nil)).Elem()
}

type PrometheusInput interface {
	pulumi.Input

	ToPrometheusOutput() PrometheusOutput
	ToPrometheusOutputWithContext(ctx context.Context) PrometheusOutput
}

func (*Prometheus) ElementType() reflect.Type {
	return reflect.TypeOf((**Prometheus)(nil)).Elem()
}

func (i *Prometheus) ToPrometheusOutput() PrometheusOutput {
	return i.ToPrometheusOutputWithContext(context.Background())
}

func (i *Prometheus) ToPrometheusOutputWithContext(ctx context.Context) PrometheusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusOutput)
}

// PrometheusArrayInput is an input type that accepts PrometheusArray and PrometheusArrayOutput values.
// You can construct a concrete instance of `PrometheusArrayInput` via:
//
//	PrometheusArray{ PrometheusArgs{...} }
type PrometheusArrayInput interface {
	pulumi.Input

	ToPrometheusArrayOutput() PrometheusArrayOutput
	ToPrometheusArrayOutputWithContext(context.Context) PrometheusArrayOutput
}

type PrometheusArray []PrometheusInput

func (PrometheusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Prometheus)(nil)).Elem()
}

func (i PrometheusArray) ToPrometheusArrayOutput() PrometheusArrayOutput {
	return i.ToPrometheusArrayOutputWithContext(context.Background())
}

func (i PrometheusArray) ToPrometheusArrayOutputWithContext(ctx context.Context) PrometheusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusArrayOutput)
}

// PrometheusMapInput is an input type that accepts PrometheusMap and PrometheusMapOutput values.
// You can construct a concrete instance of `PrometheusMapInput` via:
//
//	PrometheusMap{ "key": PrometheusArgs{...} }
type PrometheusMapInput interface {
	pulumi.Input

	ToPrometheusMapOutput() PrometheusMapOutput
	ToPrometheusMapOutputWithContext(context.Context) PrometheusMapOutput
}

type PrometheusMap map[string]PrometheusInput

func (PrometheusMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Prometheus)(nil)).Elem()
}

func (i PrometheusMap) ToPrometheusMapOutput() PrometheusMapOutput {
	return i.ToPrometheusMapOutputWithContext(context.Background())
}

func (i PrometheusMap) ToPrometheusMapOutputWithContext(ctx context.Context) PrometheusMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusMapOutput)
}

type PrometheusOutput struct{ *pulumi.OutputState }

func (PrometheusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Prometheus)(nil)).Elem()
}

func (o PrometheusOutput) ToPrometheusOutput() PrometheusOutput {
	return o
}

func (o PrometheusOutput) ToPrometheusOutputWithContext(ctx context.Context) PrometheusOutput {
	return o
}

// Connect using only the delegates which have these tags.
func (o PrometheusOutput) DelegateSelectors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Prometheus) pulumi.StringArrayOutput { return v.DelegateSelectors }).(pulumi.StringArrayOutput)
}

// Description of the resource.
func (o PrometheusOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Prometheus) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Unique identifier of the resource.
func (o PrometheusOutput) Identifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Prometheus) pulumi.StringOutput { return v.Identifier }).(pulumi.StringOutput)
}

// Name of the resource.
func (o PrometheusOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Prometheus) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Unique identifier of the organization.
func (o PrometheusOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Prometheus) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// Unique identifier of the project.
func (o PrometheusOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Prometheus) pulumi.StringPtrOutput { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// Tags to associate with the resource. Tags should be in the form `name:value`.
func (o PrometheusOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Prometheus) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// Url of the Prometheus server.
func (o PrometheusOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Prometheus) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

type PrometheusArrayOutput struct{ *pulumi.OutputState }

func (PrometheusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Prometheus)(nil)).Elem()
}

func (o PrometheusArrayOutput) ToPrometheusArrayOutput() PrometheusArrayOutput {
	return o
}

func (o PrometheusArrayOutput) ToPrometheusArrayOutputWithContext(ctx context.Context) PrometheusArrayOutput {
	return o
}

func (o PrometheusArrayOutput) Index(i pulumi.IntInput) PrometheusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Prometheus {
		return vs[0].([]*Prometheus)[vs[1].(int)]
	}).(PrometheusOutput)
}

type PrometheusMapOutput struct{ *pulumi.OutputState }

func (PrometheusMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Prometheus)(nil)).Elem()
}

func (o PrometheusMapOutput) ToPrometheusMapOutput() PrometheusMapOutput {
	return o
}

func (o PrometheusMapOutput) ToPrometheusMapOutputWithContext(ctx context.Context) PrometheusMapOutput {
	return o
}

func (o PrometheusMapOutput) MapIndex(k pulumi.StringInput) PrometheusOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Prometheus {
		return vs[0].(map[string]*Prometheus)[vs[1].(string)]
	}).(PrometheusOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrometheusInput)(nil)).Elem(), &Prometheus{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrometheusArrayInput)(nil)).Elem(), PrometheusArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrometheusMapInput)(nil)).Elem(), PrometheusMap{})
	pulumi.RegisterOutputType(PrometheusOutput{})
	pulumi.RegisterOutputType(PrometheusArrayOutput{})
	pulumi.RegisterOutputType(PrometheusMapOutput{})
}