// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.harness.platform.inputs.SecretSshkeyKerberosArgs;
import com.pulumi.harness.platform.inputs.SecretSshkeySshArgs;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretSshkeyState extends com.pulumi.resources.ResourceArgs {

    public static final SecretSshkeyState Empty = new SecretSshkeyState();

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
     * Kerberos authentication scheme
     * 
     */
    @Import(name="kerberos")
    private @Nullable Output<SecretSshkeyKerberosArgs> kerberos;

    /**
     * @return Kerberos authentication scheme
     * 
     */
    public Optional<Output<SecretSshkeyKerberosArgs>> kerberos() {
        return Optional.ofNullable(this.kerberos);
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
     * SSH port
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return SSH port
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
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
     * Kerberos authentication scheme
     * 
     */
    @Import(name="ssh")
    private @Nullable Output<SecretSshkeySshArgs> ssh;

    /**
     * @return Kerberos authentication scheme
     * 
     */
    public Optional<Output<SecretSshkeySshArgs>> ssh() {
        return Optional.ofNullable(this.ssh);
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

    private SecretSshkeyState() {}

    private SecretSshkeyState(SecretSshkeyState $) {
        this.description = $.description;
        this.identifier = $.identifier;
        this.kerberos = $.kerberos;
        this.name = $.name;
        this.orgId = $.orgId;
        this.port = $.port;
        this.projectId = $.projectId;
        this.ssh = $.ssh;
        this.tags = $.tags;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretSshkeyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretSshkeyState $;

        public Builder() {
            $ = new SecretSshkeyState();
        }

        public Builder(SecretSshkeyState defaults) {
            $ = new SecretSshkeyState(Objects.requireNonNull(defaults));
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
         * @param kerberos Kerberos authentication scheme
         * 
         * @return builder
         * 
         */
        public Builder kerberos(@Nullable Output<SecretSshkeyKerberosArgs> kerberos) {
            $.kerberos = kerberos;
            return this;
        }

        /**
         * @param kerberos Kerberos authentication scheme
         * 
         * @return builder
         * 
         */
        public Builder kerberos(SecretSshkeyKerberosArgs kerberos) {
            return kerberos(Output.of(kerberos));
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
         * @param port SSH port
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port SSH port
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
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
         * @param ssh Kerberos authentication scheme
         * 
         * @return builder
         * 
         */
        public Builder ssh(@Nullable Output<SecretSshkeySshArgs> ssh) {
            $.ssh = ssh;
            return this;
        }

        /**
         * @param ssh Kerberos authentication scheme
         * 
         * @return builder
         * 
         */
        public Builder ssh(SecretSshkeySshArgs ssh) {
            return ssh(Output.of(ssh));
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

        public SecretSshkeyState build() {
            return $;
        }
    }

}
