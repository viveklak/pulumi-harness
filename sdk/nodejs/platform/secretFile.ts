// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource for creating a secret of type secret file in Harness.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as harness from "@pulumi/harness";
 *
 * const example = new harness.platform.SecretFile("example", {
 *     description: "test",
 *     filePath: "file_path",
 *     identifier: "identifier",
 *     secretManagerIdentifier: "harnessSecretManager",
 *     tags: ["foo:bar"],
 * });
 * ```
 *
 * ## Import
 *
 * Import using secret file id
 *
 * ```sh
 *  $ pulumi import harness:platform/secretFile:SecretFile example <secret_file_id>
 * ```
 */
export class SecretFile extends pulumi.CustomResource {
    /**
     * Get an existing SecretFile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretFileState, opts?: pulumi.CustomResourceOptions): SecretFile {
        return new SecretFile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'harness:platform/secretFile:SecretFile';

    /**
     * Returns true if the given object is an instance of SecretFile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretFile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretFile.__pulumiType;
    }

    /**
     * Description of the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Path of the file containing secret value
     */
    public readonly filePath!: pulumi.Output<string>;
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
     * Identifier of the Secret Manager used to manage the secret.
     */
    public readonly secretManagerIdentifier!: pulumi.Output<string>;
    /**
     * Tags to associate with the resource. Tags should be in the form `name:value`.
     */
    public readonly tags!: pulumi.Output<string[] | undefined>;

    /**
     * Create a SecretFile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretFileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretFileArgs | SecretFileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretFileState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["filePath"] = state ? state.filePath : undefined;
            resourceInputs["identifier"] = state ? state.identifier : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["projectId"] = state ? state.projectId : undefined;
            resourceInputs["secretManagerIdentifier"] = state ? state.secretManagerIdentifier : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as SecretFileArgs | undefined;
            if ((!args || args.filePath === undefined) && !opts.urn) {
                throw new Error("Missing required property 'filePath'");
            }
            if ((!args || args.identifier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identifier'");
            }
            if ((!args || args.secretManagerIdentifier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secretManagerIdentifier'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["filePath"] = args ? args.filePath : undefined;
            resourceInputs["identifier"] = args ? args.identifier : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["projectId"] = args ? args.projectId : undefined;
            resourceInputs["secretManagerIdentifier"] = args ? args.secretManagerIdentifier : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecretFile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretFile resources.
 */
export interface SecretFileState {
    /**
     * Description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Path of the file containing secret value
     */
    filePath?: pulumi.Input<string>;
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
     * Identifier of the Secret Manager used to manage the secret.
     */
    secretManagerIdentifier?: pulumi.Input<string>;
    /**
     * Tags to associate with the resource. Tags should be in the form `name:value`.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a SecretFile resource.
 */
export interface SecretFileArgs {
    /**
     * Description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Path of the file containing secret value
     */
    filePath: pulumi.Input<string>;
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
     * Identifier of the Secret Manager used to manage the secret.
     */
    secretManagerIdentifier: pulumi.Input<string>;
    /**
     * Tags to associate with the resource. Tags should be in the form `name:value`.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}
