/**
 * Created by zc on 2020/8/4.
**/
package drone

import (
	"github.com/drone/drone-runtime/engine"
	"github.com/drone/drone-yaml/yaml"
	"github.com/drone/drone-yaml/yaml/compiler"
	"github.com/drone/drone-yaml/yaml/linter"
	"luban/pkg/logger"
	"strings"
)

func Compile(pipeline *yaml.Pipeline) (*engine.Spec, error) {
	if err := linter.Lint(pipeline, false); err != nil {
		logger.Warnln("runner: yaml lint errors", err)
		return nil, err
	}
	comp := new(compiler.Compiler)
	comp.PrivilegedFunc = compiler.DindFunc([]string{"plugins/docker"})
	spec := comp.Compile(pipeline)
	return spec, nil
}

func Preset(spec *engine.Spec, configData []byte) {
	name := "luban.config"
	spec.Files = append(spec.Files, &engine.File{
		Metadata: engine.Metadata{
			UID:       RandString(),
			Namespace: spec.Metadata.Namespace,
			Name:      name,
		},
		Data: configData,
	})
	for _, step := range spec.Steps {
		if step.WorkingDir == "" {
			step.WorkingDir = "/luban"
		} else {
			step.WorkingDir = strings.TrimRight(step.WorkingDir, "/")
		}
		step.Files = append(step.Files, &engine.FileMount{
			Name: name,
			Path: step.WorkingDir + "/" + name,
		})
	}
}
