package verification

import (
	"log"
	"os"
	"os/exec"
)

func ExecuteDeployment() error {
	if err := os.Chmod("../deployment/backend-deployment.sh", 0755); err != nil {
		log.Printf("Failed to make script executable: %v", err)
		return err
	}
	cmd := exec.Command("../deployment/backend-deployment.sh")
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Command execution error: %v\nOutput: %s", err, output)
		return err
	}
	log.Printf("Command output: %s", output)
	return nil
}