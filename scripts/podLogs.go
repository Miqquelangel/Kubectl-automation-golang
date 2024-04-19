package main

import (
        "bufio"
        "fmt"
        "os"
        "os/exec"
        "strings"
)

func main() {
        // Ask for the name of the pod
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Enter the pod name to extract logs: ")
        serviceName, _ := reader.ReadString('\n')
        serviceName = strings.TrimSpace(serviceName)

        // Execute "kubectl get services -A" to get all the services
        serviceOutput, err := exec.Command("kubectl", "get", "pods", "-A").CombinedOutput()
        if err != nil {
                fmt.Println(err)
                return
        }

        // Parse the output to find the namespace containing the service
        lines := strings.Split(string(serviceOutput), "\n")
        var targetNamespace string
        for _, line := range lines { // Skip the headeri
                columns := strings.Fields(line)
                if len(columns) >=2 && serviceName == columns[1] {
                        targetNamespace = columns[0]
                        break
                }
        }

         // Execute "kubectl describe pod" command
         describeCmd := exec.Command("kubectl", "logs", "-n", targetNamespace, serviceName)
         describeOutput, err := describeCmd.CombinedOutput()
         if err != nil {
                 fmt.Println(err)
                 return
         }

         // Print the describe output
         fmt.Println("Logs of the Pod:")
         fmt.Println(string(describeOutput))
}