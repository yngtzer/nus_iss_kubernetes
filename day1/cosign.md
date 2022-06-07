# Cosign

## cosign workflow

### Generate a key-pair
cosign generate-key-pair

### Sign an image that has been pushed to container registry
cosign sign --key cosign.key ghcr.io/fred/dov-bear:v1

### Verify an image
cosign verify --key cosign.pub ghcr.io/fred/dov-bear:v1
