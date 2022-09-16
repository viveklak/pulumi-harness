// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetGitConnectorCommitDetail {
    private String authorEmailId;
    private String authorName;
    private String message;

    private GetGitConnectorCommitDetail() {}
    public String authorEmailId() {
        return this.authorEmailId;
    }
    public String authorName() {
        return this.authorName;
    }
    public String message() {
        return this.message;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGitConnectorCommitDetail defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String authorEmailId;
        private String authorName;
        private String message;
        public Builder() {}
        public Builder(GetGitConnectorCommitDetail defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.authorEmailId = defaults.authorEmailId;
    	      this.authorName = defaults.authorName;
    	      this.message = defaults.message;
        }

        @CustomType.Setter
        public Builder authorEmailId(String authorEmailId) {
            this.authorEmailId = Objects.requireNonNull(authorEmailId);
            return this;
        }
        @CustomType.Setter
        public Builder authorName(String authorName) {
            this.authorName = Objects.requireNonNull(authorName);
            return this;
        }
        @CustomType.Setter
        public Builder message(String message) {
            this.message = Objects.requireNonNull(message);
            return this;
        }
        public GetGitConnectorCommitDetail build() {
            final var o = new GetGitConnectorCommitDetail();
            o.authorEmailId = authorEmailId;
            o.authorName = authorName;
            o.message = message;
            return o;
        }
    }
}