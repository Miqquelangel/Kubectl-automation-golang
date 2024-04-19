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


         // Execute "kubectl delete ingress" command
         deleteCmd := exec.Command("kubectl", "delete", "ingress",  ingressName, "-n", targetNamespace)
         deleteOutput, err := deleteCmd.CombinedOutput()
         if err != nil {
                         fmt.Println("Error executing 'kubectl delete ingress':", err)
                         return
         }

         // Print the delete output
                         fmt.Printf("\nDescription of %s GCP Gateway:\n\n%s\n\n", ingressName, string(deleteOutput))

}
