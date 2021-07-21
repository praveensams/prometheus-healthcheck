# prometheus-healthcheck
## Custom health monitor tool which exports metrics to prometheus

##Image preparation
docker build -t validator .
docker push validator

##Kubernetes Charts

make the relevent changes to chart value.yaml in the place of image

cat values.yaml
-------------------

image:
  repository: <validator>
  pullPolicy: IfNotPresent
  # Overrides the image tag whose default is the chart appVersion.
  tag: "latest"


---------------------

Add the scrape in the prometheus

  - name: Production Down
      rules:
      - alert: ProdDown
        expr:  sam_health{job="health_exporter"} != 1
        for: 15m
        labels:
          severity: P1
          team: infra
        annotations:
          description: 'Prod URL  {{ $labels.exported_instance }} down for 5 minutes'
          summary: 'Prod url  {{ $labels.exported_instance }} down for 5 minutes'

