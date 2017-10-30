package restful_skel

import (
	"os"
	"os/user"
	"time"

	"restful_skel/config"
	log "restful_skel/log"
)

func init() {
}

func Run() int {
	log.Logging()

	if u, _ := user.Current(); u.Gid == "0" {
		log.Info("ROOT permission granted.")
	}

	RESTfulAPIServe()

	for {
		time.Sleep(time.Second)
	}
	finish(config.EXITNORMAL)
	return config.EXITNORMAL
}

func finish(returnCode int) {
	log.Infof("Terminates: %d", returnCode)
	os.Exit(returnCode)
}
