package s3

import (
	"fmt"
	"future-admin/internal/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3 struct {
	Addition
	client  *s3.S3
	session *session.Session
	config  model.DriverConfig
}

func (d *S3) Init() error {
	if err := d.doInitSession(); err != nil {
		return err
	}

	d.doInitClient()

	return nil
}

func (d *S3) doInitClient() {
	d.client = s3.New(d.session)
}

func (d *S3) doInitSession() (err error) {
	cfg := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(d.AccessKeyID, d.SecretAccessKey, d.SessionToken),
		Region:           &d.Region,
		Endpoint:         &d.Endpoint,
		S3ForcePathStyle: aws.Bool(d.ForcePathStyle),
	}
	d.session, err = session.NewSession(cfg)
	return err
}

func (d *S3) Config() model.DriverConfig {
	return d.config
}

func (d *S3) List() error {
	var (
		continuationToken *string
		startAfter        *string
		prefix            = "/"
	)

	input := &s3.ListObjectsV2Input{
		Bucket:            &d.Bucket,
		ContinuationToken: continuationToken,
		Prefix:            &prefix,
		Delimiter:         aws.String("/"),
		StartAfter:        startAfter,
	}
	listObjectsResult, err := d.client.ListObjectsV2(input)
	if err != nil {
		return err
	}
	fmt.Println(listObjectsResult)
	return nil
}

func (d *S3) GetAddition() any {
	return d.Addition
}
