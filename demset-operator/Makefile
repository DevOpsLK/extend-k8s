SHELL = bash

clean:
	microk8s kubectl delete deployment demset-operator
	microk8s kubectl delete crd webapps.demoapp.devops.lk

delete-crd:
	microk8s kubectl delete crd webapps.demoapp.devops.lk

create-cr:
	microk8s kubectl create -f deploy/crds/demoapp.devops.lk_v1alpha1_webapp_cr.yaml 

apply-resources:
	microk8s kubectl create -f deploy/service_account.yaml
	microk8s kubectl create -f deploy/role.yaml
	microk8s kubectl create -f deploy/role_binding.yaml
	microk8s kubectl create -f deploy/operator.yaml

create-crd:
	microk8s kubectl create -f deploy/crds/demoapp.devops.lk_webapps_crd.yaml 
	microk8s kubectl create -f deploy/operator.yaml

push-img: 
	docker push  amilaku/demset-operator:v0.0.3

build-img:
	operator-sdk build amilaku/demset-operator:v0.0.3

add-controller:
	operator-sdk add controller --api-version=demoapp.devops.lk/v1alpha1 --kind=WebApp

generate:
	operator-sdk generate k8s; operator-sdk generate crds

add-api:
	operator-sdk add api --api-version=demoapp.devops.lk/v1alpha1 --kind=WebApp

create-proj:
	operator-sdk new demset-operator --repo=github.com/DevOpsLK/demset-operator
	cd demset-operator
	go mod tidy
