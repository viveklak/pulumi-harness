// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.harness.platform.inputs.AppDynamicsConnectorApiTokenArgs;
import com.pulumi.harness.platform.inputs.AppDynamicsConnectorUsernamePasswordArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AppDynamicsConnectorState extends com.pulumi.resources.ResourceArgs {

    public static final AppDynamicsConnectorState Empty = new AppDynamicsConnectorState();

    /**
     * The App Dynamics account name.
     * 
     */
    @Import(name="accountName")
    private @Nullable Output<String> accountName;

    /**
     * @return The App Dynamics account name.
     * 
     */
    public Optional<Output<String>> accountName() {
        return Optional.ofNullable(this.accountName);
    }

    /**
     * Authenticate to App Dynamics using api token.
     * 
     */
    @Import(name="apiToken")
    private @Nullable Output<AppDynamicsConnectorApiTokenArgs> apiToken;

    /**
     * @return Authenticate to App Dynamics using api token.
     * 
     */
    public Optional<Output<AppDynamicsConnectorApiTokenArgs>> apiToken() {
        return Optional.ofNullable(this.apiToken);
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
     * Url of the App Dynamics controller.
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return Url of the App Dynamics controller.
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    /**
     * Authenticate to App Dynamics using username and password.
     * 
     */
    @Import(name="usernamePassword")
    private @Nullable Output<AppDynamicsConnectorUsernamePasswordArgs> usernamePassword;

    /**
     * @return Authenticate to App Dynamics using username and password.
     * 
     */
    public Optional<Output<AppDynamicsConnectorUsernamePasswordArgs>> usernamePassword() {
        return Optional.ofNullable(this.usernamePassword);
    }

    private AppDynamicsConnectorState() {}

    private AppDynamicsConnectorState(AppDynamicsConnectorState $) {
        this.accountName = $.accountName;
        this.apiToken = $.apiToken;
        this.delegateSelectors = $.delegateSelectors;
        this.description = $.description;
        this.identifier = $.identifier;
        this.name = $.name;
        this.orgId = $.orgId;
        this.projectId = $.projectId;
        this.tags = $.tags;
        this.url = $.url;
        this.usernamePassword = $.usernamePassword;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AppDynamicsConnectorState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AppDynamicsConnectorState $;

        public Builder() {
            $ = new AppDynamicsConnectorState();
        }

        public Builder(AppDynamicsConnectorState defaults) {
            $ = new AppDynamicsConnectorState(Objects.requireNonNull(defaults));
        }

        /**
         * @param accountName The App Dynamics account name.
         * 
         * @return builder
         * 
         */
        public Builder accountName(@Nullable Output<String> accountName) {
            $.accountName = accountName;
            return this;
        }

        /**
         * @param accountName The App Dynamics account name.
         * 
         * @return builder
         * 
         */
        public Builder accountName(String accountName) {
            return accountName(Output.of(accountName));
        }

        /**
         * @param apiToken Authenticate to App Dynamics using api token.
         * 
         * @return builder
         * 
         */
        public Builder apiToken(@Nullable Output<AppDynamicsConnectorApiTokenArgs> apiToken) {
            $.apiToken = apiToken;
            return this;
        }

        /**
         * @param apiToken Authenticate to App Dynamics using api token.
         * 
         * @return builder
         * 
         */
        public Builder apiToken(AppDynamicsConnectorApiTokenArgs apiToken) {
            return apiToken(Output.of(apiToken));
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
         * @param url Url of the App Dynamics controller.
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url Url of the App Dynamics controller.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        /**
         * @param usernamePassword Authenticate to App Dynamics using username and password.
         * 
         * @return builder
         * 
         */
        public Builder usernamePassword(@Nullable Output<AppDynamicsConnectorUsernamePasswordArgs> usernamePassword) {
            $.usernamePassword = usernamePassword;
            return this;
        }

        /**
         * @param usernamePassword Authenticate to App Dynamics using username and password.
         * 
         * @return builder
         * 
         */
        public Builder usernamePassword(AppDynamicsConnectorUsernamePasswordArgs usernamePassword) {
            return usernamePassword(Output.of(usernamePassword));
        }

        public AppDynamicsConnectorState build() {
            return $;
        }
    }

}
