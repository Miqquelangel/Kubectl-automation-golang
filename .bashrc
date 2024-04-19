
#KUBECTL ALIASES OF KIND ‘APPLY’

alias ka='kubectl apply -f'

# KUBECTL ALIASES OF KIND 'GET'

alias knamespace='kubectl get namespaces -o wide'
alias kpods='kubectl get pods'
alias knodes='kubectl get nodes -o wide'
alias ksvc='kubectl get services'
alias kdep='kubectl get deployments'
alias ksecrets='kubectl get secrets'
alias kpv='kubectl get pv'
alias kpvc='kubectl get pvc'
alias kconfigmap='kubectl get configmap'
alias kcnp='kubectl get cnp'
alias kga='kubectl get gateway -o wide -A'
alias khttproute='kubectl get httproute'
alias kgw='kubectl get gw'
alias kvirtualservice='kubectl get virtualservice'
alias kdestinationrule='kubectl get destinationrule'
alias kingress='kubectl get ingress'

#KUBECTL ALIASES OF KIND ‘DESCRIBE’

alias kdesnamespace='go run /scripts/describeNamespace.go'
alias kdespod='go run /scripts/describePod.go'
alias kdesnodes='kubectl describe nodes'
alias kdessvc='go run /scripts/describeService.go'
alias kdesdep='go run /scripts/describeDeployment.go'
alias kdessecret='go run /scripts/describeSecret.go'
alias kdespv='kubectl describe pv'
alias kdespvc='kubectl describe pvc'
alias kdesconfigmap='go run /scripts/describeConfigmap.go'
alias kdescnp='kubectl describe cnp'
alias kdesga='go run /scripts/describeGatewayGCP.go'
alias kdeshttproute='go run /scripts/describeHTTPRoute.go'
alias kdesgw='go run /scripts/describeGwIstio.go'
alias kdesvirtualservice='go run /scripts/describeVirtualService.go'
alias kdesdestinationrule='go run /scripts/describeDestinationRule.go'
alias kdesingress='go run /scripts/describeIngress.go'

#KUBECTL ALIASES OF KIND ‘DELETE’

alias kdeletenamespace='go run /scripts/deleteNamespace.go'
alias kdeletepod='go run /scripts/deletePod.go'
alias kdeletenodelete='kubectl delete nodelete'
alias kdeletesvc='go run /scripts/deleteService.go'
alias kdeletedep='go run /scripts/deleteDeployment.go'
alias kdeletesecret='go run /scripts/deleteSecret.go'
alias kdeletepv='kubectl delete pv'
alias kdeletepvc='kubectl delete pvc'
alias kdeleteconfigmap='go run /scripts/deleteConfigmap.go'
alias kdeletecnp='kubectl delete cnp'
alias kdeletega='go run /scripts/deleteGatewayGCP.go'
alias kdeleteht='go run /scripts/deleteHTTPRoute.go'
alias kdeletegw='go run /scripts/deleteGwIstio.go'
alias kdeletevs='go run /scripts/deleteVirtualService.go'
alias kdeleteds='go run /scripts/deleteDestinationRule.go'
alias kdeleteingress='go run /scripts/deleteIngress.go'

#KUBECTL ALIASES OF KIND ‘SCALE’

alias kscale='kubectl scale deployment'

#KUBECTL ALIASES OF KIND ‘EXEC’

alias kexec='go run /scripts/executePod.go'
alias kcontainer='go run /scripts/executeContainer.go'
alias kexeccommand='go run /scripts/executeCommandPod.go'

#KUBECTL ALIASES OF KIND ‘LOGS’

alias klogs='go run /scripts/podLogs.go'
alias klogscontainer='go run /scripts/containerLogs.go'

#KUBECTL ALIASES OF KIND ‘USERS’

alias kwho='kubectl whoami'

