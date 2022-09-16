// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Resource for looking up an SSH Key type secret.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as harness from "@pulumi/harness";
 *
 * const example = pulumi.output(harness.platform.getSecretSshkey({
 *     identifier: "identifier",
 * }));
 * ```
 */
export function getSecretSshkey(args?: GetSecretSshkeyArgs, opts?: pulumi.InvokeOptions): Promise<GetSecretSshkeyResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("harness:platform/getSecretSshkey:getSecretSshkey", {
        "identifier": args.identifier,
        "name": args.name,
        "orgId": args.orgId,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecretSshkey.
 */
export interface GetSecretSshkeyArgs {
    /**
     * Unique identifier of the resource.
     */
    identifier?: string;
    /**
     * Name of the resource.
     */
    name?: string;
    /**
     * Unique identifier of the organization.
     */
    orgId?: string;
    /**
     * Unique identifier of the project.
     */
    projectId?: string;
}

/**
 * A collection of values returned by getSecretSshkey.
 */
export interface GetSecretSshkeyResult {
    /**
     * Description of the resource.
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Unique identifier of the resource.
     */
    readonly identifier?: string;
    /**
     * Kerberos authentication scheme
     */
    readonly kerberos: outputs.platform.GetSecretSshkeyKerbero[];
    /**
     * Name of the resource.
     */
    readonly name?: string;
    /**
     * Unique identifier of the organization.
     */
    readonly orgId?: string;
    /**
     * SSH port
     */
    readonly port: number;
    /**
     * Unique identifier of the project.
     */
    readonly projectId?: string;
    /**
     * Kerberos authentication scheme
     */
    readonly sshes: outputs.platform.GetSecretSshkeySsh[];
    /**
     * Tags to associate with the resource. Tags should be in the form `name:value`.
     */
    readonly tags: string[];
}

export function getSecretSshkeyOutput(args?: GetSecretSshkeyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSecretSshkeyResult> {
    return pulumi.output(args).apply(a => getSecretSshkey(a, opts))
}

/**
 * A collection of arguments for invoking getSecretSshkey.
 */
export interface GetSecretSshkeyOutputArgs {
    /**
     * Unique identifier of the resource.
     */
    identifier?: pulumi.Input<string>;
    /**
     * Name of the resource.
     */
    name?: pulumi.Input<string>;
    /**
     * Unique identifier of the organization.
     */
    orgId?: pulumi.Input<string>;
    /**
     * Unique identifier of the project.
     */
    projectId?: pulumi.Input<string>;
}