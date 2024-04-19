package main

import (
        "bufio"
        "fmt"
        "os"
        "os/exec"
        "strings"
)

func main() {

        // Ask for the name of the configmap
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("\nEnter the configmap name to delete: ")
        configmapName, _ := reader.ReadString('\n')
        configmapName = strings.TrimSpace(configmapName)

        // Execute "kubectl get configmaps -A" to get all the configmaps
        getSvcCmd := exec.Command("kubectl", "get", "configmaps", "-A")
        configmapOutput, err := getSvcCmd.CombinedOutput()
        if err != nil {
                        fmt.Println("Error executing 'kubectl get configmaps -A':", err)
                        return
        }

        // Parse the output to find the configmap containing the configmap
        lines := strings.Split(string(configmapOutput), "\n")
        var targetNamespace string
        for _, line := range lines[1:] { // Skip the header
                        columns := strings.Fields(line)
                        if len(columns) >= 2 && configmapName == columns[1] {
                                        targetNamespace = columns[0]
                                        break
                        }
        }


         // Execute "kubectl delete configmap" command
         deleteCmd := exec.Command("kubectl", "delete", "configmap",  configmapName, "-n", targetNamespace)
         deleteOutput, err := deleteCmd.CombinedOutput()
         if err != nil {
                         fmt.Println("Error executing 'kubectl delete configmap':", err)
                         return
         }

         // Print the delete output
                         fmt.Printf("\nThe following %s configmap was deleted! \n\n%s\n\n", configmapName, string(deleteOutput))

}