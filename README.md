### Example BOSH Release

* Update network name to available network in `manifest/manifest.yml`
* `bosh create-release --force`
* `bosh upload-release`
* `export SUBNET=<your-subnet>`
* `bosh -d example-deployment deploy --var=subnet=$SUBNET manifests/manifest.yml`
