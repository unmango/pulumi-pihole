// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pihole from "@pulumi/pihole";
 *
 * const records = pihole.getCnameRecords({});
 * ```
 */
export function getGroups(opts?: pulumi.InvokeOptions): Promise<GetGroupsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("pihole:index/getGroups:getGroups", {
    }, opts);
}

/**
 * A collection of values returned by getGroups.
 */
export interface GetGroupsResult {
    /**
     * List of groups to manage client lists and block lists
     */
    readonly groups: outputs.GetGroupsGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pihole from "@pulumi/pihole";
 *
 * const records = pihole.getCnameRecords({});
 * ```
 */
export function getGroupsOutput(opts?: pulumi.InvokeOptions): pulumi.Output<GetGroupsResult> {
    return pulumi.output(getGroups(opts))
}
