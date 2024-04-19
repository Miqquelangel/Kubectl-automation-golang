package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	// Ask for the name of the namespace
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter the namespace name: ")
	namespaceName, _ := reader.ReadString('\n')
	namespaceName = strings.TrimSpace(namespaceName)

	// Execute "kubectl get namespaces -A" to get all the namespaces
	getSvcCmd := exec.Command("kubectl", "get", "namespaces", "-A")
	namespaceOutput, err := getSvcCmd.CombinedOutput()
	if err != nil {
			fmt.Println("Error executing 'kubectl get namespaces -A':", err)
			return
	}

	// Parse the output to find the namespace containing the namespace
	lines := strings.Split(string(namespaceOutput), "\n")
	var targetNamespace string
	for _, line := range lines[1:] { // Skip the header
			columns := strings.Fields(line)
			if len(columns) >= 2 && namespaceName == columns[1] {
					targetNamespace = columns[0]
					break
			}
	}


	 // Execute "kubectl describe namespace" command
	 describeCmd := exec.Command("kubectl", "describe", "namespace",  namespaceName, "-n", targetNamespace)
	 describeOutput, err := describeCmd.CombinedOutput()
	 if err != nil {
			 fmt.Println("Error executing 'kubectl describe namespace':", err)
			 return
	 }

	 // Print the describe output
			 fmt.Printf("\nDescription of %s namespace:\n\n%s\n\n", namespaceName, string(describeOutput))

}