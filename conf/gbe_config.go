// Copyright 2019 GitBitEx.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	confFileName = "conf.json"
)

type GbeConfig struct {
	DataSource       DataSourceConfig       `json:"dataSource"`
	Redis            RedisConfig            `json:"redis"`
	Kafka            KafkaConfig            `json:"kafka"`
	PushServer       PushServerConfig       `json:"pushServer"`
	RestServer       RestServerConfig       `json:"restServer"`
	JwtSecret        string                 `json:"jwtSecret"`
	FrontEndBaseURL  string                 `json:"frontEndBaseURL"`
	BackEndBaseURL   string                 `json:"backEndBaseURL"`
	SecretMainServer SecretMainServerConfig `json:"secretMainServer"`
	ApiStackKey      string                 `json:"apiStackKey"`
	ProjectName      string                 `json:"projectName"`
}

type DataSourceConfig struct {
	DriverName        string `json:"driverName"`
	Addr              string `json:"addr"`
	Database          string `json:"database"`
	User              string `json:"user"`
	Password          string `json:"password"`
	EnableAutoMigrate bool   `json:"enableAutoMigrate"`
}

type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
}
type SecretMainServerConfig struct {
	Addr string `json:"addr"`
}

type KafkaConfig struct {
	Brokers []string `json:"brokers"`
}

type PushServerConfig struct {
	Addr string `json:"addr"`
	Path string `json:"path"`
}

type RestServerConfig struct {
	Addr string `json:"addr"`
}

var config GbeConfig
var configOnce sync.Once

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d) + "/"
}
func SetConfFileName(fileName string) {
	confFileName = fileName
}

func GetConfig() *GbeConfig {
	configOnce.Do(func() {
		rootPath := RootDir()
		rootPath = "/home/wm/backend-exchange/wegaswap-spot/"
		fmt.Println("rootPath:", rootPath)
		configPath := rootPath + confFileName
		fmt.Println("confFileName:", confFileName)
		bytes, err := ioutil.ReadFile(configPath)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &config)
		if err != nil {
			panic(err)
		}
	})
	return &config
}
