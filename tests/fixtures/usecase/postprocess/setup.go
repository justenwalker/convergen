//go:build convergen

package postprocess

import (
	"github.com/reedom/convergen/tests/fixtures/data/domain"
	"github.com/reedom/convergen/tests/fixtures/data/model"
	_ "github.com/reedom/convergen/tests/fixtures/usecase/postprocess/local"
)

//go:generate go run github.com/reedom/convergen
type Convergen interface {
	// :postprocess PostDomainToModel
	DomainToModel(*domain.Pet) *model.Pet
	// :postprocess local.PostModelToDomain
	ModelToDomain(*model.Pet) (*domain.Pet, error)
}

func PostDomainToModel(lhs *model.Pet, rhs domain.Pet) {

}