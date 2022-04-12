// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.2.0

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type LogicHTTPServer interface {
	GroupRecall(context.Context, *GroupRecallRequest) (*RecallReplay, error)
	GroupSend(context.Context, *GroupSendRequest) (*SendReplay, error)
	GroupSendMention(context.Context, *GroupSendMentionRequest) (*SendReplay, error)
	RoomBroadcast(context.Context, *GroupSendRequest) (*SendReplay, error)
	RoomSend(context.Context, *GroupSendRequest) (*SendReplay, error)
	SingleRecall(context.Context, *SingleRecallRequest) (*RecallReplay, error)
	SingleSend(context.Context, *SingleSendRequest) (*SendReplay, error)
}

func RegisterLogicHTTPServer(s *http.Server, srv LogicHTTPServer) {
	r := s.Route("/")
	r.POST("/single/send", _Logic_SingleSend0_HTTP_Handler(srv))
	r.POST("/single/recall", _Logic_SingleRecall0_HTTP_Handler(srv))
	r.POST("/group/send", _Logic_GroupSend0_HTTP_Handler(srv))
	r.POST("/group/send_mention", _Logic_GroupSendMention0_HTTP_Handler(srv))
	r.POST("/group/recall", _Logic_GroupRecall0_HTTP_Handler(srv))
	r.POST("/room/send", _Logic_RoomSend0_HTTP_Handler(srv))
	r.POST("/room/broadcast", _Logic_RoomBroadcast0_HTTP_Handler(srv))
}

func _Logic_SingleSend0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SingleSendRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.v1.Logic/SingleSend")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SingleSend(ctx, req.(*SingleSendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendReplay)
		return ctx.Result(200, reply)
	}
}

func _Logic_SingleRecall0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SingleRecallRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.v1.Logic/SingleRecall")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SingleRecall(ctx, req.(*SingleRecallRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RecallReplay)
		return ctx.Result(200, reply)
	}
}

func _Logic_GroupSend0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GroupSendRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.v1.Logic/GroupSend")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GroupSend(ctx, req.(*GroupSendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendReplay)
		return ctx.Result(200, reply)
	}
}

func _Logic_GroupSendMention0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GroupSendMentionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.v1.Logic/GroupSendMention")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GroupSendMention(ctx, req.(*GroupSendMentionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendReplay)
		return ctx.Result(200, reply)
	}
}

func _Logic_GroupRecall0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GroupRecallRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.v1.Logic/GroupRecall")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GroupRecall(ctx, req.(*GroupRecallRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RecallReplay)
		return ctx.Result(200, reply)
	}
}

func _Logic_RoomSend0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GroupSendRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.v1.Logic/RoomSend")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RoomSend(ctx, req.(*GroupSendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendReplay)
		return ctx.Result(200, reply)
	}
}

func _Logic_RoomBroadcast0_HTTP_Handler(srv LogicHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GroupSendRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.logic.v1.Logic/RoomBroadcast")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RoomBroadcast(ctx, req.(*GroupSendRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SendReplay)
		return ctx.Result(200, reply)
	}
}

type LogicHTTPClient interface {
	GroupRecall(ctx context.Context, req *GroupRecallRequest, opts ...http.CallOption) (rsp *RecallReplay, err error)
	GroupSend(ctx context.Context, req *GroupSendRequest, opts ...http.CallOption) (rsp *SendReplay, err error)
	GroupSendMention(ctx context.Context, req *GroupSendMentionRequest, opts ...http.CallOption) (rsp *SendReplay, err error)
	RoomBroadcast(ctx context.Context, req *GroupSendRequest, opts ...http.CallOption) (rsp *SendReplay, err error)
	RoomSend(ctx context.Context, req *GroupSendRequest, opts ...http.CallOption) (rsp *SendReplay, err error)
	SingleRecall(ctx context.Context, req *SingleRecallRequest, opts ...http.CallOption) (rsp *RecallReplay, err error)
	SingleSend(ctx context.Context, req *SingleSendRequest, opts ...http.CallOption) (rsp *SendReplay, err error)
}

type LogicHTTPClientImpl struct {
	cc *http.Client
}

func NewLogicHTTPClient(client *http.Client) LogicHTTPClient {
	return &LogicHTTPClientImpl{client}
}

func (c *LogicHTTPClientImpl) GroupRecall(ctx context.Context, in *GroupRecallRequest, opts ...http.CallOption) (*RecallReplay, error) {
	var out RecallReplay
	pattern := "/group/recall"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.logic.v1.Logic/GroupRecall"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) GroupSend(ctx context.Context, in *GroupSendRequest, opts ...http.CallOption) (*SendReplay, error) {
	var out SendReplay
	pattern := "/group/send"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.logic.v1.Logic/GroupSend"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) GroupSendMention(ctx context.Context, in *GroupSendMentionRequest, opts ...http.CallOption) (*SendReplay, error) {
	var out SendReplay
	pattern := "/group/send_mention"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.logic.v1.Logic/GroupSendMention"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) RoomBroadcast(ctx context.Context, in *GroupSendRequest, opts ...http.CallOption) (*SendReplay, error) {
	var out SendReplay
	pattern := "/room/broadcast"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.logic.v1.Logic/RoomBroadcast"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) RoomSend(ctx context.Context, in *GroupSendRequest, opts ...http.CallOption) (*SendReplay, error) {
	var out SendReplay
	pattern := "/room/send"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.logic.v1.Logic/RoomSend"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) SingleRecall(ctx context.Context, in *SingleRecallRequest, opts ...http.CallOption) (*RecallReplay, error) {
	var out RecallReplay
	pattern := "/single/recall"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.logic.v1.Logic/SingleRecall"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *LogicHTTPClientImpl) SingleSend(ctx context.Context, in *SingleSendRequest, opts ...http.CallOption) (*SendReplay, error) {
	var out SendReplay
	pattern := "/single/send"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.logic.v1.Logic/SingleSend"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
