package collectors

import (
	"context"

	"github.com/lovoo/goka"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type CollectorUsecase[T protoreflect.ProtoMessage] func(context.Context, []*T, *T) []*T

type CallbackUsecasInterface[T protoreflect.ProtoMessage] interface {
	Run(context.Context, []T, interface{}) []T
}

func callback[T protoreflect.ProtoMessage](ucase CallbackUsecasInterface[T]) goka.ProcessCallback {
	return func(ctx goka.Context, msg interface{}) {
		var messageList []T
		if v := ctx.Value(); v != nil {
			messageList = v.([]T)
		}

		messageList = ucase.Run(ctx.Context(), messageList, msg)
		ctx.SetValue(messageList)
	}
}
