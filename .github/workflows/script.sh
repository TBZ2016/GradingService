#!/bin/bash

# Parameters
app_name=$1
image=$2
namespace=${3:-default} # Use 'default' namespace if not provided

echo "Starting script..."
echo "App name: $app_name"
echo "Image: $image"
echo "Namespace: $namespace"

# Ensure the OpenShift project (namespace) exists
if ! oc get project "$namespace" > /dev/null 2>&1; then
    echo "Project $namespace does not exist. Creating..."
    oc new-project "$namespace"
fi

# Check if the deployment already exists
if oc get deployment "$app_name" -n "$namespace" > /dev/null 2>&1; then
    # Deployment exists, update the image
    oc set image deployment/"$app_name" "$app_name"="$image" -n "$namespace"
    oc rollout restart deployment/"$app_name" -n "$namespace"
else
    # Deployment doesn't exist, create a new app
    oc new-app "$image" --name "$app_name" -n "$namespace"
    # The following line is a suggestion to expose the application if needed
    # oc expose svc/"$app_name" -n "$namespace"
fi

echo "Script completed."
