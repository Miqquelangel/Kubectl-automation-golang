package main

import (
        "bufio"
        "fmt"
        "os"
        "os/exec"
        "strings"
)

func main() {

        // Ask for the name of the destinationrule
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("\nEnter the destinationrule name to delete: ")
        destinationruleName, _ := reader.ReadString('\n')
        destinationruleName = strings.TrimSpace(destinationruleName)

        // Execute "kubectl get destinationrules -A" to get all the destinationrules
        getSvcCmd := exec.Command("kubectl", "get", "destinationrules", "-A")
        destinationruleOutput, err := getSvcCmd.CombinedOutput()
        if err != nil {
                        fmt.Println("Error executing 'kubectl get destinationrules -A':", err)
                        return
        }

        // Parse the output to find the destinationrule containing the destinationrule
        lines := strings.Split(string(destinationruleOutput), "\n")
        var targetNamespace string
        for _, line := range lines[1:] { // Skip the header
                        columns := strings.Fields(line)
                        if len(columns) >= 2 && destinationruleName == columns[1] {
                                        targetNamespace = columns[0]
                                        break
                        }
        }


         // Execute "kubectl delete destinationrule" command
         deleteCmd := exec.Command("kubectl", "delete", "destinationrule",  destinationruleName, "-n", targetNamespace)
         deleteOutput, err := deleteCmd.CombinedOutput()
         if err != nil {
                         fmt.Println("Error executing 'kubectl delete destinationrule':", err)
                         return
         }

         // Print the delete output
                         fmt.Printf("\nThe following %s destinationrule was deleted! \n\n%s\n\n", destinationruleName, string(deleteOutput))

}