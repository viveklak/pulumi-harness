// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.harness.platform.outputs.GetAppDynamicsConnectorApiToken;
import com.pulumi.harness.platform.outputs.GetAppDynamicsConnectorUsernamePassword;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetAppDynamicsConnectorResult {
    /**
     * @return The App Dynamics account name.
     * 
     */
    private String accountName;
    /**
     * @return Authenticate to App Dynamics using api token.
     * 
     */
    private List<GetAppDynamicsConnectorApiToken> apiTokens;
    /**
     * @return Connect using only the delegates which have these tags.
     * 
     */
    private List<String> delegateSelectors;
    /**
     * @return Description of the resource.
     * 
     */
    private String description;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Unique identifier of the resource.
     * 
     */
    private @Nullable String identifier;
    /**
     * @return Name of the resource.
     * 
     */
    private @Nullable String name;
    /**
     * @return Unique identifier of the organization.
     * 
     */
    private @Nullable String orgId;
    /**
     * @return Unique identifier of the project.
     * 
     */
    private @Nullable String projectId;
    /**
     * @return Tags to associate with the resource. Tags should be in the form `name:value`.
     * 
     */
    private List<String> tags;
    /**
     * @return Url of the App Dynamics controller.
     * 
     */
    private String url;
    /**
     * @return Authenticate to App Dynamics using username and password.
     * 
     */
    private List<GetAppDynamicsConnectorUsernamePassword> usernamePasswords;

    private GetAppDynamicsConnectorResult() {}
    /**
     * @return The App Dynamics account name.
     * 
     */
    public String accountName() {
        return this.accountName;
    }
    /**
     * @return Authenticate to App Dynamics using api token.
     * 
     */
    public List<GetAppDynamicsConnectorApiToken> apiTokens() {
        return this.apiTokens;
    }
    /**
     * @return Connect using only the delegates which have these tags.
     * 
     */
    public List<String> delegateSelectors() {
        return this.delegateSelectors;
    }
    /**
     * @return Description of the resource.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Unique identifier of the resource.
     * 
     */
    public Optional<String> identifier() {
        return Optional.ofNullable(this.identifier);
    }
    /**
     * @return Name of the resource.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return Unique identifier of the organization.
     * 
     */
    public Optional<String> orgId() {
        return Optional.ofNullable(this.orgId);
    }
    /**
     * @return Unique identifier of the project.
     * 
     */
    public Optional<String> projectId() {
        return Optional.ofNullable(this.projectId);
    }
    /**
     * @return Tags to associate with the resource. Tags should be in the form `name:value`.
     * 
     */
    public List<String> tags() {
        return this.tags;
    }
    /**
     * @return Url of the App Dynamics controller.
     * 
     */
    public String url() {
        return this.url;
    }
    /**
     * @return Authenticate to App Dynamics using username and password.
     * 
     */
    public List<GetAppDynamicsConnectorUsernamePassword> usernamePasswords() {
        return this.usernamePasswords;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAppDynamicsConnectorResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String accountName;
        private List<GetAppDynamicsConnectorApiToken> apiTokens;
        private List<String> delegateSelectors;
        private String description;
        private String id;
        private @Nullable String identifier;
        private @Nullable String name;
        private @Nullable String orgId;
        private @Nullable String projectId;
        private List<String> tags;
        private String url;
        private List<GetAppDynamicsConnectorUsernamePassword> usernamePasswords;
        public Builder() {}
        public Builder(GetAppDynamicsConnectorResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accountName = defaults.accountName;
    	      this.apiTokens = defaults.apiTokens;
    	      this.delegateSelectors = defaults.delegateSelectors;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.identifier = defaults.identifier;
    	      this.name = defaults.name;
    	      this.orgId = defaults.orgId;
    	      this.projectId = defaults.projectId;
    	      this.tags = defaults.tags;
    	      this.url = defaults.url;
    	      this.usernamePasswords = defaults.usernamePasswords;
        }

        @CustomType.Setter
        public Builder accountName(String accountName) {
            this.accountName = Objects.requireNonNull(accountName);
            return this;
        }
        @CustomType.Setter
        public Builder apiTokens(List<GetAppDynamicsConnectorApiToken> apiTokens) {
            this.apiTokens = Objects.requireNonNull(apiTokens);
            return this;
        }
        public Builder apiTokens(GetAppDynamicsConnectorApiToken... apiTokens) {
            return apiTokens(List.of(apiTokens));
        }
        @CustomType.Setter
        public Builder delegateSelectors(List<String> delegateSelectors) {
            this.delegateSelectors = Objects.requireNonNull(delegateSelectors);
            return this;
        }
        public Builder delegateSelectors(String... delegateSelectors) {
            return delegateSelectors(List.of(delegateSelectors));
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder identifier(@Nullable String identifier) {
            this.identifier = identifier;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder orgId(@Nullable String orgId) {
            this.orgId = orgId;
            return this;
        }
        @CustomType.Setter
        public Builder projectId(@Nullable String projectId) {
            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder tags(List<String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }
        @CustomType.Setter
        public Builder url(String url) {
            this.url = Objects.requireNonNull(url);
            return this;
        }
        @CustomType.Setter
        public Builder usernamePasswords(List<GetAppDynamicsConnectorUsernamePassword> usernamePasswords) {
            this.usernamePasswords = Objects.requireNonNull(usernamePasswords);
            return this;
        }
        public Builder usernamePasswords(GetAppDynamicsConnectorUsernamePassword... usernamePasswords) {
            return usernamePasswords(List.of(usernamePasswords));
        }
        public GetAppDynamicsConnectorResult build() {
            final var o = new GetAppDynamicsConnectorResult();
            o.accountName = accountName;
            o.apiTokens = apiTokens;
            o.delegateSelectors = delegateSelectors;
            o.description = description;
            o.id = id;
            o.identifier = identifier;
            o.name = name;
            o.orgId = orgId;
            o.projectId = projectId;
            o.tags = tags;
            o.url = url;
            o.usernamePasswords = usernamePasswords;
            return o;
        }
    }
}
