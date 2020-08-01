package configs

import (
	"os"
)

/*_ConfigsApp export configurations app */
func _ConfigsApp() {
	/*Name of the application */
	subscribe("APP_NAME", os.Getenv("APP_NAME"), "Twittor")
	/*Application port */
	subscribe("PORT", os.Getenv("PORT"), "8080")
	/*Application secret */
	subscribe("APP_SECRET", os.Getenv("APP_SECRET"), "123456")
}
