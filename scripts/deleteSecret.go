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
        fmt.Print("\nEnter the secret name to delete: ")
        secretName, _ := reader.ReadString('\n')
        secretName = strings.TrimSpace(secretName)

        // Execute "kubectl get secrets -A" to get all the secrets
        getSvcCmd := exec.Command("kubectl", "get", "secrets", "-A")
        secretOutput, err := getSvcCmd.CombinedOutput()
        if err != nil {
                        fmt.Println("Error executing 'kubectl get secrets -A':", err)
                        return
        }

        // Parse the output to find the secret containing the secret
        lines := strings.Split(string(secretOutput), "\n")
        var targetsecret string
        for _, line := range lines[1:] { // Skip the header
                        columns := strings.Fields(line)
                        if len(columns) >= 2 && secretName == columns[1] {
                                        targetsecret = columns[0]
                                        break
                        }
        }


         // Execute "kubectl delete secret" command
         deleteCmd := exec.Command("kubectl", "delete", "secret",  secretName, "-n", targetsecret)
         deleteOutput, err := deleteCmd.CombinedOutput()
         if err != nil {
                         fmt.Println("Error executing 'kubectl delete secret':", err)
                         return
         }

         // Print the delete output
                         fmt.Printf("\nThe following %s secret was deleted! \n\n%s\n\n", secretName, string(deleteOutput))

}