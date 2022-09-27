//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v2alpha1

import (
	common "k8s.io/kube-openapi/pkg/common"
	spec "k8s.io/kube-openapi/pkg/validation/spec"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./apis/datadoghq/v2alpha1.CustomConfig":                      schema__apis_datadoghq_v2alpha1_CustomConfig(ref),
		"./apis/datadoghq/v2alpha1.DatadogAgent":                      schema__apis_datadoghq_v2alpha1_DatadogAgent(ref),
		"./apis/datadoghq/v2alpha1.DatadogAgentGenericContainer":      schema__apis_datadoghq_v2alpha1_DatadogAgentGenericContainer(ref),
		"./apis/datadoghq/v2alpha1.DatadogAgentStatus":                schema__apis_datadoghq_v2alpha1_DatadogAgentStatus(ref),
		"./apis/datadoghq/v2alpha1.DatadogCredentials":                schema__apis_datadoghq_v2alpha1_DatadogCredentials(ref),
		"./apis/datadoghq/v2alpha1.DatadogFeatures":                   schema__apis_datadoghq_v2alpha1_DatadogFeatures(ref),
		"./apis/datadoghq/v2alpha1.DogstatsdFeatureConfig":            schema__apis_datadoghq_v2alpha1_DogstatsdFeatureConfig(ref),
		"./apis/datadoghq/v2alpha1.EventCollectionFeatureConfig":      schema__apis_datadoghq_v2alpha1_EventCollectionFeatureConfig(ref),
		"./apis/datadoghq/v2alpha1.KubeStateMetricsCoreFeatureConfig": schema__apis_datadoghq_v2alpha1_KubeStateMetricsCoreFeatureConfig(ref),
		"./apis/datadoghq/v2alpha1.LocalService":                      schema__apis_datadoghq_v2alpha1_LocalService(ref),
		"./apis/datadoghq/v2alpha1.MultiCustomConfig":                 schema__apis_datadoghq_v2alpha1_MultiCustomConfig(ref),
		"./apis/datadoghq/v2alpha1.NetworkPolicyConfig":               schema__apis_datadoghq_v2alpha1_NetworkPolicyConfig(ref),
		"./apis/datadoghq/v2alpha1.OTLPFeatureConfig":                 schema__apis_datadoghq_v2alpha1_OTLPFeatureConfig(ref),
		"./apis/datadoghq/v2alpha1.OTLPGRPCConfig":                    schema__apis_datadoghq_v2alpha1_OTLPGRPCConfig(ref),
		"./apis/datadoghq/v2alpha1.OTLPHTTPConfig":                    schema__apis_datadoghq_v2alpha1_OTLPHTTPConfig(ref),
		"./apis/datadoghq/v2alpha1.OTLPProtocolsConfig":               schema__apis_datadoghq_v2alpha1_OTLPProtocolsConfig(ref),
		"./apis/datadoghq/v2alpha1.OTLPReceiverConfig":                schema__apis_datadoghq_v2alpha1_OTLPReceiverConfig(ref),
		"./apis/datadoghq/v2alpha1.OrchestratorExplorerFeatureConfig": schema__apis_datadoghq_v2alpha1_OrchestratorExplorerFeatureConfig(ref),
		"./apis/datadoghq/v2alpha1.PrometheusScrapeFeatureConfig":     schema__apis_datadoghq_v2alpha1_PrometheusScrapeFeatureConfig(ref),
		"./apis/datadoghq/v2alpha1.UnixDomainSocketConfig":            schema__apis_datadoghq_v2alpha1_UnixDomainSocketConfig(ref),
	}
}

func schema__apis_datadoghq_v2alpha1_CustomConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CustomConfig provides a place for custom configuration of the Agent or Cluster Agent, corresponding to datadog.yaml, system-probe.yaml, security-agent.yaml or datadog-cluster.yaml. The configuration can be provided in the ConfigData field as raw data, or referenced in a ConfigMap. Note: `ConfigData` and `ConfigMap` cannot be set together.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"configData": {
						SchemaProps: spec.SchemaProps{
							Description: "ConfigData corresponds to the configuration file content.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"configMap": {
						SchemaProps: spec.SchemaProps{
							Description: "ConfigMap references an existing ConfigMap with the configuration file content.",
							Ref:         ref("github.com/DataDog/datadog-operator/apis/datadoghq/common/v1.ConfigMapConfig"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/DataDog/datadog-operator/apis/datadoghq/common/v1.ConfigMapConfig"},
	}
}

func schema__apis_datadoghq_v2alpha1_DatadogAgent(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DatadogAgent Deployment with the Datadog Operator.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("./apis/datadoghq/v2alpha1.DatadogAgentSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("./apis/datadoghq/v2alpha1.DatadogAgentStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./apis/datadoghq/v2alpha1.DatadogAgentSpec", "./apis/datadoghq/v2alpha1.DatadogAgentStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema__apis_datadoghq_v2alpha1_DatadogAgentGenericContainer(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DatadogAgentGenericContainer is the generic structure describing any container's common configuration.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name of the container that is overridden",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"logLevel": {
						SchemaProps: spec.SchemaProps{
							Description: "LogLevel sets logging verbosity (overrides global setting) Valid log levels are: trace, debug, info, warn, error, critical, and off. Default: 'info'",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"env": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"name",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Specify additional environmental variables in the container See also: https://docs.datadoghq.com/agent/kubernetes/?tab=helm#environment-variables",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/api/core/v1.EnvVar"),
									},
								},
							},
						},
					},
					"volumeMounts": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"name",
									"mountPath",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Specify additional volume mounts in the container.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/api/core/v1.VolumeMount"),
									},
								},
							},
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Description: "Specify the Request and Limits of the pods To get guaranteed QoS class, specify requests and limits equal. See also: http://kubernetes.io/docs/user-guide/compute-resources/",
							Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"command": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Command allows the specification of a custom entrypoint for container",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"args": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Args allows the specification of extra args to the `Command` parameter",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"healthPort": {
						SchemaProps: spec.SchemaProps{
							Description: "HealthPort of the container for the internal liveness probe. Must be the same as the Liveness/Readiness probes.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"readinessProbe": {
						SchemaProps: spec.SchemaProps{
							Description: "Configure the Readiness Probe of the container",
							Ref:         ref("k8s.io/api/core/v1.Probe"),
						},
					},
					"livenessProbe": {
						SchemaProps: spec.SchemaProps{
							Description: "Configure the Liveness Probe of the container",
							Ref:         ref("k8s.io/api/core/v1.Probe"),
						},
					},
					"securityContext": {
						SchemaProps: spec.SchemaProps{
							Description: "Container-level SecurityContext.",
							Ref:         ref("k8s.io/api/core/v1.SecurityContext"),
						},
					},
					"appArmorProfileName": {
						SchemaProps: spec.SchemaProps{
							Description: "AppArmorProfileName specify a apparmor profile.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.EnvVar", "k8s.io/api/core/v1.Probe", "k8s.io/api/core/v1.ResourceRequirements", "k8s.io/api/core/v1.SecurityContext", "k8s.io/api/core/v1.VolumeMount"},
	}
}

func schema__apis_datadoghq_v2alpha1_DatadogAgentStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DatadogAgentStatus defines the observed state of DatadogAgent.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"conditions": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-map-keys": []interface{}{
									"type",
								},
								"x-kubernetes-list-type": "map",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Conditions Represents the latest available observations of a DatadogAgent's current state.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.Condition"),
									},
								},
							},
						},
					},
					"agent": {
						SchemaProps: spec.SchemaProps{
							Description: "The actual state of the Agent as an extended daemonset.",
							Ref:         ref("github.com/DataDog/datadog-operator/apis/datadoghq/common/v1.DaemonSetStatus"),
						},
					},
					"clusterAgent": {
						SchemaProps: spec.SchemaProps{
							Description: "The actual state of the Cluster Agent as a deployment.",
							Ref:         ref("github.com/DataDog/datadog-operator/apis/datadoghq/common/v1.DeploymentStatus"),
						},
					},
					"clusterChecksRunner": {
						SchemaProps: spec.SchemaProps{
							Description: "The actual state of the Cluster Checks Runner as a deployment.",
							Ref:         ref("github.com/DataDog/datadog-operator/apis/datadoghq/common/v1.DeploymentStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/DataDog/datadog-operator/apis/datadoghq/common/v1.DaemonSetStatus", "github.com/DataDog/datadog-operator/apis/datadoghq/common/v1.DeploymentStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.Condition"},
	}
}

func schema__apis_datadoghq_v2alpha1_DatadogCredentials(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DatadogCredentials is a generic structure that holds credentials to access Datadog.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"apiKey": {
						SchemaProps: spec.SchemaProps{
							Description: "APIKey configures your Datadog API key. See also: https://app.datadoghq.com/account/settings#agent/kubernetes",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiSecret": {
						SchemaProps: spec.SchemaProps{
							Description: "APISecret references an existing Secret which stores the API key instead of creating a new one. If set, this parameter takes precedence over \"APIKey\".",
							Ref:         ref("github.com/DataDog/datadog-operator/apis/datadoghq/common/v1.SecretConfig"),
						},
					},
					"appKey": {
						SchemaProps: spec.SchemaProps{
							Description: "AppKey configures your Datadog application key. If you are using clusterAgent.metricsProvider.enabled = true, you must set a Datadog application key for read access to your metrics.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"appSecret": {
						SchemaProps: spec.SchemaProps{
							Description: "AppSecret references an existing Secret which stores the application key instead of creating a new one. If set, this parameter takes precedence over \"AppKey\".",
							Ref:         ref("github.com/DataDog/datadog-operator/apis/datadoghq/common/v1.SecretConfig"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/DataDog/datadog-operator/apis/datadoghq/common/v1.SecretConfig"},
	}
}

func schema__apis_datadoghq_v2alpha1_DatadogFeatures(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DatadogFeatures are features running on the Agent and Cluster Agent.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"logCollection": {
						SchemaProps: spec.SchemaProps{
							Description: "LogCollection configuration.",
							Ref:         ref("./apis/datadoghq/v2alpha1.LogCollectionFeatureConfig"),
						},
					},
					"liveProcessCollection": {
						SchemaProps: spec.SchemaProps{
							Description: "LiveProcessCollection configuration.",
							Ref:         ref("./apis/datadoghq/v2alpha1.LiveProcessCollectionFeatureConfig"),
						},
					},
					"liveContainerCollection": {
						SchemaProps: spec.SchemaProps{
							Description: "LiveContainerCollection configuration.",
							Ref:         ref("./apis/datadoghq/v2alpha1.LiveContainerCollectionFeatureConfig"),
						},
					},
					"oomKill": {
						SchemaProps: spec.SchemaProps{
							Description: "OOMKill configuration.",
							Ref:         ref("./apis/datadoghq/v2alpha1.OOMKillFeatureConfig"),
						},
					},
					"tcpQueueLength": {
						SchemaProps: spec.SchemaProps{
							Description: "TCPQueueLength configuration.",
							Ref:         ref("./apis/datadoghq/v2alpha1.TCPQueueLengthFeatureConfig"),
						},
					},
					"apm": {
						SchemaProps: spec.SchemaProps{
							Description: "APM (Application Performance Monitoring) configuration.",
							Ref:         ref("./apis/datadoghq/v2alpha1.APMFeatureConfig"),
						},
					},
					"cspm": {
						SchemaProps: spec.SchemaProps{
							Description: "CSPM (Cloud Security Posture Management) configuration.",
							Ref:         ref("./apis/datadoghq/v2alpha1.CSPMFeatureConfig"),
						},
					},
					"cws": {
						SchemaProps: spec.SchemaProps{
							Description: "CWS (Cloud Workload Security) configuration.",
							Ref:         ref("./apis/datadoghq/v2alpha1.CWSFeatureConfig"),
						},
					},
					"npm": {
						SchemaProps: spec.SchemaProps{
							Description: "NPM (Network Performance Monitoring) configuration.",
							Ref:         ref("./apis/datadoghq/v2alpha1.NPMFeatureConfig"),
						},
					},
					"usm": {
						SchemaProps: spec.SchemaProps{
							Description: "USM (Universal Service Monitoring) configuration.",
							Ref:         ref("./apis/datadoghq/v2alpha1.USMFeatureConfig"),
						},
					},
					"dogstatsd": {
						SchemaProps: spec.SchemaProps{
							Description: "Dogstatsd configuration.",
							Ref:         ref("./apis/datadoghq/v2alpha1.DogstatsdFeatureConfig"),
						},
					},
					"otlp": {
						SchemaProps: spec.SchemaProps{
							Description: "OTLP ingest configuration",
							Ref:         ref("./apis/datadoghq/v2alpha1.OTLPFeatureConfig"),
						},
					},
					"eventCollection": {
						SchemaProps: spec.SchemaProps{
							Description: "EventCollection configuration.",
							Ref:         ref("./apis/datadoghq/v2alpha1.EventCollectionFeatureConfig"),
						},
					},
					"orchestratorExplorer": {
						SchemaProps: spec.SchemaProps{
							Description: "OrchestratorExplorer check configuration.",
							Ref:         ref("./apis/datadoghq/v2alpha1.OrchestratorExplorerFeatureConfig"),
						},
					},
					"kubeStateMetricsCore": {
						SchemaProps: spec.SchemaProps{
							Description: "KubeStateMetricsCore check configuration.",
							Ref:         ref("./apis/datadoghq/v2alpha1.KubeStateMetricsCoreFeatureConfig"),
						},
					},
					"admissionController": {
						SchemaProps: spec.SchemaProps{
							Description: "AdmissionController configuration.",
							Ref:         ref("./apis/datadoghq/v2alpha1.AdmissionControllerFeatureConfig"),
						},
					},
					"externalMetricsServer": {
						SchemaProps: spec.SchemaProps{
							Description: "ExternalMetricsServer configuration.",
							Ref:         ref("./apis/datadoghq/v2alpha1.ExternalMetricsServerFeatureConfig"),
						},
					},
					"clusterChecks": {
						SchemaProps: spec.SchemaProps{
							Description: "ClusterChecks configuration.",
							Ref:         ref("./apis/datadoghq/v2alpha1.ClusterChecksFeatureConfig"),
						},
					},
					"prometheusScrape": {
						SchemaProps: spec.SchemaProps{
							Description: "PrometheusScrape configuration.",
							Ref:         ref("./apis/datadoghq/v2alpha1.PrometheusScrapeFeatureConfig"),
						},
					},
					"datadogMonitor": {
						SchemaProps: spec.SchemaProps{
							Description: "DatadogMonitor configuration.",
							Ref:         ref("./apis/datadoghq/v2alpha1.DatadogMonitorFeatureConfig"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./apis/datadoghq/v2alpha1.APMFeatureConfig", "./apis/datadoghq/v2alpha1.AdmissionControllerFeatureConfig", "./apis/datadoghq/v2alpha1.CSPMFeatureConfig", "./apis/datadoghq/v2alpha1.CWSFeatureConfig", "./apis/datadoghq/v2alpha1.ClusterChecksFeatureConfig", "./apis/datadoghq/v2alpha1.DatadogMonitorFeatureConfig", "./apis/datadoghq/v2alpha1.DogstatsdFeatureConfig", "./apis/datadoghq/v2alpha1.EventCollectionFeatureConfig", "./apis/datadoghq/v2alpha1.ExternalMetricsServerFeatureConfig", "./apis/datadoghq/v2alpha1.KubeStateMetricsCoreFeatureConfig", "./apis/datadoghq/v2alpha1.LiveContainerCollectionFeatureConfig", "./apis/datadoghq/v2alpha1.LiveProcessCollectionFeatureConfig", "./apis/datadoghq/v2alpha1.LogCollectionFeatureConfig", "./apis/datadoghq/v2alpha1.NPMFeatureConfig", "./apis/datadoghq/v2alpha1.OOMKillFeatureConfig", "./apis/datadoghq/v2alpha1.OTLPFeatureConfig", "./apis/datadoghq/v2alpha1.OrchestratorExplorerFeatureConfig", "./apis/datadoghq/v2alpha1.PrometheusScrapeFeatureConfig", "./apis/datadoghq/v2alpha1.TCPQueueLengthFeatureConfig", "./apis/datadoghq/v2alpha1.USMFeatureConfig"},
	}
}

func schema__apis_datadoghq_v2alpha1_DogstatsdFeatureConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "DogstatsdFeatureConfig contains the Dogstatsd configuration parameters.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"originDetectionEnabled": {
						SchemaProps: spec.SchemaProps{
							Description: "OriginDetectionEnabled enables origin detection for container tagging. See also: https://docs.datadoghq.com/developers/dogstatsd/unix_socket/#using-origin-detection-for-container-tagging",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"hostPortConfig": {
						SchemaProps: spec.SchemaProps{
							Description: "HostPortConfig contains host port configuration. Enabled Default: true Port Default: 8125",
							Ref:         ref("./apis/datadoghq/v2alpha1.HostPortConfig"),
						},
					},
					"unixDomainSocketConfig": {
						SchemaProps: spec.SchemaProps{
							Description: "UnixDomainSocketConfig contains socket configuration. See also: https://docs.datadoghq.com/agent/kubernetes/apm/?tab=helm#agent-environment-variables Enabled Default: true Path Default: `/var/run/datadog/statsd/dsd.socket`",
							Ref:         ref("./apis/datadoghq/v2alpha1.UnixDomainSocketConfig"),
						},
					},
					"mapperProfiles": {
						SchemaProps: spec.SchemaProps{
							Description: "Configure the Dogstasd Mapper Profiles. Can be passed as raw data or via a json encoded string in a config map. See also: https://docs.datadoghq.com/developers/dogstatsd/dogstatsd_mapper/",
							Ref:         ref("./apis/datadoghq/v2alpha1.CustomConfig"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./apis/datadoghq/v2alpha1.CustomConfig", "./apis/datadoghq/v2alpha1.HostPortConfig", "./apis/datadoghq/v2alpha1.UnixDomainSocketConfig"},
	}
}

func schema__apis_datadoghq_v2alpha1_EventCollectionFeatureConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "EventCollectionFeatureConfig contains the Event Collection configuration.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"collectKubernetesEvents": {
						SchemaProps: spec.SchemaProps{
							Description: "CollectKubernetesEvents enables Kubernetes event collection. Default: true",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema__apis_datadoghq_v2alpha1_KubeStateMetricsCoreFeatureConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KubeStateMetricsCoreFeatureConfig contains the Kube State Metrics Core check feature configuration. The Kube State Metrics Core check runs in the Cluster Agent (or Cluster Check Runners). See also: https://docs.datadoghq.com/integrations/kubernetes_state_core",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"enabled": {
						SchemaProps: spec.SchemaProps{
							Description: "Enabled enables Kube State Metrics Core. Default: true",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"conf": {
						SchemaProps: spec.SchemaProps{
							Description: "Conf overrides the configuration for the default Kubernetes State Metrics Core check. This must point to a ConfigMap containing a valid cluster check configuration.",
							Ref:         ref("./apis/datadoghq/v2alpha1.CustomConfig"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./apis/datadoghq/v2alpha1.CustomConfig"},
	}
}

func schema__apis_datadoghq_v2alpha1_LocalService(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "LocalService provides the internal traffic policy service configuration.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"nameOverride": {
						SchemaProps: spec.SchemaProps{
							Description: "NameOverride defines the name of the internal traffic service to target the agent running on the local node.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"forceEnableLocalService": {
						SchemaProps: spec.SchemaProps{
							Description: "ForceEnableLocalService forces the creation of the internal traffic policy service to target the agent running on the local node. This parameter only applies to Kubernetes 1.21, where the feature is in alpha and is disabled by default. (On Kubernetes 1.22+, the feature entered beta and the internal traffic service is created by default, so this parameter is ignored.) Default: false",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema__apis_datadoghq_v2alpha1_MultiCustomConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "MultiCustomConfig provides a place for custom configuration of the Agent or Cluster Agent, corresponding to /confd/*.yaml. The configuration can be provided in the ConfigDataMap field as raw data, or referenced in a single ConfigMap. Note: `ConfigDataMap` and `ConfigMap` cannot be set together.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"configDataMap": {
						SchemaProps: spec.SchemaProps{
							Description: "ConfigDataMap corresponds to the content of the configuration files. They key should be the filename the contents get mounted to; for instance check.py or check.yaml.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Allows: true,
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"configMap": {
						SchemaProps: spec.SchemaProps{
							Description: "ConfigMap references an existing ConfigMap with the content of the configuration files.",
							Ref:         ref("github.com/DataDog/datadog-operator/apis/datadoghq/common/v1.ConfigMapConfig"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/DataDog/datadog-operator/apis/datadoghq/common/v1.ConfigMapConfig"},
	}
}

func schema__apis_datadoghq_v2alpha1_NetworkPolicyConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NetworkPolicyConfig provides Network Policy configuration for the agents.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"create": {
						SchemaProps: spec.SchemaProps{
							Description: "Create defines whether to create a NetworkPolicy for the current deployment.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"flavor": {
						SchemaProps: spec.SchemaProps{
							Description: "Flavor defines Which network policy to use.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"dnsSelectorEndpoints": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "DNSSelectorEndpoints defines the cilium selector of the DNS server entity.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"},
	}
}

func schema__apis_datadoghq_v2alpha1_OTLPFeatureConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OTLPFeatureConfig contains configuration for OTLP ingest.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"receiver": {
						SchemaProps: spec.SchemaProps{
							Description: "Receiver contains configuration for the OTLP ingest receiver.",
							Default:     map[string]interface{}{},
							Ref:         ref("./apis/datadoghq/v2alpha1.OTLPReceiverConfig"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./apis/datadoghq/v2alpha1.OTLPReceiverConfig"},
	}
}

func schema__apis_datadoghq_v2alpha1_OTLPGRPCConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OTLPGRPCConfig contains configuration for the OTLP ingest OTLP/gRPC receiver.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"enabled": {
						SchemaProps: spec.SchemaProps{
							Description: "Enable the OTLP/gRPC endpoint.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"endpoint": {
						SchemaProps: spec.SchemaProps{
							Description: "Endpoint for OTLP/gRPC. gRPC supports several naming schemes: https://github.com/grpc/grpc/blob/master/doc/naming.md The Datadog Operator supports only 'host:port' (usually `0.0.0.0:port`). Default: `0.0.0.0:4317`.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema__apis_datadoghq_v2alpha1_OTLPHTTPConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OTLPHTTPConfig contains configuration for the OTLP ingest OTLP/HTTP receiver.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"enabled": {
						SchemaProps: spec.SchemaProps{
							Description: "Enable the OTLP/HTTP endpoint.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"endpoint": {
						SchemaProps: spec.SchemaProps{
							Description: "Endpoint for OTLP/HTTP. Default: '0.0.0.0:4318'.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema__apis_datadoghq_v2alpha1_OTLPProtocolsConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OTLPProtocolsConfig contains configuration for the OTLP ingest receiver protocols.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"grpc": {
						SchemaProps: spec.SchemaProps{
							Description: "GRPC contains configuration for the OTLP ingest OTLP/gRPC receiver.",
							Ref:         ref("./apis/datadoghq/v2alpha1.OTLPGRPCConfig"),
						},
					},
					"http": {
						SchemaProps: spec.SchemaProps{
							Description: "HTTP contains configuration for the OTLP ingest OTLP/HTTP receiver.",
							Ref:         ref("./apis/datadoghq/v2alpha1.OTLPHTTPConfig"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./apis/datadoghq/v2alpha1.OTLPGRPCConfig", "./apis/datadoghq/v2alpha1.OTLPHTTPConfig"},
	}
}

func schema__apis_datadoghq_v2alpha1_OTLPReceiverConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OTLPReceiverConfig contains configuration for the OTLP ingest receiver.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"protocols": {
						SchemaProps: spec.SchemaProps{
							Description: "Protocols contains configuration for the OTLP ingest receiver protocols.",
							Default:     map[string]interface{}{},
							Ref:         ref("./apis/datadoghq/v2alpha1.OTLPProtocolsConfig"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./apis/datadoghq/v2alpha1.OTLPProtocolsConfig"},
	}
}

func schema__apis_datadoghq_v2alpha1_OrchestratorExplorerFeatureConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "OrchestratorExplorerFeatureConfig contains the Orchestrator Explorer check feature configuration. The Orchestrator Explorer check runs in the Process and Cluster Agents (or Cluster Check Runners). See also: https://docs.datadoghq.com/infrastructure/livecontainers/#kubernetes-resources",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"enabled": {
						SchemaProps: spec.SchemaProps{
							Description: "Enabled enables the Orchestrator Explorer. Default: true",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"conf": {
						SchemaProps: spec.SchemaProps{
							Description: "Conf overrides the configuration for the default Orchestrator Explorer check. This must point to a ConfigMap containing a valid cluster check configuration.",
							Ref:         ref("./apis/datadoghq/v2alpha1.CustomConfig"),
						},
					},
					"scrubContainers": {
						SchemaProps: spec.SchemaProps{
							Description: "ScrubContainers enables scrubbing of sensitive container data (passwords, tokens, etc. ). Default: true",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"extraTags": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Additional tags to associate with the collected data in the form of `a b c`. This is a Cluster Agent option distinct from DD_TAGS that is used in the Orchestrator Explorer.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"ddUrl": {
						SchemaProps: spec.SchemaProps{
							Description: "Override the API endpoint for the Orchestrator Explorer. URL Default: \"https://orchestrator.datadoghq.com\".",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./apis/datadoghq/v2alpha1.CustomConfig"},
	}
}

func schema__apis_datadoghq_v2alpha1_PrometheusScrapeFeatureConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PrometheusScrapeFeatureConfig allows configuration of the Prometheus Autodiscovery feature.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"enabled": {
						SchemaProps: spec.SchemaProps{
							Description: "Enable autodiscovery of pods and services exposing Prometheus metrics. Default: false",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"enableServiceEndpoints": {
						SchemaProps: spec.SchemaProps{
							Description: "EnableServiceEndpoints enables generating dedicated checks for service endpoints. Default: false",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"additionalConfigs": {
						SchemaProps: spec.SchemaProps{
							Description: "AdditionalConfigs allows adding advanced Prometheus check configurations with custom discovery rules.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema__apis_datadoghq_v2alpha1_UnixDomainSocketConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "UnixDomainSocketConfig contains the Unix Domain Socket configuration.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"enabled": {
						SchemaProps: spec.SchemaProps{
							Description: "Enabled enables Unix Domain Socket. Default: true",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"path": {
						SchemaProps: spec.SchemaProps{
							Description: "Path defines the socket path used when enabled.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}
