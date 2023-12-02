// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface GetCnameRecordsRecord {
    domain: string;
    target: string;
}

export interface GetDnsRecordsRecord {
    domain: string;
    ip: string;
}

export interface GetDomainsDomain {
    comment: string;
    domain: string;
    enabled: boolean;
    groupIds: number[];
    /**
     * The ID of this resource.
     */
    id: number;
    /**
     * Filter on allowed or denied domains. Must be either 'allow' or 'deny'.
     */
    type: string;
    wildcard: boolean;
}

export interface GetGroupsGroup {
    description: string;
    enabled: boolean;
    /**
     * The ID of this resource.
     */
    id: number;
    name: string;
}
