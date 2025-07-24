package task

var Tasks map[string]Interface

func init() {
	Tasks = map[string]Interface{
		"alpine-linux": &AlpineLinux{},
		"go":           &GoDev{},
		"jetbrains":    &JetBrains{},
		"openssl":      &OpenSSL{},
		"python":       &Python{},
	}
}
