package main

import (
        "bufio"
        "fmt"
        "os"
        "os/exec"
        "regexp"
        "strings"
)

func main() {
        // Ask for the command to execute
        reade := bufio.NewReader(os.Stdin)
        fmt.Print("Introduce the command to execute: ")
        input, _ := reade.ReadString('\n')
        input = strings.TrimSpace(input)

        // Ask for the name of the pod
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Enter the name of the pod where the command will be executed: ")
        podName, _ := reader.ReadString('\n')
        podName = strings.TrimSpace(podName)

        // Execute "kubectl get pods -A" to get all the pods
        getPodsCmd := exec.Command("kubectl", "get", "pods", "-A")
        podsOutput, err := getPodsCmd.CombinedOutput()
        if err != nil {
                fmt.Println("Error executing 'kubectl get pods -A':", err)
                return
        }

        // Define the regular expression to capture namespace and pod columns
        re := regexp.MustCompile(`(\S+)\s+(\S+)\s+`)
        matches := re.FindAllStringSubmatch(string(podsOutput), -1)

        // Find the namespace for the specified pod name
        var targetNamespace string
        for _, match := range matches {
                if match[2] == podName {
                        targetNamespace = match[1]
                        break
                }
        }



        command := exec.Command("kubectl", "exec", "-it", podName, "-n", targetNamespace, "--", "bash", "-c", input)
        command.Stdout = os.Stdout
        command.Stdin = os.Stdin
        command.Stderr = os.Stderr

        err = command.Run()
        if err != nil {
                fmt.Println("Error:", err)
        }
}
