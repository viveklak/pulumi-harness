// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.harness.platform.outputs.GetSecretSshkeySshSshPasswordCredential;
import com.pulumi.harness.platform.outputs.GetSecretSshkeySshSshkeyPathCredential;
import com.pulumi.harness.platform.outputs.GetSecretSshkeySshSshkeyReferenceCredential;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetSecretSshkeySsh {
    private String credentialType;
    private List<GetSecretSshkeySshSshPasswordCredential> sshPasswordCredentials;
    private List<GetSecretSshkeySshSshkeyPathCredential> sshkeyPathCredentials;
    private List<GetSecretSshkeySshSshkeyReferenceCredential> sshkeyReferenceCredentials;

    private GetSecretSshkeySsh() {}
    public String credentialType() {
        return this.credentialType;
    }
    public List<GetSecretSshkeySshSshPasswordCredential> sshPasswordCredentials() {
        return this.sshPasswordCredentials;
    }
    public List<GetSecretSshkeySshSshkeyPathCredential> sshkeyPathCredentials() {
        return this.sshkeyPathCredentials;
    }
    public List<GetSecretSshkeySshSshkeyReferenceCredential> sshkeyReferenceCredentials() {
        return this.sshkeyReferenceCredentials;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSecretSshkeySsh defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String credentialType;
        private List<GetSecretSshkeySshSshPasswordCredential> sshPasswordCredentials;
        private List<GetSecretSshkeySshSshkeyPathCredential> sshkeyPathCredentials;
        private List<GetSecretSshkeySshSshkeyReferenceCredential> sshkeyReferenceCredentials;
        public Builder() {}
        public Builder(GetSecretSshkeySsh defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.credentialType = defaults.credentialType;
    	      this.sshPasswordCredentials = defaults.sshPasswordCredentials;
    	      this.sshkeyPathCredentials = defaults.sshkeyPathCredentials;
    	      this.sshkeyReferenceCredentials = defaults.sshkeyReferenceCredentials;
        }

        @CustomType.Setter
        public Builder credentialType(String credentialType) {
            this.credentialType = Objects.requireNonNull(credentialType);
            return this;
        }
        @CustomType.Setter
        public Builder sshPasswordCredentials(List<GetSecretSshkeySshSshPasswordCredential> sshPasswordCredentials) {
            this.sshPasswordCredentials = Objects.requireNonNull(sshPasswordCredentials);
            return this;
        }
        public Builder sshPasswordCredentials(GetSecretSshkeySshSshPasswordCredential... sshPasswordCredentials) {
            return sshPasswordCredentials(List.of(sshPasswordCredentials));
        }
        @CustomType.Setter
        public Builder sshkeyPathCredentials(List<GetSecretSshkeySshSshkeyPathCredential> sshkeyPathCredentials) {
            this.sshkeyPathCredentials = Objects.requireNonNull(sshkeyPathCredentials);
            return this;
        }
        public Builder sshkeyPathCredentials(GetSecretSshkeySshSshkeyPathCredential... sshkeyPathCredentials) {
            return sshkeyPathCredentials(List.of(sshkeyPathCredentials));
        }
        @CustomType.Setter
        public Builder sshkeyReferenceCredentials(List<GetSecretSshkeySshSshkeyReferenceCredential> sshkeyReferenceCredentials) {
            this.sshkeyReferenceCredentials = Objects.requireNonNull(sshkeyReferenceCredentials);
            return this;
        }
        public Builder sshkeyReferenceCredentials(GetSecretSshkeySshSshkeyReferenceCredential... sshkeyReferenceCredentials) {
            return sshkeyReferenceCredentials(List.of(sshkeyReferenceCredentials));
        }
        public GetSecretSshkeySsh build() {
            final var o = new GetSecretSshkeySsh();
            o.credentialType = credentialType;
            o.sshPasswordCredentials = sshPasswordCredentials;
            o.sshkeyPathCredentials = sshkeyPathCredentials;
            o.sshkeyReferenceCredentials = sshkeyReferenceCredentials;
            return o;
        }
    }
}