ts = $(shell date +%s)

deploy:
	docker build . \
		-t us-central1-docker.pkg.dev/hargow/toolbox/toolbox:$(ts) \
		--platform linux/amd64
	docker push us-central1-docker.pkg.dev/hargow/toolbox/toolbox:$(ts)
	sed 's/%tag%/$(ts)/' deployment.yaml | kubectl apply -f -

