// Copyright 2018 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validation

import (
	"testing"

	torchv1 "github.com/kubeflow/pytorch-operator/pkg/apis/pytorch/v1alpha1"
	torchv2 "github.com/kubeflow/pytorch-operator/pkg/apis/pytorch/v1alpha2"
	torchv1beta1 "github.com/kubeflow/pytorch-operator/pkg/apis/pytorch/v1beta1"
	common "github.com/kubeflow/tf-operator/pkg/apis/common/v1beta1"

	"github.com/gogo/protobuf/proto"
	"k8s.io/api/core/v1"
)

func TestValidateBetaOnePyTorchJobSpec(t *testing.T) {
	testCases := []torchv1beta1.PyTorchJobSpec{
		{
			PyTorchReplicaSpecs: nil,
		},
		{
			PyTorchReplicaSpecs: map[torchv1beta1.PyTorchReplicaType]*common.ReplicaSpec{
				torchv1beta1.PyTorchReplicaTypeWorker: &common.ReplicaSpec{
					Template: v1.PodTemplateSpec{
						Spec: v1.PodSpec{
							Containers: []v1.Container{},
						},
					},
				},
			},
		},
		{
			PyTorchReplicaSpecs: map[torchv1beta1.PyTorchReplicaType]*common.ReplicaSpec{
				torchv1beta1.PyTorchReplicaTypeWorker: &common.ReplicaSpec{
					Template: v1.PodTemplateSpec{
						Spec: v1.PodSpec{
							Containers: []v1.Container{
								v1.Container{
									Image: "",
								},
							},
						},
					},
				},
			},
		},
		{
			PyTorchReplicaSpecs: map[torchv1beta1.PyTorchReplicaType]*common.ReplicaSpec{
				torchv1beta1.PyTorchReplicaTypeWorker: &common.ReplicaSpec{
					Template: v1.PodTemplateSpec{
						Spec: v1.PodSpec{
							Containers: []v1.Container{
								v1.Container{
									Name:  "",
									Image: "gcr.io/kubeflow-ci/pytorch-dist-mnist_test:1.0",
								},
							},
						},
					},
				},
			},
		},
		{
			PyTorchReplicaSpecs: map[torchv1beta1.PyTorchReplicaType]*common.ReplicaSpec{
				torchv1beta1.PyTorchReplicaTypeMaster: &common.ReplicaSpec{
					Replicas: torchv1beta1.Int32(2),
					Template: v1.PodTemplateSpec{
						Spec: v1.PodSpec{
							Containers: []v1.Container{
								v1.Container{
									Name:  "pytorch",
									Image: "gcr.io/kubeflow-ci/pytorch-dist-mnist_test:1.0",
								},
							},
						},
					},
				},
			},
		},
		{
			PyTorchReplicaSpecs: map[torchv1beta1.PyTorchReplicaType]*common.ReplicaSpec{
				torchv1beta1.PyTorchReplicaTypeWorker: &common.ReplicaSpec{
					Replicas: torchv1beta1.Int32(1),
					Template: v1.PodTemplateSpec{
						Spec: v1.PodSpec{
							Containers: []v1.Container{
								v1.Container{
									Name:  "pytorch",
									Image: "gcr.io/kubeflow-ci/pytorch-dist-mnist_test:1.0",
								},
							},
						},
					},
				},
			},
		},
	}
	for _, c := range testCases {
		err := ValidateBetaOnePyTorchJobSpec(&c)
		if err.Error() != "PyTorchJobSpec is not valid" {
			t.Error("Failed validate the v1beta1.PyTorchJobSpec")
		}
	}
}

func TestValidateAlphaTwoPyTorchJobSpec(t *testing.T) {
	testCases := []torchv2.PyTorchJobSpec{
		{
			PyTorchReplicaSpecs: nil,
		},
		{
			PyTorchReplicaSpecs: map[torchv2.PyTorchReplicaType]*torchv2.PyTorchReplicaSpec{
				torchv2.PyTorchReplicaTypeWorker: &torchv2.PyTorchReplicaSpec{
					Template: v1.PodTemplateSpec{
						Spec: v1.PodSpec{
							Containers: []v1.Container{},
						},
					},
				},
			},
		},
		{
			PyTorchReplicaSpecs: map[torchv2.PyTorchReplicaType]*torchv2.PyTorchReplicaSpec{
				torchv2.PyTorchReplicaTypeWorker: &torchv2.PyTorchReplicaSpec{
					Template: v1.PodTemplateSpec{
						Spec: v1.PodSpec{
							Containers: []v1.Container{
								v1.Container{
									Image: "",
								},
							},
						},
					},
				},
			},
		},
		{
			PyTorchReplicaSpecs: map[torchv2.PyTorchReplicaType]*torchv2.PyTorchReplicaSpec{
				torchv2.PyTorchReplicaTypeWorker: &torchv2.PyTorchReplicaSpec{
					Template: v1.PodTemplateSpec{
						Spec: v1.PodSpec{
							Containers: []v1.Container{
								v1.Container{
									Name:  "",
									Image: "gcr.io/kubeflow-ci/pytorch-dist-mnist_test:1.0",
								},
							},
						},
					},
				},
			},
		},
		{
			PyTorchReplicaSpecs: map[torchv2.PyTorchReplicaType]*torchv2.PyTorchReplicaSpec{
				torchv2.PyTorchReplicaTypeMaster: &torchv2.PyTorchReplicaSpec{
					Replicas: torchv2.Int32(2),
					Template: v1.PodTemplateSpec{
						Spec: v1.PodSpec{
							Containers: []v1.Container{
								v1.Container{
									Name:  "pytorch",
									Image: "gcr.io/kubeflow-ci/pytorch-dist-mnist_test:1.0",
								},
							},
						},
					},
				},
			},
		},
		{
			PyTorchReplicaSpecs: map[torchv2.PyTorchReplicaType]*torchv2.PyTorchReplicaSpec{
				torchv2.PyTorchReplicaTypeWorker: &torchv2.PyTorchReplicaSpec{
					Replicas: torchv2.Int32(1),
					Template: v1.PodTemplateSpec{
						Spec: v1.PodSpec{
							Containers: []v1.Container{
								v1.Container{
									Name:  "pytorch",
									Image: "gcr.io/kubeflow-ci/pytorch-dist-mnist_test:1.0",
								},
							},
						},
					},
				},
			},
		},
	}
	for _, c := range testCases {
		err := ValidateAlphaTwoPyTorchJobSpec(&c)
		if err.Error() != "PyTorchJobSpec is not valid" {
			t.Error("Failed validate the alpha2.PyTorchJobSpec")
		}
	}
}

func TestValidate(t *testing.T) {
	type testCase struct {
		in             *torchv1.PyTorchJobSpec
		expectingError bool
	}

	testCases := []testCase{
		{
			in: &torchv1.PyTorchJobSpec{
				ReplicaSpecs: []*torchv1.PyTorchReplicaSpec{
					{
						Template: &v1.PodTemplateSpec{
							Spec: v1.PodSpec{
								Containers: []v1.Container{
									{
										Name: "pytorch",
									},
								},
							},
						},
						PyTorchReplicaType: torchv1.MASTER,
						Replicas:           proto.Int32(1),
					},
				},
				PyTorchImage: "pytorch/pytorch:v0.2",
			},
			expectingError: false,
		},
		{
			in: &torchv1.PyTorchJobSpec{
				ReplicaSpecs: []*torchv1.PyTorchReplicaSpec{
					{
						Template: &v1.PodTemplateSpec{
							Spec: v1.PodSpec{
								Containers: []v1.Container{
									{
										Name: "pytorch",
									},
								},
							},
						},
						PyTorchReplicaType: torchv1.WORKER,
						Replicas:           proto.Int32(1),
					},
				},
				PyTorchImage: "pytorch/pytorch:v0.2",
			},
			expectingError: true,
		},
		{
			in: &torchv1.PyTorchJobSpec{
				ReplicaSpecs: []*torchv1.PyTorchReplicaSpec{
					{
						Template: &v1.PodTemplateSpec{
							Spec: v1.PodSpec{
								Containers: []v1.Container{
									{
										Name: "pytorch",
									},
								},
							},
						},
						PyTorchReplicaType: torchv1.WORKER,
						Replicas:           proto.Int32(1),
					},
				},
				PyTorchImage: "pytorch/pytorch:v0.2",
				TerminationPolicy: &torchv1.TerminationPolicySpec{
					Master: &torchv1.MasterSpec{
						ReplicaName: "WORKER",
						ReplicaRank: 0,
					},
				},
			},
			expectingError: false,
		},
	}

	for _, c := range testCases {
		job := &torchv1.PyTorchJob{
			Spec: *c.in,
		}
		torchv1.SetObjectDefaults_PyTorchJob(job)
		if err := ValidatePyTorchJobSpec(&job.Spec); (err != nil) != c.expectingError {
			t.Errorf("unexpected validation result: %v", err)
		}
	}
}
