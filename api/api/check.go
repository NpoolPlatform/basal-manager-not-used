package api

import (
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/basal/mgr/v1/api"

	"github.com/google/uuid"
)

func validate(info *npool.APIReq) error { //nolint
	if info.ID != nil {
		if _, err := uuid.Parse(info.GetID()); err != nil {
			logger.Sugar().Errorw("validate", "ID", info.ID)
			return fmt.Errorf("id is empty")
		}
	}

	switch info.GetProtocol() {
	case npool.Protocol_GRPC:
	case npool.Protocol_HTTP:
	default:
		logger.Sugar().Errorw("validate", "Protocol", info.Protocol)
		return fmt.Errorf("protocol is invalid")
	}

	if info.ServiceName != nil && info.GetServiceName() == "" {
		logger.Sugar().Errorw("validate", "ServiceName", info.ServiceName)
		return fmt.Errorf("servicename is invalid")
	}

	switch info.GetMethod() {
	case npool.Method_GET:
	case npool.Method_POST:
	case npool.Method_STREAM:
	default:
		logger.Sugar().Errorw("validate", "Method", info.Method)
		return fmt.Errorf("method is invalid")
	}

	if info.MethodName != nil && info.GetMethodName() == "" {
		logger.Sugar().Errorw("validate", "MethodName", info.MethodName)
		return fmt.Errorf("methodname is invalid")
	}

	if info.Path != nil && info.GetPath() == "" {
		logger.Sugar().Errorw("validate", "Path", info.Path)
		return fmt.Errorf("path is invalid")
	}

	if info.PathPrefix != nil && info.GetPathPrefix() == "" {
		logger.Sugar().Errorw("validate", "PathPrefix", info.PathPrefix)
		return fmt.Errorf("pathprefix is invalid")
	}

	for _, domain := range info.GetDomains() {
		if domain == "" {
			logger.Sugar().Errorw("validate", "Domains", info.Domains)
			return fmt.Errorf("domains is invalid")
		}
	}

	return nil
}

func validateMany(infos []*npool.APIReq) error {
	keys := map[string]struct{}{}

	for _, info := range infos {
		if err := validate(info); err != nil {
			return err
		}

		key := fmt.Sprintf(
			"%v:%v:%v:%v",
			info.GetServiceName(),
			info.GetProtocol(),
			info.GetMethod(),
			info.GetPath())

		if _, ok := keys[key]; ok {
			return fmt.Errorf("infos is invalid")
		}

		keys[key] = struct{}{}
	}

	return nil
}

func Validate(info *npool.APIReq) error {
	return validate(info)
}

func ValidateMany(infos []*npool.APIReq) error {
	return validateMany(infos)
}
