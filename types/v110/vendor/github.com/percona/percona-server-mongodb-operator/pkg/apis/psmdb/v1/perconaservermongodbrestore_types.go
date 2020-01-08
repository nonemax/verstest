package v1

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PerconaServerMongoDBRestoreSpec defines the desired state of PerconaServerMongoDBRestore
type PerconaServerMongoDBRestoreSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	BackupName  string `json:"backupName,omitempty"`
	ClusterName string `json:"clusterName,omitempty"`
	Destination string `json:"destination,omitempty"`
	StorageName string `json:"storageName,omitempty"`
}

// PerconaSMDBRestoreStatusState is for restore status states
type PerconaSMDBRestoreStatusState string

const (
	RestoreStateRequested PerconaSMDBRestoreStatusState = "requested"
	RestoreStateReady     PerconaSMDBRestoreStatusState = "ready"
	RestoreStateRejected  PerconaSMDBRestoreStatusState = "rejected"
)

// PerconaServerMongoDBRestoreStatus defines the observed state of PerconaServerMongoDBRestore
type PerconaServerMongoDBRestoreStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	State PerconaSMDBRestoreStatusState `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PerconaServerMongoDBRestore is the Schema for the perconaservermongodbrestores API
// +k8s:openapi-gen=true
type PerconaServerMongoDBRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PerconaServerMongoDBRestoreSpec   `json:"spec,omitempty"`
	Status PerconaServerMongoDBRestoreStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PerconaServerMongoDBRestoreList contains a list of PerconaServerMongoDBRestore
type PerconaServerMongoDBRestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PerconaServerMongoDBRestore `json:"items"`
}

func (r *PerconaServerMongoDBRestore) CheckFields() error {
	if len(r.Spec.ClusterName) == 0 {
		return fmt.Errorf("spec clusterName field is empty")
	}
	if len(r.Spec.BackupName) == 0 && (len(r.Spec.StorageName) == 0 || len(r.Spec.Destination) == 0) {
		return fmt.Errorf("fields backupName or StorageNmae and destination is empty")
	}

	return nil
}
