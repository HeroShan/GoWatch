    redis IP限流插件
    [timeMax] 访问时间最大值
    [limiter] 限制次数
    每timeMax秒访问limiter次，超过返回false
    每次记录访问的ip
## ==基准测试
    goos: windows
    goarch: amd64
    pkg: GoWatch/current_limiter
    BenchmarkSerlock-4           100          11360000 ns/op
    PASS
    ok      GoWatch/current_limiter 1.386s
