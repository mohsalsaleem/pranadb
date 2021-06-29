package pull

import (
	"errors"
	"fmt"
	"github.com/pingcap/tidb/infoschema"
	"github.com/pingcap/tidb/planner/core"
	"github.com/squareup/pranadb/common"
	"github.com/squareup/pranadb/pull/exec"
)

func (p *PullEngine) buildPullQueryExecution(schema *common.Schema, is infoschema.InfoSchema, query string) (queryDAG exec.PullExecutor, err error) {
	// Build the physical plan
	physicalPlan, err := p.planner.QueryToPlan(query, is, true)
	if err != nil {
		return nil, err
	}
	// Build initial dag from the plan
	dag, err := p.buildPullDAG(physicalPlan, schema)

	// TODO TODO

	// We need to create out own remote call executor which corresponds to the TableReader
	// Based on the ranges in the table reader we can figure out what remote nodes we need to call
	// In the case of a point lookup it will be a single node, otherwise it will be all nodes.
	// In the first execution of getRows on the remote call executor we make a gRPC call to the
	// node(s) and in the case of a non point lookup we pass the serialized dag fragment to the node
	// when that is received, it is instantiated and a reference stored in a "PullQueryManager" and
	// a unique id returned.
	// The next call to getRows just needs to pass the unique id of the query. When the query is complete
	// the remote node can automatically unregister the query. Queries should also be unregistered after
	// a no activity timeout.

	return dag, nil
}

// TODO maybe combine with similar logic in buildPushDAG?
func (p *PullEngine) buildPullDAG(plan core.PhysicalPlan, schema *common.Schema) (exec.PullExecutor, error) {
	cols := plan.Schema().Columns
	colTypes := make([]common.ColumnType, 0, len(cols))
	colNames := make([]string, 0, len(cols))
	for _, col := range cols {
		colType := col.GetType()
		pranaType, err := common.ConvertTiDBTypeToPranaType(colType)
		if err != nil {
			return nil, err
		}
		colTypes = append(colTypes, pranaType)
		colNames = append(colNames, col.OrigName)
	}
	var executor exec.PullExecutor
	var err error
	switch plan.(type) {
	case *core.PhysicalProjection:
		physProj := plan.(*core.PhysicalProjection)
		var exprs []*common.Expression
		for _, expr := range physProj.Exprs {
			exprs = append(exprs, common.NewExpression(expr))
		}
		executor, err = exec.NewPullProjection(colNames, colTypes, exprs)
		if err != nil {
			return nil, err
		}
	case *core.PhysicalSelection:
		physSel := plan.(*core.PhysicalSelection)
		var exprs []*common.Expression
		for _, expr := range physSel.Conditions {
			exprs = append(exprs, common.NewExpression(expr))
		}
		executor, err = exec.NewPullSelect(colNames, colTypes, exprs)
		if err != nil {
			return nil, err
		}
	case *core.PhysicalHashAgg:
		physAgg := plan.(*core.PhysicalHashAgg)
		// TODO
		println("%v", physAgg)

	case *core.PhysicalTableReader:
		physTabReader := plan.(*core.PhysicalTableReader)
		if len(physTabReader.TablePlans) != 1 {
			panic("expected one t plan")
		}
		tabPlan := physTabReader.TablePlans[0]
		physTableScan, ok := tabPlan.(*core.PhysicalTableScan)
		if !ok {
			return nil, errors.New("expected PhysicalTableScan")
		}
		tableName := physTableScan.Table.Name.L
		var t *common.TableInfo
		mv, ok := schema.Mvs[tableName]
		if !ok {
			source, ok := schema.Sources[tableName]
			if !ok {
				return nil, fmt.Errorf("unknown source or materialized view %s", tableName)
			}
			t = source.TableInfo
		} else {
			t = mv.TableInfo
		}
		executor, err = exec.NewPullTableScan(colTypes, t)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unexpected plan type %T", plan)
	}

	var childExecutors []exec.PullExecutor
	for _, child := range plan.Children() {
		childExecutor, err := p.buildPullDAG(child, schema)
		if err != nil {
			return nil, err
		}
		if childExecutor != nil {
			childExecutors = append(childExecutors, childExecutor)
		}
	}
	exec.ConnectPullExecutors(childExecutors, executor)

	return executor, nil
}
