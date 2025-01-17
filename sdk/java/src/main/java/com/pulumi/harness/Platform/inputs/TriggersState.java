// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class TriggersState extends com.pulumi.resources.ResourceArgs {

    public static final TriggersState Empty = new TriggersState();

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
     * Unique identifier of the resource.
     * 
     */
    @Import(name="identifier")
    private @Nullable Output<String> identifier;

    /**
     * @return Unique identifier of the resource.
     * 
     */
    public Optional<Output<String>> identifier() {
        return Optional.ofNullable(this.identifier);
    }

    /**
     * if-Match
     * 
     */
    @Import(name="ifMatch")
    private @Nullable Output<String> ifMatch;

    /**
     * @return if-Match
     * 
     */
    public Optional<Output<String>> ifMatch() {
        return Optional.ofNullable(this.ifMatch);
    }

    /**
     * ignore error default false
     * 
     */
    @Import(name="ignoreError")
    private @Nullable Output<Boolean> ignoreError;

    /**
     * @return ignore error default false
     * 
     */
    public Optional<Output<Boolean>> ignoreError() {
        return Optional.ofNullable(this.ignoreError);
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

    /**
     * Identifier of the target pipeline
     * 
     */
    @Import(name="targetId")
    private @Nullable Output<String> targetId;

    /**
     * @return Identifier of the target pipeline
     * 
     */
    public Optional<Output<String>> targetId() {
        return Optional.ofNullable(this.targetId);
    }

    /**
     * trigger yaml
     * 
     */
    @Import(name="yaml")
    private @Nullable Output<String> yaml;

    /**
     * @return trigger yaml
     * 
     */
    public Optional<Output<String>> yaml() {
        return Optional.ofNullable(this.yaml);
    }

    private TriggersState() {}

    private TriggersState(TriggersState $) {
        this.description = $.description;
        this.identifier = $.identifier;
        this.ifMatch = $.ifMatch;
        this.ignoreError = $.ignoreError;
        this.name = $.name;
        this.orgId = $.orgId;
        this.projectId = $.projectId;
        this.tags = $.tags;
        this.targetId = $.targetId;
        this.yaml = $.yaml;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TriggersState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TriggersState $;

        public Builder() {
            $ = new TriggersState();
        }

        public Builder(TriggersState defaults) {
            $ = new TriggersState(Objects.requireNonNull(defaults));
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
         * @param identifier Unique identifier of the resource.
         * 
         * @return builder
         * 
         */
        public Builder identifier(@Nullable Output<String> identifier) {
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
         * @param ifMatch if-Match
         * 
         * @return builder
         * 
         */
        public Builder ifMatch(@Nullable Output<String> ifMatch) {
            $.ifMatch = ifMatch;
            return this;
        }

        /**
         * @param ifMatch if-Match
         * 
         * @return builder
         * 
         */
        public Builder ifMatch(String ifMatch) {
            return ifMatch(Output.of(ifMatch));
        }

        /**
         * @param ignoreError ignore error default false
         * 
         * @return builder
         * 
         */
        public Builder ignoreError(@Nullable Output<Boolean> ignoreError) {
            $.ignoreError = ignoreError;
            return this;
        }

        /**
         * @param ignoreError ignore error default false
         * 
         * @return builder
         * 
         */
        public Builder ignoreError(Boolean ignoreError) {
            return ignoreError(Output.of(ignoreError));
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

        /**
         * @param targetId Identifier of the target pipeline
         * 
         * @return builder
         * 
         */
        public Builder targetId(@Nullable Output<String> targetId) {
            $.targetId = targetId;
            return this;
        }

        /**
         * @param targetId Identifier of the target pipeline
         * 
         * @return builder
         * 
         */
        public Builder targetId(String targetId) {
            return targetId(Output.of(targetId));
        }

        /**
         * @param yaml trigger yaml
         * 
         * @return builder
         * 
         */
        public Builder yaml(@Nullable Output<String> yaml) {
            $.yaml = yaml;
            return this;
        }

        /**
         * @param yaml trigger yaml
         * 
         * @return builder
         * 
         */
        public Builder yaml(String yaml) {
            return yaml(Output.of(yaml));
        }

        public TriggersState build() {
            return $;
        }
    }

}
