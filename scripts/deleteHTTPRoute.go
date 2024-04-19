package main

import (
        "bufio"
        "fmt"
        "os"
        "os/exec"
        "strings"
)

func main() {

        // Ask for the name of the httproute
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("\nEnter the httproute name to delete: ")
        httprouteName, _ := reader.ReadString('\n')
        httprouteName = strings.TrimSpace(httprouteName)

        // Execute "kubectl get httproutes -A" to get all the httproutes
        getSvcCmd := exec.Command("kubectl", "get", "httproutes", "-A")
        httprouteOutput, err := getSvcCmd.CombinedOutput()
        if err != nil {
                        fmt.Println("Error executing 'kubectl get httproutes -A':", err)
                        return
        }

        // Parse the output to find the httproute containing the httproute
        lines := strings.Split(string(httprouteOutput), "\n")
        var targetNamespace string
        for _, line := range lines[1:] { // Skip the header
                        columns := strings.Fields(line)
                        if len(columns) >= 2 && httprouteName == columns[1] {
                                        targetNamespace = columns[0]
                                        break
                        }
        }


         // Execute "kubectl delete httproute" command
         deleteCmd := exec.Command("kubectl", "delete", "httproute",  httprouteName, "-n", targetNamespace)
         deleteOutput, err := deleteCmd.CombinedOutput()
         if err != nil {
                         fmt.Println("Error executing 'kubectl delete httproute':", err)
                         return
         }

         // Print the delete output
                         fmt.Printf("\nThe following %s httproute was deleted! \n\n%s\n\n", httprouteName, string(deleteOutput))

}