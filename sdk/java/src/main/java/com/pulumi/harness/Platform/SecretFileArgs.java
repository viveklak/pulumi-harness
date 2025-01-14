// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretFileArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretFileArgs Empty = new SecretFileArgs();

    /**
     * Description of the resource.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the resource.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Path of the file containing secret value
     * 
     */
    @Import(name="filePath", required=true)
    private Output<String> filePath;

    /**
     * @return Path of the file containing secret value
     * 
     */
    public Output<String> filePath() {
        return this.filePath;
    }

    /**
     * Unique identifier of the resource.
     * 
     */
    @Import(name="identifier", required=true)
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
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the resource.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Unique identifier of the organization.
     * 
     */
    @Import(name="orgId")
    private @Nullable Output<String> orgId;

    /**
     * @return Unique identifier of the organization.
     * 
     */
    public Optional<Output<String>> orgId() {
        return Optional.ofNullable(this.orgId);
    }

    /**
     * Unique identifier of the project.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return Unique identifier of the project.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * Identifier of the Secret Manager used to manage the secret.
     * 
     */
    @Import(name="secretManagerIdentifier", required=true)
    private Output<String> secretManagerIdentifier;

    /**
     * @return Identifier of the Secret Manager used to manage the secret.
     * 
     */
    public Output<String> secretManagerIdentifier() {
        return this.secretManagerIdentifier;
    }

    /**
     * Tags to associate with the resource. Tags should be in the form `name:value`.
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return Tags to associate with the resource. Tags should be in the form `name:value`.
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    private SecretFileArgs() {}

    private SecretFileArgs(SecretFileArgs $) {
        this.description = $.description;
        this.filePath = $.filePath;
        this.identifier = $.identifier;
        this.name = $.name;
        this.orgId = $.orgId;
        this.projectId = $.projectId;
        this.secretManagerIdentifier = $.secretManagerIdentifier;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretFileArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretFileArgs $;

        public Builder() {
            $ = new SecretFileArgs();
        }

        public Builder(SecretFileArgs defaults) {
            $ = new SecretFileArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Description of the resource.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the resource.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param filePath Path of the file containing secret value
         * 
         * @return builder
         * 
         */
        public Builder filePath(Output<String> filePath) {
            $.filePath = filePath;
            return this;
        }

        /**
         * @param filePath Path of the file containing secret value
         * 
         * @return builder
         * 
         */
        public Builder filePath(String filePath) {
            return filePath(Output.of(filePath));
        }

        /**
         * @param identifier Unique identifier of the resource.
         * 
         * @return builder
         * 
         */
        public Builder identifier(Output<String> identifier) {
            $.identifier = identifier;
            return this;
        }

        /**
         * @param identifier Unique identifier of the resource.
         * 
         * @return builder
         * 
         */
        public Builder identifier(String identifier) {
            return identifier(Output.of(identifier));
        }

        /**
         * @param name Name of the resource.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the resource.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param orgId Unique identifier of the organization.
         * 
         * @return builder
         * 
         */
        public Builder orgId(@Nullable Output<String> orgId) {
            $.orgId = orgId;
            return this;
        }

        /**
         * @param orgId Unique identifier of the organization.
         * 
         * @return builder
         * 
         */
        public Builder orgId(String orgId) {
            return orgId(Output.of(orgId));
        }

        /**
         * @param projectId Unique identifier of the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId Unique identifier of the project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param secretManagerIdentifier Identifier of the Secret Manager used to manage the secret.
         * 
         * @return builder
         * 
         */
        public Builder secretManagerIdentifier(Output<String> secretManagerIdentifier) {
            $.secretManagerIdentifier = secretManagerIdentifier;
            return this;
        }

        /**
         * @param secretManagerIdentifier Identifier of the Secret Manager used to manage the secret.
         * 
         * @return builder
         * 
         */
        public Builder secretManagerIdentifier(String secretManagerIdentifier) {
            return secretManagerIdentifier(Output.of(secretManagerIdentifier));
        }

        /**
         * @param tags Tags to associate with the resource. Tags should be in the form `name:value`.
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Tags to associate with the resource. Tags should be in the form `name:value`.
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags Tags to associate with the resource. Tags should be in the form `name:value`.
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        public SecretFileArgs build() {
            $.filePath = Objects.requireNonNull($.filePath, "expected parameter 'filePath' to be non-null");
            $.identifier = Objects.requireNonNull($.identifier, "expected parameter 'identifier' to be non-null");
            $.secretManagerIdentifier = Objects.requireNonNull($.secretManagerIdentifier, "expected parameter 'secretManagerIdentifier' to be non-null");
            return $;
        }
    }

}
