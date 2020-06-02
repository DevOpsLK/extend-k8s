SHELL = bash

clean:
	 kubectl delete deployment demset-operator
	 kubectl delete crd webapps.demoapp.devops.lk

delete-crd:
	 kubectl delete crd webapps.demoapp.devops.lk

create-cr:
	 kubectl create -f deploy/crds/demoapp.devops.lk_v1alpha1_webapp_cr.yaml 

apply-resources:
	 kubectl create -f deploy/service_account.yaml
	 kubectl create -f deploy/role.yaml
	 kubectl create -f deploy/role_binding.yaml
	 kubectl create -f deploy/operator.yaml

create-crd:
	 kubectl create -f deploy/crds/demoapp.devops.lk_webapps_crd.yaml 
	 kubectl create -f deploy/operator.yaml

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
