Use this plugin for deplying an application to CodeDeploy.

## Config

The following parameters are used to configure the plugin:

* `access_key` - AWS access key ID
* `secret_key` - AWS secret access key
* `region` - AWS availability zone
* `application` - Application name, defaults to repo name
* `deployment_group` - Name of the deployment group
* `deployment_config` - Name of the deployment config, optional
* `description` - A description about the deployment, optional
* `ignore_stop_failures` - Causes the ApplicationStop deployment lifecycle
  event to fail to a specific instance, defaults to `false`
* `revision_type` - Revision type, defaults to `GitHub`, can be set to `S3`
* `bundle_type` - File type of the application for `S3` revision type
* `bucket_name` - Bucket for `S3` revision type
* `bucket_key` - Key for `S3` revision type
* `bucket_etag` - ETag for `S3` revision type, optional
* `bucket_version` - Version for `S3` revision type, optional

The following secret values can be set to configure the plugin.

* **AWS_ACCESS_KEY_ID** - corresponds to `access_key`
* **AWS_SECRET_ACCESS_KEY** - corresponds to `secret_key`

It is highly recommended to put the **AWS_ACCESS_KEY_ID** and
**AWS_SECRET_ACCESS_KEY** into a secret so it is not exposed to users. This can
be done using the drone-cli.

```bash
drone secret add --image=tralamazza/codedeploy \
    octocat/hello-world AWS_ACCESS_KEY_ID <YOUR_ACCESS_KEY_ID>

drone secret add --image=tralamazza/codedeploy \
    octocat/hello-world AWS_SECRET_ACCESS_KEY <YOUR_SECRET_ACCESS_KEY>
```

Then sign the YAML file after all secrets are added.

```bash
drone sign octocat/hello-world
```

See [secrets](http://readme.drone.io/0.5/usage/secrets/) for additional
information on secrets.

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
pipeline:
  deploy:
    image: tralamazza/codedeploy
    access_key: 970d28f4dd477bc184fbd10b376de753
    secret_key: 9c5785d3ece6a9cdefa42eb99b58986f9095ff1c
    region: us-east-1
    deployment_group: my-deployment
    ignore_stop_failures: true
```
