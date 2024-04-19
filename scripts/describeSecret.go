package main

import (
        "bufio"
        "fmt"
        "os"
        "os/exec"
        "strings"
)

func main() {

        // Ask for the name of the secret
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("\nEnter the secret name: ")
        secretName, _ := reader.ReadString('\n')
        secretName = strings.TrimSpace(secretName)

        // Execute "kubectl get secrets -A" to get all the secrets
        getSvcCmd := exec.Command("kubectl", "get", "secrets", "-A")
        secretOutput, err := getSvcCmd.CombinedOutput()
        if err != nil {
                fmt.Println("Error executing 'kubectl get secrets -A':", err)
                return
        }

        // Parse the output to find the namespace containing the secret
        lines := strings.Split(string(secretOutput), "\n")
        var targetNamespace string
        for _, line := range lines[1:] { // Skip the header
                columns := strings.Fields(line)
                if len(columns) >= 2 && secretName == columns[1] {
                        targetNamespace = columns[0]
                        break
                }
        }


         // Execute "kubectl describe secret" command
         describeCmd := exec.Command("kubectl", "describe", "secret",  secretName, "-n", targetNamespace)
         describeOutput, err := describeCmd.CombinedOutput()
         if err != nil {
                 fmt.Println("Error executing 'kubectl describe secret':", err)
                 return
         }

         // Print the describe output
                 fmt.Printf("\nDescription of %s secret:\n\n%s\n\n", secretName, string(describeOutput))

}