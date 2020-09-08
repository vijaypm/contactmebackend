.PHONY: deploy delete test

check-env:
ifndef ENV_VARS_FILE
	$(error ENV_VARS_FILE is not set)
endif

check-project: 
ifndef GCP_PROJECT
	$(error GCP_PROJECT is not set)
endif


test: check-env
	godotenv -f $(ENV_VARS_FILE) go test

deploy: check-env check-project
	gcloud functions deploy SendEmail --runtime go113 --trigger-http --allow-unauthenticated --max-instances=1 --env-vars-file $(ENV_VARS_FILE)  --project=$(GCP_PROJECT)

delete: check-env check-project
	gcloud functions delete SendEmail --project=$(GCP_PROJECT)
