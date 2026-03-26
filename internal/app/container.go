package app

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"mytro-backend-content/internal/app/repository"
	"mytro-backend-content/internal/app/service"
	"mytro-backend-content/internal/infrastructure/config"
	"mytro-backend-content/internal/infrastructure/grpc/pb"

	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	Config *config.Config
	Logger *zap.Logger
	GRPC   *pb.ContentStorageClient
	DB     *gorm.DB

	// TODO: implement transaction manager
	// TxManager    *repository.TxManager
	Services     *AppServices
	repositories *AppRepositories
}

type AppServices struct {
	ContentService service.ContentService
}

type AppRepositories struct {
	ContentRepository repository.ContentRepository
	FileRepository    repository.FileRepository
}

func NewApp(config *config.Config, db *gorm.DB, logger *zap.Logger, grpcClient *pb.ContentStorageClient) *App {
	app := &App{
		Config: config,
		Logger: logger,
		GRPC:   grpcClient,
		DB:     db,

		// TxManager:    repository.NewTxManager(db),
		Services:     &AppServices{},
		repositories: &AppRepositories{},
	}

	app.createRepositories()
	app.createServices()

	return app
}

func (app *App) createRepositories() {
	app.repositories.ContentRepository = repository.NewContentRepository(app.Logger, app.DB)
	app.repositories.FileRepository = repository.NewFileRepository(*app.GRPC)
}

func (app *App) createServices() {
	app.Services.ContentService = service.NewContentService(app.repositories.ContentRepository, app.repositories.FileRepository)
}

func (app *App) GetPublicKey() (*rsa.PublicKey, error) {
	publicKeyBlock, _ := pem.Decode(([]byte)(app.Config.Keys.PublicKey()))

	if publicKeyBlock == nil {
		app.Logger.Fatal("failed to parse PEM block containing the key")
	}

	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		logger.Fatal("failed to parse public key: " + err.Error())
	}

	return publicKey.(*rsa.PublicKey), nil
}
