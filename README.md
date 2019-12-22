
# grpc_go

```bash
go get -u github.com/golang/protobuf/protoc-gen-go
protoc -I protos/ protos/pong.proto --go_out=plugins=grpc:protos
```

```bash
protoc -I protos/ protos/spy.proto \
    --include_imports \
    --include_source_info \
    --proto_path=. \
    --descriptor_set_out=api_descriptor.pb \
    --go_out=plugins=grpc:protos
cp protos/* $HOME/Documents/flutter/pong_grpc_flutter/protos/
```

```bash
gcloud iam service-accounts create myaccount \
--display-name "myaccount"
gcloud iam service-accounts keys create ./service-account-creds.json \
--iam-account myaccount@louis030195-256110.iam.gserviceaccount.com
gcloud projects add-iam-policy-binding louis030195-256110 \
--member serviceAccount:myaccount@louis030195-256110.iam.gserviceaccount.com \
--role roles/servicemanagement.serviceController
gcloud projects add-iam-policy-binding louis030195-256110 \
--member serviceAccount:myaccount@louis030195-256110.iam.gserviceaccount.com \
--role roles/cloudtrace.agent
kubectl create secret generic service-account-creds \
  --from-file=service-account-creds.json
kubectl create -f k8s-grpc.yaml
kubectl get service # get ext ip
export SERVER_IP=YOUR_EXTERNAL_IP

```
