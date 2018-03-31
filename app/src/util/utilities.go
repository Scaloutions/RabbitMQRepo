package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/spf13/viper"

	"../data"
)

type (
	Utilities struct {
		v *viper.Viper
	}
)

func NewUtilites(v *viper.Viper) *Utilities {
	// will have another function to set up Redis pool
	return &Utilities{v}
}

func SerializeObject(i interface{}) []byte {

	serrialized, err := json.Marshal(i)
	if err != nil {
		return nil // TODO: add error handling later
	}
	return serrialized
}

func UnserializeObject(data []byte, i interface{}) {

	err := json.Unmarshal(data, i)
	if err != nil {
		// TODO: add error handling later
	}
}

/*
	Utility methods
*/

func (u Utilities) GetIntConfigValue(key string) int {
	return u.v.GetInt(key)
}

func (u Utilities) GetStringConfigValue(key string) string {
	return u.v.GetString(key)
}

func (u Utilities) GetBooleanConfigValue(key string) bool {
	return u.v.GetBool(key)
}

func (u Utilities) GetMapArrConfigValue(key string) []map[string]interface{} {

	return toArrayMap(u.v.Get(key))
}

func (u Utilities) GetActiveEnv() int {
	return u.GetIntConfigValue("general.active")
}

func (u Utilities) GetActiveEnvHost() string {
	prefix := u.getActiveEnvPrefix()
	hostKey := fmt.Sprintf("%s.host", prefix)
	return u.GetStringConfigValue(hostKey)
}

func (u Utilities) GetRabbitmqHost() string {
	prefix := u.getActiveEnvPrefix()
	hostKey := fmt.Sprintf("%s.rabbitmq.host", prefix)
	return u.GetStringConfigValue(hostKey)
}

func (u Utilities) GetRabbitmqPort() int {
	return u.GetIntConfigValue("general.rabbitmq.port")
}

func (u Utilities) GetRabbitmqConnType() string {
	return u.GetStringConfigValue("general.rabbitmq.connection_type")
}

func (u Utilities) GetRabbitmqPass() string {
	prefix := u.getActiveEnvPrefix()
	passKey := fmt.Sprintf("%s.rabbitmq.pass", prefix)
	return u.GetStringConfigValue(passKey)
}

func (u Utilities) GetRabbitmqUser() string {
	prefix := u.getActiveEnvPrefix()
	userKey := fmt.Sprintf("%s.rabbitmq.user", prefix)
	return u.GetStringConfigValue(userKey)
}

func (u Utilities) IsRabbitmqConnEnabled() bool {
	prefix := u.getActiveEnvPrefix()
	enabledKey := fmt.Sprintf("%s.rabbitmq.enabled", prefix)
	return u.GetBooleanConfigValue(enabledKey)
}

func (u Utilities) GetQueueName() string {
	return u.GetStringConfigValue("general.rabbitmq.queue.name")
}

func (u Utilities) GetQueueDurable() bool {
	return u.GetBooleanConfigValue("general.rabbitmq.queue.durable")
}

func (u Utilities) GetQueueAutoDelete() bool {
	return u.GetBooleanConfigValue("general.rabbitmq.queue.auto_delete")
}

func (u Utilities) GetQueueExclusive() bool {
	return u.GetBooleanConfigValue("general.rabbitmq.queue.exclusive")
}

func (u Utilities) GetQueueNoWait() bool {
	return u.GetBooleanConfigValue("general.rabbitmq.queue.no_wait")
}

func (u Utilities) GetQueueConfig() *data.QueueConfig {

	return &data.QueueConfig{
		u.GetQueueName(),
		u.GetQueueDurable(),
		u.GetQueueAutoDelete(),
		u.GetQueueExclusive(),
		u.GetQueueNoWait(),
	}
}

func (u Utilities) GetAuditURL(index int) string {

	uri := u.getUri(index)
	api := u.GetStringConfigValue("api.cache-server.api")
	host := u.GetStringConfigValue("general.cache-server.host")
	port := u.GetStringConfigValue("general.cache-server.port")

	return fmt.Sprintf("http://%s:%s%s%s", host, port, api, uri)
}

func (u Utilities) SendPostRequest(jsonStr []byte, index int) <-chan int {

	url := u.GetAuditURL(index)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	c := make(chan int)

	go func() {

		resp, err := client.Do(req)
		if err != nil {
			// TODO: add error handling later
			fmt.Println("Error: ", err)
		}
		defer resp.Body.Close()
		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		// TODO: add error handling later
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		c <- 0
	}()

	return c

}

/*
	Private methods
*/

func (u Utilities) getUri(index int) string {

	rMap := u.GetMapArrConfigValue("req-map")

	switch index {
	case toIntFromInt64Inteface(rMap[0]["index"]):
		return rMap[0]["uri"].(string)
	case toIntFromInt64Inteface(rMap[1]["index"]):
		return rMap[1]["uri"].(string)
	case toIntFromInt64Inteface(rMap[2]["index"]):
		return rMap[2]["uri"].(string)
	default:
		return ""
	}

}

func (u Utilities) getActiveEnvPrefix() string {
	env := u.GetActiveEnv()
	envMap := u.GetMapArrConfigValue("env-map")

	switch env {
	case toIntFromInt64Inteface(envMap[0]["index"]):
		return envMap[0]["type"].(string)
	case toIntFromInt64Inteface(envMap[1]["index"]):
		return envMap[1]["type"].(string)
	case toIntFromInt64Inteface(envMap[2]["index"]):
		return envMap[2]["type"].(string)
	case toIntFromInt64Inteface(envMap[3]["index"]):
		return envMap[3]["type"].(string)
	default:
		return ""
	}
}

func toArray(i interface{}) []interface{} {
	arr, ok := i.([]interface{})
	if !ok {
		// TODO: add error handling later
	}
	return arr
}

func toMap(i interface{}) map[string]interface{} {
	m, ok := i.(map[string]interface{})
	if !ok {
		// TODO: add error handling later
	}
	return m
}

func toArrayMap(i interface{}) []map[string]interface{} {
	arr := toArray(i)

	mapArr := make([]map[string]interface{}, len(arr))

	for i := 0; i < len(arr); i++ {
		mapArr[i] = toMap(arr[i])
	}

	return mapArr
}

func toIntFromInt64Inteface(i interface{}) int {

	intString := fmt.Sprintf("%d", i)
	integer, _ := strconv.ParseInt(intString, 10, 0)

	return int(integer)
}
