// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NewrelicConnectorState extends com.pulumi.resources.ResourceArgs {

    public static final NewrelicConnectorState Empty = new NewrelicConnectorState();

    /**
     * Account ID of the NewRelic account.
     * 
     */
    @Import(name="accountId")
    private @Nullable Output<String> accountId;

    /**
     * @return Account ID of the NewRelic account.
     * 
     */
    public Optional<Output<String>> accountId() {
        return Optional.ofNullable(this.accountId);
    }

    /**
     * Reference to the Harness secret containing the api key.
     * 
     */
    @Import(name="apiKeyRef")
    private @Nullable Output<String> apiKeyRef;

    /**
     * @return Reference to the Harness secret containing the api key.
     * 
     */
    public Optional<Output<String>> apiKeyRef() {
        return Optional.ofNullable(this.apiKeyRef);
    }

    /**
     * Connect using only the delegates which have these tags.
     * 
     */
    @Import(name="delegateSelectors")
    private @Nullable Output<List<String>> delegateSelectors;

    /**
     * @return Connect using only the delegates which have these tags.
     * 
     */
    public Optional<Output<List<String>>> delegateSelectors() {
        return Optional.ofNullable(this.delegateSelectors);
    }

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
     * Url of the NewRelic server.
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return Url of the NewRelic server.
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    private NewrelicConnectorState() {}

    private NewrelicConnectorState(NewrelicConnectorState $) {
        this.accountId = $.accountId;
        this.apiKeyRef = $.apiKeyRef;
        this.delegateSelectors = $.delegateSelectors;
        this.description = $.description;
        this.identifier = $.identifier;
        this.name = $.name;
        this.orgId = $.orgId;
        this.projectId = $.projectId;
        this.tags = $.tags;
        this.url = $.url;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NewrelicConnectorState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NewrelicConnectorState $;

        public Builder() {
            $ = new NewrelicConnectorState();
        }

        public Builder(NewrelicConnectorState defaults) {
            $ = new NewrelicConnectorState(Objects.requireNonNull(defaults));
        }

        /**
         * @param accountId Account ID of the NewRelic account.
         * 
         * @return builder
         * 
         */
        public Builder accountId(@Nullable Output<String> accountId) {
            $.accountId = accountId;
            return this;
        }

        /**
         * @param accountId Account ID of the NewRelic account.
         * 
         * @return builder
         * 
         */
        public Builder accountId(String accountId) {
            return accountId(Output.of(accountId));
        }

        /**
         * @param apiKeyRef Reference to the Harness secret containing the api key.
         * 
         * @return builder
         * 
         */
        public Builder apiKeyRef(@Nullable Output<String> apiKeyRef) {
            $.apiKeyRef = apiKeyRef;
            return this;
        }

        /**
         * @param apiKeyRef Reference to the Harness secret containing the api key.
         * 
         * @return builder
         * 
         */
        public Builder apiKeyRef(String apiKeyRef) {
            return apiKeyRef(Output.of(apiKeyRef));
        }

        /**
         * @param delegateSelectors Connect using only the delegates which have these tags.
         * 
         * @return builder
         * 
         */
        public Builder delegateSelectors(@Nullable Output<List<String>> delegateSelectors) {
            $.delegateSelectors = delegateSelectors;
            return this;
        }

        /**
         * @param delegateSelectors Connect using only the delegates which have these tags.
         * 
         * @return builder
         * 
         */
        public Builder delegateSelectors(List<String> delegateSelectors) {
            return delegateSelectors(Output.of(delegateSelectors));
        }

        /**
         * @param delegateSelectors Connect using only the delegates which have these tags.
         * 
         * @return builder
         * 
         */
        public Builder delegateSelectors(String... delegateSelectors) {
            return delegateSelectors(List.of(delegateSelectors));
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
         * @param url Url of the NewRelic server.
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url Url of the NewRelic server.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        public NewrelicConnectorState build() {
            return $;
        }
    }

}
