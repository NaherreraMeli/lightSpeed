package lightSpeed

import "fmt"

var lightSpeed int = 300000

func SayLightSpeed() string {
	return fmt.Sprintf("The light speed is a constant wich value is: %v Km/s", lightSpeed)
}
