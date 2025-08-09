package builder

import (
//"fmt"
//"github.com/jeremie-izzo/dctl/pkg/profile"
//"github.com/jeremie-izzo/dctl/pkg/template"
//"path/filepath"
//"strings"
)

//
//func BuildTemplate(serviceName string, data any) (string, error) {
//	// Assume the template is at templates/docker-compose.tmpl
//	path := filepath.Join("templates", "docker-compose.tmpl")
//
//	tmplStr, err := template.Load(path)
//	if err != nil {
//		return "", err
//	}
//
//	return template.Render("compose", tmplStr, data, nil)
//}
//
//// TODO : Move this to a more appropriate location
//var config = map[string]string{
//	"MYSQL_PASSWORD":     "testest",
//	"MYSQL_USER":         "test",
//	"MYSQL_DATABASE":     "test",
//	"MYSQL_PORT":         "13306",
//	"GCP_PROJECT_ID":     "gcp-local",
//	"DATASTORE_PORT":     "18081",
//	"PUBSUB_PORT":        "18085",
//	"ELASTICSEARCH_PORT": "19200",
//	"REDIS_PORT":         "16379",
//	"TEMPORAL_PORT":      "17233",
//	"TEMPORAL_UI_PORT":   "18080",
//	"KIBANA_PORT":        "15601",
//	"DATASTORE_UI_PORT":  "18282",
//	"VOLUME_PATH":        "/Users/jeremie/Repositories/lightspeed-invoicing/godev/volume",
//}
//
//// TODO: Move this to a more appropriate location
//func DefaultFuncMap(pro profile.Profile) map[string]any {
//	return map[string]any{
//		"substituteVars": func(val interface{}) string {
//			// Convert to string first
//			strVal := fmt.Sprintf("%v", val)
//			// Replace {VAR_NAME} patterns within the string
//			result := strVal
//			// Replace all variables from config
//			for varName, value := range config {
//				result = strings.ReplaceAll(result, "{"+varName+"}", value)
//			}
//			return result
//		},
//		"contains": func(s, substr string) bool {
//			return strings.Contains(s, substr)
//		},
//		"joinPath": func(service string, path string) string {
//			base := pro.GetServicePath(service)
//			return filepath.Join(base, path)
//		},
//	}
//}
