//go:build e2e

package dynakube

import (
	dynakubev1beta1 "github.com/Dynatrace/dynatrace-operator/pkg/api/v1beta1/dynakube"
	"github.com/Dynatrace/dynatrace-operator/pkg/version"
	"github.com/Dynatrace/dynatrace-operator/test/helpers/components/operator"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Builder struct {
	dynakube dynakubev1beta1.DynaKube
}

func NewBuilder() Builder {
	return Builder{
		dynakube: dynakubev1beta1.DynaKube{
			ObjectMeta: metav1.ObjectMeta{
				Annotations: map[string]string{},
			},
			Spec: dynakubev1beta1.DynaKubeSpec{},
		},
	}
}

func (dynakubeBuilder Builder) Name(name string) Builder {
	dynakubeBuilder.dynakube.Name = name
	return dynakubeBuilder
}

func (dynakubeBuilder Builder) Namespace(namespace string) Builder {
	dynakubeBuilder.dynakube.Namespace = namespace
	return dynakubeBuilder
}

func (dynakubeBuilder Builder) WithDefaultObjectMeta() Builder {
	dynakubeBuilder.dynakube.ObjectMeta = metav1.ObjectMeta{
		Name:        defaultName,
		Namespace:   operator.DefaultNamespace,
		Annotations: map[string]string{},
	}

	return dynakubeBuilder
}

func (dynakubeBuilder Builder) WithCustomPullSecret(secretName string) Builder {
	dynakubeBuilder.dynakube.Spec.CustomPullSecret = secretName
	return dynakubeBuilder
}

func (dynakubeBuilder Builder) WithCustomCAs(configMapName string) Builder {
	dynakubeBuilder.dynakube.Spec.TrustedCAs = configMapName
	return dynakubeBuilder
}

func (dynakubeBuilder Builder) WithAnnotations(annotations map[string]string) Builder {
	for key, value := range annotations {
		dynakubeBuilder.dynakube.ObjectMeta.Annotations[key] = value
	}
	return dynakubeBuilder
}

func (dynakubeBuilder Builder) ApiUrl(apiUrl string) Builder {
	dynakubeBuilder.dynakube.Spec.APIURL = apiUrl
	return dynakubeBuilder
}

func (dynakubeBuilder Builder) WithActiveGate() Builder {
	dynakubeBuilder.dynakube.Spec.ActiveGate = dynakubev1beta1.ActiveGateSpec{
		Capabilities: []dynakubev1beta1.CapabilityDisplayName{
			dynakubev1beta1.KubeMonCapability.DisplayName,
			dynakubev1beta1.DynatraceApiCapability.DisplayName,
			dynakubev1beta1.RoutingCapability.DisplayName,
			dynakubev1beta1.MetricsIngestCapability.DisplayName,
		},
	}
	return dynakubeBuilder
}

func (dynakubeBuilder Builder) Tokens(secretName string) Builder {
	dynakubeBuilder.dynakube.Spec.Tokens = secretName
	return dynakubeBuilder
}

func (dynakubeBuilder Builder) NamespaceSelector(selector metav1.LabelSelector) Builder {
	dynakubeBuilder.dynakube.Spec.NamespaceSelector = selector
	return dynakubeBuilder
}

func (dynakubeBuilder Builder) WithDynakubeNamespaceSelector() Builder {
	return dynakubeBuilder.NamespaceSelector(metav1.LabelSelector{
		MatchLabels: map[string]string{
			"inject": dynakubeBuilder.dynakube.Name,
		},
	})
}

func (dynakubeBuilder Builder) Proxy(proxy *dynakubev1beta1.DynaKubeProxy) Builder {
	dynakubeBuilder.dynakube.Spec.Proxy = proxy
	return dynakubeBuilder
}

func (dynakubeBuilder Builder) WithIstio() Builder {
	dynakubeBuilder.dynakube.Spec.EnableIstio = true
	return dynakubeBuilder
}

func (dynakubeBuilder Builder) Privileged() Builder {
	dynakubeBuilder.dynakube.Annotations[dynakubev1beta1.AnnotationFeatureRunOneAgentContainerPrivileged] = "true"
	return dynakubeBuilder
}

func (dynakubeBuilder Builder) ClassicFullstack(classicFullStackSpec *dynakubev1beta1.HostInjectSpec) Builder {
	dynakubeBuilder.dynakube.Spec.OneAgent.ClassicFullStack = classicFullStackSpec
	return dynakubeBuilder
}

func (dynakubeBuilder Builder) CloudNative(cloudNativeFullStackSpec *dynakubev1beta1.CloudNativeFullStackSpec) Builder {
	dynakubeBuilder.dynakube.Spec.OneAgent.CloudNativeFullStack = cloudNativeFullStackSpec
	return dynakubeBuilder
}

func (dynakubeBuilder Builder) CloudNativeWithAgentVersion(cloudNativeFullStackSpec *dynakubev1beta1.CloudNativeFullStackSpec, version version.SemanticVersion) Builder {
	dynakubeBuilder.dynakube.Spec.OneAgent.CloudNativeFullStack = cloudNativeFullStackSpec
	dynakubeBuilder.dynakube.Spec.OneAgent.CloudNativeFullStack.Version = version.String()
	return dynakubeBuilder
}

func (dynakubeBuilder Builder) ApplicationMonitoring(applicationMonitoringSpec *dynakubev1beta1.ApplicationMonitoringSpec) Builder {
	dynakubeBuilder.dynakube.Spec.OneAgent.ApplicationMonitoring = applicationMonitoringSpec
	return dynakubeBuilder
}

func (builder Builder) WithSyntheticLocation(entityId string) Builder {
	builder.dynakube.Annotations[dynakubev1beta1.AnnotationFeatureSyntheticLocationEntityId] = entityId
	return builder
}

func (builder Builder) ResetOneAgent() Builder {
	builder.dynakube.Spec.OneAgent.ClassicFullStack = nil
	builder.dynakube.Spec.OneAgent.CloudNativeFullStack = nil
	builder.dynakube.Spec.OneAgent.ApplicationMonitoring = nil
	builder.dynakube.Spec.OneAgent.HostMonitoring = nil
	return builder
}

func (builder Builder) NetworkZone(networkZone string) Builder {
	builder.dynakube.Spec.NetworkZone = networkZone
	return builder
}

func (dynakubeBuilder Builder) Build() dynakubev1beta1.DynaKube {
	return dynakubeBuilder.dynakube
}
