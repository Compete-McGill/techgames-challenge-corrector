package corrector

import (
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"sync"
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
func Run(users []string) []*UserServer {
	userServers := make([]*UserServer, 0, len(users))
	var wg sync.WaitGroup

	for _, user := range users {
		wg.Add(1)
		go runHelper(user, &userServers, &wg)
	}

	wg.Wait()

	return userServers
}

func runHelper(user string, userServers *([]*UserServer), wg *sync.WaitGroup) {
	defer wg.Done()

	port, err := getFreePort()
	if err != nil {
		log.Fatal(err)
	}

	os.Setenv("PORT", strconv.Itoa(port))
	server := exec.Command("npm", "start", "--prefix", os.Getenv("HOME")+"/test-repos/"+user)
	server.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}

	userServer := &UserServer{
		port:   strconv.Itoa(port),
		server: server,
		name:   user,
	}

	*userServers = append(*userServers, userServer)

	// TODO: Add error handling
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Waiting for %v's server to start on port %v\n", user, userServer.port)
	sleepTime, _ := time.ParseDuration("5s")
	time.Sleep(sleepTime)
}

// Kill terminates all the user servers
func Kill(users []*UserServer) {
	for _, user := range users {
		pgid, err := syscall.Getpgid(user.server.Process.Pid)
		if err != nil {
			log.Fatal(err)
		}

		if err := syscall.Kill(-pgid, 9); err != nil {
			log.Fatal(err)
		}
	}
}

func getFreePort() (int, error) {
	mutex := sync.Mutex{}
	mutex.Lock()
	defer mutex.Unlock()

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
