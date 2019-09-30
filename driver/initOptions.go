package driver

import (
	"fmt"
	"github.com/tebeka/selenium"
	"net"
	"runtime"
)

type DriverOptions struct {
	port int
	browserName string
	opts []selenium.ServiceOption
	binpath string
	caps selenium.Capabilities
	Service *selenium.Service
}

func DefaultChromeDrvOpt()*DriverOptions{
	o := &DriverOptions{}
	o.port, _ = GetUnUsedPort()
	if o.port == 0 {
		panic("Get unused port failed !")
	}
	o.browserName = "chrome"
	o.binpath = getChromeDrvPath()
	o.opts = []selenium.ServiceOption{}


	service, err := selenium.NewChromeDriverService(o.binpath,o.port,o.opts...)
	if err != nil {
		panic(err)
	}
	o.Service = service
	o.caps = selenium.Capabilities{"browserName":o.browserName}
	return o
}



func (options *DriverOptions)NewDriver()(Driver,error){
	urlprefix := fmt.Sprintf("http://localhost:%d/wd/hub",options.port)
	d ,err := selenium.NewRemote(options.caps, urlprefix)
	if err != nil {
		return nil,err
	}
	return d,nil
}


func getChromeDrvPath()string{
	switch runtime.GOOS{
	case "windows":
		return "./chromedriver.exe"
	case "linux":
		return "./chromedriver"
	case "darwin":
		return "./chromedriver"
	default:
		return "./chromedriver"
	}
}

func GetUnUsedPort() (int,error) {
	addr, err :=  net.ResolveTCPAddr("tcp","127.0.0.1:")
	if err != nil {
		return 0, err
	}
	l ,err :=  net.ListenTCP("tcp",addr)
	if err != nil {
		return 0,err
	}
	port := l.Addr().(*net.TCPAddr).Port
	if err := l.Close(); err != nil {
		return 0,err
	}
	return port, nil
}