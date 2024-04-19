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
        fmt.Print("\nEnter the pod name: ")
        podName, _ := reader.ReadString('\n')
        podName = strings.TrimSpace(podName)

        // Execute "kubectl get pods -A" to get all the pods
        getSvcCmd := exec.Command("kubectl", "get", "pods", "-A")
        podOutput, err := getSvcCmd.CombinedOutput()
        if err != nil {
                fmt.Println("Error executing 'kubectl get pods -A':", err)
                return
        }

        // Parse the output to find the namespace containing the pod
        lines := strings.Split(string(podOutput), "\n")
        var targetNamespace string
        for _, line := range lines[1:] { // Skip the header
                columns := strings.Fields(line)
                if len(columns) >= 2 && podName == columns[1] {
                        targetNamespace = columns[0]
                        break
                }
        }


         // Execute "kubectl describe pod" command
         describeCmd := exec.Command("kubectl", "describe", "pod",  podName, "-n", targetNamespace)
         describeOutput, err := describeCmd.CombinedOutput()
         if err != nil {
                 fmt.Println("Error executing 'kubectl describe pod':", err)
                 return
         }

         // Print the describe output
                 fmt.Printf("\nPod description of %s pod:\n\n%s\n\n", podName, string(describeOutput))

}