package corrector

import (
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"time"
)

// UserServer contains information regarding the running server of a user
type UserServer struct {
	server *exec.Cmd
	port   string
	name   string
}

// Run runs the user servers
func Run(users []string) ([]*UserServer, error) {
	userServers := make([]*UserServer, 0, len(users))

	// TODO: Upgrade to goroutine
	for _, user := range users {
		port, err := getFreePort()
		if err != nil {
			return nil, err
		}

		os.Setenv("PORT", strconv.Itoa(port))
		server := exec.Command("npm", "start", "--prefix", os.Getenv("HOME")+"/test-repos/"+user)
		server.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

		userServer := &UserServer{
			port:   strconv.Itoa(port),
			server: server,
			name:   user,
		}

		userServers = append(userServers, userServer)

		if err := server.Start(); err != nil {
			return nil, err
		}

		log.Printf("Waiting for %v's server to start on port %v\n", user, userServer.port)
		sleepTime, _ := time.ParseDuration("2s")
		time.Sleep(sleepTime)
	}

	return userServers, nil
}

// Kill terminates all the user servers
func Kill(users []*UserServer) error {
	// TODO: Upgrade to goroutine
	for _, user := range users {
		pgid, err := syscall.Getpgid(user.server.Process.Pid)
		if err != nil {
			return err
		}

		if err := syscall.Kill(-pgid, 9); err != nil {
			return err
		}
	}

	return nil
}

func getFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer listener.Close()

	return listener.Addr().(*net.TCPAddr).Port, nil
}
