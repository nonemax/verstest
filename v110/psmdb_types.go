package v110

import (
	"encoding/json"
	"log"

	"github.com/pkg/errors"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1 "github.com/percona/percona-server-mongodb-operator/pkg/apis/psmdb/v1"
)

// PerconaServerMongoDB is the Schema for the perconaservermongodbs API
type PerconaServerMongoDB struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   v1.PerconaServerMongoDBSpec   `json:"spec,omitempty"`
	Status v1.PerconaServerMongoDBStatus `json:"status,omitempty"`
}

func (cr *PerconaServerMongoDB) GetSpec() interface{} {
	rs := v1.ReplsetSpec{}
	cr.Spec.Replsets = []*v1.ReplsetSpec{&rs}
	return cr.Spec
}

func (cr *PerconaServerMongoDB) GetName() string {
	return cr.ObjectMeta.Name
}

func (cr *PerconaServerMongoDB) SetName(name string) {
	cr.ObjectMeta.Name = name
}

func (cr *PerconaServerMongoDB) SetUsersSecretName(name string) {
	cr.Spec.Secrets = &v1.SecretsSpec{
		Users: name + "-psmdb-users-secrets",
	}
}

func (cr *PerconaServerMongoDB) GetOperatorImage() string {
	return "percona/percona-server-mongodb-operator:1.1.0"
}

func (cr *PerconaServerMongoDB) SetLabels(labels map[string]string) {
	cr.ObjectMeta.Labels = labels
}

func (cr *PerconaServerMongoDB) MarshalRequests() error {
	_, err := cr.Spec.Replsets[0].VolumeSpec.PersistentVolumeClaim.Resources.Requests[corev1.ResourceStorage].MarshalJSON()
	return err
}

func (cr *PerconaServerMongoDB) GetCR() (string, error) {
	b, err := json.Marshal(cr)
	if err != nil {
		return "", errors.Wrap(err, "marshal cr template")
	}

	return string(b), nil
}

// Upgrade upgrades culster with given images
func (cr *PerconaServerMongoDB) Upgrade(imgs map[string]string) {
	if img, ok := imgs["psmdb"]; ok {
		cr.Spec.Image = img
	}
	if img, ok := imgs["backup"]; ok {
		cr.Spec.Backup.Image = img
	}
}

func (cr *PerconaServerMongoDB) SetDefaults() error {
	cr.Spec.SchedulerName = "test"
	log.Println(cr.Spec.SchedulerName)
	rsName := "rs0"
	rs := &v1.ReplsetSpec{
		Name: rsName,
	}

	volSizeFlag := "6G"
	volSize, err := resource.ParseQuantity(volSizeFlag)
	if err != nil {
		return errors.Wrap(err, "storage-size")
	}
	rs.VolumeSpec = &v1.VolumeSpec{
		PersistentVolumeClaim: &corev1.PersistentVolumeClaimSpec{
			Resources: corev1.ResourceRequirements{
				Requests: corev1.ResourceList{corev1.ResourceStorage: volSize},
			},
		},
	}
	rs.Size = int32(3)
	rs.Resources = &v1.ResourcesSpec{
		Requests: &v1.ResourceSpecRequirements{
			CPU:    "600m",
			Memory: "1G",
		},
	}
	psmdbtpk := "none" //"kubernetes.io/hostname"
	rs.Affinity = &v1.PodAffinity{
		TopologyKey: &psmdbtpk,
	}
	cr.Spec.Replsets = []*v1.ReplsetSpec{
		rs,
	}
	cr.TypeMeta.APIVersion = "psmdb.percona.com/v1"
	cr.TypeMeta.Kind = "PerconaServerMongoDB"

	cr.Spec.Image = "percona/percona-server-mongodb-operator:1.1.0-mongod4.0"

	return nil
}

/*
module github.com/Percona-Lab/percona-dbaas-cli

go 1.13

require (
	github.com/Percona-Lab/percona-dbaas-cli/dbaas-lib/engines/k8s-psmdb/types/v110 v110.0.0-00010101000000-000000000000
	github.com/Percona-Lab/percona-dbaas-cli/dbaas-lib/engines/k8s-psmdb/types/v120 v120.0.0-00010101000000-000000000000
	github.com/hashicorp/go-version v1.2.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.5
	k8s.io/api v0.17.0
	k8s.io/apimachinery v0.17.0
	k8s.io/client-go v11.0.0+incompatible
)

replace github.com/Percona-Lab/percona-dbaas-cli/dbaas-lib/engines/k8s-psmdb/types/v110 => ./dbaas-lib/engines/k8s-psmdb/types/v110

replace github.com/Percona-Lab/percona-dbaas-cli/dbaas-lib/engines/k8s-psmdb/types/v120 => ./dbaas-lib/engines/k8s-psmdb/types/v120

*/

/*
module github.com/Percona-Lab/percona-dbaas-cli/dbaas-lib/engines/k8s-psmdb/types/v110

go 1.13

require (
	github.com/percona/percona-server-mongodb-operator v0.0.0-20190707075059-f6a9dada369e
	k8s.io/api v0.17.0 // indirect
	sigs.k8s.io/controller-runtime v0.4.0 // indirect
    k8s.io/apimachinery v0.17.0
)
*/
