package verification

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func ExecuteDeployment() error {
	// Create command with nohup to ensure it keeps running
	cmd := exec.Command("nohup", "/home/orangepi/dev/deployment/backend-deployment.sh")

	// Redirect stdout and stderr to files
	stdout, err := os.OpenFile("/tmp/deployment.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	cmd.Stdout = stdout
	cmd.Stderr = stdout

	// Detach from parent process
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
		Pgid:    0,
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		log.Printf("Failed to start deployment: %v", err)
		return err
	}

	log.Printf("Deployment script started successfully with PID %d", cmd.Process.Pid)
	return nil
}

func ExecuteWebsiteDeployment() error {

	cmd := exec.Command("/home/orangepi/dev/deployment/website-deployment.sh")
	err := cmd.Start()
	if err != nil {
		log.Printf("Command execution error: %v\n", err)
		return err
	}
	log.Printf("Command output: %s", output)
	return nil
}
