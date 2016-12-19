package main

import (
    "fmt"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/codedeploy"
)

type (
    Repo struct {
        Name     string
        Fullname string
    }

    Commit struct {
        Ref string
    }

    Config struct {
        Key                string
        Secret             string
        Region             string
        Application        string
        DeploymentGroup    string
        DeploymentConfig   string
        Description        string
        IgnoreStopFailures bool
        RevisionType       string
        BundleType         string
        BucketName         string
        BucketKey          string
        BucketEtag         string
        BucketVersion      string
    }

    Plugin struct {
        Repo   Repo
        Commit Commit
        Config Config
    }
)

func newSession(p Plugin) (*session.Session, error) {
    creds := credentials.NewStaticCredentials(p.Config.Key, p.Config.Secret, "")
    if _, err := creds.Get(); err != nil {
        return nil, fmt.Errorf("Bad AWS credentials: %s", err)
    }
    cfg := aws.NewConfig().WithRegion(p.Config.Region).WithCredentials(creds)
    return session.NewSession(cfg)
}

func getRevision(p Plugin) (*codedeploy.RevisionLocation, error) {
    switch p.Config.RevisionType {
    case codedeploy.RevisionLocationTypeGitHub:
        return &codedeploy.RevisionLocation{
            RevisionType: aws.String(p.Config.RevisionType),
            GitHubLocation: &codedeploy.GitHubLocation{
                CommitId:   aws.String(p.Commit.Ref),
                Repository: aws.String(p.Repo.Fullname),
            },
        }, nil
    case codedeploy.RevisionLocationTypeS3:
        if p.Config.BundleType == "" {
            return nil, fmt.Errorf("Please provide a bundle type")
        }
        if p.Config.BucketName == "" {
            return nil, fmt.Errorf("Please provide a bucket name")
        }
        if p.Config.BucketKey == "" {
            return nil, fmt.Errorf("Please provide a bucket key")
        }
        switch p.Config.BundleType {
        case codedeploy.BundleTypeTar:
        case codedeploy.BundleTypeTgz:
        case codedeploy.BundleTypeZip:
        default:
            return nil, fmt.Errorf("Invalid bundle type: %s", p.Config.BundleType)
        }
        s3loc := &codedeploy.S3Location{
            Bucket:     aws.String(p.Config.BucketName),
            BundleType: aws.String(p.Config.BundleType),
            Key:        aws.String(p.Config.BucketKey),
        }
        if p.Config.BucketEtag != "" {
            s3loc.ETag = aws.String(p.Config.BucketEtag)
        }
        if p.Config.BucketVersion != "" {
            s3loc.Version = aws.String(p.Config.BucketVersion)
        }
        return &codedeploy.RevisionLocation{
            RevisionType: aws.String(p.Config.RevisionType),
            S3Location:   s3loc,
        }, nil
    default:
        return nil, fmt.Errorf("Invalid revision type: %s", p.Config.RevisionType)
    }
}

func (p Plugin) Exec() error {
    if p.Config.Key == "" {
        return fmt.Errorf("Please provide an access key id")
    }
    if p.Config.Secret == "" {
        return fmt.Errorf("Please provide a secret access key")
    }
    if p.Config.Region == "" {
        return fmt.Errorf("Please provide a region")
    }
    if p.Config.DeploymentGroup == "" {
        return fmt.Errorf("Please provide a deployment group")
    }

    if p.Config.Application == "" {
        p.Config.Application = p.Repo.Name
    }
    if p.Config.RevisionType == "" {
        p.Config.RevisionType = codedeploy.RevisionLocationTypeGitHub
    }

    sess, err := newSession(p)
    if err != nil {
        return err
    }

    svc := codedeploy.New(sess)

    revision, err := getRevision(p)
    if err != nil {
        return err
    }

    params := &codedeploy.CreateDeploymentInput{
        ApplicationName:               aws.String(p.Config.Application),
        DeploymentGroupName:           aws.String(p.Config.DeploymentGroup),
        Description:                   aws.String(p.Config.Description),
        IgnoreApplicationStopFailures: aws.Bool(p.Config.IgnoreStopFailures),
        Revision:                      revision,
    }

    if p.Config.DeploymentConfig != "" {
        params.DeploymentConfigName = aws.String(p.Config.DeploymentConfig)
    }

    if _, err := svc.CreateDeployment(params); err != nil {
        return err
    }

    return nil
}
