// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// SelfSubjectAccessReview checks whether or the current user can perform an action.  Not filling in a spec.namespace means "in all namespaces".  Self is a special case, because users should always be able to check whether they can perform an action
type SelfSubjectAccessReview struct {
	pulumi.CustomResourceState

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrOutput `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrOutput `pulumi:"metadata"`
	// Spec holds information about the request being evaluated.  user and groups must be empty
	Spec SelfSubjectAccessReviewSpecOutput `pulumi:"spec"`
	// Status is filled in by the server and indicates whether the request is allowed or not
	Status SubjectAccessReviewStatusPtrOutput `pulumi:"status"`
}

// NewSelfSubjectAccessReview registers a new resource with the given unique name, arguments, and options.
func NewSelfSubjectAccessReview(ctx *pulumi.Context,
	name string, args *SelfSubjectAccessReviewArgs, opts ...pulumi.ResourceOption) (*SelfSubjectAccessReview, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Spec == nil {
		return nil, errors.New("invalid value for required argument 'Spec'")
	}
	args.ApiVersion = pulumi.StringPtr("authorization.k8s.io/v1")
	args.Kind = pulumi.StringPtr("SelfSubjectAccessReview")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("kubernetes:authorization.k8s.io/v1beta1:SelfSubjectAccessReview"),
		},
	})
	opts = append(opts, aliases)
	var resource SelfSubjectAccessReview
	err := ctx.RegisterResource("kubernetes:authorization.k8s.io/v1:SelfSubjectAccessReview", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSelfSubjectAccessReview gets an existing SelfSubjectAccessReview resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSelfSubjectAccessReview(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SelfSubjectAccessReviewState, opts ...pulumi.ResourceOption) (*SelfSubjectAccessReview, error) {
	var resource SelfSubjectAccessReview
	err := ctx.ReadResource("kubernetes:authorization.k8s.io/v1:SelfSubjectAccessReview", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SelfSubjectAccessReview resources.
type selfSubjectAccessReviewState struct {
}

type SelfSubjectAccessReviewState struct {
}

func (SelfSubjectAccessReviewState) ElementType() reflect.Type {
	return reflect.TypeOf((*selfSubjectAccessReviewState)(nil)).Elem()
}

type selfSubjectAccessReviewArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion *string `pulumi:"apiVersion"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind *string `pulumi:"kind"`
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *metav1.ObjectMeta `pulumi:"metadata"`
	// Spec holds information about the request being evaluated.  user and groups must be empty
	Spec SelfSubjectAccessReviewSpec `pulumi:"spec"`
}

// The set of arguments for constructing a SelfSubjectAccessReview resource.
type SelfSubjectAccessReviewArgs struct {
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	ApiVersion pulumi.StringPtrInput
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind pulumi.StringPtrInput
	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata metav1.ObjectMetaPtrInput
	// Spec holds information about the request being evaluated.  user and groups must be empty
	Spec SelfSubjectAccessReviewSpecInput
}

func (SelfSubjectAccessReviewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*selfSubjectAccessReviewArgs)(nil)).Elem()
}

type SelfSubjectAccessReviewInput interface {
	pulumi.Input

	ToSelfSubjectAccessReviewOutput() SelfSubjectAccessReviewOutput
	ToSelfSubjectAccessReviewOutputWithContext(ctx context.Context) SelfSubjectAccessReviewOutput
}

func (*SelfSubjectAccessReview) ElementType() reflect.Type {
	return reflect.TypeOf((*SelfSubjectAccessReview)(nil))
}

func (i *SelfSubjectAccessReview) ToSelfSubjectAccessReviewOutput() SelfSubjectAccessReviewOutput {
	return i.ToSelfSubjectAccessReviewOutputWithContext(context.Background())
}

func (i *SelfSubjectAccessReview) ToSelfSubjectAccessReviewOutputWithContext(ctx context.Context) SelfSubjectAccessReviewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSubjectAccessReviewOutput)
}

func (i *SelfSubjectAccessReview) ToSelfSubjectAccessReviewPtrOutput() SelfSubjectAccessReviewPtrOutput {
	return i.ToSelfSubjectAccessReviewPtrOutputWithContext(context.Background())
}

func (i *SelfSubjectAccessReview) ToSelfSubjectAccessReviewPtrOutputWithContext(ctx context.Context) SelfSubjectAccessReviewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSubjectAccessReviewPtrOutput)
}

type SelfSubjectAccessReviewPtrInput interface {
	pulumi.Input

	ToSelfSubjectAccessReviewPtrOutput() SelfSubjectAccessReviewPtrOutput
	ToSelfSubjectAccessReviewPtrOutputWithContext(ctx context.Context) SelfSubjectAccessReviewPtrOutput
}

type selfSubjectAccessReviewPtrType SelfSubjectAccessReviewArgs

func (*selfSubjectAccessReviewPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SelfSubjectAccessReview)(nil))
}

func (i *selfSubjectAccessReviewPtrType) ToSelfSubjectAccessReviewPtrOutput() SelfSubjectAccessReviewPtrOutput {
	return i.ToSelfSubjectAccessReviewPtrOutputWithContext(context.Background())
}

func (i *selfSubjectAccessReviewPtrType) ToSelfSubjectAccessReviewPtrOutputWithContext(ctx context.Context) SelfSubjectAccessReviewPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSubjectAccessReviewPtrOutput)
}

// SelfSubjectAccessReviewArrayInput is an input type that accepts SelfSubjectAccessReviewArray and SelfSubjectAccessReviewArrayOutput values.
// You can construct a concrete instance of `SelfSubjectAccessReviewArrayInput` via:
//
//          SelfSubjectAccessReviewArray{ SelfSubjectAccessReviewArgs{...} }
type SelfSubjectAccessReviewArrayInput interface {
	pulumi.Input

	ToSelfSubjectAccessReviewArrayOutput() SelfSubjectAccessReviewArrayOutput
	ToSelfSubjectAccessReviewArrayOutputWithContext(context.Context) SelfSubjectAccessReviewArrayOutput
}

type SelfSubjectAccessReviewArray []SelfSubjectAccessReviewInput

func (SelfSubjectAccessReviewArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SelfSubjectAccessReview)(nil)).Elem()
}

func (i SelfSubjectAccessReviewArray) ToSelfSubjectAccessReviewArrayOutput() SelfSubjectAccessReviewArrayOutput {
	return i.ToSelfSubjectAccessReviewArrayOutputWithContext(context.Background())
}

func (i SelfSubjectAccessReviewArray) ToSelfSubjectAccessReviewArrayOutputWithContext(ctx context.Context) SelfSubjectAccessReviewArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSubjectAccessReviewArrayOutput)
}

// SelfSubjectAccessReviewMapInput is an input type that accepts SelfSubjectAccessReviewMap and SelfSubjectAccessReviewMapOutput values.
// You can construct a concrete instance of `SelfSubjectAccessReviewMapInput` via:
//
//          SelfSubjectAccessReviewMap{ "key": SelfSubjectAccessReviewArgs{...} }
type SelfSubjectAccessReviewMapInput interface {
	pulumi.Input

	ToSelfSubjectAccessReviewMapOutput() SelfSubjectAccessReviewMapOutput
	ToSelfSubjectAccessReviewMapOutputWithContext(context.Context) SelfSubjectAccessReviewMapOutput
}

type SelfSubjectAccessReviewMap map[string]SelfSubjectAccessReviewInput

func (SelfSubjectAccessReviewMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SelfSubjectAccessReview)(nil)).Elem()
}

func (i SelfSubjectAccessReviewMap) ToSelfSubjectAccessReviewMapOutput() SelfSubjectAccessReviewMapOutput {
	return i.ToSelfSubjectAccessReviewMapOutputWithContext(context.Background())
}

func (i SelfSubjectAccessReviewMap) ToSelfSubjectAccessReviewMapOutputWithContext(ctx context.Context) SelfSubjectAccessReviewMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SelfSubjectAccessReviewMapOutput)
}

type SelfSubjectAccessReviewOutput struct{ *pulumi.OutputState }

func (SelfSubjectAccessReviewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SelfSubjectAccessReview)(nil))
}

func (o SelfSubjectAccessReviewOutput) ToSelfSubjectAccessReviewOutput() SelfSubjectAccessReviewOutput {
	return o
}

func (o SelfSubjectAccessReviewOutput) ToSelfSubjectAccessReviewOutputWithContext(ctx context.Context) SelfSubjectAccessReviewOutput {
	return o
}

func (o SelfSubjectAccessReviewOutput) ToSelfSubjectAccessReviewPtrOutput() SelfSubjectAccessReviewPtrOutput {
	return o.ToSelfSubjectAccessReviewPtrOutputWithContext(context.Background())
}

func (o SelfSubjectAccessReviewOutput) ToSelfSubjectAccessReviewPtrOutputWithContext(ctx context.Context) SelfSubjectAccessReviewPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SelfSubjectAccessReview) *SelfSubjectAccessReview {
		return &v
	}).(SelfSubjectAccessReviewPtrOutput)
}

type SelfSubjectAccessReviewPtrOutput struct{ *pulumi.OutputState }

func (SelfSubjectAccessReviewPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SelfSubjectAccessReview)(nil))
}

func (o SelfSubjectAccessReviewPtrOutput) ToSelfSubjectAccessReviewPtrOutput() SelfSubjectAccessReviewPtrOutput {
	return o
}

func (o SelfSubjectAccessReviewPtrOutput) ToSelfSubjectAccessReviewPtrOutputWithContext(ctx context.Context) SelfSubjectAccessReviewPtrOutput {
	return o
}

func (o SelfSubjectAccessReviewPtrOutput) Elem() SelfSubjectAccessReviewOutput {
	return o.ApplyT(func(v *SelfSubjectAccessReview) SelfSubjectAccessReview {
		if v != nil {
			return *v
		}
		var ret SelfSubjectAccessReview
		return ret
	}).(SelfSubjectAccessReviewOutput)
}

type SelfSubjectAccessReviewArrayOutput struct{ *pulumi.OutputState }

func (SelfSubjectAccessReviewArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SelfSubjectAccessReview)(nil))
}

func (o SelfSubjectAccessReviewArrayOutput) ToSelfSubjectAccessReviewArrayOutput() SelfSubjectAccessReviewArrayOutput {
	return o
}

func (o SelfSubjectAccessReviewArrayOutput) ToSelfSubjectAccessReviewArrayOutputWithContext(ctx context.Context) SelfSubjectAccessReviewArrayOutput {
	return o
}

func (o SelfSubjectAccessReviewArrayOutput) Index(i pulumi.IntInput) SelfSubjectAccessReviewOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SelfSubjectAccessReview {
		return vs[0].([]SelfSubjectAccessReview)[vs[1].(int)]
	}).(SelfSubjectAccessReviewOutput)
}

type SelfSubjectAccessReviewMapOutput struct{ *pulumi.OutputState }

func (SelfSubjectAccessReviewMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SelfSubjectAccessReview)(nil))
}

func (o SelfSubjectAccessReviewMapOutput) ToSelfSubjectAccessReviewMapOutput() SelfSubjectAccessReviewMapOutput {
	return o
}

func (o SelfSubjectAccessReviewMapOutput) ToSelfSubjectAccessReviewMapOutputWithContext(ctx context.Context) SelfSubjectAccessReviewMapOutput {
	return o
}

func (o SelfSubjectAccessReviewMapOutput) MapIndex(k pulumi.StringInput) SelfSubjectAccessReviewOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SelfSubjectAccessReview {
		return vs[0].(map[string]SelfSubjectAccessReview)[vs[1].(string)]
	}).(SelfSubjectAccessReviewOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SelfSubjectAccessReviewInput)(nil)).Elem(), &SelfSubjectAccessReview{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelfSubjectAccessReviewPtrInput)(nil)).Elem(), &SelfSubjectAccessReview{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelfSubjectAccessReviewArrayInput)(nil)).Elem(), SelfSubjectAccessReviewArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SelfSubjectAccessReviewMapInput)(nil)).Elem(), SelfSubjectAccessReviewMap{})
	pulumi.RegisterOutputType(SelfSubjectAccessReviewOutput{})
	pulumi.RegisterOutputType(SelfSubjectAccessReviewPtrOutput{})
	pulumi.RegisterOutputType(SelfSubjectAccessReviewArrayOutput{})
	pulumi.RegisterOutputType(SelfSubjectAccessReviewMapOutput{})
}