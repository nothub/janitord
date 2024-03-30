package systemd

// https://github.com/systemd/systemd/blob/main/src/basic/unit-def.c

type LoadState string

const (
	LoadStateStub       LoadState = "stub"
	LoadStateLoaded     LoadState = "loaded"
	LoadStateNotFound   LoadState = "not-found"
	LoadStateBadSetting LoadState = "bad-setting"
	LoadStateError      LoadState = "error"
	LoadStateMerged     LoadState = "merged"
	LoadStateMasked     LoadState = "masked"
)

type ActiveState string

const (
	ActiveStateActive       ActiveState = "active"
	ActiveStateReloading    ActiveState = "reloading"
	ActiveStateInactive     ActiveState = "inactive"
	ActiveStateFailed       ActiveState = "failed"
	ActiveStateActivating   ActiveState = "activating"
	ActiveStateDeactivating ActiveState = "deactivating"
	ActiveStateMaintenance  ActiveState = "maintenance"
)

type ServiceState string

const (
	ServiceStateDead                    ServiceState = "dead"
	ServiceStateCondition               ServiceState = "condition"
	ServiceStateStartPre                ServiceState = "start-pre"
	ServiceStateStart                   ServiceState = "start"
	ServiceStateStartPost               ServiceState = "start-post"
	ServiceStateRunning                 ServiceState = "running"
	ServiceStateExited                  ServiceState = "exited"
	ServiceStateReload                  ServiceState = "reload"
	ServiceStateReloadSignal            ServiceState = "reload-signal"
	ServiceStateReloadNotify            ServiceState = "reload-notify"
	ServiceStateSTOP                    ServiceState = "stop"
	ServiceStateStopWatchdog            ServiceState = "stop-watchdog"
	ServiceStateStopSigterm             ServiceState = "stop-sigterm"
	ServiceStateStopSigkill             ServiceState = "stop-sigkill"
	ServiceStateStopPost                ServiceState = "stop-post"
	ServiceStateFinalWatchdog           ServiceState = "final-watchdog"
	ServiceStateFinalSigterm            ServiceState = "final-sigterm"
	ServiceStateFinalSigkill            ServiceState = "final-sigkill"
	ServiceStateFailed                  ServiceState = "failed"
	ServiceStateDeadBeforeAutoRestart   ServiceState = "dead-before-auto-restart"
	ServiceStateFailedBeforeAutoRestart ServiceState = "failed-before-auto-restart"
	ServiceStateDeadResourcesPinned     ServiceState = "dead-resources-pinned"
	ServiceStateAutoRestart             ServiceState = "auto-restart"
	ServiceStateAutoRestartQueued       ServiceState = "auto-restart-queued"
	ServiceStateCleaning                ServiceState = "cleaning"
)
