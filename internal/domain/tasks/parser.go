package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/DanielTitkov/dashboars/internal/domain"
	"github.com/DanielTitkov/dashboars/internal/service/parser"
	"github.com/DanielTitkov/dashboars/internal/util"
	"github.com/hashicorp/go-multierror"
)

const (
	argParser  = "parser"
	argMetrics = "metrics"
)

type ParserTaskArgs struct {
	Parser  []parser.Config `json:"parser"`
	Metrics []domain.Metric `json:"metrics"`
}

func NewParserTask(args domain.CreateTaskArgs, cat *domain.TaskCategory, tags []*domain.TaskTag) (*domain.Task, error) {
	var taskArgs ParserTaskArgs
	err := json.Unmarshal(args.Args, &taskArgs)
	if err != nil {
		return nil, err
	}

	return &domain.Task{
		ID:          args.ID,
		Type:        args.Type,
		Code:        args.Code,
		Title:       args.Title,
		Description: args.Description,
		Active:      args.Active,
		Display:     args.Display,
		Schedule:    args.Schedule,
		Args:        &taskArgs,
		Category:    cat,
		Tags:        tags,
		ResolveFn:   ParserTaskResolveFn,
	}, nil
}

func (a *ParserTaskArgs) Get(g string) (interface{}, bool) {
	var ret interface{}
	switch g {
	case argParser:
		return a.Parser, true
	case argMetrics:
		return a.Metrics, true
	}
	return ret, false
}

func (a *ParserTaskArgs) GetString(string) (string, bool) {
	return "", false
}

func (a *ParserTaskArgs) GetFloat(s string) (float64, bool) {
	return 0, false
}

func (a *ParserTaskArgs) GetInt(string) (int, bool) {
	return 0, false
}

func (a *ParserTaskArgs) ToMap() map[string]interface{} {
	return util.ToMap(a)
}

func (a *ParserTaskArgs) ToJSON() json.RawMessage {
	ret, err := json.Marshal(a)
	if err != nil {
		panic(err) // FIXME
	}
	return ret
}

func ParserTaskResolveFn(ctx context.Context, t *domain.Task, ti *domain.TaskInstance) (*domain.TaskInstance, error) {
	parserCfg, err := getParserConfig(t.Args)
	if err != nil {
		return ti.WithError(err, nil), err
	}

	metricsMap, err := makeMetricsMap(t)
	if err != nil {
		return ti.WithError(err, nil), err
	}

	var resultItems []*domain.Item
	var errs error
	for _, cfg := range parserCfg {
		p := parser.New(&cfg)
		// fmt.Println("RUNNING PARSER TASK", cfg) // FIXME
		result, err := p.Run()
		if err != nil {
			errs = multierror.Append(errs, err)
			continue
		}

		for _, it := range result.Items {
			metric, ok := metricsMap[it.Name]
			if !ok {
				errs = multierror.Append(errs, fmt.Errorf("metric not found: %s, must be set in task arguments", it.Name))
				continue
			}

			value, err := strconv.ParseFloat(it.Value, 64)
			if err != nil {
				errs = multierror.Append(errs, err)
				continue
			}

			var dimensions []*domain.Dimension
			for _, attr := range it.Attrs {
				dimensions = append(dimensions, &domain.Dimension{
					Title: attr.Name,
					Value: attr.Value,
				})
			}

			resultItems = append(resultItems, &domain.Item{
				Value:      value,
				Timestamp:  time.Now(),
				Dimensions: dimensions,
				Meta:       make(map[string]interface{}),
				Metric:     &metric,
			})
		}

		// fmt.Println("PARSER RESULT", result) // FIXME
	}

	if errs != nil {
		return ti.WithError(errs, resultItems), errs
	}

	return ti.WithSuccess(resultItems), nil
}

func makeMetricsMap(task *domain.Task) (map[string]domain.Metric, error) {
	metricsI, ok := task.Args.Get(argMetrics)
	if !ok {
		return nil, newArgError(argMetrics)
	}

	metrics, ok := metricsI.([]domain.Metric)
	if !ok {
		return nil, newArgTypeError(argMetrics, "[]domain.Metric", fmt.Sprintf("%T", metricsI))
	}

	metricsMap := make(map[string]domain.Metric)
	for _, m := range metrics {
		m.TaskID = task.ID
		metricsMap[m.Title] = m
	}

	return metricsMap, nil
}

func getParserConfig(args domain.TaskArgs) ([]parser.Config, error) {
	parserCfgsI, ok := args.Get(argParser)
	if !ok {
		return nil, newArgError(argParser)
	}

	parserCfg, ok := parserCfgsI.([]parser.Config)
	if !ok {
		return nil, newArgTypeError(argParser, "[]parser.Config", fmt.Sprintf("%T", parserCfgsI))
	}

	return parserCfg, nil
}
