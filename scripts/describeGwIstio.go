package main

import (
        "bufio"
        "fmt"
        "os"
        "os/exec"
        "strings"
)

func main() {

        // Ask for the name of the gw
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("\nEnter the Istio gateway name: ")
        gwName, _ := reader.ReadString('\n')
        gwName = strings.TrimSpace(gwName)

        // Execute "kubectl get gw -A" to get all the gw
        getSvcCmd := exec.Command("kubectl", "get", "gw", "-A")
        gwOutput, err := getSvcCmd.CombinedOutput()
        if err != nil {
                        fmt.Println("Error executing 'kubectl get gw -A':", err)
                        return
        }

        // Parse the output to find the gw containing the gw
        lines := strings.Split(string(gwOutput), "\n")
        var targetNamespace string
        for _, line := range lines[1:] { // Skip the header
                        columns := strings.Fields(line)
                        if len(columns) >= 2 && gwName == columns[1] {
                                        targetNamespace = columns[0]
                                        break
                        }
        }


         // Execute "kubectl describe gw" command
         describeCmd := exec.Command("kubectl", "describe", "gw",  gwName, "-n", targetNamespace)
         describeOutput, err := describeCmd.CombinedOutput()
         if err != nil {
                         fmt.Println("Error executing 'kubectl describe gw':", err)
                         return
         }

         // Print the describe output
                         fmt.Printf("\nDescription of %s Istio gateway:\n\n%s\n\n", gwName, string(describeOutput))

}