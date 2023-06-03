package terraform

import "github.com/hashicorp/terraform/internal/configs"

var (
	DefaultEvaluator           = new(Evaluator)
	DefaultEvaluationStateData = new(evaluationStateData)
)

func init() {
	DefaultEvaluator.Meta = &ContextMeta{Env: "default"}
	DefaultEvaluator.Config = &configs.Config{Module: &configs.Module{SourceDir: "."}}
	DefaultEvaluationStateData.Evaluator = DefaultEvaluator
}
