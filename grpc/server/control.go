package server

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/benbjohnson/clock"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	commonv1 "github.com/nanzhong/tstr/api/common/v1"
	controlv1 "github.com/nanzhong/tstr/api/control/v1"
	"github.com/nanzhong/tstr/db"
	"github.com/nanzhong/tstr/grpc/types"
	"github.com/nanzhong/tstr/scheduler"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ControlServer struct {
	controlv1.UnimplementedControlServiceServer
	pgxPool    *pgxpool.Pool
	dbQuerier  db.Querier
	cronParser cron.Parser
	clock      clock.Clock
}

func NewControlServer(pgxPool *pgxpool.Pool) controlv1.ControlServiceServer {
	return &ControlServer{
		pgxPool:    pgxPool,
		dbQuerier:  db.New(),
		cronParser: cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor),
		clock:      clock.New(),
	}
}

func (s *ControlServer) RegisterTest(ctx context.Context, r *controlv1.RegisterTestRequest) (*controlv1.RegisterTestResponse, error) {
	ns, err := namespaceFromContext(ctx)
	if err != nil {
		return nil, err
	}

	labels := pgtype.JSONB{}
	if err := labels.Set(r.Labels); err != nil {
		log.Error().Err(err).Msg("failed to parse test labels")
		return nil, status.Error(codes.InvalidArgument, "failed to parse labels")
	}

	matrix := pgtype.JSONB{}
	if err := matrix.Set(types.FromProtoTestMatrix(r.Matrix)); err != nil {
		log.Error().Err(err).Msg("failed to parse test matrix")
		return nil, status.Error(codes.InvalidArgument, "failed to parse matrix")
	}

	runConfig := pgtype.JSONB{}
	if err := runConfig.Set(types.FromProtoTestRunConfig(r.RunConfig)); err != nil {
		log.Error().Err(err).Msg("failed to parse run config")
		return nil, status.Error(codes.InvalidArgument, "failed to parse run config")
	}

	var nextRunAt sql.NullTime
	if r.CronSchedule != "" {
		schedule, err := s.cronParser.Parse(r.CronSchedule)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid cron schedule")
		}
		nextRunAt.Valid = true
		nextRunAt.Time = schedule.Next(s.clock.Now())
	}

	var test db.Test
	err = s.pgxPool.BeginFunc(ctx, func(tx pgx.Tx) error {
		var err error
		test, err = s.dbQuerier.RegisterTest(ctx, tx, db.RegisterTestParams{
			Namespace:    ns,
			Name:         r.Name,
			Labels:       labels,
			Matrix:       matrix,
			RunConfig:    runConfig,
			CronSchedule: sql.NullString{Valid: r.CronSchedule != "", String: r.CronSchedule},
			NextRunAt:    nextRunAt,
		})
		if err != nil {
			return fmt.Errorf("registering test: %w", err)
		}

		_, err = s.scheduleRuns(ctx, tx, &test)
		if err != nil {
			return fmt.Errorf("scheduling runs for test: %w", err)
		}

		return nil
	})
	if err != nil {
		log.Error().Err(err).Msg("failed to store test")
		return nil, status.Error(codes.Internal, "failed to register test")
	}

	pbTest, err := types.ToProtoTest(&test)
	if err != nil {
		log.Error().
			Err(err).
			Stringer("test_id", test.ID).
			Msg("failed to format test")
		return nil, status.Error(codes.Internal, "failed to format test")
	}

	return &controlv1.RegisterTestResponse{
		Test: pbTest,
	}, nil
}

func (s *ControlServer) GetTest(ctx context.Context, r *controlv1.GetTestRequest) (*controlv1.GetTestResponse, error) {
	ns, err := namespaceFromContext(ctx)
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(r.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid test id")
	}

	test, err := s.dbQuerier.GetTest(ctx, s.pgxPool, db.GetTestParams{
		ID:        id,
		Namespace: ns,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("namespace", ns).
			Stringer("test_id", id).
			Msg("failed to get test")
		return nil, status.Error(codes.Internal, "failed to get test")
	}

	pbTest, err := types.ToProtoTest(&test)
	if err != nil {
		log.Error().
			Err(err).
			Stringer("test_id", test.ID).
			Msg("failed to format test")
		return nil, status.Error(codes.Internal, "failed to format test")
	}

	return &controlv1.GetTestResponse{
		Test: pbTest,
	}, nil
}

func (s *ControlServer) ListTests(ctx context.Context, r *controlv1.ListTestsRequest) (*controlv1.ListTestsResponse, error) {
	ns, err := namespaceFromContext(ctx)
	if err != nil {
		return nil, err
	}

	tests, err := s.dbQuerier.ListTests(ctx, s.pgxPool, ns)
	if err != nil {
		log.Error().
			Err(err).
			Str("namespace", ns).
			Msg("failed to list tests")
		return nil, status.Error(codes.Internal, "failed to list tests")
	}

	var pbTests []*commonv1.Test
	for _, test := range tests {
		pbTest, err := types.ToProtoTest(&test)
		if err != nil {
			log.Error().
				Err(err).
				Stringer("test_id", test.ID).
				Msg("failed to format test")
			return nil, status.Error(codes.Internal, "failed to format test")
		}
		pbTests = append(pbTests, pbTest)
	}

	return &controlv1.ListTestsResponse{
		Tests: pbTests,
	}, nil
}

func (s *ControlServer) UpdateTest(ctx context.Context, r *controlv1.UpdateTestRequest) (*controlv1.UpdateTestResponse, error) {
	ns, err := namespaceFromContext(ctx)
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(r.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid test id")
	}

	test, err := s.dbQuerier.GetTest(ctx, s.pgxPool, db.GetTestParams{
		ID:        id,
		Namespace: ns,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("namespace", ns).
			Stringer("test_id", id).
			Msg("failed to get test")
		return nil, status.Error(codes.Internal, "failed to get test for update")
	}

	var runConfig db.TestRunConfig
	if err := test.RunConfig.AssignTo(&runConfig); err != nil {
		log.Error().
			Err(err).
			Stringer("test_id", id).
			Msg("failed to parse run config")
		return nil, status.Error(codes.Internal, "failed to parse run config")
	}

	for _, p := range r.FieldMask.Paths {
		switch p {
		case "name":
			test.Name = r.Name
		case "labels":
			if err := test.Labels.Set(r.Labels); err != nil {
				log.Error().
					Err(err).
					Dict("labels", zerolog.Dict().Fields(r.Labels)).
					Msg("failed to format labels")
				return nil, status.Error(codes.Internal, "failed to format labels")
			}
		case "matrix":
			if err := test.Matrix.Set(types.FromProtoTestMatrix(r.Matrix)); err != nil {
				log.Error().
					Err(err).
					Str("matrix", r.Matrix.String()).
					Msg("failed to format matrix")
				return nil, status.Error(codes.Internal, "failed to format matrix")

			}
		case "run_config.container_image":
			runConfig.ContainerImage = r.RunConfig.ContainerImage
		case "run_config.command":
			runConfig.Command = r.RunConfig.Command
		case "run_config.args":
			runConfig.Args = r.RunConfig.Args
		case "run_config.env":
			runConfig.Env = r.RunConfig.Env
		case "run_config.timeout":
			runConfig.TimeoutSeconds = uint(r.RunConfig.Timeout.Seconds)
		case "cron_schedule":
			test.CronSchedule = sql.NullString{Valid: r.CronSchedule != "", String: r.CronSchedule}
		}
	}

	if err := test.RunConfig.Set(&runConfig); err != nil {
		log.Error().
			Err(err).
			Msg("failed to format run config")
		return nil, status.Error(codes.Internal, "failed to format run config")
	}

	err = s.pgxPool.BeginFunc(ctx, func(tx pgx.Tx) error {
		_, err = s.dbQuerier.UpdateTest(ctx, tx, db.UpdateTestParams{
			Namespace:    test.Namespace,
			ID:           id,
			Name:         test.Name,
			Labels:       test.Labels,
			Matrix:       test.Matrix,
			RunConfig:    test.RunConfig,
			CronSchedule: test.CronSchedule,
			NextRunAt:    test.NextRunAt,
		})
		if err != nil {
			return fmt.Errorf("updating test: %w", err)
		}

		_, err = s.scheduleRuns(ctx, tx, &test)
		if err != nil {
			return fmt.Errorf("scheduling runs for test: %w", err)
		}

		return nil
	})
	if err != nil {
		log.Error().Err(err).Stringer("test_id", id).Msg("failed to update test")
		return nil, status.Error(codes.Internal, "failed to update test")
	}
	return &controlv1.UpdateTestResponse{}, nil
}

func (s *ControlServer) DeleteTest(ctx context.Context, r *controlv1.DeleteTestRequest) (*controlv1.DeleteTestResponse, error) {
	ns, err := namespaceFromContext(ctx)
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(r.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid test id")
	}

	if _, err := s.dbQuerier.GetTest(ctx, s.pgxPool, db.GetTestParams{
		ID:        id,
		Namespace: ns,
	}); err != nil {
		log.Error().
			Err(err).
			Str("namespace", ns).
			Stringer("test_id", id).
			Msg("failed to get test for delete")
		return nil, status.Error(codes.Internal, "failed to get test for delete")
	}

	err = s.pgxPool.BeginFunc(ctx, func(tx pgx.Tx) error {
		if err := s.dbQuerier.DeleteRunsForTest(ctx, tx, id); err != nil {
			return err
		}
		if _, err := s.dbQuerier.DeleteTest(ctx, tx, db.DeleteTestParams{
			ID:        id,
			Namespace: ns,
		}); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error().Err(err).Stringer("test_id", id).Msg("failed to delete test")
		return nil, status.Error(codes.Internal, "failed to delete test")
	}

	return &controlv1.DeleteTestResponse{}, nil
}

func (s *ControlServer) GetRun(ctx context.Context, r *controlv1.GetRunRequest) (*controlv1.GetRunResponse, error) {
	ns, err := namespaceFromContext(ctx)
	if err != nil {
		return nil, err
	}

	id, err := uuid.Parse(r.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid run id")
	}

	run, err := s.dbQuerier.GetRun(ctx, s.pgxPool, db.GetRunParams{
		Namespace: ns,
		ID:        id,
	})
	if err != nil {
		log.Error().
			Err(err).
			Str("namespace", ns).
			Stringer("run_id", run.ID).
			Msg("failed to get run")
		return nil, status.Error(codes.Internal, "failed to get run")
	}

	pbRun, err := types.ToProtoRun(&run)
	if err != nil {
		log.Error().
			Err(err).
			Str("namespace", ns).
			Stringer("run_id", run.ID).
			Msg("failed to format run")
		return nil, status.Error(codes.Internal, "failed to format run")
	}

	return &controlv1.GetRunResponse{
		Run: pbRun,
	}, nil
}

func (s *ControlServer) ListRuns(ctx context.Context, r *controlv1.ListRunsRequest) (*controlv1.ListRunsResponse, error) {
	ns, err := namespaceFromContext(ctx)
	if err != nil {
		return nil, err
	}

	runs, err := s.dbQuerier.ListRuns(ctx, s.pgxPool, ns)
	if err != nil {
		log.Error().
			Err(err).
			Str("namespace", ns).
			Msg("failed to list runs")
		return nil, status.Error(codes.Internal, "failed to list runs")
	}

	var pbRuns []*commonv1.Run
	for _, run := range runs {
		pbRun, err := types.ToProtoRun(&run)
		if err != nil {
			log.Error().
				Err(err).
				Str("namespace", ns).
				Stringer("run_id", run.ID).
				Msg("failed to format run")
			return nil, status.Error(codes.Internal, "failed to format run")
		}
		pbRuns = append(pbRuns, pbRun)
	}

	return &controlv1.ListRunsResponse{
		Runs: pbRuns,
	}, nil
}

func (s *ControlServer) ScheduleRun(ctx context.Context, r *controlv1.ScheduleRunRequest) (*controlv1.ScheduleRunResponse, error) {
	ns, err := namespaceFromContext(ctx)
	if err != nil {
		return nil, err
	}

	testID, err := uuid.Parse(r.TestId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid test id")
	}

	test, err := s.dbQuerier.GetTest(ctx, s.pgxPool, db.GetTestParams{
		ID:        testID,
		Namespace: ns,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to lookup test to schedule run for")
	}

	if test.Namespace != ns {
		return nil, status.Error(codes.NotFound, "failed to find test")
	}

	var labels pgtype.JSONB
	if len(r.Labels) > 0 {
		if err := labels.Set(&r.Labels); err != nil {
			log.Error().
				Err(err).
				Dict("labels", zerolog.Dict().Fields(r.Labels)).
				Msg("failed to format custom labels")
			return nil, status.Error(codes.Internal, "failed to format labels")
		}
	} else {
		labels = test.Labels
	}

	var runs []db.Run
	err = s.pgxPool.BeginFunc(ctx, func(tx pgx.Tx) error {
		var err error
		runs, err = s.scheduleRuns(ctx, tx, &test)
		return err
	})
	if err != nil {
		log.Error().
			Err(err).
			Stringer("test_id", test.ID).
			Msg("failed to schedule runs")
		return nil, status.Error(codes.Internal, "failed to schedule runs for test")
	}

	var pbRuns []*commonv1.Run
	for _, run := range runs {
		pbRun, err := types.ToProtoRun(&run)
		if err != nil {
			log.Error().
				Err(err).
				Stringer("run_id", run.ID).
				Msg("failed to format run")
			return nil, status.Error(codes.Internal, "failed to format run")
		}
		pbRuns = append(pbRuns, pbRun)
	}

	return &controlv1.ScheduleRunResponse{
		Runs: pbRuns,
	}, nil
}

func (s *ControlServer) ListRunners(ctx context.Context, r *controlv1.ListRunnersRequest) (*controlv1.ListRunnersResponse, error) {
	ns, err := namespaceFromContext(ctx)
	if err != nil {
		return nil, err
	}

	runners, err := s.dbQuerier.ListRunners(ctx, s.pgxPool)
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to list runners")
		return nil, status.Error(codes.Internal, "failed to list runners")
	}

	var pbRunners []*commonv1.Runner
	for _, runner := range runners {
		nsMatch, err := matchNamespace(runner.NamespaceSelectors, ns)
		if err != nil {
			log.Error().
				Err(err).
				Stringer("runner_id", runner.ID).
				Msg("failed to compile namespace selector")
			return nil, status.Error(codes.Internal, "failed to list runners")
		}
		if !nsMatch {
			continue
		}

		pbRun, err := types.ToProtoRunner(&runner)
		if err != nil {
			log.Error().
				Err(err).
				Stringer("runner_id", runner.ID).
				Msg("failed to format runner")
			return nil, status.Error(codes.Internal, "failed to format runner")
		}
		pbRunners = append(pbRunners, pbRun)
	}

	return &controlv1.ListRunnersResponse{
		Runners: pbRunners,
	}, nil
}

func (s ControlServer) scheduleRuns(ctx context.Context, dbtx db.DBTX, test *db.Test) ([]db.Run, error) {
	var runs []db.Run
	runParams, err := scheduler.RunsForTest(*test)
	if err != nil {
		return nil, fmt.Errorf("generating runs for test: %w", err)
	}

	for _, runParam := range runParams {
		run, err := s.dbQuerier.ScheduleRun(ctx, dbtx, runParam)
		if err != nil {
			return nil, fmt.Errorf("scheduling run for test %s: %w", test.ID, err)
		}
		runs = append(runs, run)
	}

	return runs, nil
}
