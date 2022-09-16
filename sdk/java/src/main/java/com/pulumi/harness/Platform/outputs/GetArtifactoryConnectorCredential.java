// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetArtifactoryConnectorCredential {
    private String passwordRef;
    private String username;
    private String usernameRef;

    private GetArtifactoryConnectorCredential() {}
    public String passwordRef() {
        return this.passwordRef;
    }
    public String username() {
        return this.username;
    }
    public String usernameRef() {
        return this.usernameRef;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetArtifactoryConnectorCredential defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String passwordRef;
        private String username;
        private String usernameRef;
        public Builder() {}
        public Builder(GetArtifactoryConnectorCredential defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.passwordRef = defaults.passwordRef;
    	      this.username = defaults.username;
    	      this.usernameRef = defaults.usernameRef;
        }

        @CustomType.Setter
        public Builder passwordRef(String passwordRef) {
            this.passwordRef = Objects.requireNonNull(passwordRef);
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            this.username = Objects.requireNonNull(username);
            return this;
        }
        @CustomType.Setter
        public Builder usernameRef(String usernameRef) {
            this.usernameRef = Objects.requireNonNull(usernameRef);
            return this;
        }
        public GetArtifactoryConnectorCredential build() {
            final var o = new GetArtifactoryConnectorCredential();
            o.passwordRef = passwordRef;
            o.username = username;
            o.usernameRef = usernameRef;
            return o;
        }
    }
}