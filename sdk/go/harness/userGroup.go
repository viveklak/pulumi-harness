// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package harness

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resource for creating a Harness user group
//
// ## Import
//
// # Import using the id of the user group
//
// ```sh
//
//	$ pulumi import harness:index/userGroup:UserGroup example <USER_GROUP_ID>
//
// ```
type UserGroup struct {
	pulumi.CustomResourceState

	// The description of the user group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Indicates whether the user group was imported by SCIM.
	ImportedByScim pulumi.BoolOutput `pulumi:"importedByScim"`
	// Indicates whether the user group is linked to an SSO provider.
	IsSsoLinked pulumi.BoolOutput `pulumi:"isSsoLinked"`
	// The LDAP settings for the user group.
	LdapSettings UserGroupLdapSettingsPtrOutput `pulumi:"ldapSettings"`
	// The name of the user group.
	Name pulumi.StringOutput `pulumi:"name"`
	// The notification settings of the user group.
	NotificationSettings UserGroupNotificationSettingsPtrOutput `pulumi:"notificationSettings"`
	// The permissions of the user group.
	Permissions UserGroupPermissionsPtrOutput `pulumi:"permissions"`
	// The SAML settings for the user group.
	SamlSettings UserGroupSamlSettingsPtrOutput `pulumi:"samlSettings"`
}

// NewUserGroup registers a new resource with the given unique name, arguments, and options.
func NewUserGroup(ctx *pulumi.Context,
	name string, args *UserGroupArgs, opts ...pulumi.ResourceOption) (*UserGroup, error) {
	if args == nil {
		args = &UserGroupArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource UserGroup
	err := ctx.RegisterResource("harness:index/userGroup:UserGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserGroup gets an existing UserGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserGroupState, opts ...pulumi.ResourceOption) (*UserGroup, error) {
	var resource UserGroup
	err := ctx.ReadResource("harness:index/userGroup:UserGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserGroup resources.
type userGroupState struct {
	// The description of the user group.
	Description *string `pulumi:"description"`
	// Indicates whether the user group was imported by SCIM.
	ImportedByScim *bool `pulumi:"importedByScim"`
	// Indicates whether the user group is linked to an SSO provider.
	IsSsoLinked *bool `pulumi:"isSsoLinked"`
	// The LDAP settings for the user group.
	LdapSettings *UserGroupLdapSettings `pulumi:"ldapSettings"`
	// The name of the user group.
	Name *string `pulumi:"name"`
	// The notification settings of the user group.
	NotificationSettings *UserGroupNotificationSettings `pulumi:"notificationSettings"`
	// The permissions of the user group.
	Permissions *UserGroupPermissions `pulumi:"permissions"`
	// The SAML settings for the user group.
	SamlSettings *UserGroupSamlSettings `pulumi:"samlSettings"`
}

type UserGroupState struct {
	// The description of the user group.
	Description pulumi.StringPtrInput
	// Indicates whether the user group was imported by SCIM.
	ImportedByScim pulumi.BoolPtrInput
	// Indicates whether the user group is linked to an SSO provider.
	IsSsoLinked pulumi.BoolPtrInput
	// The LDAP settings for the user group.
	LdapSettings UserGroupLdapSettingsPtrInput
	// The name of the user group.
	Name pulumi.StringPtrInput
	// The notification settings of the user group.
	NotificationSettings UserGroupNotificationSettingsPtrInput
	// The permissions of the user group.
	Permissions UserGroupPermissionsPtrInput
	// The SAML settings for the user group.
	SamlSettings UserGroupSamlSettingsPtrInput
}

func (UserGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*userGroupState)(nil)).Elem()
}

type userGroupArgs struct {
	// The description of the user group.
	Description *string `pulumi:"description"`
	// The LDAP settings for the user group.
	LdapSettings *UserGroupLdapSettings `pulumi:"ldapSettings"`
	// The name of the user group.
	Name *string `pulumi:"name"`
	// The notification settings of the user group.
	NotificationSettings *UserGroupNotificationSettings `pulumi:"notificationSettings"`
	// The permissions of the user group.
	Permissions *UserGroupPermissions `pulumi:"permissions"`
	// The SAML settings for the user group.
	SamlSettings *UserGroupSamlSettings `pulumi:"samlSettings"`
}

// The set of arguments for constructing a UserGroup resource.
type UserGroupArgs struct {
	// The description of the user group.
	Description pulumi.StringPtrInput
	// The LDAP settings for the user group.
	LdapSettings UserGroupLdapSettingsPtrInput
	// The name of the user group.
	Name pulumi.StringPtrInput
	// The notification settings of the user group.
	NotificationSettings UserGroupNotificationSettingsPtrInput
	// The permissions of the user group.
	Permissions UserGroupPermissionsPtrInput
	// The SAML settings for the user group.
	SamlSettings UserGroupSamlSettingsPtrInput
}

func (UserGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userGroupArgs)(nil)).Elem()
}

type UserGroupInput interface {
	pulumi.Input

	ToUserGroupOutput() UserGroupOutput
	ToUserGroupOutputWithContext(ctx context.Context) UserGroupOutput
}

func (*UserGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**UserGroup)(nil)).Elem()
}

func (i *UserGroup) ToUserGroupOutput() UserGroupOutput {
	return i.ToUserGroupOutputWithContext(context.Background())
}

func (i *UserGroup) ToUserGroupOutputWithContext(ctx context.Context) UserGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGroupOutput)
}

// UserGroupArrayInput is an input type that accepts UserGroupArray and UserGroupArrayOutput values.
// You can construct a concrete instance of `UserGroupArrayInput` via:
//
//	UserGroupArray{ UserGroupArgs{...} }
type UserGroupArrayInput interface {
	pulumi.Input

	ToUserGroupArrayOutput() UserGroupArrayOutput
	ToUserGroupArrayOutputWithContext(context.Context) UserGroupArrayOutput
}

type UserGroupArray []UserGroupInput

func (UserGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserGroup)(nil)).Elem()
}

func (i UserGroupArray) ToUserGroupArrayOutput() UserGroupArrayOutput {
	return i.ToUserGroupArrayOutputWithContext(context.Background())
}

func (i UserGroupArray) ToUserGroupArrayOutputWithContext(ctx context.Context) UserGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGroupArrayOutput)
}

// UserGroupMapInput is an input type that accepts UserGroupMap and UserGroupMapOutput values.
// You can construct a concrete instance of `UserGroupMapInput` via:
//
//	UserGroupMap{ "key": UserGroupArgs{...} }
type UserGroupMapInput interface {
	pulumi.Input

	ToUserGroupMapOutput() UserGroupMapOutput
	ToUserGroupMapOutputWithContext(context.Context) UserGroupMapOutput
}

type UserGroupMap map[string]UserGroupInput

func (UserGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserGroup)(nil)).Elem()
}

func (i UserGroupMap) ToUserGroupMapOutput() UserGroupMapOutput {
	return i.ToUserGroupMapOutputWithContext(context.Background())
}

func (i UserGroupMap) ToUserGroupMapOutputWithContext(ctx context.Context) UserGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserGroupMapOutput)
}

type UserGroupOutput struct{ *pulumi.OutputState }

func (UserGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserGroup)(nil)).Elem()
}

func (o UserGroupOutput) ToUserGroupOutput() UserGroupOutput {
	return o
}

func (o UserGroupOutput) ToUserGroupOutputWithContext(ctx context.Context) UserGroupOutput {
	return o
}

// The description of the user group.
func (o UserGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Indicates whether the user group was imported by SCIM.
func (o UserGroupOutput) ImportedByScim() pulumi.BoolOutput {
	return o.ApplyT(func(v *UserGroup) pulumi.BoolOutput { return v.ImportedByScim }).(pulumi.BoolOutput)
}

// Indicates whether the user group is linked to an SSO provider.
func (o UserGroupOutput) IsSsoLinked() pulumi.BoolOutput {
	return o.ApplyT(func(v *UserGroup) pulumi.BoolOutput { return v.IsSsoLinked }).(pulumi.BoolOutput)
}

// The LDAP settings for the user group.
func (o UserGroupOutput) LdapSettings() UserGroupLdapSettingsPtrOutput {
	return o.ApplyT(func(v *UserGroup) UserGroupLdapSettingsPtrOutput { return v.LdapSettings }).(UserGroupLdapSettingsPtrOutput)
}

// The name of the user group.
func (o UserGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UserGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The notification settings of the user group.
func (o UserGroupOutput) NotificationSettings() UserGroupNotificationSettingsPtrOutput {
	return o.ApplyT(func(v *UserGroup) UserGroupNotificationSettingsPtrOutput { return v.NotificationSettings }).(UserGroupNotificationSettingsPtrOutput)
}

// The permissions of the user group.
func (o UserGroupOutput) Permissions() UserGroupPermissionsPtrOutput {
	return o.ApplyT(func(v *UserGroup) UserGroupPermissionsPtrOutput { return v.Permissions }).(UserGroupPermissionsPtrOutput)
}

// The SAML settings for the user group.
func (o UserGroupOutput) SamlSettings() UserGroupSamlSettingsPtrOutput {
	return o.ApplyT(func(v *UserGroup) UserGroupSamlSettingsPtrOutput { return v.SamlSettings }).(UserGroupSamlSettingsPtrOutput)
}

type UserGroupArrayOutput struct{ *pulumi.OutputState }

func (UserGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserGroup)(nil)).Elem()
}

func (o UserGroupArrayOutput) ToUserGroupArrayOutput() UserGroupArrayOutput {
	return o
}

func (o UserGroupArrayOutput) ToUserGroupArrayOutputWithContext(ctx context.Context) UserGroupArrayOutput {
	return o
}

func (o UserGroupArrayOutput) Index(i pulumi.IntInput) UserGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserGroup {
		return vs[0].([]*UserGroup)[vs[1].(int)]
	}).(UserGroupOutput)
}

type UserGroupMapOutput struct{ *pulumi.OutputState }

func (UserGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserGroup)(nil)).Elem()
}

func (o UserGroupMapOutput) ToUserGroupMapOutput() UserGroupMapOutput {
	return o
}

func (o UserGroupMapOutput) ToUserGroupMapOutputWithContext(ctx context.Context) UserGroupMapOutput {
	return o
}

func (o UserGroupMapOutput) MapIndex(k pulumi.StringInput) UserGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserGroup {
		return vs[0].(map[string]*UserGroup)[vs[1].(string)]
	}).(UserGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserGroupInput)(nil)).Elem(), &UserGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserGroupArrayInput)(nil)).Elem(), UserGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserGroupMapInput)(nil)).Elem(), UserGroupMap{})
	pulumi.RegisterOutputType(UserGroupOutput{})
	pulumi.RegisterOutputType(UserGroupArrayOutput{})
	pulumi.RegisterOutputType(UserGroupMapOutput{})
}
