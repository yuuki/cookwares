package main

import (
	"log"
	"os"

	"github.com/yuuki/cookwares/pkg/resource/aptpackage"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	log.Println("Starting to cook localhost...")
	if err := run(os.Args); err != nil {
		log.Println(err)
		os.Exit(2)
	}
}

func run(args []string) error {
	var err error

	if err = aptpackage.Install("nginx", ""); err != nil {
		return err
	}

	// vars := struct {
	// 	LogFile string
	// }{
	// 	"/var/log/nginx/nginx.conf",
	// }
	// err = file.Template(&Template{
	// 	Dest:  "/etc/nginx/nginx.conf",
	// 	Src:   "templates/nginx.conf.tmpl",
	// 	Mode:  "0644",
	// 	Owner: "root",
	// 	Group: "root",
	// 	Vars:  vars,
	// 	Notify: func() error {
	// 		if err = systemd.Enable("nginx"); err != nil {
	// 			return err
	// 		}
	// 		return systemd.Restart("nginx")
	// 	},
	// })
	// if err != nil {
	// 	return err
	// }
	//
	// if err = systemd.Enable("nginx"); err != nil {
	// 	return err
	// }
	// if err = systemd.Start("nginx"); err != nil {
	// 	return err
	// }

	return nil
}
