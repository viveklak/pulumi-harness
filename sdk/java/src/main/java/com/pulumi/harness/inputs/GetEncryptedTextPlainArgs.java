// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.harness.inputs.GetEncryptedTextUsageScope;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetEncryptedTextPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetEncryptedTextPlainArgs Empty = new GetEncryptedTextPlainArgs();

    /**
     * Unique identifier of the encrypted secret
     * 
     */
    @Import(name="id")
    private @Nullable String id;

    /**
     * @return Unique identifier of the encrypted secret
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }

    /**
     * The name of the encrypted secret
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return The name of the encrypted secret
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * This block is used for scoping the resource to a specific set of applications or environments.
     * 
     */
    @Import(name="usageScopes")
    private @Nullable List<GetEncryptedTextUsageScope> usageScopes;

    /**
     * @return This block is used for scoping the resource to a specific set of applications or environments.
     * 
     */
    public Optional<List<GetEncryptedTextUsageScope>> usageScopes() {
        return Optional.ofNullable(this.usageScopes);
    }

    private GetEncryptedTextPlainArgs() {}

    private GetEncryptedTextPlainArgs(GetEncryptedTextPlainArgs $) {
        this.id = $.id;
        this.name = $.name;
        this.usageScopes = $.usageScopes;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetEncryptedTextPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetEncryptedTextPlainArgs $;

        public Builder() {
            $ = new GetEncryptedTextPlainArgs();
        }

        public Builder(GetEncryptedTextPlainArgs defaults) {
            $ = new GetEncryptedTextPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param id Unique identifier of the encrypted secret
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable String id) {
            $.id = id;
            return this;
        }

        /**
         * @param name The name of the encrypted secret
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param usageScopes This block is used for scoping the resource to a specific set of applications or environments.
         * 
         * @return builder
         * 
         */
        public Builder usageScopes(@Nullable List<GetEncryptedTextUsageScope> usageScopes) {
            $.usageScopes = usageScopes;
            return this;
        }

        /**
         * @param usageScopes This block is used for scoping the resource to a specific set of applications or environments.
         * 
         * @return builder
         * 
         */
        public Builder usageScopes(GetEncryptedTextUsageScope... usageScopes) {
            return usageScopes(List.of(usageScopes));
        }

        public GetEncryptedTextPlainArgs build() {
            return $;
        }
    }

}