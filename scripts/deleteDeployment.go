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
        fmt.Print("\nEnter the deployment name to delete: ")
        deploymentName, _ := reader.ReadString('\n')
        deploymentName = strings.TrimSpace(deploymentName)

        // Execute "kubectl get deployments -A" to get all the deployments
        getSvcCmd := exec.Command("kubectl", "get", "deployments", "-A")
        deploymentOutput, err := getSvcCmd.CombinedOutput()
        if err != nil {
                        fmt.Println("Error executing 'kubectl get deployments -A':", err)
                        return
        }

        // Parse the output to find the deployment containing the deployment
        lines := strings.Split(string(deploymentOutput), "\n")
        var targetNamespace string
        for _, line := range lines[1:] { // Skip the header
                        columns := strings.Fields(line)
                        if len(columns) >= 2 && deploymentName == columns[1] {
                                        targetNamespace = columns[0]
                                        break
                        }
        }


         // Execute "kubectl delete deployment" command
         deleteCmd := exec.Command("kubectl", "delete", "deployment",  deploymentName, "-n", targetNamespace)
         deleteOutput, err := deleteCmd.CombinedOutput()
         if err != nil {
                         fmt.Println("Error executing 'kubectl delete deployment':", err)
                         return
         }

         // Print the delete output
                         fmt.Printf("\nThe following %s deployment was deleted! \n\n%s\n\n", deploymentName, string(deleteOutput))

}