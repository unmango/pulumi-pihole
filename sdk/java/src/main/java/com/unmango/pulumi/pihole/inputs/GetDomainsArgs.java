// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.pulumi.pihole.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDomainsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDomainsArgs Empty = new GetDomainsArgs();

    /**
     * Filter on allowed or denied domains. Must be either &#39;allow&#39; or &#39;deny&#39;.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Filter on allowed or denied domains. Must be either &#39;allow&#39; or &#39;deny&#39;.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private GetDomainsArgs() {}

    private GetDomainsArgs(GetDomainsArgs $) {
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDomainsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDomainsArgs $;

        public Builder() {
            $ = new GetDomainsArgs();
        }

        public Builder(GetDomainsArgs defaults) {
            $ = new GetDomainsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param type Filter on allowed or denied domains. Must be either &#39;allow&#39; or &#39;deny&#39;.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Filter on allowed or denied domains. Must be either &#39;allow&#39; or &#39;deny&#39;.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public GetDomainsArgs build() {
            return $;
        }
    }

}
