package engine

type Config struct {
	AutoStart         bool   `json:"auto_start,omitempty"`
	DisableEncryption bool   `json:"disable_encryption,omitempty"`
	DownloadDirectory string `json:"download_directory,omitempty"`
	EnableUpload      bool   `json:"enable_upload,omitempty"`
	EnableSeeding     bool   `json:"enable_seeding,omitempty"`
	IncomingPort      int    `json:"incoming_port,omitempty"`
}
