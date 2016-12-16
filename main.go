package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"
    "github.com/urfave/cli"
)

var version = "0"

func main() {
    app := cli.NewApp()
    app.Name = "codedeploy plugin"
    app.Usage = "codedeploy plugin"
    app.Action = run
    app.Version = version
    app.Flags = []cli.Flag{
        //
        // Config args
        //
        cli.StringFlag{
            Name:   "access-key",
            Usage:  "AWS access key ID",
            EnvVar: "PLUGIN_ACCESS_KEY,AWS_ACCESS_KEY_ID",
        },
        cli.StringFlag{
            Name:   "secret-key",
            Usage:  "AWS secret access key",
            EnvVar: "PLUGIN_SECRET_KEY,AWS_SECRET_ACCESS_KEY",
        },
        cli.StringFlag{
            Name:   "region",
            Usage:  "AWS availability zone",
            EnvVar: "PLUGIN_REGION",
        },
        cli.StringFlag{
            Name:   "application",
            Usage:  "Application name, defaults to repo name",
            EnvVar: "PLUGIN_APPLICATION",
        },
        cli.StringFlag{
            Name:   "deployment_group",
            Usage:  "Name of the deployment group",
            EnvVar: "PLUGIN_DEPLOYMENT_GROUP",
        },
        cli.StringFlag{
            Name:   "deployment_config",
            Usage:  "Name of the deployment config, optional",
            EnvVar: "PLUGIN_DEPLOYMENT_CONFIG",
        },
        cli.StringFlag{
            Name:   "description",
            Usage:  "A description about the deployment, optional",
            EnvVar: "PLUGIN_DESCRIPTION",
        },
        cli.BoolFlag{
            Name:   "ignore_stop_failures",
            Usage:  "Causes the ApplicationStop deployment lifecycle event to fail to a specific instance, defaults to `false`",
            EnvVar: "PLUGIN_IGNORE_STOP_FAILURES",
        },
        cli.StringFlag{
            Name:   "revision_type",
            Usage:  "Revision type, defaults to GitHub, can be set to S3",
            EnvVar: "PLUGIN_REVISION_TYPE",
        },
        cli.StringFlag{
            Name:   "bundle_type",
            Usage:  "File type of the application for S3 revision type",
            EnvVar: "PLUGIN_BUNDLE_TYPE",
        },
        cli.StringFlag{
            Name:   "bucket_name",
            Usage:  "Bucket for S3 revision type",
            EnvVar: "PLUGIN_BUCKET_NAME",
        },
        cli.StringFlag{
            Name:   "bucket_key",
            Usage:  "Key for S3 revision type",
            EnvVar: "PLUGIN_BUCKET_KEY",
        },
        cli.StringFlag{
            Name:   "bucket_etag",
            Usage:  "ETag for S3 revision type, optional",
            EnvVar: "PLUGIN_BUCKET_ETAG",
        },
        cli.StringFlag{
            Name:   "bucket_version",
            Usage:  "Version for S3 revision type, optional",
            EnvVar: "PLUGIN_BUCKET_VERSION",
        },
        cli.StringFlag{
            Name:  "env-file",
            Usage: "source env file",
        },

        //
        // repo args
        //
        cli.StringFlag{
            Name:   "repo.fullname",
            Usage:  "repository full name",
            EnvVar: "DRONE_REPO",
        },
        cli.StringFlag{
            Name:   "repo.name",
            Usage:  "repository name",
            EnvVar: "DRONE_REPO_NAME",
        },

        //
        // commit args
        //
        cli.StringFlag{
            Name:   "commit.ref",
            Value:  "refs/heads/master",
            Usage:  "git commit ref",
            EnvVar: "DRONE_COMMIT_REF",
        },
    }
    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
}

func run(c *cli.Context) error {
    if c.String("env-file") != "" {
        _ = godotenv.Load(c.String("env-file"))
    }

    plugin := Plugin{
        Repo: Repo{
            Name:     c.String("repo.name"),
            Fullname: c.String("repo.fullname"),
        },
        Commit: Commit{
            Ref: c.String("commit.ref"),
        },
        Config: Config{
            Key:                c.String("access-key"),
            Secret:             c.String("secret-key"),
            Region:             c.String("region"),
            Application:        c.String("application"),
            DeploymentGroup:    c.String("deployment_group"),
            DeploymentConfig:   c.String("deployment_config"),
            Description:        c.String("description"),
            IgnoreStopFailures: c.Bool("ignore_stop_failures"),
            RevisionType:       c.String("revision_type"),
            BundleType:         c.String("bundle_type"),
            BucketName:         c.String("bucket_name"),
            BucketKey:          c.String("bucket_key"),
            BucketEtag:         c.String("bucket_etag"),
            BucketVersion:      c.String("bucket_version"),
        },
    }
    return plugin.Exec()
}
