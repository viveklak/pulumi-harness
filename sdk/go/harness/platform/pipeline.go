// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package platform

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for creating a Harness pipeline.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/lbrlabs/pulumi-harness/sdk/go/harness/platform"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := platform.NewPipeline(ctx, "example", &platform.PipelineArgs{
//				Identifier: pulumi.String("identifier"),
//				OrgId:      pulumi.Any(harness_platform_project.Test.Org_id),
//				ProjectId:  pulumi.Any(harness_platform_project.Test.Id),
//				Yaml: pulumi.String(fmt.Sprintf(`pipeline:
//	    name: name
//	    identifier: identifier
//	    allowStageExecutions: false
//	    projectIdentifier: projectIdentifier
//	    orgIdentifier: orgIdentifier
//	    tags: {}
//	    stages:
//	        - stage:
//	            name: dep
//	            identifier: dep
//	            description: ""
//	            type: Deployment
//	            spec:
//	                serviceConfig:
//	                    serviceRef: service
//	                    serviceDefinition:
//	                        type: Kubernetes
//	                        spec:
//	                            variables: []
//	                infrastructure:
//	                    environmentRef: testenv
//	                    infrastructureDefinition:
//	                        type: KubernetesDirect
//	                        spec:
//	                            connectorRef: testconf
//	                            namespace: test
//	                            releaseName: release-<+INFRA_KEY>
//	                    allowSimultaneousDeployments: false
//	                execution:
//	                    steps:
//	                        - stepGroup:
//	                                name: Canary Deployment
//	                                identifier: canaryDepoyment
//	                                steps:
//	                                    - step:
//	                                        name: Canary Deployment
//	                                        identifier: canaryDeployment
//	                                        type: K8sCanaryDeploy
//	                                        timeout: 10m
//	                                        spec:
//	                                            instanceSelection:
//	                                                type: Count
//	                                                spec:
//	                                                    count: 1
//	                                            skipDryRun: false
//	                                    - step:
//	                                        name: Canary Delete
//	                                        identifier: canaryDelete
//	                                        type: K8sCanaryDelete
//	                                        timeout: 10m
//	                                        spec: {}
//	                                rollbackSteps:
//	                                    - step:
//	                                        name: Canary Delete
//	                                        identifier: rollbackCanaryDelete
//	                                        type: K8sCanaryDelete
//	                                        timeout: 10m
//	                                        spec: {}
//	                        - stepGroup:
//	                                name: Primary Deployment
//	                                identifier: primaryDepoyment
//	                                steps:
//	                                    - step:
//	                                        name: Rolling Deployment
//	                                        identifier: rollingDeployment
//	                                        type: K8sRollingDeploy
//	                                        timeout: 10m
//	                                        spec:
//	                                            skipDryRun: false
//	                                rollbackSteps:
//	                                    - step:
//	                                        name: Rolling Rollback
//	                                        identifier: rollingRollback
//	                                        type: K8sRollingRollback
//	                                        timeout: 10m
//	                                        spec: {}
//	                    rollbackSteps: []
//	            tags: {}
//	            failureStrategies:
//	                - onFailure:
//	                        errors:
//	                            - AllErrors
//	                        action:
//	                            type: StageRollback
//
// `)),
//
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type Pipeline struct {
	pulumi.CustomResourceState

	// Description of the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Unique identifier of the resource.
	Identifier pulumi.StringOutput `pulumi:"identifier"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// YAML of the pipeline.
	Yaml pulumi.StringOutput `pulumi:"yaml"`
}

// NewPipeline registers a new resource with the given unique name, arguments, and options.
func NewPipeline(ctx *pulumi.Context,
	name string, args *PipelineArgs, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Identifier == nil {
		return nil, errors.New("invalid value for required argument 'Identifier'")
	}
	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.Yaml == nil {
		return nil, errors.New("invalid value for required argument 'Yaml'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Pipeline
	err := ctx.RegisterResource("harness:platform/pipeline:Pipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPipeline gets an existing Pipeline resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineState, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	var resource Pipeline
	err := ctx.ReadResource("harness:platform/pipeline:Pipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Pipeline resources.
type pipelineState struct {
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
	// YAML of the pipeline.
	Yaml *string `pulumi:"yaml"`
}

type PipelineState struct {
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
	// YAML of the pipeline.
	Yaml pulumi.StringPtrInput
}

func (PipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineState)(nil)).Elem()
}

type pipelineArgs struct {
	// Description of the resource.
	Description *string `pulumi:"description"`
	// Unique identifier of the resource.
	Identifier string `pulumi:"identifier"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Unique identifier of the organization.
	OrgId string `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId string `pulumi:"projectId"`
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags []string `pulumi:"tags"`
	// YAML of the pipeline.
	Yaml string `pulumi:"yaml"`
}

// The set of arguments for constructing a Pipeline resource.
type PipelineArgs struct {
	// Description of the resource.
	Description pulumi.StringPtrInput
	// Unique identifier of the resource.
	Identifier pulumi.StringInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// Unique identifier of the organization.
	OrgId pulumi.StringInput
	// Unique identifier of the project.
	ProjectId pulumi.StringInput
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags pulumi.StringArrayInput
	// YAML of the pipeline.
	Yaml pulumi.StringInput
}

func (PipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineArgs)(nil)).Elem()
}

type PipelineInput interface {
	pulumi.Input

	ToPipelineOutput() PipelineOutput
	ToPipelineOutputWithContext(ctx context.Context) PipelineOutput
}

func (*Pipeline) ElementType() reflect.Type {
	return reflect.TypeOf((**Pipeline)(nil)).Elem()
}

func (i *Pipeline) ToPipelineOutput() PipelineOutput {
	return i.ToPipelineOutputWithContext(context.Background())
}

func (i *Pipeline) ToPipelineOutputWithContext(ctx context.Context) PipelineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineOutput)
}

// PipelineArrayInput is an input type that accepts PipelineArray and PipelineArrayOutput values.
// You can construct a concrete instance of `PipelineArrayInput` via:
//
//	PipelineArray{ PipelineArgs{...} }
type PipelineArrayInput interface {
	pulumi.Input

	ToPipelineArrayOutput() PipelineArrayOutput
	ToPipelineArrayOutputWithContext(context.Context) PipelineArrayOutput
}

type PipelineArray []PipelineInput

func (PipelineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Pipeline)(nil)).Elem()
}

func (i PipelineArray) ToPipelineArrayOutput() PipelineArrayOutput {
	return i.ToPipelineArrayOutputWithContext(context.Background())
}

func (i PipelineArray) ToPipelineArrayOutputWithContext(ctx context.Context) PipelineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineArrayOutput)
}

// PipelineMapInput is an input type that accepts PipelineMap and PipelineMapOutput values.
// You can construct a concrete instance of `PipelineMapInput` via:
//
//	PipelineMap{ "key": PipelineArgs{...} }
type PipelineMapInput interface {
	pulumi.Input

	ToPipelineMapOutput() PipelineMapOutput
	ToPipelineMapOutputWithContext(context.Context) PipelineMapOutput
}

type PipelineMap map[string]PipelineInput

func (PipelineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Pipeline)(nil)).Elem()
}

func (i PipelineMap) ToPipelineMapOutput() PipelineMapOutput {
	return i.ToPipelineMapOutputWithContext(context.Background())
}

func (i PipelineMap) ToPipelineMapOutputWithContext(ctx context.Context) PipelineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineMapOutput)
}

type PipelineOutput struct{ *pulumi.OutputState }

func (PipelineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pipeline)(nil)).Elem()
}

func (o PipelineOutput) ToPipelineOutput() PipelineOutput {
	return o
}

func (o PipelineOutput) ToPipelineOutputWithContext(ctx context.Context) PipelineOutput {
	return o
}

// Description of the resource.
func (o PipelineOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Unique identifier of the resource.
func (o PipelineOutput) Identifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringOutput { return v.Identifier }).(pulumi.StringOutput)
}

// Name of the resource.
func (o PipelineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Unique identifier of the organization.
func (o PipelineOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

// Unique identifier of the project.
func (o PipelineOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Tags to associate with the resource. Tags should be in the form `name:value`.
func (o PipelineOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// YAML of the pipeline.
func (o PipelineOutput) Yaml() pulumi.StringOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringOutput { return v.Yaml }).(pulumi.StringOutput)
}

type PipelineArrayOutput struct{ *pulumi.OutputState }

func (PipelineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Pipeline)(nil)).Elem()
}

func (o PipelineArrayOutput) ToPipelineArrayOutput() PipelineArrayOutput {
	return o
}

func (o PipelineArrayOutput) ToPipelineArrayOutputWithContext(ctx context.Context) PipelineArrayOutput {
	return o
}

func (o PipelineArrayOutput) Index(i pulumi.IntInput) PipelineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Pipeline {
		return vs[0].([]*Pipeline)[vs[1].(int)]
	}).(PipelineOutput)
}

type PipelineMapOutput struct{ *pulumi.OutputState }

func (PipelineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Pipeline)(nil)).Elem()
}

func (o PipelineMapOutput) ToPipelineMapOutput() PipelineMapOutput {
	return o
}

func (o PipelineMapOutput) ToPipelineMapOutputWithContext(ctx context.Context) PipelineMapOutput {
	return o
}

func (o PipelineMapOutput) MapIndex(k pulumi.StringInput) PipelineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Pipeline {
		return vs[0].(map[string]*Pipeline)[vs[1].(string)]
	}).(PipelineOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineInput)(nil)).Elem(), &Pipeline{})
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineArrayInput)(nil)).Elem(), PipelineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PipelineMapInput)(nil)).Elem(), PipelineMap{})
	pulumi.RegisterOutputType(PipelineOutput{})
	pulumi.RegisterOutputType(PipelineArrayOutput{})
	pulumi.RegisterOutputType(PipelineMapOutput{})
}