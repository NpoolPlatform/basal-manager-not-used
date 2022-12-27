//nolint:dupl
package api

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/basal/mgr/v1/api"

	constant "github.com/NpoolPlatform/basal-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get api connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateAPI(ctx context.Context, in *npool.APIReq) (*npool.API, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateAPI(ctx, &npool.CreateAPIRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create api: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create api: %v", err)
	}
	return info.(*npool.API), nil
}

func CreateAPIs(ctx context.Context, in []*npool.APIReq) ([]*npool.API, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateAPIs(ctx, &npool.CreateAPIsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create apis: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create apis: %v", err)
	}
	return infos.([]*npool.API), nil
}

func UpdateAPI(ctx context.Context, in *npool.APIReq) (*npool.API, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateAPI(ctx, &npool.UpdateAPIRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail update api: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update api: %v", err)
	}
	return info.(*npool.API), nil
}

func GetAPI(ctx context.Context, id string) (*npool.API, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetAPI(ctx, &npool.GetAPIRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get api: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get api: %v", err)
	}
	return info.(*npool.API), nil
}

func GetAPIOnly(ctx context.Context, conds *npool.Conds) (*npool.API, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetAPIOnly(ctx, &npool.GetAPIOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get api: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get api: %v", err)
	}
	return info.(*npool.API), nil
}

func GetAPIs(ctx context.Context, conds *npool.Conds, limit, offset int32) ([]*npool.API, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetAPIs(ctx, &npool.GetAPIsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get apis: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get apis: %v", err)
	}
	return infos.([]*npool.API), total, nil
}

func ExistAPI(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistAPI(ctx, &npool.ExistAPIRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get api: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get api: %v", err)
	}
	return infos.(bool), nil
}

func ExistAPIConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistAPIConds(ctx, &npool.ExistAPICondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get api: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get api: %v", err)
	}
	return infos.(bool), nil
}

func CountAPIs(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountAPIs(ctx, &npool.CountAPIsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count api: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count api: %v", err)
	}
	return infos.(uint32), nil
}
