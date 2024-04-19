package main

import (
        "bufio"
        "fmt"
        "os"
        "os/exec"
        "strings"
)

func main() {

        // Ask for the name of the pod
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("\nEnter the pod name to delete: ")
        podName, _ := reader.ReadString('\n')
        podName = strings.TrimSpace(podName)

        // Execute "kubectl get pods -A" to get all the pods
        getSvcCmd := exec.Command("kubectl", "get", "pods", "-A")
        podOutput, err := getSvcCmd.CombinedOutput()
        if err != nil {
                        fmt.Println("Error executing 'kubectl get pods -A':", err)
                        return
        }

        // Parse the output to find the pod containing the pod
        lines := strings.Split(string(podOutput), "\n")
        var targetpod string
        for _, line := range lines[1:] { // Skip the header
                        columns := strings.Fields(line)
                        if len(columns) >= 2 && podName == columns[1] {
                                        targetpod = columns[0]
                                        break
                        }
        }


         // Execute "kubectl delete pod" command
         deleteCmd := exec.Command("kubectl", "delete", "pod",  podName, "-n", targetpod)
         deleteOutput, err := deleteCmd.CombinedOutput()
         if err != nil {
                         fmt.Println("Error executing 'kubectl delete pod':", err)
                         return
         }

         // Print the delete output
                         fmt.Printf("\nThe following %s pod was deleted! \n\n%s\n\n", podName, string(deleteOutput))

}