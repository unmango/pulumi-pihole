// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manages a Pi-hole CNAME record
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pihole from "@unmango/pulumi-pihole";
 *
 * const record = new pihole.CnameRecord("record", {
 *     domain: "foo.com",
 *     target: "bar.com",
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import pihole:index/cnameRecord:CnameRecord record foo.com
 * ```
 */
export class CnameRecord extends pulumi.CustomResource {
    /**
     * Get an existing CnameRecord resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CnameRecordState, opts?: pulumi.CustomResourceOptions): CnameRecord {
        return new CnameRecord(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'pihole:index/cnameRecord:CnameRecord';

    /**
     * Returns true if the given object is an instance of CnameRecord.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CnameRecord {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CnameRecord.__pulumiType;
    }

    /**
     * Domain to create a CNAME record for
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * Value of the CNAME record where traffic will be directed to from the configured domain value
     */
    public readonly target!: pulumi.Output<string>;

    /**
     * Create a CnameRecord resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CnameRecordArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CnameRecordArgs | CnameRecordState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CnameRecordState | undefined;
            resourceInputs["domain"] = state ? state.domain : undefined;
            resourceInputs["target"] = state ? state.target : undefined;
        } else {
            const args = argsOrState as CnameRecordArgs | undefined;
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            if ((!args || args.target === undefined) && !opts.urn) {
                throw new Error("Missing required property 'target'");
            }
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["target"] = args ? args.target : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CnameRecord.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CnameRecord resources.
 */
export interface CnameRecordState {
    /**
     * Domain to create a CNAME record for
     */
    domain?: pulumi.Input<string>;
    /**
     * Value of the CNAME record where traffic will be directed to from the configured domain value
     */
    target?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CnameRecord resource.
 */
export interface CnameRecordArgs {
    /**
     * Domain to create a CNAME record for
     */
    domain: pulumi.Input<string>;
    /**
     * Value of the CNAME record where traffic will be directed to from the configured domain value
     */
    target: pulumi.Input<string>;
}
