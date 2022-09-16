// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Resource for creating an AWS CodeDeploy service. This resource uses the config-as-code API's. When updating the `name` or `path` of this resource you should typically also set the `createBeforeDestroy = true` lifecycle setting.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as harness from "@lbrlabs/pulumi-harness";
 *
 * const exampleApplication = new harness.Application("exampleApplication", {});
 * const exampleCodedeploy = new harness.service.Codedeploy("exampleCodedeploy", {
 *     appId: exampleApplication.id,
 *     description: "Service for AWS codedeploy applications.",
 * });
 * ```
 *
 * ## Import
 *
 * # Import using the Harness application id and service id
 *
 * ```sh
 *  $ pulumi import harness:service/codedeploy:Codedeploy example <app_id>/<svc_id>
 * ```
 */
export class Codedeploy extends pulumi.CustomResource {
    /**
     * Get an existing Codedeploy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CodedeployState, opts?: pulumi.CustomResourceOptions): Codedeploy {
        return new Codedeploy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'harness:service/codedeploy:Codedeploy';

    /**
     * Returns true if the given object is an instance of Codedeploy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Codedeploy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Codedeploy.__pulumiType;
    }

    /**
     * The id of the application the service belongs to
     */
    public readonly appId!: pulumi.Output<string>;
    /**
     * Description of th service
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the service
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Variables to be used in the service
     */
    public readonly variables!: pulumi.Output<outputs.service.CodedeployVariable[] | undefined>;

    /**
     * Create a Codedeploy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CodedeployArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CodedeployArgs | CodedeployState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CodedeployState | undefined;
            resourceInputs["appId"] = state ? state.appId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["variables"] = state ? state.variables : undefined;
        } else {
            const args = argsOrState as CodedeployArgs | undefined;
            if ((!args || args.appId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appId'");
            }
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["variables"] = args ? args.variables : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Codedeploy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Codedeploy resources.
 */
export interface CodedeployState {
    /**
     * The id of the application the service belongs to
     */
    appId?: pulumi.Input<string>;
    /**
     * Description of th service
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the service
     */
    name?: pulumi.Input<string>;
    /**
     * Variables to be used in the service
     */
    variables?: pulumi.Input<pulumi.Input<inputs.service.CodedeployVariable>[]>;
}

/**
 * The set of arguments for constructing a Codedeploy resource.
 */
export interface CodedeployArgs {
    /**
     * The id of the application the service belongs to
     */
    appId: pulumi.Input<string>;
    /**
     * Description of th service
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the service
     */
    name?: pulumi.Input<string>;
    /**
     * Variables to be used in the service
     */
    variables?: pulumi.Input<pulumi.Input<inputs.service.CodedeployVariable>[]>;
}