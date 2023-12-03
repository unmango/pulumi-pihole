// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.pulumi.pihole.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetGroupsGroup {
    private String description;
    private Boolean enabled;
    /**
     * @return The ID of this resource.
     * 
     */
    private Integer id;
    private String name;

    private GetGroupsGroup() {}
    public String description() {
        return this.description;
    }
    public Boolean enabled() {
        return this.enabled;
    }
    /**
     * @return The ID of this resource.
     * 
     */
    public Integer id() {
        return this.id;
    }
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGroupsGroup defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String description;
        private Boolean enabled;
        private Integer id;
        private String name;
        public Builder() {}
        public Builder(GetGroupsGroup defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.enabled = defaults.enabled;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        @CustomType.Setter
        public Builder enabled(Boolean enabled) {
            this.enabled = Objects.requireNonNull(enabled);
            return this;
        }
        @CustomType.Setter
        public Builder id(Integer id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public GetGroupsGroup build() {
            final var _resultValue = new GetGroupsGroup();
            _resultValue.description = description;
            _resultValue.enabled = enabled;
            _resultValue.id = id;
            _resultValue.name = name;
            return _resultValue;
        }
    }
}