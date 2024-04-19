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
        fmt.Print("\nEnter the Istio Gateway name to delete: ")
        gwName, _ := reader.ReadString('\n')
        gwName = strings.TrimSpace(gwName)

        // Execute "kubectl get gws -A" to get all the gws
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


         // Execute "kubectl delete gw" command
         deleteCmd := exec.Command("kubectl", "delete", "gw",  gwName, "-n", targetNamespace)
         deleteOutput, err := deleteCmd.CombinedOutput()
         if err != nil {
                         fmt.Println("Error executing 'kubectl delete gw':", err)
                         return
         }

         // Print the delete output
                         fmt.Printf("\nThe following %s Istio Gateway was deleted! \n\n%s\n\n", gwName, string(deleteOutput))

}