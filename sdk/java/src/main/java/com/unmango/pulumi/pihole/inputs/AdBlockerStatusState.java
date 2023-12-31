// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.pulumi.pihole.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AdBlockerStatusState extends com.pulumi.resources.ResourceArgs {

    public static final AdBlockerStatusState Empty = new AdBlockerStatusState();

    /**
     * Whether to enable the Pi-hole ad blocker
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return Whether to enable the Pi-hole ad blocker
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    private AdBlockerStatusState() {}

    private AdBlockerStatusState(AdBlockerStatusState $) {
        this.enabled = $.enabled;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AdBlockerStatusState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AdBlockerStatusState $;

        public Builder() {
            $ = new AdBlockerStatusState();
        }

        public Builder(AdBlockerStatusState defaults) {
            $ = new AdBlockerStatusState(Objects.requireNonNull(defaults));
        }

        /**
         * @param enabled Whether to enable the Pi-hole ad blocker
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled Whether to enable the Pi-hole ad blocker
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        public AdBlockerStatusState build() {
            return $;
        }
    }

}
