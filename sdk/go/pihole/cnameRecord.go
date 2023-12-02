// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pihole

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/unmango/pulumi-pihole/sdk/go/pihole/internal"
)

// Manages a Pi-hole CNAME record
//
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
//			_, err := pihole.NewCnameRecord(ctx, "record", &pihole.CnameRecordArgs{
//				Domain: pulumi.String("foo.com"),
//				Target: pulumi.String("bar.com"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ```sh
//
//	$ pulumi import pihole:index/cnameRecord:CnameRecord record foo.com
//
// ```
type CnameRecord struct {
	pulumi.CustomResourceState

	// Domain to create a CNAME record for
	Domain pulumi.StringOutput `pulumi:"domain"`
	// Value of the CNAME record where traffic will be directed to from the configured domain value
	Target pulumi.StringOutput `pulumi:"target"`
}

// NewCnameRecord registers a new resource with the given unique name, arguments, and options.
func NewCnameRecord(ctx *pulumi.Context,
	name string, args *CnameRecordArgs, opts ...pulumi.ResourceOption) (*CnameRecord, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Domain == nil {
		return nil, errors.New("invalid value for required argument 'Domain'")
	}
	if args.Target == nil {
		return nil, errors.New("invalid value for required argument 'Target'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CnameRecord
	err := ctx.RegisterResource("pihole:index/cnameRecord:CnameRecord", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCnameRecord gets an existing CnameRecord resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCnameRecord(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CnameRecordState, opts ...pulumi.ResourceOption) (*CnameRecord, error) {
	var resource CnameRecord
	err := ctx.ReadResource("pihole:index/cnameRecord:CnameRecord", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CnameRecord resources.
type cnameRecordState struct {
	// Domain to create a CNAME record for
	Domain *string `pulumi:"domain"`
	// Value of the CNAME record where traffic will be directed to from the configured domain value
	Target *string `pulumi:"target"`
}

type CnameRecordState struct {
	// Domain to create a CNAME record for
	Domain pulumi.StringPtrInput
	// Value of the CNAME record where traffic will be directed to from the configured domain value
	Target pulumi.StringPtrInput
}

func (CnameRecordState) ElementType() reflect.Type {
	return reflect.TypeOf((*cnameRecordState)(nil)).Elem()
}

type cnameRecordArgs struct {
	// Domain to create a CNAME record for
	Domain string `pulumi:"domain"`
	// Value of the CNAME record where traffic will be directed to from the configured domain value
	Target string `pulumi:"target"`
}

// The set of arguments for constructing a CnameRecord resource.
type CnameRecordArgs struct {
	// Domain to create a CNAME record for
	Domain pulumi.StringInput
	// Value of the CNAME record where traffic will be directed to from the configured domain value
	Target pulumi.StringInput
}

func (CnameRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cnameRecordArgs)(nil)).Elem()
}

type CnameRecordInput interface {
	pulumi.Input

	ToCnameRecordOutput() CnameRecordOutput
	ToCnameRecordOutputWithContext(ctx context.Context) CnameRecordOutput
}

func (*CnameRecord) ElementType() reflect.Type {
	return reflect.TypeOf((**CnameRecord)(nil)).Elem()
}

func (i *CnameRecord) ToCnameRecordOutput() CnameRecordOutput {
	return i.ToCnameRecordOutputWithContext(context.Background())
}

func (i *CnameRecord) ToCnameRecordOutputWithContext(ctx context.Context) CnameRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CnameRecordOutput)
}

func (i *CnameRecord) ToOutput(ctx context.Context) pulumix.Output[*CnameRecord] {
	return pulumix.Output[*CnameRecord]{
		OutputState: i.ToCnameRecordOutputWithContext(ctx).OutputState,
	}
}

// CnameRecordArrayInput is an input type that accepts CnameRecordArray and CnameRecordArrayOutput values.
// You can construct a concrete instance of `CnameRecordArrayInput` via:
//
//	CnameRecordArray{ CnameRecordArgs{...} }
type CnameRecordArrayInput interface {
	pulumi.Input

	ToCnameRecordArrayOutput() CnameRecordArrayOutput
	ToCnameRecordArrayOutputWithContext(context.Context) CnameRecordArrayOutput
}

type CnameRecordArray []CnameRecordInput

func (CnameRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CnameRecord)(nil)).Elem()
}

func (i CnameRecordArray) ToCnameRecordArrayOutput() CnameRecordArrayOutput {
	return i.ToCnameRecordArrayOutputWithContext(context.Background())
}

func (i CnameRecordArray) ToCnameRecordArrayOutputWithContext(ctx context.Context) CnameRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CnameRecordArrayOutput)
}

func (i CnameRecordArray) ToOutput(ctx context.Context) pulumix.Output[[]*CnameRecord] {
	return pulumix.Output[[]*CnameRecord]{
		OutputState: i.ToCnameRecordArrayOutputWithContext(ctx).OutputState,
	}
}

// CnameRecordMapInput is an input type that accepts CnameRecordMap and CnameRecordMapOutput values.
// You can construct a concrete instance of `CnameRecordMapInput` via:
//
//	CnameRecordMap{ "key": CnameRecordArgs{...} }
type CnameRecordMapInput interface {
	pulumi.Input

	ToCnameRecordMapOutput() CnameRecordMapOutput
	ToCnameRecordMapOutputWithContext(context.Context) CnameRecordMapOutput
}

type CnameRecordMap map[string]CnameRecordInput

func (CnameRecordMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CnameRecord)(nil)).Elem()
}

func (i CnameRecordMap) ToCnameRecordMapOutput() CnameRecordMapOutput {
	return i.ToCnameRecordMapOutputWithContext(context.Background())
}

func (i CnameRecordMap) ToCnameRecordMapOutputWithContext(ctx context.Context) CnameRecordMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CnameRecordMapOutput)
}

func (i CnameRecordMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*CnameRecord] {
	return pulumix.Output[map[string]*CnameRecord]{
		OutputState: i.ToCnameRecordMapOutputWithContext(ctx).OutputState,
	}
}

type CnameRecordOutput struct{ *pulumi.OutputState }

func (CnameRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CnameRecord)(nil)).Elem()
}

func (o CnameRecordOutput) ToCnameRecordOutput() CnameRecordOutput {
	return o
}

func (o CnameRecordOutput) ToCnameRecordOutputWithContext(ctx context.Context) CnameRecordOutput {
	return o
}

func (o CnameRecordOutput) ToOutput(ctx context.Context) pulumix.Output[*CnameRecord] {
	return pulumix.Output[*CnameRecord]{
		OutputState: o.OutputState,
	}
}

// Domain to create a CNAME record for
func (o CnameRecordOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *CnameRecord) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// Value of the CNAME record where traffic will be directed to from the configured domain value
func (o CnameRecordOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v *CnameRecord) pulumi.StringOutput { return v.Target }).(pulumi.StringOutput)
}

type CnameRecordArrayOutput struct{ *pulumi.OutputState }

func (CnameRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CnameRecord)(nil)).Elem()
}

func (o CnameRecordArrayOutput) ToCnameRecordArrayOutput() CnameRecordArrayOutput {
	return o
}

func (o CnameRecordArrayOutput) ToCnameRecordArrayOutputWithContext(ctx context.Context) CnameRecordArrayOutput {
	return o
}

func (o CnameRecordArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*CnameRecord] {
	return pulumix.Output[[]*CnameRecord]{
		OutputState: o.OutputState,
	}
}

func (o CnameRecordArrayOutput) Index(i pulumi.IntInput) CnameRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CnameRecord {
		return vs[0].([]*CnameRecord)[vs[1].(int)]
	}).(CnameRecordOutput)
}

type CnameRecordMapOutput struct{ *pulumi.OutputState }

func (CnameRecordMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CnameRecord)(nil)).Elem()
}

func (o CnameRecordMapOutput) ToCnameRecordMapOutput() CnameRecordMapOutput {
	return o
}

func (o CnameRecordMapOutput) ToCnameRecordMapOutputWithContext(ctx context.Context) CnameRecordMapOutput {
	return o
}

func (o CnameRecordMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*CnameRecord] {
	return pulumix.Output[map[string]*CnameRecord]{
		OutputState: o.OutputState,
	}
}

func (o CnameRecordMapOutput) MapIndex(k pulumi.StringInput) CnameRecordOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CnameRecord {
		return vs[0].(map[string]*CnameRecord)[vs[1].(string)]
	}).(CnameRecordOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CnameRecordInput)(nil)).Elem(), &CnameRecord{})
	pulumi.RegisterInputType(reflect.TypeOf((*CnameRecordArrayInput)(nil)).Elem(), CnameRecordArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CnameRecordMapInput)(nil)).Elem(), CnameRecordMap{})
	pulumi.RegisterOutputType(CnameRecordOutput{})
	pulumi.RegisterOutputType(CnameRecordArrayOutput{})
	pulumi.RegisterOutputType(CnameRecordMapOutput{})
}
