// TODO: Scaffold for API types that will live in k8s.io/api/core/v1.
// Delete this package and swap references to the corev1 equivalents.

package volumehealth

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type VolumeHealthStatusType string

const (
	VolumeHealthInaccessible VolumeHealthStatusType = "Inaccessible"
	VolumeHealthDataLoss     VolumeHealthStatusType = "DataLoss"
	VolumeHealthDegraded     VolumeHealthStatusType = "Degraded"
)

type VolumeHealthCondition struct {
	Status  VolumeHealthStatusType `json:"status"`
	Reason  string                 `json:"reason"`
	Message string                 `json:"message,omitempty"`
}

type VolumeHealthStatus struct {
	Conditions         []VolumeHealthCondition `json:"conditions,omitempty"`
	LastTransitionTime *metav1.Time            `json:"lastTransitionTime,omitempty"`
}
