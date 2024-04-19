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
        fmt.Print("\nEnter the destinationrule name: ")
        destinationruleName, _ := reader.ReadString('\n')
        destinationruleName = strings.TrimSpace(destinationruleName)

        // Execute "kubectl get destinationrule -A" to get all the destinationrule
        getSvcCmd := exec.Command("kubectl", "get", "destinationrule", "-A")
        destinationruleOutput, err := getSvcCmd.CombinedOutput()
        if err != nil {
                        fmt.Println("Error executing 'kubectl get destinationrule -A':", err)
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


         // Execute "kubectl describe destinationrule" command
         describeCmd := exec.Command("kubectl", "describe", "destinationrule",  destinationruleName, "-n", targetNamespace)
         describeOutput, err := describeCmd.CombinedOutput()
         if err != nil {
                         fmt.Println("Error executing 'kubectl describe destinationrule':", err)
                         return
         }

         // Print the describe output
                         fmt.Printf("\nDescription of %s destinationrule:\n\n%s\n\n", destinationruleName, string(describeOutput))

}