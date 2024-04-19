package main

import (
        "bufio"
        "fmt"
        "os"
        "os/exec"
        "strings"
)

func main() {

        // Ask for the name of the gateway
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("\nEnter the GCP Gateway name: ")
        gatewayName, _ := reader.ReadString('\n')
        gatewayName = strings.TrimSpace(gatewayName)

        // Execute "kubectl get gateway -A" to get all the gateway
        getSvcCmd := exec.Command("kubectl", "get", "gateway", "-A")
        gatewayOutput, err := getSvcCmd.CombinedOutput()
        if err != nil {
                        fmt.Println("Error executing 'kubectl get gateway -A':", err)
                        return
        }

        // Parse the output to find the gateway containing the gateway
        lines := strings.Split(string(gatewayOutput), "\n")
        var targetNamespace string
        for _, line := range lines[1:] { // Skip the header
                        columns := strings.Fields(line)
                        if len(columns) >= 2 && gatewayName == columns[1] {
                                        targetNamespace = columns[0]
                                        break
                        }
        }


         // Execute "kubectl describe gateway" command
         describeCmd := exec.Command("kubectl", "describe", "gateway",  gatewayName, "-n", targetNamespace)
         describeOutput, err := describeCmd.CombinedOutput()
         if err != nil {
                         fmt.Println("Error executing 'kubectl describe gateway':", err)
                         return
         }

         // Print the describe output
                         fmt.Printf("\nDescription of %s GCP Gateway:\n\n%s\n\n", gatewayName, string(describeOutput))

}