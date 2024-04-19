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

        // Ask for the name of the container
        reade := bufio.NewReader(os.Stdin)
        fmt.Print("Enter the name of the container: ")
        container, _ := reade.ReadString('\n')
        container = strings.TrimSpace(container)

        // Ask for the name of the pod
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Enter the name of the pod: ")
        podName, _ := reader.ReadString('\n')
        podName = strings.TrimSpace(podName)

		// Ask for the name of the shell
        reader2 := bufio.NewReader(os.Stdin)
        fmt.Print("\nWhich shell you want? (OPTIONS --> bash or sh):")
        shellName, _ := reader2.ReadString('\n')
        shellName = strings.TrimSpace(shellName)

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

        command := exec.Command("kubectl", "exec", "-it", podName, "-n", targetNamespace, "-c", container, "--", shellName)
        command.Stdout = os.Stdout
        command.Stdin = os.Stdin
        command.Stderr = os.Stderr

        err = command.Run()
        if err != nil {
                fmt.Println("Error:", err)
        }
}
