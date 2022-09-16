// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Datasource for looking up a Kubernetes connector.
 */
export function getKubernetesConnector(args?: GetKubernetesConnectorArgs, opts?: pulumi.InvokeOptions): Promise<GetKubernetesConnectorResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("harness:platform/getKubernetesConnector:getKubernetesConnector", {
        "identifier": args.identifier,
        "name": args.name,
        "orgId": args.orgId,
        "projectId": args.projectId,
    }, opts);
}

/**
 * A collection of arguments for invoking getKubernetesConnector.
 */
export interface GetKubernetesConnectorArgs {
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
 * A collection of values returned by getKubernetesConnector.
 */
export interface GetKubernetesConnectorResult {
    /**
     * Client key and certificate config for the connector.
     */
    readonly clientKeyCerts: outputs.platform.GetKubernetesConnectorClientKeyCert[];
    /**
     * Selectors to use for the delegate.
     */
    readonly delegateSelectors: string[];
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
     * Credentials are inherited from the delegate.
     */
    readonly inheritFromDelegates: outputs.platform.GetKubernetesConnectorInheritFromDelegate[];
    /**
     * Name of the resource.
     */
    readonly name?: string;
    /**
     * OpenID configuration for the connector.
     */
    readonly openidConnects: outputs.platform.GetKubernetesConnectorOpenidConnect[];
    /**
     * Unique identifier of the organization.
     */
    readonly orgId?: string;
    /**
     * Unique identifier of the project.
     */
    readonly projectId?: string;
    /**
     * Service account for the connector.
     */
    readonly serviceAccounts: outputs.platform.GetKubernetesConnectorServiceAccount[];
    /**
     * Tags to associate with the resource. Tags should be in the form `name:value`.
     */
    readonly tags: string[];
    /**
     * Username and password for the connector.
     */
    readonly usernamePasswords: outputs.platform.GetKubernetesConnectorUsernamePassword[];
}

export function getKubernetesConnectorOutput(args?: GetKubernetesConnectorOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetKubernetesConnectorResult> {
    return pulumi.output(args).apply(a => getKubernetesConnector(a, opts))
}

/**
 * A collection of arguments for invoking getKubernetesConnector.
 */
export interface GetKubernetesConnectorOutputArgs {
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