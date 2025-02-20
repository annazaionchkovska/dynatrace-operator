package dynakube

import (
	"context"
	"fmt"

	dynatracev1beta1 "github.com/Dynatrace/dynatrace-operator/pkg/api/v1beta1/dynakube"
)

const (
	errorInvalidSyntheticNodeType = `The DynaKube's specification requires illegally the synthetic node type: %v.
Make sure such a node is valid.
`
)

func invalidSyntheticNodeType(_ context.Context, dv *dynakubeValidator, dynakube *dynatracev1beta1.DynaKube) string {
	isTypeValid := func() bool {
		switch dynakube.FeatureSyntheticNodeType() {
		case dynatracev1beta1.SyntheticNodeXs,
			dynatracev1beta1.SyntheticNodeS,
			dynatracev1beta1.SyntheticNodeM:
			return true
		}
		return false
	}

	if dynakube.IsSyntheticMonitoringEnabled() && !isTypeValid() {
		log.Info(
			"requested dynakube has the invalid synthetic node type",
			"name", dynakube.Name,
			"namespace", dynakube.Namespace)
		return fmt.Sprintf(errorInvalidSyntheticNodeType, dynakube.FeatureSyntheticNodeType())
	}
	return ""
}
