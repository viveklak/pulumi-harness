// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.harness.platform.outputs.GetSecretSshkeyKerbero;
import com.pulumi.harness.platform.outputs.GetSecretSshkeySsh;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetSecretSshkeyResult {
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
     * @return Kerberos authentication scheme
     * 
     */
    private List<GetSecretSshkeyKerbero> kerberos;
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
     * @return SSH port
     * 
     */
    private Integer port;
    /**
     * @return Unique identifier of the project.
     * 
     */
    private @Nullable String projectId;
    /**
     * @return Kerberos authentication scheme
     * 
     */
    private List<GetSecretSshkeySsh> sshes;
    /**
     * @return Tags to associate with the resource. Tags should be in the form `name:value`.
     * 
     */
    private List<String> tags;

    private GetSecretSshkeyResult() {}
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
     * @return Kerberos authentication scheme
     * 
     */
    public List<GetSecretSshkeyKerbero> kerberos() {
        return this.kerberos;
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
     * @return SSH port
     * 
     */
    public Integer port() {
        return this.port;
    }
    /**
     * @return Unique identifier of the project.
     * 
     */
    public Optional<String> projectId() {
        return Optional.ofNullable(this.projectId);
    }
    /**
     * @return Kerberos authentication scheme
     * 
     */
    public List<GetSecretSshkeySsh> sshes() {
        return this.sshes;
    }
    /**
     * @return Tags to associate with the resource. Tags should be in the form `name:value`.
     * 
     */
    public List<String> tags() {
        return this.tags;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSecretSshkeyResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private String id;
        private @Nullable String identifier;
        private List<GetSecretSshkeyKerbero> kerberos;
        private @Nullable String name;
        private @Nullable String orgId;
        private Integer port;
        private @Nullable String projectId;
        private List<GetSecretSshkeySsh> sshes;
        private List<String> tags;
        public Builder() {}
        public Builder(GetSecretSshkeyResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.identifier = defaults.identifier;
    	      this.kerberos = defaults.kerberos;
    	      this.name = defaults.name;
    	      this.orgId = defaults.orgId;
    	      this.port = defaults.port;
    	      this.projectId = defaults.projectId;
    	      this.sshes = defaults.sshes;
    	      this.tags = defaults.tags;
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
        public Builder kerberos(List<GetSecretSshkeyKerbero> kerberos) {
            this.kerberos = Objects.requireNonNull(kerberos);
            return this;
        }
        public Builder kerberos(GetSecretSshkeyKerbero... kerberos) {
            return kerberos(List.of(kerberos));
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
        public Builder port(Integer port) {
            this.port = Objects.requireNonNull(port);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(@Nullable String projectId) {
            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder sshes(List<GetSecretSshkeySsh> sshes) {
            this.sshes = Objects.requireNonNull(sshes);
            return this;
        }
        public Builder sshes(GetSecretSshkeySsh... sshes) {
            return sshes(List.of(sshes));
        }
        @CustomType.Setter
        public Builder tags(List<String> tags) {
            this.tags = Objects.requireNonNull(tags);
            return this;
        }
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }
        public GetSecretSshkeyResult build() {
            final var o = new GetSecretSshkeyResult();
            o.description = description;
            o.id = id;
            o.identifier = identifier;
            o.kerberos = kerberos;
            o.name = name;
            o.orgId = orgId;
            o.port = port;
            o.projectId = projectId;
            o.sshes = sshes;
            o.tags = tags;
            return o;
        }
    }
}
