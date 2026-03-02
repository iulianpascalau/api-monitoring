package factory

import (
	"github.com/iulianpascalau/api-monitoring/services/aggregation/api"
	"github.com/iulianpascalau/api-monitoring/services/aggregation/config"
	"github.com/iulianpascalau/api-monitoring/services/aggregation/storage"
)

//var log = logger.GetOrCreate("factory")

type componentsHandler struct {
	store  api.Storage
	server Server
	// alarmService *alarm.AlarmService
	// cancelFunc   context.CancelFunc
}

// NewComponentsHandler creates a new components handler
func NewComponentsHandler(
	sqlitePath string,
	serviceKeyApi string,
	authUsername string,
	authPassword string,
	cfg config.Config,
) (*componentsHandler, error) {
	store, err := storage.NewSQLiteStorage(sqlitePath, cfg.RetentionSeconds)
	if err != nil {
		return nil, err
	}

	serverArgs := api.ArgsWebServer{
		ServiceKeyApi:  serviceKeyApi,
		AuthUsername:   authUsername,
		AuthPassword:   authPassword,
		ListenAddress:  cfg.ListenAddress,
		StaticDir:      cfg.StaticDir,
		Storage:        store,
		GeneralHandler: api.CORSMiddleware,
	}

	server, err := api.NewServer(serverArgs)
	if err != nil {
		return nil, err
	}

	// TODO: fix this
	//// Initialize Notifiers
	//var activeNotifiers []alarm.Notifier
	//
	//logNotifier, err := notifiers.NewLogNotifier(log)
	//if err == nil {
	//	activeNotifiers = append(activeNotifiers, logNotifier)
	//}
	//
	//alarmService, err := alarm.NewAlarmService(store, activeNotifiers)
	//if err != nil {
	//	return nil, err
	//}
	//
	//_, cancel := context.WithCancel(context.Background())

	return &componentsHandler{
		store:  store,
		server: server,
		//alarmService: alarmService,
		//cancelFunc:   cancel,
	}, nil
}

// GetStore returns the storage component
func (ch *componentsHandler) GetStore() api.Storage {
	return ch.store
}

// GetServer returns the server component
func (ch *componentsHandler) GetServer() Server {
	return ch.server
}

// Start starts the inner components
func (ch *componentsHandler) Start() {
	ch.server.Start()
	//if ch.alarmService != nil && ch.cancelFunc != nil {
	//	ch.alarmService.Start(context.Background())
	//}
}

// Close closes the inner components
func (ch *componentsHandler) Close() {
	//if ch.cancelFunc != nil {
	//	ch.cancelFunc()
	//}
	//if ch.alarmService != nil {
	//	_ = ch.alarmService.Close()
	//}
	_ = ch.server.Close()
	_ = ch.store.Close()
}
