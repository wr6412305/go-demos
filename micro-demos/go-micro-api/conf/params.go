package conf

// 关于环境的定义
const (
	ENV_PROD = "prod" // 生产环境
	ENV_PRE  = "pre"  // 预发环境
	ENV_TEST = "test" // 测试环境
	ENV_DEV  = "dev"  // 开发环境
)

// 服务名称
const (
	APP_CONF_PREFIX = "/micro/config/demo" // 配置文件前缀
	APP_API_GATEWAY = "demo.api.gateway"   // 用户网关app名称
	APP_SRV_USER    = "demo.srv.user"      // 用户服务app名称
)

// redis相关的key
const (
	REDIS_EMAIL_BACKTIMES = "demo:email:backtimes_"
)
