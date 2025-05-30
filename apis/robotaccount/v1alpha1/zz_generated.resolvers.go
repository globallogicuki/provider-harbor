/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	v1alpha1 "github.com/globallogicuki/provider-harbor/apis/project/v1alpha1"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this RobotAccount.
func (mg *RobotAccount) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Permissions); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Permissions[i3].Namespace),
			Extract:      resource.ExtractParamPath("name", true),
			Reference:    mg.Spec.ForProvider.Permissions[i3].NamespaceRef,
			Selector:     mg.Spec.ForProvider.Permissions[i3].NamespaceSelector,
			To: reference.To{
				List:    &v1alpha1.ProjectList{},
				Managed: &v1alpha1.Project{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Permissions[i3].Namespace")
		}
		mg.Spec.ForProvider.Permissions[i3].Namespace = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Permissions[i3].NamespaceRef = rsp.ResolvedReference

	}

	return nil
}
