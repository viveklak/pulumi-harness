// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource for creating a Datadog connector.
 */
export class DatadogConnector extends pulumi.CustomResource {
    /**
     * Get an existing DatadogConnector resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatadogConnectorState, opts?: pulumi.CustomResourceOptions): DatadogConnector {
        return new DatadogConnector(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'harness:platform/datadogConnector:DatadogConnector';

    /**
     * Returns true if the given object is an instance of DatadogConnector.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DatadogConnector {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DatadogConnector.__pulumiType;
    }

    /**
     * Reference to the Harness secret containing the api key.
     */
    public readonly apiKeyRef!: pulumi.Output<string>;
    /**
     * Reference to the Harness secret containing the application key.
     */
    public readonly applicationKeyRef!: pulumi.Output<string>;
    /**
     * Connect using only the delegates which have these tags.
     */
    public readonly delegateSelectors!: pulumi.Output<string[] | undefined>;
    /**
     * Description of the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Unique identifier of the resource.
     */
    public readonly identifier!: pulumi.Output<string>;
    /**
     * Name of the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Unique identifier of the organization.
     */
    public readonly orgId!: pulumi.Output<string | undefined>;
    /**
     * Unique identifier of the project.
     */
    public readonly projectId!: pulumi.Output<string | undefined>;
    /**
     * Tags to associate with the resource. Tags should be in the form `name:value`.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;
    /**
     * Url of the Datadog server.
     */
    public readonly url!: pulumi.Output<string>;

    /**
     * Create a DatadogConnector resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatadogConnectorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatadogConnectorArgs | DatadogConnectorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatadogConnectorState | undefined;
            resourceInputs["apiKeyRef"] = state ? state.apiKeyRef : undefined;
            resourceInputs["applicationKeyRef"] = state ? state.applicationKeyRef : undefined;
            resourceInputs["delegateSelectors"] = state ? state.delegateSelectors : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["identifier"] = state ? state.identifier : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["url"] = state ? state.url : undefined;
        } else {
            const args = argsOrState as DatadogConnectorArgs | undefined;
            if ((!args || args.apiKeyRef === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiKeyRef'");
            }
            if ((!args || args.applicationKeyRef === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationKeyRef'");
            }
            if ((!args || args.identifier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identifier'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["apiKeyRef"] = args ? args.apiKeyRef : undefined;
            resourceInputs["applicationKeyRef"] = args ? args.applicationKeyRef : undefined;
            resourceInputs["delegateSelectors"] = args ? args.delegateSelectors : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["identifier"] = args ? args.identifier : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DatadogConnector.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatadogConnector resources.
 */
export interface DatadogConnectorState {
    /**
     * Reference to the Harness secret containing the api key.
     */
    apiKeyRef?: pulumi.Input<string>;
    /**
     * Reference to the Harness secret containing the application key.
     */
    applicationKeyRef?: pulumi.Input<string>;
    /**
     * Connect using only the delegates which have these tags.
     */
    delegateSelectors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Description of the resource.
     */
    description?: pulumi.Input<string>;
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
    /**
     * Tags to associate with the resource. Tags should be in the form `name:value`.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Url of the Datadog server.
     */
    url?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DatadogConnector resource.
 */
export interface DatadogConnectorArgs {
    /**
     * Reference to the Harness secret containing the api key.
     */
    apiKeyRef: pulumi.Input<string>;
    /**
     * Reference to the Harness secret containing the application key.
     */
    applicationKeyRef: pulumi.Input<string>;
    /**
     * Connect using only the delegates which have these tags.
     */
    delegateSelectors?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Unique identifier of the resource.
     */
    identifier: pulumi.Input<string>;
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
    /**
     * Tags to associate with the resource. Tags should be in the form `name:value`.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Url of the Datadog server.
     */
    url: pulumi.Input<string>;
}
