package csi_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/kubernetes-csi/external-health-monitor/pkg/apis/volumehealth"
)

const fieldManager = "csi-external-health-monitor-controller"

// Called only when the condition set changed, so the transition time is always fresh.
func buildHealthStatus(desired []volumehealth.VolumeHealthCondition) *volumehealth.VolumeHealthStatus {
	now := metav1.Now()
	return &volumehealth.VolumeHealthStatus{
		Conditions:         desired,
		LastTransitionTime: &now,
	}
}

// TODO replace with clientset.CoreV1().PersistentVolumeClaims(ns).ApplyStatus
func (checker *PVHealthConditionChecker) patchPVCHealthStatus(
	ctx context.Context,
	pvc *v1.PersistentVolumeClaim,
	status *volumehealth.VolumeHealthStatus,
) error {
	// A nil/empty status serializes to null, clearing the field server-side.
	var healthValue interface{}
	if status != nil && len(status.Conditions) > 0 {
		healthValue = status
	}

	patch := map[string]interface{}{
		"status": map[string]interface{}{
			"healthStatus": healthValue,
		},
	}
	data, err := json.Marshal(patch)
	if err != nil {
		return fmt.Errorf("failed to marshal health status patch for PVC %s/%s: %w", pvc.Namespace, pvc.Name, err)
	}

	_, err = checker.k8sClient.CoreV1().PersistentVolumeClaims(pvc.Namespace).Patch(
		ctx,
		pvc.Name,
		types.MergePatchType,
		data,
		metav1.PatchOptions{FieldManager: fieldManager},
		"status",
	)
	if err != nil {
		return fmt.Errorf("failed to patch health status for PVC %s/%s: %w", pvc.Namespace, pvc.Name, err)
	}
	return nil
}

// normalizeConditions returns a stable-sorted copy so equality comparison and patch output are deterministic.
func normalizeConditions(in []volumehealth.VolumeHealthCondition) []volumehealth.VolumeHealthCondition {
	if len(in) == 0 {
		return nil
	}
	out := make([]volumehealth.VolumeHealthCondition, len(in))
	copy(out, in)
	sort.Slice(out, func(i, j int) bool {
		if out[i].Status != out[j].Status {
			return out[i].Status < out[j].Status
		}
		return out[i].Reason < out[j].Reason
	})
	return out
}

// Compares status, reason, and message. A message-only change still triggers a patch.
// Inputs must be normalized first.
func conditionsEqual(a, b []volumehealth.VolumeHealthCondition) bool {
	return reflect.DeepEqual(a, b)
}
