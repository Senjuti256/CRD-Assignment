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

NAME                           CREATED AT
crdtrainings.crdtraining.com   2023-02-05T15:50:28Z

cat mysrc.yaml

cat resourcedefinition.yaml

kubectl get crdtraining.crdtraining.com/crd-assignment

kubectl describe crdtraining.crdtraining.com/crd-assignment

Next I tried to write the custom controller and to get an idea of it I executed the following commands
  git clone https://github.com/kubernetes/sample-controller.git
  cd sample-controller/
  
  ls -lrth
  
  more controller.go
  
  go build -o sample-controller .
  On doing this:-
  
go: downloading k8s.io/client-go v0.0.0-20230130210700-b1350830d0e9
go: downloading k8s.io/klog/v2 v2.80.1
go: downloading github.com/gogo/protobuf v1.3.2
go: downloading k8s.io/utils v0.0.0-20221107191617-1a15be271d1d
go: downloading github.com/google/gofuzz v1.1.0
go: downloading github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da
go: downloading golang.org/x/time v0.0.0-20220210224613-90d013bbcef8
go: downloading github.com/imdario/mergo v0.3.6
go: downloading github.com/spf13/pflag v1.0.5
go: downloading golang.org/x/term v0.4.0
go: downloading sigs.k8s.io/structured-merge-diff/v4 v4.2.3
go: downloading gopkg.in/inf.v0 v0.9.1
go: downloading k8s.io/kube-openapi v0.0.0-20230123231816-1cb3ae25d79a
go: downloading golang.org/x/net v0.5.0
go: downloading github.com/golang/protobuf v1.5.2
go: downloading github.com/google/gnostic v0.5.7-v3refs
go: downloading github.com/davecgh/go-spew v1.1.1
go: downloading github.com/google/go-cmp v0.5.9
go: downloading github.com/go-logr/logr v1.2.3
go: downloading sigs.k8s.io/json v0.0.0-20220713155537-f223a00ba0e2
go: downloading sigs.k8s.io/yaml v1.3.0
go: downloading golang.org/x/sys v0.4.0
go: downloading golang.org/x/oauth2 v0.0.0-20220223155221-ee480838109b
go: downloading google.golang.org/protobuf v1.28.1
go: downloading github.com/json-iterator/go v1.1.12
go: downloading gopkg.in/yaml.v3 v3.0.1
go: downloading gopkg.in/yaml.v2 v2.4.0
go: downloading golang.org/x/text v0.6.0
go: downloading github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822
go: downloading github.com/emicklei/go-restful/v3 v3.9.0
go: downloading github.com/go-openapi/swag v0.22.3
go: downloading github.com/go-openapi/jsonreference v0.20.1
go: downloading github.com/modern-go/reflect2 v1.0.2
go: downloading github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
go: downloading github.com/go-openapi/jsonpointer v0.19.6
go: downloading github.com/mailru/easyjson v0.7.7
go: downloading github.com/josharian/intern v1.0.0



[root@fedora sample-controller]# kubectl create -f artifacts/examples/crd-status-subresource.yaml
customresourcedefinition.apiextensions.k8s.io/foos.samplecontroller.k8s.io created

[root@fedora sample-controller]# kubectl create -f artifacts/examples/example-foo.yaml
foo.samplecontroller.k8s.io/example-foo created

[root@fedora sample-controller]# kubectl get deployments

NAME        READY   UP-TO-DATE   AVAILABLE   AGE
my-go-app   1/1     1            1           3d14h

[root@fedora sample-controller]# kubectl get crd

NAME                           CREATED AT
crdtrainings.crdtraining.com   2023-02-05T15:50:28Z
foos.samplecontroller.k8s.io   2023-02-06T06:07:38Z
