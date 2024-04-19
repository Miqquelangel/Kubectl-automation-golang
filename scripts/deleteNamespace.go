package main

import (
        "bufio"
        "fmt"
        "os"
        "os/exec"
        "strings"
)

func main() {

        // Ask for the name of the namespace
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("\nEnter the namespace name to delete: ")
        namespaceName, _ := reader.ReadString('\n')
        namespaceName = strings.TrimSpace(namespaceName)

        // Execute "kubectl get namespaces -A" to get all the namespaces
        getSvcCmd := exec.Command("kubectl", "get", "namespaces", "-A")
        namespaceOutput, err := getSvcCmd.CombinedOutput()
        if err != nil {
                        fmt.Println("Error executing 'kubectl get namespaces -A':", err)
                        return
        }

        // Parse the output to find the namespace containing the namespace
        lines := strings.Split(string(namespaceOutput), "\n")
        var targetNamespace string
        for _, line := range lines[1:] { // Skip the header
                        columns := strings.Fields(line)
                        if len(columns) >= 2 && namespaceName == columns[1] {
                                        targetNamespace = columns[0]
                                        break
                        }
        }


         // Execute "kubectl delete namespace" command
         deleteCmd := exec.Command("kubectl", "delete", "namespace",  namespaceName, "-n", targetNamespace)
         deleteOutput, err := deleteCmd.CombinedOutput()
         if err != nil {
                         fmt.Println("Error executing 'kubectl delete namespace':", err)
                         return
         }

         // Print the delete output
                         fmt.Printf("\nThe following %s namespace was deleted! \n\n%s\n\n", namespaceName, string(deleteOutput))

}