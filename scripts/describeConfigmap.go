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
        fmt.Print("\nEnter the configmap name: ")
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


         // Execute "kubectl describe configmap" command
         describeCmd := exec.Command("kubectl", "describe", "configmap",  configmapName, "-n", targetNamespace)
         describeOutput, err := describeCmd.CombinedOutput()
         if err != nil {
                         fmt.Println("Error executing 'kubectl describe configmap':", err)
                         return
         }

         // Print the describe output
                         fmt.Printf("\nDescription of %s configmap:\n\n%s\n\n", configmapName, string(describeOutput))

}