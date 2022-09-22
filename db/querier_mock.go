// Code generated by MockGen. DO NOT EDIT.
// Source: querier.go

// Package db is a generated GoMock package.
package db

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// AppendLogsToRun mocks base method.
func (m *MockQuerier) AppendLogsToRun(ctx context.Context, db DBTX, arg AppendLogsToRunParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendLogsToRun", ctx, db, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppendLogsToRun indicates an expected call of AppendLogsToRun.
func (mr *MockQuerierMockRecorder) AppendLogsToRun(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendLogsToRun", reflect.TypeOf((*MockQuerier)(nil).AppendLogsToRun), ctx, db, arg)
}

// AssignRun mocks base method.
func (m *MockQuerier) AssignRun(ctx context.Context, db DBTX, arg AssignRunParams) (Run, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AssignRun", ctx, db, arg)
	ret0, _ := ret[0].(Run)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AssignRun indicates an expected call of AssignRun.
func (mr *MockQuerierMockRecorder) AssignRun(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AssignRun", reflect.TypeOf((*MockQuerier)(nil).AssignRun), ctx, db, arg)
}

// AuthAccessToken mocks base method.
func (m *MockQuerier) AuthAccessToken(ctx context.Context, db DBTX, tokenHash string) (AuthAccessTokenRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthAccessToken", ctx, db, tokenHash)
	ret0, _ := ret[0].(AuthAccessTokenRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthAccessToken indicates an expected call of AuthAccessToken.
func (mr *MockQuerierMockRecorder) AuthAccessToken(ctx, db, tokenHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthAccessToken", reflect.TypeOf((*MockQuerier)(nil).AuthAccessToken), ctx, db, tokenHash)
}

// DeleteRunsForTest mocks base method.
func (m *MockQuerier) DeleteRunsForTest(ctx context.Context, db DBTX, testID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRunsForTest", ctx, db, testID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRunsForTest indicates an expected call of DeleteRunsForTest.
func (mr *MockQuerierMockRecorder) DeleteRunsForTest(ctx, db, testID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRunsForTest", reflect.TypeOf((*MockQuerier)(nil).DeleteRunsForTest), ctx, db, testID)
}

// DeleteTest mocks base method.
func (m *MockQuerier) DeleteTest(ctx context.Context, db DBTX, arg DeleteTestParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTest", ctx, db, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTest indicates an expected call of DeleteTest.
func (mr *MockQuerierMockRecorder) DeleteTest(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTest", reflect.TypeOf((*MockQuerier)(nil).DeleteTest), ctx, db, arg)
}

// GetAccessToken mocks base method.
func (m *MockQuerier) GetAccessToken(ctx context.Context, db DBTX, id uuid.UUID) (GetAccessTokenRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessToken", ctx, db, id)
	ret0, _ := ret[0].(GetAccessTokenRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessToken indicates an expected call of GetAccessToken.
func (mr *MockQuerierMockRecorder) GetAccessToken(ctx, db, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessToken", reflect.TypeOf((*MockQuerier)(nil).GetAccessToken), ctx, db, id)
}

// GetRun mocks base method.
func (m *MockQuerier) GetRun(ctx context.Context, db DBTX, arg GetRunParams) (Run, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRun", ctx, db, arg)
	ret0, _ := ret[0].(Run)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRun indicates an expected call of GetRun.
func (mr *MockQuerierMockRecorder) GetRun(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRun", reflect.TypeOf((*MockQuerier)(nil).GetRun), ctx, db, arg)
}

// GetRunner mocks base method.
func (m *MockQuerier) GetRunner(ctx context.Context, db DBTX, id uuid.UUID) (Runner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunner", ctx, db, id)
	ret0, _ := ret[0].(Runner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRunner indicates an expected call of GetRunner.
func (mr *MockQuerierMockRecorder) GetRunner(ctx, db, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunner", reflect.TypeOf((*MockQuerier)(nil).GetRunner), ctx, db, id)
}

// GetTest mocks base method.
func (m *MockQuerier) GetTest(ctx context.Context, db DBTX, arg GetTestParams) (Test, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTest", ctx, db, arg)
	ret0, _ := ret[0].(Test)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTest indicates an expected call of GetTest.
func (mr *MockQuerierMockRecorder) GetTest(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTest", reflect.TypeOf((*MockQuerier)(nil).GetTest), ctx, db, arg)
}

// IssueAccessToken mocks base method.
func (m *MockQuerier) IssueAccessToken(ctx context.Context, db DBTX, arg IssueAccessTokenParams) (IssueAccessTokenRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IssueAccessToken", ctx, db, arg)
	ret0, _ := ret[0].(IssueAccessTokenRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IssueAccessToken indicates an expected call of IssueAccessToken.
func (mr *MockQuerierMockRecorder) IssueAccessToken(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IssueAccessToken", reflect.TypeOf((*MockQuerier)(nil).IssueAccessToken), ctx, db, arg)
}

// ListAccessTokens mocks base method.
func (m *MockQuerier) ListAccessTokens(ctx context.Context, db DBTX, arg ListAccessTokensParams) ([]ListAccessTokensRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccessTokens", ctx, db, arg)
	ret0, _ := ret[0].([]ListAccessTokensRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccessTokens indicates an expected call of ListAccessTokens.
func (mr *MockQuerierMockRecorder) ListAccessTokens(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccessTokens", reflect.TypeOf((*MockQuerier)(nil).ListAccessTokens), ctx, db, arg)
}

// ListAllNamespaces mocks base method.
func (m *MockQuerier) ListAllNamespaces(ctx context.Context, db DBTX) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllNamespaces", ctx, db)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllNamespaces indicates an expected call of ListAllNamespaces.
func (mr *MockQuerierMockRecorder) ListAllNamespaces(ctx, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllNamespaces", reflect.TypeOf((*MockQuerier)(nil).ListAllNamespaces), ctx, db)
}

// ListPendingRuns mocks base method.
func (m *MockQuerier) ListPendingRuns(ctx context.Context, db DBTX) ([]ListPendingRunsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPendingRuns", ctx, db)
	ret0, _ := ret[0].([]ListPendingRunsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPendingRuns indicates an expected call of ListPendingRuns.
func (mr *MockQuerierMockRecorder) ListPendingRuns(ctx, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPendingRuns", reflect.TypeOf((*MockQuerier)(nil).ListPendingRuns), ctx, db)
}

// ListRunners mocks base method.
func (m *MockQuerier) ListRunners(ctx context.Context, db DBTX) ([]Runner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRunners", ctx, db)
	ret0, _ := ret[0].([]Runner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRunners indicates an expected call of ListRunners.
func (mr *MockQuerierMockRecorder) ListRunners(ctx, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRunners", reflect.TypeOf((*MockQuerier)(nil).ListRunners), ctx, db)
}

// ListRuns mocks base method.
func (m *MockQuerier) ListRuns(ctx context.Context, db DBTX, namespace string) ([]Run, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRuns", ctx, db, namespace)
	ret0, _ := ret[0].([]Run)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRuns indicates an expected call of ListRuns.
func (mr *MockQuerierMockRecorder) ListRuns(ctx, db, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRuns", reflect.TypeOf((*MockQuerier)(nil).ListRuns), ctx, db, namespace)
}

// ListTests mocks base method.
func (m *MockQuerier) ListTests(ctx context.Context, db DBTX, namespace string) ([]Test, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTests", ctx, db, namespace)
	ret0, _ := ret[0].([]Test)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTests indicates an expected call of ListTests.
func (mr *MockQuerierMockRecorder) ListTests(ctx, db, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTests", reflect.TypeOf((*MockQuerier)(nil).ListTests), ctx, db, namespace)
}

// ListTestsToSchedule mocks base method.
func (m *MockQuerier) ListTestsToSchedule(ctx context.Context, db DBTX) ([]Test, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTestsToSchedule", ctx, db)
	ret0, _ := ret[0].([]Test)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTestsToSchedule indicates an expected call of ListTestsToSchedule.
func (mr *MockQuerierMockRecorder) ListTestsToSchedule(ctx, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTestsToSchedule", reflect.TypeOf((*MockQuerier)(nil).ListTestsToSchedule), ctx, db)
}

// QueryRunners mocks base method.
func (m *MockQuerier) QueryRunners(ctx context.Context, db DBTX, arg QueryRunnersParams) ([]Runner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryRunners", ctx, db, arg)
	ret0, _ := ret[0].([]Runner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryRunners indicates an expected call of QueryRunners.
func (mr *MockQuerierMockRecorder) QueryRunners(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRunners", reflect.TypeOf((*MockQuerier)(nil).QueryRunners), ctx, db, arg)
}

// QueryRuns mocks base method.
func (m *MockQuerier) QueryRuns(ctx context.Context, db DBTX, arg QueryRunsParams) ([]Run, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryRuns", ctx, db, arg)
	ret0, _ := ret[0].([]Run)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryRuns indicates an expected call of QueryRuns.
func (mr *MockQuerierMockRecorder) QueryRuns(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryRuns", reflect.TypeOf((*MockQuerier)(nil).QueryRuns), ctx, db, arg)
}

// QueryTests mocks base method.
func (m *MockQuerier) QueryTests(ctx context.Context, db DBTX, arg QueryTestsParams) ([]Test, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryTests", ctx, db, arg)
	ret0, _ := ret[0].([]Test)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryTests indicates an expected call of QueryTests.
func (mr *MockQuerierMockRecorder) QueryTests(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryTests", reflect.TypeOf((*MockQuerier)(nil).QueryTests), ctx, db, arg)
}

// RegisterRunner mocks base method.
func (m *MockQuerier) RegisterRunner(ctx context.Context, db DBTX, arg RegisterRunnerParams) (Runner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterRunner", ctx, db, arg)
	ret0, _ := ret[0].(Runner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterRunner indicates an expected call of RegisterRunner.
func (mr *MockQuerierMockRecorder) RegisterRunner(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRunner", reflect.TypeOf((*MockQuerier)(nil).RegisterRunner), ctx, db, arg)
}

// RegisterTest mocks base method.
func (m *MockQuerier) RegisterTest(ctx context.Context, db DBTX, arg RegisterTestParams) (Test, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterTest", ctx, db, arg)
	ret0, _ := ret[0].(Test)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterTest indicates an expected call of RegisterTest.
func (mr *MockQuerierMockRecorder) RegisterTest(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterTest", reflect.TypeOf((*MockQuerier)(nil).RegisterTest), ctx, db, arg)
}

// ResetOrphanedRuns mocks base method.
func (m *MockQuerier) ResetOrphanedRuns(ctx context.Context, db DBTX, before time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetOrphanedRuns", ctx, db, before)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetOrphanedRuns indicates an expected call of ResetOrphanedRuns.
func (mr *MockQuerierMockRecorder) ResetOrphanedRuns(ctx, db, before interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetOrphanedRuns", reflect.TypeOf((*MockQuerier)(nil).ResetOrphanedRuns), ctx, db, before)
}

// RevokeAccessToken mocks base method.
func (m *MockQuerier) RevokeAccessToken(ctx context.Context, db DBTX, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeAccessToken", ctx, db, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeAccessToken indicates an expected call of RevokeAccessToken.
func (mr *MockQuerierMockRecorder) RevokeAccessToken(ctx, db, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeAccessToken", reflect.TypeOf((*MockQuerier)(nil).RevokeAccessToken), ctx, db, id)
}

// RunSummariesForRunner mocks base method.
func (m *MockQuerier) RunSummariesForRunner(ctx context.Context, db DBTX, arg RunSummariesForRunnerParams) ([]RunSummariesForRunnerRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunSummariesForRunner", ctx, db, arg)
	ret0, _ := ret[0].([]RunSummariesForRunnerRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunSummariesForRunner indicates an expected call of RunSummariesForRunner.
func (mr *MockQuerierMockRecorder) RunSummariesForRunner(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunSummariesForRunner", reflect.TypeOf((*MockQuerier)(nil).RunSummariesForRunner), ctx, db, arg)
}

// RunSummariesForTest mocks base method.
func (m *MockQuerier) RunSummariesForTest(ctx context.Context, db DBTX, arg RunSummariesForTestParams) ([]RunSummariesForTestRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunSummariesForTest", ctx, db, arg)
	ret0, _ := ret[0].([]RunSummariesForTestRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RunSummariesForTest indicates an expected call of RunSummariesForTest.
func (mr *MockQuerierMockRecorder) RunSummariesForTest(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunSummariesForTest", reflect.TypeOf((*MockQuerier)(nil).RunSummariesForTest), ctx, db, arg)
}

// ScheduleRun mocks base method.
func (m *MockQuerier) ScheduleRun(ctx context.Context, db DBTX, arg ScheduleRunParams) (Run, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScheduleRun", ctx, db, arg)
	ret0, _ := ret[0].(Run)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScheduleRun indicates an expected call of ScheduleRun.
func (mr *MockQuerierMockRecorder) ScheduleRun(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleRun", reflect.TypeOf((*MockQuerier)(nil).ScheduleRun), ctx, db, arg)
}

// SummarizeRunsBreakdownResult mocks base method.
func (m *MockQuerier) SummarizeRunsBreakdownResult(ctx context.Context, db DBTX, arg SummarizeRunsBreakdownResultParams) ([]SummarizeRunsBreakdownResultRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SummarizeRunsBreakdownResult", ctx, db, arg)
	ret0, _ := ret[0].([]SummarizeRunsBreakdownResultRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SummarizeRunsBreakdownResult indicates an expected call of SummarizeRunsBreakdownResult.
func (mr *MockQuerierMockRecorder) SummarizeRunsBreakdownResult(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SummarizeRunsBreakdownResult", reflect.TypeOf((*MockQuerier)(nil).SummarizeRunsBreakdownResult), ctx, db, arg)
}

// SummarizeRunsBreakdownTest mocks base method.
func (m *MockQuerier) SummarizeRunsBreakdownTest(ctx context.Context, db DBTX, arg SummarizeRunsBreakdownTestParams) ([]SummarizeRunsBreakdownTestRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SummarizeRunsBreakdownTest", ctx, db, arg)
	ret0, _ := ret[0].([]SummarizeRunsBreakdownTestRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SummarizeRunsBreakdownTest indicates an expected call of SummarizeRunsBreakdownTest.
func (mr *MockQuerierMockRecorder) SummarizeRunsBreakdownTest(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SummarizeRunsBreakdownTest", reflect.TypeOf((*MockQuerier)(nil).SummarizeRunsBreakdownTest), ctx, db, arg)
}

// TimeoutRuns mocks base method.
func (m *MockQuerier) TimeoutRuns(ctx context.Context, db DBTX, arg TimeoutRunsParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TimeoutRuns", ctx, db, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// TimeoutRuns indicates an expected call of TimeoutRuns.
func (mr *MockQuerierMockRecorder) TimeoutRuns(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimeoutRuns", reflect.TypeOf((*MockQuerier)(nil).TimeoutRuns), ctx, db, arg)
}

// UpdateResultData mocks base method.
func (m *MockQuerier) UpdateResultData(ctx context.Context, db DBTX, arg UpdateResultDataParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateResultData", ctx, db, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateResultData indicates an expected call of UpdateResultData.
func (mr *MockQuerierMockRecorder) UpdateResultData(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateResultData", reflect.TypeOf((*MockQuerier)(nil).UpdateResultData), ctx, db, arg)
}

// UpdateRun mocks base method.
func (m *MockQuerier) UpdateRun(ctx context.Context, db DBTX, arg UpdateRunParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRun", ctx, db, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRun indicates an expected call of UpdateRun.
func (mr *MockQuerierMockRecorder) UpdateRun(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRun", reflect.TypeOf((*MockQuerier)(nil).UpdateRun), ctx, db, arg)
}

// UpdateRunnerHeartbeat mocks base method.
func (m *MockQuerier) UpdateRunnerHeartbeat(ctx context.Context, db DBTX, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRunnerHeartbeat", ctx, db, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRunnerHeartbeat indicates an expected call of UpdateRunnerHeartbeat.
func (mr *MockQuerierMockRecorder) UpdateRunnerHeartbeat(ctx, db, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRunnerHeartbeat", reflect.TypeOf((*MockQuerier)(nil).UpdateRunnerHeartbeat), ctx, db, id)
}

// UpdateTest mocks base method.
func (m *MockQuerier) UpdateTest(ctx context.Context, db DBTX, arg UpdateTestParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTest", ctx, db, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTest indicates an expected call of UpdateTest.
func (mr *MockQuerierMockRecorder) UpdateTest(ctx, db, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTest", reflect.TypeOf((*MockQuerier)(nil).UpdateTest), ctx, db, arg)
}
