package main

import (
	"time"
	"runtime"
	"github.com/tarm/serial"
)

func main() {
	// read serial
	
	// update reading to user interface
}


type EasyPowerData struct {
	InputVoltage 	float64
	InputCurrent 	float64
	InputPower 	 	float64
	OutputVoltage   float64
	OutputCurrent   float64
	OutputPower 	float64
	IntakeAirTemp 	float64
	OuttakeAirTemp 	float64
	FanSpeed			int

}

func OpenSerial(serialPortName string) (*serial.Port, error) {
	c := &serial.Config{Name: serialPortName, Baud: 115200, ReadTimeout: time.Second * 5}
    s, err := serial.OpenPort(c)
    if err != nil {
        return nil, error
    }
    return s, nil
}

func ReadSerial(*serial.Port) {

}



func openBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}