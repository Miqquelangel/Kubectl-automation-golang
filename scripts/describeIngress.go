package main

import (
        "bufio"
        "fmt"
        "os"
        "os/exec"
        "strings"
)

func main() {

        // Ask for the name of the Ingress
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("\nEnter the Ingress name: ")
        ingressName, _ := reader.ReadString('\n')
        ingressName = strings.TrimSpace(ingressName)

        // Execute "kubectl get ingress -A" to get all the ingress
        getSvcCmd := exec.Command("kubectl", "get", "ingress", "-A")
        ingressOutput, err := getSvcCmd.CombinedOutput()
        if err != nil {
                        fmt.Println("Error executing 'kubectl get ingress -A':", err)
                        return
        }

        // Parse the output to find the ingress containing the ingress
        lines := strings.Split(string(ingressOutput), "\n")
        var targetNamespace string
        for _, line := range lines[1:] { // Skip the header
                        columns := strings.Fields(line)
                        if len(columns) >= 2 && ingressName == columns[1] {
                                        targetNamespace = columns[0]
                                        break
                        }
        }


         // Execute "kubectl describe ingress" command
         describeCmd := exec.Command("kubectl", "describe", "ingress",  ingressName, "-n", targetNamespace)
         describeOutput, err := describeCmd.CombinedOutput()
         if err != nil {
                         fmt.Println("Error executing 'kubectl describe ingress':", err)
                         return
         }

         // Print the describe output
                         fmt.Printf("\nDescription of %s GCP Gateway:\n\n%s\n\n", ingressName, string(describeOutput))

}
