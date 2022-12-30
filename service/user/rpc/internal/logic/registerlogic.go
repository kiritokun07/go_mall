package logic

import (
	"context"

	"go_mall/common/cryptx"
	"go_mall/service/user/model"
	"go_mall/service/user/rpc/internal/svc"
	"go_mall/service/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// 判断手机号是否已经注册
	_, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err == nil {
		return nil, status.Error(100, "该用户已存在")
	}
	if err != model.ErrNotFound {
		return nil, status.Error(500, err.Error())
	}
	newUser := model.User{
		Name:     in.Name,
		Gender:   uint64(in.Gender),
		Mobile:   in.Mobile,
		Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
	}
	res, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	newUserId, err := res.LastInsertId()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	return &user.RegisterResp{
		Id:     newUserId,
		Name:   newUser.Name,
		Gender: int64(newUser.Gender),
		Mobile: newUser.Mobile,
	}, nil
}
