protoCompile:
	protoc --go_out=./src/services/protos --go_opt=paths=source_relative --go-grpc_out=./src/services/protos --go-grpc_opt=paths=source_relative ./PerimeterX.proto

rollout:
	kubectl rollout restart deployments perimeterx -n antibots
	kubectl rollout status deployments perimeterx -n antibots

dockerBuild:
	docker build -t registry.digitalocean.com/athenabot/antibots/perimeterx:latest .

kubeApply:
	kubectl apply -f Deployment.yml

kubeCreate:
	kubectl create -f Deployment.yml

scratchDeploy:
	doctl auth switch --context athena
	doctl registry login
	doctl kubernetes cluster kubeconfig save athena
	make dockerBuild
	docker push registry.digitalocean.com/athenabot/antibots/perimeterx:latest
	make kubeCreate
	kubectl rollout status deployments perimeterx -n antibots