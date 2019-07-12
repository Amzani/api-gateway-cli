package commands

import (
	"bytes"
	"fmt"
	"regexp"
	"text/template"

	"github.com/Amzani/api-gateway-cli/swagger"
	"github.com/go-openapi/spec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Path type
type Path struct {
	URI         string
	ServiceName string
	ServicePort int
}

// Ingress kong Type
type Ingress struct {
	Name             string
	NameSpace        string
	UpstreamHost     string
	UpstreamPort     int
	BasePath         string
	ConnectTimeout   int
	Retries          int
	ReadTimeout      int
	WriteTimeout     int
	RouteProtocols   []string
	RouteMethods     []string
	IsolationEnabled bool
	Host             string
	Paths            []Path
}

var oasFile string

func generate(oasSpec *spec.Swagger, api *Ingress) {
	var buf bytes.Buffer
	api.Paths = make([]Path, len(swagger.PathList(oasSpec)))
	for i, path := range swagger.PathList(oasSpec) {
		r := regexp.MustCompile(`({[\w]+})`)
		path = r.ReplaceAllLiteralString(path, "[a-zA-Z0-9_-]+")
		api.Paths[i] = Path{URI: path, ServiceName: api.UpstreamHost, ServicePort: api.UpstreamPort}
	}
	t, _ := template.ParseFiles("template.yaml")

	if err := t.Execute(&buf, api); err != nil {
		fmt.Println("Erreur")
	}
	fmt.Println(buf.String())
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate an API Gateway format for deployement (k8s, deck...)",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		api := &Ingress{Name: "checkout",
			NameSpace:        viper.GetString("namespace"),
			BasePath:         viper.GetString("basepath"),
			ConnectTimeout:   viper.GetInt("connect_timeout"),
			Retries:          viper.GetInt("retries"),
			ReadTimeout:      viper.GetInt("read_timeout"),
			WriteTimeout:     viper.GetInt("write_timeout"),
			RouteProtocols:   viper.GetStringSlice("protocols"),
			RouteMethods:     viper.GetStringSlice("methods"),
			UpstreamHost:     viper.GetString("upstream_host"),
			UpstreamPort:     viper.GetInt("upstream_port"),
			IsolationEnabled: false,
			Host:             viper.GetString("host"),
		}
		OasSpec, err := swagger.Parse(oasFile)
		if err != nil {
			fmt.Println("Error while parsing your swagger.yaml")
		}
		generate(OasSpec, api)
	},
}

func init() {
	generateCmd.Flags().StringVarP(&oasFile, "spec", "s", "./swagger.yaml", "Help message for toggle")
}
