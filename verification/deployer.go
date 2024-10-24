package verification

import (
	"log"
	"os"
	"os/exec"
)

func ExecuteDeployment() error {
	if err := os.Chmod("/home/orangepi/GreatMailer/deployment/backend-deployment.sh", 0755); err != nil {
		log.Printf("Failed to make script executable: %v", err)
		return err
	}
	resetCmd := exec.Command("git", "reset", "/home/orangepi/GreatMailer/deployment/backend-deployment.sh")
	resetOutput, err := resetCmd.CombinedOutput()
	if err != nil {
		log.Printf("Command execution error: %v\nOutput: %s", err, resetOutput)
		return err
	}
	cmd := exec.Command("/home/orangepi/GreatMailer/deployment/backend-deployment.sh")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Command execution error: %v\nOutput: %s", err, output)
		return err
	}
	log.Printf("Command output: %s", output)
	return nil
}
