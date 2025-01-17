// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.harness.Utilities;
import com.pulumi.harness.platform.AwsKmsConnectorArgs;
import com.pulumi.harness.platform.inputs.AwsKmsConnectorState;
import com.pulumi.harness.platform.outputs.AwsKmsConnectorCredentials;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Resource for creating an AWS KMS connector.
 * 
 */
@ResourceType(type="harness:platform/awsKmsConnector:AwsKmsConnector")
public class AwsKmsConnector extends com.pulumi.resources.CustomResource {
    /**
     * A reference to the Harness secret containing the ARN of the AWS KMS.
     * 
     */
    @Export(name="arnRef", type=String.class, parameters={})
    private Output<String> arnRef;

    /**
     * @return A reference to the Harness secret containing the ARN of the AWS KMS.
     * 
     */
    public Output<String> arnRef() {
        return this.arnRef;
    }
    /**
     * The credentials to use for connecting to aws.
     * 
     */
    @Export(name="credentials", type=AwsKmsConnectorCredentials.class, parameters={})
    private Output<AwsKmsConnectorCredentials> credentials;

    /**
     * @return The credentials to use for connecting to aws.
     * 
     */
    public Output<AwsKmsConnectorCredentials> credentials() {
        return this.credentials;
    }
    /**
     * Connect using only the delegates which have these tags.
     * 
     */
    @Export(name="delegateSelectors", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> delegateSelectors;

    /**
     * @return Connect using only the delegates which have these tags.
     * 
     */
    public Output<Optional<List<String>>> delegateSelectors() {
        return Codegen.optional(this.delegateSelectors);
    }
    /**
     * Description of the resource.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the resource.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Unique identifier of the resource.
     * 
     */
    @Export(name="identifier", type=String.class, parameters={})
    private Output<String> identifier;

    /**
     * @return Unique identifier of the resource.
     * 
     */
    public Output<String> identifier() {
        return this.identifier;
    }
    /**
     * Name of the resource.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name of the resource.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Unique identifier of the organization.
     * 
     */
    @Export(name="orgId", type=String.class, parameters={})
    private Output</* @Nullable */ String> orgId;

    /**
     * @return Unique identifier of the organization.
     * 
     */
    public Output<Optional<String>> orgId() {
        return Codegen.optional(this.orgId);
    }
    /**
     * Unique identifier of the project.
     * 
     */
    @Export(name="projectId", type=String.class, parameters={})
    private Output</* @Nullable */ String> projectId;

    /**
     * @return Unique identifier of the project.
     * 
     */
    public Output<Optional<String>> projectId() {
        return Codegen.optional(this.projectId);
    }
    /**
     * The AWS region where the AWS Secret Manager is.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return The AWS region where the AWS Secret Manager is.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Tags to associate with the resource. Tags should be in the form `name:value`.
     * 
     */
    @Export(name="tags", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return Tags to associate with the resource. Tags should be in the form `name:value`.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AwsKmsConnector(String name) {
        this(name, AwsKmsConnectorArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AwsKmsConnector(String name, AwsKmsConnectorArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AwsKmsConnector(String name, AwsKmsConnectorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("harness:platform/awsKmsConnector:AwsKmsConnector", name, args == null ? AwsKmsConnectorArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AwsKmsConnector(String name, Output<String> id, @Nullable AwsKmsConnectorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("harness:platform/awsKmsConnector:AwsKmsConnector", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static AwsKmsConnector get(String name, Output<String> id, @Nullable AwsKmsConnectorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AwsKmsConnector(name, id, state, options);
    }
}
