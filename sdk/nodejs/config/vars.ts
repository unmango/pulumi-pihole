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
        return __config.get("apiToken") ?? utilities.getEnv("PIHOLE_API_TOKEN");
    },
    enumerable: true,
});

/**
 * CA file to connect to Pi-hole with TLS
 */
export declare const caFile: string | undefined;
Object.defineProperty(exports, "caFile", {
    get() {
        return __config.get("caFile") ?? utilities.getEnv("PIHOLE_CA_FILE");
    },
    enumerable: true,
});

/**
 * The admin password used to login to the admin dashboard. Conflicts with `apiToken`.
 */
export declare const password: string | undefined;
Object.defineProperty(exports, "password", {
    get() {
        return __config.get("password") ?? utilities.getEnv("PIHOLE_PASSWORD");
    },
    enumerable: true,
});

/**
 * URL where Pi-hole is deployed
 */
export declare const url: string | undefined;
Object.defineProperty(exports, "url", {
    get() {
        return __config.get("url") ?? utilities.getEnv("PIHOLE_URL");
    },
    enumerable: true,
});

