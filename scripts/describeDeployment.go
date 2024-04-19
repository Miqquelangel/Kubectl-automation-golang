package main

import (
        "bufio"
        "fmt"
        "os"
        "os/exec"
        "strings"
)

func main() {

        // Ask for the name of the deployment
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("\nEnter the deployment name: ")
        deploymentName, _ := reader.ReadString('\n')
        deploymentName = strings.TrimSpace(deploymentName)

        // Execute "kubectl get deployments -A" to get all the deployments
        getSvcCmd := exec.Command("kubectl", "get", "deployments", "-A")
        deploymentOutput, err := getSvcCmd.CombinedOutput()
        if err != nil {
                fmt.Println("Error executing 'kubectl get deployments -A':", err)
                return
        }

        // Parse the output to find the namespace containing the deployment
        lines := strings.Split(string(deploymentOutput), "\n")
        var targetNamespace string
        for _, line := range lines[1:] { // Skip the header
                columns := strings.Fields(line)
                if len(columns) >= 2 && deploymentName == columns[1] {
                        targetNamespace = columns[0]
                        break
                }
        }


         // Execute "kubectl describe deployment" command
         describeCmd := exec.Command("kubectl", "describe", "deployment",  deploymentName, "-n", targetNamespace)
         describeOutput, err := describeCmd.CombinedOutput()
         if err != nil {
                 fmt.Println("Error executing 'kubectl describe deployment':", err)
                 return
         }

         // Print the describe output
                 fmt.Printf("\nDeployment description of %s deployment:\n\n%s\n\n", deploymentName, string(describeOutput))
}