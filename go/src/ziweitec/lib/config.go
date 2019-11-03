package lib

type Config struct {
	Version int `toml:"version"`
	Debug   bool `toml:"debug"`
	Addr   string `toml:"addr"`

	Mysql   struct{ Enable bool `toml:"enable"`; DataSource string `toml:"data_source"` } `toml:"mysql"`
	Sqlite3   struct{ Enable bool `toml:"enable"`; Filepath string `toml:"filepath"` } `toml:"sqlite3"`
	Weapp   struct{ Enable bool `toml:"enable"`; AppId string `toml:"app_id"`; AppSecret string `toml:"app_secret"` } `toml:"weapp"`
	Cors   struct{ Enable bool `toml:"enable"` } `toml:"cors"`
	Html   struct{ Enable bool `toml:"enable"`; StaticDir string `toml:"static_dir"`; TemplateDir string `toml:"template_dir"` } `toml:"html"`
	Session   struct{ Enable bool `toml:"enable"`; Key string `toml:"key"` } `toml:"session"`
}