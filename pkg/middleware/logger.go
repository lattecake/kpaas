package middleware

//type loggingMiddleware struct {
//	logger log.Logger
//}
//
//func LoggerMiddleware() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		defer func(begin time.Time) {
//			logger.Info("took", time.Since(begin))
//		}(time.Now())
//
//		ctx.Next()
//	}
//}
