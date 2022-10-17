// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Resource for adding permissions to an existing Harness user group.
 */
export class UserGroupPermissions extends pulumi.CustomResource {
    /**
     * Get an existing UserGroupPermissions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserGroupPermissionsState, opts?: pulumi.CustomResourceOptions): UserGroupPermissions {
        return new UserGroupPermissions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'harness:index/userGroupPermissions:UserGroupPermissions';

    /**
     * Returns true if the given object is an instance of UserGroupPermissions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserGroupPermissions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserGroupPermissions.__pulumiType;
    }

    /**
     * The account permissions of the user group. Valid options are ADMINISTER*OTHER*ACCOUNT*FUNCTIONS, CREATE*AND*DELETE*APPLICATION, CREATE*CUSTOM*DASHBOARDS, MANAGE*ALERT*NOTIFICATION*RULES, MANAGE*API*KEYS, MANAGE*APPLICATION*STACKS, MANAGE*AUTHENTICATION*SETTINGS, MANAGE*CLOUD*PROVIDERS, MANAGE*CONFIG*AS*CODE, MANAGE*CONNECTORS, MANAGE*CUSTOM*DASHBOARDS, MANAGE*DELEGATE*PROFILES, MANAGE*DELEGATES, MANAGE*DEPLOYMENT*FREEZES, MANAGE*IP*WHITELIST, MANAGE*PIPELINE*GOVERNANCE*STANDARDS, MANAGE*RESTRICTED*ACCESS, MANAGE*SECRET*MANAGERS, MANAGE*SECRETS, MANAGE*SSH*AND*WINRM, MANAGE*TAGS, MANAGE*TEMPLATE*LIBRARY, MANAGE*USER*AND*USER*GROUPS*AND*API*KEYS, MANAGE*USERS*AND*GROUPS, READ*USERS*AND*GROUPS, VIEW*AUDITS, VIEW*USER*AND*USER*GROUPS*AND*API_KEYS
     */
    public readonly accountPermissions!: pulumi.Output<string[] | undefined>;
    /**
     * Application specific permissions
     */
    public readonly appPermissions!: pulumi.Output<outputs.UserGroupPermissionsAppPermissions | undefined>;
    /**
     * Unique identifier of the user group.
     */
    public readonly userGroupId!: pulumi.Output<string>;

    /**
     * Create a UserGroupPermissions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserGroupPermissionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserGroupPermissionsArgs | UserGroupPermissionsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserGroupPermissionsState | undefined;
            resourceInputs["accountPermissions"] = state ? state.accountPermissions : undefined;
            resourceInputs["appPermissions"] = state ? state.appPermissions : undefined;
            resourceInputs["userGroupId"] = state ? state.userGroupId : undefined;
        } else {
            const args = argsOrState as UserGroupPermissionsArgs | undefined;
            if ((!args || args.userGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userGroupId'");
            }
            resourceInputs["accountPermissions"] = args ? args.accountPermissions : undefined;
            resourceInputs["appPermissions"] = args ? args.appPermissions : undefined;
            resourceInputs["userGroupId"] = args ? args.userGroupId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserGroupPermissions.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserGroupPermissions resources.
 */
export interface UserGroupPermissionsState {
    /**
     * The account permissions of the user group. Valid options are ADMINISTER*OTHER*ACCOUNT*FUNCTIONS, CREATE*AND*DELETE*APPLICATION, CREATE*CUSTOM*DASHBOARDS, MANAGE*ALERT*NOTIFICATION*RULES, MANAGE*API*KEYS, MANAGE*APPLICATION*STACKS, MANAGE*AUTHENTICATION*SETTINGS, MANAGE*CLOUD*PROVIDERS, MANAGE*CONFIG*AS*CODE, MANAGE*CONNECTORS, MANAGE*CUSTOM*DASHBOARDS, MANAGE*DELEGATE*PROFILES, MANAGE*DELEGATES, MANAGE*DEPLOYMENT*FREEZES, MANAGE*IP*WHITELIST, MANAGE*PIPELINE*GOVERNANCE*STANDARDS, MANAGE*RESTRICTED*ACCESS, MANAGE*SECRET*MANAGERS, MANAGE*SECRETS, MANAGE*SSH*AND*WINRM, MANAGE*TAGS, MANAGE*TEMPLATE*LIBRARY, MANAGE*USER*AND*USER*GROUPS*AND*API*KEYS, MANAGE*USERS*AND*GROUPS, READ*USERS*AND*GROUPS, VIEW*AUDITS, VIEW*USER*AND*USER*GROUPS*AND*API_KEYS
     */
    accountPermissions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Application specific permissions
     */
    appPermissions?: pulumi.Input<inputs.UserGroupPermissionsAppPermissions>;
    /**
     * Unique identifier of the user group.
     */
    userGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserGroupPermissions resource.
 */
export interface UserGroupPermissionsArgs {
    /**
     * The account permissions of the user group. Valid options are ADMINISTER*OTHER*ACCOUNT*FUNCTIONS, CREATE*AND*DELETE*APPLICATION, CREATE*CUSTOM*DASHBOARDS, MANAGE*ALERT*NOTIFICATION*RULES, MANAGE*API*KEYS, MANAGE*APPLICATION*STACKS, MANAGE*AUTHENTICATION*SETTINGS, MANAGE*CLOUD*PROVIDERS, MANAGE*CONFIG*AS*CODE, MANAGE*CONNECTORS, MANAGE*CUSTOM*DASHBOARDS, MANAGE*DELEGATE*PROFILES, MANAGE*DELEGATES, MANAGE*DEPLOYMENT*FREEZES, MANAGE*IP*WHITELIST, MANAGE*PIPELINE*GOVERNANCE*STANDARDS, MANAGE*RESTRICTED*ACCESS, MANAGE*SECRET*MANAGERS, MANAGE*SECRETS, MANAGE*SSH*AND*WINRM, MANAGE*TAGS, MANAGE*TEMPLATE*LIBRARY, MANAGE*USER*AND*USER*GROUPS*AND*API*KEYS, MANAGE*USERS*AND*GROUPS, READ*USERS*AND*GROUPS, VIEW*AUDITS, VIEW*USER*AND*USER*GROUPS*AND*API_KEYS
     */
    accountPermissions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Application specific permissions
     */
    appPermissions?: pulumi.Input<inputs.UserGroupPermissionsAppPermissions>;
    /**
     * Unique identifier of the user group.
     */
    userGroupId: pulumi.Input<string>;
}
