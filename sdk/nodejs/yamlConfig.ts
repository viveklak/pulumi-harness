// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Resource for creating a raw YAML configuration in Harness. Note: This works for all objects EXCEPT application objects. This resource uses the config-as-code API's. When updating the `name` or `path` of this resource you should typically also set the `createBeforeDestroy = true` lifecycle setting.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as harness from "@pulumi/harness";
 *
 * const test = new harness.YamlConfig("test", {
 *     content: `harnessApiVersion: '1.0'
 * type: KUBERNETES_CLUSTER
 * delegateSelectors:
 * - k8s
 * skipValidation: true
 * useKubernetesDelegate: true
 * `,
 *     path: "Setup/Cloud Providers/Kubernetes.yaml",
 * });
 * ```
 *
 * ## Import
 *
 * # Importing a global config only using the yaml path
 *
 * ```sh
 *  $ pulumi import harness:index/yamlConfig:YamlConfig k8s_cloudprovider "Setup/Cloud Providers/kubernetes.yaml"
 * ```
 *
 * # Importing a service which requires both the application id and the yaml path.
 *
 * ```sh
 *  $ pulumi import harness:index/yamlConfig:YamlConfig k8s_cloudprovider "Setup/Applications/MyApp/Services/MyService/Index.yaml:<APPLICATION_ID>"
 * ```
 */
export class YamlConfig extends pulumi.CustomResource {
    /**
     * Get an existing YamlConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: YamlConfigState, opts?: pulumi.CustomResourceOptions): YamlConfig {
        return new YamlConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'harness:index/yamlConfig:YamlConfig';

    /**
     * Returns true if the given object is an instance of YamlConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is YamlConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === YamlConfig.__pulumiType;
    }

    /**
     * The id of the application. This is required for all resources except global ones.
     */
    public readonly appId!: pulumi.Output<string | undefined>;
    /**
     * The raw YAML configuration.
     */
    public readonly content!: pulumi.Output<string>;
    /**
     * The name of the resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The path of the resource.
     */
    public readonly path!: pulumi.Output<string>;

    /**
     * Create a YamlConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: YamlConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: YamlConfigArgs | YamlConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as YamlConfigState | undefined;
            resourceInputs["appId"] = state ? state.appId : undefined;
            resourceInputs["content"] = state ? state.content : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
        } else {
            const args = argsOrState as YamlConfigArgs | undefined;
            if ((!args || args.content === undefined) && !opts.urn) {
                throw new Error("Missing required property 'content'");
            }
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["content"] = args ? args.content : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(YamlConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering YamlConfig resources.
 */
export interface YamlConfigState {
    /**
     * The id of the application. This is required for all resources except global ones.
     */
    appId?: pulumi.Input<string>;
    /**
     * The raw YAML configuration.
     */
    content?: pulumi.Input<string>;
    /**
     * The name of the resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The path of the resource.
     */
    path?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a YamlConfig resource.
 */
export interface YamlConfigArgs {
    /**
     * The id of the application. This is required for all resources except global ones.
     */
    appId?: pulumi.Input<string>;
    /**
     * The raw YAML configuration.
     */
    content: pulumi.Input<string>;
    /**
     * The path of the resource.
     */
    path: pulumi.Input<string>;
}