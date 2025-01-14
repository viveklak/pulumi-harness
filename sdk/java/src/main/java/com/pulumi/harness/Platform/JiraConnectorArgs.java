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


public final class JiraConnectorArgs extends com.pulumi.resources.ResourceArgs {

    public static final JiraConnectorArgs Empty = new JiraConnectorArgs();

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
     * Reference to a secret containing the password to use for authentication.
     * 
     */
    @Import(name="passwordRef", required=true)
    private Output<String> passwordRef;

    /**
     * @return Reference to a secret containing the password to use for authentication.
     * 
     */
    public Output<String> passwordRef() {
        return this.passwordRef;
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
     * Url of the Jira server.
     * 
     */
    @Import(name="url", required=true)
    private Output<String> url;

    /**
     * @return Url of the Jira server.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     * Username to use for authentication.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return Username to use for authentication.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    /**
     * Reference to a secret containing the username to use for authentication.
     * 
     */
    @Import(name="usernameRef")
    private @Nullable Output<String> usernameRef;

    /**
     * @return Reference to a secret containing the username to use for authentication.
     * 
     */
    public Optional<Output<String>> usernameRef() {
        return Optional.ofNullable(this.usernameRef);
    }

    private JiraConnectorArgs() {}

    private JiraConnectorArgs(JiraConnectorArgs $) {
        this.delegateSelectors = $.delegateSelectors;
        this.description = $.description;
        this.identifier = $.identifier;
        this.name = $.name;
        this.orgId = $.orgId;
        this.passwordRef = $.passwordRef;
        this.projectId = $.projectId;
        this.tags = $.tags;
        this.url = $.url;
        this.username = $.username;
        this.usernameRef = $.usernameRef;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(JiraConnectorArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private JiraConnectorArgs $;

        public Builder() {
            $ = new JiraConnectorArgs();
        }

        public Builder(JiraConnectorArgs defaults) {
            $ = new JiraConnectorArgs(Objects.requireNonNull(defaults));
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
         * @param passwordRef Reference to a secret containing the password to use for authentication.
         * 
         * @return builder
         * 
         */
        public Builder passwordRef(Output<String> passwordRef) {
            $.passwordRef = passwordRef;
            return this;
        }

        /**
         * @param passwordRef Reference to a secret containing the password to use for authentication.
         * 
         * @return builder
         * 
         */
        public Builder passwordRef(String passwordRef) {
            return passwordRef(Output.of(passwordRef));
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
         * @param url Url of the Jira server.
         * 
         * @return builder
         * 
         */
        public Builder url(Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url Url of the Jira server.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        /**
         * @param username Username to use for authentication.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username Username to use for authentication.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        /**
         * @param usernameRef Reference to a secret containing the username to use for authentication.
         * 
         * @return builder
         * 
         */
        public Builder usernameRef(@Nullable Output<String> usernameRef) {
            $.usernameRef = usernameRef;
            return this;
        }

        /**
         * @param usernameRef Reference to a secret containing the username to use for authentication.
         * 
         * @return builder
         * 
         */
        public Builder usernameRef(String usernameRef) {
            return usernameRef(Output.of(usernameRef));
        }

        public JiraConnectorArgs build() {
            $.identifier = Objects.requireNonNull($.identifier, "expected parameter 'identifier' to be non-null");
            $.passwordRef = Objects.requireNonNull($.passwordRef, "expected parameter 'passwordRef' to be non-null");
            $.url = Objects.requireNonNull($.url, "expected parameter 'url' to be non-null");
            return $;
        }
    }

}
