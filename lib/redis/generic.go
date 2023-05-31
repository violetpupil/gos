package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// Expire 设置过期时间
// 键不存在或者选项条件不满足，返回 false 代表设置失败
// The timeout will only be cleared by commands that delete or overwrite the contents of the key, including DEL, SET, GETSET and all the *STORE commands.
// This means that all the operations that conceptually alter the value stored at the key without replacing it with a new one will leave the timeout untouched.
//
// 支持选项如下
// nx Set expiry only when the key has no expiry
// xx Set expiry only when the key has an existing expiry
// gt Set expiry only when the new expiry is greater than current one
// lt Set expiry only when the new expiry is less than current one
//
// https://redis.io/commands/expire/
func Expire(ctx context.Context, key string, e time.Duration, o string) (bool, error) {
	var f func(context.Context, string, time.Duration) *redis.BoolCmd
	switch o {
	case "nx":
		f = client.ExpireNX
	case "xx":
		f = client.ExpireXX
	case "gt":
		f = client.ExpireGT
	case "lt":
		f = client.ExpireLT
	default:
		f = client.Expire
	}
	return f(ctx, key, e).Result()
}
