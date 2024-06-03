// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"gerrit.o-ran-sc.org/r/ric-plt/appmgr/pkg/restapi/operations"
	"gerrit.o-ran-sc.org/r/ric-plt/appmgr/pkg/restapi/operations/health"
	"gerrit.o-ran-sc.org/r/ric-plt/appmgr/pkg/restapi/operations/xapp"
)

//go:generate swagger generate server --target ../../pkg --name AppManager --spec ../../api/appmgr_rest_api.yaml --exclude-main

func configureFlags(api *operations.AppManagerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.AppManagerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.XappGetConfigElementHandler == nil {
		api.XappGetConfigElementHandler = xapp.GetConfigElementHandlerFunc(func(params xapp.GetConfigElementParams) middleware.Responder {
			return middleware.NotImplemented("operation xapp.GetConfigElement has not yet been implemented")
		})
	}
	if api.XappModifyXappConfigHandler == nil {
		api.XappModifyXappConfigHandler = xapp.ModifyXappConfigHandlerFunc(func(params xapp.ModifyXappConfigParams) middleware.Responder {
			return middleware.NotImplemented("operation xapp.ModifyXappConfig has not yet been implemented")
		})
	}
	if api.AddSubscriptionHandler == nil {
		api.AddSubscriptionHandler = operations.AddSubscriptionHandlerFunc(func(params operations.AddSubscriptionParams) middleware.Responder {
			return middleware.NotImplemented("operation .AddSubscription has not yet been implemented")
		})
	}
	if api.DeleteSubscriptionHandler == nil {
		api.DeleteSubscriptionHandler = operations.DeleteSubscriptionHandlerFunc(func(params operations.DeleteSubscriptionParams) middleware.Responder {
			return middleware.NotImplemented("operation .DeleteSubscription has not yet been implemented")
		})
	}
	if api.XappDeployXappHandler == nil {
		api.XappDeployXappHandler = xapp.DeployXappHandlerFunc(func(params xapp.DeployXappParams) middleware.Responder {
			return middleware.NotImplemented("operation xapp.DeployXapp has not yet been implemented")
		})
	}
	if api.DeregisterXappHandler == nil {
		api.DeregisterXappHandler = operations.DeregisterXappHandlerFunc(func(params operations.DeregisterXappParams) middleware.Responder {
			return middleware.NotImplemented("operation .DeregisterXapp has not yet been implemented")
		})
	}
	if api.XappGetAllXappConfigHandler == nil {
		api.XappGetAllXappConfigHandler = xapp.GetAllXappConfigHandlerFunc(func(params xapp.GetAllXappConfigParams) middleware.Responder {
			return middleware.NotImplemented("operation xapp.GetAllXappConfig has not yet been implemented")
		})
	}
	if api.XappGetAllXappsHandler == nil {
		api.XappGetAllXappsHandler = xapp.GetAllXappsHandlerFunc(func(params xapp.GetAllXappsParams) middleware.Responder {
			return middleware.NotImplemented("operation xapp.GetAllXapps has not yet been implemented")
		})
	}
	if api.HealthGetHealthAliveHandler == nil {
		api.HealthGetHealthAliveHandler = health.GetHealthAliveHandlerFunc(func(params health.GetHealthAliveParams) middleware.Responder {
			return middleware.NotImplemented("operation health.GetHealthAlive has not yet been implemented")
		})
	}
	if api.HealthGetHealthReadyHandler == nil {
		api.HealthGetHealthReadyHandler = health.GetHealthReadyHandlerFunc(func(params health.GetHealthReadyParams) middleware.Responder {
			return middleware.NotImplemented("operation health.GetHealthReady has not yet been implemented")
		})
	}
	if api.GetSubscriptionByIDHandler == nil {
		api.GetSubscriptionByIDHandler = operations.GetSubscriptionByIDHandlerFunc(func(params operations.GetSubscriptionByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation .GetSubscriptionByID has not yet been implemented")
		})
	}
	if api.GetSubscriptionsHandler == nil {
		api.GetSubscriptionsHandler = operations.GetSubscriptionsHandlerFunc(func(params operations.GetSubscriptionsParams) middleware.Responder {
			return middleware.NotImplemented("operation .GetSubscriptions has not yet been implemented")
		})
	}
	if api.XappGetXappByNameHandler == nil {
		api.XappGetXappByNameHandler = xapp.GetXappByNameHandlerFunc(func(params xapp.GetXappByNameParams) middleware.Responder {
			return middleware.NotImplemented("operation xapp.GetXappByName has not yet been implemented")
		})
	}
	if api.XappGetXappInstanceByNameHandler == nil {
		api.XappGetXappInstanceByNameHandler = xapp.GetXappInstanceByNameHandlerFunc(func(params xapp.GetXappInstanceByNameParams) middleware.Responder {
			return middleware.NotImplemented("operation xapp.GetXappInstanceByName has not yet been implemented")
		})
	}
	if api.XappListAllXappsHandler == nil {
		api.XappListAllXappsHandler = xapp.ListAllXappsHandlerFunc(func(params xapp.ListAllXappsParams) middleware.Responder {
			return middleware.NotImplemented("operation xapp.ListAllXapps has not yet been implemented")
		})
	}
	if api.ModifySubscriptionHandler == nil {
		api.ModifySubscriptionHandler = operations.ModifySubscriptionHandlerFunc(func(params operations.ModifySubscriptionParams) middleware.Responder {
			return middleware.NotImplemented("operation .ModifySubscription has not yet been implemented")
		})
	}
	if api.RegisterXappHandler == nil {
		api.RegisterXappHandler = operations.RegisterXappHandlerFunc(func(params operations.RegisterXappParams) middleware.Responder {
			return middleware.NotImplemented("operation .RegisterXapp has not yet been implemented")
		})
	}
	if api.XappUndeployXappHandler == nil {
		api.XappUndeployXappHandler = xapp.UndeployXappHandlerFunc(func(params xapp.UndeployXappParams) middleware.Responder {
			return middleware.NotImplemented("operation xapp.UndeployXapp has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
