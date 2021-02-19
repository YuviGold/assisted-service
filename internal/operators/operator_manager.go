package operators

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/openshift/assisted-service/internal/operators/lso"
	"github.com/openshift/assisted-service/internal/operators/ocs"
	"github.com/openshift/assisted-service/models"
)

var OperatorCVO models.Operator = models.Operator{
	Name:           swag.String("cvo"),
	OperatorType:   swag.String(models.OperatorOperatorTypeBuiltin),
	TimeoutSeconds: 60 * 60,
}

var OperatorConsole models.Operator = models.Operator{
	Name:           swag.String("console"),
	OperatorType:   swag.String(models.OperatorOperatorTypeBuiltin),
	TimeoutSeconds: 60 * 60,
}

var monitoredOperators = [...]*models.Operator{
	&OperatorCVO,
	&OperatorConsole,
	&lso.Operator,
	&ocs.Operator,
}

type OperatorManagerAPI interface {
	GetMonitoredOperatorsList() []models.Operator
	GetOperatorByName(operatorName string) (*models.Operator, error)
	GetOperatorsByType(operatorType string) []*models.Operator
}

type operatorManager struct {
	OperatorManagerAPI
	monitoredOperators []*models.Operator
}

func NewOperatorManager() *operatorManager {
	return &operatorManager{
		monitoredOperators: monitoredOperators[:],
	}
}

func (manager *operatorManager) GetMonitoredOperatorsList() []*models.Operator {
	return manager.monitoredOperators[:]
}

func (manager *operatorManager) GetOperatorByName(operatorName string) (*models.Operator, error) {
	for _, operator := range manager.monitoredOperators {
		if swag.StringValue(operator.Name) == operatorName {
			return operator, nil
		}
	}

	return nil, fmt.Errorf("Operator %s isn't supported", operatorName)
}

func (manager *operatorManager) GetOperatorsByType(operatorType string) []*models.Operator {
	operators := make([]*models.Operator, 0)

	for _, operator := range manager.monitoredOperators {
		if swag.StringValue(operator.OperatorType) == operatorType {
			operators = append(operators, operator)
		}
	}

	return operators
}
