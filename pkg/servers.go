package corrector

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"time"
)

// UserServer contains information regarding the running server of a user
type UserServer struct {
	server *exec.Cmd
	port   int
	name   string
}

// Run runs the user servers
func Run(users []string) []*UserServer {
	userServers := make([]*UserServer, 0, len(users))

	// TODO: Upgrade to goroutine
	for _, user := range users {
		// TODO: Make servers start on random ports
		port := 6069
		exec.Command("export", "SERVER_PORT="+strconv.Itoa(port))
		server := exec.Command("npm", "start", "--prefix", os.Getenv("HOME")+"/test-repos/"+user)
		server.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

		user := &UserServer{
			port:   port,
			server: server,
			name:   user,
		}

		userServers = append(userServers, user)

		if err := server.Start(); err != nil {
			fmt.Print(err)
		}

		sleepTime, _ := time.ParseDuration("2s")
		time.Sleep(sleepTime)
	}

	return userServers
}

// Kill terminates all the user servers
func Kill(users []*UserServer) {
	// TODO: Upgrade to goroutine
	for _, user := range users {
		pgid, err := syscall.Getpgid(user.server.Process.Pid)
		if err != nil {
			fmt.Print(err)
		}

		if err := syscall.Kill(-pgid, 9); err != nil {
			fmt.Print(err)
		}
	}
}
