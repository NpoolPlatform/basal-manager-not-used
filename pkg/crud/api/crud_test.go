package api

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/basal-manager/pkg/db/ent"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	testinit "github.com/NpoolPlatform/basal-manager/pkg/testinit"
	valuedef "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/basal/mgr/v1/api"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = ent.API{
	ID:          uuid.New(),
	Protocol:    npool.Protocol_GRPC.String(),
	ServiceName: uuid.NewString(),
	Method:      npool.Method_POST.String(),
	MethodName:  uuid.NewString(),
	Path:        uuid.NewString(),
	Exported:    true,
	PathPrefix:  uuid.NewString(),
	Domains:     []string{"1", "2"},
}

var (
	id       = ret.ID.String()
	protocol = npool.Protocol_GRPC
	method   = npool.Method_POST

	req = npool.APIReq{
		ID:          &id,
		Protocol:    &protocol,
		ServiceName: &ret.ServiceName,
		Method:      &method,
		MethodName:  &ret.MethodName,
		Path:        &ret.Path,
		Exported:    &ret.Exported,
		PathPrefix:  &ret.PathPrefix,
		Domains:     ret.Domains,
	}
)

func create(t *testing.T) {
	info, err := Create(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), ret.String())
	}
}

func createBulk(t *testing.T) {
	entities := []*ent.API{
		{
			ID:          uuid.New(),
			Protocol:    npool.Protocol_GRPC.String(),
			ServiceName: uuid.NewString(),
			Method:      npool.Method_POST.String(),
			MethodName:  uuid.NewString(),
			Path:        uuid.NewString(),
			Exported:    true,
			PathPrefix:  uuid.NewString(),
			Domains:     []string{"1", "2"},
		},
		{
			ID:          uuid.New(),
			Protocol:    npool.Protocol_HTTP.String(),
			ServiceName: uuid.NewString(),
			Method:      npool.Method_GET.String(),
			MethodName:  uuid.NewString(),
			Path:        uuid.NewString(),
			Exported:    true,
			PathPrefix:  uuid.NewString(),
			Domains:     []string{"1", "2"},
		},
	}

	reqs := []*npool.APIReq{}
	for _, _ret := range entities {
		_id := _ret.ID.String()
		_protocol := npool.Protocol(npool.Protocol_value[_ret.Protocol])
		_method := npool.Method(npool.Method_value[_ret.Method])

		reqs = append(reqs, &npool.APIReq{
			ID:          &_id,
			Protocol:    &_protocol,
			ServiceName: &_ret.ServiceName,
			Method:      &_method,
			MethodName:  &_ret.MethodName,
			Path:        &_ret.Path,
			Exported:    &_ret.Exported,
			PathPrefix:  &_ret.PathPrefix,
			Domains:     _ret.Domains,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func update(t *testing.T) {
	depracated := true

	req.Depracated = &depracated
	ret.Depracated = depracated

	info, err := Update(context.Background(), &req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), ret.String())
	}
}

func row(t *testing.T) {
	info, err := Row(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), ret.String())
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		if assert.Equal(t, total, 1) {
			assert.Equal(t, infos[0].String(), ret.String())
		}
	}
}

func rowOnly(t *testing.T) {
	info, err := RowOnly(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), ret.String())
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, uint32(1))
	}
}

func exist(t *testing.T) {
	exist, err := Exist(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existConds(t *testing.T) {
	exist, err := ExistConds(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteA(t *testing.T) {
	info, err := Delete(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		ret.DeletedAt = info.DeletedAt
		assert.Equal(t, info.String(), ret.String())
	}
}

func TestAPI(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("createBulk", createBulk)
	t.Run("update", update)
	t.Run("row", row)
	t.Run("rows", rows)
	t.Run("rowOnly", rowOnly)
	t.Run("exist", exist)
	t.Run("existConds", existConds)
	t.Run("count", count)
	t.Run("delete", deleteA)
}
