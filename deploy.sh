#!/bin/bash

init_app="app.yaml"
env=".env"

echo "" > secret.yaml

cat "$init_app" > secret.yaml
echo -e "\nenv_variables:" >> secret.yaml
sed 's/=/: /g' "$env" | sed 's/^/  /'  >> secret.yaml

cat secret.yaml

gcloud app deploy secret.yaml