#sh ./install_docker
# echo "logout and log back in"

sudo snap install kubectl --classic

gcloud container clusters create lab06cluster --zone us-west1-c

kubectl run foobar --image=nginx

kubectl get deployments

kubectl get pods

kubectl delete deployment foobar


# gcloud container clusters create finalprojectcluster --zone us-west1-c
