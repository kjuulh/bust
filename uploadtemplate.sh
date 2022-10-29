#!/bin/bash

drone template add --namespace kjuulh --name dagger_go_template.yaml --data @dagger_go_template.yaml || true
drone template update --namespace kjuulh --name dagger_go_template.yaml --data @dagger_go_template.yaml
