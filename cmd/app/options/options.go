package options

type ServerRunOptions struct {
	Mongodb mongodbConfig
	Redis   redisConfig
}

type DatabaseConfig struct {
	Address string
	Port    uint
}

type RedisConfig struct {
	Address string
	Port    uint
}

func NewServerRunOptions() *ServerRunOptions {
	s := ServerRunOptions{}
}

func (s *ServerRunOptions) AddFlags(fs *pflag.FlagSet) {
	fs.DurationVar(&s.EventTTL, "event-ttl", s.EventTTL,
		"Amount of time to retain events.")
	fs.BoolVar(&s.EnableLogsHandler, "enable-logs-handler", s.EnableLogsHandler,
		"If true, install a /logs handler for the apiserver logs.")
	fs.StringVar(&s.SSHKeyfile, "ssh-keyfile", s.SSHKeyfile,
		"If non-empty, use secure SSH proxy to the nodes, using this user keyfile")

	fs.Int64Var(&s.MaxConnectionBytesPerSec, "max-connection-bytes-per-sec", s.MaxConnectionBytesPerSec, ""+
		"If non-zero, throttle each user connection to this number of bytes/sec. "+
		"Currently only applies to long-running requests.")

	fs.IntVar(&s.MasterCount, "apiserver-count", s.MasterCount,
		"The number of apiservers running in the cluster, must be a positive number.")

	fs.StringVar(&s.EndpointReconcilerType, "alpha-endpoint-reconciler-type", string(s.EndpointReconcilerType),
		"Use an endpoint reconciler ("+strings.Join(reconcilers.AllTypes.Names(), ", ")+")")
}
