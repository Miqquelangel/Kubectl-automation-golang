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
        fmt.Print("\nEnter the httproute name: ")
        httprouteName, _ := reader.ReadString('\n')
        httprouteName = strings.TrimSpace(httprouteName)

        // Execute "kubectl get httproute -A" to get all the httproute
        getSvcCmd := exec.Command("kubectl", "get", "httproute", "-A")
        httprouteOutput, err := getSvcCmd.CombinedOutput()
        if err != nil {
                        fmt.Println("Error executing 'kubectl get httproute -A':", err)
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


         // Execute "kubectl describe httproute" command
         describeCmd := exec.Command("kubectl", "describe", "httproute",  httprouteName, "-n", targetNamespace)
         describeOutput, err := describeCmd.CombinedOutput()
         if err != nil {
                         fmt.Println("Error executing 'kubectl describe httproute':", err)
                         return
         }

         // Print the describe output
                         fmt.Printf("\nDescription of %s httproute:\n\n%s\n\n", httprouteName, string(describeOutput))

}