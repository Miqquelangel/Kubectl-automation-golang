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
        fmt.Print("\nEnter the service name to delete: ")
        serviceName, _ := reader.ReadString('\n')
        serviceName = strings.TrimSpace(serviceName)

        // Execute "kubectl get services -A" to get all the services
        getSvcCmd := exec.Command("kubectl", "get", "services", "-A")
        serviceOutput, err := getSvcCmd.CombinedOutput()
        if err != nil {
                        fmt.Println("Error executing 'kubectl get services -A':", err)
                        return
        }

        // Parse the output to find the service containing the service
        lines := strings.Split(string(serviceOutput), "\n")
        var targetservice string
        for _, line := range lines[1:] { // Skip the header
                        columns := strings.Fields(line)
                        if len(columns) >= 2 && serviceName == columns[1] {
                                        targetservice = columns[0]
                                        break
                        }
        }


         // Execute "kubectl delete service" command
         deleteCmd := exec.Command("kubectl", "delete", "service",  serviceName, "-n", targetservice)
         deleteOutput, err := deleteCmd.CombinedOutput()
         if err != nil {
                         fmt.Println("Error executing 'kubectl delete service':", err)
                         return
         }

         // Print the delete output
                         fmt.Printf("\nThe following %s service was deleted! \n\n%s\n\n", serviceName, string(deleteOutput))

}