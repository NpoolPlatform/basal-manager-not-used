package api

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/basal-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/basal-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/basal-manager/pkg/tracer/api"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/basal-manager/pkg/db"
	"github.com/NpoolPlatform/basal-manager/pkg/db/ent"
	"github.com/NpoolPlatform/basal-manager/pkg/db/ent/api"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/basal/mgr/v1/api"

	"github.com/google/uuid"
)

func CreateSet(c *ent.APICreate, in *npool.APIReq) *ent.APICreate {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.Protocol != nil {
		c.SetProtocol(in.GetProtocol().String())
	}
	if in.ServiceName != nil {
		c.SetServiceName(in.GetServiceName())
	}
	if in.Method != nil {
		c.SetMethod(in.GetMethod().String())
	}
	if in.MethodName != nil {
		c.SetMethodName(in.GetMethodName())
	}
	if in.Path != nil {
		c.SetPath(in.GetPath())
	}
	if in.Exported != nil {
		c.SetExported(in.GetExported())
	}
	if in.PathPrefix != nil {
		c.SetPathPrefix(in.GetPathPrefix())
	}
	c.SetDomains(in.GetDomains())
	c.SetDepracated(false)
	return c
}

func Create(ctx context.Context, in *npool.APIReq) (*ent.API, error) {
	var info *ent.API
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.API.Create()
		c = CreateSet(c, in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.APIReq) ([]*ent.API, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBulk")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceMany(span, in)

	rows := []*ent.API{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.APICreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.API.Create(), info)
		}
		rows, err = tx.API.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func UpdateSet(info *ent.API, in *npool.APIReq) *ent.APIUpdateOne {
	stm := info.Update()

	if in.Depracated != nil {
		stm = stm.SetDepracated(in.GetDepracated())
	}

	return stm
}

func Update(ctx context.Context, in *npool.APIReq) (*ent.API, error) {
	var info *ent.API
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err = tx.API.Query().Where(api.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return fmt.Errorf("fail query api: %v", err)
		}

		stm := UpdateSet(info, in)

		info, err = stm.Save(_ctx)
		if err != nil {
			return fmt.Errorf("fail update api: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update api: %v", err)
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.API, error) {
	var info *ent.API
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Row")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.API.Query().Where(api.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.APIQuery, error) {
	stm := cli.API.Query()
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(api.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid api field")
		}
	}
	if conds.Protocol != nil {
		switch conds.GetProtocol().GetOp() {
		case cruder.EQ:
			stm.Where(api.Protocol(npool.Protocol(conds.GetProtocol().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid api field")
		}
	}
	if conds.Method != nil {
		switch conds.GetMethod().GetOp() {
		case cruder.EQ:
			stm.Where(api.Method(npool.Method(conds.GetMethod().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid api field")
		}
	}
	if conds.Exported != nil {
		switch conds.GetExported().GetOp() {
		case cruder.LIKE:
			stm.Where(api.Exported(conds.GetExported().GetValue()))
		default:
			return nil, fmt.Errorf("invalid api field")
		}
	}
	if conds.Depracated != nil {
		switch conds.GetDepracated().GetOp() {
		case cruder.LIKE:
			stm.Where(api.Depracated(conds.GetDepracated().GetValue()))
		default:
			return nil, fmt.Errorf("invalid api field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.API, int, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)
	span = commontracer.TraceOffsetLimit(span, offset, limit)

	rows := []*ent.API{}
	var total int
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(api.FieldUpdatedAt)).
			Limit(limit).
			All(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.API, error) {
	var info *ent.API
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Count(ctx context.Context, conds *npool.Conds) (uint32, error) {
	var err error
	var total int

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Count")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return uint32(total), nil
}

func Exist(ctx context.Context, id uuid.UUID) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Exist")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.API.Query().Where(api.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func Delete(ctx context.Context, id uuid.UUID) (*ent.API, error) {
	var info *ent.API
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Delete")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.API.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
