package temp

var Scheduler = `
#wing values are used to configure the kubernetes scheduler

# defaults from config and scheduler should be adequate

# Add your own!
KUBE_SCHEDULER_ARGS="--leader-elect=true --address=127.0.0.1"
`
