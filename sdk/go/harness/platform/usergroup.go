// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package platform

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for creating a Harness User Group.
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
//			_, err := platform.NewUsergroup(ctx, "example", &platform.UsergroupArgs{
//				ExternallyManaged:    pulumi.Bool(false),
//				Identifier:           pulumi.String("identifier"),
//				LinkedSsoDisplayName: pulumi.String("linked_sso_display_name"),
//				LinkedSsoId:          pulumi.String("linked_sso_id"),
//				LinkedSsoType:        pulumi.String("SAML"),
//				NotificationConfigs: platform.UsergroupNotificationConfigArray{
//					&platform.UsergroupNotificationConfigArgs{
//						SlackWebhookUrl: pulumi.String("https://google.com"),
//						Type:            pulumi.String("SLACK"),
//					},
//					&platform.UsergroupNotificationConfigArgs{
//						GroupEmail: pulumi.String("email@email.com"),
//						Type:       pulumi.String("EMAIL"),
//					},
//					&platform.UsergroupNotificationConfigArgs{
//						MicrosoftTeamsWebhookUrl: pulumi.String("https://google.com"),
//						Type:                     pulumi.String("MSTEAMS"),
//					},
//					&platform.UsergroupNotificationConfigArgs{
//						PagerDutyKey: pulumi.String("pagerDutyKey"),
//						Type:         pulumi.String("PAGERDUTY"),
//					},
//				},
//				OrgId:        pulumi.String("org_id"),
//				ProjectId:    pulumi.String("project_id"),
//				SsoGroupId:   pulumi.String("sso_group_id"),
//				SsoGroupName: pulumi.String("sso_group_name"),
//				SsoLinked:    pulumi.Bool(true),
//				Users: pulumi.StringArray{
//					pulumi.String("user_id"),
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
// # Import using user group id
//
// ```sh
//
//	$ pulumi import harness:platform/usergroup:Usergroup example <usergroup_id>
//
// ```
type Usergroup struct {
	pulumi.CustomResourceState

	// Description of the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether the user group is externally managed.
	ExternallyManaged pulumi.BoolPtrOutput `pulumi:"externallyManaged"`
	// Unique identifier of the resource.
	Identifier pulumi.StringOutput `pulumi:"identifier"`
	// Name of the linked SSO.
	LinkedSsoDisplayName pulumi.StringPtrOutput `pulumi:"linkedSsoDisplayName"`
	// The SSO account ID that the user group is linked to.
	LinkedSsoId pulumi.StringPtrOutput `pulumi:"linkedSsoId"`
	// Type of linked SSO
	LinkedSsoType pulumi.StringPtrOutput `pulumi:"linkedSsoType"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of notification settings.
	NotificationConfigs UsergroupNotificationConfigArrayOutput `pulumi:"notificationConfigs"`
	// Unique identifier of the organization.
	OrgId pulumi.StringPtrOutput `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrOutput `pulumi:"projectId"`
	// Identifier of the userGroup in SSO.
	SsoGroupId pulumi.StringPtrOutput `pulumi:"ssoGroupId"`
	// Name of the SSO userGroup.
	SsoGroupName pulumi.StringPtrOutput `pulumi:"ssoGroupName"`
	// Whether sso is linked or not
	SsoLinked pulumi.BoolOutput `pulumi:"ssoLinked"`
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// List of users in the UserGroup.
	Users pulumi.StringArrayOutput `pulumi:"users"`
}

// NewUsergroup registers a new resource with the given unique name, arguments, and options.
func NewUsergroup(ctx *pulumi.Context,
	name string, args *UsergroupArgs, opts ...pulumi.ResourceOption) (*Usergroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Identifier == nil {
		return nil, errors.New("invalid value for required argument 'Identifier'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Usergroup
	err := ctx.RegisterResource("harness:platform/usergroup:Usergroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUsergroup gets an existing Usergroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUsergroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UsergroupState, opts ...pulumi.ResourceOption) (*Usergroup, error) {
	var resource Usergroup
	err := ctx.ReadResource("harness:platform/usergroup:Usergroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Usergroup resources.
type usergroupState struct {
	// Description of the resource.
	Description *string `pulumi:"description"`
	// Whether the user group is externally managed.
	ExternallyManaged *bool `pulumi:"externallyManaged"`
	// Unique identifier of the resource.
	Identifier *string `pulumi:"identifier"`
	// Name of the linked SSO.
	LinkedSsoDisplayName *string `pulumi:"linkedSsoDisplayName"`
	// The SSO account ID that the user group is linked to.
	LinkedSsoId *string `pulumi:"linkedSsoId"`
	// Type of linked SSO
	LinkedSsoType *string `pulumi:"linkedSsoType"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// List of notification settings.
	NotificationConfigs []UsergroupNotificationConfig `pulumi:"notificationConfigs"`
	// Unique identifier of the organization.
	OrgId *string `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId *string `pulumi:"projectId"`
	// Identifier of the userGroup in SSO.
	SsoGroupId *string `pulumi:"ssoGroupId"`
	// Name of the SSO userGroup.
	SsoGroupName *string `pulumi:"ssoGroupName"`
	// Whether sso is linked or not
	SsoLinked *bool `pulumi:"ssoLinked"`
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags []string `pulumi:"tags"`
	// List of users in the UserGroup.
	Users []string `pulumi:"users"`
}

type UsergroupState struct {
	// Description of the resource.
	Description pulumi.StringPtrInput
	// Whether the user group is externally managed.
	ExternallyManaged pulumi.BoolPtrInput
	// Unique identifier of the resource.
	Identifier pulumi.StringPtrInput
	// Name of the linked SSO.
	LinkedSsoDisplayName pulumi.StringPtrInput
	// The SSO account ID that the user group is linked to.
	LinkedSsoId pulumi.StringPtrInput
	// Type of linked SSO
	LinkedSsoType pulumi.StringPtrInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// List of notification settings.
	NotificationConfigs UsergroupNotificationConfigArrayInput
	// Unique identifier of the organization.
	OrgId pulumi.StringPtrInput
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrInput
	// Identifier of the userGroup in SSO.
	SsoGroupId pulumi.StringPtrInput
	// Name of the SSO userGroup.
	SsoGroupName pulumi.StringPtrInput
	// Whether sso is linked or not
	SsoLinked pulumi.BoolPtrInput
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags pulumi.StringArrayInput
	// List of users in the UserGroup.
	Users pulumi.StringArrayInput
}

func (UsergroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*usergroupState)(nil)).Elem()
}

type usergroupArgs struct {
	// Description of the resource.
	Description *string `pulumi:"description"`
	// Whether the user group is externally managed.
	ExternallyManaged *bool `pulumi:"externallyManaged"`
	// Unique identifier of the resource.
	Identifier string `pulumi:"identifier"`
	// Name of the linked SSO.
	LinkedSsoDisplayName *string `pulumi:"linkedSsoDisplayName"`
	// The SSO account ID that the user group is linked to.
	LinkedSsoId *string `pulumi:"linkedSsoId"`
	// Type of linked SSO
	LinkedSsoType *string `pulumi:"linkedSsoType"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// List of notification settings.
	NotificationConfigs []UsergroupNotificationConfig `pulumi:"notificationConfigs"`
	// Unique identifier of the organization.
	OrgId *string `pulumi:"orgId"`
	// Unique identifier of the project.
	ProjectId *string `pulumi:"projectId"`
	// Identifier of the userGroup in SSO.
	SsoGroupId *string `pulumi:"ssoGroupId"`
	// Name of the SSO userGroup.
	SsoGroupName *string `pulumi:"ssoGroupName"`
	// Whether sso is linked or not
	SsoLinked *bool `pulumi:"ssoLinked"`
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags []string `pulumi:"tags"`
	// List of users in the UserGroup.
	Users []string `pulumi:"users"`
}

// The set of arguments for constructing a Usergroup resource.
type UsergroupArgs struct {
	// Description of the resource.
	Description pulumi.StringPtrInput
	// Whether the user group is externally managed.
	ExternallyManaged pulumi.BoolPtrInput
	// Unique identifier of the resource.
	Identifier pulumi.StringInput
	// Name of the linked SSO.
	LinkedSsoDisplayName pulumi.StringPtrInput
	// The SSO account ID that the user group is linked to.
	LinkedSsoId pulumi.StringPtrInput
	// Type of linked SSO
	LinkedSsoType pulumi.StringPtrInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// List of notification settings.
	NotificationConfigs UsergroupNotificationConfigArrayInput
	// Unique identifier of the organization.
	OrgId pulumi.StringPtrInput
	// Unique identifier of the project.
	ProjectId pulumi.StringPtrInput
	// Identifier of the userGroup in SSO.
	SsoGroupId pulumi.StringPtrInput
	// Name of the SSO userGroup.
	SsoGroupName pulumi.StringPtrInput
	// Whether sso is linked or not
	SsoLinked pulumi.BoolPtrInput
	// Tags to associate with the resource. Tags should be in the form `name:value`.
	Tags pulumi.StringArrayInput
	// List of users in the UserGroup.
	Users pulumi.StringArrayInput
}

func (UsergroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*usergroupArgs)(nil)).Elem()
}

type UsergroupInput interface {
	pulumi.Input

	ToUsergroupOutput() UsergroupOutput
	ToUsergroupOutputWithContext(ctx context.Context) UsergroupOutput
}

func (*Usergroup) ElementType() reflect.Type {
	return reflect.TypeOf((**Usergroup)(nil)).Elem()
}

func (i *Usergroup) ToUsergroupOutput() UsergroupOutput {
	return i.ToUsergroupOutputWithContext(context.Background())
}

func (i *Usergroup) ToUsergroupOutputWithContext(ctx context.Context) UsergroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsergroupOutput)
}

// UsergroupArrayInput is an input type that accepts UsergroupArray and UsergroupArrayOutput values.
// You can construct a concrete instance of `UsergroupArrayInput` via:
//
//	UsergroupArray{ UsergroupArgs{...} }
type UsergroupArrayInput interface {
	pulumi.Input

	ToUsergroupArrayOutput() UsergroupArrayOutput
	ToUsergroupArrayOutputWithContext(context.Context) UsergroupArrayOutput
}

type UsergroupArray []UsergroupInput

func (UsergroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Usergroup)(nil)).Elem()
}

func (i UsergroupArray) ToUsergroupArrayOutput() UsergroupArrayOutput {
	return i.ToUsergroupArrayOutputWithContext(context.Background())
}

func (i UsergroupArray) ToUsergroupArrayOutputWithContext(ctx context.Context) UsergroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsergroupArrayOutput)
}

// UsergroupMapInput is an input type that accepts UsergroupMap and UsergroupMapOutput values.
// You can construct a concrete instance of `UsergroupMapInput` via:
//
//	UsergroupMap{ "key": UsergroupArgs{...} }
type UsergroupMapInput interface {
	pulumi.Input

	ToUsergroupMapOutput() UsergroupMapOutput
	ToUsergroupMapOutputWithContext(context.Context) UsergroupMapOutput
}

type UsergroupMap map[string]UsergroupInput

func (UsergroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Usergroup)(nil)).Elem()
}

func (i UsergroupMap) ToUsergroupMapOutput() UsergroupMapOutput {
	return i.ToUsergroupMapOutputWithContext(context.Background())
}

func (i UsergroupMap) ToUsergroupMapOutputWithContext(ctx context.Context) UsergroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsergroupMapOutput)
}

type UsergroupOutput struct{ *pulumi.OutputState }

func (UsergroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Usergroup)(nil)).Elem()
}

func (o UsergroupOutput) ToUsergroupOutput() UsergroupOutput {
	return o
}

func (o UsergroupOutput) ToUsergroupOutputWithContext(ctx context.Context) UsergroupOutput {
	return o
}

// Description of the resource.
func (o UsergroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Usergroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether the user group is externally managed.
func (o UsergroupOutput) ExternallyManaged() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Usergroup) pulumi.BoolPtrOutput { return v.ExternallyManaged }).(pulumi.BoolPtrOutput)
}

// Unique identifier of the resource.
func (o UsergroupOutput) Identifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Usergroup) pulumi.StringOutput { return v.Identifier }).(pulumi.StringOutput)
}

// Name of the linked SSO.
func (o UsergroupOutput) LinkedSsoDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Usergroup) pulumi.StringPtrOutput { return v.LinkedSsoDisplayName }).(pulumi.StringPtrOutput)
}

// The SSO account ID that the user group is linked to.
func (o UsergroupOutput) LinkedSsoId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Usergroup) pulumi.StringPtrOutput { return v.LinkedSsoId }).(pulumi.StringPtrOutput)
}

// Type of linked SSO
func (o UsergroupOutput) LinkedSsoType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Usergroup) pulumi.StringPtrOutput { return v.LinkedSsoType }).(pulumi.StringPtrOutput)
}

// Name of the resource.
func (o UsergroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Usergroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// List of notification settings.
func (o UsergroupOutput) NotificationConfigs() UsergroupNotificationConfigArrayOutput {
	return o.ApplyT(func(v *Usergroup) UsergroupNotificationConfigArrayOutput { return v.NotificationConfigs }).(UsergroupNotificationConfigArrayOutput)
}

// Unique identifier of the organization.
func (o UsergroupOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Usergroup) pulumi.StringPtrOutput { return v.OrgId }).(pulumi.StringPtrOutput)
}

// Unique identifier of the project.
func (o UsergroupOutput) ProjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Usergroup) pulumi.StringPtrOutput { return v.ProjectId }).(pulumi.StringPtrOutput)
}

// Identifier of the userGroup in SSO.
func (o UsergroupOutput) SsoGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Usergroup) pulumi.StringPtrOutput { return v.SsoGroupId }).(pulumi.StringPtrOutput)
}

// Name of the SSO userGroup.
func (o UsergroupOutput) SsoGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Usergroup) pulumi.StringPtrOutput { return v.SsoGroupName }).(pulumi.StringPtrOutput)
}

// Whether sso is linked or not
func (o UsergroupOutput) SsoLinked() pulumi.BoolOutput {
	return o.ApplyT(func(v *Usergroup) pulumi.BoolOutput { return v.SsoLinked }).(pulumi.BoolOutput)
}

// Tags to associate with the resource. Tags should be in the form `name:value`.
func (o UsergroupOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Usergroup) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// List of users in the UserGroup.
func (o UsergroupOutput) Users() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Usergroup) pulumi.StringArrayOutput { return v.Users }).(pulumi.StringArrayOutput)
}

type UsergroupArrayOutput struct{ *pulumi.OutputState }

func (UsergroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Usergroup)(nil)).Elem()
}

func (o UsergroupArrayOutput) ToUsergroupArrayOutput() UsergroupArrayOutput {
	return o
}

func (o UsergroupArrayOutput) ToUsergroupArrayOutputWithContext(ctx context.Context) UsergroupArrayOutput {
	return o
}

func (o UsergroupArrayOutput) Index(i pulumi.IntInput) UsergroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Usergroup {
		return vs[0].([]*Usergroup)[vs[1].(int)]
	}).(UsergroupOutput)
}

type UsergroupMapOutput struct{ *pulumi.OutputState }

func (UsergroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Usergroup)(nil)).Elem()
}

func (o UsergroupMapOutput) ToUsergroupMapOutput() UsergroupMapOutput {
	return o
}

func (o UsergroupMapOutput) ToUsergroupMapOutputWithContext(ctx context.Context) UsergroupMapOutput {
	return o
}

func (o UsergroupMapOutput) MapIndex(k pulumi.StringInput) UsergroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Usergroup {
		return vs[0].(map[string]*Usergroup)[vs[1].(string)]
	}).(UsergroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UsergroupInput)(nil)).Elem(), &Usergroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*UsergroupArrayInput)(nil)).Elem(), UsergroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UsergroupMapInput)(nil)).Elem(), UsergroupMap{})
	pulumi.RegisterOutputType(UsergroupOutput{})
	pulumi.RegisterOutputType(UsergroupArrayOutput{})
	pulumi.RegisterOutputType(UsergroupMapOutput{})
}