package verification

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func ExecuteDeployment() error {
	// Set executable permissions
	if err := os.Chmod("/home/orangepi/GreatMailer/deployment/backend-deployment.sh", 0755); err != nil {
		log.Printf("Failed to make script executable: %v", err)
		return err
	}

	// Reset git changes if needed
	resetCmd := exec.Command("git", "reset", "/home/orangepi/GreatMailer/deployment/backend-deployment.sh")
	resetOutput, err := resetCmd.CombinedOutput()
	if err != nil {
		log.Printf("Command execution error: %v\nOutput: %s", err, resetOutput)
		return err
	}

	// Create command with nohup to ensure it keeps running
	cmd := exec.Command("nohup", "/home/orangepi/GreatMailer/deployment/backend-deployment.sh")

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
	if err := os.Chmod("/home/orangepi/GreatMailer/deployment/website-deployment.sh", 0755); err != nil {
		log.Printf("Failed to make script executable: %v", err)
		return err
	}
	resetCmd := exec.Command("git", "reset", "/home/orangepi/GreatMailer/deployment/website-deployment.sh")
	resetOutput, err := resetCmd.CombinedOutput()
	if err != nil {
		log.Printf("Command execution error: %v\nOutput: %s", err, resetOutput)
		return err
	}
	cmd := exec.Command("nohup", "bash", "/home/orangepi/GreatMailer/deployment/backend-deployment.sh")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Command execution error: %v\nOutput: %s", err, output)
		return err
	}
	log.Printf("Command output: %s", output)
	return nil
}
