.PHONY: deploy delete

ifndef ENV_VARS_FILE
$(error ENV_VARS_FILE is not set)
endif

ifndef GCP_PROJECT
$(error GCP_PROJECT is not set)
endif

deploy:
	gcloud functions deploy SendEmail --runtime go113 --trigger-http --allow-unauthenticated --max-instances=1 --env-vars-file $(ENV_VARS_FILE)  --project=$(GCP_PROJECT)

delete:
	gcloud functions delete SendEmail --project=$(GCP_PROJECT)
