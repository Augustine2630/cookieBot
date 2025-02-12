package bootstrap

type App struct {
	Config Config
}

func NewApp() *App {
	return &App{
		Config: *NewConfig(),
	}
}
