// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package service

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for creating a Kubernetes service. This resource uses the config-as-code API's. When updating the `name` or `path` of this resource you should typically also set the `createBeforeDestroy = true` lifecycle setting.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/lbrlabs/pulumi-harness/sdk/go/harness"
//	"github.com/lbrlabs/pulumi-harness/sdk/go/harness/service"
//	"github.com/pulumi/pulumi-harness/sdk/go/harness/service"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleApplication, err := harness.NewApplication(ctx, "exampleApplication", nil)
//			if err != nil {
//				return err
//			}
//			_, err = service.NewKubernetes(ctx, "exampleKubernetes", &service.KubernetesArgs{
//				AppId:       exampleApplication.ID(),
//				HelmVersion: pulumi.String("V3"),
//				Description: pulumi.String("Service for deploying Kubernetes manifests"),
//				Variables: service.KubernetesVariableArray{
//					&service.KubernetesVariableArgs{
//						Name:  pulumi.String("test"),
//						Value: pulumi.String("test_value"),
//						Type:  pulumi.String("TEXT"),
//					},
//					&service.KubernetesVariableArgs{
//						Name:  pulumi.String("test2"),
//						Value: pulumi.String("test_value2"),
//						Type:  pulumi.String("TEXT"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # Import using the Harness application id and service id
//
// ```sh
//
//	$ pulumi import harness:service/kubernetes:Kubernetes example <app_id>/<svc_id>
//
// ```
type Kubernetes struct {
	pulumi.CustomResourceState

	// The id of the application the service belongs to
	AppId pulumi.StringOutput `pulumi:"appId"`
	// Description of th service
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The version of Helm to use. Options are `V2` and `V3`. Defaults to 'V2'. Only used when `type` is `KUBERNETES` or `HELM`.
	HelmVersion pulumi.StringPtrOutput `pulumi:"helmVersion"`
	// Name of the service
	Name pulumi.StringOutput `pulumi:"name"`
	// Variables to be used in the service
	Variables KubernetesVariableArrayOutput `pulumi:"variables"`
}

// NewKubernetes registers a new resource with the given unique name, arguments, and options.
func NewKubernetes(ctx *pulumi.Context,
	name string, args *KubernetesArgs, opts ...pulumi.ResourceOption) (*Kubernetes, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Kubernetes
	err := ctx.RegisterResource("harness:service/kubernetes:Kubernetes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKubernetes gets an existing Kubernetes resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKubernetes(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubernetesState, opts ...pulumi.ResourceOption) (*Kubernetes, error) {
	var resource Kubernetes
	err := ctx.ReadResource("harness:service/kubernetes:Kubernetes", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Kubernetes resources.
type kubernetesState struct {
	// The id of the application the service belongs to
	AppId *string `pulumi:"appId"`
	// Description of th service
	Description *string `pulumi:"description"`
	// The version of Helm to use. Options are `V2` and `V3`. Defaults to 'V2'. Only used when `type` is `KUBERNETES` or `HELM`.
	HelmVersion *string `pulumi:"helmVersion"`
	// Name of the service
	Name *string `pulumi:"name"`
	// Variables to be used in the service
	Variables []KubernetesVariable `pulumi:"variables"`
}

type KubernetesState struct {
	// The id of the application the service belongs to
	AppId pulumi.StringPtrInput
	// Description of th service
	Description pulumi.StringPtrInput
	// The version of Helm to use. Options are `V2` and `V3`. Defaults to 'V2'. Only used when `type` is `KUBERNETES` or `HELM`.
	HelmVersion pulumi.StringPtrInput
	// Name of the service
	Name pulumi.StringPtrInput
	// Variables to be used in the service
	Variables KubernetesVariableArrayInput
}

func (KubernetesState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesState)(nil)).Elem()
}

type kubernetesArgs struct {
	// The id of the application the service belongs to
	AppId string `pulumi:"appId"`
	// Description of th service
	Description *string `pulumi:"description"`
	// The version of Helm to use. Options are `V2` and `V3`. Defaults to 'V2'. Only used when `type` is `KUBERNETES` or `HELM`.
	HelmVersion *string `pulumi:"helmVersion"`
	// Name of the service
	Name *string `pulumi:"name"`
	// Variables to be used in the service
	Variables []KubernetesVariable `pulumi:"variables"`
}

// The set of arguments for constructing a Kubernetes resource.
type KubernetesArgs struct {
	// The id of the application the service belongs to
	AppId pulumi.StringInput
	// Description of th service
	Description pulumi.StringPtrInput
	// The version of Helm to use. Options are `V2` and `V3`. Defaults to 'V2'. Only used when `type` is `KUBERNETES` or `HELM`.
	HelmVersion pulumi.StringPtrInput
	// Name of the service
	Name pulumi.StringPtrInput
	// Variables to be used in the service
	Variables KubernetesVariableArrayInput
}

func (KubernetesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubernetesArgs)(nil)).Elem()
}

type KubernetesInput interface {
	pulumi.Input

	ToKubernetesOutput() KubernetesOutput
	ToKubernetesOutputWithContext(ctx context.Context) KubernetesOutput
}

func (*Kubernetes) ElementType() reflect.Type {
	return reflect.TypeOf((**Kubernetes)(nil)).Elem()
}

func (i *Kubernetes) ToKubernetesOutput() KubernetesOutput {
	return i.ToKubernetesOutputWithContext(context.Background())
}

func (i *Kubernetes) ToKubernetesOutputWithContext(ctx context.Context) KubernetesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesOutput)
}

// KubernetesArrayInput is an input type that accepts KubernetesArray and KubernetesArrayOutput values.
// You can construct a concrete instance of `KubernetesArrayInput` via:
//
//	KubernetesArray{ KubernetesArgs{...} }
type KubernetesArrayInput interface {
	pulumi.Input

	ToKubernetesArrayOutput() KubernetesArrayOutput
	ToKubernetesArrayOutputWithContext(context.Context) KubernetesArrayOutput
}

type KubernetesArray []KubernetesInput

func (KubernetesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Kubernetes)(nil)).Elem()
}

func (i KubernetesArray) ToKubernetesArrayOutput() KubernetesArrayOutput {
	return i.ToKubernetesArrayOutputWithContext(context.Background())
}

func (i KubernetesArray) ToKubernetesArrayOutputWithContext(ctx context.Context) KubernetesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesArrayOutput)
}

// KubernetesMapInput is an input type that accepts KubernetesMap and KubernetesMapOutput values.
// You can construct a concrete instance of `KubernetesMapInput` via:
//
//	KubernetesMap{ "key": KubernetesArgs{...} }
type KubernetesMapInput interface {
	pulumi.Input

	ToKubernetesMapOutput() KubernetesMapOutput
	ToKubernetesMapOutputWithContext(context.Context) KubernetesMapOutput
}

type KubernetesMap map[string]KubernetesInput

func (KubernetesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Kubernetes)(nil)).Elem()
}

func (i KubernetesMap) ToKubernetesMapOutput() KubernetesMapOutput {
	return i.ToKubernetesMapOutputWithContext(context.Background())
}

func (i KubernetesMap) ToKubernetesMapOutputWithContext(ctx context.Context) KubernetesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubernetesMapOutput)
}

type KubernetesOutput struct{ *pulumi.OutputState }

func (KubernetesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Kubernetes)(nil)).Elem()
}

func (o KubernetesOutput) ToKubernetesOutput() KubernetesOutput {
	return o
}

func (o KubernetesOutput) ToKubernetesOutputWithContext(ctx context.Context) KubernetesOutput {
	return o
}

// The id of the application the service belongs to
func (o KubernetesOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *Kubernetes) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

// Description of th service
func (o KubernetesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Kubernetes) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The version of Helm to use. Options are `V2` and `V3`. Defaults to 'V2'. Only used when `type` is `KUBERNETES` or `HELM`.
func (o KubernetesOutput) HelmVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Kubernetes) pulumi.StringPtrOutput { return v.HelmVersion }).(pulumi.StringPtrOutput)
}

// Name of the service
func (o KubernetesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Kubernetes) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Variables to be used in the service
func (o KubernetesOutput) Variables() KubernetesVariableArrayOutput {
	return o.ApplyT(func(v *Kubernetes) KubernetesVariableArrayOutput { return v.Variables }).(KubernetesVariableArrayOutput)
}

type KubernetesArrayOutput struct{ *pulumi.OutputState }

func (KubernetesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Kubernetes)(nil)).Elem()
}

func (o KubernetesArrayOutput) ToKubernetesArrayOutput() KubernetesArrayOutput {
	return o
}

func (o KubernetesArrayOutput) ToKubernetesArrayOutputWithContext(ctx context.Context) KubernetesArrayOutput {
	return o
}

func (o KubernetesArrayOutput) Index(i pulumi.IntInput) KubernetesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Kubernetes {
		return vs[0].([]*Kubernetes)[vs[1].(int)]
	}).(KubernetesOutput)
}

type KubernetesMapOutput struct{ *pulumi.OutputState }

func (KubernetesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Kubernetes)(nil)).Elem()
}

func (o KubernetesMapOutput) ToKubernetesMapOutput() KubernetesMapOutput {
	return o
}

func (o KubernetesMapOutput) ToKubernetesMapOutputWithContext(ctx context.Context) KubernetesMapOutput {
	return o
}

func (o KubernetesMapOutput) MapIndex(k pulumi.StringInput) KubernetesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Kubernetes {
		return vs[0].(map[string]*Kubernetes)[vs[1].(string)]
	}).(KubernetesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesInput)(nil)).Elem(), &Kubernetes{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesArrayInput)(nil)).Elem(), KubernetesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubernetesMapInput)(nil)).Elem(), KubernetesMap{})
	pulumi.RegisterOutputType(KubernetesOutput{})
	pulumi.RegisterOutputType(KubernetesArrayOutput{})
	pulumi.RegisterOutputType(KubernetesMapOutput{})
}