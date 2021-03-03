package stepper

import (
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	v1 "k8s.io/api/core/v1"
)

// OverrideStep overrides the step with the given overrides
func OverrideStep(step *v1beta1.Step, override *v1beta1.Step) {
	if override.Name != "" {
		step.Name = override.Name
	}
	if len(override.Command) > 0 {
		step.Script = override.Script
		step.Command = override.Command
		step.Args = override.Args
	}
	if override.Script != "" {
		step.Script = override.Script
		step.Command = nil
		step.Args = nil
	}
	if override.Timeout != nil {
		step.Timeout = override.Timeout
	}
	if override.WorkingDir != "" {
		step.WorkingDir = override.WorkingDir
	}
	if string(override.ImagePullPolicy) != "" {
		step.ImagePullPolicy = override.ImagePullPolicy
	}
	step.Env = OverrideEnv(step.Env, override.Env)
	step.EnvFrom = OverrideEnvFrom(step.EnvFrom, override.EnvFrom)
	step.VolumeMounts = OverrideVolumeMounts(step.VolumeMounts, override.VolumeMounts)
}

// OverrideEnv override either replaces or adds the given env vars
func OverrideEnv(from []v1.EnvVar, overrides []v1.EnvVar) []v1.EnvVar {
	for _, override := range overrides {
		found := false
		for i := range from {
			f := &from[i]
			if f.Name == override.Name {
				found = true
				*f = override
				break
			}
		}
		if !found {
			from = append(from, override)
		}
	}
	return from
}

// OverrideEnvFrom override either replaces or adds the given env froms
func OverrideEnvFrom(from []v1.EnvFromSource, overrides []v1.EnvFromSource) []v1.EnvFromSource {
	for _, override := range overrides {
		found := false
		for i := range from {
			f := &from[i]
			if f.ConfigMapRef != nil && override.ConfigMapRef != nil && f.ConfigMapRef.Name == override.ConfigMapRef.Name {
				found = true
				*f = override
				break
			}
			if f.SecretRef != nil && override.SecretRef != nil && f.SecretRef.Name == override.SecretRef.Name {
				found = true
				*f = override
				break
			}
		}
		if !found {
			from = append(from, override)
		}
	}
	return from
}

// OverrideVolumeMounts override either replaces or adds the given volume mounts
func OverrideVolumeMounts(from []v1.VolumeMount, overrides []v1.VolumeMount) []v1.VolumeMount {
	for _, override := range overrides {
		found := false
		for i := range from {
			f := &from[i]
			if f.Name == override.Name {
				found = true
				*f = override
				break
			}
		}
		if !found {
			from = append(from, override)
		}
	}
	return from
}
