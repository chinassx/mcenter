package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/mcenter/apps/endpoint"
)

func init() {
	ioc.RegistryApi(&handler{})
}

type handler struct {
	service endpoint.Service
	log     logger.Logger
	ioc.IocObjectImpl
}

func (h *handler) Init() error {
	h.log = zap.L().Named(h.Name())
	h.service = ioc.GetController(endpoint.AppName).(endpoint.Service)
	return nil
}

func (h *handler) Name() string {
	return endpoint.AppName
}

func (h *handler) Version() string {
	return "v1"
}

func (h *handler) Registry(ws *restful.WebService) {
	tags := []string{"服务功能"}
	ws.Route(ws.POST("/").To(h.RegistryEndpoint).
		Doc("注册服务功能列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(endpoint.RegistryRequest{}).
		Writes(endpoint.EndpointSet{}))

	ws.Route(ws.GET("/").To(h.QueryEndpoints).
		Doc("查询服务功能列表").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(endpoint.QueryEndpointRequest{}).
		Writes(endpoint.EndpointSet{}).
		Returns(200, "OK", endpoint.NewEndpointSet()))

	ws.Route(ws.GET("/{id}").To(h.DescribeEndpoint).
		Doc("查询服务功能详情").
		Param(ws.PathParameter("id", "identifier of the service").DataType("integer").DefaultValue("1")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(endpoint.Endpoint{}).
		Returns(200, "OK", endpoint.Endpoint{}).
		Returns(404, "Not Found", nil))
}
