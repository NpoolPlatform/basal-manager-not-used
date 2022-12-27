package api

import (
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/basal/mgr/v1/api"
)

func trace(span trace1.Span, in *npool.APIReq, index int) trace1.Span {
	return span
}

func Trace(span trace1.Span, in *npool.APIReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	return span
}

func TraceMany(span trace1.Span, infos []*npool.APIReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
