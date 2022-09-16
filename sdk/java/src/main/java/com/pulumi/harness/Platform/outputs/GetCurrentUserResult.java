// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.harness.platform.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetCurrentUserResult {
    /**
     * @return Whether the user is an administrator.
     * 
     */
    private Boolean admin;
    /**
     * @return Billing frequency of the user.
     * 
     */
    private String billingFrequency;
    /**
     * @return Default account ID of the user.
     * 
     */
    private String defaultAccountId;
    /**
     * @return Edition of the platform being used.
     * 
     */
    private String edition;
    /**
     * @return Email address of the user.
     * 
     */
    private String email;
    /**
     * @return Whether the user&#39;s email address has been verified.
     * 
     */
    private Boolean emailVerified;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Intent of the user.
     * 
     */
    private String intent;
    /**
     * @return Whether 2FA is enabled for the user.
     * 
     */
    private Boolean isTwoFactorAuthEnabled;
    /**
     * @return Whether or not the user account is locked.
     * 
     */
    private Boolean locked;
    /**
     * @return Name of the user.
     * 
     */
    private String name;
    /**
     * @return Signup action of the user.
     * 
     */
    private String signupAction;
    /**
     * @return Token used to authenticate the user.
     * 
     */
    private String token;
    /**
     * @return Unique identifier of the user.
     * 
     */
    private String uuid;

    private GetCurrentUserResult() {}
    /**
     * @return Whether the user is an administrator.
     * 
     */
    public Boolean admin() {
        return this.admin;
    }
    /**
     * @return Billing frequency of the user.
     * 
     */
    public String billingFrequency() {
        return this.billingFrequency;
    }
    /**
     * @return Default account ID of the user.
     * 
     */
    public String defaultAccountId() {
        return this.defaultAccountId;
    }
    /**
     * @return Edition of the platform being used.
     * 
     */
    public String edition() {
        return this.edition;
    }
    /**
     * @return Email address of the user.
     * 
     */
    public String email() {
        return this.email;
    }
    /**
     * @return Whether the user&#39;s email address has been verified.
     * 
     */
    public Boolean emailVerified() {
        return this.emailVerified;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Intent of the user.
     * 
     */
    public String intent() {
        return this.intent;
    }
    /**
     * @return Whether 2FA is enabled for the user.
     * 
     */
    public Boolean isTwoFactorAuthEnabled() {
        return this.isTwoFactorAuthEnabled;
    }
    /**
     * @return Whether or not the user account is locked.
     * 
     */
    public Boolean locked() {
        return this.locked;
    }
    /**
     * @return Name of the user.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Signup action of the user.
     * 
     */
    public String signupAction() {
        return this.signupAction;
    }
    /**
     * @return Token used to authenticate the user.
     * 
     */
    public String token() {
        return this.token;
    }
    /**
     * @return Unique identifier of the user.
     * 
     */
    public String uuid() {
        return this.uuid;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCurrentUserResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean admin;
        private String billingFrequency;
        private String defaultAccountId;
        private String edition;
        private String email;
        private Boolean emailVerified;
        private String id;
        private String intent;
        private Boolean isTwoFactorAuthEnabled;
        private Boolean locked;
        private String name;
        private String signupAction;
        private String token;
        private String uuid;
        public Builder() {}
        public Builder(GetCurrentUserResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.admin = defaults.admin;
    	      this.billingFrequency = defaults.billingFrequency;
    	      this.defaultAccountId = defaults.defaultAccountId;
    	      this.edition = defaults.edition;
    	      this.email = defaults.email;
    	      this.emailVerified = defaults.emailVerified;
    	      this.id = defaults.id;
    	      this.intent = defaults.intent;
    	      this.isTwoFactorAuthEnabled = defaults.isTwoFactorAuthEnabled;
    	      this.locked = defaults.locked;
    	      this.name = defaults.name;
    	      this.signupAction = defaults.signupAction;
    	      this.token = defaults.token;
    	      this.uuid = defaults.uuid;
        }

        @CustomType.Setter
        public Builder admin(Boolean admin) {
            this.admin = Objects.requireNonNull(admin);
            return this;
        }
        @CustomType.Setter
        public Builder billingFrequency(String billingFrequency) {
            this.billingFrequency = Objects.requireNonNull(billingFrequency);
            return this;
        }
        @CustomType.Setter
        public Builder defaultAccountId(String defaultAccountId) {
            this.defaultAccountId = Objects.requireNonNull(defaultAccountId);
            return this;
        }
        @CustomType.Setter
        public Builder edition(String edition) {
            this.edition = Objects.requireNonNull(edition);
            return this;
        }
        @CustomType.Setter
        public Builder email(String email) {
            this.email = Objects.requireNonNull(email);
            return this;
        }
        @CustomType.Setter
        public Builder emailVerified(Boolean emailVerified) {
            this.emailVerified = Objects.requireNonNull(emailVerified);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder intent(String intent) {
            this.intent = Objects.requireNonNull(intent);
            return this;
        }
        @CustomType.Setter
        public Builder isTwoFactorAuthEnabled(Boolean isTwoFactorAuthEnabled) {
            this.isTwoFactorAuthEnabled = Objects.requireNonNull(isTwoFactorAuthEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder locked(Boolean locked) {
            this.locked = Objects.requireNonNull(locked);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder signupAction(String signupAction) {
            this.signupAction = Objects.requireNonNull(signupAction);
            return this;
        }
        @CustomType.Setter
        public Builder token(String token) {
            this.token = Objects.requireNonNull(token);
            return this;
        }
        @CustomType.Setter
        public Builder uuid(String uuid) {
            this.uuid = Objects.requireNonNull(uuid);
            return this;
        }
        public GetCurrentUserResult build() {
            final var o = new GetCurrentUserResult();
            o.admin = admin;
            o.billingFrequency = billingFrequency;
            o.defaultAccountId = defaultAccountId;
            o.edition = edition;
            o.email = email;
            o.emailVerified = emailVerified;
            o.id = id;
            o.intent = intent;
            o.isTwoFactorAuthEnabled = isTwoFactorAuthEnabled;
            o.locked = locked;
            o.name = name;
            o.signupAction = signupAction;
            o.token = token;
            o.uuid = uuid;
            return o;
        }
    }
}