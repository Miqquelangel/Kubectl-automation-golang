package main

import (
        "bufio"
        "fmt"
        "os"
        "os/exec"
        "strings"
)

func main() {

        // Ask for the name of the service
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("\nEnter the service name: ")
        serviceName, _ := reader.ReadString('\n')
        serviceName = strings.TrimSpace(serviceName)

        // Execute "kubectl get services -A" to get all the services
        getSvcCmd := exec.Command("kubectl", "get", "services", "-A")
        serviceOutput, err := getSvcCmd.CombinedOutput()
        if err != nil {
                fmt.Println("Error executing 'kubectl get services -A':", err)
                return
        }

        // Parse the output to find the namespace containing the service
        lines := strings.Split(string(serviceOutput), "\n")
        var targetNamespace string
        for _, line := range lines[1:] { // Skip the header
                columns := strings.Fields(line)
                if len(columns) >= 2 && serviceName == columns[1] {
                        targetNamespace = columns[0]
                        break
                }
        }


         // Execute "kubectl describe service" command
         describeCmd := exec.Command("kubectl", "describe", "service",  serviceName, "-n", targetNamespace)
         describeOutput, err := describeCmd.CombinedOutput()
         if err != nil {
                 fmt.Println("Error executing 'kubectl describe service':", err)
                 return
         }

         // Print the describe output
		 fmt.Printf("\nService description of %s service:\n\n%s\n\n", serviceName, string(describeOutput))

}