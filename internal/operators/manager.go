package operators

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/internal/operators/lso"
	"github.com/openshift/assisted-service/internal/operators/ocs"
	"github.com/openshift/assisted-service/models"
	"github.com/sirupsen/logrus"
)

var OperatorCVO models.MonitoredOperator = models.MonitoredOperator{
	Name:           "cvo",
	OperatorType:   models.OperatorTypeBuiltin,
	TimeoutSeconds: 60 * 60,
}

var OperatorConsole models.MonitoredOperator = models.MonitoredOperator{
	Name:           "console",
	OperatorType:   models.OperatorTypeBuiltin,
	TimeoutSeconds: 60 * 60,
}

var monitoredOperators = [...]*models.MonitoredOperator{
	&OperatorCVO,
	&OperatorConsole,
	&lso.Operator,
	&ocs.Operator,
}

//go:generate mockgen -package=operators -destination=mock_operators_api.go . API
type API interface {
	// ValidateOCSRequirements validates OCS requirements
	ValidateOCSRequirements(cluster *common.Cluster) string
	// GenerateManifests generates manifests for all enabled operators.
	// Returns map assigning manifest content to its desired file name
	GenerateManifests(cluster *common.Cluster) (map[string]string, error)
	// AnyOperatorEnabled checks whether any operator has been enabled for the given cluster
	AnyOperatorEnabled(cluster *common.Cluster) bool
	// GetOperatorStatusInfo gets status info of an operator of given name.
	GetOperatorStatusInfo(cluster *common.Cluster, operatorName string) string
	// GetMonitoredOperatorsList returns the monitored operators available by the manager.
	GetMonitoredOperatorsList() []*models.MonitoredOperator
	// GetOperatorByName the manager's supported operator object by name.
	GetOperatorByName(operatorName string) (*models.MonitoredOperator, error)
	// GetOperatorsByType returns the manager's supported operator objects by type.
	GetOperatorsByType(operatorType models.OperatorType) []*models.MonitoredOperator
}

// Manager is responsible for performing operations against additional operators
type Manager struct {
	API
	log                logrus.FieldLogger
	ocsValidatorConfig *ocs.Config
	ocsValidator       ocs.OcsValidator
	monitoredOperators []*models.MonitoredOperator
}

// NewManager creates new instance of an Operator Manager
func NewManager(log logrus.FieldLogger) *Manager {
	cfg := ocs.Config{}
	err := envconfig.Process("myapp", &cfg)
	if err != nil {
		log.Fatal(err.Error())
	}
	return NewManagerWithConfig(log, &cfg)
}

// NewManagerWithConfig creates new instance of an Operator Manager
func NewManagerWithConfig(log logrus.FieldLogger, cfg *ocs.Config) *Manager {
	ocsValidator := ocs.NewOCSValidator(log.WithField("pkg", "ocs-operator-state"), cfg)
	return &Manager{
		log:                log,
		ocsValidatorConfig: cfg,
		ocsValidator:       ocsValidator,
		monitoredOperators: monitoredOperators[:],
	}
}

// GenerateManifests generates manifests for all enabled operators.
// Returns map assigning manifest content to its desired file name
func (mgr *Manager) GenerateManifests(cluster *common.Cluster) (map[string]string, error) {
	lsoEnabled := false
	ocsEnabled := mgr.checkOCSEnabled(cluster)
	if ocsEnabled {
		lsoEnabled = true // if OCS is enabled, LSO must be enabled by default
	} else {
		lsoEnabled = mgr.checkLSOEnabled(cluster)
	}
	operatorManifests := make(map[string]string)

	if lsoEnabled {
		manifests, err := mgr.generateLSOManifests(cluster)
		if err != nil {
			mgr.log.Error("Cannot generate LSO manifests due to ", err)
			return nil, err
		}
		for k, v := range manifests {
			operatorManifests[k] = v
		}
	}

	if ocsEnabled {
		manifests, err := mgr.generateOCSManifests(cluster)
		if err != nil {
			mgr.log.Error("Cannot generate OCS manifests due to ", err)
			return nil, err
		}
		for k, v := range manifests {
			operatorManifests[k] = v
		}
	}
	return operatorManifests, nil
}

// AnyOperatorEnabled checks whether any operator has been enabled for the given cluster
func (mgr *Manager) AnyOperatorEnabled(cluster *common.Cluster) bool {
	return mgr.checkLSOEnabled(cluster) || mgr.checkOCSEnabled(cluster)
}

// ValidateOCSRequirements validates OCS requirements. Returns "true" if OCS operator is not deployed
func (mgr *Manager) ValidateOCSRequirements(cluster *common.Cluster) string {
	if IsEnabled(cluster, ocs.Operator.Name) {
		return mgr.ocsValidator.ValidateOCSRequirements(&cluster.Cluster)
	}
	return "success"
}

// GetOperatorStatusInfo gets status of an operator of given type.
func (mgr *Manager) GetOperatorStatusInfo(cluster *common.Cluster, operatorName string) string {
	operator := findOperator(cluster, operatorName)
	if operator != nil {
		return operator.StatusInfo
	}
	return "OCS is disabled"
}

func (mgr *Manager) generateLSOManifests(cluster *common.Cluster) (map[string]string, error) {
	mgr.log.Info("Creating LSO Manifests")
	return lso.Manifests(cluster.OpenshiftVersion)
}

func (mgr *Manager) generateOCSManifests(cluster *common.Cluster) (map[string]string, error) {
	mgr.log.Info("Creating OCS Manifests")
	return ocs.Manifests(mgr.ocsValidatorConfig.OCSMinimalDeployment, cluster.OpenshiftVersion, mgr.ocsValidatorConfig.OCSDisksAvailable, len(cluster.Cluster.Hosts))
}

func (mgr *Manager) checkLSOEnabled(cluster *common.Cluster) bool {
	return IsEnabled(cluster, lso.Operator.Name)
}

func (mgr *Manager) checkOCSEnabled(cluster *common.Cluster) bool {
	return IsEnabled(cluster, ocs.Operator.Name)
}

func findOperator(cluster *common.Cluster, operatorName string) *models.MonitoredOperator {
	for _, operator := range cluster.MonitoredOperators {
		if operator.Name == operatorName {
			return operator
		}
	}
	return nil
}

func IsEnabled(cluster *common.Cluster, operatorName string) bool {
	return findOperator(cluster, operatorName) != nil
}

func (mgr *Manager) GetMonitoredOperatorsList() []*models.MonitoredOperator {
	return mgr.monitoredOperators[:]
}

func (mgr *Manager) GetOperatorByName(operatorName string) (*models.MonitoredOperator, error) {
	for _, operator := range mgr.monitoredOperators {
		if operator.Name == operatorName {
			return operator, nil
		}
	}

	return nil, fmt.Errorf("Operator %s isn't supported", operatorName)
}

func (mgr *Manager) GetOperatorsByType(operatorType models.OperatorType) []*models.MonitoredOperator {
	operators := make([]*models.MonitoredOperator, 0)

	for _, operator := range mgr.monitoredOperators {
		if operator.OperatorType == operatorType {
			operators = append(operators, operator)
		}
	}

	return operators
}
