// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pihole

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/unmango/pulumi-pihole/sdk/go/pihole/internal"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/unmango/pulumi-pihole/sdk/go/pihole"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := pihole.GetCnameRecords(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetCnameRecords(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetCnameRecordsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetCnameRecordsResult
	err := ctx.Invoke("pihole:index/getCnameRecords:getCnameRecords", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getCnameRecords.
type GetCnameRecordsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of CNAME Pi-hole records
	Records []GetCnameRecordsRecord `pulumi:"records"`
}

func GetCnameRecordsOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetCnameRecordsResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetCnameRecordsResult, error) {
		r, err := GetCnameRecords(ctx, opts...)
		var s GetCnameRecordsResult
		if r != nil {
			s = *r
		}
		return s, err
	}).(GetCnameRecordsResultOutput)
}

// A collection of values returned by getCnameRecords.
type GetCnameRecordsResultOutput struct{ *pulumi.OutputState }

func (GetCnameRecordsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCnameRecordsResult)(nil)).Elem()
}

func (o GetCnameRecordsResultOutput) ToGetCnameRecordsResultOutput() GetCnameRecordsResultOutput {
	return o
}

func (o GetCnameRecordsResultOutput) ToGetCnameRecordsResultOutputWithContext(ctx context.Context) GetCnameRecordsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetCnameRecordsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCnameRecordsResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of CNAME Pi-hole records
func (o GetCnameRecordsResultOutput) Records() GetCnameRecordsRecordArrayOutput {
	return o.ApplyT(func(v GetCnameRecordsResult) []GetCnameRecordsRecord { return v.Records }).(GetCnameRecordsRecordArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCnameRecordsResultOutput{})
}
