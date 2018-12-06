Script to run in jenkins for CI

```
#!/bin/bash
project="${JOB_NAME}"
author="mustakim"
branch=$(git branch | grep \* | cut -d ' ' -f2)
kubectl="kubectl --namespace ${branch}"
docker_image="${author}/${project}:${branch}"

# build the project
go build app

# build the new docker image
echo "building docker image: ${docker_image}"
sudo docker build -t ${docker_image} ${WORKSPACE}

# create namespace if it doesn't exist
if ! kubectl get namespace ${branch} > /dev/null 2>&1; then
        echo "${branch} namespace doesn't exist, creating now..."
        kubectl create namespace ${branch}
fi

# create deployment if it doesn't exist
if ! kubectl get deployment ${project} > /dev/null 2>&1; then
		echo "${branch} deployment doesn't exist, creating now..."
        ${kubectl} create -f ${WORKSPACE}/kubernetes/01_deployment.yml
fi

# create service if it doesn't exist
if ! kubectl get service ${project} > /dev/null 2>&1; then
		echo "${branch} service doesn't exist, creating now..."
        ${kubectl} create -f ${WORKSPACE}/kubernetes/02_service.yml
fi

# set the image for the deployment
echo "setting image: ${docker_image} for deployment: ${project}"
${kubectl} set image deployment ${project} ${project}=${docker_image}
```
