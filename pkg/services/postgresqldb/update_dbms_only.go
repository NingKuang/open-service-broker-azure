package postgresqldb

import (
	"github.com/Azure/open-service-broker-azure/pkg/service"
)

func (d *dbmsOnlyManager) ValidateUpdatingParameters(
	updatingParameters service.UpdatingParameters,
) error {
	return nil
}

func (d *dbmsOnlyManager) GetUpdater(service.Plan) (service.Updater, error) {
	return service.NewUpdater()
}
