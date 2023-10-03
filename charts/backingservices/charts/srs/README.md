# Search and Reporting Service Helm chart

The Pega `Search and Reporting Service` or `SRS` backing service can replace the embedded search feature of Pega Infinity Platform. To use it in your deployment, you provision and deploy it independently as an external service which provides search and reporting capabilities with a Pega Infinity environment.

## Configuring a backing service with your pega environment

You can provision this SRS into your `pega` environment namespace, with the SRS endpoint configured with the Pega Infinity environment. When you include the SRS into your pega namespace, the service endpoint is included within your Pega Infinity environment network to ensure isolation of your application data.

## Search and Reporting Service support

The Search and Reporting Service provides next generation search and reporting capabilities for Pega Infinity 8.6 and later.

This service replaces the legacy search module from the platform with an independently deployable and scalable service along with the built-in capabilities to support more than one Pega environments with its data isolation features in Pega Infinity 8.6 and later.
The service deployment provisions runtime service pods along with a dependency on a backing technology ElasticSearch service for storage and retrieval of data.

### SRS Version compatibility matrix

| Pega Infinity version | SRS version | Elasticsearch version | Description                                                                                                                                                                                                                                                                                                           |
|-----------------------|-------------|-----------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| < 8.6                 | NA          | NA                    | SRS can be used with Pega Infinity 8.6 and later                                                                                                                                                                                                                                                                      |
| \>= 8.6              | 1.25.3  | 7.10.2, 7.16.3, and 7.17.9      | While SRS Docker images are certified against Elasticsearch versions 7.10.2, 7.16.3 and 7.17.9, Pega recommends using Elasticsearch version 7.17.9. To stay current with Pega releases, use the latest available SRS image 1.25.3.

**Note**: 

**If your deployment uses the internally-provisioned Elasticsearch:** To migrate to Elasticsearch version 7.17.9 from the Elasticsearch version 7.10.2 or 7.16.3 use the process that applies to your deployment:

* Update the SRS Docker image version to use v1.25.3, which supports both Elasticsearch versions 7.10.x and 7.16.x.
* Update the Elasticsearch `dependencies.version` parameter in the [requirement.yaml](../../requirements.yaml) to 7.17.3.
* Update Elasticsearch to 7.17.9.

**If your deployment connects to an externally-managed Elasticsearch service:** To migrate to Elasticsearch version 7.17.9 from the Elasticsearch version 7.10.2 or 7.16.3 use the process that applies to your deployment:

* Update the SRS Docker image version to use v1.25.3, which supports both Elasticsearch versions 7.10.x and 7.16.x.
* Complete the version upgrade to 7.17.9. Refer to Elasticsearch version 7.17 documentation. For example, see [Upgrade Elasticsearch](https://www.elastic.co/guide/en/elasticsearch/reference/7.17/setup-upgrade.html).

### SRS runtime configuration

The values.yaml provides configuration options to define the deployment resources along with option to either provision ElasticSearch cluster automatically for data storage, or you can choose to configure an existing externally managed elasticsearch cluster to use as a datastore with the SRS runtime.

If an externally managed elasticsearch cluster is being used, make sure the service is accessible to the k8s cluster where SRS is deployed.

You may enable the component of [Elasticsearch](https://github.com/helm/charts/tree/master/stable/elasticsearch/values.yaml) in the backingservices by configuring the 'srs.srsStorage' section in values.yaml file to deploy ElasticSearch cluster automatically. For more configuration options available for each of the components, see their Helm Charts.

Note: Pega does **not** actively update the elasticsearch dependency in `requirements.yaml`. To leverage SRS, you must do one of the following:

* To use the internally-provided Elasticsearch service in the SRS cluster, use the default `srs.enabled.true` parameter and set the Elasticsearch version by updating the `elasticsearch.imageTag` parameter in the [values.yaml](./values.yaml) to match the `dependencies.version` parameter in the [requirements.yaml](../../requirements.yaml).
* To use an externally-provided Elasticsearch service with SRS, use the default `srs.enabled.true` parameter, update the `srs.srsStorage.provisionInternalESCluster` parameter in the [values.yaml](./values.yaml) to `false` and then provide connection details as documented below.

### Deploying SRS with Pega-provided busybox images
To deploy Pega Platform with the SRS backing service, the SRS helm chart requires the use of the busybox image.  For clients who want to pull this image from a registry other than Docker Hub, they must tag and push their image to another registry, and then pull it by specifying `busybox.image` and `busybox.imagePullPolicy`.

### Configuration settings

| Configuration                           | Usage                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
|-----------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `enabled`                               | Enable the Search and Reporting Service deployment as a backing service. Set this parameter to `true` to use SRS.                                                                                                                                                                                                                                                                                                                                                                                                                            |
| `deploymentName`                        | Specify the name of your SRS cluster. Your deployment creates resources prefixed with this string. This is also the service name for the SRS.                                                                                                                                                                                                                                                                                                                                                          |
| `srsRuntime`                            | Use this section to define specific resource configuration options like image, replica count, cpu and memory resource settings in the SRS.                                                                                                                                                                                                                                                                                                                                                             |
| `busybox`                               | When provisioning an internally managed Elasticsearch cluster, you can customize the location and pull policy of the Alpine image used during the deployment process by specifying `busybox.image` and `busybox.imagePullPolicy`.                                                                                                                                                                                                                                                                      |
| `elasticsearch`                         | Define the elasticsearch cluster configurations. The [Elasticsearch](https://github.com/helm/charts/tree/master/stable/elasticsearch/values.yaml) chart defines the values for Elasticsearch provisioning in the SRS cluster. For internally provisioned Elasticsearch the default version is set to `7.10.2`. Set the `elasticsearch.imageTag` parameter in values.yaml to `7.16.3` to use this supported version in the SRS cluster.                                                                |
| `k8sProvider`                               | Specify your Kubernetes provider name. Supported values are [`eks`, `aks`, `minikube`, `gke`, `openshift`, `pks`].. 

### Enabling security between SRS and Elasticsearch
To configure a secure connection between the SRS cluster and Elasticsearch, add the following the settings in your backingservices configuration file to reflect your organization's connectivity setup.

| Configuration                            | Usage                                                                                                                                                                                                                                            |
|------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `tls`                                    | Set to `true` to enable the SRS service to authenticate to your organization's available Elasticsearch service.                                                                                                                                  |
| `srsStorage.provisionInternalESCluster`  | <ul><li>Set to `true` to enable this parameter to provide an internally managed and secured Elasticsearch cluster to be used with the SRS cluster. After you specify an Elasticsearch version in the SRS Helm chart and save the file, run `$ make es-prerequisite NAMESPACE=<NAMESPACE> ELASTICSEARCH_VERSION=<ELASTICSEARCH_VERSION>`. </li><li>Where `NAMESPACE` references your deployment namespace of the SRS cluster and `ELASTICSEARCH_VERSION` matches the Elasticsearch version you want to use in [values.yaml](../../values.yaml) and [requirements.yaml](../../requirements.yaml).</li></ul> |

To connect to external elasticsearch below configuration needs to be made.
Certificates used by external elasticsearch need to be placed in an accessible location for make command to use.
eg: If certs are placed under /home/certs. Make command will look like this:
make external-es-secrets NAMESPACE=pegabackingservices ELASTICSEARCH_VERSION=7.10.2 PATH_TO_CERTIFICATE=/home/certs/truststore.jks

| Configuration                           | Usage                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
|-----------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `tls`                                   | Set to `true` to enable the SRS service to authenticate to your organization's available Elasticsearch service.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| `esCredentials.username`                | Enter the username for your available Elasticsearch service. This username value must match the values you set in the connection info section of esCredentials.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| `esCredentials.password`                | Enter the required password for your available Elasticsearch service. This password value must match the values you set in the connection info section of esCredentials.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| `srsStorage.provisionInternalESCluster` | <ul><li>Set to false to disable this parameter and connect to your available Elasticsearch service from the SRS cluster. Disabling this setting requires you to provide connectivity details to your organization's external Elasticsearch service along with an appropriate TLS certificate with which you authenticate with the service. To pass the required certificate to the cluster using a secrets file, run the command, `$ make external-es-secrets NAMESPACE=<NAMESPACE_USED_FOR_DEPLOYMENT> ELASTICSEARCH_VERSION=<ELASTICSEARCH_VERSION> PATH_TO_CERTIFICATE=<PATH_TO_CERTS>`. </li><li>where NAMESPACE references your deployment namespace of the SRS cluster, `ELASTICSEARCH_VERSION` matches the Elasticsearch version you want to use, and `PATH_TO_CERTIFICATE` points to the location where you copied the required certificates on your location machine.</li></ul> |
| `domain`                                | Enter the DNS entry associated with your external Elasticsearch service.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |

Note: Only .p12 and .jks certificates are supported.


### Enable request authentication/authorization mechanism using identity provider(IdP) between SRS and Pega Infinity
To configure authentication/authorization mechanism using identity provider(IdP) between SRS and Pega Infinity, add the following the settings in your backingservices configuration file and then in pega chart's `values.yml` / pega helm inspected `pega.yaml`.

| Configuration                      | Usage                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
|------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `srsRuntime.env.AuthEnabled`       | <ul><li>Set to `false` to disable authentication/authorization between SRS and Pega Infinity.</li><li>Set to `true` to enable authentication/authorization using identity provider(IdP) between SRS and Pega Infinity.</li></ul>                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| `srsRuntime.env.OAuthPublicKeyURL` | <ul><li>Provide a valid OAuth Public Key URL from IdP, which can be used in SRS to fetch public keys from IdP and validate the request that is coming into SRS.</li><li> Make sure the authorization server of IdP is having 1) A scope with name `pega.search:full` 2) `scp` & `guid` claims under `pega.search:full` scope with scope name(`pega.search:full`) and `<Customer Deployment Id>` respectively</li><li>Example of JWT token payload having `scp` and `guid` claims <pre>{   "ver": 1,   "jti": "AT.EmHCGDFHE18hC5j3stjbarVonh46twW7tWutB9v8hsw",   "iss": "https://prod-pega.okta.com/oauth2/aus8ahm2k777777",   "aud": "srs-cmc-stg",   "iat": 1678097157,   "exp": 1678100757,   "cid": "0oa8acv8gbBPRCZ7I5d7",   "scp": [     "pega.search:full"   ],   "auth_time": 1000,   "sub": "0oa8acv8gbBPRCZ7I5d7",   "guid": "bf0d4cb6-e09f-1111-ab19-9aac5156b618" } </pre></ul> |

### Enable TLS/HTTPS between SRS and Pega Infinity
There are multiple ways to enable TLS/HTTPS for service which deployed in k8s. It all depends on the k8s setup, environment setup and network policies/restrictions etc. You can work with IT departments to get this done. Some ways to enable TLS/HTTPS are 1. [Using ingress with TLS ](https://kubernetes.io/docs/concepts/services-networking/ingress/#tls). 2. [Using load balancer ](https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer) etc. Some cloud provided k8s specific ways are 1) AWS: Using combination of [AWS Certificate Manager](https://aws.amazon.com/certificate-manager/), [AWS Load Balancer Controller](https://aws.amazon.com/blogs/opensource/kubernetes-ingress-aws-alb-ingress-controller/), [Kubernetes Nginx Ingress Controller](https://kubernetes.github.io/ingress-nginx/deploy/#aws). 2. Azure: Using combination of [Azure Key Vault](https://azure.microsoft.com/en-us/services/key-vault/), [Kubernetes Nginx Ingress Controller](https://kubernetes.github.io/ingress-nginx/deploy/#azure), [Azure Application Gateway](https://docs.microsoft.com/en-us/azure/application-gateway/). 3. GCP: Using combination of [Google-managed SSL certificates](https://cloud.google.com/load-balancing/docs/ssl-certificates/google-managed-certs), [GCP HTTP(S) Load Balancer](https://cloud.google.com/load-balancing/docs/https/), [Kubernetes Nginx Ingress Controller](https://kubernetes.github.io/ingress-nginx/deploy/#gce-gke) etc.

Example:

```yaml
srs:
  # Set srs.enabled=true to enable SRS
  enabled: true

  # specify unique name for the deployment based on org app and/or srs applicable environment name. eg: acme-demo-dev-srs
  deploymentName: "YOUR_SRS_DEPLOYMENT_NAME"

  # Configure the location of the busybox image that is used during the deployment process of
  # the internal Elasticsearch cluster
  busybox:
    image: "alpine:3.16.0"
    imagePullPolicy: "IfNotPresent"

  srsRuntime:
    # Number of pods to provision
    replicaCount: 2

    # docker image of the srs-service, platform-services/search-n-reporting-service:dockerTag
    srsImage: "YOUR_SRS_IMAGE:TAG"

    env:
      # AuthEnabled may be set to true when there is an authentication mechanism in place between SRS and Pega Infinity.
      AuthEnabled: false
      # When `AuthEnabled` is `true`, enter the appropriate public key URL. When `AuthEnabled` is `false`(default), leave this parameter empty.
      OAuthPublicKeyURL: ""

  # This section specifies the elasticsearch cluster configuration.
  srsStorage:
    # Setting srsStorage.provisionInternalESCluster to true will provision an internal elasticsearch cluster using the configuration
    # specified in the `elasticsearch` section
    # IF you do not enable SRS and the srs.enabled parameter is set to false, always set srs.srsStorage.provisionInternalESCluster=false
    provisionInternalESCluster: true
    # To use your own Elasticsearch cluster, set srsStorage.provisionInternalESCluster to false and then
    # set the external Elasticsearch cluster URL and port details below when using an externally managed elasticsearch
    # Ensure that the specified endpoint is accessible from the kubernetes cluster pods.
    # domain: ""
    # port: 9200
    # protocol: https
    # The elasticsearch connection supports three authentication methods: basic authentication , AWS IAM role-based authentication and Elasticsearch secure connection(tls).
    # Set srs.srsStorage.tls.enabled: true to enable the use of TLS-based authentication to your Elasticsearch service whether is it running as an internalized or externalized service in your SRS cluster.
    tls:
      enabled: false
    # To specify a certificate used to authenticate an external Elasticsearch service (with tls.enabled: true and srsStorage.provisionInternalESCluster: false), uncomment the following line to specify the TLS certificate name for your Elasticsearch service.
    # certificateName: "Certificate_Name"
    # Set srs.srsStorage.basicAuthentication.enabled: true to enable the use of basic authentication to your Elasticsearch service whether is it running as an internalized or externalized service in your SRS cluster.
    basicAuthentication:
      enabled: true
    # To configure basic authentication or TLS-based authentication to your externally-managed Elasticsearch service in your SRS cluster, uncomment and add the parameter details: srs.srsStorage.esCredentials.username and srs.srsStorage.esCredentials.password.
    # esCredentials:
    #   username: "username"
    #   password: "password"
    # To configure AWS IAM role-based authentication to your externally-managed Elasticsearch cluster, uncomment
    # and add the parameter details: srs.srsStorage.awsIAM and its associated region, srs.srsStorage.awsIAM.region
    # awsIAM:
    #   region: "AWS_ELASTICSEARCH_REGION"
    # To configure either authentication method, when the elasticsearch domain requires an open internet connection,set the requireInternetAccess parameter to "true".
    requireInternetAccess: false

```
