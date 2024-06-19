package s3

type (
	Addition struct {
		Bucket                   string `json:"bucket" required:"true"`
		Endpoint                 string `json:"endpoint" required:"true"`
		Region                   string `json:"region"`
		AccessKeyID              string `json:"access_key_id" required:"true"`
		SecretAccessKey          string `json:"secret_access_key" required:"true"`
		SessionToken             string `json:"session_token"`
		CustomHost               string `json:"custom_host"`
		SignURLExpire            int    `json:"sign_url_expire" type:"number" default:"4"`
		Placeholder              string `json:"placeholder"`
		ForcePathStyle           bool   `json:"force_path_style"`
		ListObjectVersion        string `json:"list_object_version" type:"select" options:"v1,v2" default:"v1"`
		RemoveBucket             bool   `json:"remove_bucket" help:"remove_bucket_tips"`
		AddFilenameToDisposition bool   `json:"add_filename_to_disposition" help:"Add filename to Content-Disposition header."`
	}
)
