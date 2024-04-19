package main

import (
        "bufio"
        "fmt"
        "os"
        "os/exec"
        "strings"
)

func main() {

        // Ask for the name of the virtualservice
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("\nEnter the virtualservice name: ")
        virtualserviceName, _ := reader.ReadString('\n')
        virtualserviceName = strings.TrimSpace(virtualserviceName)

        // Execute "kubectl get virtualservice -A" to get all the virtualservice
        getSvcCmd := exec.Command("kubectl", "get", "virtualservice", "-A")
        virtualserviceOutput, err := getSvcCmd.CombinedOutput()
        if err != nil {
                        fmt.Println("Error executing 'kubectl get virtualservice -A':", err)
                        return
        }

        // Parse the output to find the virtualservice containing the virtualservice
        lines := strings.Split(string(virtualserviceOutput), "\n")
        var targetNamespace string
        for _, line := range lines[1:] { // Skip the header
                        columns := strings.Fields(line)
                        if len(columns) >= 2 && virtualserviceName == columns[1] {
                                        targetNamespace = columns[0]
                                        break
                        }
        }


         // Execute "kubectl describe virtualservice" command
         describeCmd := exec.Command("kubectl", "describe", "virtualservice",  virtualserviceName, "-n", targetNamespace)
         describeOutput, err := describeCmd.CombinedOutput()
         if err != nil {
                         fmt.Println("Error executing 'kubectl describe virtualservice':", err)
                         return
         }

         // Print the describe output
                         fmt.Printf("\nDescription of %s virtualservice:\n\n%s\n\n", virtualserviceName, string(describeOutput))

}