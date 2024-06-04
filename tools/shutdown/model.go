package shutdown

const (
	WorkerPoolPriority = iota
	RoutinePoolPriority
	LogPriority
	CronPriority
	GrpcServerPriority //grpc server 優先關閉避免request不斷進來
	HttpServerShutdownPriority
)

///关闭钩子
type ShutdownHook interface {
	///名字
	Name() string
	///优先级
	ShutdownPriority() int
	///应用程序退出前
	BeforeShutdown()
	/////应用程序退出后
	//AfterShutdown()
}

///排序
type ShutdownHooks []ShutdownHook

func (object ShutdownHooks) Len() int {
	return len(object)
}
func (object ShutdownHooks) Less(i, j int) bool {
	return object[i].ShutdownPriority() > object[j].ShutdownPriority()
}
func (object ShutdownHooks) Swap(i, j int) {
	object[i], object[j] = object[j], object[i]
}

type ImplementShutdown struct {
	Priority           int
	EventName          string
	BeforeShutdownFunc func()
}

func (imp *ImplementShutdown) Name() string {
	return imp.EventName
}

func (imp *ImplementShutdown) ShutdownPriority() int {
	return imp.Priority
}

func (imp *ImplementShutdown) BeforeShutdown() {
	imp.BeforeShutdownFunc()
}
