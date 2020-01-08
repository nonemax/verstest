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
	/*metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PerconaServerMongoDBSpec   `json:"spec,omitempty"`
	Status PerconaServerMongoDBStatus `json:"status,omitempty"`

	StorageClassesAllocated PerconaServerMongoDBStorageClasses `json:"StorageClassesAllocated,omitempty"`
	StorageSizeAllocated    PerconaServerMongoDBStorageSizes   `json:"StorageSizeAllocated,omitempty"`*/

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   v1.PerconaServerMongoDBSpec   `json:"spec,omitempty"`
	Status v1.PerconaServerMongoDBStatus `json:"status,omitempty"`
}

/*
type PerconaServerMongoDBStorageClasses struct {
	BackupCoordinator string `json:"backup-coordinator,omitempty"`
	DataPod           string `json:"mongod,omitempty"`
}

type PerconaServerMongoDBStorageSizes struct {
	BackupCoordinator string `json:"backup-coordinator,omitempty"`
	DataPod           string `json:"mongod,omitempty"`
}

// PerconaServerMongoDBList contains a list of PerconaServerMongoDB
type PerconaServerMongoDBList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PerconaServerMongoDB `json:"items"`
}

type ClusterRole string

const (
	ClusterRoleShardSvr  ClusterRole = "shardsvr"
	ClusterRoleConfigSvr ClusterRole = "configsvr"
)

// PerconaServerMongoDBSpec defines the desired state of PerconaServerMongoDB
type PerconaServerMongoDBSpec struct {
	Pause            bool                          `json:"pause,omitempty"`
	Platform         *Platform                     `json:"platform,omitempty"`
	Image            string                        `json:"image,omitempty"`
	ImagePullSecrets []corev1.LocalObjectReference `json:"imagePullSecrets,omitempty"`
	RunUID           int64                         `json:"runUid,omitempty"`
	UnsafeConf       bool                          `json:"allowUnsafeConfigurations"`
	Mongod           *MongodSpec                   `json:"mongod,omitempty"`
	Replsets         []*ReplsetSpec                `json:"replsets,omitempty"`
	Secrets          *SecretsSpec                  `json:"secrets,omitempty"`
	Backup           BackupSpec                    `json:"backup,omitempty"`
	ImagePullPolicy  corev1.PullPolicy             `json:"imagePullPolicy,omitempty"`
	PMM              PMMSpec                       `json:"pmm,omitempty"`
}

type ReplsetMemberStatus struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

type ReplsetStatus struct {
	Members     []*ReplsetMemberStatus `json:"members,omitempty"`
	ClusterRole ClusterRole            `json:"clusterRole,omitempty"`

	Initialized bool     `json:"initialized,omitempty"`
	Size        int32    `json:"size"`
	Ready       int32    `json:"ready"`
	Status      AppState `json:"status,omitempty"`
	Message     string   `json:"message,omitempty"`
}

type AppState string

const (
	AppStatePending AppState = "pending"
	AppStateInit             = "initializing"
	AppStateReady            = "ready"
	AppStateError            = "error"
)

// PerconaServerMongoDBStatus defines the observed state of PerconaServerMongoDB
type PerconaServerMongoDBStatus struct {
	Status     AppState                  `json:"state,omitempty"`
	Message    string                    `json:"message,omitempty"`
	Conditions []ClusterCondition        `json:"conditions,omitempty"`
	Replsets   map[string]*ReplsetStatus `json:"replsets,omitempty"`
}

type ConditionStatus string

const (
	ConditionTrue    ConditionStatus = "True"
	ConditionFalse                   = "False"
	ConditionUnknown                 = "Unknown"
)

type ClusterConditionType string

const (
	ClusterReady   ClusterConditionType = "ClusterReady"
	ClusterInit                         = "ClusterInitializing"
	ClusterRSInit                       = "ReplsetInitialized"
	ClusterRSReady                      = "ReplsetReady"
	ClusterError                        = "Error"
)

type ClusterCondition struct {
	Status             ConditionStatus      `json:"status"`
	Type               ClusterConditionType `json:"type"`
	LastTransitionTime metav1.Time          `json:"lastTransitionTime,omitempty"`
	Reason             string               `json:"reason,omitempty"`
	Message            string               `json:"message,omitempty"`
}

type PMMSpec struct {
	Enabled    bool   `json:"enabled,omitempty"`
	ServerHost string `json:"serverHost,omitempty"`
	Image      string `json:"image,omitempty"`
}

type MultiAZ struct {
	Affinity            *PodAffinity             `json:"affinity,omitempty"`
	NodeSelector        map[string]string        `json:"nodeSelector,omitempty"`
	Tolerations         []corev1.Toleration      `json:"tolerations,omitempty"`
	PriorityClassName   string                   `json:"priorityClassName,omitempty"`
	Annotations         map[string]string        `json:"annotations,omitempty"`
	Labels              map[string]string        `json:"labels,omitempty"`
	PodDisruptionBudget *PodDisruptionBudgetSpec `json:"podDisruptionBudget,omitempty"`
}

type PodDisruptionBudgetSpec struct {
	MinAvailable   *intstr.IntOrString `json:"minAvailable,omitempty"`
	MaxUnavailable *intstr.IntOrString `json:"maxUnavailable,omitempty"`
}

type PodAffinity struct {
	TopologyKey *string          `json:"antiAffinityTopologyKey,omitempty"`
	Advanced    *corev1.Affinity `json:"advanced,omitempty"`
}

type ReplsetSpec struct {
	Resources                    *ResourcesSpec `json:"resources,omitempty"`
	Name                         string         `json:"name"`
	Size                         int32          `json:"size"`
	ClusterRole                  ClusterRole    `json:"clusterRole,omitempty"`
	Arbiter                      Arbiter        `json:"arbiter,omitempty"`
	Expose                       Expose         `json:"expose,omitempty"`
	VolumeSpec                   *VolumeSpec    `json:"volumeSpec,omitempty"`
	ReadinessInitialDelaySeconds *int32         `json:"readinessDelaySec,omitempty"`
	LivenessInitialDelaySeconds  *int32         `json:"livenessDelaySec,omitempty"`
	MultiAZ
}

type VolumeSpec struct {
	// EmptyDir represents a temporary directory that shares a pod's lifetime.
	EmptyDir *corev1.EmptyDirVolumeSource `json:"emptyDir,omitempty"`

	// HostPath represents a pre-existing file or directory on the host machine
	// that is directly exposed to the container.
	HostPath *corev1.HostPathVolumeSource `json:"hostPath,omitempty"`

	// PersistentVolumeClaim represents a reference to a PersistentVolumeClaim.
	// It has the highest level of precedence, followed by HostPath and
	// EmptyDir. And represents the PVC specification.
	PersistentVolumeClaim *corev1.PersistentVolumeClaimSpec `json:"persistentVolumeClaim,omitempty"`
}

type ResourceSpecRequirements struct {
	CPU    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
}

type ResourcesSpec struct {
	Limits   *ResourceSpecRequirements `json:"limits,omitempty"`
	Requests *ResourceSpecRequirements `json:"requests,omitempty"`
}

type SecretsSpec struct {
	Users       string `json:"users,omitempty"`
	SSL         string `json:"ssl,omitempty"`
	SSLInternal string `json:"sslInternal,omitempty"`
}

type MongosSpec struct {
	*ResourcesSpec `json:"resources,omitempty"`
	Port           int32 `json:"port,omitempty"`
	HostPort       int32 `json:"hostPort,omitempty"`
}

type MongodSpec struct {
	Net                *MongodSpecNet                `json:"net,omitempty"`
	AuditLog           *MongodSpecAuditLog           `json:"auditLog,omitempty"`
	OperationProfiling *MongodSpecOperationProfiling `json:"operationProfiling,omitempty"`
	Replication        *MongodSpecReplication        `json:"replication,omitempty"`
	Security           *MongodSpecSecurity           `json:"security,omitempty"`
	SetParameter       *MongodSpecSetParameter       `json:"setParameter,omitempty"`
	Storage            *MongodSpecStorage            `json:"storage,omitempty"`
}

type MongodSpecNet struct {
	Port     int32 `json:"port,omitempty"`
	HostPort int32 `json:"hostPort,omitempty"`
}

type MongodSpecReplication struct {
	OplogSizeMB int `json:"oplogSizeMB,omitempty"`
}

// MongodChiperMode is a cipher mode used by Data-at-Rest Encryption
type MongodChiperMode string

const (
	MongodChiperModeUnset MongodChiperMode = ""
	MongodChiperModeCBC                    = "AES256-CBC"
	MongodChiperModeGCM                    = "AES256-GCM"
)

type MongodSpecSecurity struct {
	RedactClientLogData  bool             `json:"redactClientLogData,omitempty"`
	EnableEncryption     *bool            `json:"enableEncryption,omitempty"`
	EncryptionKeySecret  string           `json:"encryptionKeySecret,omitempty"`
	EncryptionCipherMode MongodChiperMode `json:"encryptionCipherMode,omitempty"`
}

type MongodSpecSetParameter struct {
	TTLMonitorSleepSecs                   int `json:"ttlMonitorSleepSecs,omitempty"`
	WiredTigerConcurrentReadTransactions  int `json:"wiredTigerConcurrentReadTransactions,omitempty"`
	WiredTigerConcurrentWriteTransactions int `json:"wiredTigerConcurrentWriteTransactions,omitempty"`
}

type StorageEngine string

var (
	StorageEngineWiredTiger StorageEngine = "wiredTiger"
	StorageEngineInMemory   StorageEngine = "inMemory"
	StorageEngineMMAPv1     StorageEngine = "mmapv1"
)

type MongodSpecStorage struct {
	Engine         StorageEngine         `json:"engine,omitempty"`
	DirectoryPerDB bool                  `json:"directoryPerDB,omitempty"`
	SyncPeriodSecs int                   `json:"syncPeriodSecs,omitempty"`
	InMemory       *MongodSpecInMemory   `json:"inMemory,omitempty"`
	MMAPv1         *MongodSpecMMAPv1     `json:"mmapv1,omitempty"`
	WiredTiger     *MongodSpecWiredTiger `json:"wiredTiger,omitempty"`
}

type MongodSpecMMAPv1 struct {
	NsSize     int  `json:"nsSize,omitempty"`
	Smallfiles bool `json:"smallfiles,omitempty"`
}

type WiredTigerCompressor string

var (
	WiredTigerCompressorNone   WiredTigerCompressor = "none"
	WiredTigerCompressorSnappy WiredTigerCompressor = "snappy"
	WiredTigerCompressorZlib   WiredTigerCompressor = "zlib"
)

type MongodSpecWiredTigerEngineConfig struct {
	CacheSizeRatio      float64               `json:"cacheSizeRatio,omitempty"`
	DirectoryForIndexes bool                  `json:"directoryForIndexes,omitempty"`
	JournalCompressor   *WiredTigerCompressor `json:"journalCompressor,omitempty"`
}

type MongodSpecWiredTigerCollectionConfig struct {
	BlockCompressor *WiredTigerCompressor `json:"blockCompressor,omitempty"`
}

type MongodSpecWiredTigerIndexConfig struct {
	PrefixCompression bool `json:"prefixCompression,omitempty"`
}

type MongodSpecWiredTiger struct {
	CollectionConfig *MongodSpecWiredTigerCollectionConfig `json:"collectionConfig,omitempty"`
	EngineConfig     *MongodSpecWiredTigerEngineConfig     `json:"engineConfig,omitempty"`
	IndexConfig      *MongodSpecWiredTigerIndexConfig      `json:"indexConfig,omitempty"`
}

type MongodSpecInMemoryEngineConfig struct {
	InMemorySizeRatio float64 `json:"inMemorySizeRatio,omitempty"`
}

type MongodSpecInMemory struct {
	EngineConfig *MongodSpecInMemoryEngineConfig `json:"engineConfig,omitempty"`
}

type AuditLogDestination string

var AuditLogDestinationFile AuditLogDestination = "file"

type AuditLogFormat string

var (
	AuditLogFormatBSON AuditLogFormat = "BSON"
	AuditLogFormatJSON AuditLogFormat = "JSON"
)

type MongodSpecAuditLog struct {
	Destination AuditLogDestination `json:"destination,omitempty"`
	Format      AuditLogFormat      `json:"format,omitempty"`
	Filter      string              `json:"filter,omitempty"`
}

type OperationProfilingMode string

const (
	OperationProfilingModeAll    OperationProfilingMode = "all"
	OperationProfilingModeSlowOp OperationProfilingMode = "slowOp"
)

type MongodSpecOperationProfiling struct {
	Mode              OperationProfilingMode `json:"mode,omitempty"`
	SlowOpThresholdMs int                    `json:"slowOpThresholdMs,omitempty"`
	RateLimit         int                    `json:"rateLimit,omitempty"`
}

type BackupCoordinatorSpec struct {
	Resources                   *corev1.ResourceRequirements `json:"resources,omitempty"`
	StorageClass                string                       `json:"storageClass,omitempty"`
	EnableClientsLogging        bool                         `json:"enableClientsLogging,omitempty"`
	LivenessInitialDelaySeconds *int32                       `json:"livenessDelaySec,omitempty"`
	MultiAZ
}

type BackupDestinationType string

var (
	BackupDestinationS3   BackupDestinationType = "s3"
	BackupDestinationFile BackupDestinationType = "file"
)

type BackupCompressionType string

var (
	BackupCompressionGzip BackupCompressionType = "gzip"
)

type BackupTaskSpec struct {
	Name            string                `json:"name"`
	Enabled         bool                  `json:"enabled"`
	Schedule        string                `json:"schedule,omitempty"`
	StorageName     string                `json:"storageName,omitempty"`
	CompressionType BackupCompressionType `json:"compressionType,omitempty"`
}

type BackupSpec struct {
	Enabled          bool                             `json:"enabled"`
	Debug            bool                             `json:"debug"`
	RestartOnFailure *bool                            `json:"restartOnFailure,omitempty"`
	Coordinator      BackupCoordinatorSpec            `json:"coordinator,omitempty"`
	Storages         map[string]k8s.BackupStorageSpec `json:"storages,omitempty"`
	Image            string                           `json:"image,omitempty"`
	Tasks            []BackupTaskSpec                 `json:"tasks,omitempty"`
}

type Arbiter struct {
	Enabled bool  `json:"enabled"`
	Size    int32 `json:"size"`
	MultiAZ
}

type Expose struct {
	Enabled    bool               `json:"enabled"`
	ExposeType corev1.ServiceType `json:"exposeType,omitempty"`
}

type Platform string

const (
	PlatformUndef      Platform = ""
	PlatformKubernetes          = "kubernetes"
	PlatformOpenshift           = "openshift"
)
*/
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
