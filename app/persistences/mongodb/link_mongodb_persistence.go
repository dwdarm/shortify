package persistences

import (
	"github.com/kamva/mgm/v3"
)

type LinkMongodbInstance struct {
	mgm.DefaultModel `bson:",inline"`
}

type LinkMongodbRepository struct {
}
