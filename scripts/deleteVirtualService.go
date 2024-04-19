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
        fmt.Print("\nEnter the virtualservice name to delete: ")
        virtualserviceName, _ := reader.ReadString('\n')
        virtualserviceName = strings.TrimSpace(virtualserviceName)

        // Execute "kubectl get virtualservices -A" to get all the virtualservices
        getSvcCmd := exec.Command("kubectl", "get", "virtualservices", "-A")
        virtualserviceOutput, err := getSvcCmd.CombinedOutput()
        if err != nil {
                        fmt.Println("Error executing 'kubectl get virtualservices -A':", err)
                        return
        }

        // Parse the output to find the virtualservice containing the virtualservice
        lines := strings.Split(string(virtualserviceOutput), "\n")
        var targetvirtualservice string
        for _, line := range lines[1:] { // Skip the header
                        columns := strings.Fields(line)
                        if len(columns) >= 2 && virtualserviceName == columns[1] {
                                        targetvirtualservice = columns[0]
                                        break
                        }
        }


         // Execute "kubectl delete virtualservice" command
         deleteCmd := exec.Command("kubectl", "delete", "virtualservice",  virtualserviceName, "-n", targetvirtualservice)
         deleteOutput, err := deleteCmd.CombinedOutput()
         if err != nil {
                         fmt.Println("Error executing 'kubectl delete virtualservice':", err)
                         return
         }

         // Print the delete output
                         fmt.Printf("\nThe following %s virtualservice was deleted! \n\n%s\n\n", virtualserviceName, string(deleteOutput))

}