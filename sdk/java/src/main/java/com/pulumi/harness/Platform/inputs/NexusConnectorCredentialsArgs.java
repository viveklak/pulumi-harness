// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NexusConnectorCredentialsArgs extends com.pulumi.resources.ResourceArgs {

    public static final NexusConnectorCredentialsArgs Empty = new NexusConnectorCredentialsArgs();

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

    private NexusConnectorCredentialsArgs() {}

    private NexusConnectorCredentialsArgs(NexusConnectorCredentialsArgs $) {
        this.passwordRef = $.passwordRef;
        this.username = $.username;
        this.usernameRef = $.usernameRef;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NexusConnectorCredentialsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NexusConnectorCredentialsArgs $;

        public Builder() {
            $ = new NexusConnectorCredentialsArgs();
        }

        public Builder(NexusConnectorCredentialsArgs defaults) {
            $ = new NexusConnectorCredentialsArgs(Objects.requireNonNull(defaults));
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

        public NexusConnectorCredentialsArgs build() {
            $.passwordRef = Objects.requireNonNull($.passwordRef, "expected parameter 'passwordRef' to be non-null");
            return $;
        }
    }

}