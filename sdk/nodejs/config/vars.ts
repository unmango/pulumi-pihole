// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("pihole");

/**
 * Experimental: Pi-hole API token. Conflicts with `password`.
 */
export declare const apiToken: string | undefined;
Object.defineProperty(exports, "apiToken", {
    get() {
        return __config.get("apiToken");
    },
    enumerable: true,
});

/**
 * CA file to connect to Pi-hole with TLS
 */
export declare const caFile: string | undefined;
Object.defineProperty(exports, "caFile", {
    get() {
        return __config.get("caFile");
    },
    enumerable: true,
});

/**
 * The admin password used to login to the admin dashboard. Conflicts with `api_token`.
 */
export declare const password: string | undefined;
Object.defineProperty(exports, "password", {
    get() {
        return __config.get("password");
    },
    enumerable: true,
});

/**
 * URL where Pi-hole is deployed
 */
export declare const url: string | undefined;
Object.defineProperty(exports, "url", {
    get() {
        return __config.get("url");
    },
    enumerable: true,
});
