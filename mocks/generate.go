package mocks

//go:generate mockgen -package mocks -destination mock_assisted_service.go github.com/openshift/assisted-service/restapi InstallerAPI
//go:generate mockgen -package mocks -destination mock_manifests.go github.com/openshift/assisted-service/restapi ManifestsAPI
