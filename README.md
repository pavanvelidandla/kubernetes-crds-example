# crd-code-generation

Example repository for the blog post [Kubernetes Deep Dive: Code Generation for CustomResources](https://blog.openshift.com/kubernetes-deep-dive-code-generation-customresources/).

## Getting Started

First register the custom resource definition:

```
kubectl apply -f artifacts/ConfigGit.yaml
```

Then add an example of the `Database` kind:

```
kubectl apply -f artifacts/configsfromgit.yaml
```

Finally build and run the example:

```
cd wrapper 
go run main.go -kubeconfig ~/.kube/config
```
# kubernetes-crds-example
