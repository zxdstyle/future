package s3

import (
	"encoding/json"
	"future-admin/internal/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"path"
	"strings"
)

type S3 struct {
	Addition
	client  *s3.S3
	session *session.Session
	config  model.DriverConfig
}

func (d *S3) Init(addition string) error {
	if err := json.Unmarshal([]byte(addition), &d.Addition); err != nil {
		return err
	}

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

func getKey(path string, dir bool) string {
	path = strings.TrimPrefix(path, "/")
	if path != "" && dir {
		path += "/"
	}
	return path
}
func (d *S3) List(dir string) ([]*model.Object, error) {
	var (
		marker string
		prefix = getKey(dir, true)
		input  = &s3.ListObjectsInput{
			Bucket: &d.Bucket,
			Marker: &marker,
			//ContinuationToken: continuationToken,
			Prefix:    &prefix,
			Delimiter: aws.String("/"),
			//StartAfter:        startAfter,
		}

		files = make([]*model.Object, 0)
	)

	listObjectsResult, err := d.client.ListObjects(input)
	if err != nil {
		return nil, err
	}
	for _, object := range listObjectsResult.CommonPrefixes {
		name := path.Base(strings.Trim(*object.Prefix, "/"))
		file := model.Object{
			//Id:        *object.Key,
			Name:     name,
			IsFolder: true,
		}
		files = append(files, &file)
	}

	for _, object := range listObjectsResult.Contents {
		name := path.Base(*object.Key)
		file := model.Object{
			//Id:        *object.Key,
			Name:     name,
			Size:     *object.Size,
			Modified: *object.LastModified,
		}
		files = append(files, &file)
	}

	return files, nil
}

func (d *S3) GetAddition() any {
	return d.Addition
}
