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
        fmt.Print("\nEnter the GCP gateway name to delete: ")
        gatewayName, _ := reader.ReadString('\n')
        gatewayName = strings.TrimSpace(gatewayName)

        // Execute "kubectl get gateways -A" to get all the gateways
        getSvcCmd := exec.Command("kubectl", "get", "gateways", "-A")
        gatewayOutput, err := getSvcCmd.CombinedOutput()
        if err != nil {
                        fmt.Println("Error executing 'kubectl get gateways -A':", err)
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


         // Execute "kubectl delete gateway" command
         deleteCmd := exec.Command("kubectl", "delete", "gateway",  gatewayName, "-n", targetNamespace)
         deleteOutput, err := deleteCmd.CombinedOutput()
         if err != nil {
                         fmt.Println("Error executing 'kubectl delete gateway':", err)
                         return
         }

         // Print the delete output
                         fmt.Printf("\nThe following %s gateway was deleted! \n\n%s\n\n", gatewayName, string(deleteOutput))

}