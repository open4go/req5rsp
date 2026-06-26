package rpc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/open4go/log"
	"github.com/open4go/req5rsp/req"
	"github.com/open4go/util9s"
	"github.com/spf13/viper"
)

const (
	// RPCUpdateUserLevel 更新会员层级
	updateUserLevel = "/rpc/level/"
)

// UpdateUserLevel 更新会员层级
func UpdateUserLevel(c *gin.Context, userId string, payload req.CondReq) error {
	if userId == "" {
		return fmt.Errorf("userId 不能为空")
	}

	baseURL := viper.GetString("rpc.member")
	if baseURL == "" {
		return fmt.Errorf("rpc.member 配置为空")
	}

	url := baseURL + updateUserLevel + userId

	ctx := c.Request.Context()
	logger := log.Log(ctx).WithField("user_id", userId).
		WithField("payload", payload).
		WithField("url", url)

	logger.Debug("开始调用 RPC 更新会员层级")

	// 使用 POST（推荐）或 PUT
	result, err := util9s.Post(ctx, url, payload) // ← 改成 Post
	if err != nil {
		logger.WithError(err).Error("调用 RPC 更新会员层级失败")
		return fmt.Errorf("更新会员层级失败: %w", err)
	}
	logger.WithField("result", result).Info("会员层级更新成功")
	return nil
}
