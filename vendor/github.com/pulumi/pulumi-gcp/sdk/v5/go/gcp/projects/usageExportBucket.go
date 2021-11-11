// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package projects

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows creation and management of a Google Cloud Platform project.
//
// Projects created with this resource must be associated with an Organization.
// See the [Organization documentation](https://cloud.google.com/resource-manager/docs/quickstarts) for more details.
//
// The user or service account that is running this provider when creating a `organizations.Project`
// resource must have `roles/resourcemanager.projectCreator` on the specified organization. See the
// [Access Control for Organizations Using IAM](https://cloud.google.com/resource-manager/docs/access-control-org)
// doc for more information.
//
// > This resource reads the specified billing account on every provider apply and plan operation so you must have permissions on the specified billing account.
//
// To get more information about projects, see:
//
// * [API documentation](https://cloud.google.com/resource-manager/reference/rest/v1/projects)
// * How-to Guides
//     * [Creating and managing projects](https://cloud.google.com/resource-manager/docs/creating-managing-projects)
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := organizations.NewProject(ctx, "myProject", &organizations.ProjectArgs{
// 			OrgId:     pulumi.String("1234567"),
// 			ProjectId: pulumi.String("your-project-id"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// To create a project under a specific folder
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		department1, err := organizations.NewFolder(ctx, "department1", &organizations.FolderArgs{
// 			DisplayName: pulumi.String("Department 1"),
// 			Parent:      pulumi.String("organizations/1234567"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = organizations.NewProject(ctx, "myProject_in_a_folder", &organizations.ProjectArgs{
// 			ProjectId: pulumi.String("your-project-id"),
// 			FolderId:  department1.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Projects can be imported using the `project_id`, e.g.
//
// ```sh
//  $ pulumi import gcp:projects/usageExportBucket:UsageExportBucket my_project your-project-id
// ```
type UsageExportBucket struct {
	pulumi.CustomResourceState

	// The bucket to store reports in.
	BucketName pulumi.StringOutput `pulumi:"bucketName"`
	// A prefix for the reports, for instance, the project name.
	Prefix pulumi.StringPtrOutput `pulumi:"prefix"`
	// The project to set the export bucket on. If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewUsageExportBucket registers a new resource with the given unique name, arguments, and options.
func NewUsageExportBucket(ctx *pulumi.Context,
	name string, args *UsageExportBucketArgs, opts ...pulumi.ResourceOption) (*UsageExportBucket, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BucketName == nil {
		return nil, errors.New("invalid value for required argument 'BucketName'")
	}
	var resource UsageExportBucket
	err := ctx.RegisterResource("gcp:projects/usageExportBucket:UsageExportBucket", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUsageExportBucket gets an existing UsageExportBucket resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUsageExportBucket(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UsageExportBucketState, opts ...pulumi.ResourceOption) (*UsageExportBucket, error) {
	var resource UsageExportBucket
	err := ctx.ReadResource("gcp:projects/usageExportBucket:UsageExportBucket", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UsageExportBucket resources.
type usageExportBucketState struct {
	// The bucket to store reports in.
	BucketName *string `pulumi:"bucketName"`
	// A prefix for the reports, for instance, the project name.
	Prefix *string `pulumi:"prefix"`
	// The project to set the export bucket on. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type UsageExportBucketState struct {
	// The bucket to store reports in.
	BucketName pulumi.StringPtrInput
	// A prefix for the reports, for instance, the project name.
	Prefix pulumi.StringPtrInput
	// The project to set the export bucket on. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (UsageExportBucketState) ElementType() reflect.Type {
	return reflect.TypeOf((*usageExportBucketState)(nil)).Elem()
}

type usageExportBucketArgs struct {
	// The bucket to store reports in.
	BucketName string `pulumi:"bucketName"`
	// A prefix for the reports, for instance, the project name.
	Prefix *string `pulumi:"prefix"`
	// The project to set the export bucket on. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a UsageExportBucket resource.
type UsageExportBucketArgs struct {
	// The bucket to store reports in.
	BucketName pulumi.StringInput
	// A prefix for the reports, for instance, the project name.
	Prefix pulumi.StringPtrInput
	// The project to set the export bucket on. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (UsageExportBucketArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*usageExportBucketArgs)(nil)).Elem()
}

type UsageExportBucketInput interface {
	pulumi.Input

	ToUsageExportBucketOutput() UsageExportBucketOutput
	ToUsageExportBucketOutputWithContext(ctx context.Context) UsageExportBucketOutput
}

func (*UsageExportBucket) ElementType() reflect.Type {
	return reflect.TypeOf((*UsageExportBucket)(nil))
}

func (i *UsageExportBucket) ToUsageExportBucketOutput() UsageExportBucketOutput {
	return i.ToUsageExportBucketOutputWithContext(context.Background())
}

func (i *UsageExportBucket) ToUsageExportBucketOutputWithContext(ctx context.Context) UsageExportBucketOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsageExportBucketOutput)
}

func (i *UsageExportBucket) ToUsageExportBucketPtrOutput() UsageExportBucketPtrOutput {
	return i.ToUsageExportBucketPtrOutputWithContext(context.Background())
}

func (i *UsageExportBucket) ToUsageExportBucketPtrOutputWithContext(ctx context.Context) UsageExportBucketPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsageExportBucketPtrOutput)
}

type UsageExportBucketPtrInput interface {
	pulumi.Input

	ToUsageExportBucketPtrOutput() UsageExportBucketPtrOutput
	ToUsageExportBucketPtrOutputWithContext(ctx context.Context) UsageExportBucketPtrOutput
}

type usageExportBucketPtrType UsageExportBucketArgs

func (*usageExportBucketPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UsageExportBucket)(nil))
}

func (i *usageExportBucketPtrType) ToUsageExportBucketPtrOutput() UsageExportBucketPtrOutput {
	return i.ToUsageExportBucketPtrOutputWithContext(context.Background())
}

func (i *usageExportBucketPtrType) ToUsageExportBucketPtrOutputWithContext(ctx context.Context) UsageExportBucketPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsageExportBucketPtrOutput)
}

// UsageExportBucketArrayInput is an input type that accepts UsageExportBucketArray and UsageExportBucketArrayOutput values.
// You can construct a concrete instance of `UsageExportBucketArrayInput` via:
//
//          UsageExportBucketArray{ UsageExportBucketArgs{...} }
type UsageExportBucketArrayInput interface {
	pulumi.Input

	ToUsageExportBucketArrayOutput() UsageExportBucketArrayOutput
	ToUsageExportBucketArrayOutputWithContext(context.Context) UsageExportBucketArrayOutput
}

type UsageExportBucketArray []UsageExportBucketInput

func (UsageExportBucketArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UsageExportBucket)(nil)).Elem()
}

func (i UsageExportBucketArray) ToUsageExportBucketArrayOutput() UsageExportBucketArrayOutput {
	return i.ToUsageExportBucketArrayOutputWithContext(context.Background())
}

func (i UsageExportBucketArray) ToUsageExportBucketArrayOutputWithContext(ctx context.Context) UsageExportBucketArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsageExportBucketArrayOutput)
}

// UsageExportBucketMapInput is an input type that accepts UsageExportBucketMap and UsageExportBucketMapOutput values.
// You can construct a concrete instance of `UsageExportBucketMapInput` via:
//
//          UsageExportBucketMap{ "key": UsageExportBucketArgs{...} }
type UsageExportBucketMapInput interface {
	pulumi.Input

	ToUsageExportBucketMapOutput() UsageExportBucketMapOutput
	ToUsageExportBucketMapOutputWithContext(context.Context) UsageExportBucketMapOutput
}

type UsageExportBucketMap map[string]UsageExportBucketInput

func (UsageExportBucketMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UsageExportBucket)(nil)).Elem()
}

func (i UsageExportBucketMap) ToUsageExportBucketMapOutput() UsageExportBucketMapOutput {
	return i.ToUsageExportBucketMapOutputWithContext(context.Background())
}

func (i UsageExportBucketMap) ToUsageExportBucketMapOutputWithContext(ctx context.Context) UsageExportBucketMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UsageExportBucketMapOutput)
}

type UsageExportBucketOutput struct{ *pulumi.OutputState }

func (UsageExportBucketOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UsageExportBucket)(nil))
}

func (o UsageExportBucketOutput) ToUsageExportBucketOutput() UsageExportBucketOutput {
	return o
}

func (o UsageExportBucketOutput) ToUsageExportBucketOutputWithContext(ctx context.Context) UsageExportBucketOutput {
	return o
}

func (o UsageExportBucketOutput) ToUsageExportBucketPtrOutput() UsageExportBucketPtrOutput {
	return o.ToUsageExportBucketPtrOutputWithContext(context.Background())
}

func (o UsageExportBucketOutput) ToUsageExportBucketPtrOutputWithContext(ctx context.Context) UsageExportBucketPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UsageExportBucket) *UsageExportBucket {
		return &v
	}).(UsageExportBucketPtrOutput)
}

type UsageExportBucketPtrOutput struct{ *pulumi.OutputState }

func (UsageExportBucketPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UsageExportBucket)(nil))
}

func (o UsageExportBucketPtrOutput) ToUsageExportBucketPtrOutput() UsageExportBucketPtrOutput {
	return o
}

func (o UsageExportBucketPtrOutput) ToUsageExportBucketPtrOutputWithContext(ctx context.Context) UsageExportBucketPtrOutput {
	return o
}

func (o UsageExportBucketPtrOutput) Elem() UsageExportBucketOutput {
	return o.ApplyT(func(v *UsageExportBucket) UsageExportBucket {
		if v != nil {
			return *v
		}
		var ret UsageExportBucket
		return ret
	}).(UsageExportBucketOutput)
}

type UsageExportBucketArrayOutput struct{ *pulumi.OutputState }

func (UsageExportBucketArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UsageExportBucket)(nil))
}

func (o UsageExportBucketArrayOutput) ToUsageExportBucketArrayOutput() UsageExportBucketArrayOutput {
	return o
}

func (o UsageExportBucketArrayOutput) ToUsageExportBucketArrayOutputWithContext(ctx context.Context) UsageExportBucketArrayOutput {
	return o
}

func (o UsageExportBucketArrayOutput) Index(i pulumi.IntInput) UsageExportBucketOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UsageExportBucket {
		return vs[0].([]UsageExportBucket)[vs[1].(int)]
	}).(UsageExportBucketOutput)
}

type UsageExportBucketMapOutput struct{ *pulumi.OutputState }

func (UsageExportBucketMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UsageExportBucket)(nil))
}

func (o UsageExportBucketMapOutput) ToUsageExportBucketMapOutput() UsageExportBucketMapOutput {
	return o
}

func (o UsageExportBucketMapOutput) ToUsageExportBucketMapOutputWithContext(ctx context.Context) UsageExportBucketMapOutput {
	return o
}

func (o UsageExportBucketMapOutput) MapIndex(k pulumi.StringInput) UsageExportBucketOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UsageExportBucket {
		return vs[0].(map[string]UsageExportBucket)[vs[1].(string)]
	}).(UsageExportBucketOutput)
}

func init() {
	pulumi.RegisterOutputType(UsageExportBucketOutput{})
	pulumi.RegisterOutputType(UsageExportBucketPtrOutput{})
	pulumi.RegisterOutputType(UsageExportBucketArrayOutput{})
	pulumi.RegisterOutputType(UsageExportBucketMapOutput{})
}