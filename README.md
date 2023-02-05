# CRD-Assignment
sudo su

systemctl start docker

minikube start --driver=docker --force

First I wrote the mysrc.yaml file which is the custom resource that I want to add to my cluster

Next I wrote the resourcedefinition.yaml file which the custom resource definition that contains my custom api-version

Next I ran the commands :
kubectl apply -f resourcedefinition.yaml
kubectl apply -f mysrc.yaml

kubectl get crd

cat mysrc.yaml

cat resourcedefinition.yaml

kubectl get crdtraining.crdtraining.com/crd-assignment

kubectl describe crdtraining.crdtraining.com/crd-assignment

Next I tried to write the custom controller and to get an idea of it I executed the following commands
  git clone https://github.com/kubernetes/sample-controller.git
  cd sample-controller/
  ls -lrth
  more controller.go
