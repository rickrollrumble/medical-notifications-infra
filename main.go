package main

import (
	"example.com/medical-notifications-infra/imports/k8s"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

func NewChart(scope constructs.Construct, id string, ns string, appLabel string) cdk8s.Chart {
	chart := cdk8s.NewChart(scope, jsii.String(id), &cdk8s.ChartProps{
		Namespace: jsii.String(ns),
	})

	labels := map[string]*string{
		"app": jsii.String(appLabel),
	}

	k8s.NewKubeDeployment(chart, jsii.String("medical-notifications"), &k8s.KubeDeploymentProps{
		Spec: &k8s.DeploymentSpec{
			Replicas: jsii.Number(3),
			Selector: &k8s.LabelSelector{
				MatchLabels: &labels,
			},
			Template: &k8s.PodTemplateSpec{
				Metadata: &k8s.ObjectMeta{
					Labels: &labels,
				},
				Spec: &k8s.PodSpec{
					Containers: &[]*k8s.Container{
						{
							Name:  jsii.String("medical-notifications-backend"),
							Image: jsii.String("nginx:1.19.10"),
							Ports: &[]*k8s.ContainerPort{{
								ContainerPort: jsii.Number(443),
							}},
							// Resources: &k8s.ResourceRequirements{
							// 	Limits: &map[string]k8s.Quantity{
							// 		"cpu":    k8s.Quantity_FromString(jsii.String(("100m"))),
							// 		"memory": k8s.Quantity_FromString(jsii.String("100Mi")),
							// 	},
							// },
						},
						{
							Name:  jsii.String("medical-notifications-frontend"),
							Image: jsii.String("nginx:1.19.10"),
							Ports: &[]*k8s.ContainerPort{{
								ContainerPort: jsii.Number(443),
							}},
							// Resources: &k8s.ResourceRequirements{
							// 	Limits: &map[string]k8s.Quantity{
							// 		"cpu":    k8s.Quantity_FromString(jsii.String(("100m"))),
							// 		"memory": k8s.Quantity_FromString(jsii.String("100Mi")),
							// 	},
							// },
						},
					},
				},
			},
		},
	})

	return chart
}

func main() {
	app := cdk8s.NewApp(nil)

	// @param id indicates the name of the manifest
	// @ns indicats the namespace
	NewChart(app, "medical-notifications", "default", "my-app")

	app.Synth()
}
