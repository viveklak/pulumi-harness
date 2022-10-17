// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Resource for creating a Harness user group
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as harness from "@pulumi/harness";
 *
 * const example = new harness.UserGroup("example", {
 *     description: "This group demonstrates account level and resource level permissions.",
 *     permissions: {
 *         accountPermissions: [
 *             "ADMINISTER_OTHER_ACCOUNT_FUNCTIONS",
 *             "MANAGE_API_KEYS",
 *         ],
 *         appPermissions: {
 *             alls: [{
 *                 actions: [
 *                     "CREATE",
 *                     "READ",
 *                     "UPDATE",
 *                     "DELETE",
 *                 ],
 *             }],
 *             deployments: [
 *                 {
 *                     actions: [
 *                         "READ",
 *                         "ROLLBACK_WORKFLOW",
 *                         "EXECUTE_PIPELINE",
 *                         "EXECUTE_WORKFLOW",
 *                     ],
 *                     filters: ["NON_PRODUCTION_ENVIRONMENTS"],
 *                 },
 *                 {
 *                     actions: ["READ"],
 *                     filters: ["PRODUCTION_ENVIRONMENTS"],
 *                 },
 *             ],
 *             environments: [
 *                 {
 *                     actions: [
 *                         "CREATE",
 *                         "READ",
 *                         "UPDATE",
 *                         "DELETE",
 *                     ],
 *                     filters: ["NON_PRODUCTION_ENVIRONMENTS"],
 *                 },
 *                 {
 *                     actions: ["READ"],
 *                     filters: ["PRODUCTION_ENVIRONMENTS"],
 *                 },
 *             ],
 *             pipelines: [
 *                 {
 *                     actions: [
 *                         "CREATE",
 *                         "READ",
 *                         "UPDATE",
 *                         "DELETE",
 *                     ],
 *                     filters: ["NON_PRODUCTION_PIPELINES"],
 *                 },
 *                 {
 *                     actions: ["READ"],
 *                     filters: ["PRODUCTION_PIPELINES"],
 *                 },
 *             ],
 *             provisioners: [
 *                 {
 *                     actions: [
 *                         "UPDATE",
 *                         "DELETE",
 *                     ],
 *                 },
 *                 {
 *                     actions: [
 *                         "CREATE",
 *                         "READ",
 *                     ],
 *                 },
 *             ],
 *             services: [
 *                 {
 *                     actions: [
 *                         "UPDATE",
 *                         "DELETE",
 *                     ],
 *                 },
 *                 {
 *                     actions: [
 *                         "UPDATE",
 *                         "DELETE",
 *                     ],
 *                 },
 *             ],
 *             templates: [{
 *                 actions: [
 *                     "CREATE",
 *                     "READ",
 *                     "UPDATE",
 *                     "DELETE",
 *                 ],
 *             }],
 *             workflows: [
 *                 {
 *                     actions: [
 *                         "UPDATE",
 *                         "DELETE",
 *                     ],
 *                     filters: ["NON_PRODUCTION_WORKFLOWS"],
 *                 },
 *                 {
 *                     actions: [
 *                         "CREATE",
 *                         "READ",
 *                     ],
 *                     filters: [
 *                         "PRODUCTION_WORKFLOWS",
 *                         "WORKFLOW_TEMPLATES",
 *                     ],
 *                 },
 *             ],
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Import using the id of the user group
 *
 * ```sh
 *  $ pulumi import harness:index/userGroup:UserGroup example <USER_GROUP_ID>
 * ```
 */
export class UserGroup extends pulumi.CustomResource {
    /**
     * Get an existing UserGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserGroupState, opts?: pulumi.CustomResourceOptions): UserGroup {
        return new UserGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'harness:index/userGroup:UserGroup';

    /**
     * Returns true if the given object is an instance of UserGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserGroup.__pulumiType;
    }

    /**
     * The description of the user group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Indicates whether the user group was imported by SCIM.
     */
    public /*out*/ readonly importedByScim!: pulumi.Output<boolean>;
    /**
     * Indicates whether the user group is linked to an SSO provider.
     */
    public /*out*/ readonly isSsoLinked!: pulumi.Output<boolean>;
    /**
     * The LDAP settings for the user group.
     */
    public readonly ldapSettings!: pulumi.Output<outputs.UserGroupLdapSettings | undefined>;
    /**
     * The name of the user group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The notification settings of the user group.
     */
    public readonly notificationSettings!: pulumi.Output<outputs.UserGroupNotificationSettings | undefined>;
    /**
     * The permissions of the user group.
     */
    public readonly permissions!: pulumi.Output<outputs.UserGroupPermissions | undefined>;
    /**
     * The SAML settings for the user group.
     */
    public readonly samlSettings!: pulumi.Output<outputs.UserGroupSamlSettings | undefined>;

    /**
     * Create a UserGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: UserGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserGroupArgs | UserGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserGroupState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["importedByScim"] = state ? state.importedByScim : undefined;
            resourceInputs["isSsoLinked"] = state ? state.isSsoLinked : undefined;
            resourceInputs["ldapSettings"] = state ? state.ldapSettings : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["notificationSettings"] = state ? state.notificationSettings : undefined;
            resourceInputs["permissions"] = state ? state.permissions : undefined;
            resourceInputs["samlSettings"] = state ? state.samlSettings : undefined;
        } else {
            const args = argsOrState as UserGroupArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["ldapSettings"] = args ? args.ldapSettings : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notificationSettings"] = args ? args.notificationSettings : undefined;
            resourceInputs["permissions"] = args ? args.permissions : undefined;
            resourceInputs["samlSettings"] = args ? args.samlSettings : undefined;
            resourceInputs["importedByScim"] = undefined /*out*/;
            resourceInputs["isSsoLinked"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserGroup resources.
 */
export interface UserGroupState {
    /**
     * The description of the user group.
     */
    description?: pulumi.Input<string>;
    /**
     * Indicates whether the user group was imported by SCIM.
     */
    importedByScim?: pulumi.Input<boolean>;
    /**
     * Indicates whether the user group is linked to an SSO provider.
     */
    isSsoLinked?: pulumi.Input<boolean>;
    /**
     * The LDAP settings for the user group.
     */
    ldapSettings?: pulumi.Input<inputs.UserGroupLdapSettings>;
    /**
     * The name of the user group.
     */
    name?: pulumi.Input<string>;
    /**
     * The notification settings of the user group.
     */
    notificationSettings?: pulumi.Input<inputs.UserGroupNotificationSettings>;
    /**
     * The permissions of the user group.
     */
    permissions?: pulumi.Input<inputs.UserGroupPermissions>;
    /**
     * The SAML settings for the user group.
     */
    samlSettings?: pulumi.Input<inputs.UserGroupSamlSettings>;
}

/**
 * The set of arguments for constructing a UserGroup resource.
 */
export interface UserGroupArgs {
    /**
     * The description of the user group.
     */
    description?: pulumi.Input<string>;
    /**
     * The LDAP settings for the user group.
     */
    ldapSettings?: pulumi.Input<inputs.UserGroupLdapSettings>;
    /**
     * The name of the user group.
     */
    name?: pulumi.Input<string>;
    /**
     * The notification settings of the user group.
     */
    notificationSettings?: pulumi.Input<inputs.UserGroupNotificationSettings>;
    /**
     * The permissions of the user group.
     */
    permissions?: pulumi.Input<inputs.UserGroupPermissions>;
    /**
     * The SAML settings for the user group.
     */
    samlSettings?: pulumi.Input<inputs.UserGroupSamlSettings>;
}
